package utils

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"net/http"
	"solity/schemas"
	"solity/utils"
	"solity/utils/ethereum/node"
	evmUtils "solity/utils/evm"
	"solity/utils/evm/evmStructs"
	kafkaUtils "solity/utils/kafka"
	"solity/utils/logger"
	"strings"
)

func InitilializeEnvironment() map[string]string {
	envNames := []string{"KAFKA_URI", "KAFKA_LISTEN_CHANNEL", "KAFKA_OUTPUT_CHANNEL", "DUNE_KEY"}
	status, missingName, ENV := utils.InitializeENV(envNames, "eigenlayerTracker.env")
	if !status {
		logger.LogE(missingName + " was not initialized... Aborting!")
	}

	return ENV
}

func GetDuneAVSMetadata(envMap map[string]string, avsAddress string) structs.Response {
	url := "https://api.dune.com/api/v1/eigenlayer/operator-statsfilters=avs_contract_address%20%3D%20" + strings.ToLower(avsAddress)
	apiKey := envMap["DUNE_KEY"]

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.LogE("Failed to create request: %v", err)
	}

	// Add the API key header
	req.Header.Add("X-DUNE-API-KEY", apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.LogE("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogE("Failed to read response body: %v", err)
	}

	// Parse the JSON response
	var stats structs.Response
	err = json.Unmarshal(body, &stats)
	if err != nil {
		logger.LogE("Failed to parse JSON response: %v", err)
	}

	// Print the response
	fmt.Printf("%+v\n", stats)

	return stats
}

func GetDuneOperatorMetadata(envMap map[string]string, operatorAddress string) structs.ResponseOp {

	url := "https://api.dune.com/api/v1/eigenlayer/operator-stats?filters=operator_contract_address%20%3D%20" + strings.ToLower(operatorAddress)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.LogE("Failed to create request: %v", err)
	}

	req.Header.Add("X-DUNE-API-KEY", envMap["DUNE_KEY"])

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.LogE("Failed to send request: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// Parse the JSON response
	var stats structs.ResponseOp
	err = json.Unmarshal(body, &stats)
	if err != nil {
		logger.LogE("Failed to parse JSON response: %v", err)
	}

	// Print the response
	fmt.Printf("%+v\n", stats)

	return stats
}

func WriteToSmartContract(envMap map[string]string, payload structs.EigenlayerPayload) {
	client, _ := ethclient.Dial("https://eth.dev-solity.net/rpc")
	newRegistery, _ := registery.NewRegistery(common.HexToAddress(""), client)
	privateKey, _, fromAddress, _ := node.GenerateKeypairFromPrivateKeyHex(envMap["PRV_KEY"])
	txOpts, _ := node.BuildTransactionOptions(client, fromAddress, privateKey, 1090000)
	_, _ = newRegistery.RegisterEvent(txOpts, payload.AvsName, payload.OperatorName, payload.AvsAddress, payload.OperatorAddress)
}

func CheckAVSMetadata(message schemas.SolityETHCompleteTransactionMessage,
	producer *kafka.Producer, outputChannel *string, eventSignature *evmStructs.SignatureKeeper, envMap map[string]string) {
	// This service expects types.Receipt format
	rcptInfo := new(types.Receipt)
	txInfo := new(types.Transaction)
	fMErr := rcptInfo.UnmarshalJSON(message.ReceiptData)
	err := txInfo.UnmarshalJSON(message.TransactionData)

	if err != nil {
		logger.LogW("Error while un-marshalling the txInfo data: ", err)
		return
	}

	// If the message is malformed skip the message
	if fMErr != nil {
		logger.LogW("Error while un-marshalling the rcpt data: ", fMErr)
		return
	}

	// Check the status of the transaction, if it has failed do not include it
	if rcptInfo.Status != 1 {
		logger.LogW("The tx [%s] has failed, skipping the processing for this tx!", rcptInfo.TxHash.Hex())
		return
	}

	for _, eventLog := range rcptInfo.Logs {
		decodedLog, err := evmUtils.DecodeLog(eventLog, *eventSignature)
		if err != nil {
			logger.LogW(err)
		} else {
			operatorAddress, err := decodedLog.DecodedIndexedData[0].AsAddress()
			if err != nil {
				logger.LogW(err)
			} else {
				avsAddress := strings.ToLower(eventLog.Address.String())
				operatorAddressAsStr := strings.ToLower(operatorAddress.Hex())
				resOp := GetDuneOperatorMetadata(envMap, operatorAddressAsStr)
				operatorName := resOp.Result.Rows[0].OperatorName
				resAvs := GetDuneAVSMetadata(envMap, avsAddress)
				avsName := resAvs.Result.Rows[0].AVSName
				logger.LogS(operatorName + " operator is registered to " + avsName + " AVS")
				payload := structs.EigenlayerPayload{
					AvsName:         avsName,
					AvsAddress:      common.HexToAddress(avsAddress),
					OperatorAddress: operatorAddress,
					OperatorName:    operatorName,
				}
				WriteToSmartContract(envMap, payload)
				kafkaUtils.ConvertAndSendSolityMessageSingleClient(payload,
					"",
					"",
					"TRCK-Eigenlayer",
					producer,
					outputChannel,
					"",
					"")
			}
		}
	}
}

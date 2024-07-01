package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"net/http"
	"regexp"
	contracts "solity/contracts/go"
	"solity/schemas"
	"solity/utils/logger"
	mathUtils "solity/utils/math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
GetBlock takes two arguments string arguments. Url is the ethereum node endpoint and identifier is the block identifier.
Gets the specified blocks data through a http get request. If unsuccessful it returns false and empty BlockInfo object
*/
func GetBlock(url string, identifier string) (bool, schemas.EthereumBlockInfo) {
	// Add current endpoint to the base url
	fullURL := url + "/eth/v2/beacon/blocks/" + identifier

	// Create a client object
	client := &http.Client{}

	// Fetch the data
	resp, err := client.Get(fullURL)

	if err != nil {
		log.Println("Error while fetching current block data from the given url: ")
		log.Println(err)
		return false, schemas.EthereumBlockInfo{}
	}

	// When the reading is done close the message body
	defer resp.Body.Close()

	//Parse the received message
	body, err := io.ReadAll(resp.Body)

	receivedData := schemas.EthereumBlockInfo{}
	rDErr := json.Unmarshal(body, &receivedData)

	if rDErr != nil {
		log.Println("Error at unpacking the Ethereum Block Data")
		log.Println(rDErr)

		return false, schemas.EthereumBlockInfo{}
	}

	return true, receivedData
}

/*
GetBlockViaLib takes two argument and returns a block pointer using geth. Url is the ethereum node endpoint and blockNumber is the number of the block that we want to get the data for
If unsuccessful it returns false and empty BlockInfo object. This function initializes a new client per usage in order to be tread safe. This increases the memory consumption a little bit
*/
func GetBlockViaLib(url string, blockNumber int) (*types.Block, bool) {

	// Connect to the ethereum node
	client, dialErr := ethclient.Dial(url)

	if dialErr != nil {
		// If there was an error during the connection to the node exit
		log.Println("Connection to the node cannot be established: ", dialErr)
		return nil, false
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the block data
	currentBlock, blockErr := client.BlockByNumber(ctx, mathUtils.NewBigInt(blockNumber))

	if blockErr != nil {
		// If there was an error while getting the block data
		log.Println("Block data cannot be retrieved from the node: ", blockErr)
		return nil, false
	}

	return currentBlock, true
}

/*
GetBlockViaLib64 takes two argument and returns a block pointer using geth. Url is the ethereum node endpoint and blockNumber is the number of the block that we want to get the data for
If unsuccessful it returns false and empty BlockInfo object. This function initializes a new client per usage in order to be tread safe. This increases the memory consumption a little bit
*/
func GetBlockViaLib64(url string, blockNumber int64) (*types.Block, bool) {

	// Connect to the ethereum node
	client, dialErr := ethclient.Dial(url)

	if dialErr != nil {
		// If there was an error during the connection to the node exit
		log.Println("Connection to the node cannot be established: ", dialErr)
		return nil, false
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the block data
	currentBlock, blockErr := client.BlockByNumber(ctx, mathUtils.NewBigInt(blockNumber))

	if blockErr != nil {
		// If there was an error while getting the block data
		log.Println("Block data cannot be retrieved from the node: ", blockErr)
		return nil, false
	}

	return currentBlock, true
}

/*
GetTxRcptViaLib takes two argument and returns a *types.Receipt using geth. Url is the ethereum node endpoint and txHash is the hash of the transaction that we want to get the data for
If unsuccessful it returns false and nil *types.Receipt object. This function initializes a new client per usage in order to be tread safe. This increases the memory consumption a little bit
*/
func GetTxRcptViaLib(url string, txHash common.Hash) (*types.Receipt, bool) {

	// Connect to the ethereum node
	client, dialErr := ethclient.Dial(url)

	if dialErr != nil {
		// If there was an error during the connection to the node exit
		log.Println("Connection to the node cannot be established: ", dialErr)
		return nil, false
	}

	defer client.Close()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the block data
	theRcpt, txRecptErr := client.TransactionReceipt(ctx, txHash)

	if txRecptErr != nil {
		// If there was an error while getting the block data
		log.Println("The transaction receipt for the tx", txHash, "cannot be retrieved from the node: ", txRecptErr)
		return nil, false
	}

	return theRcpt, true
}

/*
CheckType checks if an address belongs to a smart contract or a wallet
Returns type of the contract in string: "C" (for contract), "W" (for wallet)
*/
func CheckType(address string, endpoint string) (addressType string) {

	// Connect to the ethereum node
	client, dialErr := ethclient.Dial(endpoint)

	if dialErr != nil {
		// If there was an error during the connection to the node exit
		log.Fatalln("Connection to the node cannot be established: ", dialErr)
	}

	// Get the byte code belonging to the address
	bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(address), nil)

	if err != nil {
		// If there was an error during the bytecode retrieval
		log.Fatalln("Error while getting the byte code belonging to the", address, ":", err)
	}

	isContract := len(bytecode) > 0

	if isContract {
		return "C"
	}

	return "W"
}

/*
GenerateKeypairFromPrivateKeyHex takes input a private key string (both with or without 0x in the beginning is okay) and returns
ecdsa Private key, Public key, ethereum address belonging to that public key, and error (if any present)
*/
func GenerateKeypairFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, error) {

	log.Println("Generating the keypair...")

	// If hex string has "0x" at the start discard it
	if privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Convert hex key to a private key object
	privateKey, privateKeyErr := crypto.HexToECDSA(privateKeyHex)

	if privateKeyErr != nil {
		return nil, nil, common.Address{}, privateKeyErr
	}

	// Generate the public key using the private key
	publicKey := privateKey.Public()

	// Cast crypto.Publickey object to the *ecdsa.PublicKey object
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}

	// Convert publickey to a address
	pubkeyAsAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, pubkeyAsAddress, nil
}

func BuildTransactionOptions(client *ethclient.Client, fromAddress common.Address, prvKey *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {

	// Retrieve the chainID
	chainID, cIDErr := client.ChainID(context.Background())

	if cIDErr != nil {
		return nil, cIDErr
	}

	// Retrieve Nonce
	nonce, nErr := client.PendingNonceAt(context.Background(), fromAddress)

	if nErr != nil {
		return nil, nErr
	}

	// Retrieve suggested gas price
	gasPrice, gErr := client.SuggestGasPrice(context.Background())

	if gErr != nil {
		return nil, gErr
	}

	logger.LogI("Suggested gas price: ", gasPrice)

	// Create empty options object
	txOptions, txOptErr := bind.NewKeyedTransactorWithChainID(prvKey, chainID)

	if txOptErr != nil {
		return nil, txOptErr
	}

	txOptions.Nonce = big.NewInt(int64(nonce))
	txOptions.Value = big.NewInt(0) // in wei
	txOptions.GasLimit = gasLimit   // in units
	txOptions.GasPrice = gasPrice.Sub(gasPrice, big.NewInt(2000000000))

	return txOptions, nil
}
func isContractAddress(addr string, client *ethclient.Client) bool {
	if len(addr) == 0 {
		log.Fatal("feedAddress is empty.")
	}

	// Ensure it is an Ethereum address: 0x followed by 40 hexadecimal characters.
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addr) {
		log.Fatalf("address %s non valid\n", addr)
	}

	// Ensure it is a contract address.

	//Temp node for fetching the is contract data
	tmpClient, initErr := ethclient.Dial("https://rpc.ankr.com/eth/ad381396eaadb2ef18b7085c4bcef169aac71404a5359bd20421ceb427cdbd4e")

	if initErr != nil {
		logger.LogW("error while initializing the temp node: ", initErr)
		return false
	}

	address := common.HexToAddress(addr)

	//Sleep in order to not f*ck the free node rate limit
	time.Sleep(time.Second * 1)
	bytecode, err := tmpClient.CodeAt(context.Background(), address, nil)

	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func getLatestCHPrice(addr string, client *ethclient.Client) float64 {
	ok := isContractAddress(addr, client)
	if !ok {
		logger.LogWf("Address [%s] is not a contract address!", addr)
		return -1.0
	}
	chainlinkPriceFeedProxyAddress := common.HexToAddress(addr)
	chainlinkPriceFeedProxy, err := contracts.NewAggregatorV3Interface(chainlinkPriceFeedProxyAddress, client)
	if err != nil {
		logger.LogWf("Error while creating AggregatorV3 of Chainlink: [%s]", err.Error())
		return -1.0
	}
	roundData, err := chainlinkPriceFeedProxy.LatestRoundData(&bind.CallOpts{})
	if err != nil {
		logger.LogWf("Error while fetching latest round data from Chainlink: [%s]", err.Error())
		return -1.0
	}
	decimals, err := chainlinkPriceFeedProxy.Decimals(&bind.CallOpts{})
	if err != nil {
		logger.LogWf("Error while fetching decimals data from Chainlink: [%s]", err.Error())
		return -1.0
	}
	divisor := mathUtils.Pow(10, int(decimals))
	formatted_ans, _ := mathUtils.DivBigFloat(roundData.Answer, divisor).Float64()
	// Get the latest price of CH
	return formatted_ans
}
func GetChainlinkPrice(token1, token2, base string, client *ethclient.Client) float64 {
	const feed_url_src = "https://reference-data-directory.vercel.app/feeds-mainnet.json"
	const chain_eth = "ethereum"
	const asset_type_focus_crypto = "crypto"
	const asset_type_focus_forex = "forex"

	const asset_verification_status = "verified"

	token1pair := token1 + base
	token2pair := token2 + base
	token1price := 0.0
	token2price := 0.0

	usd_pair := false
	concatenated_pair := strings.ToUpper(token1) + " / " + strings.ToUpper(token2)
	if token1 == "USD" || token1 == "USDT" || token2 == "USD" || token2 == "USDT" {
		usd_pair = true
	}
	resp, err := http.Get(feed_url_src)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		logger.LogW("Unexpected response from chainlink ", resp.StatusCode, " ", resp.Status, " \nBody: \n", string(body))
		logger.LogE("Could not get the feeds from chainlink ", err)

	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogE("Could not read the feeds from chainlink ", err)
	}
	// str_body := string(body)
	var proxies []schemas.ChainlinkResponse
	if err := json.Unmarshal(body, &proxies); err != nil {
		logger.LogE(err)
	}
	for _, proxy := range proxies {
		isCrypto := strings.ToLower(string(proxy.FeedType)) == asset_type_focus_crypto
		isForex := strings.ToLower(string(proxy.FeedType)) == asset_type_focus_forex
		isVerified := strings.ToLower(string(proxy.FeedCategory)) == asset_verification_status
		isHidden := proxy.Docs.Hidden
		isFocused := isCrypto
		if strings.Contains(strings.ToUpper(token1), "EUR") {
			isFocused = isForex
		}
		if isFocused && isVerified && !isHidden {
			proxy_addr := proxy.ProxyAddress
			pair := proxy.Name
			if usd_pair && concatenated_pair == pair {
				return getLatestCHPrice(proxy_addr, client)
			} else if token1pair == pair {
				token1price = getLatestCHPrice(proxy_addr, client)
				if token2 == "" {
					if token1price == -1.0 {
						logger.LogWf("Chainlink does not provide price data for [%s]!", pair)
						logger.LogW("Returning -1!")
					} else {
						log.Println("Token pair price found for ", pair, " price is ", token1price)
					}
					return token1price
				}
			} else if token2pair == pair {
				token2price = getLatestCHPrice(proxy_addr, client)
			}
		}
	}

	if token1price != 0.0 && token2price != 0.0 {
		return token1price / token2price
	} else {
		log.Printf("%s/%s Token pair price not found for base %s", token1, token2, base)
		return -1.0
	}
}

/*
GetTxSender extracts the From field from the given transaction
*/
func GetTxSender(transaction *types.Transaction) (sender common.Address, err error) {

	sender, err = types.Sender(types.NewLondonSigner(transaction.ChainId()), transaction)

	if err != nil {
		return
	}

	return
}

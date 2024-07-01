package main

import (
	"eigenlayer_hack/utils"
	"encoding/json"
	"solity/schemas"
	"solity/utils/evm/evmStructs"
	kafkaUtils "solity/utils/kafka"
	"solity/utils/logger"
	"strings"
)

func main() {
	envMap := utils.InitilializeEnvironment()

	listenChannels := strings.Split(envMap["KAFKA_LISTEN_CHANNEL"], ":")
	logger.LogI("Currently listening ", listenChannels)

	outputChannels := strings.Split(envMap["KAFKA_OUTPUT_CHANNEL"], ":")
	logger.LogI("Kafka out channel ", outputChannels)
	if len(outputChannels) != len(listenChannels) {
		logger.LogE("Kafka listen channels and output channels length miss-match")
	}
	//Listen Channels
	channelMapping := map[string]string{}

	for i, listenChannel := range listenChannels {
		channelMapping[listenChannel] = outputChannels[i]
	}
	c, p, err := kafkaUtils.InitializeKafkaConsumerAndProducer(envMap["KAFKA_URI"], listenChannels, "centralisedEx",
		nil, nil)
	if err != nil {
		logger.LogE("Error while initializing kafka: ", err)
	}

	defer c.Close()
	defer p.Close()
	operatorRegisterSignature := evmStructs.NewSignatureKeeper("OperatorSubscribed (indexed address operator, indexed uint32 chainID)")

	// Start message reading loop
	for {
		// Read the message
		msg, mErr := c.ReadMessage(-1)

		if mErr == nil {
			logger.LogIf("Topic is: %s", *msg.TopicPartition.Topic)

			// This service expects schemas.SolityETHCompleteTransactionMessage format
			receivedMessage := schemas.SolityETHCompleteTransactionMessage{}
			rMErr := json.Unmarshal(msg.Value, &receivedMessage)
			if rMErr != nil {
				logger.LogW("Error while unmarshalling the message object : ", rMErr)
				// Since this message cannot be unmarshalled, we skip the execution for this message
				continue
			}

			theOutputChannel, isOk := channelMapping[*msg.TopicPartition.Topic]

			if isOk {
				// Process the message
				go utils.CheckAVSMetadata(receivedMessage,
					p,
					&theOutputChannel,
					&operatorRegisterSignature,
					envMap,
				)
			}

		} else {
			// The client will automatically try to recover from all errors.
			logger.LogWf("Consumer error: %v (%v)\n", mErr, msg)
		}
	}
}

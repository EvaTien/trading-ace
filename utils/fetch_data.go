package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"
	"trading-ace/config"
	"trading-ace/db"
)

type EtherscanResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []Log  `json:"result"`
}

type Log struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

var SwapData struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
}

func formatLogs(apiResponse string) EtherscanResponse {
	// Unmarshal the JSON response into Go struct
	var response EtherscanResponse
	err := json.Unmarshal([]byte(apiResponse), &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func getTimeStampUTC(hexTime string) time.Time {
	// Convert timestamp from hex to decimal and then to UTC time
	timestamp := new(big.Int)
	timestamp.SetString(hexTime[2:], 16) // remove "0x" prefix
	dateTime := time.Unix(timestamp.Int64(), 0).UTC()
	return dateTime
}

func getBlockNumber(BlockNumber string) string {
	blockNumber := new(big.Int)
	blockNumber.SetString(BlockNumber[2:], 16) // remove "0x" prefix
	blockNumberStr := blockNumber.String()
	return blockNumberStr
}

func decodeData(data string) {
	// Event ABI for the Swap event
	eventABI := `[{
		"anonymous": false,
		"inputs": [
			{"indexed": true, "internalType": "address", "name": "sender", "type": "address"},
			{"indexed": false, "internalType": "uint256", "name": "amount0In", "type": "uint256"},
			{"indexed": false, "internalType": "uint256", "name": "amount1In", "type": "uint256"},
			{"indexed": false, "internalType": "uint256", "name": "amount0Out", "type": "uint256"},
			{"indexed": false, "internalType": "uint256", "name": "amount1Out", "type": "uint256"},
			{"indexed": true, "internalType": "address", "name": "to", "type": "address"}
		],
		"name": "Swap",
		"type": "event"
	}]`

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(eventABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	logData := common.FromHex(data)

	// Decode the event data into the struct
	err = parsedABI.UnpackIntoInterface(&SwapData, "Swap", logData)
	if err != nil {
		log.Fatalf("Failed to decode data: %v", err)
	}
}

func insertData(logData Log) {
	existed := db.UserExisted(logData.Address)
	if !existed {
		db.CreateNewUser(logData.Address)
	}
}

func GetSwapTransactions() {
	baseURL := "https://api.etherscan.io/api"

	params := url.Values{}
	params.Add("module", "logs")
	params.Add("action", "getLogs")
	params.Add("address", config.Config.Server.SharePoolAddress)
	params.Add("topic0", config.Config.Server.TrackingHash)
	params.Add("apikey", config.Config.Server.ApiKey)
	params.Add("fromBlock", "21276410")
	params.Add("toBlock", "latest")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var response EtherscanResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to parse response body: %v", err)
	}

	for _, logData := range response.Result {
		insertData(logData)
	}
}

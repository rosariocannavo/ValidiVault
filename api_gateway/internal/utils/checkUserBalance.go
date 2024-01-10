package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/rosariocannavo/api_gateway/config"

	"github.com/rosariocannavo/api_gateway/internal/models"
	"github.com/rosariocannavo/api_gateway/internal/nats"
)

func CheckUserBalance(metamaskAddress string) string {
	//payload := strings.NewReader(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s", "latest"],"id":1}`, userForm.MetamaskAddress))
	payload := strings.NewReader(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBalance","params":["%s", "latest"],"id":1}`, metamaskAddress))

	// Sending the HTTP POST request to the Ganache endpoint
	resp, err := http.Post(config.GanacheURL, "application/json", payload)
	if err != nil {
		fmt.Println("Impossibile to contact the blockchain: check connection")

		message := fmt.Sprintf("Timestamp: %s | Handler: %s | Status: %d | Response: %s", time.Now().UTC().Format(time.RFC3339), "registration_handler/CheckUserBalance", "500", "message: Impossible to contact the blockchain")
		nats.NatsConnection.PublishMessage(message)

		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response models.BlockChainResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	balanceInt, success := new(big.Int).SetString(response.Result[2:], 16)
	if !success {
		log.Fatalf("Failed to convert balance to big.Int")
	}

	// Define a threshold balance in Wei (here, 1 Ether = 10^18 Wei)
	threshold := new(big.Int).SetUint64(1e18)

	// Compare the balance with the threshold
	if balanceInt.Cmp(threshold) >= 0 {
		return models.Admin
	} else {
		return models.NormalUser
	}

}

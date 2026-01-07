package squid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"net/http"
)

func SendSquidRouteRequest(client *http.Client, swap QuoteRequest) (string, string, string, string, error) {
	payload, err := json.Marshal(swap)
	if err != nil {
		return "", "", "", "", fmt.Errorf("json marshal squid swap data error %s", err.Error())
	}
	squidHost, squidRoute := config.GetSquidConfig()
	req, err := http.NewRequest("POST", squidHost+squidRoute, bytes.NewBuffer(payload))
	if err != nil {
		return "", "", "", "", fmt.Errorf("new squid swap request error %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("accept", "application/json")
	req.Header.Set("x-integrator-id", "mcga-2437f1e2-2e86-4acc-9de2-ba27fb9be6d7")
	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", "", fmt.Errorf("squid swap api request error %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", "", "", "", fmt.Errorf("squid swap api response status %d, body: %s", resp.StatusCode, string(body))
	}
	defer resp.Body.Close()

	var result QuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", "", "", fmt.Errorf("json decode squid swap api response error %s", err.Error())
	}

	return result.Route.TransactionRequest.Target, result.Route.TransactionRequest.Data, result.Route.TransactionRequest.Value, result.Route.TransactionRequest.GasLimit, nil
}

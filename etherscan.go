package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Etherscan struct {
	http *http.Client
}

func NewEtherscan(apiKey string) *Etherscan {
	if apiKey == "" {
		panic("Etherscan: apiKey is required")
	}
	return &Etherscan{
		http: &http.Client{
			Timeout:   10 * time.Second,
			Transport: wrapWithAPIKey(http.DefaultTransport, apiKey),
		},
	}
}

func (e *Etherscan) GetContractABI(ctx context.Context, contract common.Address) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.etherscan.io/api", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to init http request: %w", err)
	}
	query := req.URL.Query()
	query.Set("module", "contract")
	query.Set("action", "getabi")
	query.Set("address", contract.Hex())
	req.URL.RawQuery = query.Encode()

	res, err := e.http.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}
	defer res.Body.Close()

	var msg struct {
		Status  int    `json:"status,string"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}
	err = json.NewDecoder(res.Body).Decode(&msg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return []byte(msg.Result), nil
}

type transport struct {
	impl   http.RoundTripper
	apiKey string
}

func wrapWithAPIKey(rt http.RoundTripper, apiKey string) *transport {
	return &transport{
		impl:   rt,
		apiKey: apiKey,
	}
}

func (tr *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	q.Set("apikey", tr.apiKey)
	req.URL.RawQuery = q.Encode()
	return tr.impl.RoundTrip(req)
}

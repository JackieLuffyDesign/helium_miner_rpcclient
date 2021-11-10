package helium_miner_rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type p2p struct {
	endpoint string
}

type info struct {
	endpoint string
}

type Client struct {
	P2P  p2p
	Info info
}

// Create a new Helium Miner RPC client
func New(endpoint string) *Client {
	return &Client{
		P2P:  p2p{endpoint},
		Info: info{endpoint},
	}
}

func (net p2p) Status() (*P2PStatus, error) {
	var status P2PStatus
	if err := makeRequest(net.endpoint, "info_p2p_status", nil, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

func (info info) Height() (uint64, error) {
	var height MinerHeight
	if err := makeRequest(info.endpoint, "info_height", nil, &height); err != nil {
		return 0, err
	}
	return height.Height, nil
}

func (info info) InConsensus() (*InConsensus, error) {
	var ic InConsensus
	if err := makeRequest(info.endpoint, "info_height", nil, &ic); err != nil {
		return nil, err
	}
	return &ic, nil
}

func (info info) Name() (*string, error) {
	var name Name
	if err := makeRequest(info.endpoint, "info_name", nil, &name); err != nil {
		return nil, err
	}
	return &name.Name, nil
}

func (info info) BlockAge() (uint64, error) {
	var blockAge BlockAge
	if err := makeRequest(info.endpoint, "info_block_age", nil, &blockAge); err != nil {
		return 0, err
	}
	return blockAge.Age, nil
}

func (info info) Region() (*string, error) {
	var region Region
	if err := makeRequest(info.endpoint, "info_region", nil, &region); err != nil {
		return nil, err
	}
	return region.Region, nil
}

func (info info) Summary() (*Summary, error) {
	var summary Summary
	if err := makeRequest(info.endpoint, "info_summary", nil, &summary); err != nil {
		return nil, err
	}
	return &summary, nil
}

func (info info) Version() (*string, error) {
	var version Version
	if err := makeRequest(info.endpoint, "info_version", nil, &version); err != nil {
		return nil, err
	}
	return &version.Version, nil
}

func makeRequest(endpoint string, method string, payload interface{}, result interface{}) error {
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		ID:      rand.Intn(50000),
		Method:  method,
		Payload: &payload,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json-rpc")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
		}
	}

	var (
		dec   = json.NewDecoder(resp.Body)
		reply jsonRPCResponse
	)
	if err := dec.Decode(&reply); err != nil {
		return fmt.Errorf("unable to decode JSON RPC reply: %w", err)
	}

	if req.ID != reply.ID {
		fmt.Printf("ID: %d, ID: %d", req.ID, reply.ID)
		return InvalidResponseIDErr{}
	}

	if err := json.Unmarshal(reply.Result, result); err != nil {
		return fmt.Errorf("unable to decode reply body: %w", err)
	}

	return nil
}

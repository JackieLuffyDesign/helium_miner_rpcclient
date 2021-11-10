package helium_miner_rpcclient

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	StatusCode int
	Status     string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("JSONRPC reported error %d - %s", e.StatusCode, e.Status)
}

type InvalidResponseIDErr struct{}

func (e InvalidResponseIDErr) Error() string {
	return fmt.Sprintf("JSON RPC returned reply with an invalid ID")
}

type jsonRPCRequest struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Method  string       `json:"method"`
	Payload *interface{} `json:"omitempty"`
}

type jsonRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  json.RawMessage
}

type P2PStatus struct {
	Connected string
	Dialable  string
	Height    uint64
	NatType   string `json:"nat_type"`
}

type MinerHeight struct {
	Height uint64
}

type InConsensus struct {
	Epoch  uint64
	Height uint64
}

type Name struct {
	Name string
}

type BlockAge struct {
	Age uint64 `json:"block_age"`
}

type Region struct {
	Region *string
}

type MacAddress map[string]string

type Summary struct {
	BlockAge           uint64 `json:"block_age"`
	Epoch              uint64
	FirmwareVersion    string `json:"firmware_version"`
	GatewayDetails     string `json:"gateway_details"`
	Height             uint64
	MacAddresses       []MacAddress `json:"mac_addresses"`
	Name               string
	PeerBookEntryCount uint64 `json:"peer_book_entry_count"`
	SyncHeight         uint64 `json:"sync_height"`
	Uptime             uint64
	Version            string
}

type Version struct {
	Version string
}

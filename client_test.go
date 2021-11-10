package helium_miner_rpcclient_test

import (
	"context"
	"testing"

	"github.com/henet/helium_miner_rpcclient"
)

func TestP2PStatus(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	status, err := client.P2P.Status(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("status: %+v", status)
}

func TestInfoHeight(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	height, err := client.Info.Height(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("miner height: %d", height)
}

func TestInfoInConsensus(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	inConsensus, err := client.Info.InConsensus(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("epoch: %d height: %d", inConsensus.Epoch, inConsensus.Height)
}

func TestInfoName(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	name, err := client.Info.Name(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("name: %s", *name)
}

func TestInfoBlockAge(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	age, err := client.Info.BlockAge(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("age: %d", age)
}

func TestInfoRegion(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	region, err := client.Info.Region(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("region: %+v", region)
}

func TestInfoSummary(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	summary, err := client.Info.Summary(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("summary: %+v", summary)
}

func TestInfoVersion(t *testing.T) {
	client := helium_miner_rpcclient.New("http://localhost:4467")

	version, err := client.Info.Version(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("version: %s", *version)
}

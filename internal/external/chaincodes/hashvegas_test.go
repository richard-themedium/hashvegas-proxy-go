package chaincodes

import (
	"testing"
)

func getHashVegasContract() *hashVegasContract {
	return newHashVegasContractFabric("./connection-profile.json")
}

func Test_getFabricClient(t *testing.T) {
	contract := getHashVegasContract()
	_, err := contract.getFabricClient()
	if err != nil {
		t.Fatal(err)
	}

}

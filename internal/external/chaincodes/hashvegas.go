package chaincodes

import (
	"io/ioutil"

	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
	"github.com/the-medium-tech/platform-externals/log"
)

type ContractType int

const (
	FabricType ContractType = 0
	MdlType    ContractType = 1
)

type HashVegasContract interface {
	InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error)
	GetGameLastState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error)
	GetGameHistory(tableId string, gameSeq uint64) ([]*model_dto.HoldemTableGame, error)
	getDeckInitState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error)
	AskNewCards(tableId string, gameSeq uint64, dealSeq uint8, requestCardList []map[string]interface{}) (*model_dto.HoldemServicerResponse, error)
}

func NewHashVegasContract(contractType ContractType,
	connProfilePath string,
	mdlAddress string,
	mdlHvChannel string,
	mdlHvCcName string,
	MdlHvCcVersion string,
) HashVegasContract {
	if contractType == FabricType {
		return newHashVegasContractFabric(connProfilePath)
	} else if contractType == MdlType {
		return newHashVegasContractMdl(mdlAddress, mdlHvChannel, mdlHvCcName, MdlHvCcVersion)
	} else {
		panic("could not find matched implementation of the contract: " + fmt.Sprintf("%d", contractType))
	}
}

func newHashVegasContractFabric(connProfilePath string) *hashVegasContract {
	profileData, err := ioutil.ReadFile(connProfilePath)
	if err != nil {
		panic("could not read conn profile: " + err.Error())
	}
	return &hashVegasContract{
		connProfileData: string(profileData),
		tgtOrg:          "org1",
		tgtOrgUser:      "peer0.org1.example.com",
	}
}

type hashVegasContract struct {
	connProfileData string
	tgtOrg          string
	tgtOrgUser      string
	client          *resmgmt.Client
}

func (h *hashVegasContract) getFabricClient() (*resmgmt.Client, error) {

	// log.Info("profile data: ", h.connProfileData)
	sdk, err := fabsdk.New(config.FromRaw([]byte(h.connProfileData), "json"))
	if err != nil {
		log.Error("could not new fabsdk:", err.Error())
		return nil, err
	}

	ctx := sdk.Context(
		fabsdk.WithOrg("org1"),
	)

	orgClient, err := resmgmt.New(ctx)
	if err != nil {
		log.Error("Failed generating a resource management client:", err.Error())
		return nil, err
	}

	return orgClient, nil
}

// AskNewCards implements HashVegasContract
func (*hashVegasContract) AskNewCards(tableId string, gameSeq uint64, dealSeq uint8, requestCardList []map[string]interface{}) (*model_dto.HoldemServicerResponse, error) {
	panic("unimplemented")
}

// GetGameHistory implements HashVegasContract
func (*hashVegasContract) GetGameHistory(tableId string, gameSeq uint64) ([]*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// GetGameLastState implements HashVegasContract
func (*hashVegasContract) GetGameLastState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// InitGenesisLedgerForTable implements HashVegasContract
func (*hashVegasContract) InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

// getDeckInitState implements HashVegasContract
func (*hashVegasContract) getDeckInitState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	panic("unimplemented")
}

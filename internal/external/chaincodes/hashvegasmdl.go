package chaincodes

import (
	"crypto/aes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	model_dto "github.com/the-medium-tech/mdl-manager/hashvegas-proxy-go/internal/dto/model"
	"github.com/the-medium-tech/platform-externals/log"
)

type hashVegasContractMdl struct {
	addres        string
	channel       string
	chaincodeName string
	version       string
	queryBaseUrl  string
	invokeBaseUrl string
}

func newHashVegasContractMdl(address, channel, chaincodeName, version string) *hashVegasContractMdl {
	queryBaseUrl := fmt.Sprintf("%s/api/blockchain/channels/%s/chaincodes/hashvegas/query",
		address,
		channel,
	)
	invokeBaseUrl := fmt.Sprintf("%s/api/blockchain/channels/%s/chaincodes/hashvegas/invoke",
		address,
		channel,
	)

	return &hashVegasContractMdl{
		addres:        address,
		channel:       channel,
		queryBaseUrl:  queryBaseUrl,
		invokeBaseUrl: invokeBaseUrl,
	}
}

// Invoke
// AskNewCards implements HashVegasContract
func (h *hashVegasContractMdl) AskNewCards(tableId string, gameSeq uint64, dealSeq uint8, requestCardList []map[string]interface{}) (*model_dto.HoldemServicerResponse, error) {
	var resp struct {
		TransactionId string
		Result        *model_dto.HoldemServicerResponse
	}

	ordererName, peerId, err := h.getMdlEssentials(h.channel)
	if err != nil {
		return nil, err
	}

	bodyValues := h.genEssentialValues(ordererName, peerId)
	bodyValues.Add("function", "AskNewCards")
	bodyValues.Add("args", tableId)
	bodyValues.Add("args", fmt.Sprintf("%d", gameSeq))
	bodyValues.Add("args", fmt.Sprintf("%d", dealSeq))

	respData, err := h.sendAndRecvPost(h.invokeBaseUrl, bodyValues)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// Query
// GetGameHistory implements HashVegasContract
func (h *hashVegasContractMdl) GetGameHistory(tableId string, gameSeq uint64) ([]*model_dto.HoldemTableGame, error) {
	var resp struct {
		TransactionId string
		Result        []*model_dto.HoldemTableGame
	}

	_, peerId, err := h.getMdlEssentials(h.channel)
	if err != nil {
		return nil, err
	}

	url := h.genQueryUrl(h.channel,
		h.chaincodeName,
		"GetGameHistory",
		peerId,
		[]string{tableId, fmt.Sprintf("%d", gameSeq)},
	)
	if err != nil {
		return nil, err
	}

	respData, err := h.sendAndRecvGet(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// Query
// GetGameLastState implements HashVegasContract
func (h *hashVegasContractMdl) GetGameLastState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	var resp struct {
		TransactionId string
		Result        *model_dto.HoldemTableGame
	}

	_, peerId, err := h.getMdlEssentials(h.channel)
	if err != nil {
		return nil, err
	}

	url := h.genQueryUrl(
		h.channel,
		h.chaincodeName,
		"GetGameLastState",
		peerId,
		[]string{tableId, fmt.Sprintf("%d", gameSeq)},
	)
	if err != nil {
		return nil, err
	}

	respData, err := h.sendAndRecvGet(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// Invoke
// InitGenesisLedgerForTable implements HashVegasContract
func (h *hashVegasContractMdl) InitGenesisLedgerForTable(tableId string) (*model_dto.HoldemTableGame, error) {
	var resp struct {
		TransactionId string
		Result        *model_dto.HoldemTableGame
	}

	orderer, peerId, err := h.getMdlEssentials(h.channel)
	if err != nil {
		return nil, err
	}

	bodyValues := h.genEssentialValues(orderer, peerId)
	bodyValues.Add("function", "InitGenesisLedgerForTable")
	bodyValues.Add("args", tableId)

	respData, err := h.sendAndRecvPost(h.invokeBaseUrl, bodyValues)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// Query
// getDeckInitState implements HashVegasContract
func (h *hashVegasContractMdl) getDeckInitState(tableId string, gameSeq uint64) (*model_dto.HoldemTableGame, error) {
	var resp struct {
		TransactionId string
		Result        *model_dto.HoldemTableGame
	}

	_, peerId, err := h.getMdlEssentials(h.channel)
	if err != nil {
		return nil, err
	}

	url := h.genQueryUrl(h.channel,
		h.chaincodeName,
		"GetGameHistory",
		peerId,
		[]string{tableId, fmt.Sprintf("%d", gameSeq)},
	)
	if err != nil {
		return nil, err
	}

	respData, err := h.sendAndRecvGet(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (h *hashVegasContractMdl) getTargetOrdererName() (string, error) {
	var orderers []struct {
		OrdererName string `json:"orderer_name"`
		EnrollId    string `json:"enroll_id"`
	}

	url := fmt.Sprintf("%s/api/blockchain/orderers?mdl_user=admin", h.addres)
	// log.Println("orderer url: ", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("request failed:", err)
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Error("response failed:", err)
		return "", err
	}
	defer res.Body.Close()

	if err = h.checkRespStatus(res); err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("response failed:", err)
		return "", err
	}

	if err := json.Unmarshal(body, &orderers); err != nil {
		log.Error("unmarshalling orderer failed:", err)
		return "", err
	}
	printMarshallIndentJson(orderers)

	if len(orderers) <= 0 {
		err = errors.New("could not find any orderer of network")
		log.Error(err.Error())
		return "", err
	}

	return orderers[0].OrdererName, nil
}

func (m *hashVegasContractMdl) checkRespStatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(fmt.Sprintf(
			"Got failed response (status:%d), (body:%s)",
			resp.StatusCode,
			string(body)),
		)
	}
	return nil
}

func (m *hashVegasContractMdl) sendAndRecvGet(queryUri string) ([]byte, error) {
	request, err := http.NewRequest("GET", queryUri, nil)
	if err != nil {
		log.Error("request failed:", err)
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// log.Println("result status code: ", res.StatusCode)
	err = m.checkRespStatus(res)
	if err != nil {
		return nil, err
	}

	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

func (m *hashVegasContractMdl) sendAndRecvPost(invokeUri string, bodyValues url.Values) ([]byte, error) {
	res, err := http.PostForm(invokeUri, bodyValues)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	printMarshallIndentJson(bodyValues)

	if err := m.checkRespStatus(res); err != nil {
		return nil, err
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

func (m *hashVegasContractMdl) getMdlEssentials(channel string) (string, string, error) {
	ordererName, err := m.getTargetOrdererName()
	if err != nil {
		return "", "", nil
	}

	peerId, err := m.getTargetPeerId(ordererName, channel)
	if err != nil {
		return "", "", nil
	}

	return ordererName, peerId, nil

}

func (m *hashVegasContractMdl) genEssentialValues(ordererName string, peerId string) url.Values {
	values := url.Values{}
	values.Add("mdl_user", "admin")
	values.Add("orderer_ids", ordererName)
	values.Add("peer_ids", peerId)

	return values
}

func (m *hashVegasContractMdl) genQueryUrl(channel, chaincode, function, peerId string, args []string) string {

	var argQueryParams string
	for _, arg := range args {
		argQueryParams = argQueryParams + fmt.Sprintf("&args=%s", arg)
	}

	queryUri := fmt.Sprintf(`%s/api/blockchain/channels/%s/chaincodes/%s/query?mdl_user=admin&function=%s&peer_ids=%s&%s`,
		m.addres,
		channel,
		chaincode,
		function,
		peerId,
		argQueryParams,
	)

	log.Info("queryUri:", queryUri)

	return queryUri
}

func (m *hashVegasContractMdl) getTargetPeerId(ordererName string, channelName string) (string, error) {
	var ret string

	var peers []struct {
		Id           int    `json:"id"`
		PeerEnrollId string `json:"peer_enroll_id"`
	}

	url := fmt.Sprintf("%s/api/blockchain/orderers/%s/channels/%s/joined_peers?mdl_user=admin",
		m.addres,
		ordererName,
		channelName)

	log.Infof("peer uri: %s", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("request failed:", err)
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Error("response failed:", err)
		return "", err
	}
	defer res.Body.Close()

	if err := m.checkRespStatus(res); err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("reading body failed:", err)
		return "", err
	}
	log.Infof("peer data: %s", string(data))
	json.Unmarshal(data, &peers)

	if len(peers) <= 0 {
		err = errors.New("could not find any installed peer of the channel (" + channelName + ")")
		log.Error(err.Error())
		return "", err
	}

	ret = fmt.Sprintf("%d", peers[0].Id)

	return ret, nil
}

func printMarshallIndentJson(v interface{}) error {
	respData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Printf("marshallIndent failed: %s\n", err.Error())
		return err
	}

	fmt.Printf("object data: %s\n", respData)

	return err
}

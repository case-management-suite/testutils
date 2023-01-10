package testutil

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"net/http"
	"strings"

	"github.com/case-management-suite/models"
	"github.com/rs/zerolog/log"
)

func CreateRequest(url string, method string, data io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		// handle error
		log.Fatal().Err(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode/100 == 4 {
		return nil, fmt.Errorf("ERROR HTTP %d %s: %s", resp.StatusCode, resp.Status, req.URL)
	}
	if resp.StatusCode/100 == 5 {
		return nil, fmt.Errorf("ERROR HTTP %d %s: %s", resp.StatusCode, resp.Status, req.URL)
	}
	if err != nil {
		// handle error
		log.Fatal().Err(err)
	}
	return ioutil.ReadAll(resp.Body)
}

func PutRequest(url string, data io.Reader) ([]byte, error) {
	return CreateRequest(url, http.MethodPut, data)
}

func PostRequest(url string, data io.Reader) ([]byte, error) {
	return CreateRequest(url, http.MethodPost, data)
}

func GetRequest(url string) ([]byte, error) {
	return CreateRequest(url, http.MethodGet, nil)
}

func ParseUUIDReponse(resp []byte) (models.UUIDResponse, error) {
	uuidReponse := models.UUIDResponse{}

	err := json.Unmarshal(resp, &uuidReponse)
	return uuidReponse, err
}

func ParseCaseRecord(resp []byte) (models.CaseRecord, error) {
	caseRecord := models.CaseRecord{}
	err := json.Unmarshal(resp, &caseRecord)
	return caseRecord, err
}

func ParseCaseActions(resp []byte) ([]models.CaseAction, error) {
	caseRecord := []models.CaseAction{}
	err := json.Unmarshal(resp, &caseRecord)
	return caseRecord, err
}

func CreateCase() (models.UUIDResponse, error) {
	actionUUIDBytes, err := PutRequest("http://localhost:8080/case", strings.NewReader("any thing"))
	if err != nil {
		return models.UUIDResponse{}, err
	}
	return ParseUUIDReponse(actionUUIDBytes)
}

func ExecuteAction(id models.Identifier, action string) (models.UUIDResponse, error) {
	url := fmt.Sprintf("http://localhost:8080/case/%v?action=%v", id, action)
	actionUUIDBytes, err := PostRequest(url, strings.NewReader("any thing"))
	if err != nil {
		return models.UUIDResponse{}, err
	}
	return ParseUUIDReponse(actionUUIDBytes)
}

func FindCase(id models.Identifier) (models.CaseRecord, error) {
	url := fmt.Sprintf("http://localhost:8080/case/%v", id)
	caseRecordResponse, err := GetRequest(url)
	if err != nil {
		return models.CaseRecord{}, err
	}
	return ParseCaseRecord(caseRecordResponse)
}

func FindCaseActions(id models.Identifier) ([]models.CaseAction, error) {
	url := fmt.Sprintf("http://localhost:8080/case/%v/actions", id)
	caseRecordResponse, err := GetRequest(url)
	if err != nil {
		return []models.CaseAction{}, err
	}
	return ParseCaseActions(caseRecordResponse)
}

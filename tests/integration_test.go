package tests

import (
	base "avito-pvz/internal/handlers/dto"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

const (
	baseUrl = "http://localhost:8080"
)

func TestCreatePVZAndReception(t *testing.T) {
	client := &http.Client{}

	// receiving tokens for both roles
	resp1 := GetToken(base.ModeratorRole)
	tokenModerator := resp1["Token"].(string)
	resp2 := GetToken(base.EmployeeRole)
	tokenEmployee := resp2["Token"].(string)

	pvzRequest := base.PVZ{
		ID:               uuid.New(),
		RegistrationDate: time.Now(),
		City:             base.MoscowCity,
	}
	// creating new pvz
	pvzResp := sendAuthenticatedRequest(t, client, "POST", baseUrl+"/pvz", tokenModerator, pvzRequest)

	pvzID := pvzResp["id"].(string)
	pvzUUID, _ := uuid.Parse(pvzID)

	receptionRequest := map[string]interface{}{
		"pvzId": pvzUUID,
	}

	// opening new reception
	sendAuthenticatedRequest(t, client, "POST", baseUrl+"/receptions", tokenEmployee, receptionRequest)

	// adding 50 new products to reception
	for i := 0; i < 50; i++ {
		productRequest := base.ProductRequest{
			Type:  base.ElectronicType,
			PvzID: pvzUUID,
		}
		sendAuthenticatedRequest(t, client, "POST", baseUrl+"/products", tokenEmployee, productRequest)
	}

	closeRequest := base.ReceptionRequest{
		PvzID: pvzUUID,
	}

	// closing reception
	sendAuthenticatedRequest(t, client, "POST", baseUrl+"/pvz/"+pvzID+"/close_last_reception", tokenEmployee, closeRequest)
}

func sendAuthenticatedRequest(t *testing.T, client *http.Client, method, url, token string, requestBody interface{}) map[string]interface{} {
	reqBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var result map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil
	}

	assert.NotEqual(t, resp.Status, http.StatusBadRequest)
	assert.NotEqual(t, resp.Status, http.StatusInternalServerError)
	assert.NotEqual(t, resp.Status, http.StatusForbidden)
	assert.NotEqual(t, resp.Status, http.StatusUnauthorized)

	return result
}

func GetToken(role base.UserRole) map[string]interface{} {
	req := base.DummyRequest{
		Role: role,
	}
	ioReq, _ := json.Marshal(req)
	resp, _ := http.Post(baseUrl+"/dummyLogin", "application/json", bytes.NewBuffer(ioReq))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	var token map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil
	}
	return token
}

package tests

import (
	base "avito-pvz/internal/handlers"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestCreatePVZAndReception(t *testing.T) {
	client := &http.Client{}
	baseUrl := "http://localhost:8080"

	registerRequestModerator := base.RegisterRequest{
		Email:    base.TestUser1Email,
		Password: base.TestUserPass,
		Role:     base.ModeratorRole,
	}
	sendRequest(t, client, "POST", baseUrl+"/register", registerRequestModerator)

	loginRequestModerator := base.LoginRequest{
		Email:    base.TestUser1Email,
		Password: base.TestUserPass,
	}
	loginRespModerator := sendRequest(t, client, "POST", baseUrl+"/login", loginRequestModerator)

	tokenModerator := loginRespModerator["Token"].(string)

	pvzRequest := base.PVZ{
		ID:               uuid.New(),
		RegistrationDate: time.Now(),
		City:             base.MoscowCity,
	}
	pvzResp := sendAuthenticatedRequest(t, client, "POST", baseUrl+"/pvz", tokenModerator, pvzRequest)

	pvzID := pvzResp["id"].(string)
	pvzUUID, _ := uuid.Parse(pvzID)

	receptionRequest := map[string]interface{}{
		"pvzId": pvzUUID,
	}

	registerRequestEmployee := base.RegisterRequest{
		Email:    base.TestUser2Email,
		Password: base.TestUserPass,
		Role:     base.EmployeeRole,
	}
	sendRequest(t, client, "POST", baseUrl+"/register", registerRequestEmployee)

	loginRequestEmployee := base.LoginRequest{
		Email:    base.TestUser2Email,
		Password: base.TestUserPass,
	}
	loginRespEmployee := sendRequest(t, client, "POST", baseUrl+"/login", loginRequestEmployee)

	tokenEmployee := loginRespEmployee["Token"].(string)
	sendAuthenticatedRequest(t, client, "POST", baseUrl+"/receptions", tokenEmployee, receptionRequest)

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

	sendAuthenticatedRequest(t, client, "POST", baseUrl+"/pvz/"+pvzID+"/close_last_reception", tokenEmployee, closeRequest)
}

func sendRequest(t *testing.T, client *http.Client, method, url string, requestBody interface{}) map[string]interface{} {
	reqBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AuthenticationData struct {
	Data struct {
		Ticket              string `json:"ticket"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
	} `json:"data"`
}

var credentials AuthenticationData

func authenticate() AuthenticationData {
	credentials = AuthenticationData{}
	addr := fmt.Sprintf("%s/api2/json/access/ticket", pmURL)
	data := url.Values{
		"username": {pmUsername},
		"password": {pmPassword},
	}
	requestBody := bytes.NewBufferString(data.Encode()).Bytes()
	response, _ := http.Post(addr, "application/x-www-form-urlencoded", bytes.NewReader(requestBody))
	if response.StatusCode > 299 {
		log.Fatal(fmt.Errorf("authentication returned error %d", response.StatusCode))
	}
	responseBody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(responseBody, &credentials)
	if err != nil {
		log.Fatal(err)
	}
	return credentials
}

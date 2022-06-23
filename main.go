package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	AUTH_TOKEN = 0
	AUTH_BASIC = 1
)

var pmURL string
var pmTokenId string
var pmTokenSecret string
var pmUsername string
var pmPassword string
var pmAllowInsecureTLS bool
var ansibleUser string

var authMode int

func main() {
	flag.StringVar(&pmURL, "url", "https://localhost:8006", "Proxmox API URL")
	flag.StringVar(&pmTokenId, "tokenId", "", "Proxmox Token ID - if this is set, username/password parameters are ignored. Can also be set via the PROXMOX_TOKEN_ID environment variable")
	flag.StringVar(&pmTokenSecret, "tokenSecret", "", "Proxmox Token Secret. Can also be set (and is recommended as such) via the PROXMOX_TOKEN_SECRET environment variable")
	flag.StringVar(&pmUsername, "api-user", "", "Proxmox User. Can also be set via the PROXMOX_API_USERNAME environment variable")
	flag.StringVar(&pmPassword, "api-password", "", "Proxmox Password. Can also be set (and is recommended as such) via the PROXMOX_API_PASSWORD environment variable")
	flag.BoolVar(&pmAllowInsecureTLS, "allow-insecure-tls", false, "Allow insecure TLS communication with Proxmox")
	flag.StringVar(&ansibleUser, "ansible-user", "", "SSH user on which Ansible should attempt to connect")
	flag.Parse()

	if os.Getenv("PROXMOX_TOKEN_ID") != "" {
		pmTokenId = os.Getenv("PROXMOX_TOKEN_ID")
	}
	if os.Getenv("PROXMOX_TOKEN_SECRET") != "" {
		pmTokenSecret = os.Getenv("PROXMOX_TOKEN_SECRET")
	}
	if os.Getenv("PROXMOX_API_USERNAME") != "" {
		pmUsername = os.Getenv("PROXMOX_API_USERNAME")
	}
	if os.Getenv("PROXMOX_API_PASSWORD") != "" {
		pmPassword = os.Getenv("PROXMOX_API_PASSWORD")
	}

	validateArgs()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: pmAllowInsecureTLS}

	if authMode == AUTH_BASIC {
		authenticate()
	}

	nodes := nodes()
	buildYAML(nodes)
}

func validateArgs() {
	if pmTokenId != "" {
		authMode = AUTH_TOKEN
		if pmTokenSecret == "" {
			log.Fatal("Token ID provided but no token secret. Exiting.")
		}
		if pmUsername != "" || pmPassword != "" {
			log.Fatal("Token ID provided alongside username or password - choose one authentication method only. Exiting.")
		}
	} else {
		authMode = AUTH_BASIC
		if pmUsername != "" {
			if pmPassword == "" {
				log.Fatal("API Username provided, but no password. Exiting.")
			}
		}
		if pmPassword != "" {
			if pmUsername == "" {
				log.Fatal("API Password provided, but no username. Exiting.")
			}
		}
	}
}

func sendRequest(request *http.Request) []byte {
	switch authMode {
	case AUTH_TOKEN:
		request.Header.Set("Authorization", fmt.Sprintf("%s=%s", pmTokenId, pmTokenSecret))
	case AUTH_BASIC:
		request.AddCookie(&http.Cookie{Name: "PVEAuthCookie", Value: credentials.Data.Ticket})
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(fmt.Errorf("errored when sending request to server: %s", err))
	}
	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	return responseBody
}

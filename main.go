package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var pmURL string
var pmUsername string
var pmPassword string
var pmAllowInsecureTLS bool
var ansibleUser string

func main() {
	flag.StringVar(&pmURL, "url", "https://localhost:8006", "Proxmox API URL")
	flag.StringVar(&pmUsername, "user", "", "Proxmox User")
	flag.StringVar(&pmPassword, "password", "", "Proxmox Password (recommended: use the PROXMOX_API_PASSWORD environment variable instead")
	flag.BoolVar(&pmAllowInsecureTLS, "allow-insecure-tls", false, "Allow insecure TLS communication with Proxmox")
	flag.StringVar(&ansibleUser, "ansible-user", "", "SSH user on which Ansible should attempt to connect")
	flag.Parse()

	if os.Getenv("PROXMOX_API_PASSWORD") != "" {
		pmPassword = os.Getenv("PROXMOX_API_PASSWORD")
	}

	if pmAllowInsecureTLS {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	authenticate()

	nodes := nodes()
	buildYAML(nodes)
}

func sendRequest(request *http.Request) []byte {
	if credentials != (AuthenticationData{}) {
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

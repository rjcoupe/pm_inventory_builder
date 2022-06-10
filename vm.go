package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type VM struct {
	ID     string
	Name   string
	IP     string
	Tags   []string
	Status string
}

func vms_on_node(nodeId string) []VM {
	var vms []VM
	url := fmt.Sprintf("%s/api2/json/nodes/%s/qemu", pmURL, nodeId)
	req, _ := http.NewRequest("GET", url, nil)
	var data map[string][]interface{}
	response := sendRequest(req)
	json.Unmarshal(response, &data)
	for _, item := range data["data"] {
		parsed := item.(map[string]interface{})
		vm := VM{}
		vm.ID = fmt.Sprintf("%v", parsed["vmid"])
		vm.Name = fmt.Sprintf("%v", parsed["name"])
		if parsed["tags"] != nil {
			vm.Tags = strings.Split(fmt.Sprintf("%v", parsed["tags"]), " ")
		}
		vm.Status = fmt.Sprintf("%v", parsed["status"])
		vm.IP = vm_ip_address(nodeId, vm.ID)
		if vm.Status == "running" {
			vms = append(vms, vm)
		}
	}
	return vms
}

func vm_ip_address(nodeID string, vmID string) string {
	var ip string
	url := fmt.Sprintf("%s/api2/json/nodes/%s/qemu/%s/agent/network-get-interfaces", pmURL, nodeID, vmID)
	req, _ := http.NewRequest("GET", url, nil)
	var data map[string]map[string][]interface{}
	response := sendRequest(req)
	json.Unmarshal(response, &data)
	for _, iface := range data["data"]["result"] {
		parsed := iface.(map[string]interface{})
		if parsed["name"] == "lo" {
			continue
		}
		ipaddresses := parsed["ip-addresses"].([]interface{})
		for _, ip := range ipaddresses {
			parsedIP := ip.(map[string]interface{})
			if parsedIP["ip-address-type"] == "ipv4" {
				return fmt.Sprintf("%v", parsedIP["ip-address"])
			}
		}
	}
	return ip
}

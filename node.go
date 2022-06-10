package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Node struct {
	ID  string `json:"id"`
	VMs []VM
}

func nodes() []Node {
	var nodes []Node
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api2/json/nodes", pmURL), nil)
	response := sendRequest(req)
	var data map[string][]interface{}
	json.Unmarshal(response, &data)
	for _, item := range data["data"] {
		id := fmt.Sprintf("%v", item.(map[string]interface{})["node"])
		node := Node{ID: id}
		node.VMs = vms_on_node(node.ID)
		nodes = append(nodes, node)
	}
	return nodes
}

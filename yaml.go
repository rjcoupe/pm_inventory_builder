package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func buildYAML(nodes []Node) {
	output := map[string]map[string]map[string]map[string]string{}
	output["all"] = make(map[string]map[string]map[string]string)
	output["all"]["hosts"] = make(map[string]map[string]string)
	output["ungrouped"] = make(map[string]map[string]map[string]string)
	output["ungrouped"]["hosts"] = make(map[string]map[string]string)
	for _, node := range nodes {
		for _, vm := range node.VMs {
			output["all"]["hosts"][vm.Name] = make(map[string]string)
			output["all"]["hosts"][vm.Name]["ansible_host"] = vm.IP
			if len(vm.Tags) == 0 {
				output["ungrouped"]["hosts"][vm.Name] = make(map[string]string)
				output["ungrouped"]["hosts"][vm.Name]["ansible_host"] = vm.IP
				if ansibleUser != "" {
					output["ungrouped"]["hosts"][vm.Name]["ansible_user"] = ansibleUser
				}
			} else {
				for _, tag := range vm.Tags {
					if _, ok := output[tag]; ok {
						output[tag]["hosts"][vm.Name] = make(map[string]string)
						output[tag]["hosts"][vm.Name]["ansible_host"] = vm.IP
						if ansibleUser != "" {
							output[tag]["hosts"][vm.Name]["ansible_user"] = ansibleUser
						}
					} else {
						output[tag] = make(map[string]map[string]map[string]string)
						output[tag]["hosts"] = make(map[string]map[string]string)
						output[tag]["hosts"][vm.Name] = make(map[string]string)
						output[tag]["hosts"][vm.Name]["ansible_host"] = vm.IP
						if ansibleUser != "" {
							output[tag]["hosts"][vm.Name]["ansible_user"] = ansibleUser
						}
					}
				}
			}
		}
	}
	yamlOutput, _ := yaml.Marshal(&output)
	fmt.Println(string(yamlOutput))
}

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type response struct {
	Page    int      `json:"page"`
	Entries []string `json:"entries"`
}

func main() {

	j := `[{"key":"any","blob":["{\"reason\":\"UNKNOWN\"}\"]},{\"key\":\"confirmation\",\"blob\":[\"{\"related_qtty\":1,\"input_source\":\"1\",\"tag\":\"service_confirmation\"}"]},{"key":"service","blob":["{\"account_owner\":\"teste\",\"bank_account\":\"d750a2de874243642086ef67d0805a8e73e3deba1c28b296c755168b56b42862a5045e51d0dff7fa1f\",\"sub_type\":\"flex_frac\",\"type\":\"insurance\"}"]}]`

	var itemInfoCollection []ItemInfo
	err := json.Unmarshal([]byte(j), &itemInfoCollection)
	if err != nil {
		panic(err)
	}
	for _, itemInfo := range itemInfoCollection {
		jsonMap := make(map[string]interface{})
		if strings.EqualFold(itemInfo.Key, "service") {
			for _, blob := range itemInfo.Blob {

				err := json.Unmarshal([]byte(blob), &jsonMap)
				if err != nil {
					panic(err)
				}
				valueInterface, exists := jsonMap["type"]
				if !exists {
					fmt.Println("key not exists")
				}

				valueString, ok := valueInterface.(string)
				if !ok {
					fmt.Println("value not exists")
				}
				fmt.Println(valueString)
			}
		}
	}
}

type ItemInfo struct {
	Key  string   `json:"key"`
	Blob []string `json:"blob"`
}

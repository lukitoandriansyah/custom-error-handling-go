package util

import (
	"encoding/json"
	"fmt"
)

func StructToJson(StructToJsonValue any) string {
	jsonData, err := json.Marshal(StructToJsonValue)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}
	return string(jsonData)
}

func JsonToStruct(JsonToStructValue []byte, StructValue any) error {
	err := json.Unmarshal([]byte(JsonToStructValue), &StructValue)
	return err
}

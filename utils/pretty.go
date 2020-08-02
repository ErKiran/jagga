package utils

import (
	"encoding/json"
	"fmt"
)

func Pretty(name string, obj interface{}) {
	json, _ := json.MarshalIndent(obj, "", " ")
	fmt.Println(name, string(json))
}

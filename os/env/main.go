package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type envValue interface {
	GetString() string
}

// envValString is a string environment value.
type envValString struct {
	Value string
}

func (e *envValString) GetString() string {
	return e.Value
}

func (e *envValString) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Value)
}

// envValPathList is a slice with PATH values.
type envValPathList struct {
	Values []string
}

// GetString returns a string with all values in a slice, separated by os.PathListSeparator.
func (e *envValPathList) GetString() string {
	return strings.Join(e.Values, string(os.PathListSeparator))
}

func (e *envValPathList) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Values)
}

func main() {
	env := os.Environ()
	envMap := make(map[string]envValue) // also possible to use an "interface{}" as a value, but an empty interface says nothing.
	fmt.Println(string(os.PathListSeparator))

	for _, e := range env {
		split := strings.Split(e, "=")
		if split[0] == "PATH" {
			envArr := &envValPathList{
				Values: strings.Split(split[1], string(os.PathListSeparator)),
			}
			envMap["PATH"] = envArr
		} else {
			envStr := &envValString{
				Value: split[1],
			}
			envMap[split[0]] = envStr
		}
	}

	json, _ := json.MarshalIndent(envMap, "", "    ")
	fmt.Println(string(json))

}

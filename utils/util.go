package utils

import (
	"encoding/json"
	"github.com/dminGod/D30-HectorDA/logger"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// IsJSON validates a JSON string
func IsJSON(input interface{}) bool {

	var output map[string]interface{}

	return json.Unmarshal([]byte(input.(string)), &output) == nil
}

// DecodeJSON converts a JSON string to a map of string and interface
func DecodeJSON(input interface{}) map[string]interface{} {

	var payload map[string]interface{}

	if !IsJSON(input) {
		return payload
	}
	err := json.Unmarshal([]byte(input.(string)), &payload)
	HandleError(err)

	return payload

}

// EncodeJSON converts a map of string and interface to a JSON string
func EncodeJSON(input interface{}) string {
	jsonString, err := json.Marshal(input)
	HandleError(err)	

	return string(jsonString)
}

// KeyInMap checks if a given key exists in a map of string and interface
func KeyInMap(key string, attributes map[string]interface{}) bool {

	// iterate over each route
	for k := range attributes {

		if key == k {
			return true
		}
	}

	return false
}

// FindMap checks if a given key matches a given value and returns the entire map
func FindMap(key string, value interface{}, input map[string]interface{}) map[string]interface{} {

	output := make(map[string]interface{})
	// iterate over each map
	for _, v := range input {
		meta := v.(map[string]interface{})
		if meta[key] == value {
			output = meta
		}
	}

	return output

}

// ReadFile returns the contents of the file
func ReadFile(path string) string {

	raw, err := ioutil.ReadFile(path)
	HandleError(err)
	return string(raw)
}

// ParseFilter is used to convert an LDAP type query filter to a map of string and interface
func ParseFilter(input string) map[string]string {

	output := make(map[string]string)

	pattern := `(^\(*\&?\(*)(.*)(\)?\)$)`

	if !RegexMatch(input, pattern) {
		return output
	}

	var validID = regexp.MustCompile(pattern)
	input = (validID.FindStringSubmatch(input))[2]

	/*if len(input) == 0 {
		return output
	}

	input = input[1:]
	input = strings.Trim(input,")")
	if string(input[0]) == "&" {
		input = input[1:]
	}*/

	filters := strings.Split(input, ")(")

	for _, v := range filters {

		v = strings.Replace(v, "(", "", 1)
		v = strings.Replace(v, ")", "", 1)

		keyval := strings.Split(v, "=")
		output[keyval[0]] = keyval[1]
	}

	return output

}

// RegexMatch is used to match an input string with a regex pattern
func RegexMatch(input string, pattern string) bool {

	var validID = regexp.MustCompile(pattern)

	return validID.MatchString(input)
}

// Exit is used to Exit the application with the provided exit code
func Exit(code int) {

	os.Exit(code)

}

func ExecuteCommand(command string, args ...string) string {
	out, _ := exec.Command(command, args...).Output()
	output := string(out)
	output = strings.Trim(output, "\r")
	output = strings.Trim(output, "\n")

	return output
}


func HandleError(err error) {
	if err != nil {
		logger.Write("ERROR", err.Error())
	}
}

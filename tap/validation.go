package tap

import (
	"fmt"
	"regexp"
	"strings"
)

var ALLOWED_METHODS = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}
var rePlainMethod = regexp.MustCompile(`^([[:alpha:]]+) \S+ HTTP/1\.[01]`)
var reJsonMethod = regexp.MustCompile(`{"key":":method","value":"([^"]*)"}`)

func ValidateMethod(method string) bool {
	isAllowed := false
	for _, allowed_method := range ALLOWED_METHODS {
		if strings.ToUpper(method) == allowed_method {
			isAllowed = true
			break
		}
	}

	return isAllowed
}

func ExtractMethodFromPlain(messageBytes []byte) (string, error){
	matches := rePlainMethod.FindSubmatch(messageBytes)
	if len(matches) != 2 {
		return "", fmt.Errorf("did not find method in message")
	}

	return string(matches[1]), nil
}

func ExtractMethodFromJson(messageBytes []byte) (string, error){
	matches := reJsonMethod.FindSubmatch(messageBytes)
	if len(matches) != 2 {
		return "", fmt.Errorf("did not find method in message")
	}

	return string(matches[1]), nil
}

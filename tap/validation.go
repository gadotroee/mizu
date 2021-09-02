package tap

import "strings"

var ALLOWED_METHODS = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

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

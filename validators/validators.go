package validators

import (
	"regexp"
)

func IsValidateSizeCep(cep string) bool {
	re, _ := regexp.Compile(`[^0-9]`)
	cep = re.ReplaceAllString(cep, "")
	if len(cep) < 8 || len(cep) > 8 {
		return false
	}
	return true
}

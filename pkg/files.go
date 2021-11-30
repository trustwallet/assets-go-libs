package pkg

import (
	"strings"
)

func IsFileAllowedInPR(path string) bool {
	if strings.HasSuffix(path, "tokenlist.json") {
		return false
	}
	if strings.HasPrefix(path, "blockchains") && strings.Index(path, "assets") > 0 {
		return true
	}
	if strings.HasPrefix(path, "blockchains") && strings.HasSuffix(path, "allowlist.json") {
		return true
	}
	if strings.HasPrefix(path, "blockchains") && strings.HasSuffix(path, "validators/list.json") {
		return true
	}
	if strings.HasPrefix(path, "dapps") {
		return true
	}

	return false
}

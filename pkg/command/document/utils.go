package document

import (
	"errors"
	"strings"
)

func splitPath(path string) (namespace, slug string, err error) {
	singlePath := strings.Split(path, "/")
	if len(singlePath) != 3 {
		return "", "", errors.New("Error path")
	}

	return singlePath[0] + "/" + singlePath[1], singlePath[2], nil
}

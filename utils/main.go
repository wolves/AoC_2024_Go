package utils

import (
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInput(filename, delim string) []string {
	raw, err := os.ReadFile(filename)
	CheckErr(err)

	return strings.Split(strings.TrimSpace(string(raw)), delim)
}

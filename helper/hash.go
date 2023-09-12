package helper

import (
	"crypto/sha256"
	"strings"
)

func HashMake(text string) string {
    sha := sha256.New()
    sha.Write([]byte(text))
    var encrypted = sha.Sum(nil)
    return string(encrypted)
}

func HashVerify(text string, password string) bool {
    result := HashMake(text)
    return strings.Compare(password, result) == 0
}

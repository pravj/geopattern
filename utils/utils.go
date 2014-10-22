package utils

import (
    "fmt"
    "crypto/sha1"
)

func Hash(s string) string {
    h := sha1.New()
    h.Write([]byte(s))

    hash := h.Sum(nil)

    return fmt.Sprintf("%x", hash)
}

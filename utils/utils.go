package utils

import (
    "fmt"
    "crypto/sha1"
    "strconv"
)

func Hash(s string) string {
    h := sha1.New()
    h.Write([]byte(s))

    hash := h.Sum(nil)

    return fmt.Sprintf("%x", hash)
}

func Map(value, a_min, a_max, b_min, b_max float64) int {
    a_range := a_max - a_min
    b_range := b_max - b_min

    return int(b_max - (a_max - value) * (b_range/a_range))
}

func Hex_val(str string, index, length int) int {
    hex_str := str[index:index+length]

    hex_val, err := strconv.ParseInt(hex_str, 16, 0)
    if err != nil {
        panic(err)
    }

    return int(hex_val)
}

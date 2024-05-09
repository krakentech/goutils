package goutils

import "strconv"

// StringToInt will convert the string to int or return 0
func StringToInt(in string) int {
    out, err := strconv.Atoi(in)
    if err != nil {
        return 0
    }
    return out
}

// StringToInt64 will convert the string to int64 or return 0
func StringToInt64(in string) int64 {
    out, err := strconv.ParseInt(in, 10, 64)
    if err != nil {
        return 0
    }
    return out
}

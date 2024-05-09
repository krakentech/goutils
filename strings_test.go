package goutils

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "math"
    "testing"
)

func TestStringToInt(t *testing.T) {

    tests := []struct {
        name string
        arg  string
        want int
    }{
        {name: "empty string", arg: "", want: 0},
        {name: "non-numeric string", arg: "this is a test", want: 0},
        {name: "max int", arg: fmt.Sprintf("%v", uint(math.MaxInt)), want: int(math.MaxInt)},
        {name: "string value larger than in can handle", arg: fmt.Sprintf("%v", uint64(math.MaxInt)+1), want: 0},
        {name: "real value", arg: "9", want: 9},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, StringToInt(tt.arg), "StringToInt(%v)", tt.arg)
        })
    }
}

func TestStringToInt64(t *testing.T) {

    tests := []struct {
        name string
        arg  string
        want int64
    }{
        {name: "empty string", arg: "", want: 0},
        {name: "non-numeric string", arg: "this is a test", want: 0},
        {name: "string value larger than in can handle", arg: fmt.Sprintf("%v", uint64(math.MaxInt)+1), want: 0},
        {name: "real value", arg: "9", want: 9},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, StringToInt64(tt.arg), "StringToInt(%v)", tt.arg)
        })
    }
}

package util

import (
	"fmt"
	"testing"
)

func TestGetInput(t *testing.T) {
	ans := GetInput("test.in")
    if len(ans) != 3 {
        t.Errorf("len(GetInput(\"test.in\") = %d; want 3", len(ans))
    }
}

func TestAtoi(t *testing.T) {
	var tests = []struct {
        s string
        want int
    }{
        {"0", 0},
        {"1", 1},
        {"-1", -1},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.s)
        t.Run(testname, func(t *testing.T) {
            ans := Atoi(tt.s)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
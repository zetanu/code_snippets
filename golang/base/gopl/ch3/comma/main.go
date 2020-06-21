package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma2("1234567"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	start := len(s) % 3

	if start > 0 {
		buf.WriteString(s[:start])
		buf.WriteByte(',')
	}
	for ; start < len(s); start += 3 {
		buf.WriteString(s[start : start+3])
		if start != len(s)-3 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

package main

import (
	"strconv"
	"strings"
)

func palindrome_two() {
	m := 465

	m_slice := []string{}
	m_string_split := strings.Split(strconv.Itoa(m), "")
	m_slice = append(m_slice, m_string_split...)
}

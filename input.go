package main

import "strings"

func cleantInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	split := strings.Split(trimmed, " ")
	return split
}

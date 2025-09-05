package main

import "strings"

func cleanInput(text string) []string {
	lower_case := strings.ToLower(text)
	return strings.Fields(lower_case)
}

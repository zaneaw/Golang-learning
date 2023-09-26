package main

import "strings"

func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	names := strings.Split(name, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
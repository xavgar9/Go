package main

import "strings"

func Remove(word string, c string) string {
	var ans strings.Builder

	for i:=0; i<len(word); i++{
		if word[i]!=c[0]{
			ans.WriteString(string(word[i]))
		}
	}
	return ans.String()
}

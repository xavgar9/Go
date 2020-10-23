package main

import "testing"

func TestPalindrome(t *testing.T) {
 	table := [] struct {
 		word string
 		check bool
 	}{
 		{"ana", true},
 		{"anitalavalatina", true},
 		{"hello world!", false},
 		{"", true},
 		{"A", true},
 		{"AB", true},
 		{"12", true},
 		{"123", false},
 		{"A 1 A", true},
 	}

 	for _, test := range table{
 		check := palindrome(test.word)
 		if check!=test.check{
 			t.Errorf("[ %s ] Got: %t, want: %t", test.word, check, test.check)
 		}
 	}
}
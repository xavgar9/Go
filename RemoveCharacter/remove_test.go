package main

import "testing"

func TestRemove(t *testing.T) {
 	table := [] struct {
 		word string
 		c string
 		check string
 	}{
 		{"Hello World!", "H", "ello World!"},
 		{"Let's Code!", "e", "Lt's Cod!"},
 		{"", "X", ""},
 		{"Hey! I'm Xavier and I'm learning Golang!", "a", "Hey! I'm Xvier nd I'm lerning Golng!"},
 		{"Monday", "m", "Monday"},
 		{"AaBbCc", "A", "aBbCc"},
 		{"AaBbCc", "a", "ABbCc"},
 		{"Let's go", "'", "Lets go"},
 	}

 	for _, test := range table{
 		check := Remove(test.word, test.c)
 		if check!=test.check{
 			t.Errorf("Got: %s, want: %s", check, test.check)
 		}
 	}
}
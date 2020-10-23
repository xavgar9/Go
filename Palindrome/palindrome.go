package main 

func palindrome(word string) bool{
	ok := true
	if len(word)>2{
		i,j := 0, len(word)-1
		for ok=true;ok;{
			if ok{
				if word[i]!=word[j]{
					ok = false
				} else{
					i++
					j--
				}
			} else{ ok = false }
			if i==j{ break }
		}
	}
	return ok
}
package main

import "fmt"
import "unicode"

func main() {

	myarr := "$^HE$YE#H!"

	flag := true

	for i, j :=0, len(myarr)-1;i<j;i,j = i+1,j-1 {
		if !unicode.IsLetter(rune(myarr[i])) {
			i = i+1
			continue
		}
		if !unicode.IsLetter(rune(myarr[j])) {
			j = j-1
			continue
		}

		if myarr[i] != myarr[j] {
			flag = false 
		}
	}

	fmt.Println("is ", myarr, "palindrome ", flag)
	
}

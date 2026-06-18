//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Use break and continue to control a loop's execution flow.
*/

func main() {
	words:= "Hye I am Ashish"

	for _,ch:=range words {
		if ch == 'I' || ch == ' ' {
			continue
		}
		if ch == 'h'{
			break
		}else{
			fmt.Printf("%c",ch)
		}
	}
}

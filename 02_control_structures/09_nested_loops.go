//go:build ignore

package main

import "fmt"

/*
Difficulty: Medium

Question: Write nested loops to print a multiplication table.
*/

func main() {
	for i:=1;i<=10;i++{
		for j:=1;j<=10;j++{
			fmt.Println(i,"X",j,"=",i*j)
		}
	}
}

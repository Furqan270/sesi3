package main

import (
	"fmt"
)

func main() {
	word := "selamat malam"

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}
	wordc := make(map[rune]int)

	for _, wordb := range word {
		wordc[wordb] += 1
	}
	for wordb, o := range wordc {
		fmt.Printf("map [%c : %d]", wordb, o)
	}
}

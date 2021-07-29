package main

import (
	am "anagram/anagram"
	"fmt"
)

func main() {
	am.Changeloglevel("Debug")
	data := []string{"abc", "bac", "def", "fed", "zbc", "dbc"}
	ans := am.Solve(data)
	fmt.Println(ans)
}

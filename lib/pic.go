package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	wlist := strings.Fields(s)
	wcmap := make(map[string]int)
	for _, w := range wlist {
		_, ok := wcmap[w]
		if ok == true {
			wcmap[w] += 1
		} else {
			wcmap[w] = 1
		}
	}
	return wcmap
}

func main() {
	fmt.Println(WordCount("chu chu chu mu mu tu tu m -1"))
}

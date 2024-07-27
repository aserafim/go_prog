// Dup1 prints the text of each that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
package main

import (
	"fmt"
	"os"
)

//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

func main(){
	s, sep := "", ""
	for index, arg := range os.Args[0:]{ //ignora o índice
		s += fmt.Sprintf("%d %s %s", index, sep, arg)
		sep = " "
	}
	fmt.Println(s)
}

//func main(){
//	fmt.Println(os.Args[0:])	
//}
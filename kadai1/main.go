package main

import (
	"flag"
	"fmt"
)

func main() {
	// 左から順にオプション名、デフォルトの値、helpテキストが引数に入る
	b := flag.String("b", "jpeg", "put target file type")
	a := flag.String("a", "png", "put convert after file type")
	// d:= flag.String("d", "/", "put target directory")
	flag.Parse()
	fmt.Println(*b)
	fmt.Println(*a)
}

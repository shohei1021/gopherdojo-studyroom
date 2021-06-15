package main

import (
	"flag"

	"github.com/gopherdojo-studyroom/kadai1/convert"
)

func main() {
	flag.Parse()
	cs := convert.NewConvertService(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	cs.Convert()
}

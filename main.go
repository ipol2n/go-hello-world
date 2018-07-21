package main

import (
	"fmt"

	gotools "github.com/ipol2n/go-tools"
	"github.com/ipol2n/go-tools/stringhelper"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("dasasfasdf")
	fmt.Println(gotools.Add(1, 2))
	fmt.Println(stringhelper.Upper("asdfafd"))
}

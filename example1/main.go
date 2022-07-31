package main

import (
	"example1/examplepb"
	"fmt"

	"github.com/fananchong/versionpb"
)

func main() {
	msg := &examplepb.Msg1{}
	fmt.Printf("%v\n", versionpb.MinimalVersion(msg))
}

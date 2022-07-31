package main

import (
	"example1/examplepb"
	"fmt"

	"github.com/fananchong/versionpb"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	{
		msg := &examplepb.Msg1{}
		fmt.Printf("%v\n", versionpb.MinimalVersion(msg))
	}

	{
		msg := &examplepb.Msg1{F6: examplepb.Msg1_E2}
		fmt.Printf("%v\n", versionpb.MinimalVersion(msg))
	}

	{
		msg := &examplepb.Msg1{F6: examplepb.Msg1_E3}
		fmt.Printf("%v\n", versionpb.MinimalVersion(msg))
	}

	annotations, err := versionpb.AllVersionByFiles(protoregistry.GlobalFiles, []string{"google.protobuf"})
	if err != nil {
		panic(err)
	}
	for _, v := range annotations {
		fmt.Printf("fullname:%v version:%v\n", v.FullName, v.Version)
	}
}

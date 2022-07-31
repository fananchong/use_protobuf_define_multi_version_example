package main

import (
	"example2/examplepb"
	"fmt"

	"github.com/fananchong/versionpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	{
		msg := &examplepb.Msg1{}
		fmt.Printf("%v\n", versionpb.MinimalVersion(proto.MessageReflect(msg)))
	}

	{
		v := int64(0)
		msg := &examplepb.Msg1{F5: &v}
		fmt.Printf("%v\n", versionpb.MinimalVersion(proto.MessageReflect(msg)))
	}

	{
		v := examplepb.Msg1_E2
		msg := &examplepb.Msg1{F6: &v}
		fmt.Printf("%v\n", versionpb.MinimalVersion(proto.MessageReflect(msg)))
	}

	{
		v := examplepb.Msg1_E3
		msg := &examplepb.Msg1{F6: &v}
		fmt.Printf("%v\n", versionpb.MinimalVersion(proto.MessageReflect(msg)))
	}

	annotations, err := versionpb.AllVersionByFiles(protoregistry.GlobalFiles, []string{"google.protobuf"})
	if err != nil {
		panic(err)
	}
	for _, v := range annotations {
		fmt.Printf("fullname:%v version:%v\n", v.FullName, v.Version)
	}
}

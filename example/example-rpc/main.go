package main

import (
	"fmt"
	"os"

	"github.com/nathanielc/grpccmd"
	// Import grpccmd generated code
	_ "github.com/nathanielc/grpccmd/example/internal/pb"
)

//go:generate protoc -I ../internal/pb/ ../internal/pb/example.proto --grpccmd_out=../internal/pb/

func main() {
	grpccmd.SetCmdInfo(
		"example-rpc",
		"Make calls to the Example service",
		"example-rpc command has been autogenerated via the protoc plugin github.com/nathanielc/grpccmd",
	)
	if err := grpccmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

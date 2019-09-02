package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("protoc", "-I", "../proto_files", "../proto_files/bol.proto", "--go_out=plugins=grpc:./internal/pkg/proto")
	out, err := cmd.CombinedOutput()
	_ = cmd.Run()

	if err != nil {
		log.Printf("%s", out)
		log.Fatalln(err)
	}

	log.Printf("%s", out)
}

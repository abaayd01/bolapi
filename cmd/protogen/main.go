package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("protoc", "-I", "../bolproto", "../bolproto/bol.proto", "--go_out=plugins=grpc:./internal/pkg/bolproto")
	out, err := cmd.CombinedOutput()
	_ = cmd.Run()

	if err != nil {
		log.Printf("%s", out)
		log.Fatalln(err)
	}

	log.Printf("%s", out)
}

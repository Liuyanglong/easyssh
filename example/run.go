package main

import (
	"fmt"
	"github.com/hypersleep/easyssh.v0"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "john",
		Server: "example.com",
		// Optional key or Password without either we try to contact your agent SOCKET
		Key:  "/.ssh/id_rsa",
		Port: "22",
	}

	// Call Run method with command you want to run on remote server.
	stdout, stderr, err := ssh.Run("ps ax")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("stdout is :", stdout, ";   stderr is :", stderr)
	}

}

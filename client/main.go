package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Gedis Client v1.0.0")
	fmt.Println()

	conn, err := net.Dial("tcp", ":4455")
	if err != nil {
		fmt.Println("Error: failed to connect to the server. Probably, it's shut down.")
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("] ")

		input, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, input+"\n")

		resp, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error while reading server message: ", err)
			os.Exit(1)
		}

		fmt.Print(resp)
	}
}

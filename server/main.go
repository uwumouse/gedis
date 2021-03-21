package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

// Extra spaces and other things like \t
var space *regexp.Regexp = regexp.MustCompile(`\s+`)

func handleAccept(ln *net.Listener) {
	for {
		conn, err := (*ln).Accept()

		if err != nil {
			fmt.Println("Error occured on client connection: ", err)
			os.Exit(1)
		}

		fmt.Println("Connection: ", conn.RemoteAddr())

		for {
			var args []string

			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error while reading client message:", err)
				break
			}
			// Remove all extraspaces
			message = space.ReplaceAllString(message, " ")

			var commandString []string = strings.Split(message, " ")

			// Args are passed
			if len(commandString) > 1 {
				args = commandString[1:]
			} else if len(commandString) == 1 { // Sent nothing but whitespace
				fmt.Fprintln(conn, "Pass the command.")
				continue
			}

			switch commandString[0] {
			case "get":
				if len(args) <= 1 {
					fmt.Fprintln(conn, "Get syntax: get <key>")
					continue
				}
				value := get(args[0])
				fmt.Fprintln(conn, value)
			case "set":
				if len(args) < 2 {
					fmt.Fprintln(conn, "Set syntax: set <key> <value>")
					continue
				}
				set(args[0], args[1])
				fmt.Fprintf(conn, "%v: %v\n", args[0], args[1])
			default:
				fmt.Fprintf(conn, "Uknown command: %v\n", commandString[0])
			}
		}
	}

}

func main() {
	fmt.Println("Gedis Server v1.0.0")
	fmt.Println()
	fmt.Println("Starting at port 4455")

	ln, err := net.Listen("tcp", ":4455")

	if err != nil {
		fmt.Println("Failed to start Gedis: ", err)
		os.Exit(1)
	}

	handleAccept(&ln)
}

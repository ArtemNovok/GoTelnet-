package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
)

func main() {
	address := os.Args[2]
	network := os.Args[1]
	if len(address) == 0 {
		fmt.Println(color.RedString("address not specified"))
		os.Exit(1)
	}
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println(color.RedString("got error:", err.Error()))
		os.Exit(1)
	}
	fmt.Println(color.GreenString("Connected to "), color.BlueString(address))
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "st" || text == "stop" {
			break
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println(color.RedString("got error:", err.Error()))
		}
	}
	fmt.Println(color.GreenString("done"))
}

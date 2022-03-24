package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("1. begin to connect ...")
	con, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}
	fmt.Printf("connection = %+v\n", con)

	fmt.Println("2. new reader of term")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("3. please input one line data :")
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println("4. send data to server")
		count, err := con.Write([]byte(line))
		if err != nil {
			return
		}
		fmt.Printf("count of data = %+v\n", count)

		fmt.Println("5. receive data from server")
		buf := make([]byte, 1024)
		count, err = con.Read(buf)
		if err != nil {
			return
		}
		fmt.Printf("count of data = %+v\n", count)
		fmt.Printf("\nserver : %s", string(buf[:count]))
	}
}

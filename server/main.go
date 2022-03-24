package main

import (
	"bufio"
	"fmt"
	"net"
)

func process2(con net.Conn) {
	defer con.Close()

	for {
		// get data
		reader := bufio.NewReader(con)
		buf := make([]byte, 1024)
		fmt.Printf("4. wait data from client %+v ... \n", con.RemoteAddr().String())
		count, err := reader.Read(buf)
		if err != nil {
			return
		}
		data := buf[:count]
		fmt.Printf("data = %+v \n", string(data))

		// response data
		con.Write(data)
	}

}

func process1(con net.Conn) {
	defer con.Close()

	for {
		data := make([]byte, 1024)
		fmt.Printf("4. wait data from client %+v ... \n", con.RemoteAddr().String())
		count, err := con.Read(data)
		if err != nil {
			return
		}
		data = data[:count]
		fmt.Printf("data = %+v \n", string(data))
	}

}

func main() {
	fmt.Println("1. listening port ...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}
	defer listen.Close()

	for {
		fmt.Println("2. waiting connection ...")
		connection, err := listen.Accept()
		if err != nil {
			continue
		}

		fmt.Println("3. handler connection ...")
		go process2(connection)
	}
}

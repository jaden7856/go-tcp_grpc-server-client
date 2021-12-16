/* Client n : 1 Server (Read)*/

package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Start server!")

	listen, err := net.Listen("tcp", ":8080")

	if err != nil {
		//fmt.Printf("Fail to bind Address to %d; err : %s\n", port, err)
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Fail to Accept; err : ", err)
			continue
		}

		rcvBuf := make([]byte, 4096)
		reqLen, err := conn.Read(rcvBuf)

		if err != nil {
			if io.EOF == err {
				fmt.Println("Fail to Read : ", err)
				break
			}
		}

		if reqLen > 0 {
			data := rcvBuf[:reqLen]
			fmt.Println(string(data))
			_, err := conn.Write(data[:reqLen])

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

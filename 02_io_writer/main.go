package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	//"net/http"
	"os"
	"strings"
)

func main() {
	//2.4.1
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
	
	//2.4.2
	os.Stdout.Write([]byte("os.Stdout example\n"))
	
	//2.4.3
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example \n"))
	fmt.Println(buffer.String())
	
	io.WriteString(&buffer , "bytes, Buffer example\n")

	//2.4.4
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))//p.21 「Goのバイト列と文字列"」を参照
	fmt.Println(builder.String())

	//2.4.5
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil{
		panic(err)
	}
	/*var wrt io.Writer
	var rdr io.Reader
	wrt = conn
	rdr = conn
	wrt.Write([]byte(""))
	rdr.Read([]byte(""))*/
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	//req, err := http.NewRequest("GET", "http://example.com", nil)
	//req.Write(conn)
	io.Copy(os.Stdout, conn)
}

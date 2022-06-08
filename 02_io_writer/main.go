package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"

	//"net"
	"compress/gzip"
	"net/http"
	"os"

	//"strings"
	"encoding/json"
	"time"
)

// ctrl + / ：でコメントアウト/コメントアウト解除

func main() {
	/* 2.4.1 */
	//file, err := os.Create("test.txt")
	//if err != nil {
	//	panic(err)
	//}
	// file.Write([]byte("os.File example\n"))
	// file.Close()

	/* 2.4.2 */
	os.Stdout.Write([]byte("os.Stdout example\n"))

	/* 2.4.3 */
	// var buffer bytes.Buffer
	// buffer.Write([]byte("bytes.Buffer example \n"))
	// fmt.Println(buffer.String())

	// io.WriteString(&buffer, "bytes, Buffer example\n")

	/* 2.4.4 */
	// var builder strings.Builder
	// builder.Write([]byte("strings.Builder example\n")) //p.21 「Goのバイト列と文字列"」を参照
	// fmt.Println(builder.String())

	/* 2.4.5 - 1 */
	//conn, err := net.Dial("tcp", "example.com:80")
	// if err != nil {
	// 	panic(err)
	// }
	/*var wrt io.Writer
	var rdr io.Reader
	wrt = conn
	rdr = conn
	wrt.Write([]byte(""))
	rdr.Read([]byte(""))*/
	//io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	//req, err := http.NewRequest("GET", "http://example.com", nil)
	//req.Write(conn)
	//io.Copy(os.Stdout, conn)

	//2.4.5 - 2
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	// 2.4.6
	//writer := io.MultiWriter(file, os.Stdout)
	//io.WriteString(writer, "io.MultiWriter example \n")
	file, err := os.Create("test.txt.gz") // zipファイル名
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test1.txt" // 展開後のファイル名
	io.WriteString(writer, "gzip.Writer example 1234567890\n")
	writer.Close()

	buffer := bufio.NewWriterSize(os.Stdout, 4) // 引数はio.writer
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example \n")
	buffer.Flush()

	// 2.4.7
	fmt.Fprintf(os.Stdout, "Write with %v at %v\n", "os.Stdout", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", " ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello": "world",
	})

	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)

}

func handler(w http.ResponseWriter, r *http.Request) {
	//2.4.5 - 2
	io.WriteString(w, "http.ResposeWriter sample")
}

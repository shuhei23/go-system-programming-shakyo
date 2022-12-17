package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	go func() {
		HttpServer() // サーバー立てとく
	}()

	HttpClient()
}

func HttpServer() {
	listener, err := net.Listen("tcp", "localhost:8888") // サーバー起動
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept() // クライアント受付
		if err != nil {
			panic(err)
		}

		go func() { // go routineによる並列処理（非同期）：クライアントごとのメッセージ受付
			fmt.Printf("Accept %v\n", conn.RemoteAddr()) // クライアントのIP？

			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))
			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn) // ソケットに書き込み
			conn.Close()         // クライアントと通信切断
		}()
	}
}

func HttpClient() {
	conn, err := net.Dial("tcp", "localhost:8888") // Listenしてるサーバーとセッション成立したら帰ってきます
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(
		"Get", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}

package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"

	"strings"
	"time"
)

func main() {
	go func() {
		HttpServer() // サーバー立てとく
	}()

	HttpClient()
}

// クライアントはgzipを受け入れ可能か？
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), // 例：zip,7zip,gzip の時は2が返る
		"gzip") != -1
}

// 1セッションの処理
func proccessSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr()) // クライアントのIP？
	defer conn.Close()                           // クライアントと通信切断(遅延実行)
	for {
		//タイムアウトを設定
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			// タイムアウトもしくはソケットクローズ時は終了
			// それ以外はエラーにする
			neterr, ok := err.(net.Error) // ダウンキャスト
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		// リクエストを表示
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		// レスポンスを書き込む
		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGZipAcceptable(request) {
			content := "Hello World(gzipped)\n"
			// コンテンツをgzip化して転送
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()

			response.Body = io.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			// gzip対応してない
			content := "Hello World\n"
			response.Body = io.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn) // ソケットに書き込み
	}
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

		// go routineによる並列処理（非同期）：クライアントごとのメッセージ受付
		go proccessSession(conn)
	}
}

func HttpClient() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil
	// リトライ用にループで全体を囲う
	for {
		var err error
		// まだコネクションを張ってない / エラーでリトライ
		if conn == nil {
			// Dialから行ってconnを初期化
			conn, err = net.Dial("tcp", "localhost:8888") // Listenしてるサーバーとセッション成立したら帰ってきます
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		request.Header.Set("Accept-Encoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		//サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}

	}
	conn.Close()
}

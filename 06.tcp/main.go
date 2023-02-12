package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

var contents = []string{
	" これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	" むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	" 中山さまというおとのさまが、おられたそうです。",
	" その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	" ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	" そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

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
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content) // データサイズとバイト数分のデータブロック
		}
		fmt.Fprintf(conn, "0\r\n\r\n") // 通信完了はサイズ0を渡す

		// response := http.Response{
		// 	StatusCode: 200,
		// 	ProtoMajor: 1,
		// 	ProtoMinor: 1,
		// 	Header:     make(http.Header),
		// }

		// if isGZipAcceptable(request) {
		// 	content := "Hello World(gzipped)\n"
		// 	// コンテンツをgzip化して転送
		// 	var buffer bytes.Buffer
		// 	writer := gzip.NewWriter(&buffer)
		// 	io.WriteString(writer, content)
		// 	writer.Close()

		// 	response.Body = io.NopCloser(&buffer)
		// 	response.ContentLength = int64(buffer.Len())
		// 	response.Header.Set("Content-Encoding", "gzip")
		// } else {
		// 	// gzip対応してない
		// 	content := "Hello World\n"
		// 	response.Body = io.NopCloser(strings.NewReader(content))
		// 	response.ContentLength = int64(len(content))
		// }
		// response.Write(conn) // ソケットに書き込み
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

	// Dialから行ってconnを初期化
	conn, err := net.Dial("tcp", "localhost:8888") // Listenしてるサーバーとセッション成立したら帰ってきます
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// POSTで文字列を送るリクエストを作成
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888",
		nil)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)

	//サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 ||
		response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}
	for {
		// サイズを取得
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		size, err := strconv.ParseInt(
			string(sizeStr[:len(sizeStr)-2]), 16, 64)
		// 送信側フォーマット: fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)

		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		line := make([]byte, int(size))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf("   %d bytes: %s\n", size, string(line))

	}

	conn.Close()
}

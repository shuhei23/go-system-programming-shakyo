package main

import (
	// "bufio"
	"fmt"

	// "io"
	// "strings"

	// "net"
	// "net/http"
	"bytes"
	// "os"
	"encoding/binary"
)

func main() {
	/* 3.3.2 - 1 */
	// var reader io.Reader = strings.NewReader("テストデータ")
	// var readClose io.ReadCloser = io.NopCloser(reader)

	//print(readClose)
	/* 3.3.2 - 2 */
	// var reader = bufio.NewReader(strings.NewReader("reader"))
	// var writer = bufio.NewWriter(os.Stdout)
	// var readWriter io.ReadWriter = bufio.NewReadWriter(reader, writer)
	// print(readWriter)

	/* 3.4.1 */
	// for {
	//	buffer := make([]byte, 5)
	// 	size, err := os.Stdin.Read(buffer)
	// 	if err == io.EOF {
	// 		fmt.Println("EOF")
	// 		break
	// 	}
	// 	fmt.Printf("size = %d input = '%s'\n", size, string(buffer))
	// }

	// go run main.go < main.go
	// PowerShellでやると "<"が予約語でエラーになる
	// コマンドプロンプトでもなぜか以下のエラーが表示されることがあった。
	// main.go:1:1: expected 'package', found 'EOF'

	/* 3.4.2 */
	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close() // defer:現在のスコープが終了したら、その後ろに書かれている行の処理を実行します
	// io.Copy(os.Stdout, file)

	/* 3.4.3 */
	// conn, err := net.Dial("tcp", "example.com:80")
	// if err != nil {
	// 	panic(err)
	// }
	// conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	// res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// //ヘッダーを表示
	// fmt.Println("----------Header-----------")
	// fmt.Println(res.Header)
	// //ボディーを表示。最後にClose()
	// defer res.Body.Close()
	// fmt.Println("----------Body-----------")
	// io.Copy(os.Stdout, res.Body)

	/* 3.4.4 */
	// // 空のバッファ
	// var buffer1 bytes.Buffer
	// // バイト列で初期化
	// buffer2 := bytes.NewBuffer([]byte{0x41, 0x48, 0x4f})
	// // 文字列で初期化
	// buffer3 := bytes.NewBufferString("初期文字列")
	// fmt.Printf("%s\n", buffer1.String())
	// fmt.Printf("%s\n", buffer2.String())
	// fmt.Printf("%s\n", buffer3.String())

	// // bytes.Readerはbytes.NewReaderで作成
	// bReader1 := bytes.NewReader([]byte{0x41, 0x48, 0x4f, 0x0a})
	// bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャストして設定\n"))

	// // strings.Readerはstrings.NewReader()関数で作成
	// sReader := strings.NewReader("Readerの出力内容は文字列で渡す\n")

	// io.Copy(os.Stdout, bReader1)
	// io.Copy(os.Stdout, bReader2)
	// io.Copy(os.Stdout, sReader)

	/* 3.5.1 */
	/* 指定バイト数読みだす／指定位置から指定バイト数読みだす */
	// reader := strings.NewReader("Example of io.SectionReader\n kyouha iitenki.\n")
	// lReader := io.LimitReader(reader, 16)
	// sectionReader := io.NewSectionReader(reader, 32, 7)
	// io.Copy(os.Stdout, lReader)
	// print("\n")
	// io.Copy(os.Stdout, sectionReader)

	/* 3.5.2 */
	// 32ビットのビッグエンディアンのデータ(10000)
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	//エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i) // BigEndianのデータに変換してください
	fmt.Printf("data: %d\n", i)
}

package main

import (
	"net/http"
	//"archive/zip"
	"bytes"
	"encoding/binary"

	//"flag"
	"fmt"
	"hash/crc32"

	//"encoding/binary"
	//"fmt"
	//"hash/crc32"

	"io"
	"os"
	"strings"
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

	// /* 3.5.2 */
	// // 32ビットのビッグエンディアンのデータ(10000)
	// data := []byte{0x0, 0x0, 0x27, 0x10}
	// var i int32
	// //エンディアンの変換
	// binary.Read(bytes.NewReader(data), binary.BigEndian, &i) // BigEndianのデータに変換してください
	// fmt.Printf("data: %d\n", i)

	// /* 3.5.3 */
	// file, err := os.Open("PNG_transparency_demonstration_secret.png")
	// // 226,933 バイト
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// chunks := readChunks(file)
	// for _, chunk := range chunks {
	// 	dumpChunk(chunk)
	// }

	/* 3.5.4 */
	// file, err := os.Open("PNG_transparency_demonstration_1.png") // 226,933 バイト
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// newFile, err := os.Create("PNG_transparency_demonstration_secret.png") // 226,959 バイト
	// if err != nil {
	// 	panic(err)
	// }
	// defer newFile.Close()

	// chunks := readChunks(file)
	// io.WriteString(newFile, "\x89PNG\r\n\x1a\n")  // シグニチャ
	// io.Copy(newFile, chunks[0])                   // IHDR
	// io.Copy(newFile, textChunk("Lambda Notes++")) // 秘密のデータはあと
	// // 残りのデータ
	// for _, chunk := range chunks[1:] {
	// 	io.Copy(newFile, chunk)
	// }

	/* 3.6.1 */
	// reader := bufio.NewReader(strings.NewReader(source))
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	fmt.Printf("%#v\n", line)
	// 	fmt.Printf("%v\n", line)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	// scanenr := bufio.NewScanner(strings.NewReader(source))
	// scanenr.Split(bufio.ScanWords) // スペース区切り
	// scanenr.Split(bufio.ScanBytes)
	// scanenr.Split(bufio.ScanLines)
	// scanenr.Split(bufio.ScanRunes)

	// for scanenr.Scan() {
	// 	fmt.Printf("%#v\n", scanenr.Text())
	// }

	/* 3.6.2 */
	// reader := strings.NewReader(source)
	// var i int
	// var f, g float64
	// var s string
	// // fmt.Fscan(reader, &i, &f, &g, &s)
	// // fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
	// fmt.Fscanf(reader, "%v, %v, %v, %v", &s, &f, &g, &i)
	// fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", s, f, g, i)

	/* 3.6.3 */
	// reader := strings.NewReader(csvSource)
	// csvReader := csv.NewReader(reader)
	// line, _ := csvReader.ReadAll()
	// fmt.Println(line)
	// for {
	// 	line, err := csvReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(line[2], line[6:9])
	// }

	/* 3.7 */
	// header := bytes.NewBufferString("----- HEADER -----\n")
	// content := bytes.NewBufferString("Example of io.MultiReader\n")
	// footer := bytes.NewBufferString("----- FOOTER -----\n")

	//reader := io.MultiReader(header, content, footer)
	//io.Copy(os.Stdout, reader)

	// var buffer bytes.Buffer
	// reader := bytes.NewBufferString("Example of io.TeeReader")
	// teeReader := io.TeeReader(reader, &buffer) /* teeReaderにreaderに読み出しながら、bufferにも書き出す */

	// _, _ = io.ReadAll(teeReader) /* データを読み捨てる */
	// fmt.Println(buffer.String()) /* けど、バッファに残ってる */

	/* io.pip() はWrite()/Read()が呼ばれた時点でスレッドが待ち状態になる。
	   シングルスレッドだとこの待ち状態でデッドロックになるので使用してはいけない */

	/* Q3.1 */
	//-----OS.Argsで文字受け取る
	// go run main.go -?hogehoge

	// os.Argsのlenを表示
	// fmt.Println("count:", len(os.Args))

	// for i, v := range os.Args {
	// 	fmt.Printf("args[%d] -> %s\n", i, v)
	// }

	// // receivedString := //OS.Args ...

	// var s = flag.String("str", "default message.", "append message") /* stringの引数を指定 */

	// flag.Parse() /* 引数を分解 */

	// // ------
	// oldFile, err := os.Open("old.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer oldFile.Close()
	// newFile, err := os.Create("new.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer newFile.Close()
	// //newFile := os.NewFile()
	// io.Copy(newFile, oldFile)
	// newFile.WriteString(*s)

	/* Q3.2 */
	// reader := rand.Reader
	// newFile, err := os.Create("Q3_2.txt")
	// if err != nil {
	// 		panic(err)
	// }
	// lReader := io.LimitReader(reader, 1024)
	// io.Copy(newFile, lReader)

	/* Q3.3 zipファイル内の各ファイルにstrings.Reader()を使ってデータを書き込む */
	// newFile, err := os.Create("Q3_3.zip")
	// if err != nil {
	// 	panic(err)
	// }
	// zipWriter := zip.NewWriter(newFile)
	// defer zipWriter.Close()

	// writer, _ := zipWriter.Create("newfile.txt")
	// str := strings.NewReader("imai tomoaki")
	// io.Copy(writer, str)

	/* Q3.4 */

	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	/* Q3.5 */
	// str := strings.NewReader("Q3.5 read file")
	// file, _ := os.Create("Q3_5.txt")
	// newCopyN(file, str, 20)

	/* Q3.6 */
	var (
		computer    = strings.NewReader("COMPUTER")
		system      = strings.NewReader("SYSTEM")
		programming = strings.NewReader("PROGRAMMING")
	)
	var stream io.Reader
	aReader := io.NewSectionReader(programming, 5, 1)
	sReader := io.NewSectionReader(system, 0, 1)
	cReader := io.NewSectionReader(computer, 0, 1)
	iReader1 := io.NewSectionReader(programming, 8, 1)
	iReader2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(aReader, sReader, cReader, iReader1, iReader2)
	io.Copy(os.Stdout, stream) /* ASCIIを表示する */
}

func newCopyN(w io.Writer, r io.Reader, size int) {
	buffer := make([]byte, size)
	r.Read(buffer)
	w.Write(buffer)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
	/* Content-Disposition: ダウンロードしてローカルに保存 */
}

// var source = `1行め 1行め2部 1行め3部
// 2行め 2行め2部 2行め3部
// 3行め 3行め2部 3行め3部
// `

var source = "123 1.234 1.0e4 test"

var csvSource = `13101,"100 ","1000003"," ﾄｳｷｮｳﾄ "," ﾁﾖﾀﾞｸ "," ﾋﾄﾂﾊﾞｼ (1 ﾁｮｳﾒ )"," 東京都 "," 千代田区 "," 一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101 ","1010003"," ﾄｳｷｮｳﾄ "," ﾁﾖﾀﾞｸ "," ﾋﾄﾂﾊﾞｼ (2 ﾁｮｳﾒ )"," 東京都 "," 千代田区 "," 一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100 ","1000012"," ﾄｳｷｮｳﾄ "," ﾁﾖﾀﾞｸ "," ﾋﾋﾞﾔｺｳｴﾝ "," 東京都 "," 千代田区 "," 日比谷公園 ",0,0,0,0,0,0
13101,"102 ","1020093"," ﾄｳｷｮｳﾄ "," ﾁﾖﾀﾞｸ "," ﾋﾗｶﾜﾁｮｳ "," 東京都 "," 千代田区 "," 平河町 ",0,0,1,0,0,0
13101,"102 ","1020071"," ﾄｳｷｮｳﾄ "," ﾁﾖﾀﾞｸ "," ﾌｼﾞﾐ "," 東京都 "," 千代田区 "," 富士見 ",0,0,1,0,0,0`

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer) // type
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)

	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText) // data
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer

	/* bufferにlengthを書き込む */
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	// CRC計算とバッファ書き込みを同時に行う
	writer := io.MultiWriter(&buffer, crc)
	/* buffer, crcにtypeを書き込む */
	io.WriteString(writer, "teXt") // 2バイト目の5ビット目を立てるとプライベート(オリジナル)
	// A 0x41(0100 0001) ~ Z 0x5a(0101 1010)
	// a 0x61(0110 0001) ~ z 0x7a(0111 1010)
	/* buffer, crcにdataをかきこむ */
	writer.Write(byteText)
	// bufferにチェックサムを書き込む
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func readChunks(file *os.File) []io.Reader {
	// length(4 bytes) + type(4 bytes) + data(length bytes) + CRC(4 bytes)

	//チャンクを格納する配列
	var chunks []io.Reader

	//最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length) // 長さを読みだす
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(4+4+length+4)))
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(4バイト) + データ長(length) + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(4+length+4), 1)
	}
	return chunks
}

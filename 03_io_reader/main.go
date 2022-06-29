package main

import (
	"io"
	"strings"
)

func main() {
	var reader io.Reader = strings.NewReader("テストデータ")
	var readClose io.ReadCloser = io.NopCloser(reader)

	print(readClose)
}

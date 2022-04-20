package main

import (
	"fmt"
)

type Talker interface {
	Talk()
	// Talking()を入れるとGreeterで未定義なので、talker = &greeterの行でコンパイルエラーになる。
}

type Speaker interface {
	Speak()
}

type Greeter struct {
	name string
}

type Roarer struct {
	name2 string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func (g Greeter) Speak() {
	fmt.Printf("Hey Bro,  my name is %s\n", g.name)
}

func (r Roarer) Talk() {
	fmt.Printf("Hello, I'm %s\n", r.name2)
}

func main() {
	var talker Talker
	var speaker Speaker
	var greeter = Greeter{"hogeo"}
	var roarer = Roarer{"bowbow"}

	speaker = &greeter
	talker = &greeter
	talker.Talk()
	speaker.Speak()

	//再代入すると呼び分けられる
	talker = &roarer
	talker.Talk()
}
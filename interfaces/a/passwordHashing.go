package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (x *XEncoder) Encode(value string) {
	fmt.Println("Encoding by xencoder")
}

func (x *XEncoder) Decode(value string) {
	fmt.Println("Decoding by xencoder")
}

func (x *YEncoder) Encode(value string) {
	fmt.Println("Encoding by Yencoder")
}

func (x *YEncoder) Decode(value string) {
	fmt.Println("Decoding by Yencoder")
}

func main() {
	var encoder IEncoder = &XEncoder{}

	encoder.Encode("123456")
}

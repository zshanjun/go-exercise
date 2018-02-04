package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hi你好"
	//底层
	for _, v := range []byte(s) {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	//转化为int32，四个字节
	for i, v := range s {
		fmt.Printf("(%d %x) ", i, v)
	}
	fmt.Println()

	//可以自己手工转化为rune
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//推荐使用rune转化
	for k, v := range []rune(s) {
		fmt.Printf("(%d %c) ", k, v)
	}
}

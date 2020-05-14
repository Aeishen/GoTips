package main

import (
	"bytes"
	"fmt"
	"testing"
)

/*
串联字符串:
	当需要对一个字符串进行频繁的操作时，谨记在 go 语言中字符串是不可变的（类似 java 和 c#）。
	使用诸如 a += b 形式连接字符串效率低下，尤其在一个循环内部使用这种形式。这会导致大量的内
	存开销和拷贝。应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中
*/

var s = "a"


func getNextString()string{
	return s
}

func connectByBuffer(isF bool){
	var buf bytes.Buffer
	var all string
	for i := 0; i < 100; i++{
		buf.WriteString(getNextString())
	}

	all = buf.String()
	
	if isF{
		fmt.Println(all)
	}

}

func connectByAdd(isF bool){
	var all string
	for i := 0; i < 100; i++{
		all += getNextString()
	}

	if isF{
		fmt.Println(all)
	}
	
}


func BenchmarkByConn(b *testing.B)  {
	for i := 0; i <b.N; i++{
		connectByBuffer(false)
	}
}



func BenchmarkByAdd(b *testing.B)  {
	for i := 0; i <b.N; i++{
		connectByAdd(false)
	}
}



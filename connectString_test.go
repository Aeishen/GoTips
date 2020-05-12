package main

import (
	"bytes"
	"fmt"
	"testing"
)

/*
串联字符串:
	创建一个 buffer，通过 buffer.WriteString(s) 方法将字符串 s 追加到后面，最后再通过 buffer.String() 方法转换为 string,
	这种实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
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



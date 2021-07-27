package main

import (
	"bytes"
	"fmt"

	//"github.com/pkg/profile"
	"strings"
)

func concat(n int) string {
/*	s := ""
	for i := 0; i < n;i++ {
		s += randomString(n)
	}
	return s
	*/

	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}

	return s.String()
}

func main(){
/*	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := Generate(n)
		BubbleSort(nums)
		n *= 10
	}*/

	//defer profile.Start().Stop()
	//concat(10)

	var buf bytes.Buffer
	buf.WriteString("aaaaaaaaa")
	buf.WriteString("bbbbbbbbb")
	fmt.Println(buf.String())
}



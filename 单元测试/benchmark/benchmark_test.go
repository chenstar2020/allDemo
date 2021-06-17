package benchmark

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

//测试命令 go test -benchmem -bench .
func BenchmarkHello(b *testing.B){
	//耗时操作
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkParallel(b *testing.B){
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB){
		var buf bytes.Buffer
		for pb.Next(){
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}



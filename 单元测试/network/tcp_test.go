package network

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleError(t *testing.T, err error){
	t.Helper()
	if err != nil{
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T){
	ln, err := net.Listen("tcp", "127.0.0.1:0")  //监听服务
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)   //处理请求
	go http.Serve(ln, nil)       //启动服务

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world"{
		t.Fatal("expected hello world, but got", string(body))
	}
}

func TestHttp(t *testing.T){
	 req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	 w := httptest.NewRecorder()
	 helloHandler(w, req)
	 bytes, _ := ioutil.ReadAll(w.Result().Body)

	 if string(bytes) != "hello world"{
	 	t.Fatal("expected hello world, but got", string(bytes))
	 }
}
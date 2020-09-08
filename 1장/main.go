package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "hellow\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening: 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

/*
crul으로 요청
curl --http1.0 http://localhost:18888/greeting
헤더 전송
curl --http1.0 -H "X-Test: Hello" http://localhost:18888
--user-agent(요청한 기기) 단축형은 -A
curl --http1.0 -A "Mozilla/5.0" http://localhost:18888
--request 메소드 or -X 메소드
curl --http1.0 -X POST http://localhost:18888
curl --http1.0 --request POST http://localhost:18888

데이터전송
curl -d "{'hello': 'world'}" -H "Content-Type: application/json" http://localhost:18888
파일로 전송
curl -d @test.json -H "Content-Type: application/json" http://localhost:18888
*/

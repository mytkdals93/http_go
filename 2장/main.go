package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "두번째 방문")
	} else {
		fmt.Fprintf(w, "hellow\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening: 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

/*
쿠키란 웹사이트의 정보를 브라우저 쪽에 저장하는 작은 파일
쿠키는 쉽게 삭제될 수 있으니 사라지더라도 문제없는 정보나 서버 정보로 복원 할 수 있는 자료를 저장하는 데 적합

*/

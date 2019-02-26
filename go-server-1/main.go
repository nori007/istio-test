package main

import (
	"fmt"
	"io/ioutil"
"net/http"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, req *http.Request){
		//req.data..... token check?
		fmt.Println("auth start")
		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			str := string(respBody)
			println(str)
		}

		// cors !!!!!!
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":9992", nil)
}

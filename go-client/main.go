package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request){
		//req.data..... token check?
		//respBody, err := ioutil.ReadAll(req.Body)
		//if err == nil {
		//	str := string(respBody)
		//	println(str)
		//}
		//
		//// cors !!!!!!
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Write([]byte("ok"))
		newreq, _ := http.NewRequest("POST", "http://127.0.0.1:9998/api", req.Body)
		client := &http.Client{}
		resp, err := client.Do(newreq)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		w.Write(body)

	})

	http.ListenAndServe(":9991", nil)
}

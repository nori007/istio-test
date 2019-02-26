package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	http.HandleFunc("/auth", func(w http.ResponseWriter, req *http.Request){
		newreq, _ := http.NewRequest("POST", "http://go-server-1:9992/go-server-1/", req.Body)
		//newreq, _ := http.NewRequest("POST", "http://127.0.0.1:9992/auth", req.Body)
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
	http.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request){
		newreq, _ := http.NewRequest("POST", "http://go-server-2:9993/go-server-2/", req.Body)
		//newreq, _ := http.NewRequest("POST", "http://127.0.0.1:9993/api", req.Body)
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

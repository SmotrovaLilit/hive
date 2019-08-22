package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	headers := ""
	for k, v := range r.Header {
		headers = headers + fmt.Sprintf("Header field %v, Value %v\n", k, v)
	}
	cookies := ""
	for k, v := range r.Cookies() {
		cookies = cookies + fmt.Sprintf("Cookie field %v, Value %v\n", k, v)
	}
	loginfo := fmt.Sprintf("path: %s\n"+
		"method: %s\n"+
		"query: %s\n"+
		"body: %s\n"+
		"headers: %s\n"+
		"cookies: %s\n",
		r.URL.Path,
		r.Method,
		r.URL.RawQuery,
		body,
		headers,
		cookies)
	fmt.Fprintf(w, loginfo)
	log.Println(loginfo)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

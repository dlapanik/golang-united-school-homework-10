package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// | GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
func nameHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.RequestURI, "/name/")
	fmt.Fprintf(w, "Hello, %s!", parts[1])
}

// | GET    | `/bad`                                | Status: `500`                 |
func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

// | POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
func dataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "got error:%s", err)
	} else {
		fmt.Fprintf(w, "I got message:\n%s", string(body))
	}
}

// | GET    | `/header` + Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |
func headerHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))

	w.Header().Set("a+b", strconv.Itoa(a+b))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/name/", nameHandler)
	http.HandleFunc("/bad", badHandler)
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/", okHandler)

	Start("", 8080)
}

func Start(host string, port int) {
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
		log.Fatal(err)
	}
}

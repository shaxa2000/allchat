package web

import (
	"flag"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var addr = flag.String("addr", ":8080", "http service address")

func Router() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", viewHandler)
	mux.HandleFunc("/chatPage", chatHandler)
	mux.HandleFunc("/chatPage/chatCreate", chatCreate)

	// stripPrefix := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	// mux.Handle("/static", stripPrefix)

	// err := http.ListenAndServe(fmt.Sprintf(":%s", "5000"), mux)
	err := http.ListenAndServe(*addr, mux)
	checkErr(err)
}

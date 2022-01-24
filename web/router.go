package web

import (
	"fmt"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Router() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", viewHandler)
	mux.HandleFunc("/chatPage", chatHandler)
	mux.HandleFunc("/chatPage/chatCreate", chatCreate)

	// stripPrefix := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	// mux.Handle("/static", stripPrefix)

	err := http.ListenAndServe(fmt.Sprintf(":%s", "5000"), mux)
	checkErr(err)
}

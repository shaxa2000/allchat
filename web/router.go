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

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/", viewHandler)
	mux.HandleFunc("/chatPage", chatHandler)
	mux.HandleFunc("/chatPage/chatCreate", chatCreate)

	err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), mux)
	checkErr(err)
}

package web

import (
	"html/template"
	"net/http"
)

type Chat struct {
	History []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./index.html")
	checkErr(err)
	err = temp.Execute(w, nil)
	checkErr(err)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	ChatHistory := Chat{
		History: GetAllData("./db/dataBase.txt"),
	}

	//set all paths
	files := []string{
		"./templates/chatPage.html",
		"./templates/chatBase_template.html",
		"./templates/footer_template.html",
		"./templates/header_template.html",
	}
	temp, err := template.ParseFiles(files...)
	checkErr(err)
	err = temp.Execute(w, ChatHistory)
	checkErr(err)
}

func chatCreate(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("textMessage")
	user := r.FormValue("userNames")

	SetData(user, message)

	http.Redirect(w, r, "/chatPage", http.StatusSeeOther)

}

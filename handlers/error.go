package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	Num  int
	Text string
}

func Err(res http.ResponseWriter, req *http.Request, err int) {
	temp, er := template.ParseFiles("docs/htmlTemplates/error.html")
	if er != nil {
		log.Fatal(er)
		return
	}
	res.WriteHeader(err)
	errData := ErrorData{Num: err}
	if err == 404 {
		errData.Text = "Page Not Found"
	} else if err == 400 {
		errData.Text = "Bad Request"
	} else if err == 500 {
		errData.Text = "Internal Server Error"
	} else if err == 405 {
		errData.Text = "Method Not Allowed"
	}
	fmt.Println(errData)
	temp.Execute(res, errData)
}

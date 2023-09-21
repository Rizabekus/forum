package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("./ui/html/error.html")
	if err != nil || tmp == nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
	var Err internal.ErrorStruct
	Err.Message = http.StatusText(status)
	Err.Status = status
	err = tmp.Execute(w, Err)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}
}

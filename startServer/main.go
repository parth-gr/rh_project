package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	studentdata "github.com/parth-gr/RH_PROJECT/studentData"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../template.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)                      //router and function
	http.HandleFunc("/getData", studentdata.GetData) //from studenData package exported the functions
	http.HandleFunc("/updateData", studentdata.UpdateData)
	http.HandleFunc("/changeValue", studentdata.UpdateDataValue)
	fmt.Println("server starting")
	log.Fatal(http.ListenAndServe(":4000", nil)) //port number
}

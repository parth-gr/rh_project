package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/parth-gr/RH_PROJECT/studentData"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../template.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index) //router and function
	http.HandleFunc("/getData", studentData.GetData)
	http.HandleFunc("/updateData", studentData.UpdateData)
	http.HandleFunc("/changeValue", studentData.UpdateDataValue)
	fmt.Println("server starting")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

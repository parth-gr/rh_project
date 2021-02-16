package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
    studentdata "github.com/parth-gr/rh_project/studentData"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)                              // debug output
		w.WriteHeader(http.StatusInternalServerError) // proper HTTP response
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)                      // router and function
	http.HandleFunc("/getData", studentdata.GetData) // from studenData package exported the functions
	http.HandleFunc("/updateData", studentdata.UpdateData)
	http.HandleFunc("/changeValue", studentdata.UpdateDataValue)
	fmt.Println("server starting")
	go func() {
		log.Fatal(http.ListenAndServe(":5000", nil)) // Start Server at https://localhost:5000 
	}()
	log.Fatal(http.ListenAndServe(":5001", nil)) // another port number
}



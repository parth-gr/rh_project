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
		fmt.Println(err)                              // Ugly debug output
		w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)                      //router and function
	http.HandleFunc("/getData", studentdata.GetData) //from studenData package exported the functions
	http.HandleFunc("/updateData", studentdata.UpdateData)
	http.HandleFunc("/changeValue", studentdata.UpdateDataValue)
	fmt.Println("server starting")
	go func() {
		log.Fatal(http.ListenAndServe(":5000", nil)) //port number
	}()
	log.Fatal(http.ListenAndServe(":5001", nil))
}

//run docker file
//  docker build -t RH_PROJECT .
// docker image ls
//  docker run -p 4000:4000 -tid rh_project
// FROM golang:latest

// RUN mkdir /app
// ADD . /app
// WORKDIR /app

// RUN go build -o main .

// EXPOSE 4000

// ENTRYPOINT ["/app/main"]

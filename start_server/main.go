package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("../template.html")
	t.Execute(w, nil)

	// fmt.Fprintf(w, "hello")

	// fmt.Fprintf(w, "choose data of student")
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "choose data of student")

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Rollno := r.FormValue("first")
	fmt.Println(Rollno)

	type Student []struct {
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Job    string   `yaml: "job"`
		Skills []string `yaml: "skills"`
	}

	student1 := &Student{}

	source, err := ioutil.ReadFile("../student_data.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &student1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(student1)
	t, _ := template.ParseFiles("../template.html")
	t.Execute(w, student1)

}

func main() {
	http.HandleFunc("/", index) //router and function
	http.HandleFunc("/getData", getData)
	fmt.Println("server starting")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

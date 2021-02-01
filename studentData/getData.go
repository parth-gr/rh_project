package studentData

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

func GetData(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	Rollno := r.FormValue("first")

	type Student struct {
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Job    string   `yaml: "job"`
		Skills []string `yaml: "skills"`
	}

	studentsData := &[]Student{}

	source, err := ioutil.ReadFile("../student_data.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}

	rollno, err := strconv.Atoi(Rollno)

	if len(*studentsData) < rollno || Rollno == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println(err)
	} else {
		rollno--
		t, _ := template.ParseFiles("../template.html")
		t.Execute(w, (*studentsData)[rollno])
	}
	uri := r.URL.String()
	method := r.Method
	fmt.Println(uri, method)
}

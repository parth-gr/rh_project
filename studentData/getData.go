package studentdata

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

func GetData(w http.ResponseWriter, r *http.Request) { //GetData capital as it need to be exported

	if r.Method != "POST" { //if the method is GET so redirect it back to the Home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	Rollno := r.FormValue("first") //extract the input from the user

	type Student struct { //Structure format for reading the data from YAML file(student_data.yaml)
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Age    string   `yaml: "age"`
		Class  string   `yaml: "class"`
		Skills []string `yaml: "skills"`
	}

	studentsData := &[]Student{} //reference object of the structure Student

	source, err := ioutil.ReadFile("student_data.yaml") //red data from Yaml file
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData) //unmarsaling and passing refrence object
	if err != nil {
		log.Println(err)
	}

	rollno, err := strconv.Atoi(Rollno) //converting rollno string to integer

	if len(*studentsData) < rollno || Rollno == "" { //checking if roll no is not excedding the range
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println(err)
	} else {
		rollno--
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, (*studentsData)[rollno])
	}
	uri := r.URL.String() //printing the log output
	method := r.Method
	fmt.Println(uri, method)
}

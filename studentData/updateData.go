package studentdata

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Student struct { //structure format for reading and writing in yaml
	Name   string   `yaml: "name"`
	Roll   string   `yaml: "roll"`
	Age    string   `yaml: "age"`
	Class  string   `yaml: "class"`
	Skills []string `yaml: "skills"`
}

func UpdateDataValue(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name") //data inputed by user
	age := r.FormValue("age")
	class := r.FormValue("class")
	roll := r.FormValue("rollno")   //this is the hidden field data , updated by UpdateData route while rerouting towards UpdateDataValue
	skills := r.FormValue("skills") // UpdateData recieved it from the home page input field

	studentsData := &[]Student{}

	source, err := ioutil.ReadFile("student_data.yaml") //reading yaml data
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}

	rollno, err := strconv.Atoi(roll) //string to int conversion

	rollno--

	if name != "" { //updating data
		(*studentsData)[rollno].Name = name
	}
	if age != "" {
		(*studentsData)[rollno].Age = age
	}
	if class != "" {
		(*studentsData)[rollno].Class = class
	}
	if skills != "" {
		skillsarray := strings.Split(skills, ",") //making array of Skills data seprated by ','
		fmt.Println(skillsarray)
		for _, skill := range skillsarray {

			(*studentsData)[rollno].Skills = append((*studentsData)[rollno].Skills, skill)

		}

	}

	d, err := yaml.Marshal(&studentsData) //writing back to yaml file
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile("student_data.yaml", d, 0644)

	http.Redirect(w, r, "/", http.StatusSeeOther) //redirecting to the home page

	uri := r.URL.String() //for logging request and response
	method := r.Method
	fmt.Println(uri, method)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" { //if the method is get then redirect to home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	studentsData := &[]Student{}

	source, err := ioutil.ReadFile("student_data.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}

	Rollno := r.FormValue("first") //user input data of rollno
	rollno, err := strconv.Atoi(Rollno)

	if len(*studentsData) < rollno || (Rollno == "") {
		http.Redirect(w, r, "/", http.StatusSeeOther)

		log.Println(err)
	} else {
		type rolldata struct {
			RollNo string
		}

		rolldata1 := rolldata{RollNo: Rollno}

		t, _ := template.ParseFiles("update.html")
		t.Execute(w, rolldata1)
	}

	uri := r.URL.String() //for logging request and response
	method := r.Method
	fmt.Println(uri, method)

}

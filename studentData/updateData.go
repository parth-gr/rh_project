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

// structure format for reading and writing in yaml
type student struct { 
	Name   string   `yaml: "name"`
	Roll   string   `yaml: "roll"`
	Age    string   `yaml: "age"`
	Class  string   `yaml: "class"`
	Skills []string `yaml: "skills"`
}
// UpdateDataValue is a route used to update the data of the student 
func UpdateDataValue(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") // data inputed by user
	age := r.FormValue("age")
	class := r.FormValue("class")
	roll := r.FormValue("rollno")   // this is the hidden field data , updated by UpdateData route while rerouting towards UpdateDataValue
	skills := r.FormValue("skills") // UpdateData recieved it from the home page input field

	studentsData := &[]student{}

	source, err := ioutil.ReadFile("student_data.yaml") //reading yaml data
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}
    rollno, err := strconv.Atoi(roll) // string to int conversion
	rollno--
    // updating data
	if name != "" { 
		(*studentsData)[rollno].Name = name
	}
	if age != "" {
		(*studentsData)[rollno].Age = age
	}
	if class != "" {
		(*studentsData)[rollno].Class = class
	}
	if skills != "" {
		skillsarray := strings.Split(skills, ",") // making array of Skills data seprated by ','
		fmt.Println(skillsarray)
		for _, skill := range skillsarray {
			(*studentsData)[rollno].Skills = append((*studentsData)[rollno].Skills, skill)
		}
	}

	d, err := yaml.Marshal(&studentsData) // writing back to yaml file
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile("student_data.yaml", d, 0644)

	http.Redirect(w, r, "/", http.StatusSeeOther) // redirecting to the home page
	uri := r.URL.String() // for logging request and response
	method := r.Method
	fmt.Println(uri, method)
}
// UpdateData is a route used to update the data of the student intermediate between the UpdateDataValue route
func UpdateData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // if the method is get then redirect to home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	studentsData := &[]student{}
	
	source, err := ioutil.ReadFile("student_data.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData)
	if err != nil {
		log.Println(err)
	}
	rollno := r.FormValue("first") // user input data of rollno
	rollnoconverted, err := strconv.Atoi(rollno)

	if len(*studentsData) < rollnoconverted || (rollno == "") || rollnoconverted == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println(err)
	} else {
		type rolldata struct {
			RollNo string
		}
        rolldata1 := rolldata{RollNo: rollno}
        t, _ := template.ParseFiles("update.html")
		t.Execute(w, rolldata1)
	}
	uri := r.URL.String() // for logging request and response
	method := r.Method
	fmt.Println(uri, method)
}

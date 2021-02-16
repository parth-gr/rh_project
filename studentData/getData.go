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

// GetData capital as it need to be exported
func GetData(w http.ResponseWriter, r *http.Request) { 
    // if the method is GET so redirect it back to the Home page
	if r.Method != "POST" { 
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	rollno := r.FormValue("first") // extract the input from the user
// structure format for reading the data from YAML file(student_data.yaml)
	type student struct { 
		Name   string   `yaml: "name"`
		Roll   string   `yaml: "roll"`
		Age    string   `yaml: "age"`
		Class  string   `yaml: "class"`
		Skills []string `yaml: "skills"`
	}

	studentsData := &[]student{} // reference object of the structure Student
	source, err := ioutil.ReadFile("student_data.yaml") // read data from Yaml file
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(source), &studentsData) // unmarsaling and passing refrence object
	if err != nil {
		log.Println(err)
	}

	rollnoconverted, err := strconv.Atoi(rollno) // converting rollno string to integer

	if len(*studentsData) < rollnoconverted || rollno == "" || rollnoconverted == 0{ // checking if roll no is not excedding the range
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println(err)
	} else {
		rollnoconverted--
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, (*studentsData)[rollnoconverted])
	}
	uri := r.URL.String() // printing the log output
	method := r.Method
	fmt.Println(uri, method)
}

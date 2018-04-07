package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// Person struct
type Person struct {
	Name  string
	Phone string
}

// GetFormDataHandler handler
func GetFormDataHandler(w http.ResponseWriter, r *http.Request) {

	html := `<h1>Contact  : </h1>

               // replace example.com to your machine domain name or localhost
               <form action="http://localhost:8080/process_form_data" method="post">
                <div>
                 <label>Name : </label>
                 <input type="text" name="name" id="name" >
                </div>
                <div>
                 <label>Phone : </label>
                 <input type="text" name="phone" id="phone" >
                </div>
                <div>
                  <input type="submit" value="Send">
                </div>
              </form>`

	w.Write([]byte(fmt.Sprintf(html)))
}

// ReadFormDataHandler handler
func ReadFormDataHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(err)
	}

	person := new(Person)
	decoder := schema.NewDecoder()

	err = decoder.Decode(person, r.PostForm)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person)

	w.Write([]byte(fmt.Sprintf("Name is %v \n", person.Name)))
	w.Write([]byte(fmt.Sprintf("Phone is %v \n", person.Phone)))
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", GetFormDataHandler)
	mx.HandleFunc("/process_form_data", ReadFormDataHandler)

	http.ListenAndServe(":8080", mx)
}

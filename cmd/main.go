package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	container "sampleGoWebProject/container"
	"sampleGoWebProject/models"
	"strconv"
)


func main() {
	c := &container.Container{}
	c.WireDependencies()
	r := mux.NewRouter()
	r.Handle("/employees", rootHandler(c))
	r.Handle("/employees/{id}", pathHandler(c))
	http.ListenAndServe(":" + c.GetAppConfig().Server.Port, r)
}

func rootHandler(container *container.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			handleGet(container, r, w)
		case "POST":
			handlePost(container, w, r)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}

func pathHandler(container *container.Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			vars := mux.Vars(r)
			id, _ := strconv.Atoi(vars["id"])
			fmt.Println("id = ", id)
			emp := container.GetEmployeeUseCase().GetEmployeeById(nil, id)
			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(emp)
			w.Write(reqBodyBytes.Bytes())
		case "DELETE":
			fmt.Fprintf(w, "DELETE is not supported yet!!")
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

	})
}

func handleGet(container *container.Container, r *http.Request, w http.ResponseWriter) {
	emps := container.GetEmployeeUseCase().FindAllEmployees()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(emps)
	w.Write(reqBodyBytes.Bytes())
}

func handlePost(container *container.Container, w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var result models.Employee
	json.Unmarshal([]byte(body), &result)
	fmt.Println(result)
	emp := container.GetEmployeeUseCase().SaveEmployee(nil, &result)
	w.Write([]byte(fmt.Sprintf("Employee Saved Successfully with Id {%d} ", emp.Id)))
}



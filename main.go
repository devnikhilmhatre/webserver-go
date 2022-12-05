package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello..")
	io.WriteString(w, "Hello..")
	return
}

func Form(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Println(name)
	io.WriteString(w, "form data is "+name)
	return
}

func FlowChartDiagram(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("flow-chart.png")
	w.Header().Add("Content-Type", "image/png")
	w.Write(buf)
	return
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/form", Form)
	http.HandleFunc("/flow-chart-diagram", FlowChartDiagram)
	fmt.Println("Server started at port :8080")
	http.ListenAndServe(":8080", nil)
}

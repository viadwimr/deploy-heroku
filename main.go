package main

import "net/http"
import "fmt"
import "encoding/json"

func main() {
    http.HandleFunc("/student", ActionStudent)

    server := new(http.Server)
    server.Addr = ":9000"

    fmt.Println("server started at localhost:9000")
    server.ListenAndServe()
	/*
	    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "POST":
            w.Write([]byte("post"))
        case "GET":
            w.Write([]byte("get"))
        default:
            http.Error(w, "", http.StatusBadRequest)
        }
    })

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
	*/
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if !Auth(w, r)         { return }
    if !AllowOnlyGET(w, r) { return }

    if id := r.URL.Query().Get("id"); id != "" {
        OutputJSON(w, SelectStudent(id))
        return
    }
    
    OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
    res, err := json.Marshal(o)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}
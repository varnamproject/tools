package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/varnamproject/libvarnam-golang/libvarnam"
)

var (
	v    *libvarnam.Varnam
	lang = flag.String("lang", "ml", "Language")
)

type response struct {
	Errors []string `json:"errors"`
	Word   string   `json:"input"`
	Result []string `json:"result"`
}

func TranslitrateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	text := vars["text"]
	results := v.Transliterate(text)
	res := &response{Errors: []string{}, Word: text, Result: results}
	data, _ := json.Marshal(res)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(data)
	w.Write([]byte("\n"))
}

func LearnHandler(w http.ResponseWriter, r *http.Request) {
	rc := v.Learn(r.FormValue("text"))
	if rc != nil {
		log.Println("Error while learing word")
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	flag.Parse()
	var err *libvarnam.VarnamError
	v, err = libvarnam.Init(*lang)

	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/tl/{lang}/{text}", TranslitrateHandler)
	r.HandleFunc("/api/learn", LearnHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Println("http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

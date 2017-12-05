package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io"

	dm "github.com/JerryXie98/PathFynder/Data_Models"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, my name is %s", r.URL.Path[1:])
}

func GetGameData(w http.ResponseWriter, r *http.Request) {
	var data [10]dm.Game_Object
	for i := 0; i < 10; i++ {
		data[i] = dm.Game_Object{i,i,true}
	}

	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(data)
	io.WriteString(w, string(result))
}

func RunServer() {

	http.HandleFunc("/test", SampleHandler)
	http.HandleFunc("/run", GetGameData)
	fmt.Printf("Now serving on 8080...")
	http.ListenAndServe(":8080", nil)
}
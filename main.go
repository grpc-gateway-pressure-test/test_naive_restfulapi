package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StringMessage struct {
	Value string
}

func Echo(w http.ResponseWriter, r *http.Request) {

	buff, _ := ioutil.ReadAll(r.Body)
	request := &StringMessage{}
	json.Unmarshal(buff, request)

	result := &StringMessage{}
	result.Value = request.Value

	byte, _ := json.Marshal(result)
	fmt.Fprintf(w, string(byte))
}

func main() {
	http.HandleFunc("/v1/echo", Echo)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("ListenAndServe err", err)
	}
}

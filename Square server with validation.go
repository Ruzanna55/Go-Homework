package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/square/", squareHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func squareHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		fmt.Println("sqaureHandler called")
		//fmt.Println(r)
		array := make([]byte, r.ContentLength)
		r.Body.Read(array)
		fmt.Println("Body:", string(array))
		strBody := string(array)
		var Check = regexp.MustCompile(`^[a-z]+=+\[[0-9]+\]$`)
		Check.MatchString(strBody)
		//validate input

		split := strings.Split(strBody, "=")
		number, _ := strconv.Atoi(split[1])
		number = number * number
		w.Write([]byte(strconv.Itoa(number)))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

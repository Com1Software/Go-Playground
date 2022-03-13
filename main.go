package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/Com1Software/Go-Playground-for-Programmers/tree/main/playground"
)

func main() {
	fmt.Printf("Go Playground for Programmers (c) Copyright Com1Software 1992-2022\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//	args := os.Args

	playground.Openbrowser("http://localhost:8080/")
	fmt.Println("Server running....")
	fmt.Println("Listening on port 8080")
	fmt.Println("")
	//------------------------------------------------ Home Page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		xdata := playground.StartPage()
		fmt.Fprint(w, xdata)
	})
	//------------------------------------------------ Test Page 1
	http.HandleFunc("/testpage1", func(w http.ResponseWriter, r *http.Request) {
		xdata := playground.TestPage1()
		fmt.Fprint(w, xdata)
	})
	//-----------------------------------
	//------------------------------------------------- Start Server
	http.ListenAndServe(":8080", nil)
}

//-------------------------------------------------------------

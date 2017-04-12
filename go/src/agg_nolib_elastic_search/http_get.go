package main

import _ "github.com/go-sql-driver/mysql"

var InputData Inputs

func main() {

	InputData = ReadInputs()
	RequestMapping()

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			Agg()
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
	RequestAgg()

}

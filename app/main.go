package main

import (
	"net/http"
)

func main()  {
	fs := http.FileServer(http.Dir("../front/dist/"))
	http.Handle("/", fs)

	http.ListenAndServe(":3000", nil)
}

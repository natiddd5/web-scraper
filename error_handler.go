package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	//respond with 400 status error (client error)
	jsonResponse(w, 400, "something went wrong..")
}

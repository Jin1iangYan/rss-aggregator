package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErreo(w, 400, "Something went worng")
}

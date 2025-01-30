package webserver

import (
	//"github.com/Zarapinga/HealthTrack/infra/webserver/handlers"
	"net/http"
)

func Server() {

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

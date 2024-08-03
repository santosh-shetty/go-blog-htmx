package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {

	json.NewEncoder(res).Encode("Home Page")
}

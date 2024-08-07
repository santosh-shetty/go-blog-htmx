package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/santosh-shetty/blog/pkg/helpers"
	"github.com/santosh-shetty/blog/pkg/models"
)

type FormData struct {
	Title   string
	BaseURL string
	Message string
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := "pkg/views/backend/auth"
		layoutPath := filepath.Join(path, "layout.html")
		regFormPath := filepath.Join(path, "register.html")
		data := FormData{
			Title: "Register Form",
		}
		// Parse the templates
		temp, err := template.ParseFiles(layoutPath, regFormPath)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the layout template with the content from dashboard
		err = temp.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {

		hashPass, err := helpers.HashPassword(r.FormValue("password"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := models.User{
			FullName: r.FormValue("fullName"),
			Email:    r.FormValue("email"),
			Password: hashPass,
		}

		err = models.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Use the helper to set the flash message
		helpers.SetFlashMessage(w, r, "Registration successful, please login.")
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"redirect": "/admin/login"}
		json.NewEncoder(w).Encode(response)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := "pkg/views/backend/auth/"
		layoutPath := filepath.Join(path, "layout.html")
		regFormPath := filepath.Join(path, "login.html")
		message := helpers.GetFlashMessage(w, r)
		fmt.Println(message)
		data := FormData{
			Title:   "Login Form",
			Message: message,
		}
		// Parse the templates
		temp, err := template.ParseFiles(layoutPath, regFormPath)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the layout template with the content from dashboard
		err = temp.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// var user models.User
		user, err := models.FindUserByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		correctPass := helpers.CheckPasswordHash(password, user.Password)
		if !correctPass {
			// Use the helper to set the flash message
			w.Header().Set("Content-Type", "text/html")
			response := []byte(`<div class="alert alert-danger">Incorrect Email or Password</div>`)
			w.Write(response)
			return
		}

		tokenString, err := helpers.CreateJWTToken(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		// w.Header().Set("Content-Type", "application/json")
		// response := map[string]string{"redirect": "/admin/blog/list"}
		// json.NewEncoder(w).Encode(response)
		w.Header().Set("Content-Type", "text/html")
		response := []byte(`<div class="alert alert-success">Login Sucessfully !</div><input type="text" value="/admin/blog/list" id="redirect" hidden>`)
		w.Write(response)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	helpers.SetFlashMessage(w, r, "Logout Successully!")
	http.Redirect(w, r, "/login", http.StatusFound)
}

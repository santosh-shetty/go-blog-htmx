package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/santosh-shetty/blog/pkg/helpers"
	"github.com/santosh-shetty/blog/pkg/models"
)

func CategoryList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := helpers.NewData("Category List")
		var categories []models.Category
		categories, err := models.AllCategory()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
			return
		}

		data.Categories = categories
		helpers.ServeHtmlFile(w, "category/list.html", data)
	}

}
func AddCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := helpers.NewData("Add Category")
		helpers.ServeHtmlFile(w, "category/add.html", data)
	}
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")

		data := models.Category{
			Title:       title,
			Description: description,
		}
		err := models.AddCategory(data)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Blog Added Successfully!"})
	}

}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
		return
	}

	// For GET Method
	if r.Method == "GET" {
		data := helpers.NewData("Edit Category")
		category, err := models.CategoryById(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
			return
		}
		data.Category = category
		helpers.ServeHtmlFile(w, "category/add.html", data)
	}

	// For POSt Method
	if r.Method == "POST" {
		data := models.Category{
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		err := models.UpdateCategoryById(id, data)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Category Updated Successfully!"})
		return
	}

}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
		return
	}

	err = models.DeleteCategory(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
		return
	}
	categories, err := models.AllCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := helpers.NewData("Category Deleted")
	data.Categories = categories
	helpers.ServeHtmlComponentFile(w, "categorySearchResults.html", data)
}

package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/santosh-shetty/blog/pkg/helpers"
	"github.com/santosh-shetty/blog/pkg/models"
)

func BlogList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := helpers.NewData("Blog List")
		var blogs []models.Blog
		blogs, err := models.FindAll()
		if err != nil {
			http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
			return
		}
		data.Blogs = blogs
		helpers.ServeHtmlFile(w, "blog/list.html", data)
	}
}

func BlogAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		data := helpers.NewData("Add Blog")
		categories, err := models.AllCategory()
		if err != nil {
			http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
			return
		}
		data.Categories = categories
		// json.NewEncoder(w).Encode(data)
		helpers.ServeHtmlFile(w, "blog/add.html", data)
	}
	if r.Method == "POST" {
		// Maximum upload of 10 MB files
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get Form Data
		title := r.FormValue("title")
		shortDesc := r.FormValue("short_desc")
		description := r.FormValue("description")
		catId, err := strconv.ParseInt(r.FormValue("category"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get handler for filename, size, and headers
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		// Store Image
		imagePath, err := helpers.StoreFile(w, file, handler)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		blog := models.Blog{
			Title:       title,
			ShortDesc:   shortDesc,
			Description: description,
			Image:       imagePath,
			Category:    catId,
		}
		// Save the blog to the database
		err = models.AddBlog(blog)
		if err != nil {
			http.Error(w, "Error saving blog", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Blog Added Successfully!"})
	}
}
func EditBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "GET" {
		blog, err := models.FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := helpers.NewData("Edit Blog")
		data.Blog = blog

		categories, err := models.AllCategory()
		if err != nil {
			http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
			return
		}
		data.Categories = categories

		helpers.ServeHtmlFile(w, "blog/add.html", data)
	}
	if r.Method == "POST" {
		catId, err := strconv.ParseInt(r.FormValue("category"), 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get handler for filename, size, and headers
		file, handler, err := r.FormFile("image")
		var imagePath string
		if err != nil {
			if err == http.ErrMissingFile {
				blog, _ := models.FindById(id)
				imagePath = blog.Image
			} else {
				http.Error(w, "Error retrieving the file", http.StatusBadRequest)
				return
			}
		} else {
			defer file.Close()
			// Store Image
			imagePath, err = helpers.StoreFile(w, file, handler)
			if err != nil {
				http.Error(w, "Error saving file", http.StatusInternalServerError)
				return
			}
		}

		blog := models.Blog{
			Title:       r.FormValue("title"),
			ShortDesc:   r.FormValue("short_desc"),
			Description: r.FormValue("description"),
			Image:       imagePath,
			Category:    catId,
		}
		err = models.UpdateBlogById(id, blog)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "failed", "message": err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Blog Updated Successfully!"})
		return
	}
}
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteBlog(id)
	if err != nil {
		http.Error(w, "Error during deleting blog", http.StatusInternalServerError)
		return
	}
	data := helpers.NewData("New Data After Deleted!")
	blogs, err := models.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Blogs = blogs
	helpers.ServeHtmlComponentFile(w, "blogSearchResults.html", data)
}

func SearchBlog(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	blogs, err := models.SearchBlogbyTitle(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := helpers.NewData("Blog List Table")
	if err != nil {
		http.Error(w, "Error fetching blogs", http.StatusInternalServerError)
		return
	}
	data.Blogs = blogs
	helpers.ServeHtmlComponentFile(w, "blogSearchResults.html", data)
}

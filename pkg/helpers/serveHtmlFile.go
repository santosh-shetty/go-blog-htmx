package helpers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/santosh-shetty/blog/pkg/models"
)

type Data struct {
	Title   string
	BaseURL string
	// For Blogs
	Blogs []models.Blog
	Blog  models.Blog
	// For Category
	Categories []models.Category
	Category   models.Category
}

func NewData(title string) Data {
	return Data{
		Title:   title,
		BaseURL: "http://localhost:9000",
	}
}

// Function for serving then file without Admin Layout
func ServeHtmlFile(w http.ResponseWriter, filePath string, data Data) {
	path := "pkg/views/backend/"
	layoutPath := filepath.Join(path, "layout.html")
	dashboardPath := filepath.Join(path, filePath)
	sidebarPath := filepath.Join(path, "inc/sidebar.html")
	topbarPath := filepath.Join(path, "inc/topbar.html")

	// Parse the templates
	temp, err := template.ParseFiles(layoutPath, sidebarPath, topbarPath, dashboardPath)
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the layout template with the content from dashboard
	err = temp.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Function for serving then file without Admin Layout
func ServeHtmlComponentFile(w http.ResponseWriter, filePath string, data Data) {
	path := "pkg/views/backend/components"
	dashboardPath := filepath.Join(path, filePath)

	// Parse the templates
	temp, err := template.ParseFiles(dashboardPath)
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the layout template with the content from dashboard
	err = temp.ExecuteTemplate(w, filePath, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

package routes

import (
	"github.com/gorilla/mux"
	admin "github.com/santosh-shetty/blog/pkg/controllers/admin"
)

func Routes(router *mux.Router) {
	// router.HandleFunc("/test", controllers.Home).Methods("GET")

	//  ========= Start Admin Routes =============

	router.HandleFunc("/dashboard", admin.Dashboard).Methods("GET")

	// Blog Routes
	router.HandleFunc("/admin/blog/list", admin.BlogList).Methods("GET")
	router.HandleFunc("/admin/blog/add", admin.BlogAdd).Methods("GET", "POST")
	router.HandleFunc("/admin/blog/delete/{id}", admin.DeleteBlog).Methods("GET")
	router.HandleFunc("/admin/blog/edit/{id}", admin.EditBlog).Methods("GET")
	router.HandleFunc("/admin/blog/update/{id}", admin.EditBlog).Methods("POST")
	router.HandleFunc("/admin/blog/search", admin.SearchBlog).Methods("GET")

	// Category Routes
	router.HandleFunc("/admin/category/list", admin.CategoryList).Methods("GET")
	router.HandleFunc("/admin/category/add", admin.AddCategory).Methods("GET", "POST")
	router.HandleFunc("/admin/category/delete/{id}", admin.DeleteCategory).Methods("GET")
	router.HandleFunc("/admin/category/edit/{id}", admin.EditCategory).Methods("GET")
	router.HandleFunc("/admin/category/update/{id}", admin.EditCategory).Methods("POST")

	//  ========= End Admin Routes =============
}

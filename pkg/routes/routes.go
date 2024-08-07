package routes

import (
	"github.com/gorilla/mux"
	"github.com/santosh-shetty/blog/pkg/controllers/admin"
	"github.com/santosh-shetty/blog/pkg/middleware"
)

func Routes(router *mux.Router) {
	// router.HandleFunc("/test", controllers.Home).Methods("GET")

	//  ========= Start Admin Routes =============
	// Auth Routes
	router.HandleFunc("/login", admin.Login).Methods("GET", "POST")
	router.HandleFunc("/register", admin.Register).Methods("GET", "POST")
	router.HandleFunc("/logout", admin.Logout).Methods("GET")

	adminRoute := router.PathPrefix("/admin").Subrouter()

	adminRoute.Use(middleware.AuthMiddleware)
	adminRoute.HandleFunc("/dashboard", admin.Dashboard).Methods("GET")

	// Blog Routes
	adminRoute.HandleFunc("/blog/list", admin.BlogList).Methods("GET")
	adminRoute.HandleFunc("/blog/add", admin.BlogAdd).Methods("GET", "POST")
	adminRoute.HandleFunc("/blog/delete/{id}", admin.DeleteBlog).Methods("GET")
	adminRoute.HandleFunc("/blog/edit/{id}", admin.EditBlog).Methods("GET")
	adminRoute.HandleFunc("/blog/update/{id}", admin.EditBlog).Methods("POST")
	adminRoute.HandleFunc("/blog/search", admin.SearchBlog).Methods("GET")

	// Category Routes
	adminRoute.HandleFunc("/category/list", admin.CategoryList).Methods("GET")
	adminRoute.HandleFunc("/category/add", admin.AddCategory).Methods("GET", "POST")
	adminRoute.HandleFunc("/category/delete/{id}", admin.DeleteCategory).Methods("GET")
	adminRoute.HandleFunc("/category/edit/{id}", admin.EditCategory).Methods("GET")
	adminRoute.HandleFunc("/category/update/{id}", admin.EditCategory).Methods("POST")

	//  ========= End Admin Routes =============
}

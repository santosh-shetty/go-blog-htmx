package admin

import (
	"net/http"

	"github.com/santosh-shetty/blog/pkg/helpers"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// filePath := "blog/list.html"
	filePath := "dashboard.html"
	data := helpers.NewData("Dashboard")
	helpers.ServeHtmlFile(w, filePath, data)
}

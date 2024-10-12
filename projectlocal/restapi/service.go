package restapi

import (
	"net/http"
	"projectlocal/internal"
	"projectlocal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/getproduct", internal.GetProductWithSKU).Methods("POST")
	r.HandleFunc("/createproduct", internal.CreateProduct).Methods("POST")
	r.HandleFunc("/update", internal.UpdateProductStatus).Methods("POST")
	r.HandleFunc("/delete", internal.DeleteProduct).Methods("POST")

	r.Use(middleware.Middleware())

	http.ListenAndServe(":9999", r)
}

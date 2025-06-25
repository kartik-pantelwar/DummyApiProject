package routes

import (
	"dummyProject/external"
	userhandler "dummyProject/internal/interfaces/handler"
	"dummyProject/internal/interfaces/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(
	userHandler *userhandler.UserHandler, pHandler *external.ProductHandler) http.Handler {
	router := chi.NewRouter()

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
		r.Post("/refresh", userHandler.Refresh)
	})

	router.Route("/user", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/profile", userHandler.Profile)
		r.Post("/logout", userHandler.LogOut)
	})

	router.Route("/products", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/all", pHandler.AllProducts)
		r.Get("/{id}", pHandler.SingleProduct)
		r.Get("/category/{cg}", pHandler.CategoryProduct)
		r.Get("/category/all", pHandler.ProductCategories)
		r.Get("/category/list", pHandler.CategoryList)
		r.Get("/searchproduct/{name}", pHandler.SearchProduct)
		r.Post("/addproduct", pHandler.AddProduct)
		r.Delete("/deleteproduct/{id}", pHandler.DeleteProduct)
		r.Get("/", pHandler.Paging)
		r.Get("/sortBy={field}&order={sort}", pHandler.Sorting)
		r.Put("/updateProduct/{id}", pHandler.UpdateProduct)
	})

	return router
}

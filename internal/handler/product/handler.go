package product

import (
	"github.com/go-chi/chi"
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/service"
)

type ProductHandler struct {
	ProductService service.ProductService
	JWTMiddleware  *middleware.JWT
}

func ProvideProductHandler(productService service.ProductService, jwtMiddleware *middleware.JWT) ProductHandler {
	return ProductHandler{
		ProductService: productService,
		JWTMiddleware:  jwtMiddleware,
	}
}

func (h *ProductHandler) Router(r chi.Router) {
	r.Route("/product", func(r chi.Router) {
		r.Use(h.JWTMiddleware.VerifyToken)
		// Product Handler
		r.Post("/", h.InsertNewProduct)
		r.Get("/{id}", h.Find)
		r.Get("/", h.FindAllProductData)
		r.Put("/{id}", h.UpdateProductData)
		r.Delete("/{id}", h.DeleteProductData)
		// product customer sku
		r.Post("/customer", h.SearchSKUProduct)
		// transaction customer
		r.Post("/checkout", h.InsertNewTransaction)
		r.Get("/checkout/history", h.FindTransaction)

	})
}

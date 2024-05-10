package product

import (
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/internal/domain/product/service"

	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductService service.ProductService
	JWTMiddleware  *middleware.JWT
}

func ProvideProductHandler(productService service.ProductService, jwt *middleware.JWT) ProductHandler {
	return ProductHandler{
		ProductService: productService,
		JWTMiddleware:  jwt,
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
	})
}

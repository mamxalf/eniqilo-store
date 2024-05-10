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
		r.Group(func(r chi.Router) {
			r.Post("/customer", h.SearchSKUProduct)
		})
	})
}

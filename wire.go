//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/http"
	"github.com/mamxalf/eniqilo-store/http/middleware"
	"github.com/mamxalf/eniqilo-store/http/router"
	"github.com/mamxalf/eniqilo-store/infra"
	repositoryProduct "github.com/mamxalf/eniqilo-store/internal/domain/product/repository"
	serviceProduct "github.com/mamxalf/eniqilo-store/internal/domain/product/service"
	repositoryStaff "github.com/mamxalf/eniqilo-store/internal/domain/staff/repository"
	serviceStaff "github.com/mamxalf/eniqilo-store/internal/domain/staff/service"
	"github.com/mamxalf/eniqilo-store/internal/handler/health"
	"github.com/mamxalf/eniqilo-store/internal/handler/product"
	"github.com/mamxalf/eniqilo-store/internal/handler/staff"
)

var configurations = wire.NewSet(
	config.Get,
)

var persistences = wire.NewSet(
	infra.ProvidePostgresConn,
)

var domainStaff = wire.NewSet(
	serviceStaff.ProvideStaffServiceImpl,
	wire.Bind(new(serviceStaff.StaffService), new(*serviceStaff.StaffServiceImpl)),
	repositoryStaff.ProvideStaffRepositoryInfra,
	wire.Bind(new(repositoryStaff.StaffRepository), new(*repositoryStaff.StaffRepositoryInfra)),
)

var domainProduct = wire.NewSet(
	serviceProduct.ProvideProductServiceImpl,
	wire.Bind(new(serviceProduct.ProductService), new(*serviceProduct.ProductServiceImpl)),
	repositoryProduct.ProvideProductRepositoryInfra,
	wire.Bind(new(repositoryProduct.ProductRepository), new(*repositoryProduct.ProductRepositoryInfra)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainStaff,
	domainProduct,
)

var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	health.ProvideHealthHandler,
	staff.ProvideStaffHandler,
	product.ProvideProductHandler,
	router.ProvideRouter,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideJWTMiddleware,
)

func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		authMiddleware,
		//domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}

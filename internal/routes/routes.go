package routes //! This file has all the routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
	ServerConfig "github.com/chowdhuryrahulc/dynamodb/config" // this is our config folder in outer project
	HealthHandler "github.com/chowdhuryrahulc/dynamodb/internal/handlers/health"
	ProductHandler "github.com/chowdhuryrahulc/dynamodb/internal/handlers/product"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters() 		// these 3 functions are defined below
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) setConfigsRouters() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

// delete all r Routers below
func (r *Router) RouterHealth(repository adapter.Interface) {
	//? All Health Routes here 
	handler := HealthHandler.NewHandler(repository)

	// putting our routes, same done by gorilla mux
	r.router.Route("/health", func(route chi.Router){
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	//? All Product Routes here 
	handler := ProductHandler.NewHandler(repository)

	// putting our routes, same done by gorilla mux
	r.router.Route("/product", func(route chi.Router){
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableLogger() *Router {
	// chi router middleware gives us a logger, which we will use here
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	// chi router middleware gives us a Timeout, which we will use here
	//todo: r.config.GetTimeout()??
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	// chi router middleware gives us a logger, which we will use here
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
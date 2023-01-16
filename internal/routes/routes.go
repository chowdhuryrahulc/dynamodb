package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ServerConfig "github.com/chowdhuryrhulc/dynamodb/config" // this is our config folder in outer project
	HealthHandler "github.com/chowdhuryrhulc/dynamodb/internal/handlers/health"
	ProductHandler "github.com/chowdhuryrhulc/dynamodb/internal/handlers/product"
)

type Router struct {
	config *Config
	router *chi.Router
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters() 		// these 3 functions are defined below
	r.RouterHealth(repository)
	r.RouterProduct()

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
	handler := HealthHandler.newHandler(repository)

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
	handler := ProductHandler.newHandler(repository)

	// putting our routes, same done by gorilla mux
	r.router.Route("/product", func(route chi.Router){
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableTimeout() {

}

func (r *Router) EnableCORS() {

}

func (r *Router) EnableRecover() {

}

func (r *Router) EnableRequestID() {

}

func (r *Router) EnableRealIP() {

}

func (r *Router) EnableLogger() {
	// chi router gives us a logger, which we will use here
	r.router.Use(middleware.Logger)
	return r
}
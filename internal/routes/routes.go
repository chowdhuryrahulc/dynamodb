package routes

import "github.com/go-chi/chi/v5"

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

func (r *Router) SetRouters()*chi.Mux{

}

func (r *Router) setConfigsRouters(){
	
}

// delete all r Routers below
func (r *Router) RouterHealth(){
	
}

func (r *Router) RouterProduct(){
	
}

func (r *Router) EnableTimeout(){
	
}

func (r *Router) EnableCORS(){
	
}

func (r *Router) EnableRecover(){
	
}

func (r *Router) EnableRequestID(){
	
}

func (r *Router) EnableRealIP(){
	
}

func (r *Router) setConfigsRouters(){
	
}

func (r *Router) setConfigsRouters(){
	
}

func (r *Router) setConfigsRouters(){
	
}

func (r *Router) setConfigsRouters(){
	
}
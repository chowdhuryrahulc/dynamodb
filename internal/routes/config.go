package routes //! in this file, we have set up our CORS

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	// This function returns a empty config
	return &Config{}
}

func (c *Config) Cors(next http.Handler)http.Handler{
	// most imp function in thiss config.go
	return cors.New(cors.Options{
		//. (*) means all allowed. Below statements means all origins, methords, 
		// headers anr allowed, and all headers are exposed
		// anyone trying to hit with any ip address will be allowed
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 5,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration{
	return c.timeout
}

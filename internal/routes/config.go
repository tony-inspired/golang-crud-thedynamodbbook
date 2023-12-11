package routes

import (
	"net/http"
	"time"
	"github.com/go-chi/cors"
)

type Config struct{
	timeout time.Duration
}

func NewConfig() *Config{
	return &Config{}
}

//нам надо задать корс, а у него есть .Handler(next)
//поэтому наш метод принимает next http.Handler

/*
 вызов cors.New(...).Handler(next) создает новый HTTP-обработчик, 
 который интегрирует настройки CORS, заданные в cors.Options, 
 в обработку HTTP-запроса
*/
func (c *Config) Cors(next http.Handler) http.Handler{
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 5,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config{
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c;
}

func (c *Config) GetTimeout() time.Duration{
	return c.timeout
}
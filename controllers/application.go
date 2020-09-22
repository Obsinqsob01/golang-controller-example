package controllers

import "net/http"

type ApplicationController struct {
	Controller
}

func (a *ApplicationController) Get() {
	a.Writer.WriteHeader(http.StatusOK)
	a.Writer.Write([]byte("Hola mundo"))
}

func (c *ApplicationController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r

	switch r.Method {
	case http.MethodGet:
		c.Get()
	case http.MethodPost:
		c.Post()
	case http.MethodPut:
		c.Put()
	case http.MethodDelete:
		c.Delete()
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

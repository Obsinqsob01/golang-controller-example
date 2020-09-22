package controllers

import (
	"errors"
	"net/http"
)

var (
	controllers = map[string]ControllerInterface{}
	mux         *http.ServeMux
)

func init() {
	mux = http.NewServeMux()
}

func NewController(name string, controller ControllerInterface) error {
	for key, _ := range controllers {
		if key == name {
			return errors.New("Controller already exists")
		}
	}

	controllers[name] = controller
	mux.Handle(name, controller)

	return nil
}

func GetMux() *http.ServeMux {
	return mux
}

type Controller struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

type ControllerInterface interface {
	Get()
	Post()
	Put()
	Delete()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (c *Controller) Get() {
	http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Post() {
	http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Put() {
	http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
}

func (c *Controller) Delete() {
	http.Error(c.Writer, "Method not allowed", http.StatusMethodNotAllowed)
}

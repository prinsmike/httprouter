package httprouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Title       string           `json:"title"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	Description string           `json:"description"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Routes []Route

// Add a new Route to Routes.
func (r Routes) Add(title, method, pattern, description string, handlerFunc http.HandlerFunc) Routes {
	r = append(r, Route{title, method, pattern, description, handlerFunc})
	return r
}

// Create a new router.
func New(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Title).
			Handler(route.HandlerFunc)
	}

	return router
}

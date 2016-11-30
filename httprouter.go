package httprouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Label       string           `json:"label"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Routes []Route

// Add a new Route to Routes.
func (r *Routes) Add(label, method, pattern string, handlerFunc http.HandlerFunc) {
	r = append(r, Route{label, method, pattern, handlerFunc})
}

// Create a new router.
func New(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Label).
			Handler(route.HandlerFunc)
	}

	return router
}

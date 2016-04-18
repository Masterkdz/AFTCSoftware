package router

import (
	"../handlers"
	"net/http"
)

type Route struct {
	Name	string
	Method	string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route {
		"ServeTemplate",
		"GET",
		"/",
		handlers.ServeTemplate,
	},
		Route {
		"ServeStyle",
		"GET",
		"/style/style.css",
		handlers.ServeStyle,
	},
}
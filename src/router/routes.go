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
		"Test",
		"GET",
		"/",
		handlers.Test,
	},
	Route {
		"TestDB",
		"GET",
		"/db",
		handlers.Test,
	},
}
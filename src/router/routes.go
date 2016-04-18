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
		Route {
		"ServeJs",
		"GET",
		"/js/AJAXform.js",
		handlers.ServeJs,
	},
		Route {
		"ServeJs",
		"GET",
		"/js/jquery1.12.3.js",
		handlers.ServeJQuery,
	},
		Route {
		"GetUser",
		"POST",
		"/user",
		handlers.GetUser,
	},
		Route {
		"AddUser",
		"POST",
		"/",
		handlers.AddUser,
	},
		Route {
		"DeleteUser",
		"POST",
		"/",
		handlers.DeleteUser,
	},
}
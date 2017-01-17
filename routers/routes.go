package routers

import "github.com/apisearch/apisearch/handlers/api/v1"

var routes = Routes{
	Route{
		"Healthz",
		"GET",
		"/api/v1/status/healthz",
		v1.Healthz,
		[]string{},
	},
	Route{
		"GetSettingsByToken",
		"GET",
		"/api/v1/user/{token}",
		v1.GetSettingsByToken,
		[]string{},
	},
	Route{
		"CreateSettings",
		"POST",
		"/api/v1/user",
		v1.CreateSettings,
		[]string{},
	},
	Route{
		"UpdateSettings",
		"POST",
		"/api/v1/user/{token}",
		v1.UpdateSettings,
		[]string{},
	},
	Route{
		"DeleteSettings",
		"DELETE",
		"/api/v1/user/{token}",
		v1.DeleteSettings,
		[]string{},
	},
	Route{
		"SignIn",
		"POST",
		"/api/v1/sign/in",
		v1.SignIn,
		[]string{},
	},
	Route{
		"Search",
		"GET",
		"/api/v1/search/{userId}/{query}",
		v1.Search,
		[]string{},
	},
}

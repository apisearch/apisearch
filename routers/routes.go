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
		"GetSettingsByUserId",
		"GET",
		"/api/v1/settings/{userId}",
		v1.GetSettingsById,
		[]string{},
	},
	Route{
		"UpsertSettings",
		"POST",
		"/api/v1/settings/{userId}",
		v1.CreateSettings,
		[]string{},
	},
	Route{
		"DeleteSettings",
		"DELETE",
		"/api/v1/settings/{userId}",
		v1.DeleteSettings,
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

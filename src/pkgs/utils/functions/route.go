package utils_functions

var route *RouteType = &RouteType{
	Base: "/",
	Mail: MailType{
		Basic: "/api/mail/basic",
	},
}

func GetRouteUrls() *RouteType {
	return route
}

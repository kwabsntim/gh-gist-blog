package handlers

import (
	"AuthGo/middleware"
	"AuthGo/services"
	"net/http"
)

// amazonq-ignore-next-line

type ServiceContainer struct {
	RegisterService services.RegisterInterface
	LoginService    services.LoginInterface
	FetchService    services.FetchUsersInterface
}

// chain function to apply mutiple middleware
// Add this function before RouteSetup
func Chain(middlewares ...func(http.HandlerFunc) http.HandlerFunc) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

// recieves the service container from main.go
func RouteSetup(services *ServiceContainer) *http.ServeMux {
	//using server mux to map the requests
	handlers := NewHandlers(services)

	//creating chains for middleware to be used in the routes
	var (
		roleChain       = middleware.RoleMiddleware
		authChain       = middleware.AuthMiddleware
		methodChainGet  = middleware.MethodChecker("GET")
		methodChainPost = middleware.MethodChecker("POST")
	)
	//mapping the routes to the handlers with the middleware chains
	mux := http.NewServeMux()
	mux.HandleFunc("/api/Signup", Chain(methodChainPost, authChain, roleChain)(handlers.SignUp))
	mux.HandleFunc("/api/Login", Chain(methodChainPost, authChain)(handlers.Login))
	mux.HandleFunc("/api/FetchUsers", Chain(methodChainGet)(handlers.FetchAllUsers))
	//returning the mux
	return mux
}

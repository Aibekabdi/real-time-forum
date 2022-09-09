package http

import "net/http"

type Route struct {
	Path    string
	Handler http.HandlerFunc
	IsAuth  bool
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := h.createRoutes()
	for _, route := range routes {
		if route.IsAuth {
			// todo checks cookie
		}
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

func (h *Handler) createRoutes() []Route {
	return []Route{
		{
			Path:    "/",
			Handler: nil,
			IsAuth:  false,
		},
		{
			Path:    "/sign-up",
			Handler: nil,
			IsAuth:  false,
		},
		{
			Path:    "/sign-in",
			Handler: nil,
			IsAuth:  false,
		},
		{
			Path:    "/sign-out",
			Handler: nil,
			IsAuth:  true,
		},
		{
			Path:    "/profile",
			Handler: nil,
			IsAuth:  true,
		},
		{
			Path:    "/posts/create",
			Handler: nil,
			IsAuth:  true,
		},
		{
			Path:    "/posts",
			Handler: nil,
			IsAuth:  false,
		},
		{
			Path:    "/posts/id",
			Handler: nil,
			IsAuth:  false,
		},
		{
			Path:    "/vote",
			Handler: nil,
			IsAuth:  true,
		},
		{
			Path:    "/comment",
			Handler: nil,
			IsAuth:  true,
		},
		{
			Path:    "/chat",
			Handler: nil,
			IsAuth:  true,
		},
	}
}

package http

import (
	"forum/internal/service"
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	IsAuth  bool
}

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	routes := h.createRoutes()
	for _, route := range routes {
		if route.IsAuth {
			route.Handler = h.isSessionValid(route.Handler)
		}
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

func (h *Handler) createRoutes() []Route {
	return []Route{
		{
			Path:    "/sign-up",
			Handler: h.signup,
			IsAuth:  false,
		},
		{
			Path:    "/sign-in",
			Handler: h.signin,
			IsAuth:  false,
		},
		{
			Path:    "/sign-out",
			Handler: h.signout,
			IsAuth:  true,
		},
		{
			Path:    "/profile",
			Handler: h.profile,
			IsAuth:  true,
		},
		{
			Path:    "/posts/create",
			Handler: h.createPost,
			IsAuth:  true,
		},
		{
			Path:    "/posts",
			Handler: h.posts,
			IsAuth:  false,
		},
		{
			Path:    "/posts/:id",
			Handler: h.post,
			IsAuth:  false,
		},
		{
			Path:    "/vote/post",
			Handler: h.votePost,
			IsAuth:  true,
		},
		{
			Path:    "/vote/comment",
			Handler: h.voteComment,
			IsAuth:  true,
		},
		{
			Path:    "/comment",
			Handler: h.createComment,
			IsAuth:  true,
		},
		{
			Path:    "/chat",
			Handler: h.chat,
			IsAuth:  true,
		},
	}
}

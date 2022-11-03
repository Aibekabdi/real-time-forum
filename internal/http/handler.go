package http

import (
	"forum/internal/service"
	"log"
	"net/http"
	"text/template"
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
	fs := http.FileServer(http.Dir("./web/src"))
	mux.Handle("/src/", http.StripPrefix("/src/", fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl, err := template.ParseFiles("./web/index.html")
		if err != nil {
			jsonResponse(w, r, http.StatusInternalServerError, err.Error())
		} else if err := indexTmpl.Execute(w, nil); err != nil {
			jsonResponse(w, r, http.StatusInternalServerError, err.Error())
		}
	})
	log.Println("http://localhost:8080/")
	return mux
}

func (h *Handler) createRoutes() []Route {
	return []Route{
		{
			Path:    "/api/sign-up",
			Handler: h.signup,
			IsAuth:  false,
		},
		{
			Path:    "/api/sign-in",
			Handler: h.signin,
			IsAuth:  false,
		},
		{
			Path:    "/api/sign-out",
			Handler: h.signout,
			IsAuth:  true,
		},
		{
			Path:    "/api/profile",
			Handler: h.profile,
			IsAuth:  true,
		},
		{
			Path:    "/api/posts/create",
			Handler: h.createPost,
			IsAuth:  true,
		},
		{
			Path:    "/api/posts",
			Handler: h.posts,
			IsAuth:  false,
		},
		{
			Path:    "/api/posts/:id",
			Handler: h.post,
			IsAuth:  false,
		},
		{
			Path:    "/api/vote/post",
			Handler: h.votePost,
			IsAuth:  true,
		},
		{
			Path:    "/api/vote/comment",
			Handler: h.voteComment,
			IsAuth:  true,
		},
		{
			Path:    "/api/comment",
			Handler: h.createComment,
			IsAuth:  true,
		},
		{
			Path:    "/api/chat",
			Handler: h.chat,
			IsAuth:  true,
		},
	}
}

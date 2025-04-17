package port

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Route struct {
	Method  string
	Handler http.HandlerFunc
	Pattern string
}

type Router struct {
	routes []Route
}

type ContextKey string

const ParamsKey ContextKey = "params"

func NewRouter() *Router {
	return &Router{routes: []Route{}}
}

func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

func matchPath(pattern, path string) (map[string]string, bool) {
	pParts := strings.Split(strings.Trim(pattern, "/"), "/")
	uParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(pParts) != len(uParts) {
		return nil, false
	}

	params := map[string]string{}

	for i := range pParts {
		if strings.HasPrefix(pParts[i], "{") && strings.HasSuffix(pParts[i], "}") {
			key := pParts[i][1 : len(pParts[i])-1]
			params[key] = uParts[i]
		} else if pParts[i] != uParts[i] {
			return nil, false
		}
	}

	return params, true
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if req.Method != route.Method {
			continue
		}

		params, ok := matchPath(route.Pattern, req.URL.Path)
		if ok {
			for key, value := range params {
				if value == "" {
					http.Error(w, fmt.Sprintf("Invalid '%s' param", key), http.StatusBadRequest)
					return
				}
			}

			ctx := context.WithValue(req.Context(), ParamsKey, params)
			route.Handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
	}
	http.NotFound(w, req)
}

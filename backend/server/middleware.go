package server

import (
	"context"
	"net/http"

	"bitbucket.org/andreychernih/tweemote/ctx"
	"bitbucket.org/andreychernih/tweemote/lib"
)

// middlewares

func corsMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4000")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		handler(w, r)
	}
}

func authMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := lib.VerifyAuth(lib.GetOsinServer(), r, w)
		if err != nil {
			http.Error(w, errorJson(err), 500)
			return
		}

		if user != nil {
			c := r.Context()
			c = context.WithValue(c, ctx.CtxUserKey, user)

			next(w, r.WithContext(c))
		}
	}
}

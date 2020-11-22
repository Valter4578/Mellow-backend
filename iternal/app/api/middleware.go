package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"oblique/iternal/app/api/auth"
	"oblique/iternal/app/logger"
)

var Api *API

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	// create new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, r.Method, time.Since(start))
			}()

			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// check if real method is equal to expected
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// call the next handler or middleware in chain
			f(w, r)
		}
	}
}

// Token middleware checks jwt token for endpoints
func Token() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !Api.IsDebugMode {
				bearerToken := r.Header.Get("Authorization")
				authToken := strings.Split(bearerToken, " ")[0]

				_, err := auth.VerifyToken(authToken)
				if err != nil {
					log.Println(err)
					err = errors.New("Can't verify jwt token")
					msg := logger.JSONError(err)
					io.WriteString(w, msg)
					return
				}

				f(w, r)
			} else {
				f(w, r)
			}
		}

	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

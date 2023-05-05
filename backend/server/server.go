package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/andreychernih/tweemote/twitter"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const LISTEN = ":4001"

type Router struct {
	router *mux.Router
}

// TODO fetch from config or env variable
func GetBaseURL() string {
	return "http://localhost:4001"
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *Router) handle(path string, handler func(http.ResponseWriter, *http.Request), methods string) {
	preflightHandler := func(http.ResponseWriter, *http.Request) {
		// reply to OPTIONS request with HTTP 200
	}
	r.router.HandleFunc(path, corsMiddleware(preflightHandler)).Methods("OPTIONS")
	r.router.HandleFunc(path, corsMiddleware(handler)).Methods(methods)
}

func (r *Router) httpHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.ServeHTTP(w, req)
	})
}

func createRouter() *Router {
	r := NewRouter()

	// oAuth
	r.handle("/api/v1/auth", CreateAccessToken, "POST")
	r.handle("/api/v1/users", CreateUserHandler, "POST")
	r.handle("/api/v1/users/me", authMiddleware(GetCurrentUserHandler), "GET")

	// Twitter Accounts
	r.handle("/api/v1/twitter_accounts", authMiddleware(GetTwitterAccountsHandler), "GET")
	r.handle("/api/v1/twitter_accounts/link", authMiddleware(CreateLinkTwitterAccountHandler(twitter.DefaultTwitterClientProvider{})), "GET")
	r.handle("/api/v1/twitter_accounts/callback", CreateLinkTwitterAccountCallbackHandler(twitter.DefaultTwitterClientProvider{}), "GET")
	r.handle("/api/v1/twitter_accounts/{id}", authMiddleware(UnlinkTwitterAccountHandler), "DELETE")

	// Campaigns
	r.handle("/api/v1/twitter_accounts/{id}/campaigns", authMiddleware(GetCampaignsHandler), "GET")

	// Keywords
	r.handle("/api/v1/campaigns/{campaignId}/keywords", authMiddleware(GetKeywordsHandler), "GET")
	r.handle("/api/v1/campaigns/{campaignId}/keywords", authMiddleware(CreateKeywordHandler), "POST")
	r.handle("/api/v1/campaigns/{campaignId}/keywords/{id}", authMiddleware(DeleteKeywordHandler), "DELETE")

	// Stats
	r.handle("/api/v1/campaigns/{campaignId}/stats", authMiddleware(GetStatsHandler), "GET")

	return r
}

func Start() {
	r := createRouter()

	handler := handlers.RecoveryHandler()(r.httpHandler())
	handler = handlers.CombinedLoggingHandler(os.Stdout, handler)
	s := &http.Server{
		Addr:         LISTEN,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Listening on %s", LISTEN)
	log.Fatal(s.ListenAndServe())
}

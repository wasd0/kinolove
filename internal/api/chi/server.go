package chi

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"
	"kinolove/internal/app/apiProvider"
	"kinolove/pkg/config"
	"kinolove/pkg/logger"
	"kinolove/pkg/utils/app"
	"kinolove/pkg/utils/jwtUtils"
	"net/http"
)

type Server struct {
	server *http.Server
}

func SetupServer(cfg *config.Config, provider *apiProvider.ApiProvider, formatter *logger.LogFormatterImpl, jwt *jwtUtils.Auth) *Server {
	mux := chi.NewRouter()
	auth := jwt
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{Addr: addr, Handler: mux}

	setUpMiddlewares(cfg, mux, formatter, auth)
	setUpRouters(mux, provider)

	return &Server{
		server: server,
	}
}

func (s *Server) MustRun() app.Callback {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log().Fatal(err, "http server start failed")
		}
	}()

	return s.server.Shutdown
}

func setUpMiddlewares(cfg *config.Config, mux *chi.Mux, formatter *logger.LogFormatterImpl, auth *jwtUtils.Auth) {
	mux.Use(middleware.RequestLogger(formatter))
	mux.Use(middleware.Timeout(cfg.Server.IdleTimeout))
	mux.Use(jwtauth.Verifier(auth.GetJwt()))
	//mux.Use(cors.AllowAll().Handler)
	mux.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"https://*", "http://localhost:*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler)
}

func setUpRouters(mux *chi.Mux, provider *apiProvider.ApiProvider) {
	mux.Route(provider.UserApi().Register())
	mux.Route(provider.MovieApi().Register())
	mux.Route(provider.LoginApi().Register())
	mux.NotFound(provider.DefaultApi().NotFound)
	mux.MethodNotAllowed(provider.DefaultApi().MethodNotAllowed)
}

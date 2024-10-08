package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"kinolove/api/apiModel"
	"kinolove/api/apiModel/movie"
	"kinolove/internal/middleware"
	"kinolove/internal/service"
	"kinolove/internal/service/dto"
	"kinolove/pkg/constants/perms"
	"kinolove/pkg/logger"
	"net/http"
	"strconv"
)

type MovieApi struct {
	movieService service.MovieService
	auth         *middleware.AuthMiddleware
}

func NewMovieApi(movieService service.MovieService, auth *middleware.AuthMiddleware) *MovieApi {
	return &MovieApi{movieService: movieService, auth: auth}
}

func (u *MovieApi) Register() (string, func(router chi.Router)) {
	return "/api/v1/movies", u.Handle
}

func (u *MovieApi) Handle(router chi.Router) {
	router.Get("/", u.findAll)
	router.With(u.auth.HasPermission(perms.Movie, perms.Create)).Post("/", u.create)
	router.Get("/{id}", u.findById)
	router.With(u.auth.HasPermission(perms.Movie, perms.Edit)).Put("/{id}", u.update)
}

func (u *MovieApi) findAll(w http.ResponseWriter, r *http.Request) {
	movies, err := u.movieService.FindAll()
	response := movie.ResMovieFindAll{Data: movies}

	if err != nil {
		RenderError(w, r, err)
		return
	}

	if renderErr := render.Render(w, r, &response); renderErr != nil {
		logger.Log().Fatal(renderErr, "error rendering error")
	}

}

func (u *MovieApi) create(w http.ResponseWriter, r *http.Request) {
	request := movie.ReqMovieCreate{}

	if err := render.Bind(r, &request); err != nil {
		RenderError(w, r, service.BadRequest(err, "Failed get request body"))
		return
	}

	if id, err := u.movieService.CreateMovie(request.MovieCreateRequest); err != nil {
		RenderError(w, r, err)
	} else {
		response := apiModel.RestResponse[int64]{Data: &id}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			RenderError(w, r, service.InternalError(renderErr))
			return
		}
	}
}

func (u *MovieApi) findById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)

	if parseErr != nil {
		RenderError(w, r, service.InternalError(parseErr))
		return
	}

	if m, err := u.movieService.FindById(id); err != nil {
		RenderError(w, r, err)
	} else {
		response := apiModel.RestResponse[dto.MovieSingleResponse]{Data: &m}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			RenderError(w, r, err)
		}
	}
}

func (u *MovieApi) update(w http.ResponseWriter, r *http.Request) {
	req := movie.ReqMovieUpdate{}

	if err := render.Bind(r, &req); err != nil {
		RenderError(w, r, service.BadRequest(err, "Failed get request body"))
		return
	}

	idStr := chi.URLParam(r, "id")

	id, parseErr := strconv.ParseInt(idStr, 10, 64)

	if parseErr != nil {
		RenderError(w, r, service.InternalError(parseErr))
		return
	}

	err := u.movieService.Update(id, req.MovieUpdateRequest)

	if err != nil {
		RenderError(w, r, err)
	}
}

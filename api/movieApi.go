package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"kinolove/api/apiModel"
	"kinolove/api/apiModel/movie"
	"kinolove/internal/service"
	"kinolove/internal/service/dto"
	"kinolove/pkg/logger"
	"net/http"
	"strconv"
)

type MovieApi struct {
	movieService service.MovieService
	log          logger.Common
}

func NewMovieApi(log logger.Common, movieService service.MovieService) *MovieApi {
	return &MovieApi{log: log, movieService: movieService}
}

func (u *MovieApi) Register() (string, func(router chi.Router)) {
	return "/api/v1/movies", u.Handle
}

func (u *MovieApi) Handle(router chi.Router) {
	router.Get("/", u.findAll)
	router.Post("/", u.create)
	router.Get("/{id}", u.findById)
	router.Put("/{id}", u.update)
}

func (u *MovieApi) findAll(w http.ResponseWriter, r *http.Request) {
	movies, err := u.movieService.FindAll()
	response := movie.ResMovieFindAll{Data: movies}

	if err != nil {
		renderError(w, r, err, u.log)
		return
	}

	if renderErr := render.Render(w, r, &response); renderErr != nil {
		u.log.Fatal(renderErr, "error rendering error")
	}

}

func (u *MovieApi) create(w http.ResponseWriter, r *http.Request) {
	request := movie.ReqMovieCreate{}

	if err := render.Bind(r, &request); err != nil {
		renderError(w, r, service.BadRequest(err, "Failed get request body"), u.log)
		return
	}

	if id, err := u.movieService.CreateMovie(request.MovieCreateRequest); err != nil {
		renderError(w, r, err, u.log)
	} else {
		response := apiModel.RestResponse[int64]{Data: &id}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			renderError(w, r, service.InternalError(renderErr), u.log)
			return
		}
	}
}

func (u *MovieApi) findById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)

	if parseErr != nil {
		renderError(w, r, service.InternalError(parseErr), u.log)
		return
	}

	if m, err := u.movieService.FindById(id); err != nil {
		renderError(w, r, err, u.log)
	} else {
		response := apiModel.RestResponse[dto.MovieSingleResponse]{Data: &m}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			renderError(w, r, err, u.log)
		}
	}
}

func (u *MovieApi) update(w http.ResponseWriter, r *http.Request) {
	req := movie.ReqMovieUpdate{}

	if err := render.Bind(r, &req); err != nil {
		renderError(w, r, service.BadRequest(err, "Failed get request body"), u.log)
		return
	}

	idStr := chi.URLParam(r, "id")

	id, parseErr := strconv.ParseInt(idStr, 10, 64)

	if parseErr != nil {
		renderError(w, r, service.InternalError(parseErr), u.log)
		return
	}

	err := u.movieService.Update(id, req.MovieUpdateRequest)

	if err != nil {
		renderError(w, r, err, u.log)
	}
}

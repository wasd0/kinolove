package api

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kinolove/api/apiModel"
	"kinolove/api/apiModel/user"
	"kinolove/internal/service"
	"kinolove/internal/service/dto"
	"kinolove/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserApi struct {
	userService service.UserService
	log         logger.Common
}

func NewUserApi(log logger.Common, userService service.UserService) *UserApi {
	return &UserApi{log: log, userService: userService}
}

func (u *UserApi) Register() (string, func(router chi.Router)) {
	return "/api/v1/users", u.Handle
}

func (u *UserApi) Handle(router chi.Router) {
	router.Post("/", u.createUser)
	router.Get("/{username}", u.findByUsername)
	router.Put("/{id}", u.update)
}

func (u *UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	request := user.ReqUserCreate{}

	if err := render.Bind(r, &request); err != nil {
		renderError(w, r, service.BadRequest(err, "Failed get request body"), u.log)
		return
	}

	if id, err := u.userService.CreateUser(request.UserCreateRequest); err != nil {
		renderError(w, r, err, u.log)
	} else {
		response := apiModel.RestResponse[uuid.UUID]{Data: &id}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			renderError(w, r, service.InternalError(renderErr), u.log)
			return
		}
	}
}

func (u *UserApi) findByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	byUsername, err := u.userService.FindByUsername(username)
	if err != nil {
		renderError(w, r, err, u.log)
		return
	}

	response := apiModel.RestResponse[dto.UserSingleResponse]{Data: &byUsername}
	if renderErr := render.Render(w, r, &response); renderErr != nil {
		renderError(w, r, err, u.log)
		return
	}
}

func (u *UserApi) update(w http.ResponseWriter, r *http.Request) {
	uuidStr := chi.URLParam(r, "id")

	if err := uuid.Validate(uuidStr); err != nil {
		renderError(w, r, service.BadRequest(errors.New("Wrong id"), "wrong user id"), u.log)
		return
	}

	id, err := uuid.Parse(uuidStr)

	if err != nil {
		renderError(w, r, service.InternalError(err), u.log)
		return
	}

	request := user.ReqUserUpdate{}
	if err = render.Bind(r, &request); err != nil {
		renderError(w, r, service.BadRequest(err, "Failed get request body"), u.log)
		return
	}

	servErr := u.userService.Update(id, request.UserUpdateRequest)

	if servErr != nil {
		renderError(w, r, servErr, u.log)
		return
	}
}
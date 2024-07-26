package serviceProvider

import (
	repoProvider2 "kinolove/internal/app/repoProvider"
	"kinolove/internal/service"
)

type ServiceProvider struct {
	repoProvider *repoProvider2.RepoProvider

	user  service.UserService
	movie service.MovieService
	login service.LoginService
}

func InitServices(provider *repoProvider2.RepoProvider) *ServiceProvider {
	return &ServiceProvider{repoProvider: provider}
}

func (sp *ServiceProvider) UserService() service.UserService {
	if sp.user != nil {
		return sp.user
	}

	user := service.NewUserService(sp.repoProvider.UserRepo())
	sp.user = user
	return sp.user
}

func (sp *ServiceProvider) MovieService() service.MovieService {
	if sp.movie != nil {
		return sp.movie
	}

	movie := service.NewMovieService(sp.repoProvider.MovieRepo())
	sp.movie = movie
	return sp.movie
}

func (sp *ServiceProvider) LoginService() service.LoginService {
	if sp.login != nil {
		return sp.login
	}

	login := service.NewLoginService(sp.UserService())
	sp.login = login
	return sp.login
}

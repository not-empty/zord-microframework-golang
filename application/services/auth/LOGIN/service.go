package auth

import (
	"go-skeleton/application/domain/auth"
	"go-skeleton/application/services"
	"go-skeleton/pkg/config"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	services.BaseService
	response *Response
	config   *config.Config
}

func NewService(log services.Logger, config *config.Config) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
		config: config,
	}
}

func (s *Service) Execute(request Request) {
	s.Logger.Debug("Generating Token")

	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	s.produceResponseRule(request.Token)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(Access auth.Token) {
	tokenIndex := -1
	for i, token := range s.config.AccessToken {
		if token == Access.Token {
			tokenIndex = i
		}
	}

	if tokenIndex < 0 {
		s.Error = &services.Error{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   "Invalid Access",
		}
		return
	}

	if s.config.AccessSecret[tokenIndex] != Access.Secret {
		s.Error = &services.Error{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   "Invalid Access",
		}
		return
	}

	if s.config.AccessContext[tokenIndex] != Access.Context {
		s.Error = &services.Error{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   "Invalid Access",
		}
		return
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(s.config.JwtExpiration))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.config.Secret))

	if err != nil {
		s.Error = &services.Error{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		}
		return
	}

	if s.Error == nil {
		s.response = &Response{
			Status:  http.StatusOK,
			Message: t,
		}
	}
}

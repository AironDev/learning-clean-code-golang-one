package http

import (
	usecases "github.com/airondev/learning-clean-code-golang-one/usecase"
	"github.com/labstack/echo"
)

type Handler struct {
	usecases.ArticleUsecase
}

func NewHttpHandler(e *echo.Echo, us usecases.ArticleUsecase) {
	_ = &Handler{
		us,
	}
}

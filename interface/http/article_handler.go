package http

import (
	"github.com/airondev/learning-clean-code-golang-one/model"
	"github.com/airondev/learning-clean-code-golang-one/usecase"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	articleUse usecase.ArticleUsecase
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.articleUse.Fetch(cursor, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *ArticleHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.articleUse.GetByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *model.Article) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArticleHandler) Store(c echo.Context) error {
	var article model.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ar, err := a.articleUse.Store(&article)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}
func (a *ArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	_, err = a.articleUse.Delete(id)

	if err != nil {

		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func NewArticleHttpHandler(e *echo.Echo, us usecase.ArticleUsecase) {
	handler := &ArticleHandler{
		articleUse: us,
	}

	e.GET("/article", handler.FetchArticle)
	e.POST("/article", handler.Store)
	e.GET("/article/:id", handler.GetByID)
	e.DELETE("/article/:id", handler.Delete)

}

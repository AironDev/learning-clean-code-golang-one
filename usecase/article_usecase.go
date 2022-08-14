package usecase

import (
	"github.com/airondev/learning-clean-code-golang-one/model"
	"github.com/airondev/learning-clean-code-golang-one/repository"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// ArticleUsecase Usecase is like controller which implements a repository - concrete implementation of a repository
type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*model.Article, string, error)
	GetByID(id int64) (*model.Article, error)
	Update(ar *model.Article) (*model.Article, error)
	GetByTitle(title string) (*model.Article, error)
	Store(*model.Article) (*model.Article, error)
	Delete(id int64) (bool, error)
}

type articleUsecase struct {
	articleRepos repository.ArticleRepository
}

func NewArticleUsecase(a repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepos: a,
	}
}

func (a *articleUsecase) Fetch(cursor string, num int64) ([]*model.Article, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.articleRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (a *articleUsecase) GetByID(id int64) (*model.Article, error) {
	res, err := a.articleRepos.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *articleUsecase) Update(ar *model.Article) (*model.Article, error) {
	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ar)
}

func (a *articleUsecase) GetByTitle(title string) (*model.Article, error) {

	res, err := a.articleRepos.GetByTitle(title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *articleUsecase) Store(m *model.Article) (*model.Article, error) {

	existedArticle, _ := a.GetByTitle(m.Title)
	if existedArticle != nil {
		return nil, nil
	}

	id, err := a.articleRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(id int64) (bool, error) {
	existedArticle, _ := a.articleRepos.GetByID(id)
	logrus.Info("Masuk Sini")
	if existedArticle == nil {
		logrus.Info("Masuk Sini2")
		return false, model.NOT_FOUND_ERROR
	}
	logrus.Info("Masuk Sini3")

	return a.articleRepos.Delete(id)
}

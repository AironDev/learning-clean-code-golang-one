package repository

import (
	"database/sql"
	"fmt"
	"github.com/airondev/learning-clean-code-golang-one/model"
	"github.com/sirupsen/logrus"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) ArticleRepository {
	return &mysqlArticleRepository{Conn}
}

func (m *mysqlArticleRepository) fetch(query string, args ...interface{}) ([]*model.Article, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Article, 0)
	for rows.Next() {
		t := new(model.Article)
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlArticleRepository) Fetch(cursor string, num int64) ([]*model.Article, error) {

	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}
func (m *mysqlArticleRepository) GetByID(id int64) (*model.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &model.Article{}
	if len(list) > 0 {
		a = list[0]
	}

	return a, nil
}

func (m *mysqlArticleRepository) GetByTitle(title string) (*model.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE title = ?`

	list, err := m.fetch(query, title)
	if err != nil {
		return nil, err
	}

	a := &model.Article{}
	if len(list) > 0 {
		a = list[0]
	}
	return a, nil
}

func (m *mysqlArticleRepository) Store(a *model.Article) (int64, error) {

	query := `INSERT  article SET title=? , content=? , author_id=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, err
	}

	logrus.Debug("Created At: ", a.CreatedAt)
	res, err := stmt.Exec(a.Title, a.Content, a.UpdatedAt, a.CreatedAt)
	if err != nil {

		return 0, err
	}
	return res.LastInsertId()
}

func (m *mysqlArticleRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM article WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, err
	}
	res, err := stmt.Exec(id)
	if err != nil {

		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		logrus.Error(err)
		return false, err
	}

	return true, nil
}
func (m *mysqlArticleRepository) Update(ar *model.Article) (*model.Article, error) {
	query := `UPDATE article set title=?, content=?, author_id=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, nil
	}

	res, err := stmt.Exec(ar.Title, ar.Content, ar.UpdatedAt, ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}

	return ar, nil
}

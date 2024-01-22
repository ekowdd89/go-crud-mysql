package handler

import (
	"net/http"

	"github.com/ekowdd89/go-crud-mysql/config"
	"github.com/ekowdd89/go-crud-mysql/utils"
	"github.com/gorilla/mux"
)

type IArticle interface {
	GetArticle(w http.ResponseWriter, r *http.Request)
	FindArticle(w http.ResponseWriter, r *http.Request)
	CreateArticle(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
}
type Article struct {
	db         config.Database
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	ThumnImage string `json:"thumnImage"`
	IsPublish  bool   `json:"isPublish"`
	CreatedBy  User   `json:"createdBy,omitempty"`
	UpdatedBy  User   `json:"updatedBy,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdateAt   string `json:"updated_at"`
}

func (a *Article) GetArticle(w http.ResponseWriter, r *http.Request) {
	rowsArticles, _ := a.db.Query("SELECT * FROM sys_mt_blog_articles")
	articles := []*Article{}
	for rowsArticles.Next() {
		article := new(Article)
		rowsArticles.Scan(
			&article.ID,
			&article.Title,
			&article.Slug,
			&article.Content,
			&article.ThumnImage,
			&article.IsPublish,
			&article.CreatedBy,
			&article.UpdatedBy,
			&article.CreatedAt,
			&article.UpdateAt,
		)
		articles = append(articles, article)
	}
	userStmt := `"SELECT * FROM users WHERE id=?"`
	for _, ar := range articles {
		row, _ := a.db.Query(userStmt, 1)
		user := new(User)
		for row.Next() {
			row.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.EmailVerifiedAt,
				&user.Password,
				&user.RememberToken,
				&user.CreatedAt,
				&user.UpdatedAt,
			)
			ar.CreatedBy = *user
		}
		updatedBy := new(User)

		Uprow, _ := a.db.Query(userStmt, 1)

		for Uprow.Next() {
			Uprow.Scan(
				&updatedBy.ID,
				&updatedBy.Name,
				&updatedBy.Email,
				&updatedBy.EmailVerifiedAt,
				&updatedBy.Password,
				&updatedBy.RememberToken,
				&updatedBy.CreatedAt,
				&updatedBy.UpdatedAt,
			)
			ar.UpdatedBy = *updatedBy
		}
	}
	utils.Responder(w, utils.Response{Data: articles, Code: http.StatusOK})
}
func (a *Article) FindArticle(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	rowsArticles, err := a.db.Query("SELECT * FROM sys_mt_blog_articles WHERE id=?", ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusInternalServerError, Err: err})
	}
	for rowsArticles.Next() {
		rowsArticles.Scan(
			&a.ID,
			&a.Title,
			&a.Slug,
			&a.Content,
			&a.ThumnImage,
			&a.IsPublish,
			&a.CreatedBy,
			&a.UpdatedBy,
			&a.CreatedAt,
			&a.UpdateAt,
		)
	}
	userStmt := `"SELECT * FROM users WHERE id=?"`
	row, _ := a.db.Query(userStmt, 1)
	user := new(User)
	for row.Next() {
		row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.EmailVerifiedAt,
			&user.Password,
			&user.RememberToken,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		a.CreatedBy = *user
	}
	updatedBy := new(User)

	Uprow, _ := a.db.Query(userStmt, 1)

	for Uprow.Next() {
		Uprow.Scan(
			&updatedBy.ID,
			&updatedBy.Name,
			&updatedBy.Email,
			&updatedBy.EmailVerifiedAt,
			&updatedBy.Password,
			&updatedBy.RememberToken,
			&updatedBy.CreatedAt,
			&updatedBy.UpdatedAt,
		)
		a.UpdatedBy = *updatedBy
	}
	utils.Responder(w, utils.Response{Data: a, Code: http.StatusOK})
}
func (a *Article) CreateArticle(w http.ResponseWriter, r *http.Request) {
	_, err := a.db.Exec("INSERT INTO sys_mt_blog_articles (title, slug, content, thumnImage, isPublish, createdBy, updatedBy) VALUES(?,?,?,?,?,?,?)", a.Title, a.Slug, a.Content, a.ThumnImage, a.IsPublish, a.CreatedBy, a.UpdatedBy)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusInternalServerError, Err: err})
	}
	utils.Responder(w, utils.Response{Message: "Created Succesfully", Code: http.StatusCreated})
}
func (a *Article) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	_, err := a.db.Exec("UPDATE sys_mt_blog_articles SET title=?, slug=?, content=?, thumnImage=?, isPublish=?, createdBy=?, updatedBy=? WHERE id=?", a.Title, a.Slug, a.Content, a.ThumnImage, a.IsPublish, a.CreatedBy, a.UpdatedBy, ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusInternalServerError, Err: err})
	}
	utils.Responder(w, utils.Response{Message: "Updated Succesfully", Code: http.StatusOK})
}
func (a *Article) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	_, err := a.db.Exec("DELETE FROM sys_mt_blog_articles WHERE id=?", ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: "Deleted Successfully", Code: http.StatusOK})
}

func NewArticle(db config.Database) IArticle {
	return &Article{db: db}
}

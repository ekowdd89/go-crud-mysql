package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ekowdd89/go-crud-mysql/config"
	"github.com/ekowdd89/go-crud-mysql/utils"
	"github.com/gorilla/mux"
)

type ICategory interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	FindCategory(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type Category struct {
	db        config.Database
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Category) GetCategory(w http.ResponseWriter, r *http.Request) {
	rows, err := c.db.Query("SELECT * FROM sys_mt_categories")
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusInternalServerError, Err: err})
	}
	categorys := []*Category{}
	for rows.Next() {
		rows.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		categorys = append(categorys, c)
	}
	utils.Responder(w, utils.Response{Message: "Successfully", Code: http.StatusOK, Data: categorys})
}
func (c *Category) FindCategory(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	row, err := c.db.Query("SELECT * FROM sys_mt_categories WHERE id=?", ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusInternalServerError, Err: err})
	}
	for row.Next() {
		row.Scan(
			&c.ID,
			&c.Name,
			&c.Slug,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
	}
	utils.Responder(w, utils.Response{Message: "Successfully", Code: http.StatusOK, Data: c})
}
func (c *Category) CreateCategory(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&c)
	insertStmt := `INSERT INTO sys_mt_categories (name, slug) VALUES(?,?)`
	_, err := c.db.Exec(insertStmt, c.Name, c.Slug)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: "Created Succesfully", Code: http.StatusCreated})
}
func (c *Category) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	json.NewDecoder(r.Body).Decode(&c)
	log.Println(c)
	updateStmt := `UPDATE sys_mt_categories SET name=?, slug=? WHERE id=?`
	_, err := c.db.Exec(updateStmt, c.Name, c.Slug, ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "faild", Code: http.StatusBadRequest, Err: err})
	}

	utils.Responder(w, utils.Response{Message: "Updated Successfully", Code: http.StatusOK})
}
func (c *Category) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	_, err := c.db.Exec("DELETE FROM sys_mt_categories WHERE id=?", ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Faild", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: "Deleted Successfully", Code: http.StatusOK})
}

func NewCategory(db config.Database) ICategory {
	return &Category{db: db}
}

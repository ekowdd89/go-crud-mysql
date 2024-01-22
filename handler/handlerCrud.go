package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ekowdd89/go-crud-mysql/config"
	"github.com/ekowdd89/go-crud-mysql/utils"
	"github.com/gorilla/mux"
)

type IUser interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	FindUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type User struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password        string `json:"password,omitempty"`
	RememberToken   string `json:"remember_token,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	db              config.Database
}
type NUser struct {
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	rows, err := u.db.Query("select * from users")
	if err != nil {
		utils.Responder(w, utils.Response{
			Message: "Fail",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
	}
	defer rows.Close()
	users := []*User{}
	for rows.Next() {
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.EmailVerifiedAt,
			&u.Password,
			&u.RememberToken,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		users = append(users, u)
	}
	utils.Responder(w, utils.Response{Message: "Successfully", Code: http.StatusOK, Data: users})
}

func (u *User) FindUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	rows, err := u.db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Fail", Code: http.StatusBadRequest, Err: err})
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.EmailVerifiedAt,
			&u.Password,
			&u.RememberToken,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
	}
	utils.Responder(w, utils.Response{Message: `Successfully`, Code: http.StatusOK, Data: u})
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&u)
	queryStatement := `INSERT INTO users (name, email,password) VALUES(?,?,?) `
	resp, err := u.db.Exec(queryStatement, u.Name, u.Email, u.Password)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Fail", Code: http.StatusBadRequest, Err: err})
	}
	_, errCreated := resp.RowsAffected()
	if errCreated != nil {
		utils.Responder(w, utils.Response{Message: "Fail", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: `Created Successfully`, Code: http.StatusCreated})
}
func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	json.NewDecoder(r.Body).Decode(&u)
	updateStmt := `UPDATE users SET name=?, email=?, password=? WHERE id=?`
	_, err := u.db.Exec(updateStmt, u.Name, u.Email, u.Password, ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Fail", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: `Updated Successfully`, Code: http.StatusOK})
}
func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	stmtDelete := `DELETE FROM users WHERE id=?`
	_, err := u.db.Exec(stmtDelete, ID)
	if err != nil {
		utils.Responder(w, utils.Response{Message: "Fail", Code: http.StatusBadRequest, Err: err})
	}
	utils.Responder(w, utils.Response{Message: `Deleted Successfully ` + ID, Code: http.StatusOK})
}
func NewUser(db config.Database) IUser {
	return &User{
		db: db,
	}
}

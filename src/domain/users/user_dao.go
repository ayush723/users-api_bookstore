package users

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ayush723/users-api_bookstore/src/datasources/mysql/users_db"
	"github.com/ayush723/utils-go_bookstore/logger"

	"github.com/ayush723/users-api_bookstore/src/utils/mysql_utils"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

const (
	// indexUniqueEmail = "users.email_UNIQUE"
	ErrorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "select id, first_name, last_name,email,date_created, status FROM users WHERE id =?;"
	queryUpdateuser       = "UPDATE users SET first_name=?,last_name=?, email=? WHERE id =?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

//only has access to database

//dao = data access object

func (user *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement",err)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {

		logger.Error("error when trying to get user  by id",getErr)
		return rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	return nil
}

func (user *User) Save() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement",err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user",saveErr)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user",err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	user.Id = userId

	return nil
}

func (user *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateuser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement",err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user",err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}
	return nil
}

func (user *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement",err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user",err)
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	return nil

}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare finding user by status statement",err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status",err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
	}
	defer rows.Close() // we use it after we have a vald result in rows

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan row into users struct",err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user", errors.New("database error"))
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil

}


func (user *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to get user by email and password statement",err)
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows){
			return  rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user  by email and password",getErr)
		return rest_errors.NewInternalServerError("error when trying to find user", errors.New("database error"))
	}
	return nil
}
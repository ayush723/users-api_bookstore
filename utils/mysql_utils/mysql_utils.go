package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/ayush723/users-api_bookstore/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewBadRequestError(fmt.Sprintf("no record matching given id"))
		}
		return errors.NewInternalServerError("error parsing database respnse")
	}
	switch sqlErr.Number {

	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data"))

	}
	return errors.NewInternalServerError("error processing request")
}

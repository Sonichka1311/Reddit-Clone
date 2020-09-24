package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

func SetSelectIdFromUsersOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}))
}

func SetSelectIdFromUsersWithError(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnError(sql.ErrConnDone)
}

func SetSelectIdFromUsersWithResOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)"}).AddRow("1"))
}

func SetInsertToUsersOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectExec("[INSERT INTO users]").
		WithArgs("Sonia", "testpassword").
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func SetInsertToUsersWithError(mock sqlmock.Sqlmock) {
	mock.
		ExpectExec("[INSERT INTO users]").
		WithArgs("Sonia", "testpassword").
		WillReturnError(sql.ErrConnDone)
}

func SetSelectFromUsersOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}).AddRow("1", "Sonia", "testpassword"))
}

func SetSelectFromUsersWithoutResOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}))
}

func SetSelectFromUsersWithWrongRes(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"HEX(id)", "login", "password"}).AddRow("1", "Sonia", "wrongpassword"))
}

func SetSelectFromUsersWithError(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM users]").
		WithArgs("Sonia").
		WillReturnError(sql.ErrConnDone)
}

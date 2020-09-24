package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

func SetSelectFromSessionsOK(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("testToken").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}))
}

func SetSelectFromSessionsWithResOK(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("testToken").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "expire"}).AddRow("1", "Sonia", expire))
}

func SetSelectFromSessionsByUserOK(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("Sonia").
		WillReturnRows(sqlmock.NewRows([]string{"token", "expire"}).AddRow("testToken", expire))
}

func SetSelectFromSessionsByUserWithError(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("Sonia").
		WillReturnError(sql.ErrConnDone)
}

func SetSelectFromSessionsByUserWithoutRes(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("Sonia").
		WillReturnError(sql.ErrNoRows)
}

func SetSelectFromSessionsWithError(mock sqlmock.Sqlmock) {
	mock.
		ExpectQuery("[SELECT (.+) FROM sessions]").
		WithArgs("testToken").
		WillReturnError(sql.ErrConnDone)
}

func SetInsertToSessionsOK(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectExec("[INSERT INTO sessions]").
		WithArgs("testToken", "1", "Sonia", expire).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func SetInsertToSessionsWithError(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectExec("[INSERT INTO sessions]").
		WithArgs("testToken", "1", "Sonia", expire).
		WillReturnError(sql.ErrConnDone)
}

func SetUpdateSessionsOK(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectExec("[UPDATE sessions]").
		WithArgs("testToken", expire, "Sonia").
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func SetUpdateSessionsWithError(mock sqlmock.Sqlmock, expire string) {
	mock.
		ExpectExec("[UPDATE sessions]").
		WithArgs("testToken", expire, "Sonia").
		WillReturnError(sql.ErrConnDone)
}

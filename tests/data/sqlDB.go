package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reddit/pkg/auth"
)

func InitDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func InitRegisterDb(status int) *sql.DB {
	db, mock := InitDB()

	expire := auth.GenerateTime(12)
	switch status {
	case 0:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersOK(mock)
		SetSelectIdFromUsersWithResOK(mock)
		SetSelectFromSessionsOK(mock)
		SetInsertToSessionsOK(mock, expire)
	case 1:
		SetSelectIdFromUsersWithResOK(mock)
	case 2:
		SetSelectIdFromUsersWithError(mock)
	case 3:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersWithError(mock)
	case 4:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersOK(mock)
		SetSelectIdFromUsersOK(mock)
	case 5:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersOK(mock)
		SetSelectIdFromUsersWithResOK(mock)
		SetSelectFromSessionsWithResOK(mock, expire)
	case 6:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersOK(mock)
		SetSelectIdFromUsersWithResOK(mock)
		SetSelectFromSessionsWithError(mock)
	case 7:
		SetSelectIdFromUsersOK(mock)
		SetInsertToUsersOK(mock)
		SetSelectIdFromUsersWithResOK(mock)
		SetSelectFromSessionsOK(mock)
		SetInsertToSessionsWithError(mock, expire)
	}

	return db
}

func InitLoginDb(status int) *sql.DB {
	db, mock := InitDB()

	expire := auth.GenerateTime(12)

	switch status {
	case 0:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsOK(mock)
		SetSelectFromSessionsByUserOK(mock, expire)
		SetUpdateSessionsOK(mock, expire)
	case 1:
		SetSelectFromUsersWithoutResOK(mock)
	case 2:
		SetSelectFromUsersWithError(mock)
	case 3:
		SetSelectFromUsersWithWrongRes(mock)
	case 4:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsWithResOK(mock, expire)
	case 5:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsWithError(mock)
	case 6:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsOK(mock)
		SetSelectFromSessionsByUserWithError(mock)
	case 7:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsOK(mock)
		SetSelectFromSessionsByUserWithoutRes(mock)
		SetInsertToSessionsOK(mock, expire)
	case 8:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsOK(mock)
		SetSelectFromSessionsByUserWithoutRes(mock)
		SetInsertToSessionsWithError(mock, expire)
	case 9:
		SetSelectFromUsersOK(mock)
		SetSelectFromSessionsOK(mock)
		SetSelectFromSessionsByUserOK(mock, expire)
		SetUpdateSessionsWithError(mock, expire)
	}

	return db
}

func InitAuthDb(status int) *sql.DB {
	db, mock := InitDB()

	expire := auth.GenerateTime(12)

	switch status {
	case 0:
		SetSelectFromSessionsWithResOK(mock, expire)
	case 1:
		SetSelectFromSessionsOK(mock)
	case 2:
		SetSelectFromSessionsWithResOK(mock, auth.GenerateTime(-5))
	}
	return db
}

package database

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_ShouldUpdateStats(test *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE PRODUCT").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO PRODUCT").
		WithArgs(2, 3).
		// WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = RecordStats(db, 2, 3); err != nil {
		test.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// a failing test case
func Test_ShouldRollbackStatUpdatesOnFailure(test *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE PRODUCT").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO PRODUCT").
		WithArgs(2, 3).
		// WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = RecordStats(db, 2, 3); err == nil {
		test.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

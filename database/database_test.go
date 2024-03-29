package database

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Query(test *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT 1 FROM PRODUCT").
		// mock.ExpectQuery(regexp.QuoteMeta("SELECT 1 FROM PRODUCT")).
		WillReturnRows(sqlmock.NewRows(nil))
	// WillReturnError(fmt.Errorf("some error"))

	// now we execute our method
	if err = Query(db); err != nil {
		test.Errorf("error was not expected while updating stats: %s", err)
	}
	// if err = Query(db); err == nil {
	// 	test.Errorf("error was not expected while updating stats: %s", err)
	// }

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Execute_ShouldUpdateStats(test *testing.T) {
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
	if err = Execute(db, 2, 3); err != nil {
		test.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Execute_ShouldRollbackStatUpdatesOnFailure(test *testing.T) {
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
	if err = Execute(db, 2, 3); err == nil {
		test.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_GormQuery(test *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(
		sqlserver.Config{
			Conn: db,
		}), &gorm.Config{})
	sqlString := "SELECT * FROM \"person\" ORDER BY \"person\".\"row\" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY"

	mock.ExpectQuery(regexp.QuoteMeta(sqlString)).
		WillReturnRows(sqlmock.NewRows([]string{"row", "id", "name"}).
			AddRow(0, "", 0))

	// now we execute our method
	GormQuery(gormDB)
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("there were unfulfilled expectations: %s", err)
	}
}

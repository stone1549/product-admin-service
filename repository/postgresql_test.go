package repository_test

import (
	"database/sql"
	"github.com/stone1549/product-admin-service/repository"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
	"time"
)

func mockExpectExecTimes(mock sqlmock.Sqlmock, sqlRegexStr string, times int) {
	for i := 0; i < times; i++ {
		mock.ExpectExec(sqlRegexStr).WillReturnResult(sqlmock.NewResult(int64(i), 1))
	}
}

func makeAndTestPgSmallRepo() (*sql.DB, sqlmock.Sqlmock, repository.ProductRepository, error) {
	var err error
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	mock.ExpectBegin()
	mockExpectExecTimes(mock, "INSERT INTO product", 20)
	mock.ExpectCommit()
	repo, err := repository.MakePostgresqlProductRespository(pgSmall, db)

	return db, mock, repo, err
}

// TestMakePostgresqlProductRespository_Ds ensures that a dataset can be loaded when a pg repo is constructed.
func TestMakePostgresqlProductRespository_Ds(t *testing.T) {
	db, mock, _, err := makeAndTestPgSmallRepo()
	defer db.Close()
	ok(t, err)
	ok(t, mock.ExpectationsWereMet())
}

// TestMakePostgresqlProductRespository ensures that an empty pg repo can be constructed.
func TestMakePostgresqlProductRespository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	_, err = repository.MakePostgresqlProductRespository(pgEmpty, db)
	ok(t, err)
	ok(t, mock.ExpectationsWereMet())
}

func addExpectedProductId1Row(rows *sqlmock.Rows) *sqlmock.Rows {
	createdAt, _ := time.Parse("2006-01-15T15:20:59", "2017-01-01T00:00:00Z")
	updatedAt, _ := time.Parse("2006-01-15T15:20:59", "2018-01-01T00:00:20Z")
	return rows.AddRow(
		"1",
		"Portal Gun",
		"The Portal Gun is a gadget that allows the user(s) to travel between different universes/dimensions/"+
			"realities.\n\nThe Gun was likely created by a Rick, although it is unknown which one; if there is any "+
			"truth to C-137's fabricated origin story, then he may not be the original inventor.",
		"Travel between different dimensions!",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"2499.990000",
		1,
		createdAt,
		updatedAt,
	)
}

func addExpectedProductId2Row(rows *sqlmock.Rows) *sqlmock.Rows {
	createdAt, _ := time.Parse("2006-01-15T15:20:59", "2017-01-01T00:00:01Z")
	updatedAt, _ := time.Parse("2006-01-15T15:20:59", "2018-01-01T00:00:19Z")
	return rows.AddRow(
		"2",
		"Portal Gun",
		"The Portal Gun is a gadget that allows the user(s) to travel between different universes/dimensions/"+
			"realities.\n\nThe Gun was likely created by a Rick, although it is unknown which one; if there is any "+
			"truth to C-137's fabricated origin story, then he may not be the original inventor.",
		"Travel between different dimensions!",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"2499.990000",
		1,
		createdAt,
		updatedAt,
	)
}

func addExpectedProductId3Row(rows *sqlmock.Rows) *sqlmock.Rows {
	createdAt, _ := time.Parse("2006-01-15T15:20:59", "2017-01-01T00:00:02Z")
	updatedAt, _ := time.Parse("2006-01-15T15:20:59", "2018-01-01T00:00:18Z")
	return rows.AddRow(
		"3",
		"Portal Gun",
		"The Portal Gun is a gadget that allows the user(s) to travel between different universes/dimensions/"+
			"realities.\n\nThe Gun was likely created by a Rick, although it is unknown which one; if there is any "+
			"truth to C-137's fabricated origin story, then he may not be the original inventor.",
		"Travel between different dimensions!",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"2499.990000",
		10,
		createdAt,
		updatedAt,
	)
}

func addExpectedProductId4Row(rows *sqlmock.Rows) *sqlmock.Rows {
	createdAt, _ := time.Parse("2006-01-15T15:20:59", "2017-01-01T00:00:03Z")
	updatedAt, _ := time.Parse("2006-01-15T15:20:59", "2018-01-01T00:00:17Z")
	return rows.AddRow(
		"4",
		"Portal Gun",
		"The Portal Gun is a gadget that allows the user(s) to travel between different universes/dimensions/"+
			"realities.\n\nThe Gun was likely created by a Rick, although it is unknown which one; if there is any "+
			"truth to C-137's fabricated origin story, then he may not be the original inventor.",
		"Travel between different dimensions!",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"2499.990000",
		1,
		createdAt,
		updatedAt,
	)
}

func addExpectedProductId5Row(rows *sqlmock.Rows) *sqlmock.Rows {
	createdAt, _ := time.Parse("2006-01-15T15:20:59", "2017-01-01T00:00:04Z")
	updatedAt, _ := time.Parse("2006-01-15T15:20:59", "2018-01-01T00:00:16Z")
	return rows.AddRow(
		"5",
		"Portal Gun",
		"The Portal Gun is a gadget that allows the user(s) to travel between different universes/dimensions/"+
			"realities.\n\nThe Gun was likely created by a Rick, although it is unknown which one; if there is any "+
			"truth to C-137's fabricated origin story, then he may not be the original inventor.",
		"Travel between different dimensions!",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"https://images-na.ssl-images-amazon.com/images/I/31s7nNMzMUL.jpg",
		"2499.990000",
		1,
		createdAt,
		updatedAt,
	)
}

package models

import (
	"testing"

	gormdb "golangapi/db/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {

	db, mock, _ := sqlmock.New()
	gormdb.Db, _ = gorm.Open("mysql", db)
	//sqlRows := sqlmock.NewRows([]string{"details"}).AddRow(`{"name": "foo", "type": "bar", ... }`)
	//mock.ExpectQuery("^SELECT (.+) FROM \"products\" (.+)$").WillReturnRows(sqlRows)
	var cols = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).AddRow(1, "foobar"))

	um := NewUserModel(gormdb.Db)
	u := um.GetUserByID("1")

	//...
	// some http request recording or other operations
	// and then the usual expected := , if ... != t.Errorf combo:
	//expected := `{"products":[{"details":{"name": "foo", "type": "bar"}}]}`

	expect := User{
		ID:   1,
		Name: "foobar",
	}

	assert.Equal(t, expect, u)

	/*
		mockDB, mock, sqlxDB := test.MockDB(t)
		defer mockDB.Close()

		//var cols []string = []string{"id", "name"}
		var cols = []string{"id", "name"}
		mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
			AddRow(1, "foobar"))

		um := NewUserModel(sqlxDB)
		u := um.GetUserByID("1")

		expect := User{
			ID:   1,
			Name: "foobar",
		}
		assert.Equal(t, expect, u)
	*/
}

/*
func TestFindByID(t *testing.T) {
	mockDB, mock, sqlxDB := test.MockDB(t)
	defer mockDB.Close()

	//var cols []string = []string{"id", "name"}
	var cols = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(1, "foobar"))

	um := NewUserModel(sqlxDB)
	u := um.GetUserByID("1")

	expect := User{
		ID:   1,
		Name: "foobar",
	}
	assert.Equal(t, expect, u)
}

func TestFindAll(t *testing.T) {
	mockDB, mock, sqlxDB := test.MockDB(t)
	defer mockDB.Close()

	u1 := User{ID: 1, Name: "foobar"}
	u2 := User{ID: 2, Name: "barbaz"}

	//var cols []string = []string{"id", "name"}
	var cols = []string{"id", "name"}
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(u1.ID, u1.Name).
		AddRow(u2.ID, u2.Name))

	um := NewUserModel(sqlxDB)
	u := um.GetAllUser()

	expect := []User{}
	expect = append(expect, u1, u2)
	assert.Equal(t, expect, u)
}
*/

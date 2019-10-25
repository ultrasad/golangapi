package models

import (

	//"github.com/jinzhu/gorm"

	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	//move to use gorm
	//"golangapi/db"
	"golangapi/db"
	gormdb "golangapi/db/gorm"
)

//BaseModel is default field on table users
type BaseModel struct {
	ID        uint64     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type (

	//UserStore is user interface
	UserStore interface {
		GetUserByID(id string) User
		GetAllUser() []User
	}

	//UserModel ...
	UserModel struct {
		db *gorm.DB
	}

	//User is user
	User struct {
		//BaseModel
		ID         uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		Prefix     string    `json:"prefix"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		CreateDate string    `json:"create_date"`
		Timestamp  time.Time `json:"timestamp" gorm:"column:timestamp" sql:"DEFAULT:current_timestamp"`
	}

	//DBFunc gorm return error
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error

	//Users is user
	/* Users struct {
		Users []User
	} */
)

/*
//Users is user
type Users struct {
	Users []User
}
*/

/*
//DBFunc gorm return error
type (
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error
)
*/

//NewUserModel ...
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// WithinTransaction ...
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func WithinTransaction(fn DBFunc) (err error) {
	tx := gormdb.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

// Create ...
// Helper function to insert gorm model to database by using 'WithinTransaction'
func Create(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !gormdb.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

//CreateUserWithTransection is create user with transection
func CreateUserWithTransection(u *User) (*User, error) {
	err := Create(u)
	return u, err
}

//CreateUser is create user
func CreateUser(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !gormdb.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

//GetUserByID is get user
func (h *UserModel) GetUserByID(id string) User {

	/* fmt.Println("get user...")
	user := User{}
	return user */

	//db := gormdb.ConnectMySQL()
	db := h.db
	//defer db.Close()
	user := User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Last(&user, id).Error
	if err != nil {
		fmt.Print("Connect DB Error::", err)
	}

	//result.Users = user
	//fmt.Println("user => ", user)

	return user
}

//GetAllUser is get all user
func (h *UserModel) GetAllUser() []User {
	//db := gormdb.ConnectMySQL()
	db := h.db
	//defer db.Close()
	//result := Users{}
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print(err)
	}

	//result.Users = user
	//fmt.Println("User => ", user)

	return user
}

/* func (h *UserModel) GetAllUser() Users {
	//db := gormdb.ConnectMySQL()
	db := h.db
	defer db.Close()
	result := Users{}
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print(err)
	}

	result.Users = user
	//fmt.Println("User => ", user)

	return result
} */

//GetUserMain is get user
func (h *UserModel) GetUserMain() []User {
	//db := gormdb.ConnectMySQL()
	db := h.db
	//defer db.Close()
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print("error db debug => ", err)
	}

	for _, ar := range user {
		fmt.Println(ar.ID)
		user = append(user, ar)
	}

	return user
}

/* func GetUserMain() Users {
	db := gormdb.ConnectMySQL()
	defer db.Close()
	result := Users{}
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print("error db debug => ", err)
	}

	for _, ar := range user {
		fmt.Println(ar.ID)
		result.Users = append(result.Users, ar)
	}

	return result
} */

//RowX ...
type (
	RowX struct {
		ID    int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		Name  string `json:"name"`
		TypeX string `json:"typeX"`
		Owner string `json:"owner"`
	}
)

//RowsX is ...
type RowsX struct {
	RowsX []RowX
}

//GetUserDefault is get user
func GetUserDefault() RowsX {
	//from db/connection
	con := db.CreateCon()

	sqlStatement := "SELECT id, name, type as typeX, owner FROM campaign_rules order by id limit 10"

	rows, err := con.Query(sqlStatement)
	//fmt.Println(rows)
	//fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	result := RowsX{}

	for rows.Next() {
		r := RowX{}
		err := rows.Scan(&r.ID, &r.Name, &r.TypeX, &r.Owner)
		if err != nil {
			fmt.Print(err)
		}

		result.RowsX = append(result.RowsX, r)
	}

	fmt.Println(result)

	return result
}

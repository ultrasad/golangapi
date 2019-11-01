package _models

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

const ctLayout = "2006-01-02 15:04:05 Z07:00"

type (

	//UserStore is user interface
	UserStore interface {
		GetUserByID(id string) User
		GetAllUser() []User
	}

	//Marshaler is json marshal
	Marshaler interface {
		MarshalJSON() ([]byte, error)
	}

	//UserModel ...
	UserModel struct {
		db *gorm.DB
	}

	//JSONTime is json time custom
	/* JSONTime struct {
		*time.Time
	} */

	//CustomTime is custom time
	CustomTime struct {
		time.Time
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
		//Timestamp CustomTime `json:"timestamp" gorm:"column:timestamp" sql:"DEFAULT:current_timestamp"`
	}

	// myTime is custom datetime
	/* myTime struct {
		time.Time
	} */

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

/* func (t myTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format(time.RFC3339) + `"`), nil
} */

// MarshalJSON is custom json struct
/* func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(stamp), nil
} */

// MarshalJSON is custom json struct
/* func (d *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		Timestamp string `json:"stamp"`
	}{
		Alias:     (*User)(d),
		Timestamp: d.Timestamp.Format("Mon Jan _2"),
	})
} */

//UnmarshalJSON custom datetime
/* func (t *SpecialDateTime) UnmarshalJSON(buf []byte) error {
	//tt, err := time.Parse(time.RFC1123, strings.Trim(string(buf), `"`))
	tt, err := time.Parse(time.RFC3339, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
} */

/* // UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = CustomTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
} */

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
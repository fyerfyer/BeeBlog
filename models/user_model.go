package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"golang.org/x/crypto/bcrypt"
	// "github.com/jinzhu/gorm"
)

type User struct {
	Id             int    `orm:"auto"` // 自增主键
	Username       string // 用户名必填，长度不超过20个字符
	Email          string // Email必填且符合Email格式
	HashedPassword string
	Status         int
	CreateTime     time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
}

// type Err struct {
// 	hasError bool
// 	errMsg   string
// }

func init() {
	orm.RegisterModel(new(User))
}

func CreateUser(user *User) (int, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		log.Println("CreateUser failed")
	}
	return int(id), err
}

func GetUserById(id int) (*User, error) {
	o := orm.NewOrm()
	user := User{}

	err := o.QueryTable("user").Filter("id", id).One(&user)
	if err != nil {
		log.Printf("GetUserById failed: %v", err)
		return nil, err
	}

	return &user, nil
}

func getUserByEmail(email string) (*User, error) {
	o := orm.NewOrm()
	user := User{}

	err := o.QueryTable("user").Filter("email", email).One(&user)
	if err != nil {
		log.Printf("getUserByEmail failed: %v", err)
		return nil, err
	}

	return &user, nil
}

func GetUserStatusById(id int) (int, error) {
	user, err := GetUserById(id)
	if err != nil {
		log.Printf("GetUserStatusById failed: %v", err)
		return 0, err
	}

	return user.Status, nil
}

func (u *User) Valid() map[string]string {
	errorMap := make(map[string]string)
	valid := validation.Validation{}
	if u.Username == "" {
		errorMap["Name"] = "Username cannot be empty"
	} else if len(u.Username) > 15 {
		errorMap["Name"] = "Username cannot be longer than 15 characters"
	}

	if u.Email == "" {
		errorMap["Email"] = "Email cannot be empty"
	} else if !valid.Email(u.Email, "Email").Ok {
		errorMap["Email"] = "Invalid email format"
	}
	// we pass the unhashed value to the user in the validation step first
	if u.HashedPassword == "" {
		errorMap["Password"] = "Password cannot be empty"
	} else if len(u.HashedPassword) < 10 {
		errorMap["Password"] = "Password must be longer than 12 characters"
	}

	return errorMap
}

// there's no need to fecch other field?
func UserAuthenticate(email, password string) (int, bool, bool) {
	if email == "" {
		log.Printf("Email cannot be empty")
		return 0, false, false
	}

	user, err := getUserByEmail(email)
	if err != nil {
		log.Printf("Failed to get the user: %v", err)
		return 0, false, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return 0, true, false
	}

	return user.Id, true, true
}

func SetUserStatusById(id int, status int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("id", id).Update(orm.Params{"status": status})
	if err != nil {
		log.Printf("Failed to deactive user: %v", err)
	}

	return err
}

package data

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	FullName     string    `json:"fullName" validate:"required"`
	Nicknamne    string    `json:"nickname" validate:"required"`
	Password     string    `json:"password" validate"Password"`
	Permissions  []int     `json:"permissions" validate:"required"`
	RegisteredOn time.Time `json:"-"`
}

func (c *User) UserFromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}
func (c *User) UserToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func GetUser(id int64) (User, error) {
	command := "SELECT id, full_name, nickname, password, permissions FROM cars WHERE id = $1;"
	rows, err := db.Query(command, id)
	defer rows.Close()
	if err != nil {
		return User{}, err
	}
	user := User{}
	rows.Next()
	rows.Scan(&user.Id, &user.FullName, &user.Nicknamne, &user.Password, &user.Permissions)
	return user, nil
}

func AddUser(user *User) error {
	command := "INSERT INTO users(full_name, nickname, password, permissions, registered_on) " +
		"VALUES ($1, $2, $3, $4, $5);"
	_, err := db.Exec(command, user.FullName, user.Nicknamne, user.Password, user.Permissions, user.RegisteredOn)
	return err
}

func DeleteUser(id int64) error {
	command := "DELETE FROM users WHERE id = $1;"
	_, err := db.Exec(command, id)
	return err
}

func UpdateUser(user User) error {
	command := "UPDATE users SET full_name=$1, nickname=$2, password=$3, permissions=$4 WHERE id=$5;"
	_, err := db.Exec(command, user.FullName, user.Nicknamne, user.Password, user.Permissions, user.Id)
	return err
}

func UpdateUserPermissions(user User) error {
	command := "UPDATE users SET permissions=$1 WHERE id=$2;"
	_, err := db.Exec(command, user.Permissions, user.Id)
	return err
}

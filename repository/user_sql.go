// DBに対してsqlを発行するクラス
package repository

import (
	"gomodtest/db"
	"gomodtest/models"
)

type UserRepository struct{}

type User models.User

type UserProfile struct {
	Name string
	Id   int
}

func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Table("users").Select("name, id").Scan(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// ユーザオブジェクトの定義
package models

type User struct {
	Id    uint   `json:"id" binding:"requried"`
	Name  string `json:"name" binding:"required"`
	Posts []Post `json:"posts"`
}

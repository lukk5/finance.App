package entity

import (
	"encoding/json"
	"finance-api-v1/core/common"
	"fmt"
)

type User struct {
	common.Base
	UserName string `gorm:"column:userName"`
}

func (user User) String() string {
	return fmt.Sprintf("User{id:%d, username:%s, createdAt:%v, updatedAt:%v",
		user.ID, user.UserName, user.CreatedAt, user.UpdatedAt)
}

func (user *User) UnmarshalJSON(b []byte) error {
	var tmp struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	user.ID = tmp.ID
	user.UserName = tmp.Username
	return nil
}

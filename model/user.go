/**
* @Author: mo tan
* @Description:
* @Date 2021/1/9 18:15
 */
package model

import "time"

type User struct {
	Id        	int64       `db:"id"`
	Account     string   `db:"account"`
	Mobile	    string   `db:"mobile"`
	Nick     	string    `db:"nick"`
	Password	string 	  `db:"password"`
	Status    	int       `db:"status"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) PK() string {
	return "id"
}

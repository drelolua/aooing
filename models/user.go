package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article 文章管理结构
type User struct {
	DbBase      `bson:",omitempty"`
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name       string        `bson:",omitempty" json:"name"`        //标题
	Age        int        `bson:",omitempty" json:"age"`         // 内容
	Create_at time.Time     `bson:",omitempty" json:"create_at"` // 发布日期
}

// CName 获取当前集合名称
func (u *User) CName() string {
	return "users"
}

// Insert 插入记录
func (u *User) Insert() error {
	return u.Collection(u.CName()).Insert(u)
}

//Update 更新记录
func (u *User)Update()error{
	return u.Collection(u.CName()).Update(u.ID, u)
}

// Delete 删除记录
func (u *User)Delete() error{
	return u.Collection(u.CName()).RemoveId(u.ID)
}

// GetSingleData 获取单条记录
func (u *User)GetSingleData()(User, error){
	var user User
	err := u.Collection(u.CName()).FindId(u.ID).One(&user)
	return user, err
}

//GetData 获取所有数据
func (u *User)GetData(data *[]User) error{
	return u.Find(u.CName(), nil, nil).All(data)
}

func (u *User)AllCount()(int, error){
	return u.Find(u.CName(), nil, nil).Count()
}

// GetPageData 获取所有分页数据
func (u *User)GetPageData(skip, limit int, data *[]User)error{
	return u.Find(u.CName(), nil, nil).Skip(skip).Limit(limit).All(data)
}


package model

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"path"
	"server/public"
	"strings"
	"sync"
)

type User struct {
	Id       uint   `xml:"id" json:"-"`
	UserName string `xml:"user_name" json:"user_name"`
	Password string `xml:"password" json:"password"`
	Name     string `xml:"name" json:"name"`
	Email    string `xml:"email" json:"email"`
}

type UsersData struct {
	XMLName xml.Name `xml:"data"`
	Users   []*User  `xml:"user"`
}

func NewUsersData() *UsersData {
	return &UsersData{
		Users: make([]*User, 0),
	}
}

var lock sync.Mutex

// LoadFromXml 从xml文件中加载初始化数据
func (m *UsersData) LoadFromXml(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, &m)
	if err != nil {
		return err
	}
	return nil
}

// SaveToXml 存储数据到XML文件
func (m *UsersData) SaveToXml(path string) error {
	data, err := xml.MarshalIndent(&m, "", "	")
	if err != nil {
		return err
	}
	headerBytes := []byte(xml.Header)
	xmlData := append(headerBytes, data...)
	return ioutil.WriteFile(path, xmlData, 0666)
}

// GetUserById 通过id查找用户
func (m *UsersData) GetUserById(id uint) *User {
	for _, v := range m.Users {
		if v.Id == id {
			return v
		}
	}
	return nil
}

// GetMaxId 获取最大ID
func (m *UsersData) GetMaxId() (res uint) {
	for _, v := range m.Users {
		if v.Id > res {
			res = v.Id
		}
	}
	return
}

// AddUser 写入新用户
func (m *UsersData) AddUser(user *User) error {
	lock.Lock() //增加用户时加锁，防止并发错误
	defer lock.Unlock()
	if m.HasUserName(user.UserName) { //用户名唯一性检查
		return errors.New("User name already exists ")
	}
	user.Id = m.GetMaxId() + 1 //修改ID递增到当前最大值+1，强制修改防止重复
	m.Users = append(m.Users, user)
	err := m.SaveToXml(path.Join(public.DATA_FILE_PATH, public.DATA_FILE_NAME))
	if err != nil {
		return err
	}
	return nil
}

// HasUserName 检查用户名是否已经存在
func (m *UsersData) HasUserName(username string) bool {
	for _, v := range m.Users {
		if strings.TrimSpace(v.UserName) == strings.TrimSpace(username) { //清除空白字符
			return true
		}
	}
	return false
}

// GetUserByUserName 查找指定用户名用户，若没有返回nil
func (m *UsersData) GetUserByUserName(userName string) *User {
	for _, v := range m.Users {
		if v.UserName == strings.TrimSpace(userName) {
			return v
		}
	}
	return nil
}

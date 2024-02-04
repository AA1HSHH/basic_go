package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱重复")
	ErrRecordNotFund      = gorm.ErrRecordNotFound
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}
func (dao *UserDao) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error

	var sqlErr *mysql.MySQLError
	if errors.As(err, &sqlErr) {
		if sqlErr.Number == 1062 {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}
func (dao *UserDao) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&u).Error
	return u, err
}

func (dao *UserDao) AddInfo(ctx context.Context, id int64, me string, birthday string, nickname string) error {
	log.Println("AddInfo")
	err := dao.db.WithContext(ctx).Table("users").Where("id=?", id).
		Updates(map[string]interface{}{"nickname": nickname, "birthday": birthday, "about_me": me}).Error
	return err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
	Nickname string
	Birthday string
	AboutMe  string
}

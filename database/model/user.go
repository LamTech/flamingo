package model

import (
	"bytes"
	"compress/flate"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flamingo/database"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"io"
	"log"

	"github.com/boltdb/bolt"
)

const (
	dbName     = "myBlog.db"
	userBucket = "user"
)

// User 用户类
type User struct {
	gorm.Model
	UniqueId uuid.UUID
	Name     string `json:"name"`
	Gender   uint `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// LoginReq 登录请求参数类
type LoginRequest struct {
	Mobile string `json:"mobile"`
	PassWord   string `json:"password"`
}

// compressByte returns a compressed byte slice.
func CompressByte(src []byte) []byte {
	compressedData := new(bytes.Buffer)
	compress(src, compressedData, 9)
	return compressedData.Bytes()
}

// compress uses flate to compress a byte slice to a corresponding level
func compress(src []byte, dest io.Writer, level int) {
	compressor, _ := flate.NewWriter(dest, level)
	compressor.Write(src)
	compressor.Close()
}

// decompressByte returns a decompressed byte slice.
func DecompressByte(src []byte) []byte {
	compressedData := bytes.NewBuffer(src)
	deCompressedData := new(bytes.Buffer)
	decompress(compressedData, deCompressedData)
	return deCompressedData.Bytes()
}

// compress uses flate to decompress an io.Reader
func decompress(src io.Reader, dest io.Writer) {
	decompressor := flate.NewReader(src)
	io.Copy(dest, decompressor)
	decompressor.Close()
}

// 序列化
func dumpUser(user User) []byte {
	dumped, _ := user.MarshalJSON()
	return CompressByte(dumped)
}

// 反序列化
func loadUser(jsonByte []byte) User {
	res := User{}
	res.UnmarshalJSON(DecompressByte(jsonByte))
	return res
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(mobile string, password string) error {
	if CheckUser(mobile) {
		return fmt.Errorf("该手机号已经注册！")
	}

	newUser := User{
		UniqueId : uuid.New(),
		Name : mobile,
		Gender:0,
		Mobile:mobile,
		Password:password,
	}

	return database.DB.Create(&newUser).Error
}

// CheckUser 检查用户是否存在
func CheckUser(mobile string) bool {
	count := 0
	database.DB.Model(&User{}).Where("mobile = ?", mobile).Count(&count)
	result := false
	if count>0 {
		result = true
	}
	return result
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginRequest) (bool, User, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	resultUser := User{}
	resultBool := false
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userBucket))
		if bucket == nil {
			return fmt.Errorf(" userBuket is null")
		}
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			userTemp := loadUser(v)
			if loginReq.Mobile == userTemp.Mobile && loginReq.PassWord == userTemp.Password {
				resultUser = userTemp
				resultBool = true
				break
			}
		}
		if !resultBool {
			return fmt.Errorf("用户信息错误!")
		} else {
			return nil
		}
	})
	return resultBool, resultUser, err
}
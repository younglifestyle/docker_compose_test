package tests

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

// mysql 启动过慢
func conn() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	for {
		db, err = gorm.Open("mysql", "root:123@tcp(db:3306)/test")
		if err != nil {
			time.Sleep(time.Second)
			continue
		} else {
			return db
		}
	}
}

func conn1(t *testing.T) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open("mysql", "root:123@tcp(db:3306)/test")
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestMysqlFind(t *testing.T) {
	fmt.Println("Go MySQL Tutorial")

	//time.Sleep(100 * time.Second)

	//db := conn()
	db := conn1(t)

	resp := make([]map[string]interface{}, 0)
	//条件查询
	err := db.Table("users").Where("username = ?", "xiao").Find(&resp).Error
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	fmt.Println(resp)
}

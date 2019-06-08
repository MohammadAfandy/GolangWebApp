package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"myapp/app/models"
	"database/sql"
    // "fmt"
)

type GormController struct {
    *revel.Controller
    Txn *gorm.DB
}

var Gdb *gorm.DB

func InitDB() {
    var err error
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
    db.LogMode(true) // Print SQL statements
	if err != nil {
        panic(err)
	}
	db.AutoMigrate(&models.Post{})

    Gdb = db
}

// transactions
// This method fills the c.Txn before each transaction
func (c *GormController) Begin() revel.Result {
    txn := Gdb.Begin()
    if txn.Error != nil {
        panic(txn.Error)
    }
    c.Txn = txn
    return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Commit()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Rollback()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}
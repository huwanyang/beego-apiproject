package models

import (
	"apiproject/bean"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	BookList []bean.Book
)

// 添加 book
func AddBook(book bean.Book) (bid int64, err error) {
	o := orm.NewOrm()
	rs, err := o.Raw("insert into book set name = ?, description = ?, price = ?", book.Name, book.Description, book.Price).Exec()
	if err == nil {
		bid, _ := rs.LastInsertId()
		return bid, nil
	}
	return 0, errors.New("Insert book error.")
}

// 获取所有的 book 列表
func GetAllBook() []bean.Book {
	o := orm.NewOrm()
	num, err := o.Raw("select id,name,description,price from book").QueryRows(&BookList)
	if err == nil && num > 0 {
		for index, book := range BookList {
			fmt.Printf("index: %d, id: %d, name: %s, description: %s, price: %.2f\n",
				index+1, book.Id, book.Name, book.Description, book.Price)
		}
	}
	return BookList
}

// 根据 id 获取 book
func GetBookById(bid int) (b bean.Book, err error) {
	var book bean.Book
	o := orm.NewOrm()
	errs := o.Raw("select id,name,description,price from book where id = ?", bid).QueryRow(&book)
	if errs == nil {
		return book, nil
	}
	return b, errors.New("Book not exists")
}

// 更新 book
func UpdateBookById(bid int, book bean.Book) error {
	o := orm.NewOrm()
	_, err := o.Raw("update book set name = ?, description = ?, price = ? where id = ?",
		book.Name, book.Description, book.Price, bid).Exec()
	if err == nil {
		return nil
	}
	return errors.New("Update book fail")
}

// 删除 book
func DeleteBookById(bid int) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from book where id = ?", bid).Exec()
	if err == nil {
		return nil
	}
	return errors.New("Delete book fail")
}

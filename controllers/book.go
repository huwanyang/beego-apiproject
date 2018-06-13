package controllers

import (
	"apiproject/bean"
	"apiproject/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// 关于 Book 相关接口操作
type BookController struct {
	beego.Controller
}

// @Title Create Book
// @Description 创建一个新的 Book
// @Param	body	body	bean.Book	true	"针对 Book 的 Post 请求 Body"
// @Success 200 {int} bean.Book.id
// @Failure 403 Body 为空
// @router / [post]
func (b *BookController) PostBook() {
	var book bean.Book
	json.Unmarshal(b.Ctx.Input.RequestBody, &book)
	bid, err := models.AddBook(book)
	if err == nil {
		b.Data["json"] = bean.ResultVO{true, "Add user success", map[string]interface{}{"bid": bid}}
	} else {
		b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	b.ServeJSON()
}

// @Title Get all books
// @Description 获取所有 Book 列表
// @Success 200 {object} bean.Book
// @router / [get]
func (b *BookController) GetAllBook() {
	bookLit := models.GetAllBook()
	b.Data["json"] = bean.ResultVO{true, "", bookLit}
	b.ServeJSON()
}

// @Title GetBook by Id
// @Description 根据 id 获取 Book
// @Param	bid	path	int	true	"根据 path 的 id 获取 book 信息"
// @Success 200 {object} bean.Book
// @Failure 403 bid is not int
// @router /:bid [get]
func (b *BookController) GetBook() {
	bid, err := b.GetInt(":bid")
	if err == nil && bid != 0 {
		book, err := models.GetBookById(bid)
		if err != nil {
			b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		} else {
			b.Data["json"] = bean.ResultVO{true, "", book}
		}
	} else {
		b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	b.ServeJSON()
}

// @Title Update book by id
// @Description 根据 id 更新 book
// @Param	bid	path	int true	"根据 Path 的 id 更新 Book 信息"
// @Param	body	body	bean.Book	true	"针对 Book 的 Put 请求 Body"
// @Success 200 {object} bean.Book
// @Failure 403 bid is not int
// @router /:bid [put]
func (b *BookController) UpdateBook() {
	var book bean.Book
	bid, err := b.GetInt(":bid")
	if err == nil && bid != 0 {
		json.Unmarshal(b.Ctx.Input.RequestBody, &book)
		err := models.UpdateBookById(bid, book)
		if err == nil {
			b.Data["json"] = bean.ResultVO{true, "Update book success", book}
		} else {
			b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	b.ServeJSON()
}

// @Title Delete book
// @Description 根据 id 删除 book
// @Param	bid	path	int	true	"根据 Path 的 id 删除 Book 信息"
// @Success 200 string delete success
// @Failure 403 bid is not int
// @router /:bid [delete]
func (b *BookController) DeleteBook() {
	bid, err := b.GetInt(":bid")
	if err == nil && bid != 0 {
		err := models.DeleteBookById(bid)
		if err == nil {
			b.Data["json"] = bean.ResultVO{true, "Delete book success", ""}
		} else {
			b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
		}
	} else {
		b.Data["json"] = bean.ResultVO{false, err.Error(), ""}
	}
	b.ServeJSON()
}

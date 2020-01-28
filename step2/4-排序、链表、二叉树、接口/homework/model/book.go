package model

import "errors"

var (
	ErrStockNotEnough = errors.New("stock is not enough")
)

type Book struct {
	Name   string "书名"
	Num    int    "副本数"
	Author string "作者"
	Date   string "出版日期"
}

func CreateBook(name, author, date string, num int) *Book {
	return &Book{
		Name:   name,
		Num:    num,
		Author: author,
		Date:   date,
	}
}
func (self *Book) CanBorrow(n int) bool {
	return self.Num > n
}

func (self *Book) Borrow(n int) (err error) {
	if self.CanBorrow(n) == false {
		err = ErrStockNotEnough
		return
	}
	self.Num -= n
	return
}
func (self *Book) Back(n int) {
	self.Num += n
}

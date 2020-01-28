package model

type Student struct {
	Name, ID, Gender string "姓名、身份证、性别"
	Grade            int    "年级"
	books            []*BorrowItem
}
type BorrowItem struct {
	book *Book
	num  int
}
func (s * Student) DelBook(b *BorrowItem)(err error){
	for i:=0;i<len(s.books);i++{
		if s.books[i].book.Name == b.book.Name{
			if b.num == s.books[i].num{
				s.books = append(s.books[0:i],s.books[i+1:]...)
				return
			}
			s.books[i].num -= b.num
		}
	}
	err = ErrStockNotEnough
	return err
}
func CreateStudent(name, id, gender string, grade int) *Student {
	return &Student{
		Name:   name,
		ID:     id,
		Gender: gender,
		Grade:  grade,
	}
}

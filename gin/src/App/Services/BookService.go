package Services

import (
	"topic/src/AppInit"
	. "topic/src/Models"
)

type BookService struct {
	
}

func(this *BookService) GetBookList(request *BookListRequest) *BookList {
	products := &BookList{}
	AppInit.DB.Limit(request.Size).Order("book_id desc").Find(products)
	return products
}

func(this *BookService) GetBookDetail(request *BookDetailRequest) (*Books, *BookMetas, error) {
	products := NewBooks()
	if AppInit.DB.Find(products, request.BookID).RowsAffected == 0 {
		return nil, nil, nil
	}

	err := NewBookMeta("click","1",products.BookID).Save()
	bookMeta := &BookMetas{}
	AppInit.DB.Find(bookMeta,"item_id = ?", products.BookID)
	return products, bookMeta, err
}
package service

import (
	"github.com/akshay/libraryAssign/internal/library/middleware/auth"
	"github.com/akshay/libraryAssign/internal/library/models/dto"
	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetBooks(ctx *gin.Context) (dto.BookNames, error) {
	claims, err := auth.GetClaimsFromContext(ctx)
	if err != nil {
		return dto.BookNames{}, err
	}
	var books []string

	if claims.Role == dto.ADMIN {
		adminBooksList, err := utils.ReadCSV(dto.ADMIN)
		if err != nil {
			return dto.BookNames{}, err
		}

		for _, book := range adminBooksList {
			books = append(books, book.Name)
		}
	}

	regularBookList, err := utils.ReadCSV((dto.REGULAR))
	if err != nil {
		return dto.BookNames{}, err
	}

	for _, book := range regularBookList {
		books = append(books, book.Name)
	}

	BookList := dto.BookNames{
		Books: books,
	}

	return BookList, nil
}

func (s *Service) AddBooks(ctx *gin.Context, req *dto.BooksList) error {
	booksList := dto.BooksList{
		Books: req.Books,
	}
	err := utils.WriteCSV(dto.REGULAR, booksList)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteBook(ctx *gin.Context, req *dto.BookName) error {
	bookName := req.Book
	err := utils.DeleteBook(dto.REGULAR, bookName)
	if err != nil {
		return err
	}
	return nil
}

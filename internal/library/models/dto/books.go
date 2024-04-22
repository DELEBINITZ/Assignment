package dto

type Book struct {
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publication_year" validate:"required"`
}

type BookNames struct {
	Books []string `json:"books"`
}

type BooksList struct {
	Books []Book `json:"books"`
}

type BookName struct {
	Book string `json:"book_name"`
}

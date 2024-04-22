package dao

type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear string `json:"publication_year"`
}

type BooksList struct {
	Books []Book `json:"Books"`
}

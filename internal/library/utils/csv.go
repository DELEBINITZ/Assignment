package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akshay/libraryAssign/internal/library/models/dto"
)

func ReadCSV(fileName string) ([]dto.Book, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	newPath := strings.Replace(workingDir, "/cmd/library", "", -1)

	file, err := os.Open(fmt.Sprintf("%s/internal/library/infrastructure/datastore/%s.csv", newPath, fileName))
	if err != nil {
		fmt.Println("error while opening the file")
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var books []dto.Book

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
		return nil, err
	}

	for _, record := range records {
		PublicationYear, _ := strconv.Atoi(record[2])
		book := dto.Book{
			Name:            record[0],
			Author:          record[1],
			PublicationYear: PublicationYear,
		}
		books = append(books, book)
	}

	return books, nil
}

func WriteCSV(fileName string, data dto.BooksList) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	newPath := strings.Replace(workingDir, "/cmd/library", "", -1)

	file, err := os.OpenFile(fmt.Sprintf("%s/internal/library/infrastructure/datastore/%s.csv", newPath, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("error while opening the file")
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	if _, err := file.Stat(); os.IsNotExist(err) {
		header := []string{"Book Name", "Author", "Publication Year"}
		writer.Write(header)
	}

	var dataForCSV [][]string
	for _, book := range data.Books {
		dataForCSV = append(dataForCSV, []string{book.Name, book.Author, strconv.Itoa(book.PublicationYear)})
	}

	err = writer.WriteAll(dataForCSV)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

func DeleteBook(fileName, bookName string) error {
	data, err := ReadCSV(fileName)
	if err != nil {
		return err
	}

	filteredData := dto.BooksList{Books: []dto.Book{}}
	for _, row := range data {
		if !strings.EqualFold(row.Name, bookName) {
			filteredData.Books = append(filteredData.Books, row)
		}

	}

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	newPath := strings.Replace(workingDir, "/cmd/library", "", -1)

	filePath := fmt.Sprintf("%s/internal/library/infrastructure/datastore/%s.csv", newPath, fileName)
	err = os.Remove(filePath)
	if err != nil && !os.IsNotExist(err) { // Ignore "no such file" error
		return fmt.Errorf("failed to delete file: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return WriteCSV(fileName, filteredData)
}

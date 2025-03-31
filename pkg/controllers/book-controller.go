package controllers

import (
	"encoding/json"
	"github.com/garvit4540/go-postgres-bookstore/pkg/models"
	"github.com/garvit4540/go-postgres-bookstore/pkg/trace"
	"github.com/garvit4540/go-postgres-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	trace.PushTrace(trace.CreateBookRequestReceived, map[string]interface{}{"new_book": newBook})

	createdBook := newBook.CreateBook()
	trace.PushTrace(trace.BookSuccessFullyCreated, map[string]interface{}{"created_book": createdBook})

	res, _ := json.Marshal(createdBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {

	trace.PushTrace(trace.FetchAllBooksRequestReceived, nil)

	newBooks := models.GetAllBooks()
	trace.PushTrace(trace.BooksFetchedSuccessfully, map[string]interface{}{"fetched_books": newBooks})

	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["book_id"]
	trace.PushTrace(trace.FetchBookRequestReceived, map[string]interface{}{"book_id": bookId})

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		trace.PushTrace(trace.ErrorJsonMarshalUnMarshall, map[string]interface{}{"error": err})
	}
	bookDetails, _ := models.GetBookById(id)

	trace.PushTrace(trace.FetchBookSuccessfull, map[string]interface{}{"book_details": *bookDetails})

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	prevBookId := vars["book_id"]

	newBookConfig := &models.Book{}
	utils.ParseBody(r, newBookConfig)

	trace.PushTrace(trace.UpdateBookRequestReceived, map[string]interface{}{"book_id": prevBookId})

	id, err := strconv.ParseInt(prevBookId, 0, 0)
	if err != nil {
		trace.PushTrace(trace.ErrorJsonMarshalUnMarshall, map[string]interface{}{"error": err})
	}
	prevBookDetails, db := models.GetBookById(id)

	if newBookConfig.Name != "" {
		prevBookDetails.Name = newBookConfig.Name
	}
	if newBookConfig.Author != "" {
		prevBookDetails.Author = newBookConfig.Author
	}
	if newBookConfig.Publication != "" {
		prevBookDetails.Publication = newBookConfig.Publication
	}

	db.Save(&prevBookDetails)
	res, _ := json.Marshal(newBookConfig)

	trace.PushTrace(trace.BookUpdatedSuccessfully, map[string]interface{}{"new_config": prevBookDetails})

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["book_id"]
	trace.PushTrace(trace.DeleteBookRequestReceived, map[string]interface{}{"book_id": bookId})

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		trace.PushTrace(trace.ErrorJsonMarshalUnMarshall, map[string]interface{}{"error": err})
	}

	book := models.DeleteBook(id)
	trace.PushTrace(trace.BookSuccessFullyDeleted, map[string]interface{}{"book": book})
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)

}

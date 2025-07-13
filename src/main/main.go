package main

import (
	"database/sql"
	"fmt"
	"github.com/kyliancc/kyc-beginia/src/model"
	"github.com/kyliancc/kyc-beginia/src/repository"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//http.HandleFunc("/", hello)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	db, err := sql.Open("sqlite3", "./identifier.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todoDocRepo := repository.NewTodoDocsRepo(db)

	labels := []string{"电子", "工科"}

	doc := model.DocItem{Name: "数字电子技术基础", Priority: 2, Labels: labels}
	id, err := todoDocRepo.CreateTodoDoc(&doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	docs, err := todoDocRepo.QueryAllTodoDocs()
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range docs {
		fmt.Println(doc)
	}
}

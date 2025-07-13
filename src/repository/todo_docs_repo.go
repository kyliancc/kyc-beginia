package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kyliancc/kyc-beginia/src/model"
)

type TodoDocsRepo struct {
	db *sql.DB
}

func NewTodoDocsRepo(db *sql.DB) *TodoDocsRepo {
	return &TodoDocsRepo{db: db}
}

func (r *TodoDocsRepo) CreateTodoDoc(doc model.DocItem) (id int, err error) {
	stmt, err := r.db.Prepare("INSERT INTO todo_docs(name, comment, priority, labels) VALUES(?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare creation: %w", err)
	}
	defer stmt.Close()

	jsonLabels, err := json.Marshal(doc.Labels)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal labels: %w", err)
	}

	res, err := stmt.Exec(doc.Name, doc.Comment, doc.Priority, jsonLabels)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into todo_docs: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	fmt.Printf("Created todo doc with id %d", lastID)
	return int(lastID), nil
}

func (r *TodoDocsRepo) DeleteTodoDoc(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM todo_docs WHERE id=?")
	if err != nil {
		return fmt.Errorf("failed to prepare deletion: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to delete todo_docs: %w", err)
	}

	fmt.Printf("Deleted todo doc with id %d", id)
	return nil
}

func (r *TodoDocsRepo) UpdateTodoDoc(doc model.DocItem) error {
	stmt, err := r.db.Prepare("UPDATE todo_docs SET name=?, comment=?, priority=?, labels=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("failed to prepare update: %w", err)
	}
	defer stmt.Close()

	jsonLabels, err := json.Marshal(doc.Labels)
	if err != nil {
		return fmt.Errorf("failed to marshal labels: %w", err)
	}

	_, err = stmt.Exec(doc.Name, doc.Comment, doc.Priority, jsonLabels, doc.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo_docs: %w", err)
	}

	fmt.Printf("Updated todo doc with id %d", doc.ID)
	return nil
}

func (r *TodoDocsRepo) QueryAllTodoDocs() (docs []model.DocItem, err error) {
	stmt, err := r.db.Prepare("SELECT * FROM todo_docs")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	for res.Next() {
		var doc model.DocItem
		var rawLabels string
		if err := res.Scan(&doc.ID, &doc.Created, &doc.Name, &doc.Comment, &doc.Priority, &rawLabels); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		err = json.Unmarshal([]byte(rawLabels), &doc.Labels)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal labels: %w", err)
		}

		docs = append(docs, doc)
	}

	fmt.Printf("Found %d docs", len(docs))
	return docs, nil
}

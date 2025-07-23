package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kyliancc/kyc-beginia/src/model"
)

type TodoDocsRepo struct {
	db *sql.DB
}

func NewTodoDocsRepo(db *sql.DB) *TodoDocsRepo {
	return &TodoDocsRepo{db: db}
}

func (r *TodoDocsRepo) CreateTodoDoc(doc *model.DocItem) (id int, err error) {
	stmt, err := r.db.Prepare("INSERT INTO todo_docs(name, comment, priority, labels) VALUES(?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare insert: %w", err)
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
		return fmt.Errorf("failed to delete doc: %w", err)
	}

	fmt.Printf("Deleted todo doc with id %d", id)
	return nil
}

func (r *TodoDocsRepo) UpdateTodoDoc(doc *model.DocItem) error {
	stmt, err := r.db.Prepare("UPDATE todo_docs SET name=?, comment=?, labels=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("failed to prepare update: %w", err)
	}
	defer stmt.Close()

	jsonLabels, err := json.Marshal(doc.Labels)
	if err != nil {
		return fmt.Errorf("failed to marshal labels: %w", err)
	}

	_, err = stmt.Exec(doc.Name, doc.Comment, jsonLabels, doc.ID)
	if err != nil {
		return fmt.Errorf("failed to update doc: %w", err)
	}

	fmt.Printf("Updated todo doc with id %d", doc.ID)
	return nil
}

func (r *TodoDocsRepo) QueryAllTodoDocs() ([]*model.DocItem, error) {
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

	var ret []*model.DocItem

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

		ret = append(ret, &doc)
	}

	fmt.Printf("Found %d docs", len(ret))
	return ret, nil
}

func (r *TodoDocsRepo) QueryTodoDocById(id int) (*model.DocItem, error) {
	stmt, err := r.db.Prepare("SELECT * FROM todo_docs WHERE id=?")
	if err != nil {
		return &model.DocItem{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var doc model.DocItem
	var rawLabels string

	res, err := stmt.Query(id)
	if err != nil {
		return &model.DocItem{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	if err = res.Scan(&doc.ID, &doc.Created, &doc.Name, &doc.Comment, &doc.Priority, &rawLabels); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &model.DocItem{}, nil
		}
		return &model.DocItem{}, fmt.Errorf("failed to scan row: %w", err)
	}

	err = json.Unmarshal([]byte(rawLabels), &doc.Labels)
	if err != nil {
		return &model.DocItem{}, fmt.Errorf("failed to unmarshal labels: %w", err)
	}

	fmt.Printf("Found doc with id %d", doc.ID)
	return &doc, nil
}

func (r *TodoDocsRepo) MinusOneAbove(priority int) (naffected int64, err error) {
	stmt, err := r.db.Prepare("UPDATE todo_docs SET priority = priority - 1 WHERE priority > ?")
	if err != nil {
		return 0, fmt.Errorf("failed to update priority: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(priority)
	if err != nil {
		return 0, fmt.Errorf("failed to update priority: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get affected rows: %w", err)
	}

	fmt.Printf("Priority minus one affected rows: %d", affected)
	return affected, nil
}

func (r *TodoDocsRepo) SwitchPriority(id1 int, id2 int) error {
	stmt, err := r.db.Prepare(
		"UPDATE todo_docs SET priority = CASE " +
			"WHEN id = ? THEN (SELECT priority FROM todo_docs WHERE id = ?) " +
			"WHEN id = ? THEN (SELECT priority FROM todo_docs WHERE id = ?) " +
			"ELSE priority " +
			"END " +
			"WHERE id IN (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to switch priority: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id1, id2, id2, id1, id1, id2)
	if err != nil {
		return fmt.Errorf("failed to switch priority: %w", err)
	}

	fmt.Printf("Switched priority between id %d and %d", id1, id2)
	return nil
}

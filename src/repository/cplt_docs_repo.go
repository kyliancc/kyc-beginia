package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kyliancc/kyc-beginia/src/model"
)

type CpltDocsRepo struct {
	db *sql.DB
}

func NewCpltDocsRepo(db *sql.DB) *CpltDocsRepo {
	return &CpltDocsRepo{db: db}
}

func (r *CpltDocsRepo) CreateCpltDoc(doc *model.DocItem) (id int, err error) {
	stmt, err := r.db.Prepare("INSERT INTO cplt_docs (created, name, comment, labels) VALUES (?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare insert: %w", err)
	}
	defer stmt.Close()

	jsonLabels, err := json.Marshal(doc.Labels)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal labels: %w", err)
	}

	res, err := stmt.Exec(doc.Created, doc.Name, doc.Comment, jsonLabels)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into cplt_docs: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	fmt.Printf("Created completed doc with id %d", lastID)
	return int(lastID), nil
}

func (r *CpltDocsRepo) DeleteCpltDoc(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM cplt_docs WHERE id=?")
	if err != nil {
		return fmt.Errorf("failed to prepare delete: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to delete doc: %w", err)
	}

	fmt.Printf("Deleted completed doc with id %d", id)
	return nil
}

func (r *CpltDocsRepo) UpdateCpltDoc(doc *model.DocItem) error {
	stmt, err := r.db.Prepare("UPDATE cplt_docs SET created=?, name=?, comment=?, labels=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("failed to prepare update: %w", err)
	}
	defer stmt.Close()

	jsonLabels, err := json.Marshal(doc.Labels)
	if err != nil {
		return fmt.Errorf("failed to marshal labels: %w", err)
	}

	_, err = stmt.Exec(doc.Created, doc.Name, doc.Comment, jsonLabels, doc.ID)
	if err != nil {
		return fmt.Errorf("failed to update doc: %w", err)
	}

	fmt.Printf("Updated doc with id %d", doc.ID)
	return nil
}

func (r *CpltDocsRepo) QueryAllCpltDocs() ([]model.DocItem, error) {
	stmt, err := r.db.Prepare("SELECT * FROM cplt_docs")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	var ret []model.DocItem

	for res.Next() {
		var doc model.DocItem
		var rawLabels string
		if err = res.Scan(&doc.ID, &doc.Created, &doc.Completed, &doc.Name, &doc.Comment, &rawLabels); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		err = json.Unmarshal([]byte(rawLabels), &doc.Labels)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal labels: %w", err)
		}

		doc.Done = true
		ret = append(ret, doc)
	}

	fmt.Printf("Found %d docs", len(ret))
	return ret, nil
}

func (r *CpltDocsRepo) QueryCpltDocById(id int) (model.DocItem, error) {
	stmt, err := r.db.Prepare("SELECT * FROM cplt_docs WHERE id=?")
	if err != nil {
		return model.DocItem{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var doc model.DocItem
	var rawLabels string

	res, err := stmt.Query(id)
	if err != nil {
		return model.DocItem{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	if err = res.Scan(&doc.ID, &doc.Created, &doc.Completed, &doc.Name, &doc.Comment, &rawLabels); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.DocItem{}, nil
		}
		return model.DocItem{}, fmt.Errorf("failed to scan row: %w", err)
	}

	err = json.Unmarshal([]byte(rawLabels), &doc.Labels)
	if err != nil {
		return model.DocItem{}, fmt.Errorf("failed to unmarshal labels: %w", err)
	}

	doc.Done = true
	return doc, nil
}

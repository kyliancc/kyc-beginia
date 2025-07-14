package service

import (
	"database/sql"
	"fmt"
	"github.com/kyliancc/kyc-beginia/src/model"
	"github.com/kyliancc/kyc-beginia/src/repository"
)

type DocsService struct {
	globalRepo   *repository.GlobalRepo
	cpltDocsRepo *repository.CpltDocsRepo
	todoDocsRepo *repository.TodoDocsRepo
	maxPriority  int
}

func NewDocsService(db *sql.DB) *DocsService {
	return &DocsService{
		globalRepo:   repository.NewGlobalRepo(db),
		cpltDocsRepo: repository.NewCpltDocsRepo(db),
		todoDocsRepo: repository.NewTodoDocsRepo(db),
		maxPriority:  0,
	}
}

func (s *DocsService) CreateDoc(doc *model.DocItem) (int, error) {
	doc.Priority = s.maxPriority + 1
	s.maxPriority += 1
	id, err := s.todoDocsRepo.CreateTodoDoc(doc)
	if err != nil {
		return 0, fmt.Errorf("failed to create doc item: %w", err)
	}
	return id, nil
}

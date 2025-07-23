package service

import (
	"database/sql"
	"errors"
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
		return 0, err
	}
	return id, nil
}

func (s *DocsService) UpdateDoc(doc *model.DocItem) error {
	if doc.Done {
		// Completed doc item
		err := s.cpltDocsRepo.UpdateCpltDoc(doc)
		if err != nil {
			return err
		}
	} else {
		// _Todo doc item
		err := s.todoDocsRepo.UpdateTodoDoc(doc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DocsService) DeleteDoc(doc *model.DocItem) error {
	if doc.Done {
		err := s.cpltDocsRepo.DeleteCpltDoc(doc.ID)
		if err != nil {
			return err
		}
	} else {
		doc, err := s.todoDocsRepo.QueryTodoDocById(doc.ID)
		if err != nil {
			return err
		}

		err = s.todoDocsRepo.DeleteTodoDoc(doc.ID)
		if err != nil {
			return err
		}

		_, err = s.todoDocsRepo.MinusOneAbove(doc.Priority)
		if err != nil {
			return err
		}
		s.maxPriority -= 1
	}
	return nil
}

func (s *DocsService) GetAllDocs() (todo []*model.DocItem, cplt []*model.DocItem, err error) {
	todoDocs, todoErr := s.todoDocsRepo.QueryAllTodoDocs()
	if todoErr != nil {
		return nil, nil, todoErr
	}
	cpltDocs, cpltErr := s.cpltDocsRepo.QueryAllCpltDocs()
	if cpltErr != nil {
		return nil, nil, cpltErr
	}
	return todoDocs, cpltDocs, nil
}

func (s *DocsService) GetDoc(doc *model.DocItem) (*model.DocItem, error) {
	if doc.Done {
		doc, err := s.cpltDocsRepo.QueryCpltDocById(doc.ID)
		if err != nil {
			return nil, err
		}
		return doc, nil
	} else {
		doc, err := s.todoDocsRepo.QueryTodoDocById(doc.ID)
		if err != nil {
			return nil, err
		}
		return doc, nil
	}
}

func (s *DocsService) CompleteDoc(id int) error {
	doc, err := s.todoDocsRepo.QueryTodoDocById(id)
	if err != nil {
		return err
	}
	err = s.todoDocsRepo.DeleteTodoDoc(id)
	if err != nil {
		return err
	}
	_, err = s.todoDocsRepo.MinusOneAbove(doc.Priority)
	if err != nil {
		return err
	}
	_, err = s.cpltDocsRepo.CreateCpltDoc(doc)
	if err != nil {
		return err
	}
	s.maxPriority -= 1
	return nil
}

func (s *DocsService) SwitchTodoPriority(pairs [][]int) error {
	for _, pair := range pairs {
		if len(pair) != 2 {
			return errors.New("invalid pair")
		}
	}
	for _, pair := range pairs {
		err := s.todoDocsRepo.SwitchPriority(pair[0], pair[1])
		if err != nil {
			return err
		}
	}
	return nil
}

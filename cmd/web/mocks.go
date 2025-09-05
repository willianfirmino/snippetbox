package main

import (
	"time"

	"snippetbox.wfirmino.net/internal/models"
)

type snippetModelMock struct{}

func (m *snippetModelMock) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *snippetModelMock) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return &models.Snippet{
			ID:      1,
			Title:   "An old silent pond",
			Content: "An old silent pond...",
			Created: time.Now(),
			Expires: time.Now(),
		}, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *snippetModelMock) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{
		{
			ID:      1,
			Title:   "An old silent pond",
			Content: "An old silent pond...",
			Created: time.Now(),
			Expires: time.Now(),
		},
	}, nil
}

type userModelMock struct{}

func (m *userModelMock) Insert(name, email, password string) error {
	return nil
}

func (m *userModelMock) Authenticate(email, password string) (int, error) {
	return 1, nil
}

func (m *userModelMock) Exists(id int) (bool, error) {
	return true, nil
}

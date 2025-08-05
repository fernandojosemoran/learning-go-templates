package blog

import (
	"errors"
	"html/template"
	"strings"
)

type CreateArticleDto struct {
	Title   string
	Content template.HTML
}

func isEmpty(target string) bool {
	if len(strings.ToLower(strings.TrimSpace(target))) == 0 {
		return true
	}

	return false
}

func (dto CreateArticleDto) Create() (*CreateArticleDto, error) {
	// &template.Error{}

	if isEmpty(dto.Title) {
		return nil, errors.New("Title is required")
	}

	if isEmpty(string(dto.Content)) {
		return nil, errors.New("Content is required")
	}

	return &dto,
		nil
}

package blog

import "html/template"

type ArticleEntity struct {
	Id      string
	Title   string
	Content template.HTML
}

package utils

import (
	"beeblog/models"

	bee "github.com/beego/beego/v2/server/web"
)

type Pagination struct {
	Articles            *[]models.Article
	PageinationArticles []models.Article
	PageSize            int
	CurrrentPage        int
	TotalPages          int
}

func GetPageSize() int {
	pagesize, _ := bee.AppConfig.Int("pagesize")
	return pagesize
}

func GetPaginatedArticles(articles *[]models.Article, pageNum int, pageSize int) *[]models.Article {
	start := (pageNum - 1) * pageSize
	end := start + pageSize

	articleSlice := *articles

	if start < 0 {
		start = 0
	}

	if end > len(articleSlice) {
		end = len(articleSlice)
	}

	// get the slice
	paginationArticles := articleSlice[start:end]
	return &paginationArticles
}

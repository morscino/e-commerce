package repo

import (
	"strings"

	"e-commerce/db"
	"e-commerce/models"
)

type WhereObj struct {
	field     string
	condition string
	value     interface{}
}

const (
	DEFAULT_PAGE                = 1
	DEFAULT_LIMIT               = 10
	PAGE_DEFAULT_SORTED_BY      = "created_at"
	PAGE_DEFAULT_SORT_DIRECTION = "desc"
)

type Repo struct {
	Repo *db.Database
}

func NewRepo(db *db.Database) *Repo {
	return &Repo{
		Repo: db,
	}
}

func getPaginationInfo(query *models.APIPagingDto) (*models.APIPagingDto, int) {
	var offset int
	// load defaults
	if query.Page == 0 {
		query.Page = DEFAULT_PAGE
	}
	if query.Limit == 0 {
		query.Limit = DEFAULT_LIMIT
	}
	if query.Sort == "" {
		query.Sort = PAGE_DEFAULT_SORTED_BY
	}
	if query.Direction == "" {
		query.Direction = PAGE_DEFAULT_SORT_DIRECTION
	}

	if query.Page > 1 {
		offset = query.Limit * (query.Page - 1)
	}
	return query, offset
}

func getPagingInfo(query *models.APIPagingDto, count int) models.PagingInfo {
	var hasNextPage bool
	next := int64((query.Page * query.Limit) - count)
	if next < 0 && query.Limit > 0 {
		hasNextPage = true
	}

	pagingInfo := models.PagingInfo{
		TotalCount:  count,
		HasNextPage: hasNextPage,
		Page:        int(query.Page),
	}

	return pagingInfo
}

func genWhere(filtered WhereObj) *WhereObj {
	switch filtered.condition {
	case "eq":
		filtered.condition = "="
	case "like":
		filtered.condition = "like"
		filtered.value = "%" + filtered.value.(string) + "%"
	case "in":
		filtered.condition = "in"
		filtered.value = strings.Split(filtered.value.(string), ",")
	case "ne":
		filtered.condition = "<>"
	case "gt":
		filtered.condition = ">"
	case "lt":
		filtered.condition = "<"
	}
	return &filtered
}

func getFilterFromQuery(filterValue string) []*WhereObj {
	filtered := parseFilterEntries(filterValue)
	for i := 0; i < len(filtered); i++ {
		present := filtered[i]
		filtered[i] = genWhere(*present)
	}
	return filtered
}

func parseFilterEntries(filter string) []*WhereObj {

	if filter == "" {
		return nil
	}
	splitFilter := strings.Split(filter, " ")
	var allWhereObj []*WhereObj
	for i := 0; i < len(splitFilter); i++ {
		data := strings.Split(splitFilter[i], "|")
		var obj WhereObj
		if len(data) < 2 {
			return nil
		}
		obj.field = data[0]
		obj.condition = data[1]
		obj.value = data[2]
		allWhereObj = append(allWhereObj, &obj)
	}

	return allWhereObj
}

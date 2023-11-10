package repositories

import (
	"database/sql"
	"fmt"
	"galleries/modules/galleries/models"
	"galleries/packages/builders"
	"galleries/packages/database"
	"galleries/packages/helpers/converter"
	"galleries/packages/helpers/generator"
	"galleries/packages/helpers/response"
	"galleries/packages/utils/pagination"
	"math"
	"net/http"
)

func GetAllGalleries(page, pageSize int, path string) (response.Response, error) {
	// Declaration
	var obj models.GetGalleries
	var arrobj []models.GetGalleries
	var res response.Response
	var baseTable = "galleries"
	var sqlStatement string

	// Nullable column
	var GalleryDesc sql.NullString
	var GalleryTag sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("content_info", &baseTable, nil)
	whereTemplate := builders.GetTemplateLogic(baseTable, "active")
	where := "is_private = 0"
	order := "created_at, galleries_name DESC "

	sqlStatement = "SELECT " + selectTemplate + ", galleries_desc, galleries_tag, galleries_format " +
		"FROM " + baseTable + " " +
		"WHERE " + where + " AND " + whereTemplate + " " +
		"ORDER BY " + order +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	fmt.Println(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.GallerySlug,
			&obj.GalleryName,
			&GalleryDesc,
			&GalleryTag,
			&obj.GalleryFormat,
			&obj.CreatedAt,
			&obj.CreatedBy,
		)

		if err != nil {
			return res, err
		}

		obj.GalleryDesc = converter.CheckNullString(GalleryDesc)
		obj.GalleryTag = converter.CheckNullString(GalleryTag)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &where)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}

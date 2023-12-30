package repositories

import (
	"galleries/packages/builders"
	"galleries/packages/database"
	"galleries/packages/helpers/generator"
	"galleries/packages/helpers/response"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func HardDelGalleryBySlug(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "galleries"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "galleries_slug")

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "permanently delete", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func PostGallery(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "galleries"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	galName := data.FormValue("galleries_name")
	slug := generator.GetSlug(galName)
	galDesc := data.FormValue("galleries_desc")
	galUrl := data.FormValue("galleries_url")
	galTag := data.FormValue("galleries_tag")
	galFormat := data.FormValue("galleries_format")
	isPrivate := data.FormValue("is_private")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, galleries_slug, galleries_name, galleries_desc, galleries_url, galleries_tag, galleries_format, is_private, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?,null,null,null,null)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, slug, galName, galDesc, galUrl, galTag, galFormat, isPrivate, dt, "1")
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg(baseTable, "create", int(rowsAffected))
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

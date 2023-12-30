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

func HardDelDictionaryById(id string) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Command builder
	sqlStatement = builders.GetTemplateCommand("hard_delete", baseTable, "id")

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

func PostDictionary(data echo.Context) (response.Response, error) {
	// Declaration
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string
	dt := time.Now().Format("2006-01-02 15:04:05")

	// Data
	id := uuid.Must(uuid.NewRandom())
	dctName := data.FormValue("dictionaries_name")
	dctType := data.FormValue("dictionaries_type")

	// Command builder
	sqlStatement = "INSERT INTO " + baseTable + " (id, dictionaries_type, dictionaries_name, created_at, created_by) " +
		"VALUES (?,?,?,?,?)"

	// Exec
	con := database.CreateCon()
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, dctType, dctName, dt, "1")
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

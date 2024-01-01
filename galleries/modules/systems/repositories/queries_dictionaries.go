package repositories

import (
	"galleries/modules/systems/models"
	"galleries/packages/builders"
	"galleries/packages/database"
	"galleries/packages/helpers/generator"
	"galleries/packages/helpers/response"
	"net/http"
)

func GetDictionaryByType(path string, dctType string) (response.Response, error) {
	// Declaration
	var obj models.GetDictionaryByType
	var arrobj []models.GetDictionaryByType
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Query builder
	where := "dictionaries_type = '" + dctType + "' "
	order := "dictionaries_name DESC "

	sqlStatement = "SELECT id, dictionaries_type, dictionaries_name " +
		"FROM " + baseTable + " " +
		"WHERE " + where +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.ID,
			&obj.DctName,
			&obj.DctType,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &where)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetDictionaryByMyTag(path string) (response.Response, error) {
	// Declaration
	var obj models.GetDictionaryByMyTag
	var arrobj []models.GetDictionaryByMyTag
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string
	where := "dictionaries_type = 'gallery' AND created_by = '198bdb3d-374b-41e9-84ff-88a941d0f0ce'" // for now

	sqlStatement = "SELECT dictionaries_name " +
		"FROM " + baseTable + " " +
		"WHERE " + where + " " +
		"ORDER BY created_at DESC"

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.DctName,
		)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &where)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

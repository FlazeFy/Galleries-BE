package models

type (
	GetDictionaryByType struct {
		ID      string `json:"id"`
		DctType string `json:"dictionary_name"`
		DctName string `json:"dictionary_desc"`
	}
)

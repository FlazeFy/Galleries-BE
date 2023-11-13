package models

type (
	GetGalleries struct {
		GallerySlug   string `json:"galleries_slug"`
		GalleryName   string `json:"galleries_name"`
		GalleryDesc   string `json:"galleries_desc"`
		GalleryUrl    string `json:"galleries_url"`
		GalleryTag    string `json:"galleries_tag"`
		GalleryFormat string `json:"galleries_format"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
	}
	GetGalleryDetail struct {
		GallerySlug   string `json:"galleries_slug"`
		GalleryName   string `json:"galleries_name"`
		GalleryDesc   string `json:"galleries_desc"`
		GalleryUrl    string `json:"galleries_url"`
		GalleryTag    string `json:"galleries_tag"`
		GalleryFormat string `json:"galleries_format"`

		// Props
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	}
)

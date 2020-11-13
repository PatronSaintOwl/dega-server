package model

import "github.com/factly/dega-server/config"

// Organisation model
type Organisation struct {
	config.Base
	Title            string  `json:"title"`
	Slug             string  `json:"slug"`
	Description      string  `json:"description"`
	FeaturedMediumID *uint   `json:"featured_medium_id"`
	Medium           *Medium `json:"medium"`
}

// OrganisationPermission model
type OrganisationPermission struct {
	config.Base
	OrganisationID uint  `gorm:"column:organisation_id" json:"organisation_id"`
	Spaces         int64 `gorm:"column:spaces" json:"spaces"`
	Media          int64 `gorm:"column:media" json:"media"`
	Posts          int64 `gorm:"column:posts" json:"posts"`
	FactCheck      bool  `gorm:"fact_check" json:"fact_check"`
}

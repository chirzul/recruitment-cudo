package db

import (
	"github.com/chirzul/recruitment-cudo/src/model"
	"gorm.io/gorm"
)

func GenerateJSONStructure(orgID string, db *gorm.DB) (model.Organization, error) {
	var org model.Organization

	if err := db.Select("org_id,org_name").Where("org_id = ? AND org_status = ?", orgID, "1").First(&org).Error; err != nil {
		return org, err
	}

	err := getChildOrganization(&org, db)
	if err != nil {
		return org, err
	}

	return org, nil
}

func getChildOrganization(org *model.Organization, db *gorm.DB) error {
	var orgChilds []model.Organization

	if err := db.Select("org_id,org_name").Where("org_parent_id = ? AND org_status = ?", org.OrgID, "1").Find(&orgChilds).Error; err != nil {
		return err
	}

	for i := range orgChilds {
		err := getChildOrganization(&orgChilds[i], db)
		if err != nil {
			return err
		}
	}

	org.OrgChilds = orgChilds

	return nil
}

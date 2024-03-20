package model

type Organization struct {
	OrgID     string         `json:"org_id"`
	OrgName   string         `json:"org_name"`
	OrgChilds []Organization `json:"org_childs,omitempty" gorm:"-"`
}

func (Organization) TableName() string {
	return "organization"
}

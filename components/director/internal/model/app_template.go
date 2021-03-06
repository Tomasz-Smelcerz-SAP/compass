package model

import "github.com/kyma-incubator/compass/components/director/pkg/pagination"

type ApplicationTemplate struct {
	ID                   string
	Name                 string
	Description          *string
	ApplicationInputJSON string
	Placeholders         []ApplicationTemplatePlaceholder
	AccessLevel          ApplicationTemplateAccessLevel
}

type ApplicationTemplatePage struct {
	Data       []*ApplicationTemplate
	PageInfo   *pagination.Page
	TotalCount int
}

type ApplicationTemplateInput struct {
	Name                 string
	Description          *string
	ApplicationInputJSON string
	Placeholders         []ApplicationTemplatePlaceholder
	AccessLevel          ApplicationTemplateAccessLevel
}

type ApplicationTemplateAccessLevel string

const (
	GlobalApplicationTemplateAccessLevel ApplicationTemplateAccessLevel = "GLOBAL"
)

type ApplicationTemplatePlaceholder struct {
	Name        string
	Description *string
}

type ApplicationTemplateValueInput struct {
	Placeholder string
	Value       string
}

func (a *ApplicationTemplateInput) ToApplicationTemplate(id string) ApplicationTemplate {
	if a == nil {
		return ApplicationTemplate{}
	}

	return ApplicationTemplate{
		ID:                   id,
		Name:                 a.Name,
		Description:          a.Description,
		ApplicationInputJSON: a.ApplicationInputJSON,
		Placeholders:         a.Placeholders,
		AccessLevel:          a.AccessLevel,
	}
}

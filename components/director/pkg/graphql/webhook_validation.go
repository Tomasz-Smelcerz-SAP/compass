package graphql

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (i WebhookInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Type, validation.Required, validation.In(ApplicationWebhookTypeConfigurationChanged)),
		validation.Field(&i.URL, validation.Required, is.URL, validation.Length(1, 256)),
		validation.Field(&i.Auth),
	)
}

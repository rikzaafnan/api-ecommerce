package payment

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PaymentInput struct {
	TrxCode string `json:"trxCode"`
}

func (pi *PaymentInput) Validate() error {
	return validation.ValidateStruct(pi,
		validation.Field(pi.TrxCode, validation.Required),
	)
}

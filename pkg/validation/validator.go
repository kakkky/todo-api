package validation

import "github.com/go-playground/validator/v10"

// シングルトンインスタンスとして用意
var validation *validator.Validate

func NewValidation() *validator.Validate {
	if validation != nil {
		return validation
	}
	// 非ポインタ構造体でもrequiredgタグが有効になる
	return validator.New(validator.WithRequiredStructEnabled())
}

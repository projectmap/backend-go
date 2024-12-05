package product

import (
	"clean-architecture/pkg/api_errors"
	"net/http"
)

var (
	ErrInvalidProductID = api_errors.NewAPIError(http.StatusBadRequest, "Invalid product ID")
)

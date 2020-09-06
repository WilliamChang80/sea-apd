package product

import (
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
)

type GetProductByIdResponse struct {
	base.BaseResponse
	Data domain.ProductDto `json:"data"`
}

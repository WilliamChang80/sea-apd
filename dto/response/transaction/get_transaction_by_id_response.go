package transaction

import "github.com/williamchang80/sea-apd/dto/domain"

type GetTransactionByIdResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    domain.TransactionDto `json:"data"`
}

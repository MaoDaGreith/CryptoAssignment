package parser

import (
	"MaoDaGreith/CryptoAssignment/model"
)

type Parser interface {
	GetCurrentBlock() int64
	Subscribe(address string) bool
	GetTransactions(address string) []model.Transaction
}

package queries

import (
	"github.com/datrine/http/apis/dtos"
)

func GetOrdersQuery(data dtos.GetOrdersDTO) (*[]Order, error) {
	return &[]Order{}, nil
}

package business

import "github.com/google/uuid"

type (
	Orders []Order

	Order struct {
		ID       uuid.UUID
		Item     string
		Quantity int32
	}
)

package events

type OrderEventType uint64

const (
	CreateOrder OrderEventType = iota
	OrderCreated
	OrderUpdated
	CartCreated
	CartUpdated
	CartDeleted
	ItemAddedToCart
	ItemEditedInCart
	ItemRemovedFromCart
	ItemQuantityEdited
	ShippingAddressUpdated
	OrderProcessed
	DiscountApplied
	OrderDelivered
	OrderCancelled
)

func (s OrderEventType) String() string {
	switch s {
	case CreateOrder:
		return "CreateOrder"
	case OrderCreated:
		return "OrderCreated"
	case OrderUpdated:
		return "OrderUpdated"
	case OrderProcessed:
		return "OrderProcessed"
	case DiscountApplied:
		return "DiscountApplied"
	case ShippingAddressUpdated:
		return "ShippingAddressUpdated"
	case CartCreated:
		return "CartCreated"
	case CartDeleted:
		return "CartDeleted"

	case ItemEditedInCart:
		return "ItemEditedInCart"

	case ItemRemovedFromCart:
		return "ItemRemovedFromCart"
	}
	return "unknown"
}

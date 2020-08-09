package api

import (
	"github.com/dandelabs/ghostbuster-backend-libs/db"
)

/*type WebHookData struct {
	Data *Data `json:"data"`
}*/

type OrderData struct {
	Order        *db.Order         `json:"order"`
	OrderDetails []*db.OrderDetail `json:"order_details"`
}

/*type Data struct {
	OrderID         int          `json:"id"`
	PartiallyCancel bool         `json:"partially_cancel"`
	SyncCreated     string       `json:"sync_created"`
	OrderStatus     string       `json:"order_status"`
	IsBackOrder     bool         `json:"is_back_order"`
	ModifiedDate    string       `json:"modified_date"`
	EventName       string       `json:"event_name"`
	OrderItems      []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ItemID            int    `json:"item_id"`
	QuatityDelivered  int    `json:"quantity_delivered"`
	QuantityCancelled int    `json:"quantity_cancelled"`
	QuantityOrdered   int    `json:"quantity_ordered"`
	SKU               string `json:"sku"`
	ID                int    `json:"id"`
}*/

package main

import "time"

type orderModel struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Local             string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              uint64    `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`

	Delivery DeliveryModel `json:"delivery"`
	Payment  PaymentModel  `json:"payment"`
	Items    []ItemModel   `json:"items"`
}

type DeliveryModel struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type ItemModel struct {
	ChrtId      uint64 `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int64  `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int64  `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int64  `json:"total_price"`
	NmID        uint64 `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type PaymentModel struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       uint64 `json:"amount"`
	PaymentDt    uint64 `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint64 `json:"delivery_cost"`
	GoodsTotal   uint64 `json:"goods_total"`
	CustomFee    uint64 `json:"custom_fee"`
}

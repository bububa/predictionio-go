package predictionio

import "time"

type ResponseMessage struct {
	Message string `json:"message"`
}

type AddUserRequest struct {
	AppKey   string `json:"pio_appkey"`
	Uid      string `json:"pio_uid"`
	Latlng   string `json:"pio_latlng"`
	Inactive string `json:"pio_inactive"`
	Customs  map[string]string
}

type GetUserRequest struct {
	AppKey string `json:"pio_appkey"`
	Uid    string `json:"pio_uid"`
}

type DeleteUserRequest struct {
	AppKey string `json:"pio_appkey"`
	Uid    string `json:"pio_uid"`
}

type AddItemRequest struct {
	AppKey    string    `json:"pio_appkey"`
	Iid       string    `json:"pio_iid"`
	Types     string    `json:pio_itypes"`
	Latlng    string    `json:"pio_latlng"`
	Inactive  string    `json:"pio_inactive"`
	StartTime time.Time `json:"pio_startT"`
	EndTime   time.Time `json:"pio_endT"`
	Price     string    `json:"pio_price"`
	Profit    string    `json:"pio_profit"`
	Customs   map[string]string
}

type GetItemRequest struct {
	AppKey string `json:"pio_appkey"`
	Iid    string `json:"pio_iid"`
}

type DeleteItemRequest struct {
	AppKey string `json:"pio_appkey"`
	Iid    string `json:"pio_iid"`
}

type RecordActionRequest struct {
	AppKey string    `json:"pio_appkey"`
	Uid    string    `json:"pio_uid"`
	Iid    string    `json:"pio_iid"`
	Action string    `json:"pio_action"`
	Rate   uint      `json:"pio_rate"`
	Latlng string    `json:"pio_latlng"`
	T      time.Time `json:"pio_t"`
}

type GetRecommendItemsRequest struct {
	AppKey     string `json:"pio_appkey"`
	Engine     string `json:"engine"`
	Uid        string `json:"pio_uid"`
	N          uint   `json:"pio_n"`
	Types      string `json:pio_itypes"`
	Latlng     string `json:"pio_latlng"`
	Within     string `json:"pio_within"`
	Unit       string `json:"pio_unit"`
	Attributes string `json:"pio_attributes"`
}

type GetSimilarItemsRequest struct {
	AppKey     string `json:"pio_appkey"`
	Engine     string `json:"engine"`
	Iid        string `json:"pio_iid"`
	N          uint   `json:"pio_n"`
	Types      string `json:pio_itypes"`
	Latlng     string `json:"pio_latlng"`
	Within     string `json:"pio_within"`
	Unit       string `json:"pio_unit"`
	Attributes string `json:"pio_attributes"`
}

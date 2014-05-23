package predictionio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const timeLayout = "2006-01-02T15:04:05Z"

const (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	DELETE_METHOD = "DELETE"
)

type Client struct {
	GateWay string
}

func NewClient(gateWay string) *Client {
	return &Client{GateWay: gateWay}
}

func (this *Client) parseRequest(req interface{}) (httpRequest *http.Request, err error) {
	switch req.(type) {
	case *AddUserRequest:
		endPoint := fmt.Sprintf("%s/users.json", this.GateWay)
		request := req.(*AddUserRequest)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		values.Add("pio_uid", request.Uid)
		if request.Latlng != "" {
			values.Add("pio_latlng", request.Latlng)
		}
		if request.Inactive != "" {
			values.Add("pio_inactive", request.Inactive)
		}
		if request.Customs != nil {
			for k, v := range request.Customs {
				values.Add(k, v)
			}
		}
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		httpRequest, err = http.NewRequest(POST_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *GetUserRequest:
		request := req.(*GetUserRequest)
		endPoint := fmt.Sprintf("%s/users/%s.json", this.GateWay, request.Uid)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		httpRequest, err = http.NewRequest(GET_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *DeleteUserRequest:
		request := req.(*DeleteUserRequest)
		endPoint := fmt.Sprintf("%s/users/%s.json", this.GateWay, request.Uid)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		httpRequest, err = http.NewRequest(DELETE_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *AddItemRequest:
		endPoint := fmt.Sprintf("%s/items.json", this.GateWay)
		request := req.(*AddItemRequest)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		values.Add("pio_iid", request.Iid)
		values.Add("pio_itypes", request.Types)
		if request.Latlng != "" {
			values.Add("pio_latlng", request.Latlng)
		}
		if request.Inactive != "" {
			values.Add("pio_inactive", request.Inactive)
		}
		if request.Price != "" {
			values.Add("pio_price", request.Price)
		}
		if request.Profit != "" {
			values.Add("pio_profit", request.Profit)
		}
		if !request.StartTime.IsZero() {
			values.Add("pio_startT", request.StartTime.UTC().Format(timeLayout))
		}
		if !request.EndTime.IsZero() {
			values.Add("pio_endT", request.EndTime.UTC().Format(timeLayout))
		}
		if request.Customs != nil {
			for k, v := range request.Customs {
				values.Add(k, v)
			}
		}
		httpRequest, err = http.NewRequest(POST_METHOD, endPoint, strings.NewReader(values.Encode()))
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return
	case *GetItemRequest:
		request := req.(*GetItemRequest)
		endPoint := fmt.Sprintf("%s/items/%s.json", this.GateWay, request.Iid)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		httpRequest, err = http.NewRequest(GET_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *DeleteItemRequest:
		request := req.(*DeleteItemRequest)
		endPoint := fmt.Sprintf("%s/items/%s.json", this.GateWay, request.Iid)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		httpRequest, err = http.NewRequest(DELETE_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *RecordActionRequest:
		endPoint := fmt.Sprintf("%s/actions/u2i.json", this.GateWay)
		request := req.(*RecordActionRequest)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		values.Add("pio_uid", request.Uid)
		values.Add("pio_iid", request.Iid)
		values.Add("pio_action", request.Action)
		if request.Action == "rate" {
			values.Add("pio_rate", strconv.FormatUint(uint64(request.Rate), 10))
		}
		if request.Latlng != "" {
			values.Add("pio_latlng", request.Latlng)
		}
		if !request.T.IsZero() {
			values.Add("pio_t", request.T.UTC().Format(timeLayout))
		}
		httpRequest, err = http.NewRequest(GET_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *GetRecommendItemsRequest:
		request := req.(*GetRecommendItemsRequest)
		endPoint := fmt.Sprintf("%s/engines/itemrec/%s/topn.json", this.GateWay, request.Engine)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		values.Add("pio_uid", request.Uid)
		values.Add("pio_n", strconv.FormatUint(uint64(request.N), 10))
		if request.Latlng != "" {
			values.Add("pio_latlng", request.Latlng)
		}
		if request.Types != "" {
			values.Add("pio_itypes", request.Types)
		}
		if request.Within != "" {
			values.Add("pio_within", request.Within)
		}
		if request.Unit != "" {
			values.Add("pio_unit", request.Unit)
		}
		if request.Attributes != "" {
			values.Add("pio_attributes", request.Attributes)
		}
		httpRequest, err = http.NewRequest(GET_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	case *GetSimilarItemsRequest:
		request := req.(*GetSimilarItemsRequest)
		endPoint := fmt.Sprintf("%s/engines/itemsim/%s/topn.json", this.GateWay, request.Engine)
		values := url.Values{}
		values.Add("pio_appkey", request.AppKey)
		values.Add("pio_iid", request.Iid)
		values.Add("pio_n", strconv.FormatUint(uint64(request.N), 10))
		if request.Latlng != "" {
			values.Add("pio_latlng", request.Latlng)
		}
		if request.Types != "" {
			values.Add("pio_itypes", request.Types)
		}
		if request.Within != "" {
			values.Add("pio_within", request.Within)
		}
		if request.Unit != "" {
			values.Add("pio_unit", request.Unit)
		}
		if request.Attributes != "" {
			values.Add("pio_attributes", request.Attributes)
		}
		httpRequest, err = http.NewRequest(GET_METHOD, endPoint, strings.NewReader(values.Encode()))
		return
	}
	return nil, errors.New("unknown endpoint")
}

func (this *Client) Command(req interface{}) (res []byte, err error) {
	httpRequest, err := this.parseRequest(req)
	if err != nil {
		return nil, err
	}
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

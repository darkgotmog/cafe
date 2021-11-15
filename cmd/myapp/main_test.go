package main

import (
	// "errors"
	"bytes"
	"cafe/internal"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	menu string = `[{"name":"Вода","type":0},{"name":"Капучино","type":1},{"name":"Эспрессо","type":2},{"name":"Американо","type":3},{"name":"Ристретто","type":4}]`
)

func TestServer(t *testing.T) {

	// Setup
	go main()
	time.Sleep(time.Duration(10) * time.Millisecond)

	testMenu(t)
	testPlaceOrders(t)
	testOrderWork(t)
	time.Sleep(time.Duration(1) * time.Second)
	testOrderReady(t)
	testOrderReceve(t)

}

func testMenu(t *testing.T) {
	req := require.New(t)
	resp, err := http.Get("http://127.0.0.1:1323/api/v1/menu")

	req.NoError(err)
	req.Equal(http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	req.NoError(err)
	rezult := strings.Trim(string(body), "\n")
	req.Equal(menu, rezult)

}

func testPlaceOrders(t *testing.T) {

	list := []internal.ListPositon{}

	list = append(list, internal.ListPositon{List: []internal.Position{{Type: 1, Count: 2}, {Type: 3, Count: 1}}})
	list = append(list, internal.ListPositon{List: []internal.Position{{Type: 4, Count: 1}, {Type: 0, Count: 1}}})

	for _, value := range list {

		testPlaceOrder(t, value)
	}
}

func testPlaceOrder(t *testing.T, list internal.ListPositon) {
	req := require.New(t)

	json_data, err := json.Marshal(list)

	req.NoError(err)

	resp, err := http.Post("http://127.0.0.1:1323/api/v1/order", "application/json", bytes.NewBuffer(json_data))

	req.NoError(err)
	req.Equal(http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	errDecode := json.NewDecoder(resp.Body).Decode(&result)
	req.NoError(errDecode)

	_, ok := result["id"]

	req.Equal(true, ok)

}

func testOrderWork(t *testing.T) {
	req := require.New(t)
	resp, err := http.Get("http://127.0.0.1:1323/api/v1/orderWork")

	req.NoError(err)
	req.Equal(http.StatusOK, resp.StatusCode)

	var result []internal.Order
	errDecode := json.NewDecoder(resp.Body).Decode(&result)
	req.NoError(errDecode)
	req.Equal(2, len(result))

}

func testOrderReady(t *testing.T) {
	req := require.New(t)
	resp, err := http.Get("http://127.0.0.1:1323/api/v1/orderReady")

	req.NoError(err)
	req.Equal(http.StatusOK, resp.StatusCode)

	var result []internal.Order
	errDecode := json.NewDecoder(resp.Body).Decode(&result)
	req.NoError(errDecode)
}

func testOrderReceve(t *testing.T) {
	req := require.New(t)

	orderId := internal.OrderId{Id: 1}
	wantOrder := internal.Order{
		Id:    1,
		Ready: true,
		List:  []internal.Position{{Type: 1, Count: 2}, {Type: 3, Count: 1}},
	}

	json_data, err := json.Marshal(orderId)

	req.NoError(err)
	resp, err := http.Post("http://127.0.0.1:1323/api/v1/orderReceve", "application/json", bytes.NewBuffer(json_data))

	req.NoError(err)
	req.Equal(http.StatusOK, resp.StatusCode)

	var result internal.Order
	errDecode := json.NewDecoder(resp.Body).Decode(&result)
	req.NoError(errDecode)
	req.Equal(wantOrder, result)
}

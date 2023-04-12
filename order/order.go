package order

import (
	"bytes"
	"encoding/json"
	"log"
	"sort"
	"strconv"
)

type Order struct {
	Prices      map[int][]Price
	orderedNums []int
}

type Price struct {
	Client   string `json:"client"`
	Quantity int    `json:"quantity"`
}

// Creates a new order struct that returns all unique prices as a key with empty []price array
// e.g. order.prices[120][0]{}.
func NewOrder(records [][]string) *Order {
	prices := make(map[int][]Price, 0)
	nums := make([]int, 0)
	var i int = 0
	var exists bool = false

	for _, values := range records {
		for ii, value := range values {
			if ii == 2 {
				price, err := strconv.Atoi(value)
				if err != nil {
					log.Print(err)
					break
				}

				for _, num := range nums {
					if price == num {
						exists = true
					}
				}
				if !exists {
					nums = append(nums, price)
					prices[i] = make([]Price, 0)
					i++
				}

				exists = false
			}
		}
	}

	return &Order{Prices: prices, orderedNums: nums}
}

// Fill matches client and quantity values with it's price and
// appends to correct price key.
func (o *Order) Fill(records [][]string) {
	for _, values := range records {
		var name string
		var quantity int
		var price int
		var err error

		for i, value := range values {
			if i == 0 {
				name = value
			}

			if i == 2 {
				price, err = strconv.Atoi(value)
				if err != nil {
					break
				}
			}

			if i == 3 {
				quantity, err = strconv.Atoi(value)
				if err != nil {
					break
				}
			}
		}
		if err == nil {
			o.Prices[price] = append(o.Prices[price], Price{Client: name, Quantity: quantity})
		}
	}
}

// Custom JSON marshaller that preserves key order for type Order struct.
//
// see: https://github.com/golang/go/issues/27179 for more details.
func (om *Order) MarshalJSON() ([]byte, error) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	buffer.WriteRune('{')
	endOfJsonRange := len(om.orderedNums)

	sort.Sort(sort.Reverse(sort.IntSlice(om.orderedNums)))

	for i := range om.orderedNums {
		price := om.orderedNums[i]
		val, err := json.Marshal(strconv.Itoa(price))
		if err != nil {
			return nil, err
		}

		buffer.Write(val)
		buffer.WriteRune(':')
		buffer.WriteRune('[')

		arr := om.Prices[price]
		ii := 1
		for _, price := range arr {

			ll := len(arr)
			client, err := json.Marshal(price.Client)
			if err != nil {
				return nil, err
			}
			quantity, err := json.Marshal(price.Quantity)
			if err != nil {
				return nil, err
			}

			buffer.WriteRune('{')
			buffer.WriteString(`"client"`)
			buffer.WriteRune(':')
			buffer.Write(client)
			buffer.WriteRune(',')
			buffer.WriteString(`"quantity"`)
			buffer.WriteRune(':')
			buffer.Write(quantity)

			if ll != ii {
				buffer.WriteRune('}')
				buffer.WriteRune(',')
			} else {
				buffer.WriteRune('}')

			}
			ii++
		}

		if endOfJsonRange != 1 {
			buffer.WriteRune(']')
			buffer.WriteRune(',')

		} else {
			buffer.WriteRune(']')
		}
		endOfJsonRange--

	}
	buffer.WriteRune('}')
	return buffer.Bytes(), nil
}

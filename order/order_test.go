package order

import (
	"testing"
)

func TestFillingOrderWithRecords(t *testing.T) {
	records := [][]string{
		{"A", "Red", "80", "20"},
		{"B", "Red", "120", "20"},
		{"C", "Red", "100", "30"},
		{"D", "Red", "120", "10"},
	}

	order := NewOrder(records)

	if order.Prices == nil {
		t.Errorf("order prices should not be nil")
	}

	if len(order.Prices) != 3 {
		t.Errorf("order length should be 3, but is: %d", len(order.Prices))
	}

	order.Fill(records)

	assertPrice120(order, t)
	assertPrice100(order, t)
	assertPrice80(order, t)
}

func TestFillingOrderWithRecordsThatHaveBadData(t *testing.T) {
	records := [][]string{
		{"B", "Red", "this isn't a number", "20"},
		{"B", "Red", "120", "this also isn't a number"},
		{"D", "Red", "120", "10"},
	}

	order := NewOrder(records)

	if order.Prices == nil {
		t.Errorf("order prices should not be nil")
	}

	if len(order.Prices) != 1 {
		t.Errorf("order length should be 1, but is: %d", len(order.Prices))
	}

	order.Fill(records)

	prices, ok := order.Prices[120]
	if !ok {
		t.Errorf("price should contain value, but is %v", prices)
	}

	if order.Prices[120][0].Client != "D" {
		t.Errorf("client should be D, but is: %s", order.Prices[120][0].Client)
	}

	if order.Prices[120][0].Quantity != 10 {
		t.Errorf("quantity should be 10, but is: %d", order.Prices[120][0].Quantity)
	}
}

func TestFillingOrderWithRecordsThatHaveDuplicatePrices(t *testing.T) {
	records := [][]string{
		{"B", "Red", "120", "20"},
		{"D", "Red", "120", "10"},
	}

	order := NewOrder(records)

	if order.Prices == nil {
		t.Errorf("order prices should not be nil")
	}

	if len(order.Prices) != 1 {
		t.Errorf("order length should be 1, but is: %d", len(order.Prices))
	}

	order.Fill(records)

	prices, ok := order.Prices[120]
	if !ok {
		t.Errorf("price should contain value, but is %v", prices)
	}

	if order.Prices[120][0].Client != "B" {
		t.Errorf("client should be B, but is: %s", order.Prices[120][0].Client)
	}

	if order.Prices[120][0].Quantity != 20 {
		t.Errorf("quantity should be 20, but is: %d", order.Prices[120][0].Quantity)
	}

	if order.Prices[120][1].Client != "D" {
		t.Errorf("client should be D, but is: %s", order.Prices[120][1].Client)
	}

	if order.Prices[120][1].Quantity != 10 {
		t.Errorf("quantity should be 10, but is: %d", order.Prices[120][1].Quantity)
	}
}

func assertPrice120(order *Order, t *testing.T) {
	prices, ok := order.Prices[120]
	if !ok {
		t.Errorf("price should contain value, but is %v", prices)
	}

	if order.Prices[120][0].Client != "B" {
		t.Errorf("client should be B, but is: %s", order.Prices[120][0].Client)
	}

	if order.Prices[120][0].Quantity != 20 {
		t.Errorf("quantity should be 20, but is: %d", order.Prices[120][0].Quantity)
	}

	if order.Prices[120][1].Client != "D" {
		t.Errorf("client should be D, but is: %s", order.Prices[120][1].Client)
	}

	if order.Prices[120][1].Quantity != 10 {
		t.Errorf("quantity should be 10, but is: %d", order.Prices[120][1].Quantity)
	}
}

func assertPrice100(order *Order, t *testing.T) {
	prices, ok := order.Prices[100]
	if !ok {
		t.Errorf("price should contain value, but is %v", prices)
	}

	if order.Prices[100][0].Client != "C" {
		t.Errorf("client should be C, but is: %s", order.Prices[100][0].Client)
	}

	if order.Prices[100][0].Quantity != 30 {
		t.Errorf("quantity should be 30, but is: %d", order.Prices[100][0].Quantity)
	}
}

func assertPrice80(order *Order, t *testing.T) {
	prices, ok := order.Prices[80]
	if !ok {
		t.Errorf("price should contain value, but is %v", prices)
	}

	if order.Prices[80][0].Client != "A" {
		t.Errorf("client should be A, but is: %s", order.Prices[80][0].Client)
	}

	if order.Prices[80][0].Quantity != 20 {
		t.Errorf("quantity should be 20, but is: %d", order.Prices[80][0].Quantity)
	}
}

func TestMarshallingJSON(t *testing.T) {
	expected := "{\"120\":[{\"client\":\"B\",\"quantity\":20},{\"client\":\"D\",\"quantity\":10}],\"100\":[{\"client\":\"C\",\"quantity\":30}],\"80\":[{\"client\":\"A\",\"quantity\":20}]}"
	records := [][]string{
		{"A", "Red", "80", "20"},
		{"B", "Red", "120", "20"},
		{"C", "Red", "100", "30"},
		{"D", "Red", "120", "10"},
	}

	order := NewOrder(records)

	order.Fill(records)
	actual, err := order.MarshalJSON()

	if err != nil {
		t.Errorf("expected nil, but was: %v", err)
	}

	if expected != string(actual) {
		t.Errorf("expected %s but was: %s", expected, actual)
	}
}

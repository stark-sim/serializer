package serializer

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestRandomObjSerialize(t *testing.T) {
	// define obj
	type Book struct {
		ID           int64   `json:"id"`
		Name         string  `json:"name"`
		SerialNumber uint64  `json:"serial_number"`
		Price        float64 `json:"price"`
		PopularRate  float64 `json:"popular_rate"`
		StoreNum     int32   `json:"store_num"`
	}
	// init obj
	book := &Book{
		ID:           288989722028675072,
		Name:         "张三日记",
		SerialNumber: math.MaxUint64,
		Price:        99.9,
		PopularRate:  19999943563211111.5555555555,
		StoreNum:     10,
	}
	// serialize obj to frontend friendly
	res := FrontendSerialize(book)
	// check res
	if res.(map[string]interface{})["id"].(string) != strconv.FormatInt(book.ID, 10) {
		t.Errorf("book's id not serialized properly, got %v", res.(map[string]interface{})["id"].(string))
	}
	if res.(map[string]interface{})["serial_number"].(string) != strconv.FormatUint(book.SerialNumber, 10) {
		t.Errorf("book's serial_number not serialized properly, got %v", res.(map[string]interface{})["serial_number"].(string))
	}
	if res.(map[string]interface{})["name"].(string) != book.Name {
		t.Errorf("book's name not serialized properly, got %v", res.(map[string]interface{})["name"].(string))
	}
	if res.(map[string]interface{})["price"].(float64) != book.Price {
		t.Errorf("store_num not expected")
	}
	if res.(map[string]interface{})["store_num"].(int64) != int64(book.StoreNum) {
		t.Errorf("store_num not expected")
	}
}

func TestListObjSerialize(t *testing.T) {
	type Directory struct {
		ID       uint64      `json:"id"`
		Name     string      `json:"name"`
		Children []Directory `json:"children"`
	}
	list := []Directory{
		{
			ID:   288989722028675072,
			Name: "顶级菜单0",
			Children: []Directory{
				{
					ID:       288989722028675073,
					Name:     "二级菜单0",
					Children: nil,
				},
				{
					ID:       288989722028675074,
					Name:     "二级菜单1",
					Children: nil,
				},
			},
		},
		{
			ID:   288989722028675075,
			Name: "顶级菜单1",
			Children: []Directory{
				{
					ID:       288989722028675076,
					Name:     "二级菜单2",
					Children: nil,
				},
				{
					ID:       288989722028675077,
					Name:     "二级菜单3",
					Children: nil,
				},
			},
		},
	}
	res := FrontendSerialize(list)
	fmt.Printf("%+v", res)
}

func TestRandomValueSerialize(t *testing.T) {
	numbers := []int64{288989722028675072, 288989722028675073, 288989722028675074, 288989722028675075}
	res := FrontendSerialize(numbers)
	fmt.Printf("%+v", res)
	res = FrontendSerialize(numbers[0])
	fmt.Printf("%+v", res)
}

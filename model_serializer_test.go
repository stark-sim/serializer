package serializer

import (
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
	}
	// init obj
	book := &Book{
		ID:           288989722028675072,
		Name:         "张三日记",
		SerialNumber: math.MaxUint64,
		Price:        99.9,
		PopularRate:  99999999999.55,
	}
	// serialize obj to frontend friendly
	bookSerializer := &ModelSerializer{
		Model:    book,
		IsPlural: false,
	}
	res := bookSerializer.Serialize()
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
}

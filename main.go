package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ItemsRead(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		ids = nil
	}
	items := []map[string]interface{}{
		map[string]interface{}{
			"id":      1,
			"name":    "商品1",
			"barcode": "1234",
		},
		map[string]interface{}{
			"id":      2,
			"name":    "商品2",
			"barcode": "1235",
		},
	}
	if ids != nil {
		res := []map[string]interface{}{}
		id, _ := strconv.Atoi(ids[0])
		for _, item := range items {
			if id == item["id"].(int) {
				res = append(res, item)
			}
		}
		items = res
	}
	bytes, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(bytes)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("front")))

	http.HandleFunc("/api/items", ItemsRead)
	http.ListenAndServe(":8080", nil)
}

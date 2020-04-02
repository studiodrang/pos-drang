package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
	Price   string `json:"price"`
}
type Sale struct {
	Id        int
	SaleItems []struct {
		ItemId int
		Volume int
	}
	SaleAt time.Time
}

var DB *sql.DB

func main() {
	var err error
	DB, err = sql.Open("mysql", "drang:drang@/posdrang")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()

	// GET
	get := http.NewServeMux()
	get.Handle("/", http.FileServer(http.Dir("front")))

	// POST
	post := http.NewServeMux()
	post.HandleFunc("/api/items", ItemsRead)
	post.HandleFunc("/api/items/create", ItemsCreate)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			post.ServeHTTP(w, r)
		default:
			get.ServeHTTP(w, r)
		}
	})
	http.ListenAndServe(":8082", nil)
}

func ItemsRead(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id,name,barcode,price FROM items")
	if err != nil {
		panic(err.Error())
	}
	items := []Item{}
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Name, &item.Barcode, &item.Price)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}
	bytes, err := json.Marshal(items)
	if err != nil {
		panic(err.Error())
	}
	w.Write(bytes)
}
func ItemsCreate(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var item Item
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		panic(err.Error())
	}
	_, err = DB.Exec("INSERT INTO items(name,barcode,price)VALUES(?,?,?)", item.Name, item.Barcode, item.Price)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("{}"))
}
func ItemsUpdate(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var item Item
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		panic(err.Error())
	}
	_, err = DB.Exec("UPDATE items SET name = ?,barcode = ?,price = ? WHERE id = ?", item.Name, item.Barcode, item.Price, item.Id)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("{}"))
}
func ItemsDelete(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var item Item
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		panic(err.Error())
	}
	_, err = DB.Exec("DELETE FROM items WHERE id = ?", item.Id)
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte("{}"))
}

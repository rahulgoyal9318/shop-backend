package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/get-items", handler2)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func getDetailsById(id string) (string, error) {
	connStr := "host=127.0.0.1 user=postgres password=rahul dbname=shop_backend sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "", err
	}
	defer db.Close()
	q := fmt.Sprintf("select * from inventory_items where id='%s';", id)
	rows, err := db.Query(q)
	if err != nil {
		return "", err
	}
	if !rows.Next() {
		return "", nil

	}
	var x InventoryItems
	err = rows.Scan(&x.Id, &x.Name, &x.Qty, &x.Price)
	if err != nil {
		return "", err
	}
	y := fmt.Sprintf("name is %s price is %f qty is %d ", x.Name, x.Price, x.Qty)
	return y, nil
}

type InventoryItems struct {
	Id    string  `db:"id"`
	Name  string  `db:"name"`
	Qty   int     `db:"qty"`
	Price float64 `db:"price"`
}

func handler2(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query().Get("id")
	fmt.Println(x)
	a, b := getDetailsById(x)
	fmt.Println(b)
	w.Write([]byte(a))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := r.URL.RawQuery
	fmt.Println(x)
	w.Write([]byte(h))
}

var h string = `
<!DOCTYPE html>
<html lang="en">
<form action="/submit-form-data" method="POST">

  <!-- Label and Text Input -->
  <label for="name">Name:</label>
  <input type="text" id="name" name="user_name" placeholder="Enter your name" required>
  <br><br>

  <!-- Email Input -->
  <label for="email">Email:</label>
  <input type="email" id="email" name="user_email" required>
  <br><br>

  <!-- Message Area -->
  <label for="msg">Message:</label>
  <textarea id="msg" name="user_message"></textarea>
  <br><br>

  <!-- Submit Button -->
  <button type="submit">Submit</button>

</form>`

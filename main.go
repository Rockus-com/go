package main

import (
	"fmt"
	"net/http"
)

type User struct {
	
	name string
	age uint16
	money uint16
	avg_grades, happiness float64

}
func (u User) getAllInfo() string {
	return fmt.Sprintf("name: %s, age: %d, money: %d, avg_grades: %f, happiness: %f", u.name, u.age, u.money, u.avg_grades, u.happiness)
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 100, 4.0, 0.0}
	fmt.Fprintf(w, bob.getAllInfo())
} 

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts!")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	

	handleRequest()
}
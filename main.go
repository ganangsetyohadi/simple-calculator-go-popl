package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Menulis output ke browser
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Menghubungkan route "/" ke handler helloWorld
	http.HandleFunc("/", helloWorld)

	// Menjalankan server pada port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

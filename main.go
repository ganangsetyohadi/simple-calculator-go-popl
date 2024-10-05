package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Struct untuk menyimpan hasil kalkulasi
type CalculationResult struct {
	Result string
}

// Fungsi untuk menampilkan halaman utama
func renderTemplate(w http.ResponseWriter, result string) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, CalculationResult{Result: result})
}

// Handler untuk kalkulator
func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		num1, err1 := strconv.ParseFloat(r.FormValue("num1"), 64)
		num2, err2 := strconv.ParseFloat(r.FormValue("num2"), 64)
		operator := r.FormValue("operator")

		if err1 != nil {
			renderTemplate(w, "Invalid input for the first number")
			return
		}

		var result float64
		var calcError string

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if err2 != nil || num2 == 0 {
				calcError = "Error: Division by zero"
			} else {
				result = num1 / num2
			}
		case "^": // Operasi pangkat
			result = math.Pow(num1, num2)
		case "%": // Operasi modulo
			if err2 != nil || num2 == 0 {
				calcError = "Error: Modulo by zero"
			} else {
				result = math.Mod(num1, num2)
			}
		case "sqrt": // Akar kuadrat
			if num1 < 0 {
				calcError = "Error: Square root of a negative number"
			} else {
				result = math.Sqrt(num1)
			}
		default:
			calcError = "Error: Invalid operator"
		}

		if calcError != "" {
			renderTemplate(w, calcError)
		} else {
			renderTemplate(w, fmt.Sprintf("%.2f", result))
		}
	} else {
		renderTemplate(w, "")
	}
}

func main() {
	// Handler untuk file statis (CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handler untuk kalkulator
	http.HandleFunc("/", calculatorHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

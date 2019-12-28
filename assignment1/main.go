package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App is up!")
}

func isPrime(origNum int) bool {
	num := math.Abs(float64(origNum))
	result := true
	if num < 2 || num == 4 {
		result = false
	} else if num > 4 {
		limit := int(math.Sqrt(float64(num)))
		// fmt.Println("limit", limit)
		for i := 2; i <= limit; i++ {
			if int(num)%i == 0 {
				result = false
				break
			}
		}
	}

	fmt.Printf("isPrime check for %v returned %v\n", origNum, result)
	return result

}

func isTwoSidedPrime(origNum int) bool {
	fmt.Println("Checking isTwoSidedPrime for number: ", origNum)
	num := int(math.Abs(float64(origNum)))
	result := isPrime(num)
	if !result || num < 10 {
		return result
	}

	for num1 := num / 10; num1 > 0; num1 = num1 / 10 {
		check1 := isPrime(num1)
		if !check1 {
			result = false
			return result
		}
	}

	for num2 := num % int(math.Pow10(int(math.Log10(float64(num))))); num2 > 0; num2 = num2 % int(math.Pow10(int(math.Log10(float64(num2))))) {
		check2 := isPrime(num2)
		if !check2 {
			result = false
		}
	}

	return result
}

func isTwoSidedPrimeHandler(w http.ResponseWriter, r *http.Request) {
	strNum := mux.Vars(r)["num"]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		log.Fatal("Failed to convert string %s to integer", strNum)
	}
	result := isTwoSidedPrime(num)
	fmt.Fprintf(w, strconv.FormatBool(result))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/istwosidedprime/{num}", isTwoSidedPrimeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"net/http"
	"strconv"
	"strings"
)

const path = "/primes/"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handlePrimeRequest)
	http.ListenAndServe(":8080", mux)
}

func handlePrimeRequest(w http.ResponseWriter, r *http.Request) {
	max, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, path))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	primeCount := 0
	for i := 1; i < max; i++ {
		if checkPrime(i) {
			primeCount++
		}
	}
	w.WriteHeader(200)
	_, _ = w.Write([]byte(strconv.Itoa(primeCount)))
}

func checkPrime(toCheck int) bool {
	if toCheck == 1 {
		return false
	}
	if toCheck == 2 {
		return true
	}
	for i := 2; i < toCheck; i++ {
		if toCheck%i == 0 {
			return false
		}
	}
	return true
}

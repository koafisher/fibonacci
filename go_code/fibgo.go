package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
)

//our response to the get request for however many digits requested
type response1 struct {
    Digits   int
    Sequence []int
}

//Fibonacci will handle the request for fibonacci numbers
func Fibonacci(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    f := fib()
    n, err := strconv.Atoi(ps.ByName("name"))
    if err != nil {
        panic(err)
    }
    var fibs []int

    //if no fibonacci numbers to calculate return
    if n <= 0 {
        return
    }

    //0 will count as first number
    fibs = append(fibs, 0)

    //calculate fibonacci numbers and throw an error if there is integer overflow
    for i := 0; i < n-1; i++ {
        x, err := f()
        if err != nil {
            panic(err)
        }
        fibs = append(fibs, x)
    }

    //create the response
    res1D := &response1{
            Digits:   n,
            Sequence: fibs}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    //encode our response in json and write it to our response
    enc := json.NewEncoder(w)
    enc.Encode(res1D)

    //w.Header().Set("Content-Type", "application/json")
    //json.NewEncoder(w).Encode(fibs)
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() (int,error) {
	a, b := 0, 1
	return func() (int,error) {
		a, b = b, a+b
        if b < a  {
            return 0, fmt.Errorf("Integer overflow")
        }
		return a, nil
	}
}

func main() {
    router := httprouter.New()
    router.GET("/fibonacci/:name", Fibonacci)

    log.Fatal(http.ListenAndServe(":8080", router))
}




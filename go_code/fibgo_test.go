package main

import (
    "net/http"
    "testing"
)

type testpair struct {
    value int
    number int
}


var tests = []testpair{
{0,1},
{ 5, 3},
{ 13, 144 },
{ 92, 4660046610375530309 },
}

func TestFib(t *testing.T) {
    for _, pair := range tests {
        f := fib()

        for i := pair.value; i > 2; i-- {
            f()
        }

        x, err := f()

        if err != nil {
            panic(err)
        }

        if x != pair.number {
            t.Error(
            "For", pair.value,
            "expected", pair.number,
            "got", x,
            )
        }
    }
}

func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()

    resp, err := http.Get("http://localhost:8080/fibonacci/93")
    if err != nil {
        panic(err)
    }
    resp.Body.Close()

}

package main

import "testing"

func Test2is2(t *testing.T) {

	if fizzbuzz(2) != "2" {
		t.Error("fizzbuzz of 2 should be 2 and not ", fizzbuzz(2))
	}
}

func Test3isFizz(t *testing.T) {
	if fizzbuzz(3) != "FIZZ" {
		t.Error("fizzbuzz of 3 should be FIZZ, not", fizzbuzz(3))
	}
}

func Test5isBuzz(t *testing.T) {
	if fizzbuzz(5) != "BUZZ" {
		t.Error("fizzbuzz of 5 should be BUZZ, not", fizzbuzz(5))
	}
}

func Test15isBuzz(t *testing.T) {
	if fizzbuzz(15) != "FIZZBUZZ" {
		t.Error("fizzbuzz of 15 should be FIZZBUZZ, not", fizzbuzz(15))
	}
}

package errclass_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/makiuchi-d/errclass"
)

func ExampleErrClass() {
	// create an ErrClass
	var MyErr = errclass.New("My Error")

	// add classification information to the error
	err := MyErr(errors.New("some error"))

	// check the error is a MyErr
	if errors.Is(err, MyErr) {
		fmt.Println("err is a MyErr")
	}

	fmt.Printf("err: %v", err)

	// output:
	// err is a MyErr
	// err: some error
}

var (
	MyErr1 = errclass.New("My Error 1")
	MyErr2 = errclass.New("My Error 2")
)

func TestWrapped(t *testing.T) {
	base := errors.New("base error")

	w1 := fmt.Errorf("w1: %w", MyErr1(base))
	if !errors.Is(w1, MyErr1) {
		t.Fatal("w1 must be a MyErr1")
	}
	if errors.Is(w1, MyErr2) {
		t.Fatal("w1 must not be a MyErr1")
	}

	w12 := fmt.Errorf("w12: %w", MyErr2(w1))
	if !errors.Is(w12, MyErr1) {
		t.Fatal("w12 must be a MyErr1")
	}
	if !errors.Is(w12, MyErr2) {
		t.Fatal("w12 must be a MyErr2")
	}
}

var (
	samename  = "same name"
	SameName1 = errclass.New(samename)
	SameName2 = errclass.New(samename)
)

func TestSameName(t *testing.T) {
	err := SameName1(errors.New("base error"))
	t.Logf("s1:%v, s2:%v", errors.Is(err, SameName1), errors.Is(err, SameName2))

	if errors.Is(err, SameName2) {
		t.Fatal("err must not be a SameName2")
	}
}

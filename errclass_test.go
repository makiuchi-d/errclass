package errclass_test

import (
	"errors"
	"fmt"

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

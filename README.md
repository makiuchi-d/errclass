errclass is an error classification utility
===========================================

Usage:

```go
// create an ErrClass
var MyErr = errclass.New("My Error")

// add classification information to the error
err := MyErr(errors.New("some error"))

// check the error is a MyErr
if errors.Is(err, MyErr) {
	fmt.Println("err is a MyErr")
}

// output:
// err is a MyErr
```

package main

func main() {
	err := testError("x")
	if err != nil {
		print("err is not nil")
	} else {
		print("err is nil")
	}
}

func testError(string) (err error) {
	return
}

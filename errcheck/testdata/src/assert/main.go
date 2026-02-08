package assert

func main() {
	var i any
	_ = i.(string) // want "unchecked error"

	handleInterface(i.(string)) // want "unchecked error"

	if i.(string) == "hello" { // want "unchecked error"
		//
	}

	switch i.(type) {
	case string:
	case int:
		_ = i.(int) // want "unchecked error"
	case nil:
	}
}

func handleInterface(i any) string {
	return i.(string) // want "unchecked error"
}

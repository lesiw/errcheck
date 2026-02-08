package main

import (
	"os"

	"labs.lesiw.io/ops/goapp"

	"lesiw.io/ops"
)

type Ops struct{ goapp.Ops }

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "check")
	}
	ops.Handle(Ops{})
}

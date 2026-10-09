package must

import (
	"fmt"
	"os"
)

func Run(err error) {
	if err != nil {
		panic(err)
	}
}

func Return[T any](in T, err error) T {
	if err != nil {
		panic(err)
	}
	return in
}

func Return2[T1 any, T2 any](in1 T1, in2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return in1, in2
}

func RunOr(err error) func(format string, variables ...any) {
	return func(format string, variables ...any) {
		if err != nil {
			fmt.Fprintf(os.Stderr, format, variables...)
			panic(err)
		}
	}
}

func ReturnOr[T any](in T, err error) func(format string, variables ...any) T {
	return func(format string, variables ...any) T {
		if err != nil {
			fmt.Fprintf(os.Stderr, format, variables...)
			panic(err)
		}
		return in
	}
}

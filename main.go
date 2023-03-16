package main

func If[T any](c bool, v1, v2 T) T {
	if c {
		return v1
	} else {
		return v2
	}
}

package main

type set struct {
	m map[int]struct{}
}

func newSet(i int) *set {
	m := map[int]struct{}{i: struct{}{}}
	return &set{m}
}

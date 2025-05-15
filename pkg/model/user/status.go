package user

//go:generate stringer -type=Status

type Status int

const (
	none Status = iota
	active
	inActive
	lock
	banned
	pending
	suspended
	deleted
)

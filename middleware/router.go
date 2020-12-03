package middleware

type Level int

const (
	INFO Level = iota
	WARNING
	ERROR
	FATAL
)

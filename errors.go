package apiHelpers

type ClientErr interface {
	ClientReadable() bool
	ClientMessage() string
	ClientStatusCode() int
}

func IsClientReadable(err error) bool {
	ce, ok := err.(ClientErr)
	return ok && ce.ClientReadable()
}

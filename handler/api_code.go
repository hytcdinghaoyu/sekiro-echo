package handler

const (
	LoginError = 10001 // RFC 7231, 6.2.1
)

var statusText = map[int]string{
	LoginError: "wrong email or password",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}

package handlers

import "fmt"

type Handler interface {
	Register()
}

func BuildPathWithV1(path string) string {
	return fmt.Sprintf("/v1/%s", path)
}

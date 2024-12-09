package ulid

import "github.com/oklog/ulid/v2"

func NewUlid() string {
	return ulid.Make().String()
}

func IsValid(id string) bool {
	_, err := ulid.Parse(id)
	return err == nil
}

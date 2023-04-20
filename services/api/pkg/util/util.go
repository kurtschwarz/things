package util

import "github.com/google/uuid"

func StringPtr(s string) *string {
	return &s
}

func StringsToUUIDs(s []string) []uuid.UUID {
	c := make([]uuid.UUID, len(s))
	for i, v := range s {
		c[i] = uuid.MustParse(v)
	}

	return c
}

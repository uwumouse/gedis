package main

type Save struct {
	key   string
	value string
}

var storage []Save = []Save{}

func get(key string) string {
	var found string

	for _, s := range storage {
		if s.key == key {
			found = s.value
		}
	}

	return found
}

func set(key string, value string) {
	storage = append(storage, Save{key, value})
}

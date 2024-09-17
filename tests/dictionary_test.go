package tests

import (
	"reflect"
	"testing"

	"github.com/mbilaljawwad/golang-channels-learning/internal/dictionary"
)

func TestDictionary(t *testing.T) {
	t.Run("test for retrieving Bilal's data", func(t *testing.T) {
		dict, err := dictionary.GetUsersDictionary()
		if err != nil {
			t.Error(err)
		}

		user := dict["Bilal"]
		expected := dictionary.User{
			Name:  "Bilal",
			Email: "m.bilal.jawwad@gmail.com",
			Age:   36,
		}

		if ok := reflect.DeepEqual(user, expected); !ok {
			t.Errorf("Result: %v, Expected: %v", user, expected)
		}
	})
}

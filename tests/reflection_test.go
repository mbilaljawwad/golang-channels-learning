package tests

import (
	"reflect"
	"testing"

	"github.com/mbilaljawwad/golang-channels-learning/internal/reflection"
)

func TestWalk(t *testing.T) {
	t.Run("Scenario 1", func(t *testing.T) {
		expected := "Chris"

		var got []string

		x := struct {
			Name string
		}{expected}

		reflection.Walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}
	})

	t.Run("Scenario 2", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				reflection.Walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}

			})
		}
	})
}

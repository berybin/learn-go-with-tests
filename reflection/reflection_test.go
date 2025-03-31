package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type profile struct {
		Age  int
		City string
	}
	type person struct {
		Name    string
		Profile profile
	}

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Jay"},
			[]string{"Jay"},
		},

		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Jay", "Sydney"},
			[]string{"Jay", "Sydney"},
		},

		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Jay", 24},
			[]string{"Jay"},
		},

		{
			"nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Jay", struct {
				Age  int
				City string
			}{24, "Sydney"}},
			[]string{"Jay", "Sydney"},
		},

		{
			"pointer to things",
			&struct {
				Name    string
				Profile struct {
					City string
				}
			}{"Jay", struct{ City string }{"Sydney"}},
			[]string{"Jay", "Sydney"},
		},

		{
			"slices",
			[]profile{
				{24, "Sydney"},
				{21, "Berlin"},
			},
			[]string{"Sydney", "Berlin"},
		},

		{
			"arrays",
			[2]profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan profile)

		go func() {
			ch <- profile{12, "Berlin"}
			ch <- profile{43, "London"}
			close(ch)
		}()

		var got []string
		want := []string{"Berlin", "London"}

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with a function", func(t *testing.T) {
		fn := func() (profile, profile) {
			return profile{12, "Berlin"}, profile{24, "Glasgow"}
		}

		got := []string{}
		want := []string{"Berlin", "Glasgow"}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

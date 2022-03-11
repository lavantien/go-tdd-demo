package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	testCases := []struct {
		name string
		in   struct {
			name     string
			language string
		}
		want string
	}{
		{
			name: "OK in English",
			in: struct {
				name     string
				language string
			}{
				name:     "John",
				language: "English",
			},
			want: "Hello, John",
		},
		{
			name: "OK in empty",
			in: struct {
				name     string
				language string
			}{
				name:     "John",
				language: "",
			},
			want: "Hello, John",
		},
		{
			name: "OK in Spanish",
			in: struct {
				name     string
				language string
			}{
				name:     "Raul",
				language: "Spanish",
			},
			want: "Hola, Raul",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Hello(tc.in.name, tc.in.language)
			want := tc.want

			assertCorrectMessage(t, got, want)
		})
	}
}

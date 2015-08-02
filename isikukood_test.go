package isikukood

import (
	"testing"
	"time"
)

func TestBirthPlace(t *testing.T) {
	cases := []struct {
		in   Isikukood
		want string
	}{
		{Isikukood{code: "49403136526"}, "Lõuna-Eesti Haigla (Võru), Põlva Haigla"},
	}
	for _, c := range cases {
		got := c.in.BirthPlace()

		if got != c.want {
			t.Errorf("TestBirthplace(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestSex(t *testing.T) {
	cases := []struct {
		in   Isikukood
		want string
	}{
		{Isikukood{code: "17605030299"}, "M"},
		{Isikukood{code: "27605030299"}, "F"},
		{Isikukood{code: "37605030299"}, "M"},
		{Isikukood{code: "48603100244"}, "F"},
		{Isikukood{code: "58603100244"}, "M"},
		{Isikukood{code: "68603100244"}, "F"},
		{Isikukood{code: "78603100244"}, "M"},
		{Isikukood{code: "88603100244"}, "F"},
	}
	for _, c := range cases {
		got := c.in.Sex()

		if got != c.want {
			t.Errorf("TestSex(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestBirthday(t *testing.T) {
	cases := []struct {
		in   Isikukood
		want time.Time
	}{
		{Isikukood{code: "48603100244"}, time.Date(1986, 03, 10, 0, 0, 0, 0, time.UTC)},
		{Isikukood{code: "34501234215"}, time.Date(1945, time.January, 23, 0, 0, 0, 0, time.UTC)},
	}
	for _, c := range cases {
		got := c.in.birthday()

		if got != c.want {
			t.Errorf("TestBirthday(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestValid(t *testing.T) {
	cases := []struct {
		in   Isikukood
		want bool
	}{
		{Isikukood{code: "37605030299"}, true},
		{Isikukood{code: "48603100244"}, true},
		{Isikukood{code: "asdfasdfasa"}, false},
		{Isikukood{code: "0012240290"}, false},
	}
	for _, c := range cases {
		got := c.in.Valid()

		if got != c.want {
			t.Errorf("TestValid(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestChecksum(t *testing.T) {
	cases := []struct {
		in   Isikukood
		want int
	}{
		{Isikukood{code: "37605030299"}, 9},
		{Isikukood{code: "48603100244"}, 4},
		{Isikukood{code: "34501234215"}, 5},
		{Isikukood{code: "60012240290"}, 0},
	}
	for _, c := range cases {
		got := Checksum(c.in)

		if got != c.want {
			t.Errorf("ValidateChecksum(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

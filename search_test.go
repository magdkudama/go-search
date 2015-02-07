package search

import "testing"

var intProvider = []struct {
	data     int
	found    bool
	position int
}{
	{1, true, 0},
	{7, true, 2},
	{18, false, -1},
	{-1, true, 1},
}

func TestWithInt(t *testing.T) {
	elems := []int{1, -1, 7}

	for i, elem := range intProvider {
		found, pos, _ := Search(elems, elem.data)

		if (elem.found && !found) || (!elem.found && found) {
			t.Errorf("TestWithInt, dataset %d. Expected found %q, result %q", i, elem.found, found)
		}

		if elem.found && elem.position != pos {
			t.Errorf("TestWithInt, dataset %d. Expected position %q, result %q", i, elem.position, pos)
		}
	}
}

var stringProvider = []struct {
	data     string
	found    bool
	position int
}{
	{"a", true, 0},
	{"b", true, 2},
	{"c", false, -1},
	{"d", true, 1},
}

func TestWithString(t *testing.T) {
	elems := []string{"a", "d", "b"}

	for i, elem := range stringProvider {
		found, pos, _ := Search(elems, elem.data)

		if (elem.found && !found) || (!elem.found && found) {
			t.Errorf("TestWithString, dataset %d. Expected found %q, result %q", i, elem.found, found)
		}

		if elem.found && elem.position != pos {
			t.Errorf("TestWithString, dataset %d. Expected position %q, result %q", i, elem.position, pos)
		}
	}
}

type test struct {
	a string
}

var structProvider = []struct {
	data     test
	found    bool
	position int
}{
	{test{a: "a"}, true, 0},
	{test{a: "b"}, true, 2},
	{test{a: "bb"}, false, -1},
	{test{a: "d"}, true, 1},
}

func TestWithStruct(t *testing.T) {
	elems := []test{test{a: "a"}, test{a: "d"}, test{a: "b"}}

	for i, elem := range structProvider {
		found, pos, _ := Search(elems, elem.data)

		if (elem.found && !found) || (!elem.found && found) {
			t.Errorf("TestWithStruct, dataset %d. Expected found %q, result %q", i, elem.found, found)
		}

		if elem.found && elem.position != pos {
			t.Errorf("TestWithStruct, dataset %d. Expected position %q, result %q", i, elem.position, pos)
		}
	}
}

func TestDifferentTypes(t *testing.T) {
	elems := []int{1, 2, 3}
	_, _, err := Search(elems, "hello")

	if err == nil {
		t.Errorf("TestDifferentTypes, expected error, found no error")
	}
}

func TestInterface(t *testing.T) {
	elems := []interface{}{1, 2, 3}
	_, _, err := Search(elems, 1)

	if err == nil {
		t.Errorf("TestInterface, expected error, found no error")
	}
}

func TestInputNotSlice(t *testing.T) {
	elems := 1
	_, _, err := Search(elems, 1)

	if err == nil {
		t.Errorf("TestInputNotSlice, expected error, found no error")
	}
}

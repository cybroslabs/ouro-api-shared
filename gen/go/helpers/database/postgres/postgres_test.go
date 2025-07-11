package postgres

import (
	"testing"
)

func Test_1_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def"
	expected := []string{"col->'abc'->'def'", "", ""}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_2_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def"
	expected := []string{"col->'abc'->>'def'", "", ""}
	if a, b, c, err := dbPathToDbSelector(in, true); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_3_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def[*]@.p1"
	expected := []string{"col", "$.abc.def[*]", "@.p1"}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_4_dbPathToDbSelector(t *testing.T) {
	in := "\"table.col\":$.abc.def@.p1"
	expected := []string{"\"table.col\"", "$.abc.def", "@.p1"}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

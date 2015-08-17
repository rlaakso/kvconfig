package kvconfig

import "testing"

func TestReadConfig(t *testing.T) {
	cases := []struct {
		key, value string
	}{
		{"foo", "bar"},
		{"baz", "bad"},
		{"section/first", "aaa"},
		{"section/second", "bbb"},
	}

	out, err := ReadConfig("testfile.ini")
	if err != nil {
		t.Errorf("ReadConfig error", err)
	}
	for _, c := range cases {
		if out[c.key] != c.value {
			t.Errorf("key != value", c.key, c.value)
		}
	}
}

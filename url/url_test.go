package url

import "testing"

func TestParse(t *testing.T) {
	const rawurl = "foo.com"
	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err=%q, want nil", rawurl, err)
	}

	want := "https"
	got := u.Scheme
	if got != want {
		t.Logf("Parse(%q).Scheme = %q;want %q", rawurl, got, want)
		t.Fail()
	}

}

package url

import "testing"

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"
	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err=%q, want nil", rawurl, err)
	}

	want := "https"
	got := u.Scheme
	if got != want {
		t.Errorf("Parse(%q).Scheme = %q;want %q", rawurl, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q;want %q", rawurl, got, want)
	}

	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q;want %q", rawurl, got, want)
	}

}

func testPort(t *testing.T, in, wantPort string) {
	u := &URL{Host: in}
	if got := u.Port(); got != wantPort {
		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
	}
}

func TestURLPortWithEmptyPort(t *testing.T) { testPort(t, "foo.com:", "") }
func TestURLPortWithoutPort(t *testing.T)   { testPort(t, "foo.com", "") }
func TestURLPortIPWithPort(t *testing.T)    { testPort(t, "1.2.3.4:90", "90") }
func TestURLPortIPWithoutPort(t *testing.T) { testPort(t, "1.2.3.4", "") }

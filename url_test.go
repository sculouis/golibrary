package url

import "testing"

func TestParse(t *testing.T) {
	const rawurl = "Broken url"
	if err := Parse(rawurl); err != nil {
		t.Logf("Parse(%q) err=%q, want nil", rawurl, err)
		t.Fail()
	}

}

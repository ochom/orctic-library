package utils

import "testing"

func TestGenerateOTP(t *testing.T) {
	got := GenerateOTP(6)
	if len(got) != 6 {
		t.Errorf("GenerateOTP(6) = %v; want 6", len(got))
	}

	got2 := GenerateOTP(6)
	if len(got2) != 6 {
		t.Errorf("GenerateOTP(6) = %v; want 6", len(got2))
	}

	if got == got2 {
		t.Errorf("GenerateOTP(6) = %v; want different", got)
	}

	if len(got2) != len(got) {
		t.Errorf("GenerateOTP(6) = %v; want 6", len(got2))
	}

	got3 := GenerateOTP(0)
	if len(got3) != 0 {
		t.Errorf("GenerateOTP(0) = %v; want 0", len(got3))
	}

}

package restos

import "testing" 

func TestWindows(t *testing.T) {
	v := SetWindows()
	if v != nil {
		t.Error("Something wrong")
	}
}
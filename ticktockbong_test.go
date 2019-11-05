package otherpkg

import (
	"project1/pkg"
	"testing"
)

func Test_tocks(t *testing.T) {
	pkg.TickTockBong()
	for i, v := range pkg.Dataarray {
		if i%60 != 0 && v.Message == "Tock"{
			t.Error("Test Failed!")
		}
	}
}

func Test_bongs(t *testing.T) {
	for i, v := range pkg.Dataarray {
		if i%(60 * 60) != 0 && v.Message == "Bong"{
			t.Error("Test Failed!")
		}
	}
}

func Test_Ticks(t *testing.T) {
	for i, v := range pkg.Dataarray {
		if (i%(60) == 0 || i%(60 * 60) == 0) && v.Message == "Tick"{
			t.Error("Test Failed!")
		}
	}
}
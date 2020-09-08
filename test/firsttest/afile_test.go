package main
import (
	"testing"
	"github.com/truschival/gosteps/pkg/firstpkg"
)

func TestBanner(t * testing.T){
	r := firstpkg.Banner("Foo")
	if r != "Hi Foo" {
		t.Errorf("Banner is wrong! expected \"Hi Foo\" got %s ",r)
	}
}

func TestBannerRepeat_ok(t * testing.T){
	var name string = "Foo"
	r, err := firstpkg.BannerRepeat(2, &name)
	if !err || r != "Hi Foo Foo!" {
		t.Errorf("Banner is wrong! expected \"Hi Foo Foo!\" got %s ",r)
	}
}

func TestBannerRepeat_nok(t * testing.T){
	var name string = "Foo"
	_, err := firstpkg.BannerRepeat(-1, &name)
	if err {
		t.Error("expected test to fail!")
	}
}

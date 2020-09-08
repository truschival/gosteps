package firstpkg

func Banner(name string) string {
	var ret string = "Hi "
	return ret + name
}

func BannerRepeat(times int, name *string) (string, bool) {
	if times <= 0 {
		return "", false
	}
	var ret string = "Hi "
	for i := 0; i < times; i++ {
		ret = ret + *name
		if i < times-1 {
			ret = ret + " "
		} else {
			ret = ret + "!"
		}
	}
	return ret, true
}

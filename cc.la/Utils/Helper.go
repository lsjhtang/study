package Utils

func Join(strs ...string) string {
	str := ""
	for _,v := range strs{
		str += v
	}
	return str
}

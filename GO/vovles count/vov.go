package main

// "String"

func main() {
	anystring := "awnfsiopanf"
	vovcount := GetCount(anystring)
	println(vovcount)
}

func GetCount(str string) (count int) {
	// Enter solution here
	// var thething string
	count = 0
	for i, thething := range str {
		if thething == 'a' || thething == 'e' || thething == 'i' || thething == 'o' || thething == 'u' {
			// println(thething)
			// println(i)
			i = i + 1
			count = count + 1
		}
	}
	println("in func")
	return count
}

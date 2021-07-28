package reverse

func reverse(p *[6]int) {
	l := len(*p)
	for i := 0; i < l/2; i++ {
		p[i], p[l-1-i] =  p[l-1-i], p[i]
	}
}

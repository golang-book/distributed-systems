package lvl1

func Run(print func(<-chan string)) {
	in := make(chan string)

	go print(in)

	for _, str := range []string{"cat", "dog", "bird"} {
		in <- str
	}
}

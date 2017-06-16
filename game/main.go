package main

func main() {

}

func simulate() {
	for i := 0; i < 1000; i++ {
		block := Block{Color: ColorBlack}
	}
}

func Forward(
	in <-chan Block,
	out chan<- Block,
) {
	for block := range in {
		out <- block
	}
}

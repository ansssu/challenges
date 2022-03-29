package main

type X string

func main() {

}

func (x X) String() string {
	return Sprintf("<%s>", x)
}

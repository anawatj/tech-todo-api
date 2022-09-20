package main

func main() {
	a := App{}
	r := a.SetupRouter()
	r.Run()
}

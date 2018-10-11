package main

func main() {
	a := App{}
	a.InitializeConection("root", "123456", "vagas")

	a.Run(":3000")
}

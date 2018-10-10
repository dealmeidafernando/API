package main

func main() {
	a := App{}
	a.InitializeConection("root", "", "vagas")

	a.Run(":8080")
}

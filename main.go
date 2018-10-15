package main

func main() {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	a := App{}
	a.InitializeConection(config.AppUser, config.AppPassword, config.AppDB)

	a.Run(config.AppPort)
}

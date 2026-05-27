package main

func main() {
	err := config.Init()
	if err != nil{
		fmt.Printf("Erro ao inicializar config: %s", err.Error())
		return
	}

	db := config.GetDb()
}

package main

import (
	utils "SSM/main/application/utils"
)

const port = ":8054"

func main() {
	router := utils.InitRouter()

	utils.StartServer(port, router)
}

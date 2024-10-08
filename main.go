package main

import (
	"github.com/Zmohamed6991/JWT-Project/config"
	"github.com/Zmohamed6991/JWT-Project/route"
)


func main(){
	
	config.ConnectingDB()
	route.Router()

	
}
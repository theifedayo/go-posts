package models

import "github.com/theifedayo/go-posts/config"


 func init(){
	config.LoadEnvVariables()
	config.ConnectToDB()
 }

 func main(){
	config.DB.AutoMigrate(&Post{})
 }
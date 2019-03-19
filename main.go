package main

import (
	"github.com/MalikovSoft/coverted_ctx_links_validator/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := database.InitDatabase(`root@/ncfu?charset=utf8&parseTime=true&loc=Local`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

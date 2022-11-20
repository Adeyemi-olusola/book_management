package config

import(
	"github.com/jinzhu/gorm"
	
)

var(
	db *gorm.DB
)

func Connect(){
	d,err := gorm.Open("mysql" , "sola:0123456789@simple?charset=utf8&parseTime=True&loc=Local")
	if( err != nil){
		panic(err)
	}else{
		db = d

	}
}

func GetDB()*gorm.DB{
	return db
}
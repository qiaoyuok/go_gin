package main

import (
	"go_gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery, // generate mode
	})

	db, _ := gorm.Open(mysql.Open(config.GetDbDsn()))
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(g.GenerateModelAs("user", "User"))

	// Generate the code
	g.Execute()
}

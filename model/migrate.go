package model

func migration() {

	// 开发时 schema 会调整，自动迁移会让数据库 schema 一直保持最新
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})

	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}

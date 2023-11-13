package config

import "fmt"

const (
	DBHost     = "localhost"
	DBPort     = "3306"
	DBUser     = "root"
	DBPassword = "bkashWEB786"
	DBName     = "test-db"
	DBType     = "mysql"
)

func GetDBType() string {
	return DBType
}

func GetMYSQLConnectionString() string {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)
	return database
}

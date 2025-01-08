package db

import (
	"fmt"
	"new-beego-api/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func Init() {
	// Register the driver
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// Get database configuration from app.conf
	user, _ := beego.AppConfig.String("db_user")
	pass, _ := beego.AppConfig.String("db_pass")
	host, _ := beego.AppConfig.String("db_host")
	port, _ := beego.AppConfig.String("db_port")
	dbname, _ := beego.AppConfig.String("db_name")

	// Create connection string
	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname)

	// Register the database
	err := orm.RegisterDataBase("default", "postgres", dataSource)
	if err != nil {
		logs.Error("Failed to register database:", err)
		return
	}

	// Register models
	orm.RegisterModel(new(models.User))

	// Create tables
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("Failed to sync database:", err)
		return
	}

	logs.Info("Database connected successfully")
}

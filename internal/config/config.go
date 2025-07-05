package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
	Sslmode  string
}

func NewConfigDb() *config {
	return &config{}
}

func (c *config) GetConfigDbFromYaml(folderName, fileName string) error {
	if err := readYaml(fileName, folderName); err != nil {
		return err
	}

	c.User = viper.GetString("db.user")
	c.Password = os.Getenv("POSTGRES_PASSWORD")
	c.Host = viper.GetString("db.host")
	c.Port = os.Getenv("PORT_DB")
	c.Dbname = viper.GetString("db.dbname")
	c.Sslmode = viper.GetString("db.sslmode")

	return nil
}

func (c *config) GetConnectSting() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Dbname, c.Sslmode)
}

func readYaml(fileName, folderName string) error {
	viper.AddConfigPath(folderName)
	viper.SetConfigName(fileName)
	return viper.ReadInConfig()
}

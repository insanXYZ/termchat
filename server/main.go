package main

import (
	"backend/bootstrap"
	"backend/config"
)

func main() {
	viper := config.NewViper()
	validator := config.NewValidator()
	gorm := config.NewGorm(viper)
	echo := config.NewEcho(viper)

	bootstrapInit := bootstrap.Configs{
		Viper:     viper,
		Gorm:      gorm,
		Echo:      echo,
		Validator: validator,
	}

	bootstrapInit.Run()

	echo.Logger.Fatal(echo.Start(viper.GetString("WEB_PORT")))

}

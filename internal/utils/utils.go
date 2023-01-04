package utils

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// read config yaml

func GetConfig() {
	path := "configs/config.yaml"
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	fmt.Printf("%+v", c)
}

package test

import (
	"context"
	"github.com/kostanios/gotemp/services/sqlservice"
	"gorm.io/gorm"

	mysqlcontainer "github.com/testcontainers/testcontainers-go/modules/mysql"
)

type DbContainerData struct {
	ContainerData
	DB *gorm.DB
}

func GetTestingDb() (DbContainerData, error) {
	ctx := context.Background()
	container, err := mysqlcontainer.RunContainer(context.Background(),
		testcontainers.WithImage("mysql:8.1.0"),
		mysqlcontainer.WithDatabase("test"),
		mysqlcontainer.WithUsername("test-username"),
		mysqlcontainer.WithPassword("test-password"),
	)
	if err != nil {
		panic(err)
	}

	url, err := container.Endpoint(ctx, "")
	if err != nil {
		panic(err)
	}

	err, db := sqlservice.New(url, "test-username", "test-password", "test", 10, 10, 10)
	if err != nil {
		panic(err)
	}

	return DbContainerData{
		ContainerData: ContainerData{
			Url:           url,
			TerminateFunc: container.Terminate,
		},
		DB: db,
	}, err
}

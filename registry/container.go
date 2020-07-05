package registry

import (
	sqlitehandler "IkezawaYuki/craft/infrastructure/sqlite_handler"
	"IkezawaYuki/craft/interfaces/controllers"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	if err := builder.Add([]di.Def{
		{
			Name:  "bitflyer-controller",
			Build: buildBitlyerController,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildBitlyerController(ctn di.Container) (interface{}, error) {
	handler := sqlitehandler.NewSQLiteHandler(sqlitehandler.Connect())
	controller := controllers.NewBitlyerController(handler)
	return controller, nil
}

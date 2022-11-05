package main

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	app.logger.Info("Hello, World!", app.logger.Field("config", app.config.Config()))
}

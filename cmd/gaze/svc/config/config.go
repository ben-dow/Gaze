package config

type Configuration struct {
	DatabaseLocation string
	LogLevel         int
	ServerAddress    string
}

var config *Configuration

func InitializeConfiguration() {
	if config != nil {
		return
	}

	cfg := &Configuration{
		DatabaseLocation: "gaze.db",
		LogLevel:         1,
		ServerAddress:    ":3000",
	}

	config = cfg
}

func GetConfiguration() *Configuration {
	if config == nil {
		return nil
	}
	return config
}

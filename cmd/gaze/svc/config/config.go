package config

type Configuration struct {
	DatabaseLocation string
	LogLevel         int
}

var config *Configuration

func InitializeConfiguration() {
	if config != nil {
		return
	}

	cfg := &Configuration{
		DatabaseLocation: "gaze.db",
		LogLevel:         0,
	}

	config = cfg
}

func GetConfiguration() *Configuration {
	if config == nil {
		return nil
	}
	return config
}

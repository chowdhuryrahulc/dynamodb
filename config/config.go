package config //! used my main.go, to get enviournment variables

import (
	"strconv"

	"github.com/chowdhuryrahulc/dynamodb/utils/env"
)

type Config struct {
	Port int // servers will interact with this port
	// If some Error in port, change port number. Bcoz some other program might be using that same port
	Timeout int // if program doesnt work for some time, we get a timeout error, and get to know there is a error.
	// in production, timeout is very imp. Otherwise the request will keep showing pending in the front-end
	// put timeout for all apis backend servers
	Dialect     string // dilects defines if it is sql, or nosql etc. Type of database
	DatabaseURI string // golang will hit this port for dynamodb database
}

func GetConfig() Config {
	// returns values of Config
	return Config{
		Port:        parseEnvToInt("PORT", ":8080"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),    // 30 means 30sec
		Dialect:     env.GetEnv("DIALECT", "sqllite3"), // enviournment can be production, testing etc
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
		//todo Search for Dynamodb uri's in docs. eg: ":memory:", etc Database url starts with :memory:
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	// If you send envName = PORT and it returns 8080, or send TIMEOUT, returns 30
	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue)) // num is int
	if err != nil {
		return 0
	}
	return num
}

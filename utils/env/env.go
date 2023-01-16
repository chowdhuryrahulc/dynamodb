package env

import "os"

func GetEnv(env, defaultValue string) string {
	// this function gets all enviournment variables
	enviournment := os.GetEnv(env)
	if enviournment == ""{
		return defaultValue
	}
	return enviournment
}
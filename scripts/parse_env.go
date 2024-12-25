package scripts

import (
	"fmt"
	"os"
	"strings"
)

func ParseDotEnv() {
	envFile, err := os.ReadFile(".env")

	if err == nil {
		envFile := strings.Split(string(envFile), "\n")

		for i := 0; i < len(envFile)-1; i++ {
			envLine := strings.Split(envFile[i], "=")

			envKey := envLine[0]

			if envKey == "" {
				continue
			}

			envValue := envLine[1]

			if envValue == "" {
				continue
			}

			println(envValue)

			os.Setenv(envKey, envValue)
		}

	} else {
		fmt.Println("Error while reading .env file. Ensure that it exists. From now enviroment variables will be gotten directly from the enviroment of the application.")
	}
}

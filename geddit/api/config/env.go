package config

import (
	"bufio"
	"os"
	"strings"
)

type Env struct {
	DB_TYPE, DB_HOST, DB_NAME, DB_PORT, DB_PASS, DB_USER string
}

func DefaultEnv() *Env {
	file, err := os.Open(".env")
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			kv := strings.Split(line, "=")
			if len(kv) == 2 {
				os.Setenv(kv[0], kv[1])
			}
		}
	}

	return &Env{
		os.Getenv("db_type"),
		os.Getenv("db_host"),
		os.Getenv("db_name"),
		os.Getenv("db_port"),
		os.Getenv("db_pass"),
		os.Getenv("db_user"),
	}
}

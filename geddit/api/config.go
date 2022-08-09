package api

import (
	"bufio"
	"os"
	"strings"
)

type settings struct {
	db_type, db_host, db_name, db_port, db_pass, db_user string
}

var (
	s settings
)

func setenv() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
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

func init() {
	setenv()

	s = settings{
		os.Getenv("db_type"),
		os.Getenv("db_host"),
		os.Getenv("db_name"),
		os.Getenv("db_port"),
		os.Getenv("db_pass"),
		os.Getenv("db_user"),
	}
}

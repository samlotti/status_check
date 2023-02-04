package checker

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port  string
	Rules []IRule
}

func LoadConfig(configName string) {
	fmt.Printf("config: %s\n", configName)
	b, err := os.ReadFile(configName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	for num, line := range lines {
		fmt.Printf("%02d: %s\n", num, line)
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		sections := strings.Split(line, ":")
		switch sections[0] {
		case "port":
			GConfig.Port = sections[1]
		case "process":
			GConfig.Rules = append(GConfig.Rules, ProcessParser(sections))
		case "file":
			GConfig.Rules = append(GConfig.Rules, FileParser(sections))
		default:
			panic(fmt.Sprintf("Invalid rule entry: %s", line))
		}

	}
}

var GConfig = Config{
	Port:  "45241",
	Rules: nil,
}

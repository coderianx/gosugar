package gosugar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
// ENV FILE LOADER
//

func EnvFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("cannot open env file: %s", path))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// boş satır veya yorum
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			panic(fmt.Errorf("invalid env line: %q", line))
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// varsa override etme
		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		if err := os.Setenv(key, value); err != nil {
			panic(fmt.Errorf("failed to set env %s", key))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//
// ENV STRING
//

func EnvString(key string, defaultValue ...string) string {
	value, ok := os.LookupEnv(key)
	if ok && value != "" {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

//
// ENV INT
//

func EnvInt(key string, defaultValue ...int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("missing env var: %s", key))
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("invalid int env var %s=%q", key, value))
	}

	return v
}

//
// ENV BOOL
//

func EnvBool(key string, defaultValue ...bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("missing env var: %s", key))
	}

	switch strings.ToLower(value) {
	case "true", "1", "yes", "y", "on":
		return true
	case "false", "0", "no", "n", "off":
		return false
	default:
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(fmt.Errorf("invalid bool env var %s=%q", key, value))
	}
}

//
// MUST ENV (REQUIRED)
//

func MustEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		panic(fmt.Errorf("required env var missing: %s", key))
	}
	return value
}

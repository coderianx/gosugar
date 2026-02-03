package gosugar

import (
	"fmt"
	"os"
)

//
// READ FILE
//

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("cannot read file %s: %w", path, err))
	}
	return string(data)
}

//
// WRITE FILE
//

func WriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(fmt.Errorf("cannot write file %s: %w", path, err))
	}
}

func CreateFile(path string, content string) {
	// dosya var mı?
	if _, err := os.Stat(path); err == nil {
		// dosya zaten var → sessizce çık
		return
	} else if !os.IsNotExist(err) {
		// başka bir hata varsa
		panic(fmt.Errorf("cannot check file %s: %w", path, err))
	}

	// dosya yok → oluştur
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("cannot create file %s: %w", path, err))
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		panic(fmt.Errorf("cannot write to file %s: %w", path, err))
	}
}

//
// APPEND FILE (create if not exists)
//

func AppendFile(path string, content string) {
	f, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(fmt.Errorf("cannot append to file %s: %w", path, err))
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		panic(fmt.Errorf("cannot append to file %s: %w", path, err))
	}
}

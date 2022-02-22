package autoexe

import (
	"fmt"
	"os"
)

func SaveConfig(config []byte, path, name string) error {
	dst := fmt.Sprintf(`%s\%s.cfg`, path, name)
	return os.WriteFile(dst, config, 0644)
}

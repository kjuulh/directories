package xdg

import (
	"os"
	"path"
	"strings"
)

type UserDirs struct {
}

func (x *UserDirs) All(homeDir string, userDirFile string) (map[string]string, error) {
	raw, err := os.ReadFile(userDirFile)
	if err != nil {
		return nil, err
	}

	file := string(raw)

	userDirs := make(map[string]string, 0)
	for _, l := range strings.Split(file, "\n") {
		vals := strings.SplitN(l, "=", 1)
		if len(vals) != 2 {
			continue
		}
		k := vals[0]
		v := vals[1]

		k = strings.TrimSpace(k)

		if strings.HasPrefix(k, "XDG_") && strings.HasSuffix(k, "_DIR") {
			k = strings.TrimSuffix(strings.TrimPrefix(k, "XDG_"), "_DIR")
		} else {
			continue
		}

		if strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"") {
			v = strings.TrimSuffix(strings.TrimPrefix(v, "\""), "\"")
		} else {
			continue
		}

		if v == "$HOME/" {
			continue
		} else if strings.HasPrefix(v, "$HOME/") {
			v = strings.TrimPrefix(v, "$HOME/")
			v = path.Join(homeDir, v)
		} else if strings.HasPrefix(v, "/") {
			// Do nothing
		} else {
			continue
		}

		userDirs[k] = v
	}

	return userDirs, nil
}

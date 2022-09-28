//go:build linux
// +build linux

package directories

import (
	"fmt"
	"os"
	"path"
)

var _ BaseDirsContract = &baseDirs{}

func NewBaseDirs() (BaseDirsContract, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, false
	}

	return &baseDirs{
		home: homeDir,
		cache: envOrBackup("XDG_CACHE_HOME", func() string {
			return path.Join(homeDir, ".cache")
		}),
		config: envOrBackup("XDG_CONFIG_HOME", func() string {
			return path.Join(homeDir, ".config")
		}),
		data: envOrBackup("XDG_DATA_HOME", func() string {
			return path.Join(homeDir, ".local/share")
		}),
		dataLocal: envOrBackup("XDG_DATA_HOME", func() string {
			return path.Join(homeDir, ".local/share")
		}),
		executable: envOrBackup("XDG_BIN_HOME", func() string {
			return path.Join(homeDir, ".local/bin")
		}),
		preference: envOrBackup("XDG_CONFIG_HOME", func() string {
			return path.Join(homeDir, ".config")
		}),
		runtime: envOrBackup("XDG_RUNTIME_DIR", nil),
		state: envOrBackup("XDG_STATE_HOME", func() string {
			return path.Join(homeDir, ".local/state")
		}),
	}, true
}

func NewUserDirs() (UserDirsContract, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, false
	}

	dataDir := envOrBackup("XDG_DATA_HOME", func() string {
		return path.Join(homeDir, ".local/share")
	})
	fontDir := path.Join(dataDir, "fonts")
	userD := &xdg.UserDirs{}
	userDirsMap, err := userD.All(
		homeDir,
		envOrBackup("XDG_CONFIG_HOME", func() string {
			return path.Join(homeDir, ".config")
		}),
	)
	if err != nil {
		fmt.Printf("could not find userDirs: %s", err)
		return nil, false
	}

	return &userDirs{
		home: homeDir,
		audio: func() (string, bool) {
			return valueOrFalse(userDirsMap["MUSIC"])
		},
		desktop: func() (string, bool) {
			return valueOrFalse(userDirsMap["DESKTOP"])
		},
		document: func() (string, bool) {
			return valueOrFalse(userDirsMap["DOCUMENTS"])
		},
		download: func() (string, bool) {
			return valueOrFalse(userDirsMap["DOWNLOAD"])
		},
		font: func() (string, bool) {
			return valueOrFalse(fontDir)
		},
		picture: func() (string, bool) {
			return valueOrFalse(userDirsMap["PICTURES"])
		},
		public: func() (string, bool) {
			return valueOrFalse(userDirsMap["PUBLICSHARE"])
		},
		template: func() (string, bool) {
			return valueOrFalse(userDirsMap["TEMPLATES"])
		},
		video: func() (string, bool) {
			return valueOrFalse(userDirsMap["VIDEOS"])
		},
	}, true
}

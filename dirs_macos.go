//go:build darwin
// +build darwin

package directories

import (
	"os"
	"path"
)

var _ BaseDirsContract = &baseDirs{}

func NewBaseDirs() (BaseDirsContract, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, false
	}

	cacheDir := path.Join(homeDir, "Library/Caches")
	configDir := path.Join(homeDir, "Library/Application Support")
	dataDir := configDir
	dataLocalDir := dataDir
	preferenceDir := path.Join(homeDir, "Library/Preferences")

	return &baseDirs{
		home:       homeDir,
		cache:      cacheDir,
		config:     configDir,
		data:       dataDir,
		dataLocal:  dataLocalDir,
		preference: preferenceDir,
	}, true
}

func NewUserDirs() (UserDirsContract, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, false
	}

	return &userDirs{
		home: homeDir,
		audio: func() (string, bool) {
			return path.Join(homeDir, "Music"), false
		},
		desktop: func() (string, bool) {
			return path.Join(homeDir, "Desktop"), false
		},
		document: func() (string, bool) {
			return path.Join(homeDir, "Documents"), false
		},
		download: func() (string, bool) {
			return path.Join(homeDir, "Downloads"), false
		},
		font: func() (string, bool) {
			return path.Join(homeDir, "Library/Fonts"), false
		},
		picture: func() (string, bool) {
			return path.Join(homeDir, "Pictures"), false
		},
		public: func() (string, bool) {
			return path.Join(homeDir, "Public"), false
		},
		template: func() (string, bool) {
			return "", false
		},
		video: func() (string, bool) {
			return path.Join(homeDir, "Movies"), false
		},
	}, true
}

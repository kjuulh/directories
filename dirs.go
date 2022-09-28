package directories

import (
	"os"
	"path"
)

type baseDirs struct {
	home       string
	cache      string
	config     string
	data       string
	dataLocal  string
	executable string
	preference string
	runtime    string
	state      string
}

func valueOrFalse(val string) (string, bool) {
	if val == "" {
		return "", false
	}
	return val, true
}

func (bd *baseDirs) HomeDir() (string, bool)       { return valueOrFalse(bd.home) }
func (bd *baseDirs) CacheDir() (string, bool)      { return valueOrFalse(bd.cache) }
func (bd *baseDirs) ConfigDir() (string, bool)     { return valueOrFalse(bd.config) }
func (bd *baseDirs) DataDir() (string, bool)       { return valueOrFalse(bd.data) }
func (bd *baseDirs) DataLocalDir() (string, bool)  { return valueOrFalse(bd.dataLocal) }
func (bd *baseDirs) ExecutableDir() (string, bool) { return valueOrFalse(bd.executable) }
func (bd *baseDirs) PreferenceDir() (string, bool) { return valueOrFalse(bd.preference) }
func (bd *baseDirs) RuntimeDir() (string, bool)    { return valueOrFalse(bd.runtime) }
func (bd *baseDirs) StateDir() (string, bool)      { return valueOrFalse(bd.state) }

type BaseDirsContract interface {
	HomeDir() (string, bool)
	CacheDir() (string, bool)
	ConfigDir() (string, bool)
	DataDir() (string, bool)
	DataLocalDir() (string, bool)
	ExecutableDir() (string, bool)
	PreferenceDir() (string, bool)
	RuntimeDir() (string, bool)
	StateDir() (string, bool)
}

type optionalPath func() (string, bool)

type userDirs struct {
	home string

	audio    optionalPath
	desktop  optionalPath
	document optionalPath
	download optionalPath
	font     optionalPath
	picture  optionalPath
	public   optionalPath
	template optionalPath
	video    optionalPath
}

func (ud *userDirs) HomeDir() (string, bool)     { return valueOrFalse(ud.home) }
func (ud *userDirs) AudioDir() (string, bool)    { return ud.audio() }
func (ud *userDirs) DesktopDir() (string, bool)  { return ud.desktop() }
func (ud *userDirs) DocumentDir() (string, bool) { return ud.document() }
func (ud *userDirs) DownloadDir() (string, bool) { return ud.download() }
func (ud *userDirs) FontDir() (string, bool)     { return ud.font() }
func (ud *userDirs) PictureDir() (string, bool)  { return ud.picture() }
func (ud *userDirs) PublicDir() (string, bool)   { return ud.public() }
func (ud *userDirs) TemplateDir() (string, bool) { return ud.template() }
func (ud *userDirs) VideoDir() (string, bool)    { return ud.video() }

type UserDirsContract interface {
	HomeDir() (string, bool)

	AudioDir() (string, bool)
	DesktopDir() (string, bool)
	DocumentDir() (string, bool)
	DownloadDir() (string, bool)
	FontDir() (string, bool)
	PictureDir() (string, bool)
	PublicDir() (string, bool)
	TemplateDir() (string, bool)
	VideoDir() (string, bool)
}

func envOrBackup(envKey string, backup func() string) string {
	env := os.Getenv(envKey)

	if path.IsAbs(env) {
		return env
	} else {
		if backup != nil {
			return backup()
		}
	}

	return ""
}

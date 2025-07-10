package version

var (
	Version    string
	Revision   string
	DevVersion = "Pre-release"
)

func BuildVersion() string {
	if len(Version) == 0 {
		return DevVersion
	}
	return Version
}

func GetReleaseInfo() (string, string) {
	return BuildVersion(), Revision
}

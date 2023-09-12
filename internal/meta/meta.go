package meta

var (
	Version   string
	BuildTime string
	GitCommit string
	OSArch    string
	GoVersion string
)

var AppInfo = struct {
	Version   string
	BuildTime string
	GitCommit string
	OSArch    string
	GoVersion string
}{
	Version:   Version,
	BuildTime: BuildTime,
	GitCommit: GitCommit,
	OSArch:    OSArch,
	GoVersion: GoVersion,
}

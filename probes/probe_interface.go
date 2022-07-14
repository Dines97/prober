package probes

type RunResult string

const (
	Success RunResult = "Success"
	Fail    RunResult = "Fail"
	Timeout RunResult = "Timeout"
	Unknown RunResult = "Unknown"
)

type ProbeRunner interface {
	Run() (RunResult, error)
}

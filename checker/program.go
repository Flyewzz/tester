package checker

type Program struct {
	Path        string
	MemoryLimit string
	DiskLimit   string
	CpuLimit    string
}

func NewProgram(path, memoryLimit, diskLimit, cpuLimit string) *Program {
	return &Program{
		Path:        path,
		MemoryLimit: memoryLimit,
		DiskLimit:   diskLimit,
		CpuLimit:    cpuLimit,
	}
}

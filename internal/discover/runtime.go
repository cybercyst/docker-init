package discover

type RuntimeType uint8

func (t RuntimeType) ToString() string {
	return []string{"Golang", "Node", "Python", "Scratch"}[t]
}

const (
	Go RuntimeType = iota
	Node
	Python
	None
)

type Runtime struct {
	Type    RuntimeType
	Version string
}

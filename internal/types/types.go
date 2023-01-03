package types

type Target struct {
	TargetType TargetType
	Path       string
	Input      map[string]interface{}
}

type TargetType uint8

func (t TargetType) ToString() string {
	return []string{"Golang", "Angular", "None"}[t]
}

const (
	Go TargetType = iota
	Angular
	None
)

package DTreeMaster

import "io"
import "DTreeMaster/id3"

type DTree interface {
	LoadData(io.ReadSeeker)
	AppendValue()
}

var (
	treeMap = make(map[string]DTree, 0)
)

func init() {
	treeMap[id3.AlgorithmName] = new(id3.ID3)
}

func NewDTree() (DTree, error) {
	return nil, nil
}

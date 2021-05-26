package id3

import "io"

const AlgorithmName = "id3_algo"

type ID3 struct {
}

func (I *ID3) LoadData(seeker io.ReadSeeker) {
	panic("implement me")
}

func (I *ID3) AppendValue() {
	panic("implement me")
}

package gomisc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MainStruct struct {
	MID string
	SubStruct
	mNum int
}

type SubStruct struct {
	ID  string
	num int
}

type CustomType int64

func newTestStruct() *MainStruct {
	ms := new(MainStruct)
	ms.MID = "MainID"
	ms.ID = "subID"
	ms.num = 19
	ms.mNum = 734

	return ms
}

type StructA struct {
	pb *StructB
}

type StructB struct {
	pa *StructA
}

func TestGetFieldNamesRecursively(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		structVal := newTestStruct()
		var intVal int64 = 32
		var cusVal CustomType = 255

		tests := []struct {
			in  interface{}
			out []string
		}{
			{structVal, []string{"MID", "ID", "num", "mNum"}},
			{*structVal, []string{"MID", "ID", "num", "mNum"}},
			{intVal, []string{"int64"}},
			{&intVal, []string{"int64"}},
			{cusVal, []string{"CustomType"}},
			{&cusVal, []string{"CustomType"}},
		}

		for i, test := range tests {
			assert.Equal(t, test.out, GetFieldNamesRecursively(test.in), i)
		}
	})
}

func BenchmarkGetFieldNamesRecur(b *testing.B) {
	b.ReportAllocs()
	r := newTestStruct()

	for i := 0; i < b.N; i++ {
		GetFieldNamesRecursively(r)
	}
}

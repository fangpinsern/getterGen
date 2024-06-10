package demo

//go:generate go run ../tools/getterGen/getterGen.go
type Complex struct {
	NumberI64       int64
	PtrNumberI64    *int64
	ArrNumberI64    []int64
	ArrPtrNumberI64 []*int64
	NumberI32       int32
	PtrNumberI32    *int32
	NumberI16       int16
	PtrNumberI16    *int16
	Str             string
	PtrStr          *string
	ArrStr          []string
	ArrStruct       []*Complex
	BaseStruct      SecondStruct
	PtrStruct       *SecondStruct
	MapTest         map[string]*SecondStruct
	ArrMapTest      []map[string]*SecondStruct
	// no need maps at the moment. for our use case, maps will be unserialized to structs
}

type SecondStruct struct {
	ArrStr    []string
	ArrStruct []*SecondStruct
}

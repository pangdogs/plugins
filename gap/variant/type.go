package variant

import (
	"kit.golaxy.org/plugins/util/binaryutil"
	"reflect"
)

// TypeId 类型Id
type TypeId uint32

// Read implements io.Reader
func (t TypeId) Read(p []byte) (int, error) {
	bs := binaryutil.NewBigEndianStream(p)
	if err := bs.WriteUvarint(uint64(t)); err != nil {
		return bs.BytesWritten(), err
	}
	return bs.BytesWritten(), nil
}

// Write implements io.Writer
func (t *TypeId) Write(p []byte) (int, error) {
	bs := binaryutil.NewBigEndianStream(p)

	v, err := bs.ReadUvarint()
	if err != nil {
		return bs.BytesRead(), err
	}
	*t = TypeId(v)

	return bs.BytesRead(), nil
}

// Size 大小
func (t TypeId) Size() int {
	return binaryutil.SizeofUvarint(uint64(t))
}

// New 创建对象指针
func (t TypeId) New() (Value, error) {
	return variantCreator.New(t)
}

// Make 创建对象
func (t TypeId) Make() (ValueReader, error) {
	return variantCreator.Make(t)
}

// NewReflected 创建反射对象指针
func (t TypeId) NewReflected() (reflect.Value, error) {
	return variantCreator.NewReflected(t)
}

// MakeReflected 创建反射对象
func (t TypeId) MakeReflected() (reflect.Value, error) {
	return variantCreator.MakeReflected(t)
}

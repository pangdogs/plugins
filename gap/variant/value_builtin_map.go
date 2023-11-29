package variant

import (
	"kit.golaxy.org/plugins/util/binaryutil"
)

// MakeMap 创建Map
func MakeMap[K comparable, V any](m map[K]V) (Map, error) {
	var varMap Map
	var err error

	for k, v := range m {
		var kv KV

		kv.K, err = CastVariant(k)
		if err != nil {
			return nil, err
		}

		kv.V, err = CastVariant(v)
		if err != nil {
			return nil, err
		}

		varMap = append(varMap, kv)
	}

	return varMap, nil
}

// KV kv
type KV struct {
	K, V Variant
}

// Map map
type Map []KV

// Read implements io.Reader
func (v Map) Read(p []byte) (int, error) {
	rn := 0

	bs := binaryutil.NewBigEndianStream(p)
	if err := bs.WriteUvarint(uint64(len(v))); err != nil {
		return rn, err
	}
	rn += bs.BytesWritten()

	for i := range v {
		kv := &v[i]

		n, err := kv.K.Read(p[rn:])
		rn += n
		if err != nil {
			return rn, err
		}

		n, err = kv.V.Read(p[rn:])
		rn += n
		if err != nil {
			return rn, err
		}
	}

	return rn, nil
}

// Write implements io.Writer
func (v *Map) Write(p []byte) (int, error) {
	wn := 0

	bs := binaryutil.NewBigEndianStream(p)
	l, err := bs.ReadUvarint()
	if err != nil {
		return wn, err
	}
	wn += bs.BytesRead()

	kvs := make([]KV, l)

	for i := uint64(0); i < l; i++ {
		kv := &kvs[i]

		n, err := kv.K.Write(p[wn:])
		wn += n
		if err != nil {
			return wn, err
		}

		n, err = kv.V.Write(p[wn:])
		wn += n
		if err != nil {
			return wn, err
		}
	}
	*v = kvs

	return wn, nil
}

// Size 大小
func (v Map) Size() int {
	n := binaryutil.SizeofUvarint(uint64(len(v)))
	for i := range v {
		kv := &v[i]
		n += kv.K.Size()
		n += kv.V.Size()
	}
	return n
}

// Type 类型
func (Map) Type() TypeId {
	return TypeId_Map
}
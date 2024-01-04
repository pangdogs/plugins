package concurrent

func MakeLockedMap[K comparable, V any](size int) LockedMap[K, V] {
	return LockedMap[K, V]{
		Locked: MakeLocked(make(map[K]V, size)),
	}
}

type LockedMap[K comparable, V any] struct {
	Locked[map[K]V]
}

func (lm *LockedMap[K, V]) Insert(k K, v V) {
	lm.AutoLock(func(m *map[K]V) {
		(*m)[k] = v
	})
}

func (lm *LockedMap[K, V]) Delete(k K) {
	lm.AutoLock(func(m *map[K]V) {
		delete(*m, k)
	})
}

func (lm *LockedMap[K, V]) Get(k K) (v V, ok bool) {
	lm.AutoRLock(func(m *map[K]V) {
		v, ok = (*m)[k]
	})
	return
}

func (lm *LockedMap[K, V]) Len() (l int) {
	lm.AutoRLock(func(m *map[K]V) {
		l = len(*m)
	})
	return
}
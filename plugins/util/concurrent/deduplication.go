package concurrent

import (
	"sync/atomic"
	"time"
)

// IDeduplication 去重器接口
type IDeduplication interface {
	// MakeSeq 创建序号
	MakeSeq() int64
	// ValidateSeq 验证序号
	ValidateSeq(remote string, seq int64) bool
	// Remove 删除对端
	Remove(remote string)
}

// MakeDeduplication 创建去重器
func MakeDeduplication() Deduplication {
	return Deduplication{
		localSeq:     time.Now().UnixMicro(),
		remoteSeqMap: MakeLockedMap[string, *_RemoteSeq](0),
	}
}

type _RemoteSeq = Locked[int64]

// Deduplication 去重器，用于保持幂等性
type Deduplication struct {
	localSeq     int64
	remoteSeqMap LockedMap[string, *_RemoteSeq]
}

// MakeSeq 创建序号
func (d *Deduplication) MakeSeq() int64 {
	return atomic.AddInt64(&d.localSeq, 1)
}

// ValidateSeq 验证序号
func (d *Deduplication) ValidateSeq(remote string, seq int64) (passed bool) {
	remoteSeq, ok := d.remoteSeqMap.Get(remote)
	if !ok {
		var firstInsert bool

		d.remoteSeqMap.AutoLock(func(m *map[string]*_RemoteSeq) {
			remoteSeq, ok = (*m)[remote]
			if !ok {
				remoteSeq = NewLocked[int64](seq)
				(*m)[remote] = remoteSeq

				firstInsert = true
			}
		})

		if firstInsert {
			return true
		}
	}

	remoteSeq.AutoLock(func(remoteSeq *int64) {
		if seq <= *remoteSeq {
			return
		}
		*remoteSeq = seq
		passed = true
	})

	return
}

// Remove 删除对端
func (d *Deduplication) Remove(remote string) {
	d.remoteSeqMap.Delete(remote)
}
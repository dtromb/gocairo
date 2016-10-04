package cairo

import (
	"math/rand"
	"sync"
	"time"
)

type InterfaceRef struct {
	x interface{}
}

var global_ReferenceStore stdReferenceStore

func init() {
	global_ReferenceStore = stdReferenceStore{
		refs: make(map[uint32]Reference),
		rng:  rand.NewSource(time.Now().UnixNano()),
	}
}

func MakeGlobalReference(obj interface{}) Reference {
	return global_ReferenceStore.MakeReference(obj)
}

func LookupGlobalReference(key uint32) (Reference, bool) {
	return global_ReferenceStore.LookupReference(key)
}

func IncrementGlobalReferenceCount(ref Reference) {
	global_ReferenceStore.IncrementReference(ref)
}

func DecrementGlobalReferenceCount(ref Reference) {
	global_ReferenceStore.DecrementReference(ref)
}

type ReferenceStore interface {
	MakeReference(obj interface{}) Reference
	LookupReference(key uint32) (Reference, bool)
	IncrementReference(r Reference) bool
	DecrementReference(r Reference) bool
}

type Reference interface {
	Ref() interface{}
	Key() uint32
	Cleared() bool
}

type stdReferenceStore struct {
	lock sync.Mutex
	refs map[uint32]Reference
	rng  rand.Source
}

func (rs *stdReferenceStore) MakeReference(obj interface{}) Reference {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	nkey := uint32(rs.rng.Int63())
	for {
		if _, has := rs.refs[nkey]; has {
			nkey = uint32(rs.rng.Int63())
		} else {
			break
		}
	}
	ref := &stdReference{
		key: nkey,
		ref: obj,
		rc:  1,
	}
	return ref
}

func (rs *stdReferenceStore) LookupReference(key uint32) (Reference, bool) {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	ref, has := rs.refs[key]
	return ref, has
}

func (rs *stdReferenceStore) IncrementReference(r Reference) bool {
	rs.lock.Lock()
	defer rs.lock.Unlock()
	if sr, ok := r.(*stdReference); ok {
		if sr.rc > 0 {
			sr.rc += 1
			return true
		}
	}
	return false
}

func (rs *stdReferenceStore) DecrementReference(r Reference) bool {
	rs.lock.Lock()
	rs.lock.Unlock()
	if sr, ok := r.(*stdReference); ok {
		if sr.rc > 0 {
			sr.rc -= 1
			if sr.rc == 0 {
				if rs.refs[sr.key] == sr {
					delete(rs.refs, sr.key)
				}
			} else {
				return true
			}
		}
	}
	return false
}

type stdReference struct {
	key uint32
	rc  uint32
	ref interface{}
}

func (r *stdReference) Ref() interface{} {
	return r.ref
}

func (r *stdReference) Key() uint32 {
	return r.key
}

func (r *stdReference) Incref() {
	r.rc += 1
}

func (r *stdReference) Cleared() bool {
	return r.rc == 0
}

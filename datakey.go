package cairo

import (
	"time"
	"sync"
	"math/rand"
)

/*

	#cgo LDFLAGS: -lcairo
    #include <cairo/cairo.h>
	#include <inttypes.h>
	#include <stdlib.h>
	
	cairo_user_data_key_t *cgo_get_cairo_userdata_key(int32_t keyid) {
		return (cairo_user_data_key_t*)((size_t)keyid);
	}
	
	uint32_t cgo_get_refkey(void *cref) {
		return (uint32_t)((size_t)cref);
	}
	
	void* cgo_get_keyref(uint32_t key) {
		return (void*)((size_t)key);
	}

*/
import "C"


const GO_DATAKEY_KEY = 0x000010B0

type DataKey interface {
	String() string
	Key() uint32
	Return() 
}

func GetDataKey(name string) DataKey {
	global_datakeyContext.lock.Lock()
	defer global_datakeyContext.lock.Unlock()
	if entry, has := global_datakeyContext.byName[name]; has {
		entry.rc += 1
		return entry
	}
	key := uint32(global_datakeyContext.rng.Int63())
	for {
		if _, has := global_datakeyContext.byId[key]; has {
			key = uint32(global_datakeyContext.rng.Int63())
		} else {
			break
		}
	}
	nkey := &stdDataKey{
		name: name,
		key: key,
		rc: 1,
	}
	global_datakeyContext.byName[name] = nkey
	global_datakeyContext.byId[key] = nkey
	return nkey
}


func LookupDataKey(key uint32) DataKey {
	global_datakeyContext.lock.Lock()
	defer global_datakeyContext.lock.Unlock()
	if entry, has := global_datakeyContext.byId[key]; has {
		entry.rc += 1
		return entry
	}
	return nil
}

type stdDataKey struct{
	name string
	key uint32
	rc uint32
	
}

type datakeyContext struct {
	lock sync.Mutex
	byName map[string]*stdDataKey
	byId map[uint32]*stdDataKey
	rng rand.Source
}

var global_datakeyContext datakeyContext

func init() {
	global_datakeyContext = datakeyContext {
		byName: make(map[string]*stdDataKey),
		byId: make(map[uint32]*stdDataKey),
		rng: rand.NewSource(time.Now().UnixNano()),
	}
}

func (dk *stdDataKey) String() string {
	return dk.name
}

func (dk *stdDataKey) Key() uint32 {
	return dk.key
}

func (dk *stdDataKey) Return() {
	global_datakeyContext.lock.Lock()
	defer global_datakeyContext.lock.Unlock()
	if dk.rc == 0 {
		return
	}
	dk.rc -= 1
	if dk.rc == 0 {
		delete(global_datakeyContext.byId, dk.key)
		delete(global_datakeyContext.byName, dk.name)
	}
}
package refer

import (
	"reflect"
	"runtime"
	"strings"
)

type empty struct{}

var referPkgName = reflect.TypeOf(empty{}).PkgPath()

func outerPkgName() string {
	for i := 2; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		name := fn.Name()
		lastSlash := strings.LastIndex(name, "/")
		firstDot := strings.Index(name[lastSlash+1:], ".")
		name = name[:lastSlash+firstDot+1]
		if name == referPkgName {
			continue
		}
		return name
	}
	return ""
}

func newValuePtrByType(p reflect.Type) any {
	return reflect.New(p).Interface()
}

func newValuePtr[T any]() *T {
	ptrType := reflect.TypeOf((*T)(nil)).Elem()
	ptr := newValuePtrByType(ptrType)
	return ptr.(*T)
}

func getType[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}

func getActualType[T any]() reflect.Type {
	var t T
	return actualTypeOf(t)
}

func isPointer(o any) bool {
	types := reflect.TypeOf(o)
	if types.Kind() == reflect.Ptr {
		return true
	}
	return false
}

func isStructPointer(o any) bool {
	if !isPointer(o) {
		return false
	}
	kind := reflect.TypeOf(o).Elem().Kind()
	if kind == reflect.Struct {
		return true
	}
	return false
}

func actualTypeOf(o any) reflect.Type {
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func actualValueOf(o any) reflect.Value {
	t := reflect.ValueOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func toGenericPtr[T any](o any) (ptr *T) {
	if o == nil {
		return nil
	}
	if val, ok := o.(*T); ok {
		return val
	}
	return nil
}

func isAllowedKey(key string) bool {
	if key == "" {
		return false
	}
	return true
}

package refer

func Bind[T any](v *T) *T {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	ct.storeRef("", v)
	return v
}

func BindWithKey[T any](key string, v T) *T {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	ct.storeRef(key, v)
	return &v
}

func Unbind[T any]() error {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	if _, ok := ct.lookupRef(""); ok {
		ct.deleteRef("")
	}
	return ErrRefWithDefKeyNotFound
}

func UnbindWithKey[T any](key string) error {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	if !ct.hasRef(key) {
		return ErrRefNotFound
	}
	ct.deleteRef(key)
	return nil
}

func Ref[T any]() *T {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	return toGenericPtr[T](ct.loadRef(""))
}

func RefWithKey[T any](key string) *T {
	ns := curNamespace()
	ct := ns.loadRefContainer(getActualType[T]())
	return toGenericPtr[T](ct.loadRef(key))
}

func Invoke[T any](fn func(*T) error) error {
	v := Ref[T]()
	if v == nil {
		return ErrRefWithDefKeyNotFound
	}
	return fn(v)
}

func InvokeWithKey[T any](key string, fn func(*T) error) error {
	v := RefWithKey[T](key)
	if v == nil {
		return ErrRefNotFound
	}
	return fn(v)
}

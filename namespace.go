package refer

import "reflect"

type namespace struct {
	containers map[reflect.Type]*refContainer
}

func newNamespace() *namespace {
	return &namespace{
		containers: make(map[reflect.Type]*refContainer),
	}
}

func (ns *namespace) loadRefContainer(typ reflect.Type) *refContainer {
	if c, ok := ns.containers[typ]; ok {
		return c
	}
	ns.containers[typ] = newContainer()
	return ns.containers[typ]
}

var (
	globalNS           = newNamespace()
	namespaceMap       = make(map[string]*namespace)
	namespaceMapByPkg  = make(map[string]*namespace)
	namespaceMapByName = make(map[string]*namespace)
)

type NSType int

const (
	NSGlobal = iota
	NSPkg
	NSName
)

var DefaultNSType NSType = NSGlobal

func Namespace(nsType NSType, ns ...string) {
	pkgName := outerPkgName()
	switch nsType {
	case NSGlobal:
		namespaceMap[pkgName] = globalNS
	case NSPkg:
		// create a new namespace for the package
		if _, ok := namespaceMapByPkg[pkgName]; !ok {
			namespaceMapByPkg[pkgName] = newNamespace()
		}
		namespaceMap[pkgName] = namespaceMapByPkg[pkgName]
	case NSName:
		// create a new custom namespace for the given name
		if len(ns) != 0 {
			name := ns[0]
			if _, ok := namespaceMapByName[name]; !ok {
				namespaceMapByName[name] = newNamespace()
			}
			namespaceMap[name] = namespaceMapByName[name]
		} else {
			panic("Name of namespace cannot be empty")
		}
	}
}

func curNamespace() *namespace {
	name := outerPkgName()
	if _, ok := namespaceMap[name]; !ok {
		// default to global namespace
		Namespace(DefaultNSType)
	}
	return namespaceMap[name]
}

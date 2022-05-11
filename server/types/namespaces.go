package types

type NamespacedRequest interface {
	GetNamespace() string
}

func Namespace(req interface{}) string {
	if v, ok := req.(NamespacedRequest); ok {
		return v.GetNamespace()
	}
	return ""
}

type NamespaceHolder string

func (n NamespaceHolder) GetNamespace() string {
	return string(n)
}

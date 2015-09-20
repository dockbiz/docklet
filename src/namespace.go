package src

import (
	"fmt"
	"regexp"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

// validNamespace matches valid namespace names.
var validNamespace = regexp.MustCompile(`^[0-9A-Za-z._-]{0,100}$`)

// namespacedContext wraps a Context to support namespaces.
type namespacedContext struct {
	namespace string
}

func Namespace(c context.Context, namespace string) (context.Context, error) {
	if !validNamespace.MatchString(namespace) {
		return nil, fmt.Errorf("app proxy namespace %q does not match /%s/", namespace, validNamespace)
	}
	n := &namespacedContext{
		namespace: namespace,
	}
	return internal.WithNamespace(internal.WithCallOverride(c, n.call), namespace), nil
}

func (n *namespacedContext) call(ctx context.Context, service, method string, in, out proto.Message) error {
	if mod, ok := internal.NamespaceMods[service]; ok {
		mod(in, n.namespace)
	}
	return internal.Call(ctx, service, method, in, out)
}

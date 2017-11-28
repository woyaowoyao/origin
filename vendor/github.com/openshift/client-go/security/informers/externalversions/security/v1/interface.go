// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/client-go/security/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SecurityContextConstraintses returns a SecurityContextConstraintsInformer.
	SecurityContextConstraintses() SecurityContextConstraintsInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// SecurityContextConstraintses returns a SecurityContextConstraintsInformer.
func (v *version) SecurityContextConstraintses() SecurityContextConstraintsInformer {
	return &securityContextConstraintsInformer{factory: v.SharedInformerFactory}
}
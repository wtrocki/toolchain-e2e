package test

import (
	"github.com/codeready-toolchain/api/pkg/apis"
	commontest "github.com/codeready-toolchain/toolchain-common/pkg/test"

	quotav1 "github.com/openshift/api/quota/v1"
	operatorsv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func NewFakeClient(t commontest.T, initObjs ...runtime.Object) *commontest.FakeClient {
	s := scheme.Scheme
	builder := append(apis.AddToSchemes, quotav1.Install, operatorsv1alpha1.AddToScheme)
	err := builder.AddToScheme(s)
	require.NoError(t, err)
	cl := fake.NewFakeClientWithScheme(s, initObjs...)
	return &commontest.FakeClient{Client: cl, T: t}
}

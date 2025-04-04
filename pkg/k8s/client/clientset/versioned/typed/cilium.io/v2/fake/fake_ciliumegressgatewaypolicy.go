// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/typed/cilium.io/v2"
	gentype "k8s.io/client-go/gentype"
)

// fakeCiliumEgressGatewayPolicies implements CiliumEgressGatewayPolicyInterface
type fakeCiliumEgressGatewayPolicies struct {
	*gentype.FakeClientWithList[*v2.CiliumEgressGatewayPolicy, *v2.CiliumEgressGatewayPolicyList]
	Fake *FakeCiliumV2
}

func newFakeCiliumEgressGatewayPolicies(fake *FakeCiliumV2) ciliumiov2.CiliumEgressGatewayPolicyInterface {
	return &fakeCiliumEgressGatewayPolicies{
		gentype.NewFakeClientWithList[*v2.CiliumEgressGatewayPolicy, *v2.CiliumEgressGatewayPolicyList](
			fake.Fake,
			"",
			v2.SchemeGroupVersion.WithResource("ciliumegressgatewaypolicies"),
			v2.SchemeGroupVersion.WithKind("CiliumEgressGatewayPolicy"),
			func() *v2.CiliumEgressGatewayPolicy { return &v2.CiliumEgressGatewayPolicy{} },
			func() *v2.CiliumEgressGatewayPolicyList { return &v2.CiliumEgressGatewayPolicyList{} },
			func(dst, src *v2.CiliumEgressGatewayPolicyList) { dst.ListMeta = src.ListMeta },
			func(list *v2.CiliumEgressGatewayPolicyList) []*v2.CiliumEgressGatewayPolicy {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2.CiliumEgressGatewayPolicyList, items []*v2.CiliumEgressGatewayPolicy) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

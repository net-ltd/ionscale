package tests

import (
	"net/netip"
	"testing"

	"github.com/jsiebens/ionscale/pkg/client/ionscale"
	"github.com/jsiebens/ionscale/tests/sc"
	"github.com/jsiebens/ionscale/tests/tsn"
	"github.com/stretchr/testify/require"
	"tailscale.com/tailcfg"
)

func TestAppConnector_AppliesCapMapToNodeWithTag(t *testing.T) {
	testTag := "tag:test-app-connector"
	routeAllIPv4 := netip.MustParsePrefix("0.0.0.0/0")
	routeAllIPv6 := netip.MustParsePrefix("::/0")

	autoApprovers := ionscale.ACLAutoApprovers{
		Routes: map[string][]string{
			routeAllIPv4.String(): {testTag},
			routeAllIPv6.String(): {testTag},
		},
	}

	sc.Run(t, func(s *sc.Scenario) {
		aclPolicy := ionscale.ACLPolicy{
			ACLs: []ionscale.ACLEntry{
				{
					Action:      "accept",
					Source:      []string{"*"},
					Destination: []string{"*:*"},
				},
			},
			AutoApprovers: &autoApprovers,
			NodeAttrs: []ionscale.ACLNodeAttrGrant{
				{
					Target: []string{"*"},
					App: (tailcfg.NodeCapMap)(map[tailcfg.NodeCapability][]tailcfg.RawMessage{
						"tailscale.com/app-connectors": {tailcfg.RawMessage(`{
															"name": "example",
															"domains": ["example.com"],
															"connectors": ["tag:app-router"]
														}`)},
					}),
				},
			},
		}

		tailnet := s.CreateTailnet()
		s.SetACLPolicy(tailnet.Id, &aclPolicy)

		testNode := s.NewTailscaleNode()
		require.NoError(t, testNode.Up(
			s.CreateAuthKey(tailnet.Id, true),
		))

		testAppConnectorNode := s.NewTailscaleNode()
		require.NoError(t, testAppConnectorNode.Up(
			s.CreateAuthKey(tailnet.Id, true, testTag),
			tsn.WithAdvertiseTags(testTag),
			tsn.WithAdvertiseConnector(),
		))

		require.NoError(t, testNode.WaitFor(tsn.HasTailnet(tailnet.Name)))
		require.NoError(t, testAppConnectorNode.WaitFor(tsn.HasTailnet(tailnet.Name)))

		_, err := s.FindMachine(tailnet.Id, testAppConnectorNode.Hostname())
		require.NoError(t, err)

		midTestAppConnectorNode, err := s.FindMachine(tailnet.Id, testAppConnectorNode.Hostname())
		require.NoError(t, err)

		// Check the app connector isn't advertising any routes at the moment
		machineRoutes := s.GetMachineRoutes(midTestAppConnectorNode)
		require.NoError(t, err)
		require.Equal(t, []string(nil), machineRoutes.AdvertisedRoutes)
		require.Equal(t, []string(nil), machineRoutes.EnabledRoutes)

		// Check the nodes have the correct capability as assigned in the ACL Policy
		require.NoError(t, testNode.Check(tsn.HasCapability("tailscale.com/app-connectors")))
		require.NoError(t, testAppConnectorNode.Check(tsn.HasCapability("tailscale.com/app-connectors")))
	})
}

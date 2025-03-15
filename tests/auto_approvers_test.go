package tests

import (
	"github.com/jsiebens/ionscale/pkg/client/ionscale"
	"github.com/jsiebens/ionscale/pkg/defaults"
	"github.com/jsiebens/ionscale/tests/sc"
	"github.com/jsiebens/ionscale/tests/tsn"
	"github.com/stretchr/testify/require"
	"net/netip"
	"testing"
	"time"
)

// This will test that the coordinator will automatically approve
// the routes that we have ACL's for and ignore routes that we don't during
// node registration
func TestAdvertiseRoutesAutoApproverOnNewNode(t *testing.T) {
	route1 := netip.MustParsePrefix("10.1.0.0/24")
	route2 := netip.MustParsePrefix("10.2.0.0/24")

	sc.Run(t, func(s *sc.Scenario) {
		aclPolicy := defaults.DefaultACLPolicy()
		aclPolicy.AutoApprovers = &ionscale.ACLAutoApprovers{
			Routes: map[string][]string{
				route1.String(): {"tag:test-route"},
			},
		}

		tailnet := s.CreateTailnet()
		s.SetACLPolicy(tailnet.Id, aclPolicy)

		testNode := s.NewTailscaleNode()
		// Register the node and advertise routes
		require.NoError(t, testNode.Up(
			s.CreateAuthKey(tailnet.Id, true, "tag:test-route"),
			tsn.WithAdvertiseTags("tag:test-route"),
			tsn.WithAdvertiseRoutes([]string{
				route1.String(),
				route2.String()},
			),
		))

		require.NoError(t, testNode.Check(tsn.HasTailnet(tailnet.Name)))

		mid, err := s.FindMachine(tailnet.Id, testNode.Hostname())
		require.NoError(t, err)

		machineRoutes := s.GetMachineRoutes(tailnet.Id, mid)
		require.NoError(t, err)

		// Check that the coordinator server has record of the routes we requested to advertise
		require.Equal(t, []string{route1.String(), route2.String()}, machineRoutes.AdvertisedRoutes)

		// Check that the coordinator server has 'AutoApproved' the one route which the ACL allows
		require.Equal(t, []string{route1.String()}, machineRoutes.EnabledRoutes)

		require.NoError(t, testNode.Check(tsn.HasRoute(route1)))
	})
}

// This will test that the coordinator will automatically approve
// the routes that we have ACL's for and ignore routes that we don't
// after the node is already registered
func TestAdvertiseRoutesAutoApproverOnExistingNode(t *testing.T) {
	route1 := netip.MustParsePrefix("10.1.0.0/24")
	route2 := netip.MustParsePrefix("10.2.0.0/24")
	route3 := netip.MustParsePrefix("10.3.0.0/24")

	sc.Run(t, func(s *sc.Scenario) {
		aclPolicy := defaults.DefaultACLPolicy()
		aclPolicy.AutoApprovers = &ionscale.ACLAutoApprovers{
			Routes: map[string][]string{
				route1.String(): {"tag:test-route"},
				route3.String(): {"tag:test-route"},
			},
		}

		tailnet := s.CreateTailnet()
		s.SetACLPolicy(tailnet.Id, aclPolicy)

		testNode := s.NewTailscaleNode()
		// Register the node, but don't advertise any routes just yet!
		require.NoError(t, testNode.Up(
			s.CreateAuthKey(tailnet.Id, true, "tag:test-route"),
			tsn.WithAdvertiseTags("tag:test-route"),
		))

		require.NoError(t, testNode.Check(tsn.HasTailnet(tailnet.Name)))

		testNode.Set(tsn.WithAdvertiseRoutes([]string{
			route3.String(),
			route1.String(),
			route2.String()},
		))

		time.Sleep(3 * time.Second)

		mid, err := s.FindMachine(tailnet.Id, testNode.Hostname())
		require.NoError(t, err)

		machineRoutes := s.GetMachineRoutes(tailnet.Id, mid)
		require.NoError(t, err)

		// Check that the coordinator server has record of the routes we requested to advertise
		require.Equal(t, []string{route1.String(), route2.String(), route3.String()}, machineRoutes.AdvertisedRoutes)

		// Check that the coordinator server has 'AutoApproved' the routes which the ACL allows
		require.Equal(t, []string{route1.String(), route3.String()}, machineRoutes.EnabledRoutes)

		require.NoError(t, testNode.Check(tsn.HasRoute(route1)))
	})
}

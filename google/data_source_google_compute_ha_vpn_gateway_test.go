// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceComputeHaVpnGateway(t *testing.T) {
	t.Parallel()

	gwName := fmt.Sprintf("tf-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceComputeHaVpnGatewayConfig(gwName),
				Check:  CheckDataSourceStateMatchesResourceState("data.google_compute_ha_vpn_gateway.ha_gateway", "google_compute_ha_vpn_gateway.ha_gateway"),
			},
		},
	})
}

func testAccDataSourceComputeHaVpnGatewayConfig(gwName string) string {
	return fmt.Sprintf(`
resource "google_compute_ha_vpn_gateway" "ha_gateway" {
  name     = "%s"
  network  = google_compute_network.network1.id
}

resource "google_compute_network" "network1" {
  name                    = "%s"
  auto_create_subnetworks = false
}

data "google_compute_ha_vpn_gateway" "ha_gateway" {
  name = google_compute_ha_vpn_gateway.ha_gateway.name
}
`, gwName, gwName)
}

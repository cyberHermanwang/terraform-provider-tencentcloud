package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTencentCloudAuditCosRegionsDataSourc(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudAuditCosRegionsDataSourceConfigWithWebsite,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_audit_cos_regions.filter"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_audit_cos_regions.filter", "cos_region_list.#"),
				),
			},
		},
	})
}

const testAccTencentCloudAuditCosRegionsDataSourceConfigWithWebsite = `
data "tencentcloud_audit_cos_regions" "filter" {
	website_type = "zh"
}
`

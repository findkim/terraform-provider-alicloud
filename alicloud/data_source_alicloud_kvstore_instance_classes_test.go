package alicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAlicloudKVStoreInstanceClasses(t *testing.T) {
	rand := acctest.RandInt()
	resourceId := "data.alicloud_kvstore_instance_classes.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId, "KVStore", kvstoreConfigHeader)

	EngineVersionConfRedis := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":        "${data.alicloud_zones.resources.zones.0.id}",
			"engine":         "Redis",
			"engine_version": "5.0",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"zone_id":        "${data.alicloud_zones.resources.zones.0.id}",
			"engine":         "Redis",
			"engine_version": "4.9",
		}),
	}

	// At present, the datasource does not support memcache
	//EngineVersionConfMemcache := dataSourceTestAccConfig{
	//	existConfig: testAccConfig(map[string]interface{}{
	//		"zone_id": "${data.alicloud_zones.resources.zones.0.id}",
	//		"engine":  "Memcache",
	//	}),
	//}
	ChargeTypeConfPostpaid := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":              "${data.alicloud_zones.resources.zones.0.id}",
			"instance_charge_type": "PostPaid",
		}),
	}
	PerformanceTypeStandardPerformanceType := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":          "${data.alicloud_zones.resources.zones.0.id}",
			"performance_type": "standard_performance_type",
		}),
	}
	PerformanceTypeEnhancePerformanceType := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":          "${data.alicloud_zones.resources.zones.0.id}",
			"performance_type": "enhance_performance_type",
		}),
	}
	StorageTypeInmemory := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"storage_type": "inmemory",
		}),
	}
	PackageTypeStandard := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"package_type": "standard",
		}),
	}
	PackageTypeCustomized := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"package_type": "customized",
		}),
	}
	ArchitectureStandard := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"architecture": "standard",
		}),
	}
	ArchitectureCluster := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"architecture": "cluster",
		}),
	}
	ArchitectureRwsplit := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":      "${data.alicloud_zones.resources.zones.0.id}",
			"architecture": "rwsplit",
		}),
	}
	NodeTypeDouble := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":   "${data.alicloud_zones.resources.zones.0.id}",
			"node_type": "double",
		}),
	}
	NodeTypeSingle := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":   "${data.alicloud_zones.resources.zones.0.id}",
			"node_type": "single",
		}),
	}
	NodeTypeReadone := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":   "${data.alicloud_zones.resources.zones.0.id}",
			"node_type": "readone",
		}),
	}
	NodeTypeReadthree := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":   "${data.alicloud_zones.resources.zones.0.id}",
			"node_type": "readthree",
		}),
	}
	NodeTypeReadfive := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":   "${data.alicloud_zones.resources.zones.0.id}",
			"node_type": "readfive",
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"zone_id":              "${data.alicloud_zones.resources.zones.0.id}",
			"instance_charge_type": "PostPaid",
			"engine":               "Redis",
			"engine_version":       "5.0",
			"architecture":         "standard",
			"performance_type":     "standard_performance_type",
			"storage_type":         "inmemory",
			"node_type":            "double",
			"package_type":         "standard",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"zone_id":              "${data.alicloud_zones.resources.zones.0.id}",
			"instance_charge_type": "PostPaid",
			"engine":               "Fake",
			"engine_version":       "5.6",
			"architecture":         "standard",
			"performance_type":     "standard_performance_type",
			"storage_type":         "inmemory",
			"node_type":            "double",
			"package_type":         "standard",
		}),
	}

	var existKVStoreInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_classes.#": CHECKSET,
			"instance_classes.0": CHECKSET,
		}
	}

	var fakeKVStoreInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_classes.#": "0",
		}
	}

	var KVStoreInstanceCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existKVStoreInstanceMapFunc,
		fakeMapFunc:  fakeKVStoreInstanceMapFunc,
	}

	// At present, the datasource does not support memcache
	//KVStoreInstanceCheckInfo.dataSourceTestCheck(t, rand, EngineVersionConfRedis, EngineVersionConfMemcache,
	//	ChargeTypeConfPostpaid, PerformanceTypeStandardPerformanceType, PerformanceTypeEnhancePerformanceType,
	//	StorageTypeInmemory, PackageTypeStandard, PackageTypeCustomized, ArchitectureStandard, ArchitectureCluster,
	//	ArchitectureRwsplit, NodeTypeDouble, NodeTypeSingle, NodeTypeReadone, NodeTypeReadthree, NodeTypeReadfive,
	//	ArchitectureStandard, allConf)
	KVStoreInstanceCheckInfo.dataSourceTestCheck(t, rand, EngineVersionConfRedis,
		ChargeTypeConfPostpaid, PerformanceTypeStandardPerformanceType, PerformanceTypeEnhancePerformanceType,
		StorageTypeInmemory, PackageTypeStandard, PackageTypeCustomized, ArchitectureStandard, ArchitectureCluster,
		ArchitectureRwsplit, NodeTypeDouble, NodeTypeSingle, NodeTypeReadone, NodeTypeReadthree, NodeTypeReadfive,
		ArchitectureStandard, allConf)
}

func kvstoreConfigHeader(name string) string {
	return fmt.Sprintf(`
data "alicloud_zones" "resources" {
	available_resource_creation= "%s"
}
`, name)
}

// Code generated by "mapstructure-to-hcl2 -type StorageConfig,DiskConfig"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatDiskConfig is an auto-generated flat version of DiskConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDiskConfig struct {
	DiskSize            *int64 `mapstructure:"disk_size" required:"true" cty:"disk_size" hcl:"disk_size"`
	DiskThinProvisioned *bool  `mapstructure:"disk_thin_provisioned" cty:"disk_thin_provisioned" hcl:"disk_thin_provisioned"`
	DiskEagerlyScrub    *bool  `mapstructure:"disk_eagerly_scrub" cty:"disk_eagerly_scrub" hcl:"disk_eagerly_scrub"`
	DiskControllerIndex *int   `mapstructure:"disk_controller_index" cty:"disk_controller_index" hcl:"disk_controller_index"`
}

// FlatMapstructure returns a new FlatDiskConfig.
// FlatDiskConfig is an auto-generated flat version of DiskConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DiskConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDiskConfig)
}

// HCL2Spec returns the hcl spec of a DiskConfig.
// This spec is used by HCL to read the fields of DiskConfig.
// The decoded values from this spec will then be applied to a FlatDiskConfig.
func (*FlatDiskConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"disk_size":             &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_thin_provisioned": &hcldec.AttrSpec{Name: "disk_thin_provisioned", Type: cty.Bool, Required: false},
		"disk_eagerly_scrub":    &hcldec.AttrSpec{Name: "disk_eagerly_scrub", Type: cty.Bool, Required: false},
		"disk_controller_index": &hcldec.AttrSpec{Name: "disk_controller_index", Type: cty.Number, Required: false},
	}
	return s
}

// FlatStorageConfig is an auto-generated flat version of StorageConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatStorageConfig struct {
	DiskControllerType []string         `mapstructure:"disk_controller_type" cty:"disk_controller_type" hcl:"disk_controller_type"`
	Storage            []FlatDiskConfig `mapstructure:"storage" cty:"storage" hcl:"storage"`
}

// FlatMapstructure returns a new FlatStorageConfig.
// FlatStorageConfig is an auto-generated flat version of StorageConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*StorageConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatStorageConfig)
}

// HCL2Spec returns the hcl spec of a StorageConfig.
// This spec is used by HCL to read the fields of StorageConfig.
// The decoded values from this spec will then be applied to a FlatStorageConfig.
func (*FlatStorageConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"disk_controller_type": &hcldec.AttrSpec{Name: "disk_controller_type", Type: cty.List(cty.String), Required: false},
		"storage":              &hcldec.BlockListSpec{TypeName: "storage", Nested: hcldec.ObjectSpec((*FlatDiskConfig)(nil).HCL2Spec())},
	}
	return s
}

// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.

package yandexexport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Endpoint              *string           `mapstructure:"endpoint" required:"false" cty:"endpoint" hcl:"endpoint"`
	ServiceAccountKeyFile *string           `mapstructure:"service_account_key_file" required:"false" cty:"service_account_key_file" hcl:"service_account_key_file"`
	Token                 *string           `mapstructure:"token" required:"true" cty:"token" hcl:"token"`
	MaxRetries            *int              `mapstructure:"max_retries" cty:"max_retries" hcl:"max_retries"`
	SerialLogFile         *string           `mapstructure:"serial_log_file" required:"false" cty:"serial_log_file" hcl:"serial_log_file"`
	StateTimeout          *string           `mapstructure:"state_timeout" required:"false" cty:"state_timeout" hcl:"state_timeout"`
	InstanceCores         *int              `mapstructure:"instance_cores" required:"false" cty:"instance_cores" hcl:"instance_cores"`
	InstanceGpus          *int              `mapstructure:"instance_gpus" required:"false" cty:"instance_gpus" hcl:"instance_gpus"`
	InstanceMemory        *int              `mapstructure:"instance_mem_gb" required:"false" cty:"instance_mem_gb" hcl:"instance_mem_gb"`
	InstanceName          *string           `mapstructure:"instance_name" required:"false" cty:"instance_name" hcl:"instance_name"`
	PlatformID            *string           `mapstructure:"platform_id" required:"false" cty:"platform_id" hcl:"platform_id"`
	Labels                map[string]string `mapstructure:"labels" required:"false" cty:"labels" hcl:"labels"`
	Metadata              map[string]string `mapstructure:"metadata" required:"false" cty:"metadata" hcl:"metadata"`
	MetadataFromFile      map[string]string `mapstructure:"metadata_from_file" cty:"metadata_from_file" hcl:"metadata_from_file"`
	Preemptible           *bool             `mapstructure:"preemptible" cty:"preemptible" hcl:"preemptible"`
	DiskName              *string           `mapstructure:"disk_name" required:"false" cty:"disk_name" hcl:"disk_name"`
	DiskSizeGb            *int              `mapstructure:"disk_size_gb" required:"false" cty:"disk_size_gb" hcl:"disk_size_gb"`
	DiskType              *string           `mapstructure:"disk_type" required:"false" cty:"disk_type" hcl:"disk_type"`
	DiskLabels            map[string]string `mapstructure:"disk_labels" required:"false" cty:"disk_labels" hcl:"disk_labels"`
	SubnetID              *string           `mapstructure:"subnet_id" required:"false" cty:"subnet_id" hcl:"subnet_id"`
	Zone                  *string           `mapstructure:"zone" required:"false" cty:"zone" hcl:"zone"`
	UseIPv4Nat            *bool             `mapstructure:"use_ipv4_nat" required:"false" cty:"use_ipv4_nat" hcl:"use_ipv4_nat"`
	UseIPv6               *bool             `mapstructure:"use_ipv6" required:"false" cty:"use_ipv6" hcl:"use_ipv6"`
	UseInternalIP         *bool             `mapstructure:"use_internal_ip" required:"false" cty:"use_internal_ip" hcl:"use_internal_ip"`
	FolderID              *string           `mapstructure:"folder_id" required:"true" cty:"folder_id" hcl:"folder_id"`
	ServiceAccountID      *string           `mapstructure:"service_account_id" required:"true" cty:"service_account_id" hcl:"service_account_id"`
	Paths                 []string          `mapstructure:"paths" required:"true" cty:"paths" hcl:"paths"`
	SSHPrivateKeyFile     *string           `mapstructure:"ssh_private_key_file" required:"false" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"endpoint":                   &hcldec.AttrSpec{Name: "endpoint", Type: cty.String, Required: false},
		"service_account_key_file":   &hcldec.AttrSpec{Name: "service_account_key_file", Type: cty.String, Required: false},
		"token":                      &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"max_retries":                &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"serial_log_file":            &hcldec.AttrSpec{Name: "serial_log_file", Type: cty.String, Required: false},
		"state_timeout":              &hcldec.AttrSpec{Name: "state_timeout", Type: cty.String, Required: false},
		"instance_cores":             &hcldec.AttrSpec{Name: "instance_cores", Type: cty.Number, Required: false},
		"instance_gpus":              &hcldec.AttrSpec{Name: "instance_gpus", Type: cty.Number, Required: false},
		"instance_mem_gb":            &hcldec.AttrSpec{Name: "instance_mem_gb", Type: cty.Number, Required: false},
		"instance_name":              &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"platform_id":                &hcldec.AttrSpec{Name: "platform_id", Type: cty.String, Required: false},
		"labels":                     &hcldec.AttrSpec{Name: "labels", Type: cty.Map(cty.String), Required: false},
		"metadata":                   &hcldec.AttrSpec{Name: "metadata", Type: cty.Map(cty.String), Required: false},
		"metadata_from_file":         &hcldec.AttrSpec{Name: "metadata_from_file", Type: cty.Map(cty.String), Required: false},
		"preemptible":                &hcldec.AttrSpec{Name: "preemptible", Type: cty.Bool, Required: false},
		"disk_name":                  &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"disk_size_gb":               &hcldec.AttrSpec{Name: "disk_size_gb", Type: cty.Number, Required: false},
		"disk_type":                  &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"disk_labels":                &hcldec.AttrSpec{Name: "disk_labels", Type: cty.Map(cty.String), Required: false},
		"subnet_id":                  &hcldec.AttrSpec{Name: "subnet_id", Type: cty.String, Required: false},
		"zone":                       &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"use_ipv4_nat":               &hcldec.AttrSpec{Name: "use_ipv4_nat", Type: cty.Bool, Required: false},
		"use_ipv6":                   &hcldec.AttrSpec{Name: "use_ipv6", Type: cty.Bool, Required: false},
		"use_internal_ip":            &hcldec.AttrSpec{Name: "use_internal_ip", Type: cty.Bool, Required: false},
		"folder_id":                  &hcldec.AttrSpec{Name: "folder_id", Type: cty.String, Required: false},
		"service_account_id":         &hcldec.AttrSpec{Name: "service_account_id", Type: cty.String, Required: false},
		"paths":                      &hcldec.AttrSpec{Name: "paths", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":       &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
	}
	return s
}

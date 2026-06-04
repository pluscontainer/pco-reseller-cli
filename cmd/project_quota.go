/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/list"
	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
	"github.com/spf13/cobra"
)

var defaultQuota = openapi.UpdateQuota{
	Compute: &openapi.ComputeQuotas{
		Cores:              new(256),
		FloatingIps:        new(60),
		Instances:          500,
		KeyPairs:           500,
		MetadataItems:      100,
		Ram:                new(524288),
		SecurityGroupRules: new(500),
		SecurityGroups:     new(500),
		ServerGroupMembers: new(500),
		ServerGroups:       60,
	},
	Network: &openapi.NetworkQuotas{
		Floatingip:        new(60),
		Network:           60,
		Port:              new(1000),
		Router:            60,
		SecurityGroup:     500,
		SecurityGroupRule: 500,
		Subnet:            500,
	},
	Volume: &openapi.VolumeQuotas{
		BackupGigabytes: 4000,
		Backups:         100,
		Gigabytes:       new(4000),
		Snapshots:       new(100),
		Volumes:         1000,
	},
}

var quotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "Get and update project quotas",
	Long:  `Get and update compute, network and volume quotas for a project`,
}

func printQuota(quota openapi.UpdateQuota) {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)

	l.AppendItem("Quotas:")
	l.Indent()
	l.AppendItem("Compute:")
	l.Indent()
	l.AppendItems([]any{
		fmt.Sprintf("vCPU: %d", *quota.Compute.Cores),
		fmt.Sprintf("RAM: %d MiB", *quota.Compute.Ram),
		fmt.Sprintf("Instances: %d", quota.Compute.Instances),
		fmt.Sprintf("KeyPairs: %d", quota.Compute.KeyPairs),
		fmt.Sprintf("Floating IPs: %d", *quota.Compute.FloatingIps),
		fmt.Sprintf("Security Groups: %d", *quota.Compute.SecurityGroups),
		fmt.Sprintf("Security Group Rules: %d", *quota.Compute.SecurityGroupRules),
		fmt.Sprintf("Server Groups: %d", quota.Compute.ServerGroups),
		fmt.Sprintf("Server Groups Members: %d", *quota.Compute.ServerGroupMembers),
	})
	l.UnIndent()
	l.AppendItem("Volume:")
	l.Indent()
	l.AppendItems([]any{
		fmt.Sprintf("Volumes: %d", quota.Volume.Volumes),
		fmt.Sprintf("Gigabytes: %d", *quota.Volume.Gigabytes),
		fmt.Sprintf("Per Volume Gigabytes: %d", *quota.Volume.PerVolumeGigabytes),
		fmt.Sprintf("Backups: %d", quota.Volume.Backups),
		fmt.Sprintf("Backup Gigabytes: %d", quota.Volume.BackupGigabytes),
		fmt.Sprintf("Snapshots: %d", *quota.Volume.Snapshots),
		fmt.Sprintf("Groups: %d", *quota.Volume.Groups),
	})
	l.UnIndent()
	l.AppendItem("Network:")
	l.Indent()
	l.AppendItems([]any{
		fmt.Sprintf("Networks: %d", quota.Network.Network),
		fmt.Sprintf("Floating IPs: %d", *quota.Network.Floatingip),
		fmt.Sprintf("Ports: %d", *quota.Network.Port),
		fmt.Sprintf("RBAC Policies: %d", *quota.Network.RbacPolicy),
		fmt.Sprintf("Routers: %d", quota.Network.Router),
		fmt.Sprintf("Security Groups: %d", quota.Network.SecurityGroup),
		fmt.Sprintf("Security Group Rules: %d", quota.Network.SecurityGroupRule),
		fmt.Sprintf("Subnets: %d", quota.Network.Subnet),
		fmt.Sprintf("Subnet Pools: %d", *quota.Network.Subnetpool),
	})

	fmt.Println(l.Render())
}

func init() {
	projectCmd.AddCommand(quotaCmd)
}

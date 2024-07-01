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

// quotaCmd represents the quota command
var quotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "Get and update project quotas",
	Long:  `Get and update project quotas`,
}

func init() {
	projectCmd.AddCommand(quotaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quotaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quotaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printQuota(quota openapi.UpdateQuota) {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)

	l.AppendItem("Quotas:")
	l.Indent()
	l.AppendItem("Compute:")
	l.Indent()
	l.AppendItems([]interface{}{
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
	l.AppendItems([]interface{}{
		fmt.Sprintf("Volumes: %d", quota.Volume.Volumes),
		fmt.Sprintf("Gigabyes: %d", *quota.Volume.Gigabytes),
		fmt.Sprintf("Per Volume Gigabytes: %d", *quota.Volume.PerVolumeGigabytes),
		fmt.Sprintf("Backups: %d", quota.Volume.Backups),
		fmt.Sprintf("Backup Gigabytes: %d", quota.Volume.BackupGigabytes),
		fmt.Sprintf("Snapshots: %d", *quota.Volume.Snapshots),
		fmt.Sprintf("Groups: %d", *quota.Volume.Groups),
	})
	l.UnIndent()
	l.AppendItem("Network:")
	l.Indent()
	l.AppendItems([]interface{}{
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

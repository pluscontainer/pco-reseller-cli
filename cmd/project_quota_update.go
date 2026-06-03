/*
Copyright © 2022 PlusServer GmbH
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var quotaCPU, quotaRAM, quotaInstances, quotaKeyPairs, quotaServerGroups, quotaServerGroupMembers, quotaVolumes, quotaGigabytes, quotaBackups, quotaBackupGigabytes, quotaSnapshots, quotaVolumeGroups int
var quotaNetworks, quotaFloatingIPs, quotaPorts, quotaRBAC, quotaRouters, quotaSecurityGroups, quotaSecurityGroupRules, quotaSubnets int

var updateQuotaCmd = &cobra.Command{
	Use:   "update [project-id]",
	Short: "Update the quotas of the specified project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		psOsClient := fetchPsOpenStackClientOrDie()
		ctx := context.Background()
		resp, err := psOsClient.GetProjectQuota(ctx, args[0])
		if err != nil {
			return err
		}

		if quotaCPU != 0 {
			resp.Compute.Cores = &quotaCPU
		}
		if quotaRAM != 0 {
			resp.Compute.Ram = &quotaRAM
		}
		if quotaInstances != 0 {
			resp.Compute.Instances = quotaInstances
		}
		if quotaKeyPairs != 0 {
			resp.Compute.KeyPairs = quotaKeyPairs
		}
		if quotaServerGroups != 0 {
			resp.Compute.ServerGroups = quotaServerGroups
		}
		if quotaServerGroupMembers != 0 {
			resp.Compute.ServerGroupMembers = &quotaServerGroupMembers
		}
		if quotaVolumes != 0 {
			resp.Volume.Volumes = quotaVolumes
		}
		if quotaGigabytes != 0 {
			resp.Volume.Gigabytes = &quotaGigabytes
		}
		if quotaBackups != 0 {
			resp.Volume.Backups = quotaBackups
		}
		if quotaBackupGigabytes != 0 {
			resp.Volume.BackupGigabytes = quotaBackupGigabytes
		}
		if quotaSnapshots != 0 {
			resp.Volume.Snapshots = &quotaSnapshots
		}
		if quotaVolumeGroups != 0 {
			resp.Volume.Groups = &quotaVolumeGroups
		}
		if quotaNetworks != 0 {
			resp.Network.Network = quotaNetworks
		}
		if quotaFloatingIPs != 0 {
			resp.Network.Floatingip = &quotaFloatingIPs
		}
		if quotaPorts != 0 {
			resp.Network.Port = &quotaPorts
		}
		if quotaRBAC != 0 {
			resp.Network.RbacPolicy = &quotaRBAC
		}
		if quotaRouters != 0 {
			resp.Network.Router = quotaRouters
		}
		if quotaSecurityGroups != 0 {
			resp.Network.SecurityGroup = quotaSecurityGroups
		}
		if quotaSecurityGroupRules != 0 {
			resp.Network.SecurityGroupRule = quotaSecurityGroupRules
		}
		if quotaSubnets != 0 {
			resp.Network.Subnet = quotaSubnets
		}

		resp, err = psOsClient.UpdateProjectQuota(ctx, args[0], *resp)
		if err != nil {
			return err
		}

		printQuota(*resp)
		return nil
	},
}

func init() {
	quotaCmd.AddCommand(updateQuotaCmd)
	updateQuotaCmd.Flags().IntVar(&quotaCPU, "cpu", 0, "Update the number of CPUs the project can consume")
	updateQuotaCmd.Flags().IntVar(&quotaRAM, "ram", 0, "Update the amount of RAM the project can consume (in MiB)")
	updateQuotaCmd.Flags().IntVar(&quotaInstances, "instances", 0, "Update the number of instances the project can start")
	updateQuotaCmd.Flags().IntVar(&quotaKeyPairs, "keypair", 0, "Update the number of keypairs the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaServerGroups, "servergroups", 0, "Update the number of server groups the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaServerGroupMembers, "servergroupmembers", 0, "Update the number of members a server group can contain")
	updateQuotaCmd.Flags().IntVar(&quotaVolumes, "volumes", 0, "Update the number of volumes the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaGigabytes, "gigabytes", 0, "Update the gigabytes the project can allocate")
	updateQuotaCmd.Flags().IntVar(&quotaBackups, "backups", 0, "Update the number of backups the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaBackupGigabytes, "backup-gigabytes", 0, "Update the amount of storage backups can allocate (in GiB)")
	updateQuotaCmd.Flags().IntVar(&quotaSnapshots, "snapshots", 0, "Update the number of snapshots the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaVolumeGroups, "volumegroups", 0, "Update the number of volume groups the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaNetworks, "networks", 0, "Update the number of networks the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaFloatingIPs, "floatingips", 0, "Update the number of floating IPs the project can allocate")
	updateQuotaCmd.Flags().IntVar(&quotaPorts, "ports", 0, "Update the number of ports the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaRBAC, "rbac", 0, "Update the number of RBAC policies the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaRouters, "routers", 0, "Update the number of routers the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaSecurityGroups, "securitygroups", 0, "Update the number of security groups the project can create")
	updateQuotaCmd.Flags().IntVar(&quotaSecurityGroupRules, "securitygrouprules", 0, "Update the number of rules a security group can contain")
	updateQuotaCmd.Flags().IntVar(&quotaSubnets, "subnets", 0, "Update the number of subnets the project can create")
}

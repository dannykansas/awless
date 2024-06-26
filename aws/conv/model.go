/*
/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsconv

import "github.com/thunderbird86/awless/cloud"
import "github.com/thunderbird86/awless/cloud/properties"

var awsResourcesDef = map[string]map[string]*propertyTransform{
	//EC2
	cloud.Instance: {
		properties.Name:              {name: "Tags", transform: extractTagFn("Name")},
		properties.Type:              {name: "InstanceType", transform: extractValueFn},
		properties.Subnet:            {name: "SubnetId", transform: extractValueFn},
		properties.Vpc:               {name: "VpcId", transform: extractValueFn},
		properties.PublicIP:          {name: "PublicIpAddress", transform: extractValueFn},
		properties.PrivateIP:         {name: "PrivateIpAddress", transform: extractValueFn},
		properties.Image:             {name: "ImageId", transform: extractValueFn},
		properties.Launched:          {name: "LaunchTime", transform: extractValueFn},
		properties.State:             {name: "State", transform: extractFieldFn("Name")},
		properties.KeyPair:           {name: "KeyName", transform: extractValueFn},
		properties.SecurityGroups:    {name: "SecurityGroups", transform: extractStringSliceValues("GroupId")},
		properties.Affinity:          {name: "Placement", transform: extractFieldFn("Affinity")},
		properties.AvailabilityZone:  {name: "Placement", transform: extractFieldFn("AvailabilityZone")},
		properties.PlacementGroup:    {name: "Placement", transform: extractFieldFn("GroupName")},
		properties.Host:              {name: "Placement", transform: extractFieldFn("HostId")},
		properties.Architecture:      {name: "Architecture", transform: extractValueFn},
		properties.Hypervisor:        {name: "Hypervisor", transform: extractValueFn},
		properties.Profile:           {name: "IamInstanceProfile", transform: extractFieldFn("Arn")},
		properties.Lifecycle:         {name: "InstanceLifecycle", transform: extractValueFn},
		properties.NetworkInterfaces: {name: "NetworkInterfaces", transform: extractStringSliceValues("NetworkInterfaceId")},
		properties.PublicDNS:         {name: "PublicDnsName", transform: extractValueFn},
		properties.RootDevice:        {name: "RootDeviceName", transform: extractValueFn},
		properties.RootDeviceType:    {name: "RootDeviceType", transform: extractValueFn},
		properties.Tags:              {name: "Tags", transform: extractTagsFn},
	},
	cloud.Vpc: {
		properties.Name:    {name: "Tags", transform: extractTagFn("Name")},
		properties.Default: {name: "IsDefault", transform: extractValueFn},
		properties.State:   {name: "State", transform: extractValueFn},
		properties.CIDR:    {name: "CidrBlock", transform: extractValueFn},
		properties.Tags:    {name: "Tags", transform: extractTagsFn},
	},
	cloud.Subnet: {
		properties.Name:             {name: "Tags", transform: extractTagFn("Name")},
		properties.Vpc:              {name: "VpcId", transform: extractValueFn},
		properties.Public:           {name: "MapPublicIpOnLaunch", transform: extractValueFn},
		properties.State:            {name: "State", transform: extractValueFn},
		properties.CIDR:             {name: "CidrBlock", transform: extractValueFn},
		properties.AvailabilityZone: {name: "AvailabilityZone", transform: extractValueFn},
		properties.Default:          {name: "DefaultForAz", transform: extractValueFn},
		properties.Tags:             {name: "Tags", transform: extractTagsFn},
	},
	cloud.SecurityGroup: {
		properties.Name:          {name: "GroupName", transform: extractValueFn},
		properties.Description:   {name: "Description", transform: extractValueFn},
		properties.InboundRules:  {name: "IpPermissions", transform: extractIpPermissionSliceFn},
		properties.OutboundRules: {name: "IpPermissionsEgress", transform: extractIpPermissionSliceFn},
		properties.Owner:         {name: "OwnerId", transform: extractValueFn},
		properties.Vpc:           {name: "VpcId", transform: extractValueFn},
		properties.Tags:          {name: "Tags", transform: extractTagsFn},
	},
	cloud.Keypair: {
		properties.Fingerprint: {name: "KeyFingerprint", transform: extractValueFn},
	},
	cloud.Volume: {
		properties.Name:             {name: "Tags", transform: extractTagFn("Name")},
		properties.Type:             {name: "VolumeType", transform: extractValueFn},
		properties.State:            {name: "State", transform: extractValueFn},
		properties.Size:             {name: "Size", transform: extractValueFn},
		properties.Encrypted:        {name: "Encrypted", transform: extractValueFn},
		properties.Created:          {name: "CreateTime", transform: extractTimeFn},
		properties.AvailabilityZone: {name: "AvailabilityZone", transform: extractValueFn},
		properties.Instances:        {name: "Attachments", transform: extractStringSliceValues("InstanceId")},
		properties.Tags:             {name: "Tags", transform: extractTagsFn},
	},
	cloud.Snapshot: {
		properties.Description: {name: "Description", transform: extractValueFn},
		properties.Encrypted:   {name: "Encrypted", transform: extractValueFn},
		properties.Owner:       {name: "OwnerId", transform: extractValueFn},
		properties.Progress:    {name: "Progress", transform: extractValueFn},
		properties.Created:     {name: "StartTime", transform: extractValueFn},
		properties.State:       {name: "State", transform: extractValueFn},
		properties.Size:        {name: "VolumeSize", transform: extractValueFn},
		properties.Volume:      {name: "VolumeId", transform: extractValueFn},
		properties.Tags:        {name: "Tags", transform: extractTagsFn},
	},
	cloud.Image: {
		properties.Name:           {name: "Name", transform: extractValueFn},
		properties.Architecture:   {name: "Architecture", transform: extractValueFn},
		properties.Hypervisor:     {name: "Hypervisor", transform: extractValueFn},
		properties.Location:       {name: "ImageLocation", transform: extractValueFn},
		properties.Description:    {name: "Description", transform: extractValueFn},
		properties.Type:           {name: "ImageType", transform: extractValueFn},
		properties.Public:         {name: "Public", transform: extractValueFn},
		properties.State:          {name: "State", transform: extractValueFn},
		properties.Created:        {name: "CreationDate", transform: extractTimeWithZSuffixFn},
		properties.Virtualization: {name: "VirtualizationType", transform: extractValueFn},
		properties.Tags:           {name: "Tags", transform: extractTagsFn},
	},
	cloud.InstanceProfile: {
		properties.Name:    {name: "InstanceProfileName", transform: extractValueFn},
		properties.Arn:     {name: "Arn", transform: extractValueFn},
		properties.Created: {name: "CreateDate", transform: extractValueFn},
		properties.Path:    {name: "Path", transform: extractValueFn},
		properties.Roles:   {name: "Roles", transform: extractStringSliceValues("RoleId")},
	},
	cloud.ImportImageTask: {
		properties.Architecture: {name: "Architecture", transform: extractValueFn},
		properties.Description:  {name: "Description", transform: extractValueFn},
		properties.Hypervisor:   {name: "Hypervisor", transform: extractValueFn},
		properties.Image:        {name: "ImageId", transform: extractValueFn},
		properties.Progress:     {name: "Progress", transform: extractValueFn},
		properties.State:        {name: "Status", transform: extractValueFn},
		properties.StateMessage: {name: "StatusMessage", transform: extractValueFn},
	},
	cloud.InternetGateway: {
		properties.Name: {name: "Tags", transform: extractTagFn("Name")},
		properties.Vpcs: {name: "Attachments", transform: extractStringSliceValues("VpcId")},
		properties.Tags: {name: "Tags", transform: extractTagsFn},
	},
	cloud.NatGateway: {
		properties.Created: {name: "CreateTime", transform: extractValueFn},
		properties.Subnet:  {name: "SubnetId", transform: extractValueFn},
		properties.Vpc:     {name: "VpcId", transform: extractValueFn},
		properties.State:   {name: "State", transform: extractValueFn},
	},
	cloud.RouteTable: {
		properties.Name:         {name: "Tags", transform: extractTagFn("Name")},
		properties.Vpc:          {name: "VpcId", transform: extractValueFn},
		properties.Routes:       {name: "Routes", transform: extractRoutesSliceFn},
		properties.Default:      {name: "Associations", transform: extractHasATrueBoolInStructSliceFn("Main")},
		properties.Associations: {name: "Associations", transform: extractRouteTableAssociationsFn},
		properties.Tags:         {name: "Tags", transform: extractTagsFn},
	},
	cloud.AvailabilityZone: {
		properties.Name:     {name: "ZoneName", transform: extractValueFn},
		properties.State:    {name: "State", transform: extractValueFn},
		properties.Region:   {name: "RegionName", transform: extractValueFn},
		properties.Messages: {name: "Messages", transform: extractStringSliceValues("Message")},
	},
	cloud.ElasticIP: {
		properties.Name:        {name: "PublicIp", transform: extractValueFn},
		properties.PublicIP:    {name: "PublicIp", transform: extractValueFn},
		properties.PrivateIP:   {name: "PrivateIpAddress", transform: extractValueFn},
		properties.Association: {name: "AssociationId", transform: extractValueFn},
	},
	cloud.NetworkInterface: {
		properties.PublicIP:         {name: "Association", transform: extractFieldFn("PublicIp")},
		properties.PublicDNS:        {name: "Association", transform: extractFieldFn("PublicDnsName")},
		properties.Attachment:       {name: "Attachment", transform: extractFieldFn("AttachmentId")},
		properties.Instance:         {name: "Attachment", transform: extractFieldFn("InstanceId")},
		properties.InstanceOwner:    {name: "Attachment", transform: extractFieldFn("InstanceOwnerId")},
		properties.AvailabilityZone: {name: "AvailabilityZone", transform: extractValueFn},
		properties.Description:      {name: "Description", transform: extractValueFn},
		properties.SecurityGroups:   {name: "Groups", transform: extractStringSliceValues("GroupId")},
		properties.Type:             {name: "InterfaceType", transform: extractValueFn},
		properties.IPv6Addresses:    {name: "Ipv6Addresses", transform: extractStringSliceValues("Ipv6Address")},
		properties.MACAddress:       {name: "MacAddress", transform: extractValueFn},
		properties.Owner:            {name: "OwnerId", transform: extractValueFn},
		properties.PrivateDNS:       {name: "PrivateDnsName", transform: extractValueFn},
		properties.PrivateIP:        {name: "PrivateIpAddress", transform: extractValueFn},
		properties.State:            {name: "Status", transform: extractValueFn},
		properties.Subnet:           {name: "SubnetId", transform: extractValueFn},
		properties.Vpc:              {name: "VpcId", transform: extractValueFn},
		properties.Tags:             {name: "TagSet", transform: extractTagsFn},
	},
	// LoadBalancer
	cloud.LoadBalancer: {
		properties.Name:              {name: "LoadBalancerName", transform: extractValueFn},
		properties.Arn:               {name: "LoadBalancerArn", transform: extractValueFn},
		properties.AvailabilityZones: {name: "AvailabilityZones", transform: extractStringSliceValues("ZoneName")},
		properties.Subnets:           {name: "AvailabilityZones", transform: extractStringSliceValues("SubnetId")},
		properties.Zone:              {name: "CanonicalHostedZoneId", transform: extractValueFn},
		properties.Created:           {name: "CreatedTime", transform: extractTimeFn},
		properties.PublicDNS:         {name: "DNSName", transform: extractValueFn},
		properties.IPType:            {name: "IpAddressType", transform: extractValueFn},
		properties.Scheme:            {name: "Scheme", transform: extractValueFn},
		properties.State:             {name: "State", transform: extractFieldFn("Code")},
		properties.Type:              {name: "Type", transform: extractValueFn},
		properties.Vpc:               {name: "VpcId", transform: extractValueFn},
	},

	cloud.ClassicLoadBalancer: {
		properties.Name:              {name: "LoadBalancerName", transform: extractValueFn},
		properties.AvailabilityZones: {name: "AvailabilityZones", transform: extractValueFn},
		properties.Subnets:           {name: "Subnets", transform: extractValueFn},
		properties.Instances:         {name: "Instances", transform: extractStringSliceValues("InstanceId")},
		properties.Ports:             {name: "ListenerDescriptions", transform: extractClassicLoadbListenerDescriptionsFn},
		properties.Zone:              {name: "CanonicalHostedZoneNameID", transform: extractValueFn},
		properties.Created:           {name: "CreatedTime", transform: extractTimeFn},
		properties.PublicDNS:         {name: "DNSName", transform: extractValueFn},
		properties.Scheme:            {name: "Scheme", transform: extractValueFn},
		properties.Vpc:               {name: "VPCId", transform: extractValueFn},
	},

	cloud.TargetGroup: {
		properties.Name:                    {name: "TargetGroupName", transform: extractValueFn},
		properties.Arn:                     {name: "TargetGroupArn", transform: extractValueFn},
		properties.CheckInterval:           {name: "HealthCheckIntervalSeconds", transform: extractValueFn},
		properties.CheckPath:               {name: "HealthCheckPath", transform: extractValueFn},
		properties.CheckPort:               {name: "HealthCheckPort", transform: extractValueFn},
		properties.CheckProtocol:           {name: "HealthCheckProtocol", transform: extractValueFn},
		properties.CheckTimeout:            {name: "HealthCheckTimeoutSeconds", transform: extractValueFn},
		properties.CheckHTTPCode:           {name: "Matcher", transform: extractFieldFn("HttpCode")},
		properties.HealthyThresholdCount:   {name: "HealthyThresholdCount", transform: extractValueFn},
		properties.UnhealthyThresholdCount: {name: "UnhealthyThresholdCount", transform: extractValueFn},
		properties.Port:                    {name: "Port", transform: extractValueFn},
		properties.Protocol:                {name: "Protocol", transform: extractValueFn},
		properties.Vpc:                     {name: "VpcId", transform: extractValueFn},
	},
	cloud.Listener: {
		properties.Arn:          {name: "ListenerArn", transform: extractValueFn},
		properties.Certificates: {name: "Certificates", transform: extractStringSliceValues("CertificateArn")},
		properties.Actions:      {name: "DefaultActions", transform: extractStringSliceValues("Type")},
		properties.TargetGroups: {name: "DefaultActions", transform: extractStringSliceValues("TargetGroupArn")},
		properties.LoadBalancer: {name: "LoadBalancerArn", transform: extractValueFn},
		properties.Port:         {name: "Port", transform: extractValueFn},
		properties.Protocol:     {name: "Protocol", transform: extractValueFn},
		properties.CipherSuite:  {name: "SslPolicy", transform: extractValueFn},
	},
	//Database
	cloud.Database: {
		properties.Storage:                   {name: "AllocatedStorage", transform: extractValueFn},
		properties.AutoUpgrade:               {name: "AutoMinorVersionUpgrade", transform: extractValueFn},
		properties.AvailabilityZone:          {name: "AvailabilityZone", transform: extractValueFn},
		properties.BackupRetentionPeriod:     {name: "BackupRetentionPeriod", transform: extractValueFn},
		properties.CertificateAuthority:      {name: "CACertificateIdentifier", transform: extractValueFn},
		properties.Charset:                   {name: "CharacterSetName", transform: extractValueFn},
		properties.CopyTagsToSnapshot:        {name: "CopyTagsToSnapshot", transform: extractValueFn},
		properties.Cluster:                   {name: "DBClusterIdentifier", transform: extractValueFn},
		properties.Arn:                       {name: "DBInstanceArn", transform: extractValueFn},
		properties.Class:                     {name: "DBInstanceClass", transform: extractValueFn},
		properties.State:                     {name: "DBInstanceStatus", transform: extractValueFn},
		properties.Name:                      {name: "DBName", transform: extractValueFn},
		properties.ParameterGroups:           {name: "DBParameterGroups", transform: extractStringSliceValues("DBParameterGroupName")},
		properties.DBSecurityGroups:          {name: "DBSecurityGroups", transform: extractStringSliceValues("DBSecurityGroupName")},
		properties.DBSubnetGroup:             {name: "DBSubnetGroup", transform: extractFieldFn("DBSubnetGroupName")},
		properties.Port:                      {name: "DbInstancePort", transform: extractValueFn},
		properties.GlobalID:                  {name: "DbiResourceId", transform: extractValueFn},
		properties.PublicDNS:                 {name: "Endpoint", transform: extractFieldFn("Address")},
		properties.Zone:                      {name: "Endpoint", transform: extractFieldFn("HostedZoneId")},
		properties.Engine:                    {name: "Engine", transform: extractValueFn},
		properties.EngineVersion:             {name: "EngineVersion", transform: extractValueFn},
		properties.Created:                   {name: "InstanceCreateTime", transform: extractValueFn},
		properties.IOPS:                      {name: "Iops", transform: extractValueFn},
		properties.LatestRestorableTime:      {name: "LatestRestorableTime", transform: extractValueFn},
		properties.License:                   {name: "LicenseModel", transform: extractValueFn},
		properties.Username:                  {name: "MasterUsername", transform: extractValueFn},
		properties.MonitoringInterval:        {name: "MonitoringInterval", transform: extractValueFn},
		properties.MonitoringRole:            {name: "MonitoringRoleArn", transform: extractValueFn},
		properties.MultiAZ:                   {name: "MultiAZ", transform: extractValueFn},
		properties.OptionGroups:              {name: "OptionGroupMemberships", transform: extractStringSliceValues("OptionGroupName")},
		properties.PreferredBackupDate:       {name: "PreferredBackupWindow", transform: extractValueFn},
		properties.PreferredMaintenanceDate:  {name: "PreferredMaintenanceWindow", transform: extractValueFn},
		properties.Public:                    {name: "PubliclyAccessible", transform: extractValueFn},
		properties.SecondaryAvailabilityZone: {name: "SecondaryAvailabilityZone", transform: extractValueFn},
		properties.Encrypted:                 {name: "StorageEncrypted", transform: extractValueFn},
		properties.StorageType:               {name: "StorageType", transform: extractValueFn},
		properties.Timezone:                  {name: "Timezone", transform: extractValueFn},
		properties.SecurityGroups:            {name: "VpcSecurityGroups", transform: extractStringSliceValues("VpcSecurityGroupId")},
		properties.ReplicaOf:                 {name: "ReadReplicaSourceDBInstanceIdentifier", transform: extractValueFn},
	},
	cloud.DbSubnetGroup: {
		properties.Name:        {name: "DBSubnetGroupName", transform: extractValueFn},
		properties.Arn:         {name: "DBSubnetGroupArn", transform: extractValueFn},
		properties.Description: {name: "DBSubnetGroupDescription", transform: extractValueFn},
		properties.State:       {name: "SubnetGroupStatus", transform: extractValueFn},
		properties.Subnets:     {name: "Subnets", transform: extractStringSliceValues("SubnetIdentifier")},
		properties.Vpc:         {name: "VpcId", transform: extractValueFn},
	},
	//Autoscaling
	cloud.LaunchConfiguration: {
		properties.Name:           {name: "LaunchConfigurationName", transform: extractValueFn},
		properties.Arn:            {name: "LaunchConfigurationARN", transform: extractValueFn},
		properties.Public:         {name: "AssociatePublicIpAddress", transform: extractValueFn},
		properties.Created:        {name: "CreatedTime", transform: extractValueFn},
		properties.Profile:        {name: "IamInstanceProfile", transform: extractValueFn},
		properties.Image:          {name: "ImageId", transform: extractValueFn},
		properties.Type:           {name: "InstanceType", transform: extractValueFn},
		properties.KeyPair:        {name: "KeyName", transform: extractValueFn},
		properties.SecurityGroups: {name: "SecurityGroups", transform: extractStringPointerSliceValues},
		properties.SpotPrice:      {name: "SpotPrice", transform: extractValueFn},
		properties.UserData:       {name: "Userdata", transform: extractValueFn},
	},
	cloud.ScalingGroup: {
		properties.Name:                    {name: "AutoScalingGroupName", transform: extractValueFn},
		properties.Arn:                     {name: "AutoScalingGroupARN", transform: extractValueFn},
		properties.Created:                 {name: "CreatedTime", transform: extractValueFn},
		properties.DefaultCooldown:         {name: "DefaultCooldown", transform: extractValueFn},
		properties.DesiredCapacity:         {name: "DesiredCapacity", transform: extractValueFn},
		properties.HealthCheckGracePeriod:  {name: "HealthCheckGracePeriod", transform: extractValueFn},
		properties.HealthCheckType:         {name: "HealthCheckType", transform: extractValueFn},
		properties.LaunchConfigurationName: {name: "LaunchConfigurationName", transform: extractValueFn},
		properties.MaxSize:                 {name: "MaxSize", transform: extractValueFn},
		properties.MinSize:                 {name: "MinSize", transform: extractValueFn},
		properties.NewInstancesProtected:   {name: "NewInstancesProtectedFromScaleIn", transform: extractValueFn},
		properties.State:                   {name: "Status", transform: extractValueFn},
		properties.Tags:                    {name: "Tags", transform: extractTagsFn},
	},
	cloud.ScalingPolicy: {
		properties.Name:              {name: "PolicyName", transform: extractValueFn},
		properties.Arn:               {name: "PolicyARN", transform: extractValueFn},
		properties.AdjustmentType:    {name: "AdjustmentType", transform: extractValueFn},
		properties.AlarmNames:        {name: "Alarms", transform: extractStringSliceValues("AlarmName")},
		properties.ScalingGroupName:  {name: "AutoScalingGroupName", transform: extractValueFn},
		properties.Cooldown:          {name: "AutoScalingGroupName", transform: extractValueFn},
		properties.Type:              {name: "PolicyType", transform: extractValueFn},
		properties.ScalingAdjustment: {name: "ScalingAdjustment", transform: extractValueFn},
	},
	//Containers
	cloud.Repository: {
		properties.Name:    {name: "RepositoryName", transform: extractValueFn},
		properties.Arn:     {name: "RepositoryArn", transform: extractValueFn},
		properties.URI:     {name: "RepositoryUri", transform: extractValueFn},
		properties.Created: {name: "CreatedAt", transform: extractValueFn},
		properties.Account: {name: "RegistryId", transform: extractValueFn},
	},
	cloud.ContainerCluster: {
		properties.Name:                              {name: "ClusterName", transform: extractValueFn},
		properties.Arn:                               {name: "ClusterArn", transform: extractValueFn},
		properties.ActiveServicesCount:               {name: "ActiveServicesCount", transform: extractValueFn},
		properties.PendingTasksCount:                 {name: "PendingTasksCount", transform: extractValueFn},
		properties.RegisteredContainerInstancesCount: {name: "RegisteredContainerInstancesCount", transform: extractValueFn},
		properties.RunningTasksCount:                 {name: "RunningTasksCount", transform: extractValueFn},
		properties.State:                             {name: "Status", transform: extractValueFn},
	},
	cloud.ContainerTask: {
		properties.Name:             {name: "Family", transform: extractValueFn},
		properties.Arn:              {name: "TaskDefinitionArn", transform: extractValueFn},
		properties.ContainersImages: {name: "ContainerDefinitions", transform: extractContainersImagesFn},
		properties.Version:          {name: "Revision", transform: extractValueAsStringFn},
		properties.Role:             {name: "TaskRoleArn", transform: extractValueFn},
	},
	cloud.Container: {
		properties.Name:         {name: "Name", transform: extractValueFn},
		properties.Arn:          {name: "ContainerArn", transform: extractValueFn},
		properties.ExitCode:     {name: "ExitCode", transform: extractValueFn},
		properties.State:        {name: "LastStatus", transform: extractValueFn},
		properties.StateMessage: {name: "Reason", transform: extractValueFn},
	},
	cloud.ContainerInstance: {
		properties.Arn:               {name: "ContainerInstanceArn", transform: extractValueFn},
		properties.AgentConnected:    {name: "AgentConnected", transform: extractValueFn},
		properties.AgentState:        {name: "AgentUpdateStatus", transform: extractValueFn},
		properties.Attributes:        {name: "Attributes", transform: extractECSAttributesFn},
		properties.Instance:          {name: "Ec2InstanceId", transform: extractValueFn},
		properties.PendingTasksCount: {name: "PendingTasksCount", transform: extractValueFn},
		properties.RunningTasksCount: {name: "RunningTasksCount", transform: extractValueFn},
		properties.Created:           {name: "RegisteredAt", transform: extractValueFn},
		properties.State:             {name: "Status", transform: extractValueFn},
		properties.Version:           {name: "Version", transform: extractValueAsStringFn},
		properties.AgentVersion:      {name: "VersionInfo", transform: extractFieldFn("AgentVersion")},
		properties.DockerVersion:     {name: "VersionInfo", transform: extractFieldFn("DockerVersion")},
	},
	//ACM
	cloud.Certificate: {
		properties.Arn:  {name: "CertificateArn", transform: extractValueFn},
		properties.Name: {name: "DomainName", transform: extractValueFn},
	},
	//IAM
	cloud.User: {
		properties.Name:             {name: "UserName", transform: extractValueFn},
		properties.Arn:              {name: "Arn", transform: extractValueFn},
		properties.Path:             {name: "Path", transform: extractValueFn},
		properties.Created:          {name: "CreateDate", transform: extractTimeFn},
		properties.PasswordLastUsed: {name: "PasswordLastUsed", transform: extractTimeFn},
		properties.InlinePolicies:   {name: "UserPolicyList", transform: extractStringSliceValues("PolicyName")},
	},
	cloud.Role: {
		properties.Name:           {name: "RoleName", transform: extractValueFn},
		properties.Arn:            {name: "Arn", transform: extractValueFn},
		properties.Created:        {name: "CreateDate", transform: extractTimeFn},
		properties.Path:           {name: "Path", transform: extractValueFn},
		properties.InlinePolicies: {name: "RolePolicyList", transform: extractStringSliceValues("PolicyName")},
		properties.TrustPolicy:    {name: "AssumeRolePolicyDocument", transform: extractURLEncodedJson},
	},
	cloud.Group: {
		properties.Name:           {name: "GroupName", transform: extractValueFn},
		properties.Arn:            {name: "Arn", transform: extractValueFn},
		properties.Created:        {name: "CreateDate", transform: extractTimeFn},
		properties.Path:           {name: "Path", transform: extractValueFn},
		properties.InlinePolicies: {name: "GroupPolicyList", transform: extractStringSliceValues("PolicyName")},
	},
	cloud.Policy: {
		properties.Name:        {name: "PolicyName", transform: extractValueFn},
		properties.Arn:         {name: "Arn", transform: extractValueFn},
		properties.Created:     {name: "CreateDate", transform: extractTimeFn},
		properties.Updated:     {name: "UpdateDate", transform: extractTimeFn},
		properties.Description: {name: "Description", transform: extractValueFn},
		properties.Attachable:  {name: "IsAttachable", transform: extractValueFn},
		properties.Path:        {name: "Path", transform: extractValueFn},
		properties.Document:    {name: "PolicyVersionList", transform: extractDocumentDefaultVersion},
	},
	cloud.AccessKey: {
		properties.Username: {name: "UserName", transform: extractValueFn},
		properties.State:    {name: "Status", transform: extractValueFn},
		properties.Created:  {name: "CreateDate", transform: extractTimeFn},
	},
	cloud.MFADevice: {
		properties.AttachedAt: {name: "EnableDate", transform: extractTimeFn},
	},
	//S3
	cloud.Bucket: {
		properties.Created: {name: "CreationDate", transform: extractTimeFn},
	},
	cloud.S3Object: {
		properties.Key:      {name: "Key", transform: extractValueFn},
		properties.Modified: {name: "LastModified", transform: extractTimeFn},
		properties.Owner:    {name: "Owner", transform: extractFieldFn("ID")},
		properties.Size:     {name: "Size", transform: extractValueFn},
		properties.Class:    {name: "StorageClass", transform: extractValueFn},
	},
	//Notification
	cloud.Subscription: {
		properties.Endpoint: {name: "Endpoint", transform: extractValueFn},
		properties.Owner:    {name: "Owner", transform: extractValueFn},
		properties.Protocol: {name: "Protocol", transform: extractValueFn},
		properties.Arn:      {name: "SubscriptionArn", transform: extractValueFn},
		properties.Topic:    {name: "TopicArn", transform: extractValueFn},
	},
	cloud.Topic: {
		properties.Arn: {name: "TopicArn", transform: extractValueFn},
	},
	// DNS
	cloud.Zone: {
		properties.Name:            {name: "Name", transform: extractValueFn},
		properties.Comment:         {name: "Config", transform: extractFieldFn("Comment")},
		properties.Private:         {name: "Config", transform: extractFieldFn("PrivateZone")},
		properties.CallerReference: {name: "CallerReference", transform: extractValueFn},
		properties.RecordCount:     {name: "ResourceRecordSetCount", transform: extractValueFn},
	},
	cloud.Record: {
		properties.Name:                  {name: "Name", transform: extractValueFn},
		properties.Zone:                  {name: "Zone"},
		properties.Failover:              {name: "Failover", transform: extractValueFn},
		properties.Continent:             {name: "GeoLocation", transform: extractFieldFn("ContinentCode")},
		properties.Country:               {name: "GeoLocation", transform: extractFieldFn("CountryCode")},
		properties.HealthCheck:           {name: "HealthCheckId", transform: extractValueFn},
		properties.Region:                {name: "Region", transform: extractValueFn},
		properties.Records:               {name: "ResourceRecords", transform: extractStringSliceValues("Value")},
		properties.Alias:                 {name: "AliasTarget", transform: extractFieldFn("DNSName")},
		properties.Set:                   {name: "SetIdentifier", transform: extractValueFn},
		properties.TTL:                   {name: "TTL", transform: extractValueFn},
		properties.TrafficPolicyInstance: {name: "TrafficPolicyInstanceId", transform: extractValueFn},
		properties.Type:                  {name: "Type", transform: extractValueFn},
		properties.Weight:                {name: "Weight", transform: extractValueFn},
	},
	// Lambda
	cloud.Function: {
		properties.Arn:         {name: "FunctionArn", transform: extractValueFn},
		properties.Name:        {name: "FunctionName", transform: extractValueFn},
		properties.Hash:        {name: "CodeSha256", transform: extractValueFn},
		properties.Size:        {name: "CodeSize", transform: extractValueFn},
		properties.Description: {name: "Description", transform: extractValueFn},
		properties.Handler:     {name: "Handler", transform: extractValueFn},
		properties.Modified:    {name: "LastModified", transform: extractTimeFn},
		properties.Memory:      {name: "MemorySize", transform: extractValueFn},
		properties.Role:        {name: "Role", transform: extractValueFn},
		properties.Runtime:     {name: "Runtime", transform: extractValueFn},
		properties.Timeout:     {name: "Timeout", transform: extractValueFn},
		properties.Version:     {name: "Version", transform: extractValueFn},
	},
	// Monitoring
	cloud.Metric: {
		properties.Name:       {name: "MetricName", transform: extractValueFn},
		properties.Namespace:  {name: "Namespace", transform: extractValueFn},
		properties.Dimensions: {name: "Dimensions", transform: extractNameValueFn},
	},
	cloud.Alarm: {
		properties.Name:                    {name: "AlarmName", transform: extractValueFn},
		properties.ActionsEnabled:          {name: "ActionsEnabled", transform: extractValueFn},
		properties.AlarmActions:            {name: "AlarmActions", transform: extractStringPointerSliceValues},
		properties.InsufficientDataActions: {name: "InsufficientDataActions", transform: extractStringPointerSliceValues},
		properties.OKActions:               {name: "OKActions", transform: extractStringPointerSliceValues},
		properties.Arn:                     {name: "AlarmArn", transform: extractValueFn},
		properties.Description:             {name: "AlarmDescription", transform: extractValueFn},
		properties.Dimensions:              {name: "Dimensions", transform: extractNameValueFn},
		properties.MetricName:              {name: "MetricName", transform: extractValueFn},
		properties.Namespace:               {name: "Namespace", transform: extractValueFn},
		properties.Updated:                 {name: "StateUpdatedTimestamp", transform: extractValueFn},
		properties.State:                   {name: "StateValue", transform: extractValueFn},
	},
	// CDN
	cloud.Distribution: {
		properties.Arn:                {name: "ARN", transform: extractValueFn},
		properties.Aliases:            {name: "Aliases", transform: extractFieldFn("Items")},
		properties.Comment:            {name: "Comment", transform: extractValueFn},
		properties.ACMCertificate:     {name: "ViewerCertificate", transform: extractFieldFn("ACMCertificateArn")},
		properties.Certificate:        {name: "ViewerCertificate", transform: extractFieldFn("Certificate")},
		properties.TLSVersionRequired: {name: "ViewerCertificate", transform: extractFieldFn("MinimumProtocolVersion")},
		properties.SSLSupportMethod:   {name: "ViewerCertificate", transform: extractFieldFn("SSLSupportMethod")},
		properties.Origins:            {name: "Origins", transform: extractDistributionOriginFn},
		properties.PublicDNS:          {name: "DomainName", transform: extractValueFn},
		properties.Enabled:            {name: "Enabled", transform: extractValueFn},
		properties.HTTPVersion:        {name: "HttpVersion", transform: extractValueFn},
		properties.IPv6Enabled:        {name: "IsIPV6Enabled", transform: extractValueFn},
		properties.Modified:           {name: "LastModifiedTime", transform: extractValueFn},
		properties.PriceClass:         {name: "PriceClass", transform: extractValueFn},
		properties.State:              {name: "Status", transform: extractValueFn},
		properties.WebACL:             {name: "WebACLId", transform: extractValueFn},
	},
	// Cloudformation
	cloud.Stack: {
		properties.Name:            {name: "StackName", transform: extractValueFn},
		properties.Capabilities:    {name: "Capabilities", transform: extractStringPointerSliceValues},
		properties.ChangeSet:       {name: "ChangeSetId", transform: extractValueFn},
		properties.Created:         {name: "CreationTime", transform: extractValueFn},
		properties.Description:     {name: "Description", transform: extractValueFn},
		properties.DisableRollback: {name: "DisableRollback", transform: extractValueFn},
		properties.Modified:        {name: "LastUpdatedTime", transform: extractValueFn},
		properties.Notifications:   {name: "NotificationARNs", transform: extractStringPointerSliceValues},
		properties.Role:            {name: "RoleARN", transform: extractValueFn},
		properties.State:           {name: "StackStatus", transform: extractValueFn},
		properties.StateMessage:    {name: "StackStatusReason", transform: extractValueFn},
		properties.Parameters:      {name: "Parameters", transform: extractStackParametersFn},
		properties.Outputs:         {name: "Outputs", transform: extractStackOutputsFn},
	},
	//Queue
	cloud.Queue: {}, //Manually set
}

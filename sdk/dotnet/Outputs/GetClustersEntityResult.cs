// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace JacobBMay.Nutanix.Outputs
{

    [OutputType]
    public sealed class GetClustersEntityResult
    {
        /// <summary>
        /// - Map of cluster efficiency which includes numbers of inefficient vms. The value is populated by analytics on PC. (Readonly)
        /// </summary>
        public readonly ImmutableDictionary<string, string> AnalysisVmEfficiencyMap;
        /// <summary>
        /// The API version.
        /// </summary>
        public readonly string ApiVersion;
        public readonly ImmutableArray<Outputs.GetClustersEntityAuthorizedPublicKeyListResult> AuthorizedPublicKeyLists;
        /// <summary>
        /// - Cluster build details.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Build;
        /// <summary>
        /// - Zone name used in value of TZ environment variable.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersEntityCaCertificateListResult> CaCertificateLists;
        /// <summary>
        /// - Categories for the image.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersEntityCategoryResult> Categories;
        /// <summary>
        /// - Customer information used in Certificate Signing Request for creating digital certificates.
        /// </summary>
        public readonly ImmutableDictionary<string, string> CertificationSigningInfo;
        /// <summary>
        /// - Client authentication config.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ClientAuth;
        /// <summary>
        /// - Cluster architecture. (Readonly, Options: Options : X86_64 , PPC64LE)
        /// </summary>
        public readonly string ClusterArch;
        /// <summary>
        /// - Cluster domain credentials.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DomainServerCredentials;
        /// <summary>
        /// - Joined domain name. In 'put' request, empty name will unjoin the cluster from current domain.
        /// </summary>
        public readonly string DomainServerName;
        /// <summary>
        /// -  The IP of the nameserver that can resolve the domain name. Must set when joining the domain.
        /// </summary>
        public readonly string DomainServerNameserver;
        /// <summary>
        /// - Array of enabled features.
        /// </summary>
        public readonly ImmutableArray<string> EnabledFeatureLists;
        /// <summary>
        /// - Cluster encryption status.
        /// </summary>
        public readonly string EncryptionStatus;
        /// <summary>
        /// - The cluster IP address that provides external entities access to various cluster data services.
        /// </summary>
        public readonly string ExternalDataServicesIp;
        /// <summary>
        /// - The local IP of cluster visible externally.
        /// </summary>
        public readonly string ExternalIp;
        /// <summary>
        /// - External subnet for cross server communication. The format is IP/netmask. (default 172.16.0.0/255.240.0.0)
        /// </summary>
        public readonly string ExternalSubnet;
        /// <summary>
        /// - GPU driver version.
        /// </summary>
        public readonly string GpuDriverVersion;
        /// <summary>
        /// - List of proxies to connect to the service centers.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersEntityHttpProxyListResult> HttpProxyLists;
        /// <summary>
        /// - HTTP proxy whitelist.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersEntityHttpProxyWhitelistResult> HttpProxyWhitelists;
        /// <summary>
        /// - The internal subnet is local to every server - its not visible outside.iSCSI requests generated internally within the appliance (by user VMs or VMFS) are sent to the internal subnet. The format is IP/netmask.
        /// </summary>
        public readonly string InternalSubnet;
        /// <summary>
        /// - Indicates if cluster is available to contact. (Readonly)
        /// </summary>
        public readonly bool IsAvailable;
        /// <summary>
        /// - List of cluster management servers. (Readonly)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersEntityManagementServerListResult> ManagementServerLists;
        /// <summary>
        /// - The cluster NAT'd or proxy IP which maps to the cluster local IP.
        /// </summary>
        public readonly string MasqueradingIp;
        /// <summary>
        /// - Port used together with masquerading_ip to connect to the cluster.
        /// </summary>
        public readonly int MasqueradingPort;
        /// <summary>
        /// - The image kind metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// - image name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// - The list of IP addresses of the name servers.
        /// </summary>
        public readonly ImmutableArray<string> NameServerIpLists;
        /// <summary>
        /// - Comma separated list of subnets (of the form 'a.b.c.d/l.m.n.o') that are allowed to send NFS requests to this container. If not specified, the global NFS whitelist will be looked up for access permission. The internal subnet is always automatically considered part of the whitelist, even if the field below does not explicitly specify it. Similarly, all the hypervisor IPs are considered part of the whitelist. Finally, to permit debugging, all of the SVMs local IPs are considered to be implicitly part of the whitelist.
        /// </summary>
        public readonly ImmutableArray<string> NfsSubnetWhitelists;
        public readonly ImmutableArray<Outputs.GetClustersEntityNodeResult> Nodes;
        /// <summary>
        /// - The list of IP addresses or FQDNs of the NTP servers.
        /// </summary>
        public readonly ImmutableArray<string> NtpServerIpLists;
        /// <summary>
        /// - Cluster operation mode. - 'NORMAL': Cluster is operating normally. - 'READ_ONLY': Cluster is operating in read only mode. - 'STAND_ALONE': Only one node is operational in the cluster. This is valid only for single node or two node clusters. - 'SWITCH_TO_TWO_NODE': Cluster is moving from single node to two node cluster. - 'OVERRIDE': Valid only for single node cluster. If the user wants to run vms on a single node cluster in read only mode, he can set the cluster peration mode to override. Writes will be allowed in override mode.
        /// </summary>
        public readonly string OperationMode;
        /// <summary>
        /// - The reference to a user.
        /// </summary>
        public readonly ImmutableDictionary<string, string> OwnerReference;
        /// <summary>
        /// - The reference to a project.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ProjectReference;
        /// <summary>
        /// - Array of enabled cluster services. For example, a cluster can function as both AOS and cloud data gateway. - 'AOS': Regular Prism Element - 'PRISM_CENTRAL': Prism Central - 'CLOUD_DATA_GATEWAY': Cloud backup and DR gateway - 'AFS': Cluster for file server - 'WITNESS' : Witness cluster - 'XI_PORTAL': Xi cluster.
        /// </summary>
        public readonly ImmutableArray<string> ServiceLists;
        /// <summary>
        /// - SMTP Server Address.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SmtpServerAddress;
        /// <summary>
        /// - SMTP Server Credentials.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SmtpServerCredentials;
        /// <summary>
        /// - SMTP Server Email Address.
        /// </summary>
        public readonly string SmtpServerEmailAddress;
        /// <summary>
        /// - SMTP Server Proxy Type List
        /// </summary>
        public readonly ImmutableArray<string> SmtpServerProxyTypeLists;
        /// <summary>
        /// - SMTP Server type.
        /// </summary>
        public readonly string SmtpServerType;
        /// <summary>
        /// - Map of software on the cluster with software type as the key.
        /// </summary>
        public readonly ImmutableDictionary<string, object> SoftwareMapNcc;
        /// <summary>
        /// - Map of software on the cluster with software type as the key.
        /// </summary>
        public readonly ImmutableDictionary<string, object> SoftwareMapNos;
        /// <summary>
        /// - UTC date and time in RFC-3339 format when the key expires
        /// </summary>
        public readonly string SslKeyExpireDatetime;
        public readonly string SslKeyName;
        /// <summary>
        /// - Customer information used in Certificate Signing Request for creating digital certificates.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SslKeySigningInfo;
        /// <summary>
        /// - SSL key type. Key types with RSA_2048, ECDSA_256 and ECDSA_384 are supported for key generation and importing.
        /// </summary>
        public readonly string SslKeyType;
        /// <summary>
        /// - The state of the cluster entity.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// - Verbosity level settings for populating support information. - 'Nothing': Send nothing - 'Basic': Send basic information - skip core dump and hypervisor stats information - 'BasicPlusCoreDump': Send basic and core dump information - 'All': Send all information (Default value: BASIC_PLUS_CORE_DUMP)
        /// </summary>
        public readonly string SupportedInformationVerbosity;
        /// <summary>
        /// - Zone name used in value of TZ environment variable.
        /// </summary>
        public readonly string Timezone;

        [OutputConstructor]
        private GetClustersEntityResult(
            ImmutableDictionary<string, string> analysisVmEfficiencyMap,

            string apiVersion,

            ImmutableArray<Outputs.GetClustersEntityAuthorizedPublicKeyListResult> authorizedPublicKeyLists,

            ImmutableDictionary<string, string> build,

            ImmutableArray<Outputs.GetClustersEntityCaCertificateListResult> caCertificateLists,

            ImmutableArray<Outputs.GetClustersEntityCategoryResult> categories,

            ImmutableDictionary<string, string> certificationSigningInfo,

            ImmutableDictionary<string, string> clientAuth,

            string clusterArch,

            ImmutableDictionary<string, string> domainServerCredentials,

            string domainServerName,

            string domainServerNameserver,

            ImmutableArray<string> enabledFeatureLists,

            string encryptionStatus,

            string externalDataServicesIp,

            string externalIp,

            string externalSubnet,

            string gpuDriverVersion,

            ImmutableArray<Outputs.GetClustersEntityHttpProxyListResult> httpProxyLists,

            ImmutableArray<Outputs.GetClustersEntityHttpProxyWhitelistResult> httpProxyWhitelists,

            string internalSubnet,

            bool isAvailable,

            ImmutableArray<Outputs.GetClustersEntityManagementServerListResult> managementServerLists,

            string masqueradingIp,

            int masqueradingPort,

            ImmutableDictionary<string, string> metadata,

            string name,

            ImmutableArray<string> nameServerIpLists,

            ImmutableArray<string> nfsSubnetWhitelists,

            ImmutableArray<Outputs.GetClustersEntityNodeResult> nodes,

            ImmutableArray<string> ntpServerIpLists,

            string operationMode,

            ImmutableDictionary<string, string> ownerReference,

            ImmutableDictionary<string, string> projectReference,

            ImmutableArray<string> serviceLists,

            ImmutableDictionary<string, string> smtpServerAddress,

            ImmutableDictionary<string, string> smtpServerCredentials,

            string smtpServerEmailAddress,

            ImmutableArray<string> smtpServerProxyTypeLists,

            string smtpServerType,

            ImmutableDictionary<string, object> softwareMapNcc,

            ImmutableDictionary<string, object> softwareMapNos,

            string sslKeyExpireDatetime,

            string sslKeyName,

            ImmutableDictionary<string, string> sslKeySigningInfo,

            string sslKeyType,

            string state,

            string supportedInformationVerbosity,

            string timezone)
        {
            AnalysisVmEfficiencyMap = analysisVmEfficiencyMap;
            ApiVersion = apiVersion;
            AuthorizedPublicKeyLists = authorizedPublicKeyLists;
            Build = build;
            CaCertificateLists = caCertificateLists;
            Categories = categories;
            CertificationSigningInfo = certificationSigningInfo;
            ClientAuth = clientAuth;
            ClusterArch = clusterArch;
            DomainServerCredentials = domainServerCredentials;
            DomainServerName = domainServerName;
            DomainServerNameserver = domainServerNameserver;
            EnabledFeatureLists = enabledFeatureLists;
            EncryptionStatus = encryptionStatus;
            ExternalDataServicesIp = externalDataServicesIp;
            ExternalIp = externalIp;
            ExternalSubnet = externalSubnet;
            GpuDriverVersion = gpuDriverVersion;
            HttpProxyLists = httpProxyLists;
            HttpProxyWhitelists = httpProxyWhitelists;
            InternalSubnet = internalSubnet;
            IsAvailable = isAvailable;
            ManagementServerLists = managementServerLists;
            MasqueradingIp = masqueradingIp;
            MasqueradingPort = masqueradingPort;
            Metadata = metadata;
            Name = name;
            NameServerIpLists = nameServerIpLists;
            NfsSubnetWhitelists = nfsSubnetWhitelists;
            Nodes = nodes;
            NtpServerIpLists = ntpServerIpLists;
            OperationMode = operationMode;
            OwnerReference = ownerReference;
            ProjectReference = projectReference;
            ServiceLists = serviceLists;
            SmtpServerAddress = smtpServerAddress;
            SmtpServerCredentials = smtpServerCredentials;
            SmtpServerEmailAddress = smtpServerEmailAddress;
            SmtpServerProxyTypeLists = smtpServerProxyTypeLists;
            SmtpServerType = smtpServerType;
            SoftwareMapNcc = softwareMapNcc;
            SoftwareMapNos = softwareMapNos;
            SslKeyExpireDatetime = sslKeyExpireDatetime;
            SslKeyName = sslKeyName;
            SslKeySigningInfo = sslKeySigningInfo;
            SslKeyType = sslKeyType;
            State = state;
            SupportedInformationVerbosity = supportedInformationVerbosity;
            Timezone = timezone;
        }
    }
}
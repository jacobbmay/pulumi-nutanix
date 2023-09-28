import * as pulumi from "@pulumi/pulumi";
import * as std from "@pulumi/std";
import * as nutanix from "@jacobbmay/nutanix";

let config = new pulumi.Config();

const provider = new nutanix.Provider('nutanix', {
    endpoint: config.require("PRISM_CENTRAL_ENDPOINT"),
    port: "443",
    username: config.require("PRISM_CENTRAL_USERNAME"),
    password: config.requireSecret("PRISM_CENTRAL_PASSWORD"),
    insecure: false,
    waitTimeout: 10,
    sessionAuth: true
});

const imageId = nutanix.getImage({imageName: config.require("imageName")}, {provider: provider}).then((res) => {
  return res.id
});
const clusterId = nutanix.getCluster({name: config.require("clusterName")}, {provider: provider}).then((res) => {
  return res.clusterId
});
const subnetId = nutanix.getSubnet({subnetName: config.require("subnetName")}, {provider: provider}).then((res) => {
  return res.id
});

const pubKey = config.require("pubkey")
const joinToken = "examplejointoken123456"

const cloudConfig = `
#cloud-config
hostname: bootstrap
users:
  - name: cloud-user
    sudo: ['ALL=(ALL) NOPASSWD:ALL']
    shell: /bin/bash
    ssh_authorized_keys:
    - ${pubKey}
ssh_pwauth: False
runcmd:
  - |
    server_ip=$(ip route get $(ip route show 0.0.0.0/0 | grep -oP 'via \\K\\S+') | grep -oP 'src \\K\\S+')
    /root/rke2-startup.sh -t ${joinToken} -T cluster.foo.bar -s \${server_ip} -u cloud-user
`;

  const rke2Bootstrap = new nutanix.VirtualMachine("rke2_bootstrap", {
    name: `bootstrap-node`,
    clusterUuid: clusterId,
    memorySizeMib: 8192,
    numSockets: 1,
    numVcpusPerSocket: 8,
    diskLists: [
        {
            dataSourceReference: {
                kind: "image",
                uuid: imageId
            },
            // deviceProperties: {
            //     diskAddress: { adapterType: "SCSI", deviceIndex: "0"},
            //     deviceType: "DISK",
            // },
            diskSizeMib: 102400,
        },
        {
            diskSizeMib: 102400,
        },
    ],
    nicLists: [{
        subnetUuid: subnetId,
    }],
    guestCustomizationCloudInitUserData: std.base64encodeOutput({
        input: cloudConfig,
    }).result,
}, {provider: provider});
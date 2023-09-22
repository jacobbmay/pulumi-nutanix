import * as pulumi from "@pulumi/pulumi";
import * as nutanix from "@jacobbmay/nutanix";
import { input } from "@jacobbmay/nutanix/types";
import { VirtualMachineDiskList } from "@jacobbmay/nutanix/types/input";

const provider = new nutanix.Provider('nutanix', {
    endpoint: process.env.PRISM_CENTRAL_ENDPOINT,
    port: "9440",
    username: process.env.PRISM_CENTRAL_USERNAME,
    password: process.env.PRISM_CENTRAL_PASSWORD,
    insecure: true,
    waitTimeout: 10,
    sessionAuth: true
  });

let config = new pulumi.Config();

const clusters = nutanix.getClusters();
const ubuntuImage = nutanix.getImage({imageName: "Ubuntu Server"});


const test_vm = new nutanix.VirtualMachine('test_vm', {
  name: "test_vm",
  clusterUuid: clusters.then((res) => {
    return res.entities[0].metadata.clusterUuid
  }),
  numVcpusPerSocket: 2,
  numSockets: 1,
  memorySizeMib: 4096,
  powerStateMechanism: "OFF",
  // diskLists: input<VirtualMachineDiskList>
  });
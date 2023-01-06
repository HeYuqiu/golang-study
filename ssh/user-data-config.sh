#!/bin/sh
DISABLE_KUBE_PROXY=true
CONTAINERD_MOUNT_DISK=
MOUNT_DISK_INFO=[{\"size\":40,\"volumeType\":\"ESSD_FlexPL\",\"mountPath\":\"/cce\",\"containerdMount\":true,\"kubeletMount\":true},{\"size\":30,\"volumeType\":\"ESSD_PL1\",\"mountPath\":\"/test\",\"containerdMount\":false,\"kubeletMount\":false}]
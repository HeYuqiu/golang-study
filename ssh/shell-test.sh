KUBERNETES_VERSION="v1.20.15-vke.4"
#KUBERNETES_VERSION="v1.20.12"
CRD_NAME_pre="test-pod-xxx"
for i in {1..50}
do
  CRD_NAME="test-pod-"${i}
  echo $CRD_NAME
  echo ${CRD_NAME_pre}
  sed -i "" "4s/${CRD_NAME_pre}/${CRD_NAME}/g" temp.yaml
  kubectl apply -f temp.yaml
  CRD_NAME_pre=${CRD_NAME}
  sleep 0.3
done


#CONTAINERD_MOUNT_DISK="vdb"
#volumeID=$(cat /sys/block/"${CONTAINERD_MOUNT_DISK}"/serial)
#if [ ${#CONTAINERD_MOUNT_DISK} -eq 3 ] && [ "${volumeID}" != '' ]; then
#    echo "${CONTAINERD_MOUNT_DISK}"
#    echo "${volumeID}"
#else
#    echo "11111"
#fi
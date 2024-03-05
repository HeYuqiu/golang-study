## 20230629 测试环境模拟线上vke-cluster-controller informer问题，from 何玉秋
for i in {1..500};do
#  kubectl rollout restart deploy vke-cluster-controller -n vke-system
  kubectl apply -f ~/go/src/github.com/admission-webhook-example/v1/deployment/sleep-with-labels.yaml
  echo "执行完第${i}轮"
  sleep 3
  kubectl delete -f ~/go/src/github.com/admission-webhook-example/v1/deployment/sleep-with-labels.yaml
  sleep 3
done


#sed -i 's/net.ipv6.conf.all.disable_ipv6 = 1/net.ipv6.conf.all.disable_ipv6 = 0/g' fstab.txt
#sed -i 'hyqtest' fstab.txt
#
#KUBERNETES_VERSION="v1.20.15-vke.1"
#KUBE_VERSION="v1.20.15-vke.6"
#echo ${KUBERNETES_VERSION#*v1.20.15-vke.}
#echo ${#KUBERNETES_VERSION}
#

#if [[ $KUBERNETES_VERSION == "v1.20.15-vke."* ]] && [[ ${#KUBERNETES_VERSION} -gt 14 || $KUBERNETES_VERSION < "v1.20.15-vke.4" ]]; then
#    echo "fff"
#fi
#
#if [[ $KUBERNETES_VERSION == "v1.20.15-vke."* ]]; then
#  if [[ ${#KUBERNETES_VERSION} -gt 14 || $KUBERNETES_VERSION > "v1.20.15-vke.4" ]]; then
#    echo "fff"
#  fi
#fi

#CONTAINERD_MOUNT_DISK="vdb"
#volumeID=$(cat /sys/block/"${CONTAINERD_MOUNT_DISK}"/serial)
#if [ ${#CONTAINERD_MOUNT_DISK} -eq 3 ] && [ "${volumeID}" != '' ]; then
#    echo "${CONTAINERD_MOUNT_DISK}"
#    echo "${volumeID}"
#else
#    echo "11111"
#fi

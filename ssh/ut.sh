#!/bin/bash
function sendMessage() {
    token=$1
    repos=$2
    data=""
    for repo in ${repos[@]}; do
        echo $repo
        coveragePrefix="https://badge.byted.org/ci/coverage/infcp"
        hrefPrefix="https://code.byted.org/infcp"
        if [ $repo == "csi-driver" ]; then
            coveragePrefix="https://badge.byted.org/ci/coverage/infcs"
            hrefPrefix="https://code.byted.org/infcs"
        elif [ $repo == "cello" ]; then
            coveragePrefix="https://badge.byted.org/ci/coverage/iaas"
            hrefPrefix="https://code.byted.org/iaas"
        elif [ $repo == "catena" ]; then
            coveragePrefix="https://badge.byted.org/ci/coverage/containernetworking"
            hrefPrefix="https://code.byted.org/containernetworking"
        fi
        coverage=$(curl -s "$coveragePrefix/$repo/default" | awk -F "<|>" '{print $(NF-6)}')
        coverage="${coverage/\%}"
        data="$data$coverage $repo $hrefPrefix/$repo\n"
    done

    content=""
    while IFS= read -r line; do
        # skip empty lines
        if [ -z "$line" ]; then
            continue
        fi
        coverage=$(echo $line | cut -d ' ' -f1)
        repo=$(echo $line | cut -d ' ' -f2)
        href=$(echo $line | cut -d ' ' -f3)
        if [ -z "$content" ]; then
            content="[{\"tag\":\"a\",\"text\":\"$repo\",\"href\":\"$href\"},{\"tag\":\"text\",\"text\":\": $coverage%\"}]"
        else
            content="$content,[{\"tag\":\"a\",\"text\":\"$repo\",\"href\":\"$href\"},{\"tag\":\"text\",\"text\":\": $coverage%\"}]"
        fi
    done <<< "$(echo -e "$data" | sort -n)"

    echo $content
#    curl -X POST "https://open.feishu.cn/open-apis/bot/v2/hook/$token" -H 'Authorization: Bearer t-8b2d47b17baxxxde101fbee' -H 'Content-Type: application/json' -d "{ \"email\": \"xuepengfei.xuepf@bytedance.com\", \"msg_type\": \"post\", \"content\": { \"post\": { \"zh_cn\": { \"title\": \"容器基础自研组件单测覆盖率报告\", \"content\": [$content] } } } }"
}

repos=(cloud-controller-manager alb-ingress-controller shuttle-operator dns-controller load-balancer-controller kube-extended-scheduler general-webhook vke-node-local-dns-admission node-job-helper vepfs-manager node-job-controller vke-vci-admission csi-driver cello catena vci-virtual-kubelet)
sendMessage "71dbf3d0-8bba-4523-a4ae-0dbb4dbc561e" ${repos[*]}

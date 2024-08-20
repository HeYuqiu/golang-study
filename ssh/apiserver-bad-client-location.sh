#!/bin/bash

# 检查输入参数
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <logfile>"
    exit 1
fi

LOGFILE=$1

declare -A all_requests
declare -A list_requests_by_url_ua
declare -A list_requests_by_client
declare -A long_requests
declare -A etcd_requests
declare -A etcd_long_requests

while IFS= read -r line; do
    if [[ $line =~ "Trace" ]]; then
        verb=$(echo "$line" | grep -oE 'verb:\K[^ ]+')
        url=$(echo "$line" | grep -oE 'url:\K[^,]+')
        user_agent=$(echo "$line" | grep -oE 'user-agent:\K[^,]+')
        client=$(echo "$line" | grep -oE 'client:\K[^,]+')
        total_time=$(echo "$line" | grep -oE 'total time: \K[0-9]+')

        key="$verb $url $user_agent"

        # 所有请求（不包含etcd）
        if ! [[ $line =~ "etcd3" ]]; then
            all_requests["$key"]=$((all_requests["$key"] + 1))

            # List 请求
            if [[ $verb == "LIST" ]]; then
                list_requests_by_url_ua["$url $user_agent"]=$((list_requests_by_url_ua["$url $user_agent"] + 1))
                list_requests_by_client["$client"]=$((list_requests_by_client["$client"] + 1))
            fi

            # 延时超过10秒的请求
            if [[ $total_time -gt 10000 ]]; then
                long_requests["$key"]=$((long_requests["$key"] + 1))
            fi
        else
            # etcd 请求
            etcd_requests["$key"]=$((etcd_requests["$key"] + 1))

            # etcd 延时超过10秒的请求
            if [[ $total_time -gt 10000 ]]; then
                etcd_long_requests["$key"]=$((etcd_long_requests["$key"] + 1))
            fi
        fi
    fi
done < "$LOGFILE"

# 函数：按值排序并打印前 N 个结果
print_top_n() {
    local -n arr=$1
    local n=$2
    local i=1
    for key in $(for k in "${!arr[@]}"; do echo "$k ${arr[$k]}"; done | sort -k2 -nr | head -n $n | awk '{print $1}'); do
        echo "$i. $key, Request times: ${arr[$key]}"
        ((i++))
    done
}

# 1. 所有请求按照 verb + url + user-agent 分组，按照请求次数从多到少排列输出 Top 10
echo "Top 10 requests by verb + url + user-agent:"
print_top_n all_requests 10

# 2. 所有List请求按照url + user-agent 分组，按照请求次数从多到少排列输出 Top 8
echo "Top 8 List requests by url + user-agent:"
print_top_n list_requests_by_url_ua 8

# 3. 所有List请求按照client（IP） 分组，按照请求次数从多到少排列输出 Top 8
echo "Top 8 List requests by client (IP):"
print_top_n list_requests_by_client 8

# 4. 所有延时超过10秒的请求，按照 verb + url + user-agent 分组，按照请求次数从多到少排列输出 Top 6
echo "Top 6 requests with delay over 10 seconds by verb + url + user-agent:"
print_top_n long_requests 6

# 5. Apiserver 对 ETCD的请求，按照 verb + key + limit 分组排序，输出请求次数从多到少的 top 8
echo "Top 8 ETCD requests by verb + key + limit:"
print_top_n etcd_requests 8

# 6. Apiserver 对 ETCD 延时超过10秒的请求，按照 verb + key + limit 分组排序，输出请求次数从多到少 top 6
echo "Top 6 ETCD requests with delay over 10 seconds by verb + key + limit:"
print_top_n etcd_long_requests 6
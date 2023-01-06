#!/bin/bash
### This script is only applicable to the debian series of operating systems
#set -e -x -u

function main() {
    for (( i = 0; i < 1440; i=(i+1) )); do
      echo "$(date +%F\ %T) 执行第${i}轮"
      if [ $(($i%2)) -eq 0 ];then
        kubectl apply -f deploy.yaml
      else
        kubectl delete -f deploy.yaml
      fi
      echo "$(date +%F\ %T) 第${i}轮执行完毕，sleep 1分钟"
      sleep 1m
    done
    echo "finish ."
}

main "$@"

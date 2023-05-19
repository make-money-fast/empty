#!/bin/bash

## 替换app为运行程序名称

ps -axf |grep app | grep -v 'app' | awk -F ' ' '{print $1}' | xargs kill -9

nohup ./app 2>&1 &
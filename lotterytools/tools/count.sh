#!/bin/bash

if [ $1 == "" ];then 
mysql -upkpk -h127.0.0.1 -P3306 -p@Dyf19840218@ -e "use pkpk; select currentPierod,putTime,numberList from NewHistoryPush;"|grep -v 'currentPierod'|grep -v putTime|grep -v numberList > count.txt
else
sql="use pkpk; select currentPierod,putTime,numberList,CreateTime from NewHistoryPush where CreateTime like \"$1%\""
mysql -upkpk -h127.0.0.1 -P3306 -p@Dyf19840218@ -e "$sql"|grep -v 'currentPierod'|grep -v putTime|grep -v numberList> count.txt
fi
cat count.txt |awk '{print $2}'| awk 'BEGIN{a=0}{if ($1>a){a=$1}else{print a;a=1}}'|sort |uniq -c

#!/bin/bash

if [ $1 == "" ];then 
mysql -upkpk -h127.0.0.1 -P3306 -p@Dyf19840218@ -e "use pkpk; select currentPierod,putTime,numberList from HistoryPush;"|grep -v 'currentPierod'|grep -v putTime|grep -v numberList > count.txt
else
sql="use pkpk; select currentPierod,putTime,numberList,CreateTime from HistoryPush where CreateTime like \"$1%\""
mysql -upkpk -h127.0.0.1 -P3306 -p@Dyf19840218@ -e "$sql"|grep -v 'currentPierod'|grep -v putTime|grep -v numberList> count.txt
fi
cat count.txt |awk 'BEGIN{a1=0;a2=0;a3=0;a4=0;a5=0;a6=0;a7=0;a8=0;a9=0;a10=0;cp=sprintf("%s|%s|%s|%s|%s",$3,$4,$5,$6,$7);aaa=0;ca=cp} {cp=sprintf("%s|%s|%s|%s|%s",$3,$4,$5,$6,$7);if (ca==cp) {aaa+=1} else {ca=cp ;if(aaa==1){a1+=1};if(aaa==2){a2+=1};if(aaa==3){a3+=1};if(aaa==4){a4+=1};if(aaa==5){a5+=1};if(aaa==6){a6+=1};if(aaa==7){a7+=1};if(aaa==8){a8+=1};if(aaa==9){a9+=1};if(aaa==10){a10+=1};aaa=1}}END {print sprintf("1 push %s time\n2 push %s time\n3 push %s time\n4 push %s time\n5 push %s time\n6 push %s time\n7 push %s time\n8 push %s time\n9 push %s time\n10 push %s time\n",a1,a2,a3,a4,a5,a6,a7,a8,a9,a10);}' 

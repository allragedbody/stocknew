mysql -upkpk -h127.0.0.1 -P3306 -p@Dyf19840218@ -e "use pkpk; select * from PK10 where id>690 order by period;"|grep -v number1 > list;cat list |sed 's/\t/|/g'>  numbers.list;dos2unix numbers.list; 

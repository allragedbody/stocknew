cat list20180327 |awk -F'（' '{print $2}'|awk -F'）' '{print $1}'|sort|uniq -c

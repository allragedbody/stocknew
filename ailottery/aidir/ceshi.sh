for i in `ls knn13_0*.py`;do echo $i;python $i |grep 'the total error rate is';done

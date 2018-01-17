# -*- coding: utf-8 -*-

from numpy import *
import operator

def createDataSet():
    group = array([[1.0,1.1],[1.0,1.0],[0,0],[0,0.1],[2,2],[2.1,2.0]])
    labels=["A","A","B","B","C","C"]
    return group,labels


def classify0(inX ,dataSet,labels,k):
    #计算给出的二级数组的个数，此处为6个。
    dataSetSize=dataSet.shape[0]
    #print(dataSetSize)
    #计算目标点到阵列中各个点的坐标距离。
    diffMat=tile(inX,(dataSetSize,1))-dataSet
    #print(diffMat)
    #将距离做平方
    sqDiffMat=diffMat**2
    #print(sqDiffMat)
    #将平方加和
    sqDistances=sqDiffMat.sum(axis=1)
    #print(sqDistances)
    #计算平方根，即为欧氏距离。
    distances=sqDistances**0.5
    #print(distances)
    #根据到目标点的欧氏距离对样例点进行排序
    sortDistIndicies=distances.argsort()
    #print(sortDistIndicies)
    classCount={}
    for i in range(k):
        #将点按照投票来排序 450132对应 CCAABB
        voteIlabel=labels[sortDistIndicies[i]]
        #print(voteIlabel)
        #计算字典中key是否存在，存在则加1，后面那个0代表从多少开始计数
        classCount[voteIlabel]=classCount.get(voteIlabel,0)+1
        #print(classCount[voteIlabel])
    sortedClassCount=sorted(classCount.items(),key=operator.itemgetter(1),reverse=True)
    #print(sortedClassCount)#计算计数之后的结果 [('C', 2), ('A', 1)]
    return sortedClassCount[0][0]


        
#group,labels=createDataSet()
#print(classify0([2,2],group,labels,3))
def file2matrix(filename):
    fr=open(filename)
    arrayOLines=fr.readlines()
    numberOfLines=len(arrayOLines)
    returnMat=zeros((numberOfLines,3))
    #print("returnMat %s",returnMat)
    classLabelVector=[]
    index=0
    for line in arrayOLines:
        line=line.strip()
        #切分每一行数据得到一个列表形式的数据
        listFromLine=line.split('\t')
        #为二维数组的每一份儿数据赋值，每行三个数据组成数组，文件有多长，这样的数组就有多少个。
        returnMat[index,:]=listFromLine[0:3]
        #将数组中最后一个分类数据存储在列表中，最后一项数据代表分类。根据喜欢程度分成三个层次。
        classLabelVector.append(int(listFromLine[-1]))
        index+=1
    return returnMat,classLabelVector


def autoNorm(dataSet):
    minVals=dataSet.min(0)
    maxVals=dataSet.max(0)
    ranges=maxVals-minVals
    #print("min %f max %f range %f",minVals,maxVals,ranges)
    normDataSet=zeros(shape(dataSet))
    m=dataSet.shape[0]
    normDataSet=dataSet-tile(minVals,(m,1))
    normDataSet=normDataSet/tile(ranges,(m,1))
    return normDataSet,ranges,minVals


#2-4 分类器针对约会网站的测试代码
def datingClassTest():
    hoRatio=0.10
    #将文件读入到矩阵当中
    datingDataMat,datingLabels=file2matrix('datingTestSet2.txt')
    #将矩阵归一化
    normMat,ranges,minVals=autoNorm(datingDataMat)
    #拿到矩阵的第一层的个数 1000
    m=normMat.shape[0]
    #print("m =",m)
    #取得百分之10的样本的数量
    numTestVecs=int(m*hoRatio)
    errorCount=0.0
    #1000行记录循环100次。
    for i in range(numTestVecs):
        #字段1输入测试点，从第一行到numtestvecs行。
        #字段2输入测试集合。从numtestvecs行到结尾。
        #字段3输入
        #print("111111111111",datingLabels[numTestVecs:m])
        #字段4输入k的个数。
        classifierResult=classify0(normMat[i,:],normMat[numTestVecs:m,:],datingLabels[numTestVecs:m],3)
        print("经过训练得到的结果为: %d, 实际结果为: %d" % (classifierResult,datingLabels[i]))
        if (classifierResult != datingLabels[i]):
            errorCount+=1.0
    #print(len(datingLabels[numTestVecs:m]))
    #print(datingLabels[numTestVecs:m])
    print("the total error rate is :%f" % (errorCount/float(numTestVecs)))
#import importlib
#importlib.reload(kNN)


#import matplotlib
#import matplotlib.pyplot as plt
#fig=plt.figure()
#ax=fig.add_subplot(111)
#ax.scatter(datingDataMat[:,0],datingDataMat[:,1],15.0*array(datingLabels),10.0*array(datingLabels))
#plt.show()
    

def classifyPerson():
    resultList=["几乎不可能","有可能","希望很大"]
    percentTats= float(input("玩儿视频游戏所耗时间百分比："))
    ffMiles=float(input("每年在天上的公里数："))
    iceCream=float(input("每年吃冰激凌的公升数："))
    #将文件读入到矩阵当中
    datingDataMat,datingLabels=file2matrix('datingTestSet2.txt')
    #将矩阵归一化
    normMat,ranges,minVals=autoNorm(datingDataMat)
    inArr=array([ffMiles,iceCream,percentTats])
    classifierResult=classify0(((inArr-minVals)/ranges),normMat,datingLabels,3)
    print("你喜欢这个人的程度预测为",resultList[classifierResult-1])


#算法总结：
    #1.读取一个文件，文件内容为可变为矩阵的内容。
    #2.读取文件后形成矩阵，并得到每一行数据的类型数据。
    #3.将矩阵取出大部分作为训练集，设计K临近算法对训练内容进行实际分类。
    #4。将剩余的数据内容部分作为实例，测试分类器的错误率。
    #5。如果训练集得到的错误率很低，就可以使用这个训练集进行相应的预测活动。
#该算法可用用于进行一些分类工作。可以适当修改数值和意义，来实现新的分类含义。
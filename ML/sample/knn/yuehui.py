# -*- coding: utf-8 -*-
"""
Created on Sat Jan 20 22:00:26 2018

@author: dengyunfei
"""
from numpy import *
import operator

#计算我自己喜欢的姑娘的类型
#首先加载一个文件，输入姑娘的各种信息指标以|作为分割
#具体字段含义如下：
#年龄|身高|体重|家庭|心理程度|脸部|胸部|手部|臀部|腿部|足部|喜欢程度 大概7个字段

def file2matrix(filename,fn):
    fr=open(filename)
    arrayOLines=fr.readlines()
    numberOfLines=len(arrayOLines)
    returnMat=zeros((numberOfLines,fn))
    #print("returnMat %s",returnMat)
    classLabelVector=[]
    index=0
    for line in arrayOLines:
        line=line.strip()
        #切分每一行数据得到一个列表形式的数据
        listFromLine=line.split('|')
        #为二维数组的每一份儿数据赋值，每行三个数据组成数组，文件有多长，这样的数组就有多少个。
        returnMat[index,:]=listFromLine[0:fn]
        #将数组中最后一个分类数据存储在列表中，最后一项数据代表分类。根据喜欢程度打分1-5。
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


def classifyPerson():
    resultList=["没感觉","有点儿感觉","一般喜欢","特别喜欢","挚爱之人"]
    age=float(input("年龄："))
    height= float(input("身高："))
    weight= float(input("体重："))
    family= float(input("家庭："))
    psychology= float(input("心理分数："))
    face= float(input("脸部："))
    chest= float(input("胸部："))
    hands= float(input("手部："))
    hips= float(input("臀部："))
    legs= float(input("腿部："))
    feet= float(input("足部："))

    #将文件读入到矩阵当中
    datingDataMat,datingLabels=file2matrix('yuehui.txt',12)
    #将矩阵归一化
    normMat,ranges,minVals=autoNorm(datingDataMat)
    inArr=array([age,height,weight,family,psychology,face,chest,hands,hips,legs,feet])
    classifierResult=classify0(((inArr-minVals)/ranges),normMat,datingLabels,1)
    print("你喜欢这个人的程度预测为",resultList[classifierResult-1])


resultList=["没感觉","有点儿感觉","一般喜欢","特别喜欢","挚爱之人"]
age=25
height= 172
weight= 100
family= 3
psychology= 7
face= 5
chest= 7
hands= 9
hips= 8
legs= 6
feet= 6

 #将文件读入到矩阵当中
datingDataMat,datingLabels=file2matrix('yuehui.txt',11)
#将矩阵归一化
normMat,ranges,minVals=autoNorm(datingDataMat)
inArr=array([age,height,weight,family,psychology,face,chest,hands,hips,legs,feet])
classifierResult=classify0(((inArr-minVals)/ranges),normMat,datingLabels,1)
print("你喜欢这个人的程度预测为",resultList[classifierResult-1])


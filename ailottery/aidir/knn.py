#!/usr/bin/python
# -*- coding: utf-8 -*-
from numpy import *
import numpy as np
import operator
import sys
#字符识别导入目录中的数据
from os import listdir



def classify0(inX ,dataSet,labels,k):
    #计算给出的二级数组的个数，此处为6个。
    dataSetSize=dataSet.shape[0]
    #print(dataSetSize)
    #计算测试点列到阵列中各个点的坐标距离，并形成距离矩阵。
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
    #根据到目标点的欧氏距离对训练集进行排序
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

def file2matrix(filename):
    fr=open(filename)
    arrayOLines=fr.readlines()
    numberOfLines=len(arrayOLines)
    returnMat=zeros((numberOfLines,70))
    #print("returnMat %s",returnMat)
    classLabelVector=[]
    index=0
    for line in arrayOLines:
        line=line.strip()
        #切分每一行数据得到一个列表形式的数据
        listFromLine=line.split(',')
        #为二维数组的每一份儿数据赋值，每行三个数据组成数组，文件有多长，这样的数组就有多少个。
        returnMat[index,:]=listFromLine[0:70]
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

def autoNorm1(dataSet):
    minVals=dataSet.min(0)
    maxVals=dataSet.max(0)
    ranges=maxVals-minVals
    #print("min %f max %f range %f",minVals,maxVals,ranges)
    normDataSet=zeros(shape(dataSet))
    normDataSet=dataSet
    return normDataSet,ranges,minVals


def classifyPerson():
    resultList=[1,2]
    #将文件读入到矩阵当中
    datingDataMat,datingLabels=file2matrix('./missknnlist7.txt')
    #将矩阵归一化
    normMat,ranges,minVals=autoNorm(datingDataMat)
    inArrList=[]

    for i in range(1, len(sys.argv)):
        inArrList.append(int(sys.argv[i]))
    inArr=np.array(inArrList) 
    classifierResult=classify0(((inArr-minVals)/ranges),normMat,datingLabels,3)
    print(resultList[classifierResult-1])


def datingClassTest():
    hoRatio=0.15
    #将文件读入到矩阵当中
    datingDataMat,datingLabels=file2matrix('./missknnlist7.txt')
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
        #inArr=numTestVecs[i,:]    
        #classifierResult=classify0((inArr-minVals)/ranges,normMat[numTestVecs:m,:],datingLabels[numTestVecs:m],3)
        classifierResult=classify0(normMat[i,:],normMat[numTestVecs:m,:],datingLabels[numTestVecs:m],7)
        if classifierResult==datingLabels[i]:
            str=1 
        else:
            str=0
        print("经过训练得到的结果为: %d, 实际结果为: %d 结果 %d" % (classifierResult,datingLabels[i],str))
        if (classifierResult != datingLabels[i]):
            errorCount+=1.0
    #print(len(datingLabels[numTestVecs:m]))
    #print(datingLabels[numTestVecs:m])
    print("the total error rate is :%f" % (errorCount/float(numTestVecs)))

datingClassTest()

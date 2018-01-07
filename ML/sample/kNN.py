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
    print(distances)
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
    classLabelVector=[]
    index=0
    for line in arrayOLines:
        line=line.strip()
        listFromLine=line.split('\t')
        returnMat[index,:]=listFromLine[0:3]
        classLabelVector.append(int(listFromLine[-1]))
        index+=1
    return returnMat,classLabelVector


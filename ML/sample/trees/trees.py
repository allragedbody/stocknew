# -*- coding: utf-8 -*-
"""
Created on Mon Jan 22 21:48:51 2018

@author: dengyunfei
"""
#加载对数运算源代码
import operator
from math import log

def createDataSet():
    dataSet = [[1,1,"鱼类"],[1,1,"鱼类"],[1,0,"不是鱼类"],[0,1,"不是鱼类"],[0,1,"不是鱼类"]]
    labels=["不浮出水面是否可以生存","是否有脚蹼"]
    return dataSet,labels

#计算香农熵，将数据中的最后一个分类信息出现的次数分别计数，并最终计算概率和对数乘积并加在一起取负数的结果。
#说明一个特征如果能顺利将列表分成特定的多个集合，每个集合的内部混乱程度较小，就是较好的分类。
def calcShannonEnt(dataSet):
    #计算数组中实例的总数
    numEntries=len(dataSet)
    labelCounts={}
    for featVec in dataSet:
        currentLabel=featVec[-1]
        if currentLabel not in labelCounts.keys():
            labelCounts[currentLabel]=0
        labelCounts[currentLabel] += 1
    shannonEnt=0.0
    for key in labelCounts:
        #计算每种类型的情况在总数中的比例，及概率
        prob=float(labelCounts[key])/numEntries
        #计算熵
        shannonEnt-=prob*log(prob,2)
    return shannonEnt

#找到第axis项特征值为value的项目组成列表
def splitDataSet(dataSet,axis,value):
    retDataSet=[]
    for featVec in dataSet:
        #如果该数据集第axis项目的内容等于value就作如下操作
        if featVec[axis]==value:
            #去掉该特征值然后得到一个新的一维数组
            reducedFeatVec=featVec[:axis]
            reducedFeatVec.extend(featVec[axis+1:])
            #将一维数组合在一起得到一个两位数组
            retDataSet.append(reducedFeatVec)
    return retDataSet


#通过给出的数据集，得到最优化分类那一列的特征的列号。
def chooseBestFeatureToSpit(dataSet):
    #计算每一个列表中的用于描述特征的值的个数。
    numFeatures=len(dataSet[0])-1
    #print(numFeatures)
    #计算整个数据集的原始香农熵
    baseEntropy=calcShannonEnt(dataSet)
    bestInfoGain=0.0;bestFeature=-1
    for i in range(numFeatures):
        #得到每个数据集合中指定特征的值并形成列表
        featList=[example[i] for example in dataSet]
        #print(featList)
        #将特征值变成集合
        uniqueVals=set(featList)
        #print(uniqueVals)
        #前面的工作都是铺垫，是为了获得特征及值的集合
        newEntropy=0.0
        for value in uniqueVals:
            subDataSet=splitDataSet(dataSet,i,value)
            #print(subDataSet)
            prob=len(subDataSet)/float(len(dataSet))
            newEntropy+=prob*calcShannonEnt(subDataSet)
        infoGain=baseEntropy-newEntropy
        if (infoGain>bestInfoGain):
            bestInfoGain=infoGain
            bestFeature =i
    return bestFeature

def majorityCnt(classList):
    classCount={}
    for vote in classList:
        if vote not in classCount.keys():
            classCount[vote]=0
        classCount[vote]+=1
    sortedClassCount=sorted(classCount.items(),key=operator.itemgetter(1),reverse=True)
    return sortedClassCount

def createTrees(dataSet,labels):
    #确定分类的列表
    classList=[example[-1] for example in dataSet]
    #如果类型列表中的所有项都和第一项的值相同，停止划分
    if classList.count(classList[0])==len(classList):
        return classList[0]
    #如果数据集合最终只剩下类型一项，则根据多数进行类型确定
    if len(dataSet[0])==1:
        return majorityCnt(classList)
    #如果以上都不符合，就选取最有效的划分特征
    bestFeat=chooseBestFeatureToSpit(dataSet)
    #拿到该特征的特征值
    bestFeatLabel=labels[bestFeat]
    #根据该特征值创建树节点
    myTree={bestFeatLabel:{}}
    #删掉该特征
    del(labels[bestFeat])
    #得到最佳特征的集合
    featValues=[example[bestFeat] for example in dataSet]
    uniqueVals=set(featValues)
    #根据最佳特征进行下一步创建子树的过程
    for value in uniqueVals:
        subLabels=labels[:]
        #创建子树
        myTree[bestFeatLabel][value]=createTrees(splitDataSet(dataSet,bestFeat,value),subLabels)
    return myTree

    
#m,l=trees.createDataSet()
#mt=trees.createTrees(m,l)
#print(mt)
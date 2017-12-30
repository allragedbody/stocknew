# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""

import numpy as np
import matplotlib.pyplot as plt

x,y=[],[]
for sample in open("../Data/prices.txt","r"):
    _x,_y=sample.split(",")
    #字符串转化为浮点数
    x.append(float(_x))
    y.append(float(_y))
#读取完数据后，将他们转化为Numpy数据以方便进一步处理。
x,y=np.array(x),np.array(y)
#标准化
x=(x-x.mean())/x.std()
#将原始数据散点图形式画出
plt.figure()
plt.scatter(x,y,c="g",s=6)
plt.show()


x0=np.linspace(-10,10,1000)

#多项式根据x 和 p 计算y
#通过47个点的训练，得到最佳的p列表值，根据p列表值计算y值。
def polynomial(deg,potx):
    return np.polyval(np.polyfit(x,y,deg),potx)

#计算一系列的点的列表
def get_model(deg,input_x):
    return [polynomial(deg,potx) for potx in input_x]

def get_cost(deg,input_x,input_y):
    b=0
    for i in range(len(get_model(deg,input_x))):
        b=b+((get_model(deg,input_x)[i]-input_y[i])**2)
        c=0.5*b
    return c

test_set=(1,4,10)
for d in test_set:
    print(get_cost(d,x,y))
plt.scatter(x,y,c="g",s=20)
for d in test_set:
    plt.plot(x0,get_model(d,x0),label="degree = {}".format(d))
plt.xlim(-2,4)
plt.ylim(1e5,8e5)
plt.legend()
plt.show
#https://github.com/carefree0910/MachineLearning/blob/master/_Data/prices.txt
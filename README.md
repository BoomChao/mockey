# mockey
利用 mockey 写 Golang 单测

参考：https://github.com/bytedance/mockey/tree/main

1. 运行的时候禁用内联:-gcflags="all=-l -N"
2. GetMethod 无法对私有方法进行mock, 只能mock公有的方法
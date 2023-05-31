docker image tutorial 

通过制作c语言、java、go语言的docker镜像来了解制作docker镜像的过程。其中注意到c/c++/go语言生成的镜像非常小，是轻量可行的。但是java环境的docker镜像太大了，足足有150MB，看了都觉得害怕。
```
REPOSITORY           TAG       IMAGE ID       CREATED          SIZE
demo03               v0        a1f243590be5   48 seconds ago   13.6MB
demo02               v0        e7e008b639cb   37 minutes ago   9.35MB
demo01               v0        e59174d56fa0   59 minutes ago   150MB
demo00               v1        28906c9b7f0c   2 hours ago      7.34MB
demo00               v         db72820da683   2 hours ago      7.34MB
alpine               latest    5e2b554c1c45   3 weeks ago      7.33MB
redis                latest    116cad43b6af   4 weeks ago      117MB
```
其中demo03,demo02,demo00分别是go http服务器，go原生程序，c语言原生程序，demo01是一个包含jvm的Java程序。可以比较4者之间的大小关系。其中遇到的问题有
- gcc要使用-static选项将所有的依赖的静态文件汇集到编译文件中，而不是运行时链接。否则在docker运行时会报` no such file or directory`的错误
- javac的本地版本和docker环境中的版本要一致，要不要会报` no such file or directory`的错误
- go编译文件需要特定的docker image来进行编译，要不然将本地编译好的文件制作成镜像会遇到` no such file or directory`的错误

综上，如果不能很简洁地将本地编译好的文件制作成docker的image并运行，而是依赖特定的docker环境进行编译运行，那么docker的使用有和意义。跟一般的运维将本地代码上传到生产服务器编译运行有什么区别。这是一点docker image制作过后一点点想法。

[制作docker image过程中遇到的问题和解决方法](https://docs.qq.com/doc/DRm9XbEhDUkN5bUZr)

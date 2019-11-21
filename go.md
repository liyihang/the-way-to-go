#Go语言基础

##1、Go安装和配置
###1.1、Go的安装

###1.2、gopath的配置
Go安装完成后我们可以在命令行执行`echo %gopath%`命令查看相关`GOPATH`相关信息。
```
C:\Users\29473>echo %gopath%
C:\Users\29473\go
```
如果没有相关的`gopath`设置，可以使用`set GOPATH=c:/go`
###1.3、编辑器的选择
现在比较好用编辑器，个人推荐是`goland`,这个工具使用非常方便。链接地址如下：https://pan.baidu.com/s/1jmk-K8-a_bWHhrvi-W8ZNA
##2、Go语言基础
###2、1 Go语言常用关键字
```go，
break   default     func    interface   select  case
defer   go  map struct  chan    else    goto    package
switch  const   fallthrough if  range   type    continue    
for import  return  var  
```
###2、2 Go变量常量
**var**

var关键字是Go最基本的定义变量方式，与C语言不通的是Go把变量类型放到了变量名后面，如下：
```go，
var x int
x = 10
```
我们还可以使用`x:=10`来定义变量，这种方式只能在函数内部使用。
**const常量**
```go,
const PI = 3.1415926
const name string = "hello"
```
###2、3 Go数据类型
```go,
bool
rune
int8 int16 int32 int64
byte
unit8 unit16 unit32 unit64
float32 float64
complex64 complex128
string
array slice
map
```
####array
array就是数组，定义方式如下：
```go,
var x [n]int{}
```
数组可以使用另外一种方式`:=`来声明。
```go,
func main() {
	//var x = [10]int{1,3,4,5,7}
	x := [5]int{1,2,3,4,5}
	fmt.Print(x)
}
```
####slice动态数组
slice动态数组可以通过一个数组或者一个已经存在的slice中再次声明。
```go,
func main() {
	x:=[10]int{1,2,3,4,5,6,7,8,9,10}
	y:=x[1:3]
	fmt.Print(y)
}
```
我们也可以通过make来声明动态数组。
```go,
func main() {
	x:=make([]int,3,5) // 创建cap为5 length为3的slice
	x = append(x,4,5,6,77)
	//fmt.Print(x)
	fmt.Print(cap(x)) //cap 10  动态增加了
}
```
####map
**map的定义**
map是一堆键值对的未排序集合，它的格式为map[keyType]valueType，是一个key-value的hash结构。
map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是int类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。
**map的初始化**
```go,
var map[keyType]valueType
```
其中keyType为键类型，valueType为值类型。valueType是标注值类型的，可以自定义。
```go,
type personInfo struct{
    id string
    name string
    address string
}
```
直接初始化：
```go,
info := map[string]string{"name":"jim","sex":"boy"}
```
通过make初始化
```go,
func main() {
	var info map[string]string
	//map 必须分配内存 否则报错
	info= make(map[string]string)
	info["name"] = "jim"
	fmt.Printf("%v",info)
}
```
**map的遍历**
```go,
for k, v := range info {
		fmt.Print(k, v)
	}
```
**map的取值**
```go,
info:=map[string]string{
		"name":"jj",
		"age":"20",
	}
	name := info["name"]
	fmt.Print(name)
```
**map的删除**
```go,
info:map[string]string{
    "name":"jj",
    "age":"20"
}
delete(m,"name")
```
##3、流程控制
**if**
```go,
    a := int(10)
	if a > 10 {
		fmt.Print(10)
	} else if a > 50 {
		fmt.Print(50)
	} else {
		fmt.Print(1000)
	}
```
**switch**
Go语言中的switch不需要break。
```go,
func main() {
	var class int
	class = 80
	switch class {
	case 80:
		fmt.Print("优秀")
	case 90:
		fmt.Print("厉害")
	default:
		fmt.Print("牛啤")
	}
}
```

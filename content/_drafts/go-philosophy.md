---
title: Go 语言的编程哲学
tags:
  - 编程
  - 计算机
  - 设计
date: 2025-11-24T08:45:00
draft: true
toc: true
---

从这个月 10 号，水逆期开始，我就开始减少出门，宅在家里读一读书，捡一捡法语学习，看看音乐剧，还找了些新东西来学，Go 语言就是其中一个。一开始只是抱着好奇心探索，并没有什么功利心，不过越学越发觉，Go 真的很适合用来写高并发高性能的 Web 应用（而且越学越讨厌 Java）。

学任何一门编程语言我都更喜欢「不自量力」地直接上手写代码，一边写一边学基础语法，遇到不知道怎么实现的需求时再查阅文档。不过，学 Go 的一开始我发现了 [Go Proverbs](https://go-proverbs.github.io/) 这个网站，便在实战之外还花了不少时间研读这里的「名言警句」。这些 Go 语言编程相关的谚语大多来自 Rob Pike 在 2015 年的[一场演讲](https://www.youtube.com/watch?v=PAAkCSZUG1c)，每一条都很有启发性，这篇文章就来逐条阐述我对它们的理解，其中有不少也可以用在 Go 语言以外的编程中。<!--more-->

## Don't communicate by sharing memory, share memory by communicating.

这句话直接翻译过来是：不要用共享内存的方式通信，要用通信的方式共享内存。

这需要一些上下文来理解。Go 是一门可以通过非常优雅的方式实现并发的编程语言，异步编程若是要在多个线程之间同步信息，往往需要加锁，这是为了避免线程同时操作内存，导致冲突、数据准确性（integrity）被破坏等无法预测的问题。加锁一般指的是互斥锁（mutual exclusion object，简写作 `mutex`），线程操作数据时需要先获取对应的锁，如果锁被其他线程持有，则需要等待，获取到锁的线程才能执行操作，这保证数据操作是互斥的。

加锁是异步编程中的常见做法[^1]，但 Go 程序员认为「锁是痛苦的」，不仅要创建锁、管理锁，还要想办法避免死锁等运行时错误，这类错误没法被编译器察觉，程序员排查起来很困难。所以，为了保证程序的安全性和可读性，应当避免使用锁。使用互斥锁保证数据通信安全的方法，就是这句话的前半句「communicate by sharing memory」，因为共享了内存，所以必须加锁，不然就会出现问题，而锁又会带来更多难题。

Go 语言鼓励的「share memory by communicating」，要用到 Go 语言的一个特性—— `channel`。你可以把 `channel` 理解为在不同线程之间传递数据的「通道」。以下是一个整数类型的 `channel` 例子。

```go
intChan := make(chan int)

// 线程 A 传入一个数据
intChan <- 10

// 线程 B 取出这个数据
someVar := <- intChan
```

严格来说，这两个线程仍然共享了 `intChan` 这个 `channel` 类型的变量，但整体上，这种写法避免了直接读写被储存在内存中的数据。在线程 A 没有传入数据时，线程 B 执行到 `<- inChan` 这一行就会自动阻塞，直到取到 `channel` 中的数据。在 `channel` 的**抽象**下，线程不再操作一块无人管理的，需要线程自觉加锁、等待用锁的数据区域，而是相互之间直接传递数据——这就是「share memory by communicating」。

Go 语言实现并发的方式非常优雅，而 `channel` 的存在又简化了线程之间共享数据的方式，这使得用 Go 编写并发程序的体验非常舒服，这一点在接下来还会看到。

## Concurrency is not parallelism.

这句话直接翻译过来是：并发不是并行。

这其实是计算机科学的基础知识，并发（concurrency）和并行（parallelism）本来就不是一回事。并发指的是程序在处理器上交替执行，在单核处理器上，一个进程完全占用 CPU 资源，完成计算或者用完时间片之后，其运行的上下文会被保存，然后 CPU 进行上下文切换（context switch），换上下一个要执行的进程的上下文，然后运行下一个进程。每个进程执行时都完全占用一个处理器内核，时间到了之后就让下一个进程执行，只不过进程在 CPU 上切换的速度非常快，宏观上让人感觉是在同步执行。并行指的是真正的同步执行，在多核处理器上，不同的进程在不同的处理器内核上执行，而不需要交替执行。

高级语言编程一般是跨平台的，不需要考虑硬件底层，也就是说，编程语言不关心执行程序的计算机究竟有多少个处理器核心。Go 语言强调并发并非并行的原因是，Go 把并发当作一种「编程结构」，即组织程序的方式，而并行是否能实现要取决于硬件和调度器，编程语言无法决定，也就是说，在编程语言层面无法实现并行。程序员在进行多线程编程时，应该聚焦于程序结构而非硬件底层，而且也不能理所应当地认为自己用了并发的结构，程序就一定是并行的。

Go 语言里实现并发非常简单，假设我们有一个 `FetchData()` 函数，而这个函数需要进行大量 IO 操作或者需要发送 HTTP 请求，如果同步执行会长时间阻塞主线程，要把这个函数变成并发的，只需要在调用函数时，在前面加一个关键词 `go`。

```go
go FetchData()	// 这样就创建了一个线程执行 FetchData()
```

在 Go 语言里，一个线程被称作 `goroutine`，和 Java 中 `Thread` 的显著区别是：`Thread` 是一个类，需要创建和管理对象；`goroutine` 是一种**控制结构**，直接写就好了。另一个很重要的区别是，Java 线程占用的内存是 MB 级别的，而一个最基本的 `goroutine` 大小只有 2 KB。简单来说，`goroutine` 用起来很简单，创建的成本也很低。

很重要的一个点是，将并发作为一种控制结构也让代码变得更可读了，比如上面那一行代码，就可以很有喜感地读成：*Go! Fetch data!*

以下是一个 `goroutine` 配合 `channel` 的并发示例，代码来自我的项目 [wthis](https://github.com/BigCoke233/wthis)：

```go
fmlChan := make(chan *FormulaInfo)
caskChan := make(chan *CaskInfo)
rvsChan := make(chan []string)
// start 2 goroutines, fetching formula/cask info and uses
go func() {
	formula, cask := GetBrewInfo(pkgName)
	fmlChan <- formula
	caskChan <- cask
}()
go func() {
	rvsChan <- GetBrewUses(pkgName)
}()
stat = NewStatistics(<-fmlChan, <-caskChan, pkgName, <-rvsChan)
```

总结一下，关注程序结构，写出更清晰的代码，这是并发；至于并行，那是硬件和操作系统要考虑的事情。

## Channels orchestrate; mutexes serialize.

这句话直接翻译过来是：Channels 编排；互斥锁串行。

尽管 Go 提供了在大部分时候都更好用的 `channel`，也更鼓励使用 `channel`，但互斥锁并不是被禁止的，也没有从这门语言里消失，因为有些时候必须要用到锁。这句话简单回答了「什么时候该用 channel？」「什么时候该用锁？」这个问题，也能帮助 Go 程序员更好地理解 `channel` 的用途。

编排这个词很有意思，它的原文是 orchestration，又名 choreography，意思分别是「编配管弦乐曲」和「编舞」。如果你理解概念，就会觉得这个描述非常生动形象。这个词在不同的语境下的意思有一些区别，但一般来说，是指自动化的协调管理、统一掌控和分工协作。在管弦乐（orchestration）中，各个演奏者各司其职，但由指挥家统一主导；在舞蹈（choreography）中，舞者之间需要相互沟通，分工协作。这引申到编程中，对应了不同的编程结构，下面用简单的例子来解释。

假设有一项操作需要大量计算，这些计算被分散开，启动多个 `goroutines` 执行，最后由主线程将计算的结果整合起来，Go 语言代码可以这样写：

```go
// 有 100 项作业需要执行，相应地有 100 项结果需要接收
// 分别创建 channel
jobs := make(chan int, 100)
results := make(chan int, 100)

// 启动 10 个 Worker 执行 100 个作业
for _ := range 10 {
	go Worker(jobs, results)
}

// 传入 100 个作业给 jobs channel
// worker 接收到之后就会执行
for j := range 100 {
	jobs <- j
}

// 收集 100 个作业结果
for _ := range 100 {
	result := <- results
	doSomethingWith(result)
}
```

这是一个典型的 Worker Pool。这个例子中，主线程就是指挥，而 Workers 就是乐队成员，整体上，这就是一种编排（orchestration）。

假设有三个作业（A、B、C），C 作业依赖 A 和 B 的结果，而 B 依赖 A 的结果，要让这三个作业相互配合完成工作，可以这样写：

```go
a := make(chan int)
b := make(chan int)
c := make(chan int)

// A
go func() {
	result := doSomething()
	a <- result
}()

// B
go func() {
	result := doSomething(<-a)
	b <- result
}()

// C
go func() {
	result := doSomething(<-a, <-b)
	c <- result
}()
```

三个作业各司其职，使用 `channel`  相互通信完成了工作，A、B 和 C 都可以被视作相互配合的舞者，这也是一种编排（choreography）。

Go 语言还提供 `select` 控制结构，用于监听和操作 `channel`。在线程等待多个操作返回结果的时候，可以使用 `select`，数据先从哪个管道过来，就先处理哪一个。

```go
// 例子来源：https://gobyexample.com/select
c1 := make(chan string)
c2 := make(chan string)

go func() {
	time.Sleep(1 * time.Second)
	c1 <- "one"
}()
go func() {
	time.Sleep(2 * time.Second)
	c2 <- "two"
}()

for range 2 {
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}
}
```

`select` 和 `channel` 为 Go 语言提供了十分强大的编排能力，允许程序员在大多数情况下高效处理业务，而不是机械地等待上一个操作完成再做下一个。当然，严格的串行也是存在的，通常是为了保证某一操作的**原子性**（Atomic），即不可被打断性，这个时候就要用到锁。

在 Go 中使用锁需要用到标准库 `sync`，一个常见的做法是把 `sync.Mutex` 放到使用它的结构体中。以下例子实现了一个计数器容器，一个 `Container` 中有多个计数器（`counter`），`inc`（自增）方法用于自增指定计数器。默认情况下 `++` 是一个非原子（non-atomic）操作，如果多个 `goroutine` 都在对同一个计数器自增，极有可能会相互影响——例如，A 读取到计数器的值是 5，自增后写回 6；B 在 A 写回之前也读取了计数器，值也是 5，也写回 6，这就覆盖了 A 的操作；两次自增的结果本应该是 7，但由于操作的非原子性，变成了错误的 6。

为了保证自增操作的原子性，可以在自增时加锁，不允许其他 `goroutine` 读写。

```go
// 例子来源：https://gobyexample.com/mutexes
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
// ...
```

注意 `defer c.mu.Unlock()` 这一行代码。`defer` 关键词表示这行代码会在其周边代码执行完毕之后被执行，用在 `Unlock()` 和 `Close()` 这样的操作前面，就可以把它们紧挨着 `Lock()` 和 `Open()`，防止忘记关锁。`defer` 还会在程序 `panic`（遇到无法处理的错误）时执行，因此 `defer` 也经常被用来恢复错误。

以上操作也可以直接用 Go 标准库提供的原子计数器，不需要自己封装。

```go
var ops atomic.Uint64
ops.Add(1)
```

总而言之，`channel` 用于高效地编排程序，而 `mutex`（互斥锁）用于严格地保证串行。

## The bigger the interface, the weaker the abstraction.

这句话直接翻译过来是：接口越大，抽象越弱。

熟悉面向对象编程的程序员应该不会对 `interface`（接口）感到陌生，它用于规定一个类应该实现什么方法，也让静态类型语言的类型系统变得更强大。例如，Java 中的 `List` 接口规定了列表应该实现 `add()` 等方法，同时也让底层实现不同的列表（`ArrayList`、`LinkedList`...）可以通用——如果一个方法要求传入一个 `List` 作为参数，无论传入 `ArrayList` 还是 `LinkedList` 都是可以接受的。

Go 语言鼓励程序员编写「小巧」的接口，就像这样：

```go
type Writer interface {
	Write()
}
```

只有一两个方法的接口看起来很没用，但语义非常明确：Writers write; Dancers dance; Eaters eat. 用 -er 后缀结尾的词命名接口几乎是 Go 语言的惯例。小巧的抽象是强大的，它能让程序员更灵活地使用 Go 语言的类型系统。

例如，这是一个日志接口：

```go
type Logger interface {
	Log()
}
```

在某个系统中，日志可能有三种输出方式：直接输出到终端、写入文件和发送 HTTP 请求给远程服务器。

```go
type StdoutLogger struct{}

func (s StdoutLogger) Log(msg string) {
    fmt.Println(msg)
}
```

```go
type FileLogger struct {
    f *os.File
}

func (fl FileLogger) Log(msg string) {
    fmt.Fprintln(fl.f, msg)
}
```

```go
type HttpLogger struct {
    endpoint string
}

func (h HttpLogger) Log(msg string) {
    http.Post(h.endpoint, "text/plain", strings.NewReader(msg))
}
```

现在，我们正常编写业务逻辑，在处理某个数据的时候，突然有了记录日志的需求。由于我们定义了 `Logger` 接口，处理数据时不需要思考应该用哪种方式输出日志，只需要接收一个 `Logger`，直接 `Log()` 就好。调用这个函数时再传入具体的 `StdoutLogger` `FileLogger` 或者 `HttpLogger`。

```go
func Process(data string, logger Logger) {
    logger.Log("processing: " + data)
}
```

短小的接口看起来没用，但接口完全可以像搭积木一样把方法堆叠起来，比如下面这个类型就既是 `Writer` 又是 `Closer` 还是 `Logger`。

```go
type FileWriter struct { 
	f *os.File
}

func (*fw FileWriter) Write() { ... }
func (*fw FileWriter) Close() { ... }
func (*fw FileWriter) Log() { ... }
```

`interface` 实际上是对类型的「限制」，规定的方法越多，接口越大，要实现这个接口就越难，应用的场景也就越局限。Go 语言不是严格意义上的面向对象编程语言，它更关注**类型**，而短小的接口使得程序员可以更灵活地使用类型系统。

顺带一提，和面向对象编程语言不同，Go 语言不需要显式地指定实现哪些接口，只要包含了接口规定的方法的类型，都会被归为这个接口。

```java
// 在 Java 中实现接口
class FileWriter implements Writer { 
	public void write() { ... }
}
```

```go
// 在 Go 中实现接口
type FileWriter struct { ... }
func (*fw FileWriter) Write() { ... }
```

## Make the zero value useful.

这句话直接翻译过来是：让零值变得有用。

Zero value 指的是变量被声明但尚未初始化（赋值），类型系统自动为变量填充的值。对字符串来说，这个值是 `""`，整数是 `0`，浮点数是 `0.0`，布尔值是 `false`，指针是 `nil`。结构体被声明时，其中的变量也会被填充零值。

Go 语言鼓励程序员好好利用这个零值，让类型即时没有初始化也能正常使用。标准库里有很多例子，前面提到的 `sync.Mutex` 就是一个。让我们回看这个例子：

```go
// 例子来源：https://gobyexample.com/mutexes
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
// ...
```

显然，`inc()` 方法没有初始化 `c.mu` 这个 `sync.Mutex` 类型的变量，`c.mu` 里面装的是默认的零值，但零值仍然可以被正常使用。我觉得这句谚语更准确的说法应该是：Make the zero value us**able**.

另一个例子是标准库里的 `bytes.Buffer`：

```go
var b bytes.Buffer
b.Write([]buffer("Hello World"))
io.Copy(os.Stdout, &b)
```

字节缓冲区没有初始化，仅仅是声明过后就可以直接调用 `Write()` 方法使用。如果是在 Java 语言里，你大概需要 `new` 一个类的实例才行，或者用到别的构造方法，总之做不到声明即用。

Go 语言之所以强调「让零值可用」「声明即可用」，很大的一个原因是在 Go 语言里，空指针（`nil`）可以用来调用违背初始化的类型的方法。假设你有一个名为 `Config` 的结构体类型，你为它编写了 `Path()` 方法，你声明一个 `Config` 之后并没有指定路径，甚至 `Config` 本身就是空的，`Path()` 也可以返回一个默认值。

```go
func (c *Config) Path() string {
	if c == nil {
		return DEFAULT_STRING
	}
	return filepath.Join(USER_HOME, c.Filename)
}
```

简单来说，如果你可以直接使用零值，就不要把事情搞得太复杂！

## `interface{}` says nothing

这句话直接翻译过来是：空接口说明不了任何东西。

这里说的并不是在编写接口时没有编写任何方法这种明显的错误，而是指在声明类型时使用 `interface{}`。前面提到，Go 并非是面向对象编程语言，`interface` 是为类型服务而非对象服务的；在 Go 里面，实现方法不需要显式声明，只需要实现规定的方法就好了——没有方法的空接口，自然能匹配任何类型。

`interface{}` 常常被用来接收不确定类型的值，很多时候这种操作是必要的，但 Rob Pike 指出，很多人滥用了空接口类型，他说：

> You might as well write in Python.

将代码保持在 Go 严格的类型系统中，有助于保持程序的安全性。除此之外，空接口最大的问题是「不够清晰」——清晰的代码可读性是 Go 语言的核心之一。程序员在如果调用了一个函数，结果发现要传入的参数是一个空接口，即任意类型都可以，那他就要多花些时间思考，就算去读函数实现也要费些功夫。简单来说，使用空接口类型让函数的输入和行为都变得不可预测，代码更难阅读，而且更难发现 Bug。

在 Rob Pike 这场演讲之后发布的 Go 1.18 中，Go 语言添加了 `any` 类型，是 `interface{}` 的别名。添加这个别名的主要原因是让代码变得更可读，人们看到 `any` 就会知道在这里传入任何类型的值都可以，而不熟悉 Go 的人看到 `interface{}` 可能会觉得困惑。即使是 `any`，在使用前也要想清楚自己是不是真的没有办法预测数据类型，还是只是为了图省事而牺牲了安全性和可读性。

合理使用 `any` / `interface{}` 的例子是标准库里的 `fmt.Println()` 方法，程序员的确无法预知要打印的是字符串、数字还是别的什么类型。

总之，这条谚语让程序员严格遵守 Go 的类型系统。

## Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

这句话直接翻译过来是：gofmt 的样式不是任何人的最爱，但 gofmt 是所有人的最爱。

`gofmt`（或者 `go fmt`）是 Go 语言提供的命令行工具，用于格式化 Go 源代码。写完 Go 语言代码之后在项目根目录执行 `gofmt -s -w .` 就可以一键格式化代码。`gofmt` 的样式是统一的，没有自定义的空间，这看起来死板，让程序员不能用自己喜欢的风格格式化代码，但好处在于，有一个官方的格式意味着团队之间无需对代码风格做出讨论和决策，也不需要像 JavaScript 开发者一样纠结用 ESLint 还是 Prettier 等第三方代码格式化工具——用官方的 `gofmt` 就好了！这样一来，开发者就可以把心思放在更重要的事情上。

总而言之，不要纠结这个地方要不要用 Tab、那个地方要不要换行了，你的开发环境里已经有 `gofmt` 了。

## A little copying is better than a little dependency.

这句话直接翻译过来是：一点点复制比一点点依赖更好。

大部分程序员都在强调代码复用性（reusability），排斥重复造轮子，追捧 DRY（Don't Repeat Yourself）的理念，这导致许多程序员在编写简单的需求时也会使用大量的第三方依赖库，即时他们只调用其中的一两个方法，即时这些依赖库完全可以自己用三五行代码封装来代替。

Rob Pike 强调：要保持一个小巧干净的依赖树。维护依赖是痛苦的，因为依赖是**动态**的，如果依赖库有更新、出了问题，使用了这个依赖库的代码往往也需要跟进。这个教训，已经被 JavaScript 程序员吃尽了，JavaScript 庞大的生态意味着几乎没有一个 JavaScript 项目不大量使用依赖库，而一个 JavaScript 项目要是有几个月的时间没有更新，代码很可能就过时甚至用不了了。

1. 在大多数情况下都尽量依赖标准库；
2. 需要更复杂的工具时，优先考虑封装在自己的代码库里；
3. 如果几行代码就能解决，就不要引入依赖；
4. 如果考虑周全，得出结论认为一定要使用第三方库，这时候引入依赖才是最好的做法。

总而言之，头脑清晰的「C-V 工程师」比喜欢尝试各种库的「时髦工程师」写出来的代码更稳定！

## Syscall must always be guarded with build tags.

这句话直接翻译过来是：Syscall 必须总是被构建标签守护。

`syscall` 即系统调用，让程序从用户态访问内核态的操作，像操作系统发出请求，完成较高权限的读写。系统调用很强大，很多情况下也是必要的，尤其是在系统级开发下几乎必不可少，但系统调用的问题是：它不跨平台。

Windows、macOS 以及各种千奇百怪的 Linux 发行版，各自的 `syscall` 都有很大差异，几乎不能写出通用的代码。要保证全平台通用，就应该用 `os` 库；但如果必须用 `syscall`，就要用 Go 编译器的「构建标签」（build tags）。

构建标签实际上就是写在文件开头的一行注释：

```go
// go:build linux
```

上面这行代码给整个文件添加了 `linux` 的构建标签，表示只有在 Linux 操作系统下才编译这个文件。构建标签也可以有别的用处，比如软件有 Free 和 Pro 两个版本，Pro 版的功能就可以单独写在几个文件里，给这些文件加上 `pro` 的构建标签，然后在构建的时候这样编译：

```shell
go build -tags pro
# 在编译时手动加上 pro 标签
```

区分 `debug` 模式，区分生产环境和开发环境，也可以用到构建标签。

## Cgo must always be guarded with build tags.

这句话直接翻译过来是：Cgo 必须总是被构建标签守护。

Cgo（读作 Seagull，即海鸥）指的是 Go 语言标准库里的 `C` 库，这个库允许程序员在 Go 语言中调用 C 标准库，还可以在注释里写 C 代码并使用在 Go 代码中。

```go
// 例子来源：https://go.dev/wiki/cgo

/*
##include <stdio.h>
##include <stdlib.h>

void myprint(char* s) {
    printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func Example() {
    cs := C.CString("Hello from stdio\n")
    C.myprint(cs)
    C.free(unsafe.Pointer(cs))
}
```

和 `syscall` 一样，操作系统提供的 C 标准库实际上并非完全相同，C 语言并不跨平台。如果要用 Cgo，就必须使用构建标签声明这段代码是为哪个平台编写的。

## Cgo is not Go

这句话直接翻译过来是：Cgo 不是 Go。

这句话强调的是，Go 语言提供性能良好的垃圾回收机制、安全的类型系统、安全的操作内存的方式，编写纯粹的 Go 代码可以保证准确性、稳定性、安全性以及代码可读性。相反，C 语言本身并不是一门安全的语言，需要程序员手动 `free()` 释放内存。如果使用 Cgo，就会失去 Go 语言自带的各种安全机制，很容易出现难以排查的问题。

Rob Pike 建议，尽量避免使用 Cgo，你多半不需要用到它。

## With the unsafe package there are no guarantees.

这句话直接翻译过来是：用了 `unsafe` 包就没有保障。

在前面使用 Cgo 的例子里，就已经演示过了 Go 标准库提供的 `unsafe` 包。它的功能很强大，能绕过类型系统直接访问内存、把指针当作整数来用…… 能力越大，责任也越大。和 Cgo 一样，使用 `unsafe` 就等于抛弃了 Go 提供的安全机制——跑出安全区的样子好像很酷，但你为什么不直接去写 C 呢？明明用 `unsafe` 还麻烦不少。

Rob Pike 建议避免使用 `unsafe`，至少使用之前要想想自己是否真的需要用到这么强大又危险的功能，用的时候也要做好可能出错的准备。

## Clear is better than clever.

这句话直接翻译过来是：聪明比不过清晰。

这解释了 Go 语言为什么移除了三目表达式，原因就是它不够清晰。一般来说，三目表达式其实非常好用，在 C、Java 和 JavaScript 等热门语言中都存在。

```js
// ? 前面是一个布尔值或者布尔表达式
// ? 后面是为 true 的结果
// : 后面是为 false 的结果
console.log(OK ? "It's okay" : "It's not okay")
```

这看起来还挺清晰的，但 Go 的设计者发现有不少人会用三目表达式来编写令人费解的控制结构。

```js
const someString = isAnimal 
	? (isElephant ? "I'm an elephant" : "I'm just a dumb animal")
	: (CheckIfItIsAPlant() 
		? "I can't speak. I'm a plant" 
		: "Fuck, I'm not even a plant?")
```

三目表达式被滥用过后会大大降低代码可读性，Go 语言鼓励看起来很蠢，但清晰可读的写法。

```go
var someString string
if isAnimal {
	if isElephant {
		someString = "I'm an elephant"
	} else {
		someString = "I'm just a dumb animal"
	}
} else  {
	if CheckIfItIsAPlant() {
		someString = "I can't speak. I'm a plant"
	} else {
		someString = "Fuck, I'm not even a plant?"
	}
}
```

## Reflection is never clear.

## Errors are values.

## Don't just check errors, handle them gracefully.

## Design the architecture, name the components, document the details.

## Documentation is for users.

## Don't panic.

[^1]: C 语言里还有用信号量（Semaphore）实现互斥的方法，不过似乎已经没有什么人用了。

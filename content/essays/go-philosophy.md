---
title: Go 语言的编程哲学
tags:
  - 编程
  - 计算机
  - 设计
date: 2025-11-25 23:49:00
draft: false
toc: true
aliases:
  - /posts/go-philosophy
categories: Coding
---


从这个月 10 号，水逆期开始，我就开始减少出门，宅在家里读一读书，捡一捡法语学习，看看音乐剧，还找了些新东西来学，Go 语言就是其中一个。一开始只是抱着好奇心探索，并没有什么功利心，不过越学越发觉，Go 真的很适合用来写高并发高性能的 Web 应用（而且越学越讨厌 Java），也完全能胜任系统开发，就业岗位貌似不少，竞争也相对更小。更重要的是，我非常喜欢这种「观点鲜明」的编程语言，我想我会一直用下去的。

学任何一门编程语言我都更喜欢「不自量力」地直接上手写代码，一边写一边学基础语法，遇到不知道怎么实现的需求再查阅文档。不过，学 Go 的一开始我发现了 [Go Proverbs](https://go-proverbs.github.io/) 这个网站，便在实战之外还花了不少时间研读这里的「名言警句」。这些 Go 语言编程相关的谚语大多来自 Rob Pike 在 2015 年的[一场演讲](https://www.youtube.com/watch?v=PAAkCSZUG1c)，每一条都很有启发性，这篇文章就来逐条阐述我对它们的理解，其中有不少也可以用在 Go 语言以外的编程中，希望能引发读者的一些思考。<!--more-->

## Don't communicate by sharing memory, share memory by communicating.

这句话直接翻译过来是：不要用共享内存的方式通信，要用通信的方式共享内存。

这需要一些上下文来理解。Go 是一门可以通过非常优雅的方式实现并发的编程语言，异步编程若是要在多个线程之间同步信息，往往需要加锁，这是为了避免线程同时操作内存，导致冲突、数据准确性（integrity）被破坏等无法预测的问题。加锁一般指的是互斥锁（mutual exclusion object，简写作 `mutex`），线程操作数据时需要先获取对应的锁，如果锁被其他线程持有，则需要等待，获取到锁的线程才能执行操作，这保证数据操作是互斥的。

加锁是异步编程中的常见做法[^1]，但「锁是痛苦的」，不仅要创建锁、管理锁，还要想办法避免死锁等运行时错误，这类错误没法被编译器察觉，程序员排查起来很困难。所以，为了保证程序的健壮性和可读性，应当避免使用锁。线程之间之所以要用锁来实现互斥，是因为共享了同一块内存区域，而共享内存是为了传递数据，也就是这句话的前半句「communicate by sharing memory」。共享了内存就必须加锁，不然就会出现问题，而锁又会带来更多难题。

Go 语言鼓励的「share memory by communicating」，要用到 Go 语言的一个特性—— `channel`。你可以把 `channel` 理解为在不同线程之间传递数据的「通道」。以下是一个整数类型的 `channel` 例子。

```go
intChan := make(chan int)

// 线程 A 传入一个数据
intChan <- 10

// 线程 B 取出这个数据
someVar := <- intChan
```

严格来说，这两个线程仍然共享了 `intChan` 这个 `channel` 类型的变量，但整体上，这种写法避免了直接读写被储存在内存中的数据。在线程 A 没有传入数据时，线程 B 执行到 `<- intChan` 这一行就会自动阻塞，直到取到 `channel` 中的数据。在 `channel` 的**抽象**下，线程不再操作一块无人管理的，需要线程自觉加锁、等待用锁的数据区域，而是相互之间直接传递数据——这就是「share memory by communicating」。

Go 语言实现并发的方式非常优雅，而 `channel` 的存在又简化了线程之间共享数据的方式，这使得用 Go 编写并发程序的体验非常舒服，这一点在接下来还会看到。

## Concurrency is not parallelism.

这句话直接翻译过来是：并发不是并行。

这其实是计算机科学的基础知识，并发（concurrency）和并行（parallelism）本来就不是一回事。一般来说，并发指的是程序在处理器上交替执行，在单核处理器上，一个进程完全占用 CPU 资源，完成计算或者用完时间片之后，其运行所需的上下文会被保存，然后 CPU 进行上下文切换（context switch），换上下一个要执行的进程的上下文，然后运行下一个进程。每个进程执行时都完全占用一个处理器内核，时间到了之后就让下一个进程执行，只不过进程在 CPU 上切换的速度非常快，宏观上让人感觉是在同步执行。并行指的是真正的同步执行，在多核处理器上，不同的进程在不同的处理器内核上执行，而不需要交替执行。

Go 语言是跨平台的，不需要考虑硬件底层，也就是说，编程语言不关心执行程序的计算机究竟有多少个处理器核心。Go 语言强调并发并非并行的原因是，Go 把并发当作一种「编程结构」，即组织程序的方式。程序员在进行多线程编程时，应该聚焦于程序结构而非硬件底层，如果关注程序结构，就能写出更清晰、可读、可维护的多线程程序代码；而且，也不能理所应当地认为自己用了并发的结构，程序就一定是并行的。

Go 语言里实现并发非常简单，假设我们有一个 `FetchData()` 函数，而这个函数需要进行大量 IO 操作或者需要发送 HTTP 请求，如果同步执行会长时间阻塞主线程，要把这个函数变成并发的，只需要在调用函数时，在前面加一个关键词 `go`。这有点像 JavaScript 里的 `async`，但并不相同。

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

总结一下，关注程序结构，写出更清晰的代码，在逻辑结构上让代码异步执行，这是并发；至于并行，那与硬件底层和操作系统有关，不能认为按并发结构编写的代码就一定是并行的。

## Channels orchestrate; mutexes serialize.

这句话直接翻译过来是：Channels 编排；互斥锁串行。

尽管 Go 提供了在大部分时候都更好用的 `channel`，也更鼓励使用 `channel`，但互斥锁并不是被禁止的，也没有从这门语言里消失，因为有些时候必须要用到锁。这句话简单回答了「什么时候该用 channel？」「什么时候该用锁？」这两个问题，也能帮助 Go 程序员更好地理解 `channel` 的用途。

编排这个词很有意思，它的原文是 orchestration，又名 choreography，意思分别是「编配管弦乐曲」和「编舞」。如果你理解概念，就会觉得这个描述非常生动形象。这个词在不同的语境下的意思有一些区别，但一般指代一种自动化的工作流，可以统一调度、掌控多个作业，高效处理大量用户请求，或者让多个程序之间相互协作等等。在管弦乐（orchestration）中，各个演奏者各司其职，但由指挥家统一主导；在舞蹈（choreography）中，舞者之间需要相互沟通，分工协作。这引申到编程中，对应了不同的编程结构，下面用简单的例子来解释。

假设有一项操作需要大量计算，这些计算被分散开，启动多个 `goroutines` 执行，最后由主线程将计算的结果整合起来，Go 语言代码可以这样写：

```go
// 有 100 项作业需要执行，相应地有 100 项结果需要接收
// 分别创建 channel
jobs := make(chan int, 100)
results := make(chan int, 100)

// 启动 10 个 Worker 执行这些作业
for w := range 10 {
	go Worker(w, jobs, results)
}

// 传入 100 个作业给 jobs channel
// worker 接收到之后就会执行
// 执行完一个作业之后会继续接收 jobs channel 里的剩余作业
for j := range 100 {
	jobs <- j
}
close(jobs)

// 收集 100 个作业结果
for _ := range 100 {
	doSomethingWith(<- results)
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

总而言之，`channel` 用于高效、灵活地编排程序，而 `mutex`（互斥锁）用于安全、严格地保证串行。

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

Go 语言鼓励程序员好好利用这个零值，让类型即使没有初始化也能正常使用。标准库里有很多例子，前面提到的 `sync.Mutex` 就是一个。让我们回看这个例子：

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

Go 语言之所以强调「让零值可用」「声明即可用」，除了让代码更简单易读、编写程序更简单，另一个原因是在 Go 语言里，空指针（`nil`）可以用来调用未被初始化的类型的方法。假设你有一个名为 `Config` 的结构体类型，你为它编写了 `Path()` 方法，你声明一个 `Config` 之后并没有指定路径，甚至 `Config` 本身就是空的，`Path()` 也可以返回一个默认值。

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
> 
> 你还不如去写 Python 呢。[^2]

将代码保持在 Go 严格的类型系统中，有助于保持程序的安全性。除此之外，空接口最大的问题是「不够清晰」——清晰的代码可读性是 Go 语言的核心之一。使用空接口类型让函数的输入和行为都变得不可预测，代码更难阅读，而且更难发现 Bug。

在 Rob Pike 这场演讲的几年之后，Go 发布了 1.18 版本，其中添加了 `any` 类型，是 `interface{}` 的别名。添加这个别名的主要原因是让代码变得更可读，人们看到 `any` 就会知道在这里传入任何类型的值都可以，而不熟悉 Go 的人看到 `interface{}` 可能会觉得困惑。即使是 `any`，在使用前也要想清楚自己是不是真的没有办法预测数据类型，还是只是为了图省事而牺牲了安全性和可读性。

合理使用 `any` / `interface{}` 的例子是标准库里的 `fmt.Println()` 方法，程序员的确无法预知要打印的是字符串、数字还是别的什么类型。

总之，这条谚语让程序员严格遵守 Go 的类型系统。

## Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

这句话直接翻译过来是：gofmt 的样式不是任何人的最爱，但 gofmt 是所有人的最爱。

`gofmt`（或者 `go fmt`）是 Go 语言提供的命令行工具，用于格式化 Go 源代码。写完 Go 语言代码之后在项目根目录执行 `gofmt -s -w .` 就可以一键格式化代码。`gofmt` 的样式是统一的，没有自定义的空间，这看起来死板，让程序员不能用自己喜欢的风格格式化代码，但好处在于，有一个官方的格式意味着团队之间无需对代码风格做出讨论和决策，也不需要像 JavaScript 开发者一样纠结用 ESLint 还是 Prettier 等第三方代码格式化工具——用官方的 `gofmt` 就好了！这样一来，开发者就可以把心思放在更重要的事情上。

总而言之，不要纠结这个地方要不要用 Tab、那个地方要不要换行了，你的开发环境里已经有 `gofmt` 了。

## A little copying is better than a little dependency.

这句话直接翻译过来是：一点点复制比一点点依赖更好。

大部分程序员都在强调代码复用性（reusability），排斥重复造轮子，追捧 DRY（Don't Repeat Yourself）的理念，这导致许多程序员在编写简单的需求时也会使用大量的第三方依赖库，即使他们只调用其中的一两个方法，即使自己完全可以封装一个三五行代码的函数来实现相同的需求。

Rob Pike 强调：要保持一个小巧干净的依赖树。维护依赖是痛苦的，因为依赖是**动态**的，如果依赖库有更新、出了问题，使用了这个依赖库的代码往往也需要跟进。这个教训，已经被 JavaScript 程序员吃尽了，JavaScript 庞大的生态意味着几乎没有一个 JavaScript 项目不大量使用依赖库，而一个 JavaScript 项目要是有几个月的时间没有更新，代码很可能就过时甚至用不了了。

所以：

1. 在大多数情况下都尽量依赖标准库；
2. 需要更复杂的工具时，优先考虑封装在自己的代码库里；
3. 如果几行代码就能解决，就不要引入依赖；
4. 如果考虑周全，得出结论认为一定要使用第三方库，这时候才引入依赖。

总而言之，**头脑清晰**的「C-V 工程师」比喜欢尝试各种库的「时髦工程师」写出来的代码更稳定！

## Syscall must always be guarded with build tags.

这句话直接翻译过来是：Syscall 必须总是被构建标签守护。

`syscall` 即系统调用，允许程序从用户态访问内核态，像操作系统发出请求，完成需要系统级权限的操作。系统调用很强大，很多情况下也是必要的，尤其是在系统级开发下几乎必不可少，但系统调用的问题是：它不跨平台。

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

Cgo（读音类似 Seagull）指的是 Go 语言标准库里的 `C` 包，这个包提供的方法允许程序员在 Go 语言中调用 C 标准库，还可以在注释里写 C 代码并使用在 Go 代码中。

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

这句话强调的是，Go 语言提供性能良好的垃圾回收机制、安全的类型系统、安全的操作内存的方式，编写纯粹的 Go 代码可以保证准确性、稳定性、安全性以及代码可读性。相反，C 语言本身并不是一门安全的语言，需要程序员手动 `free()` 释放内存，容易出现段错误等等。如果使用 Cgo，就会失去 Go 语言自带的各种安全机制，很容易出现难以排查的问题。

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

还有那些过度复杂的条件表达式，看起来很聪明，但可读性非常差。

```go
// ❌

if thisVar <= thatVar || (isLikeThis && fetchStatus() != status.OK) {
	return
}

// ✅

if thisVar <= thatVar {
	return
}

statusOK := (fetchStatus() != status.OK)
if isLikeThis && !statusOK {
	return
}
```

更直接、简单、清晰的代码不仅提高可读性，让维护代码变得更容易（这对自己和团队成员都是好事），也变相提高了代码稳定性，因为清晰的代码更容易排查问题。

执拗地写聪明但不清晰的代码有点像某种大男子主义的攀比心，像是「我昨天熬夜到凌晨四点把工作做完了！」，许多人都会为这种刻苦但实际上是耍帅的精神鼓掌，然而，人的认知能力在凌晨四点到六点比喝醉了酒还要差[^3]，而且熬夜对身体的伤害不必多说。

## Reflection is never clear.

这句话直接翻译过来是：反射永远不清晰。

反射（Reflection）是很多高级语言里都有的机制，用于在程序**运行时**检视和操作对象。有的信息在程序运行之前无法预知，要在程序中使用这些不可预知的值，就可以用到反射；比如前面提到的空接口，如果函数接收一个空接口类型（任意类型）的值，但需要根据值的类型执行操作，就可以利用反射来获取这个值的类型。

反射也带点 Meta（元）的意味，经常被称作 Metaprogramming（元编程），因为它可以被用来检视程序自身的结构，比如 Java 语言中就可以利用反射机制获取某个类的注解、构造函数、成员变量等等，并把它们当作普通的对象来使用。

在 Go 语言提供的 `reflect` 包里，大致有以下几种方法：

```go
// 接收一个 interface{} 类型的值，即任意类型的值
func doSomeReflection(i interface{}) {
	// 获取 i 的类型
	// 一个定义在 main 包里、名为 User 的结构体会返回 main.User
	t := reflect.TypeOf(i)
	
	// 获取 i 的底层类型
	// 同样一个名为 User 的结构体，会返回 struct
	k := reflect.KindOf(i)
	
	// 获取 i 的值
	v := reflect.ValueOf(i)
	// 如果要获取值的类型，就可以这样写
	vk := reflect.ValueOf(i).Kind()
	
	// 如果 i 是一个结构体，可以用这个方法获取它的字段数量
	nf := reflect.NumField(i)
	
	// Field() 方法用于获取字段，配合上面的方法可以遍历结构体
	for n := 0; n < numberOfFields; n++ {
		// 这里会输出该字段的类型和值
		fmt.Printf("Type:%T, Value:%#v", value.Field(n), value.Field(n))
	    // 这里输出的是该字段的底层类型
		fmt.Println("Kind is ", value.Field(i).Kind())
	}
}
```

至于 Rob Pike 为什么说「反射不够清晰」，我想读完上面的代码已经不言自明了，反射就是很难读。我认为最根本的原因是，在代码中引入反射，迫使程序员用另一个完全不同的视角（元视角）来审视代码，这种视角的转换是很困难的，而且往往要同时使用常规的编程逻辑和元编程逻辑才能读懂反射。

这并不是说反射不好，有的需求必须要用反射实现。只不过，清晰永远比聪明更好，如果有更清晰的、不涉及反射的方法，那就不要用反射。

## Errors are values.

这句话直接翻译过来是：错误是值。

接下来我们正式进入到 Go 的错误处理哲学。Go 语言完全抛弃了传统的 `try ... catch` 结构，将错误简化为值，将错误处理从一种特殊的控制接口降级为简单的 `if ... else`。Go 还支持多函数返回值，也就是说，你可以将错误当成一个返回值返回给调用者，而不需要像 Java 一样抛出（throw）。

在 Go 的错误处理中，程序员最常写的就是这样一段代码：

```go
result, err := GetSomeData()
if err != nil {
	// 在这里处理错误
}
```

如果是 Java 等依赖 `try ... catch` 进行错误（异常）处理的语言，你需要这样写。

```java
try {
	// 其他相关的代码
	ResultType result = GetSomeData()
	// 继续其他相关的代码
	// ...
} catch (IOException e) {
	// 处理 IO 异常
} catch (Exception e) {
	// 处理其他类型的异常
}
```

这意味着你需要把可能会抛出错误（异常）的代码提前包裹在 `try ... catch` 块中，而且你还创建一个新的作用域（scope），如果你在 `try ... catch` 里面获取了数据，就得直接在里面处理这个数据，或者在这个作用域外面先声明函数，再在里面使用。总之是有点麻烦的。更不用说，在 Java 里面你还需要考虑不同类型的异常，光是想到这些我都不想做错误处理了。

Go 语言将错误作为值，不仅让错误处理变得简单直观，减少了一种控制结构，还让程序变得更灵活——既然错误是值，那你就可以把错误储存起来，像普通的值一样直接写进日志里，放进一个 map 里，缓存起来，或者等完成一大批操作之后统一反馈给用户，甚至发送给你的妈妈看——最后一条是 Rob Pike 在演讲时开的玩笑。其他语言也可以做到，但在 Go 里面会更容易，更重要的是，你可以对错误进行编程，而不会被局限在 `try ... catch` 这一种控制结构里。

总之，错误在 Go 语言里是值，你可以用它干任何事情！

## Don't just check errors, handle them gracefully.

这句话直接翻译过来是：不要只是检查错误，要优雅地处理错误。

所谓的「gracefully」（优雅地）其实只是「intentionally」（有意识地），也就是不理所当然地使用一种错误处理方式，不理所当然地觉得「发生错误了，我直接输出错误信息就好了」，而应该思考：我应该不应该增加更多上下文信息？应该不应该保存这个错误，做完其他事情之后再打印出来？

以下是写在[官方文档](https://go.dev/doc/effective_go#errors)里的，处理错误的例子：

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

`for` 循环的目的是尝试创建文件两次，如果第一次成功（`err == nil`），就 `return` 跳出循环，不执行第二次；如果第一次失败，那就对 `err` 进行类型断言（type assertion），判断这是不是一个 `PathError`，并进一步对错误进行检查，如果发现是空间不足导致的错误，就调用 `deleteTempFiles()` 函数删除临时文件恢复一些空间，然后 `continue` 进行下一次循环，再次尝试创建文件。

总而言之，不要把错误抛出来之后就真的抛之脑后了，想想自己能对这个错误做些什么！

## Don't panic.

这句话并不是 Rob Pike 在前文提及的演讲中发表的，而是写在 Go 官方文档里的[一句话](https://go.dev/wiki/CodeReviewComments#dont-panic)，由于它和前面两小节讨论的错误处理相关，我就插到了这里。

直接翻译过来是：不要恐慌（panic）。这里的 `panic` 指的是 Go 语言中用于错误处理的一个函数，表示程序遇到了**意料之外**且无法处理的错误，必须结束程序。这句话很直接地告诉你：不要使用 `panic()` 函数。

文档很明确地告诉 Go 程序员，不要把 `panic()` 用来处理常规错误，任何能被预料的、有办法处理的（诊断原因并打印错误也是一种处理），都应该使用 `error` 类型的值，并将这个值作为函数的额外返回值返回给调用者，然后调用者再检查 `if err != nil` 来处理错误。原因在于，`panic()` 会直接结束程序，只有在程序真的完全没有办法继续执行，发生了无法预料的致命错误时，才应该使用 `panic()`。

即使是这样，`panic()` 也可以被恢复，即使程序被 `panic()` 停止，`defer` 的函数也会执行（前文有提及），程序员可以在 `defer` 函数中使用 `recover()` 函数恢复错误。以下是官方文档中的例子。

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

简单来说，除非真的没办法处理，否则不要 `panic()`。

## Design the architecture, name the components, document the details.

这句话直接翻译过来是：设计架构，命名组件，用文档展开细节。

比起编程语言的基础知识，这条更像是软件工程的经验之谈，Rob Pike 说这是在 Go 语言中设计大型软件系统的方式，不过这对任何语言来说似乎都大有裨益。这句话总结了系统设计的三步骤：

1. 设计架构：宏观上，完成这个软件系统，我需要哪些组件（模块）？
2. 命名组件：给每个组件起个好名字，清晰的名字会让系统更容易被理解，最好是让人看到名字之后就立马能知道各个组件的职责。
3. 用文档展开细节：就像命名变量和函数一样，清晰直观固然能提升可读性，但也免不了要写注释。一些名字说不清楚的细节，就在文档里写清楚。

## Documentation is for users.

这句话直接翻译过来是：文档是给用户的。

> A lot of time when people write documentation, they say "this is what the function does", and they don't think about "what the function is for".
> 
> 很多时候人们写文档，会写「这个函数会做这个事情」，而他们不会想「这个函数是拿来干嘛的」。

这条也更像是工程智慧，而非语言的基础知识。在我的理解力，「这个函数是拿来干嘛的」其实是提供了相关的上下文，即「为什么一开始要设计这个函数」「为什么用户会用到这个函数」。在上一小节的例子中，`safelyDo()` 函数封装了 `do()` 函数，用于恢复可能出现的 `panic`，如果只考虑「做什么」，文档就应该是这样：

> `safelyDo()` 调用 `do()`，并恢复潜在的 `panic`。

但用户可能会感到困惑：我为什么会需要恢复潜在的 `panic`？不恢复会有什么问题吗？如果这样写，就会清晰不少：

> 由于 `do()` 可能会触发不可恢复的 `panic` 错误，为了保证程序能继续执行，`safelyDo()` 会恢复错误并不中断程序。

---

## 我的想法

总结下来，我认为 Go 语言最鲜明的几条观点是：

1. ==代码的可读性和可维护性比开发效率更重要。== 不要贪快写聪明的代码而忽略了其他人能不能读懂、几个月后的自己能不能读懂，不要因为自己的好恶影响了代码的风格，不要因为不愿意复制粘贴就引入动态的依赖，不要因为方便就滥用强大但不安全且不跨平台的 `syscall`、Cgo 或 `unsafe`，减少使用不清晰的反射。我认为这是**反技术债务**且**长期主义**的，这避免了给未来的自己和开发团队留下隐患、增加工作量。
2. ==程序员必须要有清晰的头脑，而且不能把任何事情当成理所当然，要思考每一个决策。== 使用 `interface{}` 接收任意类型很好用，但自己真的无法确定参数的类型吗？直接打印错误看起来很常规，但自己有没有办法恢复错误，或者打印错误时能不能提供更清晰的上下文信息以供排查问题？我要实现的需求真的只能用危险的 `unsafe` 或 Cgo 实现吗，如果一定要脱离 Go 的安全区，我愿不愿意接受代码随时可能崩溃的代价？代码复用和不重复造轮子真的是任何时候都必须遵守的铁律吗，这个涉及第三方库的需求我真的不能自己写吗？
3. ==不要轻易离开 Go 的保护。== Go 有自带的垃圾回收机制、安全的内存访问和各种安全机制，使用 Cgo 和 `unsafe` 就失去了 Go 的保障。
4. ==保持代码的灵活性。== 无论是小巧的、积木式的接口、「错误是值」的思想，还是保证零值可用的做法，在我看来都是为灵活性服务的。小接口让抽象更强大，类型系统更灵活；错误是值意味着可以对错误做任何对值能做的事情；零值可用意味着无论有没有初始化，无论变量是不是空的，方法都可以被正常地调用。

目前看来，Go 就是一门既安全又灵活，写代码既优雅又清晰，没有痛苦且观点鲜明的编程语言。Go 能用更少的代码写出更健壮的程序，很难令人不爱。

[^1]: C 语言里还有用信号量（Semaphore）实现互斥的方法，不过似乎已经没有什么人用了。

[^2]: 此处并非嘲讽 Python 语言和 Python 程序员，这么说是因为 Python 是一门非常灵活的**动态类型语言**。Rob Pike 的意思应该是：滥用 `interface{}` 就像是抛弃了类型系统，把 Go 当作动态类型语言来说，如果这么做的话，用真的动态类型语言会更好。

[^3]: 来源：《[绝佳时间](/library/2025/绝佳时间/)》

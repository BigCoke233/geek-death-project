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

学任何一门编程语言我都更喜欢直接上手写代码，一边写一边学基础语法，遇到不知道怎么实现的需求时再查阅文档。不过，这次学 Go 的一开始我发现了 [Go Proverbs](https://go-proverbs.github.io/) 这个网站，便在实战之外还花了不少时间阅读和理解这里的「名言警句」。这些 Go 语言编程相关的谚语大多来自 Rob Pike 在 2015 年的[一场演讲](https://www.youtube.com/watch?v=PAAkCSZUG1c)，每一条都很有启发性，这篇文章就来逐条阐述我对它们的理解，其中有不少也可以用在 Go 语言以外的编程下。<!--more-->

## Don't communicate by sharing memory, share memory by communicating.

这句话直接翻译过来是：不要用共享内存的方式通信，用通信的方式共享内存。

这其实有些令人费解，需要一些上下文来理解。Go 是一门可以通过非常优雅的方式实现并发的编程语言，异步编程若是要在多个线程之间同步信息，往往需要加锁，这是为了避免线程同时操作内存，导致冲突、数据准确性（integrity）被破坏等无法预测的问题。加锁一般指的是互斥锁（mutual exclusion object，简写作 `mutex`），线程操作数据时需要先获取对应的锁，如果锁被其他线程持有，则需要等待，获取到锁的线程才能执行操作，这保证数据操作是互斥的。

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

假设有一项操作需要大量计算，这些计算被分散开，启动多个 `goroutines` 执行，最后由主线程将计算的结果整合起来，Go 语言代码可以这样写。

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

这是一个典型的 Worker Pool。这个例子中，主线程就是指挥，而 Workers 就是乐队成员，整体上，这就是一种编排。

假设有三个作业（A、B、C），C 作业依赖 A 和 B 的结果，而 B 依赖 A 的结果，

[^1]: C 语言里还有用信号量（Semaphore）实现互斥的方法，不过似乎已经没有什么人用了。

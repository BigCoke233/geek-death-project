---
title: 稻草人周刊 Vol.64
subtitle:
tags:
date: 2026-01-12 11:39:31
categories: 稻草人周刊
draft: true
toc: true
toot:
---

{{< music cover="" title="Satellite" artist="Harry Styles" apple="https://music.apple.com/cn/album/satellite/1615584999?i=1615585100&l=en-GB" >}}

我很喜欢《Harry's House》这张专辑，虽然已经发布快三年了，还是经常拿出来听。这周 Apple Music 的精选集里推荐了这首《Satellite》，歌词的内容没有太多感同身受，但有点希望自己是歌里的那个「You」，知道有个人在等着我敞开心扉，而不是在小心翼翼地探出头之后受到伤害，这很令人安心。

<!--more-->

## ⋈︎ · 连接

### 2026 是自托管之年

> [!note] 📜
> [2026 is the Year of Self-hosting](https://fulghum.io/self-hosting)

作者表示，现在够买一个小巧的 Linux PC 已经很便宜，网络层面也有 Tailscale 提供支持，再加上 Claude Code 等基于命令行的 Agent 已经非常智能，那些缺少计算机系统和运维知识的人也能很方便地配置和使用自托管应用。你只需要买一个迷你主机放在家里，在这台主机和你的电脑上安装并配置 Tailscale，用你的电脑 SSH 到这台作为家庭服务器的主机上安装 Claude Code，然后告诉它你想要做什么就好了。

我自己也有一台利用率不高的 Mac Mini，还是 M4 芯片的，完全可以用来做家庭服务器。今年我也在逐步转向开源软件和自托管服务，我应该会参照作者的方案自托管密码管理器、照片管理和流媒体服务等。

### 全能型输入框

> [!note] 📜
> [A Box of Many Inputs](https://allenpike.com/2025/a-box-of-many-inputs/)

尽管我不使用并且抗拒使用任何现有的 AI 浏览器（在我看来开发 AI 浏览器是拿着锤子找钉子的做法，而且对隐私没有任何保障），不过这类浏览器中的地址栏设计很有意思。不少现代软件都内置了一个可以随时打开的输入框（常见的有 VS Code 里的 ⌘P 打开文件选择窗口，Obsidian 也可以用同样的快捷键打开命令选择窗口，Raycast 这类启动器软件更是提供了操作系统级别的全局输入框），这些输入框可以根据用户输入的内容推断意图，并提供合适的选项。AI 浏览器的地址栏也做了相同的处理，一些输入会被推断为问题，会直接打开 AI 对话框；另一些输入则被当作常规的搜索关键词。

作者用 Dia 浏览器搜索名为 *Who Framed Roger Rabbit?* 的电影，因为电影名本身是一个问句，结果被视作问题触发了 AI 对话——作者相信没有人在搜索电影的时候想要得到这样的结果。在 ChatGPT 的 Atlas 这边，这个问题并不会触发对话，而是会进入常规的搜索结果页面。区别在于，Dia 浏览器使用一个本地的小模型判断用户输入的内容到底是问题还是搜索关键词，而 Atlas 则简单粗暴地把 10 个单词以内的输入都当作常规搜索。（我很好奇这种判断法对中文的兼容性）

作者期待在未来，AI 浏览器会把输入框做得越来越好，能够更准确判断用户的意图。可我觉得这是设计的大忌。在 *The Design of Everyday Things* 这本书里，作者花了很长的篇幅批评以前的电话系统，座机上的一个按键能干的事情太多了，用户根本找不到某些功能的入口，不仔细拿说明书钻研，甚至会以为这个电话没有呼叫保持（Hold）功能。清晰比聪明更好，AI 对话和常规的搜索应该分开，至少设置明确的、不同的触发方式，比如回车是搜索，而 ⌘ 加上回车就是 AI 对话。用户使用软件的习惯不同，再精确的匹配和判断总有例外。不应该替用户做决定，而是要让用户自己做决定。

### Wikipedia 25 周年

> [!note] 🌐
> [25 Year of Wikipedia](https://wikipedia25.org/)

2026 年 1 月 15 日是维基百科的 25 岁生日。Wikipedia 发布了一个网站讲述这 25 年的发展历程，涉及技术、编辑者社区、内容政策的变更等多个方面的内容，读一读还是很有意思的。页面的最后是一个小册试题「你的 Wikiepdia 未来是什么样的？」，题目有点意思，不像是随便做着玩的。我得到的结果是 *The Consensus-Driven Collaborator*（共识驱动的合作者）。

> Welcome to the Knowl-Age
>
> You're super curious, friendly yet argumentative, but not in a bad way.
>
> Your future is shaped by open source technologies, open discourse, and a lot of behind-the-scenes teamwork. AI might be part of your future, as long as it helps make life better for all of humanity-not just select individuals.
>
> All this we see in your future: A fairer world where the internet belongs to everyone and maybe-just maybe-a Wikipedia Virtual Reality edition. Don't mind us, we will be doing the cha-cha from Wikipedia's "List of dance styles". Oh, and by the way, since you got this as a result, consider yourself officially invited to Wikipedia's 50th birthday!



## ⁂ · 星群

### AeroSpace

macOS 上的一个平铺窗口管理器。所谓平铺式窗口，就是「不让窗口互相重叠」。这类软件极大地限制了桌面的灵活度，但带来了高效的操作和绝对的整洁。AeroSpace 使用 Vim 键位操作窗口，按下 `⌥ h` 可以将焦点移动到左侧的窗口，`j` `k` `l` 也同理；在这个操作之上加上 `Shift` 键，就可以将窗口移动到指定的方位。把 `h` `j` `k` `l` 换成数字或其他字母，可以对「工作区」（workspace）进行管理，切换到某个工作区或将聚焦的窗口移动到某个工作区。

使用 AeroSpace 可以很方便地排列窗口，更有用的是，形成自己的习惯之后能够快速地在不同使用场景之前切换。比如我的 1 号工作区是运行 NeoVim 的终端和另一个用来干杂货的终端，我的 S 号（Social）工作区打开了 Mastodon 客户端、Matrix 客户端和 NetNewsWire（RSS 阅读器），D 号（Dashboard）工作区是打开了 Cloudflare 仪表盘、Spaceship 管理界面和各种后台的网站。

如果你也是 macOS 用户，我非常推荐这个软件，能够极大地提高效率。

访问：[AeroSpace](https://github.com/nikitabobko/AeroSpace)

## ↯ · 当下




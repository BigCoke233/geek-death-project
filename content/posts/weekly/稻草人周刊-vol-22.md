---
title: 稻草人周刊 Vol.22
subtitle: Brighter days ahead
description: 红与黑、自然的食欲、社会认可与个人目标、从 Git Commit 看高效沟通、如何制作小而美的游戏、🍎 Mac Mini
tags:
  - 健康
  - 习惯
  - 社会学
  - 软件
  - 沟通表达
  - 游戏开发
related_cards:
  - social-atomism
date: 2025-03-31 13:58:35
categories: 稻草人周刊
banner: https://image.guhub.cn//uPic/jez-timms-HauxSOFvh6k-unsplash.jpg
draft: false
toc: true
---

## 🙋 在做什么

Mac Mini 到手的第一件事居然是……

{{< gallery >}}
{{< img src="https://image.guhub.cn/uPic/IMAGE%202025-03-31%2013:32:40.jpg" >}}
{{< img src="https://image.guhub.cn//uPic/IMG_3134大.jpeg" >}}
{{< img src="https://image.guhub.cn//uPic/IMAGE%202025-03-31%2013:32:44.jpg" >}}
{{< img src="https://image.guhub.cn//uPic/IMAGE%202025-03-31%2013:32:46.jpg" >}}
{{</ gallery >}}

在听 Ariana Grande 的新专辑，《eternal sunshine》的豪华版，加曲《dandelion》很对胃口，这几天一直在循环。

<iframe allow="autoplay *; encrypted-media *;" frameborder="0" height="150" style="width:100%;max-width:660px;overflow:hidden;background:transparent;" sandbox="allow-forms allow-popups allow-same-origin allow-scripts allow-storage-access-by-user-activation allow-top-navigation-by-user-activation" src="https://embed.music.apple.com/cn/album/dandelion/1800579610?i=1800580284"></iframe>

## 📖 本周阅读

{{< bookcard "红与黑" >}}

---

## 💬 本期杂谈

### 食欲：从消费者到劳动者

![](https://image.guhub.cn//uPic/jV4Izn.jpg)

{{< callout "播客链接：[Apple Podcast](https://podcasts.apple.com/cn/podcast/%E7%8B%AC%E6%A0%91%E4%B8%8D%E6%88%90%E6%9E%97/id1711052890?i=1000701381415) / [小宇宙](https://www.xiaoyuzhoufm.com/episode/67e83d2c6ea60022359ddd33?utm_source=rss)" "🔗" >}}

下课回来的路上听了一期有意思的播客，其中「重新审视自己与食物的关系」这一观点让我很受启发。当自己需要花很多时间和精力才能吃到令自己垂涎欲滴的食物时，不需要刻意的饮食计划就能形成自然的食欲，因为自己离食物的因果链更近了一步，食物不再是自己动动手指就能让人送到门口的高热量精加工食品，而是自己真是的劳动成果。自己仍然可以贪吃，但贪吃的前提是克服懒惰。

我很喜欢这种自然随性的感觉。尽管我在宿舍连烧摩卡壶的炉子都不敢用[^1]，更别提自己做饭了，但我发觉自己在大部分时候食欲是相对健康的。至少，在一个暑假高强度减肥瘦了 40 斤之后，在毫无饮食计划，跑步也总是断断续续的情况下，我又瘦了 20 多斤，体重稳定地保持在 140～150 斤。[^2]

发觉自己食欲没那么自然，大概是意识到自己有一种不饿也不馋，但下课、下班之后就往超市、小卖部和烧烤店跑的习惯。==这大概是因为在城市生活的现代人难以分清食欲与习惯吧==。

[#健康](/tags/健康/) [#习惯](/tags/习惯/)

### 分清「社会认可」与「个人目标」

{{< callout "文章链接：[Don't Compete](https://invertedpassion.com/dont-compete/)" "📜" >}}

- 因为世界被「赢家获得一切」的零和思想统治，对大多数人来说，获得新闻和社交媒体里看到的那种成功，概率近乎为零。
- 仔细思考的话，社会结构其实很像彩票。成千上万的人花了大量的时间和他们的血汗钱，投入到这场赌博当中，像是交税一样为其他竞争者铺平道路，仅仅追求那微小的成功概率。
- 不要把社会认可与对你自己有利的东西搞混了。

作者的思想乍一看和「社会原子论」很像，认为个体在投入社会进步这场声势浩大的、赌博式的大赛中，也要停下来思考自己的愿望和利益，而不是被社会大流推着走，在潮流中获得的社会认可比起自己的幸福，其实无足轻重。

{{< card "social-atomism" >}}

做自己感到满意的事情，而不是大多数人认为最有用的，其实能保证自己的幸福，在大局上也能推动社会进步。

[#社会学](/tags/社会学/)

### 从 Git Commit 看高效沟通的方法

{{< callout "文章链接：[No Longer My Favorite Git Commit](https://mtlynch.io/no-longer-my-favorite-git-commit/)、[Every line of code is always documented](https://mislav.net/2014/02/hidden-documentation/)" "📜" >}}

Git 是一个版本管理软件，常用于软件开发。Commit 是 Git 的一个功能，当程序员完成一次修改之后，就可以 `git commit` 提交者次修改，修改时需要附上信息，用文字描述这次更改。

为了更好的协作开发，commit 要怎么写有很多说法。我自己用的是 [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) 的格式，像这样：

```
<type>[optional scope]: <description>
```

不过我也愈发觉得一句话完全不足以描述一次 commit 做出的改动，而且，如 Mislav Marohnić 在《[Every line of code is always documented](https://mislav.net/2014/02/hidden-documentation/)》这篇文章中指出的：每行代码都因为提交时编写的信息而有迹可循。好的 commit 信息应该要让后来者明白自己为什么要这么改。

而 Michael Lynch 在《[No Longer My Favorite Git Commit](https://mtlynch.io/no-longer-my-favorite-git-commit/)》中指出事无巨细地描述修改过程并不是好的软件工程，并展示了一个「倒金字塔」模型。

<img src="https://image.guhub.cn//uPic/inverted-pyramid.png" class="invert-if-dark center">

倒金字塔顶端是「所有人都在乎的东西」，其次是「一些人会在乎的东西」，最次是「很少人在乎的东西」。倒金字塔既表示了顺序，也表示了占比。也就是说，所有人都在乎的东西应该放在最前面，且应该是最多的；很少有人在乎的东西应该放在最后，而且应该尽可能言简意赅。

不能事无巨细，也不能总是简单掠过，搞清楚主次会让沟通更有效率。

[#软件](/tags/软件/) [#沟通表达](/tags/沟通表达/)

### 如何制作小而美的游戏

{{< callout "文章链接：[How To Make Good Small Games](https://farawaytimes.blogspot.com/2023/02/how-to-make-good-small-games.html)" "📜" >}}

- 如果刚开始做游戏就想一口吃成个大胖子，只有两种结果：要么放弃，要么在一个游戏引擎上耗了五年，指望跟掷骰子连着掷出13次六点一样低的概率、奇迹般地做出了自己梦中情游——这几乎不会发生。
- 从你的宏大计划中取出一小块，做成一个小游戏，以一种意想不到的方式实现你的梦想。
- 不要把做小游戏想成是你实现梦想之前所做的作业，不要说「做得好不好我都不在意」这样的话，因为这些小游戏也应该是值得你骄傲的作品。

关于「如何制作小而美的游戏」，作者给出了十二条建议。文章很长，可以选择感兴趣的条目阅读，不过要注意有的建议和前后文是相关的。

1. 游戏的质量独立于其规模。
2. 游戏的质量独立于其情感程度。
    - 并不是让人感动得流眼泪、产生巨大的情感波动的游戏体验才算得上好。
    - 学会欣赏微妙的艺术（subtle art），既能让自己成为更好的游戏开发者，也能让自己成为涉猎更广泛的玩家。
3. 游戏的成功是自己定义的。
4. 游戏的成功取决于是否实现了自己的承诺。
5. 承诺越小，游戏越容易成功。
6. 用一种有趣且令人愉悦的方式实现你的承诺。
7. 不要超出你承诺的内容。
8. 形式因素也是游戏承诺的一部分。
    - 你的游戏能在浏览器上运行吗？加载要花多久？手机上能不能玩？有很多文字吗？还是什么都没有？
    - 这些形式而非内容上的因素，也会极大程度上影响你的玩家。
9. 发布游戏之前请完成它。
    - 在游戏最后写上「未完待续」是未能完成自己承诺的表现，在标注 Demo 或者序集这样的字样来为自己设防之前，想想你为什么一定要现在发布它？
    - 找你的朋友玩一玩来获得反馈，会比把未完成的内容扔给陌生人玩要好得多。
    - 在一开始就做出更小的承诺。
10. 游戏系列并非捷径。
    - 好的游戏系列也应该是由多个独立的、完整的故事组成的。
11. 不要考虑游戏能不能火。
12. “不错”就很不错了（Good is good enough）。
13. 玩的开心 :)

作者还推荐了两个小巧的游戏引擎，[Bitsy](https://bitsy.org) 和 [Twine](https://twinery.org)，前者的限制不小，但两个都很适合用来交互式叙事。如果不需要自己编写高度定制化的游戏机制，不失为好的选择。

[#游戏开发](/tags/游戏开发/)

---

## 🗣️ 碎碎念

- 在书店站着读完了《我胆小如鼠》，是第一篇让我产生很强共鸣的余华的作品。一群男人聚集在一起欺负老实人，倒有点像上野千鹤子在《厌女》里讲的「男性同性社会同盟」，暗含着一种对「被当作女人」的恐惧，尤其是吕前进弯着腰以为自己要被强奸了，大喊「我是男人！我是男人！」的那一段，把男人对女性气质的厌恶和“沦为”女性的恐惧表现得淋漓尽致。这兴许与「主体性丧失」也有关系。
- 之前用 M1 的 MacBook Air 玩饥荒联机版，有很明显的卡顿，本以为是性能问题，用上 M4 的 Mac Mini 就能解决，结果才发现是饥荒洞穴服务器导致的延迟，在 Mac Mini 上也有，关掉洞穴服务器就不会有这个问题…… 之前在 Windows 平台上没有遇到过这么明显的延迟，不知道是不是处理器架构或者其他操作系统设计的差异导致的。总之，用其他方法优化之后还是能正常开启洞穴游玩的。[^3]
- 终于可以把我的暗影精灵塞箱底了，换上 Mac Mini 之后，桌面的走线马上干净了，不需要找地方塞那块砖头一样大的电源适配器。

---

封面图来源：[Jez Timms @ unsplash](https://unsplash.com/photos/a-spotlight-coming-from-a-hole-in-a-dark-underground-cave-in-minorca-HauxSOFvh6k)

[^1]: 除了因为宿舍禁止使用大功率电器，还因为我怕摩卡壶会把咖啡喷出来，把我桌面上的插排、电脑和本子搞得乱七八糟…… 地方小只能忍忍了。
[^2]: 不过，让我苦恼的一点是，尽管体重健康，但自己腰上仍然有很多赘肉，这大概也与相对不那么自然的饮食习惯和运动习惯有关。
[^3]: ~~这么说，我岂不是没必要为了玩游戏买 Mac Mini？~~罢了罢了，反正都蹭上优惠了。

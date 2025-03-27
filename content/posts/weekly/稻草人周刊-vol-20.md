---
title: 稻草人周刊 Vol.20
subtitle: Heart Hunter 发布
description: Heart Hunter、假装看手机指南、收藏知识与否、逃离推荐算法、骗人的进度条、简单不易过时的技术、缪斯的死亡、View Transition
tags:
  - Web
  - 游戏开发
  - 互联网文化
date: 2025-03-17 15:50:09
categories: 稻草人周刊
draft: false
toc: true
aliases:
  - /posts/稻草人周刊-vol-20/
---

{{< callout "你正在阅读 稻草人周刊，一个读者的自述，每周一更新。" "🔖" >}}

## 🎧 这周在听什么？

<iframe allow="autoplay *; encrypted-media *;" frameborder="0" height="450" style="width:100%;max-width:660px;overflow:hidden;background:transparent;" sandbox="allow-forms allow-popups allow-same-origin allow-scripts allow-storage-access-by-user-activation allow-top-navigation-by-user-activation" src="https://embed.music.apple.com/cn/album/pure-heroine/1440818584"></iframe>

听了听 Lorde 的老专辑（因为《Melodrama》快听包浆了）。

## 🧑‍💻 这周做了什么？

我的游戏《Heart Hunter》已经在 itch.io 上发布了！

<iframe frameborder="0" src="https://itch.io/embed/3402746" width="552" height="167"><a href="https://eltrac.itch.io/heart-hunter">Heart Hunter by Eltrac</a></iframe>

这周重构了代码，用上了 Box2D 物理引擎，让移动和碰撞更丝滑了；还添加了新的生物，调整了关卡设计和生物数据，增加了近战攻击，增加了小地图等等。最后因为碰上水逆怎么也修不好地图生成的问题，第一次下载 Cursor 让 AI 把 Bug 给我修好了（AI 可不会水逆！）。

我这周真的投入了好多时间做游戏，虽然结果还是很稚嫩，画面也十分简陋，但我自己玩得很开心！这周我打算休整一下，思考下一个游戏做什么，同时也关注一下 itch.io 上玩家的反馈。

过段时间我应该会写一篇开发日志进行总结。

## 🧠 这周思考了什么？

### 你好像在看手机

Pippin Barr 做了一个叫做「[It is as if you were on your phone](https://pippinbarr.com/it-is-as-if-you-were-on-your-phone/info/)」的网页游戏，仅限手机游玩。作者说：「假装你在看手机从来没有像现在这样重要过！」或许是因为不想说法，又或者是为了融入人类，你可能需要假装你在看手机，但又不想让任何信息占据你的大脑，那就试试这款游戏吧！

游戏的 UI 分为两个部分，上半部分是动作提示词，会随机出现一些句子提示你现在应该做什么动作，比如：

1. 盯着屏幕看
2. 皱眉
3. 摇摇头
4. 把两根手指放在嘴唇上
5. 呼气
6. 抬高眉头然后身体前倾

下半部分是操作屏幕，上面会随机出现一些你能够点击或者双击的元素，有的时候你需要拖动小球，有的时候你只需要盯着一个旋转的线条看。

虽然最好的方法其实是避免无意义社交，积极参与有意义的社交，但人也有身不由己的时候，如果不想被打扰又不想刷其他无意义的信息，这款游戏说不定能让你在什么都不做的情况下逃避社交。

### 收藏知识的后续动作？

Xiaoxiao 写了一篇《[关于收藏知识的后续动作](https://liuyuntian.com/2025/03/13/subsequent-actions-regarding-collecting-knowledge.html)》，他发现自己过去十几年收集的好文章，如今看来大部分都过时了，一方面是因为技术和趋势变化太快，另一方面是因为这些知识一直没有被转化为自己的东西。

作者指出，对于收藏的知识，应该定期回顾，删除没有用的内容，及时总结输出，才能内化，让有用的知识变成自己的。

而我在想，能不能完全移除「收藏知识」这一步？因为收藏就意味着告诉自己「我可以一会再看」，到最后回顾这些收藏项目就会变得非常痛苦。如果能在当下就稍微多花几分钟，将读到的、听见到的、看到的，及时地用自己的话写成笔记，而不是「收藏」，兴许会好很多。

### 别让算法继续操纵你的思维

《[Kill your Feeds](https://usher.dev/posts/2025-03-08-kill-your-feeds/)》这篇文章中，作者指出各种软件和网站的推荐算法造成了一个**反馈循环**，这种循环会孤立用户（让有相同想法和好恶的用户聚在一起），加强信念，让相反的观点变得没那么重要。这其实就是中文互联网上被说烂了的信息茧房。作者认为这样的反馈循环会造成观点的极端化。

作者给出的建议是，不要使用平台的推荐算法，而是直接去找信息的源头。比如，你可以继续在 YouTube 上看视频，但比起没完没了地刷新首页，你可以将你喜欢的 YouTube 频道页面添加到书签，直接去他们的主页找内容看。

对我来说，RSS 阅读器所能提供的功能就是如此，通过直接订阅信息源，可以避开各种软件平台的算法推荐。最近，像 Follow 这样「集中化管理你订阅的内容」的应用已经不少，前几日也在英文互联网上发现过类似的产品（因为我不需要所以没有收藏，现在已经忘记了），兴许人们已经在有意识地夺回自己的注意力，只不过这样的人仍然占少数。

### 真相、谎言与进度条

一位设计师在他的[博客](https://cloudfour.com/thinks/truth-lies-and-progress-bars/)中讲到，在他事业初期的时候，他和工程师团队就「进度条」的问题进行过交流。工程师笑着告诉他，用户界面里的进度条是假的，进度条只是一开始移动一段固定的大小，之后的增量会逐次乘以一个小于 1 的值，这样进度条看起来在移动，但实际上永远也不会到达 100%（很多进度条会卡在 99% 不动，大概也是这个原因）。

设计师得知「计算某个操作的进度」是一件非常复杂且困难的事情，如果要动态地估计进度，还会消耗很多不必要的资源，但如果没有进度条，只是用转圈或者像理发店的三色灯一样呈现一种动态，用户可能会怀疑程序卡住了，然后不必要地取消操作。

这位设计师提出了很多解决方法，比如「计算操作步骤的数量」或者「标识当前正在进行的步骤」，这样不需要精细化的计算，也可以让用户更清晰地知道程序运行到哪一步了。

{{< img src="https://image.guhub.cn/uPic/2025/03/progress-prep-r6.jpg" title="把一个操作分成大致的几个步骤" >}}

{{< img src="https://image.guhub.cn/uPic/2025/03/progress-count-r2.jpg" title="如果不确定总步骤，就计算当前已经完成的步骤" >}}

还有在进度条下面增加一个小游戏的设计，这个我就不太看好了。要是操作比较快的话，用户可能玩游戏玩到一半就被打断了；如果操作本身比较长，简单的小游戏可能也不太容易消磨时间；要是觉得游戏不好玩，放在那里就和没有进度条没差别，而且还觉得碍眼。而且，难道用户就没有别的事情可以做了吗？

### 使用简单、不易过时的技术

Andre Garzia [用 Lua 语言重写了他的博客](https://andregarzia.com/2025/03/why-i-choose-lua-for-this-blog.html)，因为他发现 JavaScript 更新迭代的速度太快了，新框架层出不穷，旧框架的更新也非常频繁。他听说不少人需要花费很多精力才能让他们的博客系统保持最新，如果不这么做的话就会发生很多问题，这不是什么愉快的体验。

> JavaScript can be a light-speed moving mess.

作者指出，Lua 语言更新得非常慢，从 2006 年到 2020 年，只经历了 5.1 到 5.4 的版本更新，而且就算是这样，这两个版本的变化也非常少。要运行 Lua，只需要一个 C89 编译器，这比 JavaScript 简单太多了。

不过，Lua 也有一些缺点，他不是一个自带电池（batteries-inclueded）的编程语言。自带电池的意思是，提供了各种软件库，不需要太多的配置，就能直接拿来写程序，很方便。要实现一些 JavaScript 中很容易实现的功能，Lua 可能需要引入额外的第三方库（虽然写 JavaScript 程序要用到的第三方库也不少），而且还可能需要自己实现一些基础功能。

不过作者认为，选择 Lua 这种「简单无聊」的技术可以大大减少后期维护所需要经历。他也提醒读者，如果需要构建自己的博客系统，可以思考自己愿意维护这个程序多久，然后谨慎选择自己用到的技术。

### 接受缪斯的死亡

缪斯（muse）是智慧女神，往往被人们用来比喻作家、画家等创作者的「灵感来源」。[John P. Weiss](https://johnpweiss.com/) 说[他的缪斯死去了](https://johnpweiss.com/blog/199949/the-day-the-muse-died)，在搬家之后，他对家周围的风景提不起兴趣，于是丧失了作画的灵感。

不过，他发现自己逐渐对阅读和写作提起了兴趣。他想要提醒读者，如果往日的热情不再，不必担心，因为自己总会找到出路的。

> We must be willing to let go of the life we planned so as to have the life that is waiting for us.
> 我们必须愿意放弃已经计划好的生活，才能得到正在等待我们到来的生活。
>
> {{< right "——Joseph Campbell" >}}

此外，我注意到作者在字里行间中透露出来的某些细节。

> Iearned that online readers are impatient and tend to scan. Thus, it's important to hook online readers with tantalizing headlines, subheadings, and short paragraphs.

作者开始在互联网上写作时，发现网上的读者往往没什么耐心，更喜欢略读，所以他学着去写那些更有吸引力的标题、副标题，把段落写得更短，以此来吸引读者的注意力。

他的心态引发了我的一些思考，因为我一直觉得，要是有人没耐心读我的文章，那是他的损失，而他也不是我想要的读者。可这位作者的心态完全不同，他选择接受事实并寻求解决方案。

## 💡 这周学到了什么？

### View Transition

CSS 的 View Transition API 其实在 2023 年上半年就已经开始测试了，如今已经被大多数现代浏览器支持（除了火狐）。

![image-20250316100620618](https://image.guhub.cn/uPic/2025/03/image-20250316100620618.png)

你只需要在网站的 CSS 文件中加上：

```css
@view-transition {
  navigation: auto;
}
```

浏览器就能自动帮你实现简单的渐变切页效果，而且非常丝滑，已经足够应付大部分场景了。如果你想要更细致的切页动画，比如现在从本站首页点击文章链接，就能看到链接丝滑地过渡为页面文章标题的动效，你只需要额外为这些元素添加 `view-transition-name` 属性就可以了。

```html
<!-- 在文章列表中 -->
<a 
   href="/post/xxx" 
   style="view-transition-name: post-xxx"
>文章标题</a>

<!-- 在文章页面中 -->
<header style="view-transition-name: post-xxx">
	<h1>文章标题</h1>
</header>
```

关于 View Transition，如果想要了解更多，可以阅读 [Dave 的博客](https://daverupert.com/2023/05/getting-started-view-transitions/)，他也用这个 API 在博客上实现了许多酷炫的切换效果。

Jim Nielsen 也在他的博客提出了用很多小 HTML 页面作为组件，利用 View Transition 实现很多原本要用到 JavaScript 才能实现的效果，比如下拉菜单。如果你感兴趣，可以阅读[他的博客文章](https://blog.jim-nielsen.com/2025/lots-of-little-html-pages/)，我觉得非常有意思。前端在减少 JavsScript 代码数量上又前进了一大步。

---

## 🗣️ 碎碎念

- 同样上厕所看手机，为什么我一个 5 分钟的视频看不完就把屎拉干净走人了，三个室友个个都能蹲 20 分钟的样子，而且好几次卡在我「把这个做完我就去厕所」或者即将要出门的时间点，让我在被憋死和要急死的边缘气得直跺脚。*我真的非常担心我的膀胱健康*。
- 因为每天都在电脑面前过于投入地敲代码，连上课都恋恋不舍，所以好像没什么时间碎碎念？
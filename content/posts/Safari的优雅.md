---
title: Safari 的优雅
subtitle: Chrome vs. Safari
date: 2025-01-18 22:57:00
tags:
  - Apple
  - Geek
draft: false
toc: true
---

购入 [MacBook](/posts/揽物日志-vol-3/) 之后，我就抛弃了 Chrome 浏览器，也没有去使用很火的 Arc 浏览器。与很多人不同，Safari 并不是我下载 Chrome 的工具，在过去一个月的使用体验中，它作为我的主力浏览器，极大程度地改变了我的 Web 体验。

<!--more-->

## 用户界面

先横向对比一下 Safari 和 Chrome 的用户界面。

{{< gallery >}}
![](https://image.guhub.cn/uPic/2025/01/PixPin_2025-01-18_09-56-12.png)
![](https://image.guhub.cn/uPic/2025/01/PixPin_2025-01-18_09-56-42.png)
{{</ gallery >}}

Chrome 的标签栏和地址栏分界很明显，在任何时候，浏览器最上方都被分成了两栏。而 Safari 只有在存在多个标签页时才会显示标签页栏，而且 Apple 也没有把它分成割裂的两部分，颜色仍然是相似的，当前标签页和顶部的地址栏是连在一起的，一体化的感觉会更强。

![](https://image.guhub.cn/uPic/2025/01/PixPin_2025-01-18_09-59-45.png)

你也可以在设置里将 Safari 标签设置为紧凑布局，这样地址栏和标签页栏就可以融合在一起。更棒的是，此时的顶栏会匹配当前网站顶部的颜色，和网页的设计融合在一起。

{{< gallery >}}
![](https://image.guhub.cn/uPic/2025/01/image-20250118110522601.png)
![](https://image.guhub.cn/uPic/2025/01/image-20250118111307594.png)
!{}(https://image.guhub.cn/uPic/2025/01/image-20250118111439987.png)
{{</ gallery >}}

我认为这样的设计能让我在访问网站的时候更专注网站本身提供的内容，不会有太多的工具按钮分散我的注意力，让我一直想着「我能在这个网站上做点什么」。

尽管配色很和谐，但 Chrome 颜色分界明显的设计仍然产生了一种割裂感（至少对我来说），不仅标签页和地址栏有分界，地址栏和其他的工具按钮之间也有明显的分界。

此外，Chrome 右侧工具图标也很扎眼（长时间使用 Safari 再使用 Chrome，这种感觉就非常明显），我不明白为什么那个头像不能收进菜单里，它并不常用；以及那些浏览器拓展，真的有必要吗？

## 太多的浏览器拓展

Safari 和 Chrome 的使用体验最大的不同，可能是拓展程序的使用。Safari 的拓展程序远远不如 Chrome 丰富，我确实花了很多时间适应。没了 Dark Reader，在访问一些没有夜间模式的网站时我感到非常不适；没了 Notion Webclipper、Cubox 等插件，收藏网页内容也十分不便；Safari 使用油猴脚本也没那么方便（我才发现我到现在还没有尝试在 Safari 上使用过油猴脚本）。

使用 Safari 一个月之后，我发现自己并没有花多少时间折腾浏览器拓展。我不禁开始思考，自己之前给 Chrome 安装各种拓展以获得额外功能，真的有必要吗？

> 会不会是他对于额外功能永不满足的欲望，拖慢了他的浏览器？
>
> ——[告诉你为什么Chrome这么慢（不是内存的锅）#linus谈科技](https://www.bilibili.com/video/BV1Ut421K7bP/)

我不打算花时间讨论剪藏、文献管理、自定义用户脚本等功能的必要与否，因为每个人的需求都是不同的。至少对我来说，之前我认为很重要的 Webclipper 等功能，完全可以舍弃，或者用更简单的功能替代。

作为 Apple 原生应用，Safari 可以很方便地「共享」内容。如果我发现一篇不错的文章，打算稍后再读，比起添加到 Notion、Cubox 或者 Obsidian，我可以直接使用原生的「添加到阅读列表」，之后在 Safari 的首页看到我要读的内容。

阅读列表的设计对比 Cubox 这样专门的「稍后读」应用显得有些简陋，但其实完全够用，至少对我来说。

{{< gallery >}}
![](https://image.guhub.cn/uPic/2025/01/image-20250118105724547.png)
![](https://image.guhub.cn/uPic/2025/01/image-20250118225343611.png)
{{</ gallery >}}

## 垂直标签页

Safari 也带有垂直标签页和标签分组的设计，并且单独建立的标签页组可以在 Safari 关闭后保留，下次打开依然可以访问。尽管不如 Arc 浏览器的垂直标签页一般优雅便捷，但若是没有很强的需求，偶尔需要打开大量标签页查询资料时用一用完全够了。

![](https://image.guhub.cn/uPic/2025/01/image-20250118110144024.png)

Chrome 的话，一直没有垂直标签页的设计，我试过好几个利用 Chrome 侧边栏制作的垂直标签页拓展，都不太好用。

## 不适合开发者

刚用 Safari 浏览器的时候，我怎么也找不到清除缓存的按钮。就算用过一次，下一次也要在设置里找半天才能找到。这确实是很大的缺点。

Safari 也没有开发者工具，将它作为默认浏览器的弊端是，每次从 IDE 或者一些 CLI 工具打开本地运行的 Web 应用时，我都必须手动复制链接，再用 Chrome 打开。没有开发者工具真的会让调试变得很困难。

通过一些工具跑分发现，Safari 的性能也远不如 Chrome，在加载一些页面时可能会更慢，表现不如 Chrome。

我的 Mac 上装了两个浏览器，Safari 用于日常的浏览和资料查询，而 Chrome 用作开发。这样做可以最大程度上发挥两个浏览器各自的优势，也能让我将开发者和用户这两个身份分离开来。

## 对一般用户不友好

我本来想再写写多台苹果设备一起使用 Safari 有多么方便，比如 iCloud 标签页和无感同步的书签、访问记录、不装拓展就能直接用 Apple 的密码管理器等等。

转念一些，这些都是所谓的「Apple 全家桶」用户才能享受到的好处。Apple 给其他浏览器做的书签同步插件非常难用，不太能做到设备之间的无缝切换。Apple 的密码管理器我没在 Windows 上用过（以前用的是跨平台的 Enpass），这里就不讲了。

如果主力设备中有非苹果的，那 Safari 的使用体验大概会大打折扣。

---

## ☕️ TL;DR

- 试着用用装不了什么拓展的 Safari（或者关闭 Chrome 等浏览器的所有拓展程序），说不定你会发现自己根本不需要那些额外功能
- 如果你偏好简洁的设计，Safari 会比 Chrome 更适合你
- Safari 不适合开发者，如果你有前端开发的需求，你仍然需要 Chrome 或者其他带有开发者工具的浏览器
- 如果你的主力设备中不全是苹果设备，那么使用 Safari 的体验或许没那么好；但如果你使用苹果全家桶，就试试 Safari 吧！

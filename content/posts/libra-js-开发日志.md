---
title: Libra.js 开发日志
tags:
  - Web
  - Geek
date: 2025-09-23T12:58:00
draft: true
toc: true
---

博客从去年十二月开始就一直在用 [Fancybox](https://fancyapps.com/fancybox/) 作为博客的图片灯箱插件，之前还在用 Typecho 搭建博客的时候就在用这个库。所谓灯箱，其实就是点击图片之后将图片放大并居中显示，就像现实世界中的广告灯牌一样。不过我其实一直很想要把 Fancybox 换掉，因为它太臃肿了，包含了很多我不需要的功能（比如 1:1 放大、镜像翻转、前后切换等等），体积超过 100 KB；不仅如此，Fancybox 实际上是一个商业库，并不是免费开源的。出于这些原因，我决定自己从零开始做一个灯箱插件。<!--more-->

## 原则和技术选型

我希望项目能做到足够「轻」，一方面是打包后要足够小，另一方面是运行时不需要消耗太多资源。权衡之后，还有一个点就是必要的配置项足够少，最好能做到零配置起步。我决定用原生 JavaScript 编写，毕竟需求没有很复杂。由于需求本身足够简单，除了开发时会用到的一些依赖，打包后的库不应该包含或依赖于任何第三方库，这也能避免不必要的维护工作。此外，目前我能找到的灯箱插件大多都包含了多图排版功能或者额外的图片缩放操作，我希望这个库能够专注于需求本身，只要把图片放大就好了。

轻量、原生、专一，~~非常适合用作软件的宣传噱头的~~三个基调就这么定好了！

## 模块设计

需求只有一个，点击图片后播放动画并放大以供查看，再明确不过了，完全可以跳过需求分析的步骤，直接开始设计。

由于点击图片之后需要播放动画——原图片从原位置、原大小逐渐过渡到屏幕正中间、放大显示的状态——而图片 `<img>` 本身是一个行内元素或块元素，要让其「飞离」容器会有些困难，所以我的想法是构造一个**影子元素**（Shadow），通过绝对定位覆盖在图片上方，由于影子元素是绝对定位的，要放大和位移就很方便了。

除了放在正中央的图片，要组成灯箱，往往还需要**遮罩**（Overlay）。遮罩平时隐藏，打开图片灯箱之后显示在其他元素和影子图片之间，这样在视觉上，打开灯箱之后用户就只能看到放大的图片，可以起到突出强调的作用。

不需要其他花里胡哨的功能的话，**灯箱**（LightBox）就可以理解为「影子 + 遮罩」。影子的生命周期（创建和销毁）由自己管理，遮罩也用一个单独的模块管理，灯箱负责调用这两个模块，控制灯箱的开启和关闭，也负责把对应的图片信息传递给它们。

在之后的开发过程中，我发现影子模块的代码有些复杂，因为播放动画的逻辑也写在里面了，于是把动画相关的代码分离了出来，单独作为**动画**（Animation）模块。

于是四个模块的分工就很明确了：

- Shadow：管理影子元素的生命周期；获取图片的位置和尺寸信息。
- Overlay：创建、打开和关闭遮罩。
- LightBox：初始化监听器；调用 Shadow 和 Overlay 两个模块来打开灯箱。[^1]
- Animation：计算动画的起始和结束状态，并应用相关属性。

## 定位图片时踩的坑

一开始我在一两个小时内就实现了需求，在后续测试时却发现了明显的问题。我在测试页面里能够正常开关灯箱，图片也能被放到正确的位置，但放到我的博客上时却发现动画的起始位置完全不对，图片不是从原始位置放大的，而是从页面最下方瞬间飞上来的，而且飞上来之后的位置也有偏移。

### 初始状态的获取逻辑

发现原因在于定位时使用的属性不对，我是这样获取图片的初始位置的：

```js
// shadow 是根据图片创建的影子图片
// image 是原图
shadow.style.top = image.offsetTop;
shadow.style.left = image.offsetLeft;
```

一般情况下，`offsetTop` 和 `offsetLeft` 会返回元素相对于 `<body>` 的位置，`offsetTop` 是元素顶部距离 `<body>` 顶部的距离，`offsetLeft` 同理；可以理解为以整个页面的左上角为原点建立了一个坐标轴，只不过 Y 轴的正方向是向下的。

![](https://image.guhub.cn/uPic/2025/09/lovecoordsystem.png "图片来自 [LÖVE 引擎的官方文档](https://www.love2d.org/wiki/love.graphics)，由于这个图形系统和这里的例子很相似就拿过来做演示了；其中 x 可以理解为 left，而 y 是 top。")
{.dark:invert}

然而，这并不是 ` offset ` 系列属性的用途，实际上，它们返回的是相对于最近的**定位祖先元素**（positioned ancestor element）的距离。[^2]其中，祖先元素是 DOM 树上比自己更高层且自身属于其枝干的节点，也就是父元素的父元素…… 定位（positioned）指的是 CSS ` position ` 属性不为 ` static `，受 ` top ` ` left ` ` bottom ` ` right ` 等属性影响。如果没有这样的祖先元素，就会选择 `<body>` 作为参考对象。

问题就出在这里，如果图片位于一个定位元素（positioned element）中，那定位就不是相对于 `<body>` 的；然而，影子元素的定位方式是 `position: absolute`，它在被创建时是插入到 `<body>` 的末尾的，也就是说，影子元素是相对于 `<body>` 进行定位的。由于计算得出的位置和实际应用的位置，两者的参考系不同，最终定位的结果自然是错位的。知道原因之后，解决方案就很简单了，只要获取原图相对于 `<body>` 的位置信息就好。

在 MDN 上查阅之后，我找到了 `getBoundingClientRect()` 这个方法，用于获取一个矩形对象，这个矩形的位置信息是相对于**视口**的，也就是浏览器窗口中显示页面的区域。矩形的 `top` 属性就是这个元素相对视口顶部的距离，如果加上页面滚动的距离，就能得到这个元素相对于 `<body>` 顶部的距离；`left` 属性也是同理。

```js
const rect = image.getBoundingClientRect();
this.originalPosition = {
  top: rect.top + window.scrollY,
  left: rect.left + window.scrollX,
  width: rect.width,	// 宽高的获取没什么要绕弯子的
  height: rect.height	// 这里就不追述了
};
```

着用获取的值就可以直接用来定位影子元素了，只要宽高也是一样的，就能遮盖住原图片。

<!--这段有点抽象了，之后画个图吧-->

### 结束状态的获取逻辑

灯箱打开动画的结束状态就是灯箱放大后的位置和大小，这部分的需求是这样的：

1. 放大后图片的长宽不应该超过视口大小，也就是不应该出现放大后图片显示不全的情况；
2. 放大后图片的位置应该居中与视口正中央。
3. 出于美观考虑，应该给放大后的图片设置一个边距，在四周留出一定的空白。

……

## 绘制动画时踩的坑

`requestAnimationFrame()` 和 CSS transition

## JavaScript 设置样式时踩的坑

没写单位啊！

[^1]: 其实 LightBox 模块还能拆开来，初始化监听器和开关灯箱明显是两个职责，但我担心分出多个模块会让结构变得复杂。

[^2]: 参见："The offsetTop read-only property of the HTMLElement interface returns the distance from the outer border of the current element (including its margin) to the top padding edge of the offsetParent, the closest positioned ancestor element." —— [HTMLElement: offsetTop property - Web APIs \| MDN](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetTop#:~:text=The%20offsetTop%20read%2Donly%20property%20of%20the%20HTMLElement%20interface%20returns%20the%20distance%20from%20the%20outer%20border%20of%20the%20current%20element%20(including%20its%20margin)%20to%20the%20top%20padding%20edge%20of%20the%20offsetParent%2C%20the%20closest%20positioned%20ancestor%20element.)

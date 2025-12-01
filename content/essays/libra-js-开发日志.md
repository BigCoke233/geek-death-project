---
title: Libra.js 开发日志
tags:
  - Web
date: 2025-09-23 22:22:00
draft: false
toc: true
aliases:
  - /posts/libra-js-开发日志
categories: Coding
---

博客从去年十二月开始就一直在用 [Fancybox](https://fancyapps.com/fancybox/) 作为博客的图片灯箱插件，之前还在用 Typecho 搭建博客的时候就在用这个库。所谓灯箱，其实就是点击图片之后将图片放大并居中显示，就像现实世界中的广告灯牌一样。不过我其实一直很想要把 Fancybox 换掉，因为它太臃肿了，包含了很多我不需要的功能（比如 1:1 放大、镜像翻转、前后切换等等），体积接近 100 KB；不仅如此，Fancybox 实际上是一个商业库，并不是免费开源的。出于这些原因，我决定自己从零开始做一个灯箱插件。<!--more-->

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

一开始我在一个小时内就实现了需求，在后续测试中却发现了明显的问题。我在测试页面里能够正常开关灯箱，图片也能被放到正确的位置，但在我的博客上应用插件时却发现动画的起始位置完全不对，图片不是从原始位置放大的，而是从页面最下方瞬间飞上来的，而且飞上来之后的位置也有偏移。

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

这样获取的值就可以直接用来定位影子元素了，只要宽高也是一样的，就能遮盖住原图片。

![](https://image.guhub.cn/picgo2025/IMG_0395.jpg "为了方便你理解，我在画了一个示意图。")
{.dark:invert}

### 结束状态的计算方式

灯箱打开动画的结束状态就是灯箱放大后的位置和大小，这部分的需求是这样的：

1. 图片应该尽量展开到原始尺寸，但不应该放大超过原始尺寸导致模糊；
2. 放大后图片的长宽不应该超过视口大小，也就是不应该出现放大后图片显示不全的情况；
3. 放大后图片的位置应该居中于视口正中央；
4. 出于美观考虑，应该给放大后的图片设置一个边距，在四周留出一定的空白。

图片的大小应该是窗口大小和图片自然情况下的大小共同决定的，也就是以下四个变量。

```js
const nw = element.naturalWidth;	// 图片自然宽度
const nh = element.naturalHeight;	// 图片自然高度
const ww = window.innerWidth;		// 窗口内部宽度（视口宽度）
const wh = window.innerHeight;		// 窗口内部高度（视口高度）
```

这里的逻辑其实很简单，如果优先计算宽度的话，选取图片宽度和视口宽度当中**最小**的一个就能保证图片不超过视口。不过要注意的是，并不能直接用相同的方法去获取结束状态的宽度，因为这可能会导致图片的比例变化，所以要先算出原始图片的长宽比例，然后通过比例计算出高度。

```js
const ratio = nw / nh;
finals.width =  (nw > ww) ? ww : nw; // or `Math.min(nw, ww)`
finals.height = (nw > ww) ? ww / ratio : nh;
```

不过，通过比例计算出来的高度完全有可能超出视口高度，所以要再进行一次判断。

```js
if (finals.height > wh) {
  finals.height = wh;
  finals.width = wh * ratio;
}
```

这样就基本上没问题了。尺寸确定之后，才能计算最终的位置。

已知矩形的长宽，也知道视口的长宽，要算出让矩形居中的坐标位置就很容易了，只要把视口的长度减去矩形的长度再除以二就好，宽度也同理。这个操作应该还挺经典的，我记得我在学 Java Swing[^3] 的时候为了把窗口居中就经常做这样的计算。由于影子图片的定位是基于 `<body>` 的，所以还要像刚才那样加上滚动的距离。

```js
finals.top = (wh - finals.height) / 2 + window.scrollY;
finals.left = (ww - finals.width) / 2 + window.scrollX;
```

这样就得到了放大尺寸合适、位置居中的图片状态。

还要添加边距，我本来以为要修改宽高的计算公式，后来发现，其实只要在**逻辑上**缩小视口大小、再偏移坐标位置就好了，简单来说就是这样：

```js
// margin 是边距
const ww = window.innerWidth - 2*margin;
const wh = window.innerHeight - 2*margin;

finals.top = (wh - finals.height) / 2 + window.scrollY + margin;
finals.left = (ww - finals.width) / 2 + window.scrollX + margin;

// 其他全都保持不变
```

## 绘制动画时踩的坑

前面已经获取了影子元素的初始状态和结束状态，也就确定了动画的开始帧和结束帧，现在要做的就是把影子元素平滑地移动并放大到结束状态，补全中间的关键帧。

我能想到的最直接的办法就是通过 `setInterval()` 每隔一段时间更新一次影子元素的 `top` `left` `width` 和 `height` 属性，只要设置适当的步数（帧数）和持续时间，应该就能达到比较平滑的效果。然而，无论我怎么调整参数，动画效果都非常不如意。最主要的问题是太慢了，没有人愿意盯着一个图片动那么久。

我一开始仍然觉得是参数设置的问题，于是把代码发给 ChatGPT 修改，好在 ChatGPT 没有盲目听从我的指令，而是建议我使用 `requestAnimationFrame()` 这个方法，这个方法会请求浏览器在下一次**重绘**（repaint）时执行传入的用户回调函数。这有点像 LÖVE 引擎里的 `love.draw()` 方法，但又不完全一样，因为 `requestAnimationFrame()` 只会执行一次，所以要在传入的回调函数里再次执行 `requestAnimationFrame()`，直到动画完成之后再停下来，也就是要用到递归。

```js
animate(startingState, finalState) {
    // 初始化步数，即从开始到结束一共绘制几帧
    const steps = 20;
    let step = 0;

    // 每一帧会执行的回调函数
    const frame = () => {
      // 通过当前步数和总步数计算进度
      const progress = step / steps;

      // 根据进度计算出当前的位置和尺寸
      const currentTop = startingState.top + (finalState.top - startingState.top) * progress;
      const currentLeft = startingState.left + (finalState.left - startingState.left) * progress;
      const currentWidth = startingState.width + (finalState.width - startingState.width) * progress;
      const currentHeight = startingState.height + (finalState.height - startingState.height) * progress;

      // 更新影子元素的样式
      this.element.style.top = `${currentTop}px`;
      this.element.style.left = `${currentLeft}px`;
      this.element.style.width = `${currentWidth}px`;
      this.element.style.height = `${currentHeight}px`;

      step++;	// 自增步数

      // 步数未到，则继续绘制下一帧
      if (step < steps) {
        requestAnimationFrame(frame);
      }
    };

    // 开始第一帧
    requestAnimationFrame(frame);
  }
```

[MDN](https://developer.mozilla.org/en-US/docs/Web/API/Window/requestAnimationFrame) 是这样描述这个方法的：

> The frequency of calls to the callback function will generally match the display refresh rate. The most common refresh rate is 60hz, (60 cycles/frames per second), though 75hz, 120hz, and 144hz are also widely used.
>  
> 调用回调函数的频率一般会和显示器的刷新率匹配。最常见的刷新率是 60 赫兹（每秒 60 个循环/帧），不过 75 赫兹、120 赫兹和 144 赫兹也被广泛使用。

我不确定描述是否属实，因为测试时动画偶尔会掉帧，往往是在第一次执行动画时，这有可能是浏览器或者操作系统为了平衡性能和管理系统资源导致的偶发性问题吧，不过我只是在画 2D 图像啊…… 如果图片要移动的距离比较长，也容易出现卡顿或者闪现，这是因为动画的帧数是写死的，太长的话，每帧之间的差异也就会大一些。总之，使用 `requestAnimationFrame` 的方法不太美妙。

最后我选择参考 Fancybox 的实现方式，这才发现 Fancybox 的 GitHub 仓库里居然只有一个打包好、压缩过的 JavaScript 文件——原来你不是开源的啊！无妨，用浏览器检查元素也能看到 Fancybox 是怎么实现缩放动画的。观察之后，我发现 Fancybox 并没有逐帧更新元素，而是只更新了一次元素的 `transform` 属性，更改的值是我从没见过的 `matrix()`。

查阅之后才发现 `matrix()` 是用于表示 2D 变换的 CSS 函数，实际上它表示了**齐次坐标**（Homogeneous coordinates）的变化，按照 MDN 的说法，是**齐次二维变换矩阵**（homogeneous 2D transformation matrix）。说人话的话，可以把它理解为以下几个 2D 变换的缩略写法：缩放（scale）、歪斜（skew）和位移（translate）。

简单来说：

```css
transform: matrix(a, b, c, d, e, f);
/* 等价于 */
transform: scaleX(a) skewY(b) skewX(c) scaleY(d) translateX(e) translateY(f);
```

*好险，差点就要被线性代数吓死了。* 😰

所以 Fancybox 的做法是，计算要将图片放大并居中到对应状态，所需要对图片做出的 2D 变换，也就是不直接改变元素的坐标和宽高，而是对图片进行位移和缩放。这样做的好处是可以不用 JavaScript 处理动画，利用 CSS 自带的过渡（`transition`）实现简单且丝滑的动画。

这听起来不难，不过当时的我已经精疲力尽，便把这个发现告诉了 ChatGPT，让它帮我写。事实证明，大语言模型没有任何平面想象能力，因为他最终给出的结果简直惨不忍睹，图片都不知道飞到哪里去了。

_*闭眼捏鼻 *吸气 *呼气……_

得，还是自己来吧。

我调试了半天都还是没办法把图片变换到正确的位置上，后来发现一个很关键的问题。在此之前，图片的位置一直是根据元素左上角的坐标点确定的，如果你没理解，可以回看这一张图。

![](https://image.guhub.cn/uPic/2025/09/lovecoordsystem.png "注意看小猪图片的坐标，是左上角的点的坐标。")
{.dark:invert}

然而，`transform` 对元素进行缩放时，是从中心点进行缩放的。根据中心点放大或缩小之后，图片左上角的位置会向左、向上偏移；实际上，除了中心点以外的其他点的位置都会移动。所以，要计算图片变换到最终位置需要在 X 和 Y 轴上位移多长的距离，要先找到图片始末状态的中心点，这是唯一一个不会在缩放之后偏移从而造成计算误差的点。

先计算出开始状态和结束状态下，中心点的 X、Y 坐标。

```js
const startCenterX = starts.left + starts.width / 2;
const startCenterY = starts.top + starts.height / 2;
const finalCenterX = finals.left + finals.width / 2;
const finalCenterY = finals.top + finals.height / 2;
```

然后计算两个中心点的差值就能得到位移的距离。

```js
const translateX = finalCenterX - startCenterX;
const translateY = finalCenterY - startCenterY;
```

至于缩放的倍数，就非常简单了。

```js
const scaleX = finals.width / starts.width;
const scaleY = finals.height / starts.height;
```

计算完变换矩阵之后就可以把这些属性应用到元素上了，然后借助 CSS 自带的 `transition` 就能看到图片放大和缩小的动画了。

```js
// 放大、居中的变换
element.style.transform = `matrix(${scaleX}, 0, 0, ${scaleY}, ${translateX}, ${translateY})`;

// 缩小、复位的变换
element.style.transform = `matrix(1, 0, 0, 1, 0, 0)`;
```

## JavaScript 设置样式时踩的坑

以上一切问题都排查完毕之后，图片灯箱居然还是没有办法在博客上正常工作！明明在测试页面中就好好的！我百思不得其解，通过检查元素才发现，博客上的影子元素根本没有被应用 `top` `left` `width` 和 `height` 属性，也就是说一开始影子元素就没有遮在原图片的上方，初始位置一直在页面末尾，灯箱一直是从最底下飞上来，又从正中间飞下去……

可是为什么测试页面里就是正常的啊！

为了自己的头发，我把自己观察到的问题和可能的原因连同代码一起发给了 ChatGPT，结果对方表示：

>  你没写单位。

*啊…… 果然难改的 Bug 都是这种原因吗……*

因为 JavaScript 通过操作元素 `style` 属性修改样式的方式，实际上是修改了元素的行内属性，而 HTML 属性只有字符串类型，没有整数和浮点数，所以我直接把计算得到的数字赋值给 `style.top` 等属性的方式不总是能奏效。JavaScript 作为一门弱类型语言，不总是能够正确地把数字转换为字符串，所以要加上单位，因为单位是字符串，一个数字加上字符串的操作就能让 JavaScript 明白我希望得到字符串的结果。

```js
positionShadow() {
	this.element.style.top = this.originalPosition.top + 'px';
	this.element.style.left = this.originalPosition.left + 'px';
	this.element.style.width = this.originalPosition.width + 'px';
	this.element.style.height = this.originalPosition.height + 'px';
}
```

JavaScript 真的是一坨混沌的产物！

![](https://image.guhub.cn/uPic/2025/09/programming-people-javascript.png "图片来自：[leftoversalad.com/c/015\_programmingpeople/](https://leftoversalad.com/c/015_programmingpeople/)")

## 喜闻乐见的命名环节

用 Parcel 打包，最终得到了一个不到 4 KB 的 JavaScript 文件，包的大小我很满意。

![](https://image.guhub.cn/uPic/2025/09/qP5PmF.jpg)

由于是灯箱（LightBox）插件，又注重轻量，所以名字也简短一点比较好，我取了 Light 和 Box 两个词的前两个字母，取名叫 Libo。后来又觉得这个名字看着不太顺眼，读着也没有很上口，就改了。因为 Libo 读起来很像 Libra（天秤座），干脆就改成了这个好听又好记的名字。

现在有点想凑齐十二星座了……

---

## 最后

你可以在 GitHub 上查看 Libra.js：[BigCoke233/libra](https://github.com/BigCoke233/libra)

之前用 LÖVE 引擎开发 2D 游戏的经验居然也在 Web 前端开发这一领域给了我帮助，果然天下没有白干的活，以往的经历会绕着弯子来帮助现在的自己。

嗯？你问我什么时候开发下一个游戏？呃…… _*擦汗_ 要不我们还是下次再聊？

_*转身逃跑_

[^1]: 其实 LightBox 模块还能拆开来，初始化监听器和开关灯箱明显是两个职责，但我担心分出多个模块会让结构变得复杂。

[^2]: 参见："The offsetTop read-only property of the HTMLElement interface returns the distance from the outer border of the current element (including its margin) to the top padding edge of the offsetParent, the closest positioned ancestor element." —— [HTMLElement: offsetTop property - Web APIs \| MDN](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetTop#:~:text=The%20offsetTop%20read%2Donly%20property%20of%20the%20HTMLElement%20interface%20returns%20the%20distance%20from%20the%20outer%20border%20of%20the%20current%20element%20(including%20its%20margin)%20to%20the%20top%20padding%20edge%20of%20the%20offsetParent%2C%20the%20closest%20positioned%20ancestor%20element.)

[^3]: 一个很老旧的 Java 图形库。

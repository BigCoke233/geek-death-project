---
title: 零 JavaScript 的博客统计图表
tags:
  - Web
  - Geek
date: 2025-06-28 16:09:00
draft: false
toc: true
slug: zero-js-chart
aliases:
  - /posts/zero-js-chart
categories: Coding
---


受到 [Jim Nielsen](https://blog.jim-nielsen.com/about/) 的启发，我在一个月前决定给博客添加[统计页面](/stats/)，展示字数、每年每月发布的文章数量、外部链接排行等等，在效仿的过程中我还产生了一些新点子，比如统计文章发布的时间段，看看是白天发的文章多还是晚上更多。

统计免不了用很多图表，但我使用的是本身就追求「轻」的 Hugo，一个静态博客生成器，如果要引入臃肿的 JavaScript 图表库，就显得不太优雅了。我并不是排斥任何 JavaScript，那些能带来舒适交互体验的脚本我当然会毫不犹豫地使用，但我总觉得用 JavaScript 生成简单的图表是不必要的。<!--more-->

## 如何用 HTML+CSS 制作图表

### 基于宽度比例实现的条状图

既然不用 JavaScript 绘制图形，那么就只能用 HTML 和 CSS 限定元素的宽度或高度，通过元素大小的差异来表现「比例」，以实现统计图表想要达到的效果。下面是博客的「文章发布时间分布图」。

![image.png](https://image.guhub.cn/picgo2025/20250627225917.png)

它的 HTML 结构大致如下（忽略了部分不必要的元素）：

```html
<div class="time-chart-container">
	<div class="time-chart-wrapper">
		<div class="time-icon">☀️</div>
		<div class="time-chart-bar dark:invert">
			<!-- for each time segement -->
			<div class="time-segment"
				 style="width: {{$percentage}}">
				{{$label}}
			</div>
			<!-- end for -->
		</div>
		<div class="time-icon">🌙</div>
	</div>

	<div class="time-counts">
		<!-- for each time segement -->
		<div class="count-item"
		 style="width: {{$percentage}}">
			<span class="count-number">{{$count}}</span>
		</div>
		<!-- end for -->
	</div>
</div>
```

简单来说，通过某种方式计算得出了每一个分组的百分比数据，并将这个百分比作为元素的宽度；因为父元素的大小是 `100%`，且使用 `flex` 布局，浏览器就会将这些代表各个分组的元素按照这样的比例并排显示在容器内。

一个最简单的例子：

```html
<div class="wrapper" style="width:100%;display:flex">
	<div class="segment" style="width:25%"></div>
	<div class="segment" style="width:12.5%"></div>
	<div class="segment" style="width:12.5%"></div>
	<div class="segment" style="width:50%"></div>
</div>
```

只要给各个 `segment` 添加高度和背景色，以及其他的样式，就能实现非常不错的统计图表效果。由于使用 `flex` 布局，还能轻易地对排列方式进行调整，以及添加分段的间距（`gap` 属性）。

也可以给每个 `segment` 都裹上各自的 `wrapper`，然后纵向排列显示，就能制作出一个效果不错的纵向条状图，非常适合用来制作排行表。

![image.png](https://image.guhub.cn/picgo2025/20250628115709.png "本站的外部链接排行表")

### 横向条状图

![image.png](https://image.guhub.cn/picgo2025/20250628150944.png)

如果要制作传统的横向条状图，实现方式也是类似的，只是把宽度换成了高度。不过由于排版的方向不同，还需要显示其他的文字元素，实现起来要复杂一些。

```html
<div class="chart-container">
	<!-- for each month -->
    <div class="chart-item">
        <div class="bar">
          <div class="bar-number">{{$val}}</div>
          <div class="bar-value"
               style="height: {{$percentage}}%"></div>
        </div>
        <div class="bar-label">{{$month}}</div>
    </div>
	<!-- end for -->
</div>
```

复杂的部分在于 CSS，但说实话也没有复杂到哪里去。由于基于宽度排列时，宽度往往是有自然的最大值的，百分比会根据这个最大值被计算成实际的宽度值，但高度就没有，需要给容器限定一个高度。

```css
.bar {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-end;
  width: 1rem;
  gap: 0.25rem;
  height: 200px; /* fixed height */
  position: relative;
}
```

为了保证条状图的所有「条」都在同一条水平线上（往往是靠底部排列的），最好也使用 `flex` 布局，保证元素排列在下方（`align-items: flex-end`）。同时，由于横向排列时，条状图的宽度可能会很大，超出了容器能显示的范围，最好设置水平的溢出为 `auto`，宽度为 `100%`，保证宽度溢出时会显示一个滚动条；如果希望展示最右侧的数据（最左侧的数据需要滚动到左边才能显示），那么 `justify-content` 也要设置为 `flex-end`。

```css
.chart-container {
  display: flex;
  align-items: flex-end;
  justify-content: flex-end;
  /* ... */
  width: 100%;
  overflow-x: auto;
  box-sizing: border-box;
}
```

除了条状图的「条」，如果还需要在条上显示数据，在条下显示图例，那结构就更复杂了。简单来说，条状图的所有元素都放在 `chart-container` 中；`chart-container` 中放置的是数个 `chart-item`，`chart-item` 包括「条」「图例」两部分；其中「条」表示为 `bar` 类，包含数字（`bar-number`）和表示数值大小的图形（`bar-value`）。你可以回顾上文的 HTML 结构来理解。

### 饼图和折线图

是的，不需要 JavaScript 也可以实现饼图（Pie Chart）！不过我并不喜欢饼图，也就没有在博客里实现。如果你感兴趣，可以参考 StackOverflow 上的[这个问题](https://stackoverflow.com/questions/10028182/how-to-make-a-pie-chart-in-css)。

饼图的 HTML 结构异常简单，你只需要：

```html
<div class="pie"></div>
```

剩下的全部交给 CSS，通过锥形渐变实现。

```css
.pie {
  width: 400px;
  height: 400px;
  background-image: conic-gradient(orange 64%, blue 64%, blue 81%, black 81%);
  border-radius: 50%
}
```

只不过这里的比例并不是颜色的占比，而是颜色的开始位置，所以要转换数据的话会比较麻烦。StackOverflow 里还有更多的实现方式，这里就不展开了。

[这篇文章](https://webdesign.tutsplus.com/how-to-make-line-charts-in-pure-css--cms-108657t)讲到了如何使用 CSS 制作折线图，似乎用到了 `clip-path` 属性，上手难度就更大了。条状图其实已经足够展现出「趋势」了，不想折腾的话也没必要做折线图。演示占比也可以用一开始提到的宽度方案，不需要用饼图。总而言之，宽度和高度的方案已经足够展示大部分的统计数据了。

## 通过 SSG 获取统计数据

如你所见，零 JavaScript 的统计图表依赖于准确计算的百分比值，既然没有 JavaScript，要如何获取真实的统计数据呢？如果你使用的是 Hugo 这样的静态网站生成器（SSG，Static Site Generator），可以在网站生成阶段就把这些数据写入静态网页；也就是说，让 SSG 在生成网站的同时把图表也绘制出来（所谓「绘制」，其实就是计算出百分比数据而已）。[^1]

由于我使用的是 Hugo，所以接下来会用 Go 语言和 Hugo 的模板语法来做演示。Hugo 用户可以直接拿来用，但如果你使用的是 Hexo 或者其他的 SSG，就需要自己写代码了。[^2]

我依然使用「文章发布时间分布」为例子。首先定义一个名为 `fuzzyTimes`（模糊时间）的字典，通过小时给时间分段，比如在 5 点以前就算凌晨，7 点以前算清晨等等。

```go
{{ $fuzzyTimes := slice
    (dict "limit" 5  "label" "凌晨" "class" "dawn")
    (dict "limit" 7  "label" "清晨" "class" "early-morning")
    (dict "limit" 11 "label" "早上" "class" "morning")
    (dict "limit" 13 "label" "中午" "class" "noon")
    (dict "limit" 17 "label" "下午" "class" "afternoon")
    (dict "limit" 19 "label" "傍晚" "class" "evening")
    (dict "limit" 22 "label" "晚上" "class" "night")
    (dict "limit" 24 "label" "深夜" "class" "late-night")
}}
```

然后再初始化一个词典（键值对）`timeCounts`，用来保存一会收集到的统计数据。

```go
{{ $timeCounts := dict
    "凌晨" 0
    "清晨" 0
    "早上" 0
    "中午" 0
    "下午" 0
    "傍晚" 0
    "晚上" 0
    "深夜" 0
}}
```

接下来就是统计的流程。遍历网站所有的页面（`.Site.RegularPages`），判断是否有日期属性，如果有的话就可以继续。判断得到每个页面发布的时间段之后，对应的 `timeCounts` 就加 1，所有的页面遍历完成之后就得到了所需的数据。在这个过程中一直累加 `total` 变量，最终得到总数，用于计算百分比。

```go
{{ $total := 0 }}
{{ range .Site.RegularPages }}
    {{ if .Date }}
        {{ $hour := .Date.Hour }}
        {{ $fuzzyTime := "深夜" }}
        {{ range $fuzzyTimes }}
            {{ if lt $hour .limit }}
                {{ $fuzzyTime = .label }}
                {{ break }}
            {{ end }}
        {{ end }}
        {{ $total = add $total 1 }}
        {{ $currentCount := index $timeCounts $fuzzyTime }}
        {{ $timeCounts = merge $timeCounts (dict $fuzzyTime (add $currentCount 1)) }}
    {{ end }}
{{ end }}
```

把刚才的数据用 HTML 展示出来即可，`percentage = count / total`，是小学生都会做的数学题。如果你也把文字图例显示在 `segment` 里面，就需要注意太短的 `segment` 可能会导致文字溢出，可以通过判断 `percentage` 的大小决定要不要显示文字。
```html
{{ range $fuzzyTimes }}
	{{ $count := index $timeCounts .label }}
	{{ if gt $total 0 }}
		{{ $percentage := div (mul $count 100.0) $total }}
		{{ if gt $count 0 }}
			<div class="time-segment {{ .class }}"
				 style="width: {{ printf "%.1f" $percentage }}%;">
				{{ if gt $percentage 8 }}{{ .label }}{{ end }}
			</div>
		{{ end }}
	{{ end }}
{{ end }}
```

要遍历所有的页面看起来是开销很大的事情，但 Hugo 的构建速度很快，比 Node.js 快多了。从算法的角度来看，时间复杂度是线性的，不会很慢。我把这段代码做成了一个单独的 partial，包含完整的 Go 语言代码和 HTML 以及 CSS，如果你也使用 Hugo，可以直接复制[这里的代码](https://github.com/BigCoke233/geek-death-project/blob/master/layouts/partials/components/publish-time-chart.html)来用。

如果要收集其他的统计数据，算法自然会有所不同，这里只是展示这个方法的基本思想：通过 SSG 在构建网站的同时计算数据并生成对应的 HTML 结构，完全不需要 JavaScript 就能制作直观的统计图表。

## SSG 无法统计的数据

Hugo 等 SSG 能够对发布时间做出统计，是因为每个页面的元数据中就有「日期」这个数据，可以直接拿来用。然而，有的数据没有现成可用的，比如前文提到的「外部链接排行」，没有正常的 SSG 和 CMS 会统计这个东西。

此时，静态网站的另一个好处就体现出来了。由于所有的文章都是以纯文本的 Markdown 文件保存在 `content` 目录中的，所以可以用外部脚本扫描这个目录，通过正则表达式匹配需要的模式，比如以 `http(s)://` 开头的外部链接，然后把域名作为键，域名出现的次数作为值，放在一个键值对里。[^3]

显然，使用外部的脚本统计这些数据，就需要把数据保存在一个文件里才能使用。如果要访问文件里的数据，那难道不需要编写 JavaScript 代码吗？

还真不用，Hugo 可以在构建网站时访问外部的数据源[^4]，只要提前用脚本把数据导出为 JSON 或者其他 Hugo 支持的格式，Hugo 就可以在构建的时候使用这些数据。举个例子：

1. 在构建网站之前，先执行 `extractData.py` 脚本，这个脚本在 Hugo 的 `data` 目录下生成一个名为 `datafile.json` 的文件。
2. 在 Hugo 模板里这样访问数据 `{{ .Site.Data.datafile }}`；如果是数组类型的数据，则可以这样遍历： `{{ range .Site.Data.datafile }}`。
3. 执行 `hugo` 命令构建网站即可。

听起来很麻烦，但完全可以自动化。如果你使用 Cloudflare Pages、GitHub Pages 或者 Vercel 这类支持「持续集成 / 持续部署」（CI/CD）的平台，一般可以自定义构建命令。

![](https://image.guhub.cn/picgo2025/20250628145311.png "Cloudflare Pages 的构建配置")

你只需要在 `hugo` 命令之前先允许生成数据文件的脚本，保证 Hugo 在构建网站时能够访问需要的数据文件就可以了。

```shell
# 例如：
python3 generateData.py && hugo
```

就算是用 Python 进行 I/O 操作读文件，构建时间也不会被拖得很长。我使用一个 Python 脚本生成外部链接的排行数据和内部连接的[图谱](/graph/)数据，再使用 Hugo 构建网站，最后还用 Pagefind 扫描 `public` 目录生成搜索功能所需的索引数据。这些步骤都能够在 35～40 秒内完成，跟文件多了一些就需要构建两三分钟的 Hexo 比起来实在是快太多了。

## 最后

掌握了以上知识，就可以不用 JavaScript，给静态网站加上各种图表了！在优雅、强大、直观的同时还不会拉低网站的构建速度和渲染速度，可谓是一举两得。我的博客是开源的，如果你也在使用 Hugo，以下是相关的文件，你可以直接从 GitHub 上复制代码来用。[^5]

> [!note]
> AI Disclosure：以下代码在编写时借助了 AI 工具。

- 文章字数统计：[/layouts/partials/components/word-count-chart.html](https://github.com/BigCoke233/geek-death-project/blob/master/layouts/partials/components/word-count-chart.html)
- 文章发布时间分布：[/layouts/partials/components/publish-time-chart.html](https://github.com/BigCoke233/geek-death-project/blob/master/layouts/partials/components/publish-time-chart.html)
- 外部链接排行：
	- 生成链接数据的 Python 脚本：[/scripts\/analyze_links.py](https://github.com/BigCoke233/geek-death-project/blob/master/scripts/analyze_links.py)（这个脚本同时也会生成供[内部链接图谱](/graph/)使用的 JSON 数据）
	- 组件：[/layouts/partials/components/external-link-rank.html](https://github.com/BigCoke233/geek-death-project/blob/master/layouts/partials/components/external-link-rank.html)

如果你感兴趣，[这个组件](https://github.com/BigCoke233/geek-death-project/blob/master/layouts/partials/components/toc.html)是博客的目录组件，可以用来替代 `tocbot.js`——借助 Hugo 自带的 `{{ .TableOfContents }}`，免去在前端用 JavaScript 生成目录的麻烦，同时还加上了一定的 CSS 和 JavaScript 保证易用性。

[^1]: 传统的、动态生成内容的 CMS（例如 Typecho、Wordpress……）应该也能做到。不过都用动态网站了，也没必要在这个方面追求极简了。

[^2]: 或者把这篇文章喂给 AI，就当我帮你写好提示词了。

[^3]: 等等，这为什么会是静态网站的好处，动态网站遍历数据库难道不比读文件更快吗……

[^4]: 参见官方文档：[Data sources](https://gohugo.io/content-management/data-sources/)

[^5]: 最近一直在添加新功能，没什么心思维护和重构，所以可能会看到一些冗余或者耦合度高的代码，请见谅。

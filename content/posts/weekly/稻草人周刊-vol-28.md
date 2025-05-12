---
title: 稻草人周刊 Vol.28
subtitle: 你更有可能死在自己的生日
tags:
  - 游戏开发
  - 叙事
  - 道德
  - AI
  - 设计
  - 创作
  - 效率
  - 写作
  - 死亡
date: 2025-05-12T11:17:00
categories: 稻草人周刊
banner: https://image.guhub.cn//uPic/202505121117-weekly28-banner.jpeg
draft: false
toc: true
---

<iframe allow="autoplay *; encrypted-media *;" frameborder="0" height="150" style="width:100%;max-width:660px;overflow:hidden;background:transparent;" sandbox="allow-forms allow-popups allow-same-origin allow-scripts allow-storage-access-by-user-activation allow-top-navigation-by-user-activation" src="https://embed.music.apple.com/cn/album/naked-in-manhattan/1698723205?i=1698723328"></iframe>

这首歌歌词里的 Constant like cicadas in the summertime（像夏日蝉鸣一样连续不断）被我听成了 Nonsense like gaydars in the summertime（像夏天的 gay 达一样纯属胡扯）——因为这首歌是 Chappell 的所以我一度觉得非常合理。<!--more-->

## 🙋 在做什么

到了每周的审问环节才发现自己在这七天里好像什么都没干，就是很平常地生活，甚至带着一些伴随初夏一起出现的烦躁不安（不过实际上还没到夏天吧）。

值得说道的可能是这几天做了些断舍离，该扔的扔了，可以送人的送了（鬼知道我怎么囤这么多咖啡豆罐子、票据，甚至还有衣服吊牌）。买了个桌面增高架，调整了一下布局，也理了线，现在如果不给设备充电的话，几乎是看不到什么线的（留下来的几个咖啡豆罐子也能遮挡一下线缆，原谅自己囤垃圾了）。

{{< gallery >}}
![](https://image.guhub.cn//uPic/202505112229-my-desk.jpeg "原来我断舍离之后还是有这么多东西")

![](https://image.guhub.cn//uPic/202505121107-enderman-plush.jpeg "看到这个怎么能忍住不买")
{{</ gallery >}}

另外，我的英文博客时隔一个多月终于更新了，欢迎你来读一读：[Take Ugly Notes and Do Pretty Work](/en/posts/take-ugly-notes-do-pretty-work/)

## 💬 本期杂谈

### 为什么道德对于叙事来说无足轻重

{{< callout "文章链接：[为什么道德对于叙事来说无足轻重](https://mp.weixin.qq.com/s/Syj60Uq40i7Ng1oaz9kfLQ)" "📜" >}}

==在好与坏之间选边站不是最难的，在一团糟的情况下做出最好的决策才是。==

作者认为，无论是游戏设计还是一般意义上的叙事，道德本身并不会拉进故事与玩家个人的距离。如果故事总是给出非黑即白的选项，那么玩家的抉择就是一种「伪抉择」，因为大多数人都能分清好与坏，那样的叙事只是让玩家选择要么当个老好人，要么抱着一种猎奇的心理使坏。

想让抉择变得有意义，那就要关注**抉择的结果**，而不是抉择的道德意义，前者才是举足轻重的东西。

[#游戏开发](/tags/游戏开发) [#叙事](/tags/叙事/) [#道德](/tags/道德/) 

### 不聊天的 AI

{{< callout "文章链接：[Post-Chat UI - Allen Pike](https://allenpike.com/2025/post-chat-llm-ui)" "📜" >}}

除了聊天界面，生成式 AI 的用户界面还可以怎样设计？或者说，什么样的用户界面才能更好地让生成式 AI 在不同的场景下发挥作用？作者列出了几种 UI 模式：

1. 协作（Co-authorship）<br>
   例如 ChatGPT 的画布模式和 Cursor，这种 UI 模式的主要操作区域是代码、文档或其他内容的编辑器，聊天界面被推到了一侧，允许你与 AI 对话，与它一同编辑内容。
2. 生成式右键（Generative Right-click）<br>
   之前 Dia 浏览器的宣传片里，他们展示了一种新的 UI 模式——用户可以在浏览器内部通过鼠标右键呼出菜单，让 AI 在任何地方生成内容。关于 Dia 浏览器，我之前也写过[一篇文章](/posts/the-browser-company-的新浏览器/)讨论我对它的看法。
3. 符合直觉的搜索（Intuitive Search）<br>
   Notion 推出类似的功能已经有一段时间了，简单来说，你不需要精确地记住具体的页面名字或关键词，可以用更模糊或口语化的表达请求 AI 帮你搜索内容。我个人很喜欢 AI 的这种应用。
4. 输入代替选取（Type Instead of Pick）<br>
   类似 Raycast、Alfred、Powertoys Run 等应用程序提供的功能，用户可以快速呼出一个输入框，输入命令，而不是在一个菜单里选取。另一个应用是在需要输入日期的场景中，用自然语言（例如：下周四）描述时间，而非选取年月日。
5. 行内反馈（Inline Feedback）<br>
   就像 Grammarly 或内置在许多软件里的拼写检查，生成式 AI 可以帮忙给出更详细的反馈和修改建议。
6. 清理（Clean up）<br>
   不想思考怎么给变量和组件之类的东西命名？可以把这个任务交给 AI，让自己关注于功能的实现和真正的创作。
7. 总结和聚合（Summary & Synthesis）<br>
   令人失望的 Apple Intelligence 提供的为数不多功能中，智能摘要算是一个不错的应用，AI 能屏蔽一些不必要信息。不过，我的观点是，不如在一开始就不要接收这些信息，屏蔽掉无用的消息推送；另一点考量是，我很担心愚蠢的 AI 会把重要的信息过滤掉。
8. 语音（Voice）<br>
   比起用文字精细地描述某个需求，尤其是要描述某个存在于屏幕上的文字或媒体时，直接用鼠标指向那个元素，然后张嘴向 AI 发问，要方便得多。
9. 完成显而易见的任务（Do the Obvious Thing）<br>
   例子：按下 Tab 补全代码。我自己的体会是，这很容易上瘾，然后文件里就爆出来一大串还要花时间一条条看的与需求无关的代码。Tab 补全只能做**显而易见的事情**，而不是所有的事情。
10. 生成用户界面（Generated UI）<br>
    作者说现在大语言模型可以实时为当前用户生成特定的、符合他需求的用户界面，这让我想起了 [Oasis AI](https://oasis-ai.org)，一个每一帧都用 AI 生成的 Minecraft 游戏。作者也指出了问题，尽管 AI 可以实现用户界面的高度客制化，但这也让 UI 变得不固定，不便于用户记忆和学习

不得不说，在人们表达能力日益下降的今天，用精心设计的 UI 规避用户不会写提示词的痛点，是一个非常聪明的做法。

[#AI](/tags/AI/) [#设计](/tags/设计/)

### 创造是不刻意创造的结果

{{< callout "文章链接：[The 80-Hour Myth (Why We’re Addicted To Being Busy)](https://thedankoe.com/letters/the-80-hour-myth-why-were-addicted-to-being-busy/)" "📜" >}}

作者认为现代人痴迷于忙碌，总是停不下来，我们生活在一个神圣化忙碌的世界（a world that glorifies busyness）。作者指出，我们应该停止浪漫化苦难（stop romanticizing the long hard route）。

- 休闲（leisure）比工作更重要
- 强度比持续时间更重要
- 休息是更高效的工作形式
- 结果比时间更说明问题

话说到这里，需要展开讲讲作者所说的工作和休闲分别是什么意思。

休闲（leisure）和娱乐（entertainment）完全不同。我认为，按照作者的观点，刷短视频、看剧、听歌、打游戏，其实并非休闲，因为在进行这些活动时，表意识充满了各种信息，大脑需要不断处理这些信息（我认为这也是某种形式的劳动）。作者指的休闲是散步、读书等等，他认为这种活动会激活人的「默认模式网络」（Default Mode Network），此时潜意识会更具创造力、更有效率地向表意识推送想法，例如洗澡时头脑放空，人常常在这个时候产生新想法。在《[数字极简](/library/2024/数字极简/)》一书中，Cal Newport 也提倡读者放下电子设备，重拾闲暇时光。

作者还提到了「创作者悖论」，内容是：创造力是不刻意创造的结果。我认为作者所说的「工作」是把创意和想法变成现实的过程。如果没有闲暇时产生的灵感，工作就是无聊且漫长的。人们痴迷的忙碌可能就是在刻意地创造。作为写作者，我自己的体会时，我常常在通勤时想到刚才读的一页书、看的一篇文章或者别的什么东西，然后产生了新想法；如果我在没有想法时坐在屏幕前，我就只能挤出让我自己都不忍直视的文字。

作者还认为拖延不是懒惰，而是人类的机制（How we are wired）。拖延、分心、放空，让自己从专注状态中脱离一段时间，刻意地休息，能给工作带来更多想法，提升效率和质量。作者提倡「像狮子一样工作」，关注自己的在乎的事情，受到强大的欲望推动，在短时间内高强度地完成工作，然后停下来休息。

[#创作](/tags/创作/) [#效率](/tags/效率/)

### 我宁愿读提示词

{{< callout "文章链接：[I'd rather read the prompt](https://claytonwramsey.com/blog/prompt/)" "📜" >}}

作者认为机器不会产生「想法」。比起机器生成的看似详尽准确的长篇大论，他更愿意读人写下的那一两句提示词，因为人写下的那一两句话就代表了他此时此刻的想法。比起计算机程序的东拼西凑，人的想法才是最有价值的。在我看来，「为什么问这个问题？」「为什么要这样措辞？」「为什么在这个时候问这个问题？」都是更值得关注的点。

作者的主要观点有三：

1. **思考比写出东西来更重要**。写作的目的不是写出字来（create an artifact of text），而是为了促进思考，让 AI 写东西没法促进人的思考。Peter Naur 认为，写出足够复杂的程序不止需要写出代码来（the artifact of code），还需要程序的理论，程序员必须理解代码背后的逻辑结构。
2. **Vide Coding 有严重的安全隐患**。Vibe Coding 指完全不自己写代码，而是向大语言模型描述需求，完全让 AI 编程。我自己 Vibe Code 一些脚本来实现简单的需求，但对于更大的系统，例如一个完整的应用程序，Vibe Coding 就会造成很多问题。之前有一个 Vide Coding 产出的应用就泄漏了它的 OpenAI Keys。我认为在很长一段时间内，AI 都没法独立地编程，程序员也不可能不理解任何软件工程知识，光靠 AI 工具做软件。
3. **AI 十分擅长创造垃圾**（great for making nonsense!）。这一点与第一点关联很大，AI 很擅长「写出字来」。作者是这生成了一段文本，其中包含这样一句话「Perhaps it stems from a desire for efficiency, a wish to quickly generate text without the perceived effort to crafting each sentence.」，作者认为「Perhaps people do it for efficiency.」几个词就足以表达观点。语妙不在言多，词藻的堆砌不过是思想贫瘠的遮羞布。

总的来说，作者认为创作的意义就在于**分享个人体验**——如果没有体验，为什么还要写出来？如果一个东西都不值得人自己写，那也不值得读。

[#AI](/tags/AI/) [#创作](/tags/创作/) [#写作](/tags/写作/)

### 你更有可能死在自己的生日吗？

{{< callout "文章链接：[Are you more likely to die on your birthday?](https://pudding.cool/2025/04/birthday-effect/)" "📜" >}}

很有趣的研究。作者找了当地的死亡数据，发现人们的确更有可能在自己生日当天去世。他用了各种统计学方法证实了这个观点。背后的原因可能是死亡率和出生率都有一定的季节性规律，例如冬天的死亡率更高。某些月份更容易死人，某些月份有更多的人出生。因为死亡和出生在时间上的分布是不均匀的，也就造就了死亡发生在生日当天的概率要高一些。

具体的论述可以阅读原文，我的数理统计知识较匮乏，不太能写得明白。

作者用了很多不同的数理统计方法证明这个观点，每次的结果都有些差异，他在最后写道，研究的结果会因为我们采用的不同方法而产生差异。我在《[为何以及如何建立自己的信任链](/posts/为何以及如何建立自己的信任链/)》中写道，普通人在阅读参考文献时，应当关注被采用的研究方法。

[#死亡](/tags/死亡/) <span style="font-size:80%">想不到这段内容也能归到死亡这个标签下</span>

## 🗣️ 碎碎念

没话写，在思考我究竟还需不需要在周刊设置一个「碎碎念」栏目。
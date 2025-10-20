---
title: Java Development with True IDEA
date: 2025-10-20T16:16:00
draft: false
toc: true
---

Let's be totally honest. **IntelliJ IDEA is bloated, slow, and buggy.** It's not the worst tool but certainly doesn't provide a comfortable dev experience, especially when it's so damn expensive. Not to mention, it eats up precious storage space on my MacBook (that one's probably on Apple though).<!--more-->

## Stop taking things for granted (and making newcomers do the same)

The one thing I don't understand is that people urge Java beginners to download this software filled with magic buttons when they don't understand how everything works. Personally, I never fully understand the magic behind "Run/Debug" button — It does too many things! Depending on the project type and configuration, it can compile and run a single java file, or it can build, deploy, and start a local server.

When I started out learning Spring, people just told me to configure this and configure that, simply asking me to follow a set of instructions that seemed to come out of nowhere, with no further explanation. (I seriously doubt they understand it themselves) And then, I was told to hit the "Run" button. And boom. There's your project up and running. Start Developing!

Software like IDEA makes development far more complicated than it needs to be, with too many layers of interfaces. It's probably a fair trade-off for large teams, but certainly not for everyone. And there are better solutions. I can put up with photo and video editing apps like PhotoShop having complex UIs — there's no other way to do it. But for software development? Why would I need all those sections and subsections and sub-subsections in the setting panel, when a bunch of configuration files can do the work just fine?

With the IDE interface being bloated, learning Java development becomes two things: The basic language knowledge (the syntax, common libraries and best practices) and **HOW TO GET THIS GODDAMN THING TO F\*CKING WORK**. 

I don't see C programmers losing their minds because their editors break. In fact, let's appreciate the beauty of C development. You got a simple editor, a trusty compiler, and... nothing else.[^1] Your compiler tells you what goes wrong, and you fix it. That's it! JavaScript development also requires only an editor and a set of CLI tools (build tools and package managers) to work. There might be tools like ESLint and husky that add to the workflow, but they're introduced by choice. We know what we're doing and we know **why** we need them.

An integrated development environment robs you of free choices and deliberate thinking, making you depend on defaults and automation you don't fully understand. It's got everything they think you need, covering it up with magic, and you just stop thinking and take them for granted along the way. 

That's stupid. Let's just stop. Even if you do use an IDE, that should be the result of your own judgment. And trust me. You'll probably find that you don't need one.

## The beauty of command-line interface

Almost every common task can be done with the right CLI tool. For instance, I never really used the database panel in IDEA. I just enter `mysql` or `mariadb` in my terminal and type `SELECT * FROM ...`. SQL command not only works, but also is clear. You won't get lost with interface navigation when you're just typing text commands.

It's true that CLIs are harder to learn. Beginners might take a while to understand what each command does and to remember their names. But **the clarity is precious**. You don't just hit "Run" and expect magic to happen. You type `mvn spring-boot:run` and realize you're starting a Spring Boot application with a Maven plugin. The server setup and annoying configurations are already simplified by Spring Boot, so a line of command can do the work. And if something goes wrong, you know what happened exactly. You don't have to navigate through tabs just to find where the error log is! 

The fact that I have to click through multiple menus to find what I need is insane! Everything should be at most one-click away and I should be able to navigate only using my keyboard! Mouse is the slowest input device if you're not gaming.

For IDEs based on graphic interfaces, if you want to create a Spring Boot project, you need to click a "New Project" button. And then select "Spring Boot" in the left panel. And then click the inputs appeared on the right panel to edit names and stuff. And click "OK". And scan a list of dependencies where you probably don't recognize most of the items, and then select the dependencies you think you need. And hit "Continue". Finally, after seconds (which can be really long depending on your machine) of loading, you're in a whole new interface with panels and sidebars and split-views!

For command-line interfaces, you copy-paste and edit this command in your shell:

```shell
curl https://start.spring.io/starter.zip \
  -d dependencies=web \
  -d language=java \
  -d name=demo \
  -d packageName=com.example.demo \
  -d type=maven-project \
  -o demo.zip
```

Unzip it, open it with any editor you'd like (which is usually way faster than IDEs), and voilà — You got yourself a Spring Boot project ready to go. If you create new projects very often, you can create snippets, templates, scripts, or aliases in any way you want, simplifying this action to a single line of command.

One of the things I hate about IDEA is the "Maven" button. It only appears in `pom.xml` tab and sometimes disappears for no reason, making a simple dependency update unnecessarily confusing. I'll have to navigate to a drop-down menu to sync my Maven dependencies. That's stupid. Just type `mvn clean install` and hit enter.

If you want to push to the extremes, use Vim or Helix, an editor running in your terminal. Then all of your development can be done in one place, a place faster and more flexible than any IDE.

## The True IDEA

- **Instant.** After getting rid of the bloat, your project opens instantly. Goodbye to being stuck in the loading screen!
- **Deliberate.** Every action you take is clear and judicious, for now you know exactly what every command does. No magic. Pure tech.
- **Elegant.** No over-categorization of tabs, sections, trees, or the things I don't even bother to name. In your sight is only the things you truly need, hand-picked by yourself.
- **Attentive.** Focus on coding rather than learn how to make sense of such unnecessarily complicated software.

Plus, there's more.

- **Universal.** You now possess the tools to develop in any language with any tech stack. Who needs a 5-gigabyte software that mainly targets 2 languages?
- **Production-like.** If you build using command-line, you can switch to a Linux server smoothly, where every action is done using commands.
- **Flexible.** Tear down any part of your workflow and replace it if you don't see fit. You have full control over your development environment.

My current setup is Zed + Ghostty for Java and JavaScript development. I also poke around with Go, Lua, Python, and other tech with that. With the right CLI and editor extension installed, there is zero friction.

[^1]: There are massive IDEs built for C, like CLion, also a JetBrains product. Honestly, I don't really see many people using it. Maybe I'm biased. Correct me if I'm wrong.

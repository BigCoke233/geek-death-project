---
title: Java Development with True IDEA
date: 2025-10-20T10:45:00
draft: true
toc: true
---

Let's be totally honest. **IntelliJ IDEA is bloated, slow and buggy.** It's not the worst tool but certainly does not provide a comfortable dev experience, especially when it's so damn expensive, not to mention it's also taking tons of expensive storage space on my MacBook (this one's probably on Apple tho).<!--more-->

## Stop taking things for granted (and having new-commers to do the same)

The one thing I don't understand is how people urge Java beginners to download this software filled with magic buttons when they don't understand how everything works. For me, I never fully understand the magic behind "Run/Debug" button —— It does too many things! It runs a single Java file with a main function but also build a WAR package, start a local server and runs your application. When I started out learning Spring, people just told me to configure this and configure that, simply asking me to follow a set of instructions that seems to come out of nowhere. And then, I was told to hit the "Run" button. And boom. There's your project up and running. Start Developing!

Softwares like IDEA make development unnecessarily complicated with too many user interfaces. I can put up with photo and video editing app like PhotoShop being complex in user interface, for there's no other way to do it. But for software development? Why would I need all those sections and subsections in the setting panel when a bunch of configuration files can do the work just fine?

When the interface is like this, learning Java development becomes two things: The basic language knowledge (the syntax, common libraries and best practices) and HOW TO GET THIS GODDAMN THING TO F\*CKING WORK. 

I don't see C programmers losing their minds because their editors break. In fact, let's appreciate the beauty of C development. You got a simple editor, a trusty compiler, and... nothing else. Your compiler tells you what goes wrong, and you fix it. That's it! JavaScript development also requires only an editor and NPM to work, as simple as C. There might be tools like ESLint and husky that adds to the workflow, but they're introduced by our choice. We know what we're doing and we know why we need them.

Integrated development environment robs you of free choices. It got everything they think you need, covering them up with magic, and you just stop thinking and take them for granted along the way. 

That's stupid. Let's just stop.

## The beauty of command-line interface

Apart from the convenience of creating a new project using out-of-the-box templates and archetypes (spoiler alert: CLI can do that too!), almost every functionality of IDEA can be replaced with CLI tools. For instance, I never really used the database panel. I just enter `mysql` or `mariadb` in my terminal and type `SELECT * FROM ...`. SQL command not only works, but also is clear. You won't get lost with interface navigation when you're just typing text command.

It's true that CLI are harder to learn. Beginners might take a while to understand what each command does and remember the names. But **the clarity is precious**. You don't just hit "Run" and expect magic to happen. You type `mvn spring-boot:run` and realize you're starting a Spring Boot application. The server setup and annoying configuration is already simplified by Spring Boot. A command can do the work. And if something goes wrong, you know what happened exactly. You don't have to navigate through tabs just to find where the error log is! The fact that I have to navigate multiple times to find the things I need is insane! Everything should be at most one-click away and I should be able to navigate only using my keyboard!



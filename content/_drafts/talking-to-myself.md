---
title: I like talking to myself (with Git and GitHub)
date: 2025-12-16T17:27:00
draft: true
toc: true
---

I am an introvert. I like talking to myself. I talk to myself in spoken language when I get comfortable in a place where nobody's around. Recently, I find myself talking to myself in written words on Git and GitHub, which could actually be helpful in an engineering perspective, not just emotionally.<!--more-->

It's established that following a commit message specification, like [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/), is necessary for collaboration. Clear commit messages help collaborators understand the changes you've made, and how those changes can affect them. But... What if you're doing something solo? For instance, you're working on a side project that you know is going to be small, and you don't expect anyone to help you expand it. Toy code does not need that specification, right?

Well, nope, at least not for me.

Writing commit messages is not a responsibility thrust upon you. It's a delight, actually. More precisely, any documentation should be fun for developers. You don't do it for the sake of users or other developers. You do it because it feels right.

I like having my thoughts straightened by writing. Language is a structural presentation of human ideas, which start out chaotic with countless loose ends, but then slowly fit into place when you're trying to articulate them. The ancient Greek word "λόγος" can both mean logic and words. It's inspiring if you think about it: any logical flaw is easy to identify when you're putting things into words. Either spoken or written, you're making things clear, and clear is good. 

We want our work to be clearly described so others can understand, but we also do that to make ourselves understand, in the present and after weeks or months.

When I'm working on a side project, I keep a decent git history, with atomic commits, straightforward messages, and a lot of branches. I also create Pull Requests rather than direct `git merge` or `git rebase`, even though the code is reviewed by me and only me. I write issues that nobody reads, just to keep track of a bug I was trying to fix. I comment my PRs and issues with fact checks and progress updates. I add notes on installing dependencies and starting a dev server in `README.md`, stored in a private repository. Basically, I report to myself.

How is doing extra work a delight?

1. I use words and logic to keep a clear head.
2. I enjoy seeing things well organized.
3. I feel safe when there's no line of code I don't understand.
4. It feels like I'm presenting my work to the world.
5. The future me will be happy when he's trying to get my ancient code working.

Rob Pike says "Documentation is for users". I'd like to say it's also for the writers themselves.

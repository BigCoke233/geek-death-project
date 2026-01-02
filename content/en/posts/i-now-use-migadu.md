---
title: I now use Migadu for my emails.
date: 2026-01-02 12:13:24
draft: false
---

I like emails. They're quiet and you don't have to rush. And they're decentralized. You get to choose who stores your data. If you're on WeChat, Line, or whatever IM apps you're using, you're forced to stay there because that's where all your contacts are. Emails are portable. Save your mails as `.eml` files and store anywhere. Keep your address, if you own your domain name, and migrate to wherever you want, preferably someplace you trust, and people can still mail you with the same address.

I was using Tencent Exmail for a long time. It's an evil service provided by an evil big corporate if you don't know, as evil as Google. It's hard to get a custom domain email address in China. People rely on emails less and there are fewer options. Some people use Cloudflare or similar services to create an alias for their email, but that's still your data in a BigCo's data center. What you get is merely a cool address.

Not really paranoid about privacy but I do feel more comfortable if my mails don't go through people I don't trust. So I started looking for new emails services.

I searched and asked ChatGPT, ruled out some hard "No"s, and got 3 options:

- PurelyMail
- MXRoute
- Migadu

They're all small services. Best of all, they're cheap, like insanely cheap. Migadu's micro plan costs $19.9 annually. PurelyMail costs $10. MXRoute was defintely the cheapest, $15 for 2 years! But that offer is only available when they're doing a promotion.

They all offer unlimited domains and unlimited addresses. Only Migadu comes with a few constraints, 200 emails in per day and 20 out per day, and also you only get 5GB storage (But they're all soft limits. There is tolerance).

Eventually I chose Migadu, for two major reasons:

1. PurelyMail and MXRoute are both run by one man. They're not a corporate. It's good but I keep worrying what happens if their owner gets hit by bus some day... Sorry, it's not a nice thought. Another downside is that you don't get much support from your service provider. It could be really frustrating if something unexpected happens. You're on your own.
2. Migadu is a small company located in Switzerland, where EU laws ensure the best privacy. Privacy was the whole point I'm migrating my email service after all. It's a reasonable choice.

Migadu is surprisingly honest about their service. They have a [Pro/Cons](https://www.migadu.com/procon/) page (where pro is actually absent) listing all the reasons why you probably don't want to use Miagdu. For example, they're made for human emails only, meaning they don't send bulk transactional emails. They are honest about their pricing not being the cheapest, their spam filters not being perfect, their server not being in Switzerland but in France... And they offer fully-functional trials, so you know exactly what you're paying for.

I love how opinionated they are. They do things their way and they are honest about it. They don't dazzle customers with fancy words. They don't even use Web Analytics or track their users in any way!

The setup was smooth. DNS configuration can be intimidating even if you have experience, because configuring email systems is way more complicated, with multiple TXT, CNAME, SRV records. Merely 2 MX records with not work. They have clear documentation on setting up your DNS. All you have to do is read, understand, and copy. Within 10 minutes, I can receive and send emails using Migadu with no problems.

I posted on Fediverse and asked people to send anything to me for testing. And I ended up receiving some really interesting mails. If you'd like, mail me at `hi@guhub.cn` and send whatever you want! I'd be happy to make some new friends here.

So, I now use Migadu for my emails. And I hope you also enjoy your email service :3


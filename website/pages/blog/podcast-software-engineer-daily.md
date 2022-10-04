---
title: Open Source Cloud Asset Inventory with Yevgeny Pats @ Software Engineer Daily
tag: guest
date: 2022/06/14
description: >-
  Podcast episode with Yevgeny Pats hosted by Alex DeBrie. In this episode we’ll
  discuss CloudQuery, Yevgeny’s entrepreneurial background and raising funding
  with an open source project.
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>



### Play podcast at:

- [Software Engineering Daily podcast website](https://softwareengineeringdaily.com/2022/06/10/open-source-cloud-asset-management-with-yevgeny-pats/)
- [Apple Podcasts](https://podcasts.apple.com/us/podcast/open-source-cloud-asset-management-with-yevgeny-pats/id1019576853?i=1000566051477)

---

We also provide transcript below:

## Transcript

**[00:00:00.370] - Alex DeBrie**

Nearly all new tech companies build in a public cloud, and established companies are rapidly migrating applications to the cloud from their on-prem data centers. But this move to the cloud can lead to a visibility problem. Cloud providers offer not only compute instances, but also manage services like databases, blob stores, queues and more.

**[00:00:19.350] - Alex DeBrie**

It can be difficult for SRE teams and security departments to understand what is happening across a company's cloud accounts. Yevgeny Pats is the creator of CloudQuery, an open-source cloud asset inventory powered by SQL. CloudQuery allows you to ingest and structure the resources in your cloud accounts so that you can query them using SQL.

**[00:00:39.470] - Alex DeBrie**

This allows SRE teams to understand the source of specific resources, while security teams can ensure compliance with policies. In this episode, we'll discuss CloudQuery, Yevgeny's entrepreneurial background and raising funding with an open-source project. All right. Hey Yevgeny, welcome to Software Engineering Daily.

**[00:00:57.770] - Yevgeny Pats**

Hey Alex, thanks for hosting me. Excited to be here.

**[00:02:37.570] - Alex DeBrie**

Absolutely. You're the founder and creator of CloudQuery. Can you just give us a little bit of background on what CloudQuery it is and what it's doing for people?

**[00:02:46.820] - Yevgeny Pats**

Yeah, sure. We started as an open-source project early last year. We started as an open-source called Assets Inventory. Basically, connecting to all your cloud services, cloud APIs, then striking all the configuration, transforming it, and loading into relational database into Postgres.

**[00:03:05.000] - Yevgeny Pats**

Maybe you can also think about this as like Terraform, but the other way around, not provisioning anything, but just extracting all the configuration. Kind of an ETL layer for cloud configuration. We started it because we thought, one, because I didn't find anything like that and I really needed it. Two, because I thought that it's a corner respond to a lot of use cases in the cloud, either a security, cost, visibility.

**[00:03:35.230] - Yevgeny Pats**

You have like a sprawl of enterprise closed source vendors that I think their development teams, like 90% of the time are implementing closed source version of CloudQuery and everyone has their own small closed source version and there is no open-source. So we decided to go with the open-source way and this also gives the ability to users to also contribute back because there are infinite number of APIs out there and there is no one vendor that can support all of that.

**[00:04:07.370] - Yevgeny Pats**

The only way that we thought we were able to solve it is to go the open-source way and like the similar reason why Terraform is open-source, because you want to give people the ability to contribute back something that they need and the vendor just didn't get into like implemented yet. I can talk a bit later about some of the use cases, specifically for users and how people use us and deploy it.

**[00:04:31.950] - Alex DeBrie**

Awesome. I love that description as a reverse Terraform, and I think the open-source nature is interesting. I want to get into that later, especially like as a funded company and open-source, how we do that.

**[00:04:41.300] - Alex DeBrie**

But I think you make a great point where if you're building this framework of how do you input resources, get them into this thing to make them queryable, then that open-source nature of it just allows so many people to come and do the long tail of things, whether it's a different cloud provider or different type of resource that other people aren't using, they can get those in there, so I love that.

**[00:05:00.980] - Alex DeBrie**

You said as you're looking to make this, you wanted something like this and you couldn't find anything like it. Tell us about that. Why did you want this to exist? Getting your resources into a relational database? What were you looking to do with that?

**[00:05:13.840] - Yevgeny Pats**

When the industry... Around like security, [inaudible 00:05:18], cloud, the last 12 years or so I played build SaaS apps. I was also on the other way, not only from the vendor side but also from the builder side, so it's like close to me. But the other thing that the big shift that I saw...

**[00:05:34.270] - Yevgeny Pats**

My first startup was Super Enterprise, and that may be like a long answer to short question, but my first startup was Super Enterprise email security, top-down startup and I've been there for two years. But then I understood that I don't want to run a company where there's get a demo button.

**[00:05:51.630] - Yevgeny Pats**

So last five years, I was focused solely on the developer's page because for me, all the classic vendors didn't make sense anymore because I wouldn't use it from the user perspective. But a lot of the companies still in the security space are like top-down.

**[00:06:09.580] - Yevgeny Pats**

You can also go to Palo Alto Networks checkpoint, but also like kind of next gen startups, especially like in Israel. I said, "Okay, let's see what's the DEV alternative, open-source alternative." I didn't find any, and that's basically how I started. There were a few open-source projects trying to tackle it, but nothing that really hit it out of the park with good architecture and good adoption.

**[00:06:37.330] - Alex DeBrie**

Cool. You mentioned it's open-source. You have a bunch of different providers and things. Can you give us a sense of what it supports, currently? I imagine you're thing the major cloud providers, but what other sort of providers and things? Then how deep are the resources within some of those providers?

**[00:06:52.310] - Yevgeny Pats**

Yeah, so we definitely started as focused on the big three cloud providers like AWS, GCP, Azure, Digital Ocean, and a few smaller providers. But yeah, the focus was really on the big cloud providers. We have also third-party providers. For example, Yandex Cloud is maintaining their own integration and on provider, which is pretty neat. We're trying to expand that and help users to do that more easily as we go with better developer experience.

**[00:07:23.220] - Yevgeny Pats**

But we already saw some good third-party integrations to CloudQuery. Regarding the depth of that, we try to do pretty deep and support as much as APIs as possible, especially like the important one. I think in AWS config, we support potentially like even more and same thing in GCP, even more resources than AWS config or in Google Cloud access inventory, which is like the enterprise, like the cloud solution because actually, their team, they have the same problem. They work with the same APIs and their solution is closed doors, so they hit the same. Even though they are in Google or in AWS, it doesn't help them to cover all their APIs.

**[00:08:06.310] - Alex DeBrie**

I think that's so interesting. It's the same pattern we've seen with Terraform over the years, where Terraform has better coverage than CloudFormation. I think the CloudFormation team and things, they're getting better about that but it's just amazing when you have that open-source nature. Someone in the community is going to want to contribute that so they can use it and it's just pretty amazing.

**[00:08:24.130] - Alex DeBrie**

Okay, so imagine I'm someone that I have a lot of resources in AWS. I want to get a sense of this. I pick up CloudQuery. How do I get started? What's that process look like from getting started to getting into a relational database? What's happening there?

**[00:08:36.630] - Yevgeny Pats**

Yeah, there are like a few onboarding path. Like the first step, usually it's just a go. It's written in Go, a single binary and have really similar architecture to Terraform. We use a lot of the underlying Terraform libraries, like it's called Go Plugin to download all the plugins and run them locally. But yeah, it's a single binary.

**[00:08:58.450] - Yevgeny Pats**

You can download it locally and then you just run like Postgres in the docker, in the local container. We can load all the information in your local Postgres, you connect it to your AWS account and that's basically it. We support also AWS-specific features like AWS Orgs. So you can connect us to read only your audit permissions so we can extract all your accounts under your organization, go through them, extract all the configuration across all your accounts and people ran us on thousands of accounts.

**[00:09:36.310] - Yevgeny Pats**

I think the biggest account, the biggest one at least on GCP that we saw from one of our users, is like tens of thousands of GCP projects. The nice part is it's really about rebuilding those use cases on top of a data platform. Once you have the data there in a relational structured way, then you can connect it to whatever you use like Grafana. Some people connect us to Grafana, to Metabase, to Apache Superset, whatever you have. It doesn't matter. Just a database.

**[00:10:09.250] - Alex DeBrie**

Yeah, cool. Then so once I've got all my resources into CloudQuery, what are some of the practical use cases or maybe even teams or different things that people are doing with that data?

**[00:10:20.670] - Yevgeny Pats**

Yeah. It's mainly split across like two teams: the DevOps and SRE teams that fall under the same bucket. The second one is more on the security side. Sometimes it's also like the DevOps team also takes that. But on the DevOps SRE team, it's usually just around monitoring and visibility.

**[00:10:39.730] - Yevgeny Pats**

You set up alerts, for example, in Grafana or Metabase, like sending you daily reports when something changed, when you have, for example, more compute than you want to, or some API where you have APIs enabled that you don't want to be enabled in some project, like whatever rules. You can set up any rules with SQL. Basically, you use SQL as your rule and query engine. The second use case is around security. So running like putting the guardrails like again, with the standard query language.

**[00:11:14.450] - Yevgeny Pats**

We also implemented something called CloudQuery Policies. Actually just another structure on top of SQL. It's just like you run the same SQL and you can run them like as a batch. You put all your policies and rules and then you can notifications and so on. Another thing on the DevOps side that we saw, like in this case we saw quite a bit is not only giving this visibility to DevOps team, but also giving this visibility to the developers.

**[00:11:45.840] - Yevgeny Pats**

Let's say you're like a thousand people organization and you have thousands of accounts. Not everyone has access to all the accounts and they shouldn't. But also, for example, you're debugging something or you're seeing something in the logs. Now you want to understand, okay, in which cloud it is, like this IP or located and then in which project, who should I reach out to where it's connected? Just kind of like an infrastructure search use case, I would say for developers.

**[00:12:20.590] - Alex DeBrie**

Nice, I love that. One thing I want to get is just a sense of scale for query users. Like how big are these projects they're having? You're saying people with 10,000 GCP projects, how many resources are we talking about? There are people pushing a hundred million. What's that look like?

**[00:12:37.040] - Yevgeny Pats**

We saw 100, 300,000. Most of our early doctors that we saw coming to CloudQuery is usually big companies that they don't have a good enough alternative basically that... This is also partially why we built it. But yeah, companies like Tempus, Horsley and other big organization, using us and production, which is good. We're working on bringing more.

**[00:13:06.140] - Alex DeBrie**

Yeah, absolutely. I imagine just like with 10,000 projects and trying to figure out which S three-bucket or like you're saying which IP address and who owns that, it's just quite the task unless you have something like this. With that notion of scale in mind, I want to talk a little bit about the underlying technology.

**[00:13:22.030] - Alex DeBrie**

I know you mentioned Postgres, things like that, I guess. Were there any particularly hard problems that you had to design for as you're going either in the fetching part, in the storage and querying? What was hard and tricky about this?

**[00:13:35.320] - Yevgeny Pats**

Yeah. There were like a few tricky things to solve. I think we are also still solving a lot of them and trying to put a lot of focus on the developer experience because actually we have two types of users. One is really users, people who use essentially the Postgres or deploy, it's the DevOps team, the security team, or just the developers.

**[00:13:57.130] - Yevgeny Pats**

The two is actually real developers of our SDK. Essentially, we built an SDK, where you can build your providers and we have some users that already did it. Then it's a completely different developer. You have to make sure your APIs and SDK is in very high production rate so people can really use it and it can scale.

**[00:14:22.670] - Yevgeny Pats**

The other thing is the developer experience and the usage and that it can support high-scale accounts. The quality, we put a lot of focus on quality over quantity and still especially because it's adapt tool like self-serve. Adapt tool where we don't have any sales team to support it, you come to our GitHub, you try it. Either it works and then if you stay on call free and either it's not, then we don't know about it.

**[00:14:52.870] - Yevgeny Pats**

Some of the hard problems, as you say, is around like the fetching. So we put a lot of work into scheduling and optimization and how to fetch the most resources and short amount of time, so that was one. The second one was really designing and building CloudQuery SDK on the developer perspective and then putting a lot of effort internally.

**[00:15:16.960] - Yevgeny Pats**

Also for external integration but also internally, how do you scale that in terms of both development velocity and testing? Okay, we want to increase our team and we want to support more APIs. I see full request. It's a lot of APIs to support that we want to add and we want to be able to add them fast or to other people to contribute. I need to know that, okay it works and it's easy to write tests for that.

**[00:15:44.260] - Yevgeny Pats**

Testing infrastructure is quite tricky, so we have Terraforms to test it and we try to make really contributor developer also quite easy, and we're still working on making it even easier. But that was a lot of the engineering and that's still going on and we try to put focus on.

**[00:18:23.330] - Alex DeBrie**

Do you have a lot of integration tests against live accounts with real resources? I feel like Terraform had that problem just expense-wise of creating and tearing down resources.

**[00:18:33.720] - Yevgeny Pats**

Yeah, so we had. I think it's partially solved. The best thing that we found, for every resource that we create, we want to test it against live environment, so the contributor has to write Terraform for that. But we deploy this resource in real environment and then we test against that okay, we are able to fetch like all the data that we're expecting to.

**[00:18:59.610] - Yevgeny Pats**

Actually, we are not even tearing this down because we want people to be able to... We don't want people to wait for this resource to be created which will be extremely lengthy and we don't want to do nightly tasks. You committed something in the morning and then it broke in the night and getting into all this fakeness, so it's just there.

**[00:19:22.190] - Yevgeny Pats**

Apart from maybe things like are very expensive that we are skipping. We try to have live and test environment and maybe as it will grow, we'll have to think about something smarter. But for now, the cost is not that high so we're just trying to keep it simple and test it against real environment.

**[00:19:41.540] - Alex DeBrie**

Yeah. I wonder if the clouds would help you out with some credit on that too. I imagine you're helping with adoption in different ways or reducing curriculum. It's interesting to see.

**[00:19:51.190] - Alex DeBrie**

You mentioned you spend a lot of time on scheduling and efficiently polling things like that. If I have an organization that has lots of accounts that has 100,000 resources, how long does that CloudQuery fetch take to go and pull those up? Is that minutes? Is that hours? Like what's that look like?

**[00:20:08.530] - Yevgeny Pats**

Yeah, it depends on the computer that you have and the number of CPUs. That's the ETL workload. I think one of the reasons that we also chose those, one it's releasing in terms of distribution, just one binary without dependencies. I think it was one of the main reasons.

**[00:20:28.430] - Yevgeny Pats**

But the second one is, actually, it's very good in concurrency. You can just create very lightweight working and we use that. I think if you have strong enough machine, let's say even like eight CPUs, maybe you can go like 16, but you can get it maybe even under an hour.

**[00:20:46.980] - Yevgeny Pats**

But I will say that the ballpark. If you have really that extreme big accounts which should work because you can even fetch it. We saw people fetch it and we also fetched internally like every six hours. It's usually good enough for us.

**[00:21:02.160] - Yevgeny Pats**

Some people do it like every two hours. We didn't see a lot of requests to have real time, like completely real time. If you have two, it usually should be good enough.

**[00:21:13.430] - Alex DeBrie**

Are most people hosting CloudQuery somewhere in the cloud or do you have a lot of people just run on their laptops and checking it there? How do you see that?

**[00:21:20.180] - Yevgeny Pats**

Yes, exactly. The first, onboarding guys to run it locally, just to play with it, see it works, that it answers your use cases, whether it's search or visibility or security rules. The second one is he has to deploy it on production and to run periodically.

**[00:21:38.970] - Yevgeny Pats**

You can run call query in a lot of ways. We just give some guidelines and premade deployments for Kubernetes just to be consistent and not support a lot of different types of deployments.

**[00:21:53.100] - Yevgeny Pats**

We have [inaudible 00:21:54], which you can deploy on your Kubernetes and it will run in the [inaudible 00:21:57] and you can connect it to your RDS. But we also provide Terraforms for GCP and AWS to provision the infrastructures if you need like, actually together with the Kubernetes as well so you have one deployment so you don't need to jump from Terraform and Helm.

**[00:22:17.010] - Yevgeny Pats**

But actually what the Terraform is doing is deploying the Helm after it provision the EKF cluster, the database, or in GCP, the GKE in the cloud SQL.

**[00:22:29.870] - Alex DeBrie**

Does CloudQuery support Kubernetes in the sense that I can get a more granular look at what my Kubernetes cluster is doing? Like maybe I have 100 nodes running but I have all these different services on it, can I get visibility into that?

**[00:22:44.130] - Yevgeny Pats**

Yeah. We have a Kubernetes integration provider. I would say it's like beta, so it supports some of the resources. I'm not sure what the granularity of the visibility you get. You can plug it in.

**[00:22:58.480] - Yevgeny Pats**

Something that actually people ask is to connect the Kubernetes provider, for example, with the GKE. After you patch all the information from GCP and configuration about GKE is to go ahead and use the Kubernetes provider for the Kubernetes clusters on your cloud provider. Whether a GKE or a EKS to get even more granular data.

**[00:23:22.700] - Alex DeBrie**

Yes. One last thing just on the underlying tech. I read a blog post about you looking at timescale and things like that. Do you have historical looks at the resources I have or is it always like, hey, this is what I currently have and it's going to get overwritten every time when I do a fetch?

**[00:23:39.560] - Yevgeny Pats**

Yes, this one was really tricky to get right as well. We try different things on what will work. To have historical view, you basically need to maintain handwritten migrations for every resource, which is quite impossible from development perspective.

**[00:24:01.590] - Yevgeny Pats**

But then we talked a little bit more with some of the users. For them they actually said that they don't necessarily need it all in one schema at the same table to always create the time, but it's more like for reactive or investigation purposes.

**[00:24:17.670] - Yevgeny Pats**

Okay, I want to go back and look at what I had this date. Load it and then look at what I had there. Some of the things that we'll probably look into in this scenario that they said, it's actually more of a data warehouse in data lake. You want actually to store it where it's cheap and then only like query when you need it.

**[00:24:36.100] - Yevgeny Pats**

This is something that we will look into in terms of supporting actually data lakes like Snowflake, BigQuery and putting it there. Then you can also both use other databases and also use it for historical reasons. But in a sense you can also do it outside CloudQuery.

**[00:24:54.850] - Yevgeny Pats**

Just back up Postgres periodically. But we want to give more native support right where you can connect CloudQuery to a data lake and then, okay, it's stored there. If something happens, I have another point of information that I can go back and use if I need to investigate something.

**[00:25:13.220] - Alex DeBrie**

Yeah, cool. It could also be cool to have some post fetch roll up aggregation queries. Like maybe you just aggregate some of the data into, you're saying like, hey, on this date, I have this many EC2 instances in each account. Just maintaining that, the smaller number of rows to maintain over time but then you could still get some historical stuff or things like that if you want to.

**[00:25:32.670] - Yevgeny Pats**

Yeah, exactly. For one table, this is something you can do. Like you can maintain migration for exactly one table but not for hundreds of tables. This will be one option as well.

**[00:25:43.580] - Alex DeBrie**

Yeah. Cool. That's interesting. I want to shift gears a little bit and go into your background because you said you've started a few different companies. This is your third one.

**[00:25:50.780] - Alex DeBrie**

Can you just give us a little bit on how you got entrepreneurship, maybe some of the companies you have an acquisition in there, some cool stuff?

**[00:25:57.250] - Yevgeny Pats**

Yeah. I guess I could start from the start. I guess I was always, like my childhood dream starting a company. I don't know why. Maybe I read a lot of stories about Mark Zuckerberg. It was exactly like 2008. I finished high school. Yeah, somewhere there.

**[00:26:17.370] - Yevgeny Pats**

Then I joined the Israeli Cybersecurity Intelligence Unit, and I've been there for four and a half years. I was a lot in to computers there early on and iPhone happened then. I guess it's all accumulated all together. It was kind of [inaudible 00:26:34] especially, like really from the start also started like a few projects while in the army, like bootstrap things.

**[00:26:40.980] - Yevgeny Pats**

But then after the army, I joined ASA like cyber security startup in 2013. It was a small startup, was called Hyper Wise and it was acquired by Checkpoint after eight months. There wasn't even one customer.

**[00:27:00.890] - Yevgeny Pats**

This was my first start of experience and I said, well, that's easy. You just start a company, you working on eight months, you don't need to have any customers and then just Checkpoint acquires you for a bunch of money. I didn't stay in Checkpoint, I had the retention.

**[00:27:15.490] - Yevgeny Pats**

But I said, if that's that easy, there is no sense to stay for whatever. Basically it was good experience but from the start, actually, it was pretty bad. It was learning from completely the wrong example.

**[00:27:31.660] - Alex DeBrie**

It taught you the wrong lesson. Yeah.

**[00:27:33.610] - Yevgeny Pats**

Like back then I was sure that's the way it is. I thought it was my own eyes and then I thought and the funny thing that I had two good friends that were working in small startups in the same time and they had the same experience. It's just like that, there was just some micro acquisition spree of just very small companies and we all learned the wrong lesson.

**[00:27:57.160] - Yevgeny Pats**

We went, we raised some money for short term. Initially, we didn't even know what we were doing, but cybersecurity was hot back then. I guess it's also hot now. We got some money. After we got the seed of around of $2 million, we started talking to customers, investors.

**[00:28:14.350] - Yevgeny Pats**

We understood that actually our idea it's not that good so we pivoted pretty quickly to email security, like enterprise email security and started building that for two years.

**[00:28:25.780] - Yevgeny Pats**

Then I realized, all right, no one is coming. No one is buying us. It doesn't work that way. Actually, you have to build a product that people use. Actually, it was like what I thought before I was in the [inaudible 00:28:39]. Okay, now I get it. But anyway, we had not that great product market feed.

**[00:28:49.100] - Yevgeny Pats**

Anyway, investors brought news to you. I stated it like to help with the tech, things like company, a solid place. Then I went, I said, okay, I want to focus on that first PLG. I think this is the future. We got it completely wrong here.

**[00:29:05.490] - Yevgeny Pats**

But it looks like they found like a new CEO with sales, like growth and he [inaudible 00:29:09] there. The company is still live today, but it's always early, you never know. I think it's doing a bit of revenue, which is nice. Not like a complete disaster.

**[00:29:20.940] - Yevgeny Pats**

But then I was totally focused on PLG, and for about three, four years, I was focused on bootstrap self-funded startups. My last one was really the CI for fast testing. If you know Google has cluster fuzz. It's their open source fuzzing as a service for a lot of open source projects.

**[00:29:45.460] - Yevgeny Pats**

I thought, okay, looks like the market is hot, I want to try it out. I bootstrapped it with something more user friendly. They have it like, it was a Google internal project so I wanted to build something more user friendly.

**[00:29:58.590] - Yevgeny Pats**

I built it and we had quite a good market sharer in terms of open source project using and it was called Fuzz it. Like 50% was Fuzzing and 50% was Google. 50% of the project that we're using fuzzing. Not all of them.

**[00:30:17.080] - Yevgeny Pats**

But then I realized that after the first experience that we raise money way too early, I was always very cautious to raise money because after that it's the way of no return. Okay, what happens if you realize like one week after we raise money that there is no product market fit at all? Then you're stuck.

**[00:30:39.850] - Yevgeny Pats**

Then you have to, okay, let's make something. God, I don't know. You are in to this inconvenience point. I was very cautious with raising money before I'm sure. I realized there is not a lot of product. It's a bit of too niche, like not every project needs huzzing.

**[00:30:59.630] - Yevgeny Pats**

It's a bit of like C, C++ specific if we use the market even more. A lot of our projects were really like C, C++. It's a burning problem with memory corruption, vulnerabilities. I think people understand that they don't write on C++ much anymore, but it means the margin is getting smaller. Actually, luckily GitHub said, "Okay, we want this, it's good for us."

**[00:31:23.450] - Yevgeny Pats**

I said okay, "Just in time, I'll help you integrate it into the platform." Yeah, it was a great experience also in GitHub, like looking how they work in remote company and we are a remote company now. I did a lot of the lessons there. After that, went in to CloudQuery.

**[00:31:42.170] - Alex DeBrie**

Did you stick around in GitHub for a while or did you mostly just help integrate and then say, "Hey, I want to start my own thing again?"

**[00:31:47.930] - Yevgeny Pats**

Yeah, I hope to integrate but I was always thinking about the next thing. I didn't know when I'll find the next one because I wasn't pressured and I want to be sure also, I work on the right thing. It grew faster than I expected. Some of my early experiment with the open source. So once I saw it started getting traction, I decided to live and start focusing on that full time.

**[00:36:40.990] - Alex DeBrie**

Cool. That's a good segue because I want to talk about CloudQuery is a totally open source project but you've also taken seed funding. How do you plan to balance that open versus paid nature?

**[00:36:52.180] - Alex DeBrie**

Do you have any, I guess, business role models in that category, like certain companies that you want to model after there? What are you thinking about?

**[00:36:58.670] - Yevgeny Pats**

Yeah. I think I have a few and I think also it's something it's still an unsolved problem and a lot of companies taking different approaches depending on their specific use case. I think in some sense Terraform is a good role model.

**[00:37:18.700] - Yevgeny Pats**

But again I think we like, yeah... Eventually it will be managed version. A lot of users ask us for manage version because infrastructure burden is real and a lot of companies don't have any specific constraint. It's usually a no-brainer as long as you give some competitive and reasonable pricing.

**[00:37:38.520] - Yevgeny Pats**

This will be our way into monetization to start working on a managed version of CloudQuery. But before doing that we need to have big enough adoption like let's say in the thousands of medium to big organization using us on a daily basis because let's say we release a managed version and we can take something like 10% conversion rate.

**[00:38:02.690] - Yevgeny Pats**

Like we don't have any special features, just the same thing, just manage. Like 10% conversion is I think something reasonable. If you have thousands, like 2,000 users, it's starting to be interesting in terms of revenue.

**[00:38:18.770] - Yevgeny Pats**

But if you have like even 100, which is not bad. If you have 10% conversion rate, then you wasted a lot of time and money on 10 very expensive users while you could focus on growing the community, maturing the project more. You can do both, usually because focus is constrained and also because money is well constrained.

**[00:38:41.680] - Alex DeBrie**

Yeah. There are a lot of open source companies trying to make some money. I think some of them are going to have a tougher time, some of them less. I think this one, like you're saying, people will pay for a manage version, given that there's a database involved, you want some sort of UI, maybe some access control, things like that. There's going to be a lot of opportunity there to make that work.

**[00:39:00.220] - Yevgeny Pats**

Yeah. I think there will be also potentially, and this is something that it's just too early to know, just because the things that I'll know in a year are things that I don't know now regarding what features we would do maybe in the managed version or what features will make sense to build like an extra or like how to do tiers.

**[00:39:20.720] - Yevgeny Pats**

You have a first tier and then you have a second tier. I literally can't. It will be too hard to predict it. But yeah, eventually you will have to. Even if you have this conversion for just a managed version, you'll have to start thinking, okay, how do you increase this conversion? Maybe how do you have tiers, like what features you introduce to make it profitable?

**[00:39:43.100] - Yevgeny Pats**

Again, it will depend on financial. What is the team size that you need to maintain? You have to have all those metrics to make some smart decisions.

**[00:39:53.150] - Alex DeBrie**

Yes, absolutely. It would be cool to see that go. I have a last section here that I want to talk about and I came up with some potential features or areas that CloudQuery can go in. I want you to tell me, "No way, we're not going to do that. That doesn't make sense with CloudQuery," or like this is why that might work, things like that. We'll just riff on these and you can expand on them if you want to.

**[00:40:11.760] - Alex DeBrie**

Number 1, let's start off. Currently, CloudQuery supports SQL. You can do select statements, select whatever from my EC2 instances, whatever. What about inserts or updates? Will you ever create or alter resources in CloudQuery? Or is it always going to be read-only?

**[00:40:26.750] - Yevgeny Pats**

Yeah, it's a great question. I think we discussed this also internally quite a bit. For now, for sure, we will focus on read-only because we still have a lot of features and a lot of growth to do with just what we have now and a lot of incremental features. This one is also if you want to you really like taking the company in another direction. It will also mean a lot of more testing, more development time. You want to be sure people really need it in a sense.

**[00:41:00.630] - Yevgeny Pats**

That's a good question. I'm partially sure. I found one use case which is good but I'm not sure exactly how popular it is. For example, in terms of provisioning, people already use Terraform. I would say it feels like you don't want to replace that. But on things like deleting resources in a sense.

**[00:41:24.980] - Yevgeny Pats**

For example, there is a tool even quite popular called AWS Nuke which clears AWS environments. You actually give the reconfiguration in a DSL. If you look at the code, it's also a bit similar to CloudQuery.

**[00:41:39.790] - Yevgeny Pats**

The next thing is okay, I can use SQL and the framework is the same and then we can just implement the delete thing. But I think if we are going to make this leap, we'll need some more insights into really how popular that is because it can bring a lot of support, new bugs, support burden, development burden. Yeah, it's a good question.

**[00:42:03.110] - Yevgeny Pats**

There is actually another company I think called ISQL, which is I think trying to replace Terraform with exactly [crosstalk 00:42:13] insert into. I think it's how it's going. But yeah, it's an interesting approach but I'm not sure. I don't know. It's a good question.

**[00:42:24.020] - Alex DeBrie**

Yeah, I think I agree with you. I'm a big infrastructure as code fan and I just think why would you want to have it drift from your Terraform? You've got the CloudQuery drift in there and this is going to add more drift. So I agree. But it's interesting.

**[00:42:35.040] - Alex DeBrie**

The one thing I was wondering is policy remediation stuff. If I always want to say on my S3 buckets, always enforce this, it'd be nice as a security team to be able to do that, but you probably just got to go out and talk to the teams responsible and say, "Hey, change your Terraform so that we're always provisioning this property or whatever." Probably better ways to do that.

**[00:42:53.800] - Yevgeny Pats**

Yeah, exactly. Because then you also can get in through these loops if you just remediate it. This is what we have from the infrastructure team, they are worried about sometimes about auto-remediation because then it will just go into the loop.

**[00:43:09.330] - Yevgeny Pats**

The only exclusion that I heard about is if it's this one or two rules that are really bad that, you want to remediate right away. But usually, if it's something that you just don't want, you want to find exactly the infrastructure code and fix it there. It's redeployed.

**[00:43:27.250] - Alex DeBrie**

Yes. Cool. Okay. Second idea, have you thought about merging this with metrics in some way? I'm seeing if I want to say I have all these S3 buckets or maybe all these SQS queues or whatever across my infrastructure find and show me the top 10 SQS queues by the number of messages received in the last day or something like that.

**[00:43:47.930] - Alex DeBrie**

You probably have to push this out to a foreign data wrapper or something and query that outside of CloudQuery. But would you ever mix that in?

**[00:43:54.380] - Yevgeny Pats**

Yeah, I think it's something that we will look into and I think basically the meta feature is connected with other external data. If it's metric, but also maybe cost, maybe the next one most requested one, which is also can work also in ETL approach like metrics is security issues. Collect security issues from AWS security hub or things like whatever, get a dependent bot or any other alerts or sneak in, plug it in and then you can do smart things and prioritize.

**[00:44:34.100] - Yevgeny Pats**

Usually, you have tons of those alerts and then you can ask questions. Show me the vulnerabilities that are just on in this VPC, which is important for me. The metrics is another place to get information. Probably it's a more tricky one because it's not a classic ETL. Because as you said, maybe foreign data wrapper because there is a lot of data there.

**[00:45:01.590] - Yevgeny Pats**

Yeah, it's a good question if you want to store them daily or just query directly and have a positive... Data wrapper is actually good but actually, a great way to solve it.

**[00:45:12.170] - Alex DeBrie**

Yeah, cool. What about the third idea? Currently, CloudQuery is a pull-based system. What about a more event-driven push-based system where if I create an EC2 resource, maybe you guys hook into my Cloud Trail and immediately ingest that in? I imagine that takes a lot from the cloud providers to help you out there. But have you looked into more vent-driven and push-based updates?

**[00:45:34.100] - Yevgeny Pats**

Yeah. It's also on the need long-term to investigate and try out. But yeah, reading CloudTrail or EventBridge, for example, in AWS or in Google it's... I don't remember the name, but yeah, the alternative CloudTrail, getting into that and then basically updating on a push basis. It's possible, but of course, there will be new challenges, new development, new things that you want to test.

**[00:46:05.400] - Yevgeny Pats**

Just for example, CloudTrail also has 30, sometimes 60 minutes delay. Sometimes it only has partial information. Sometimes it will be easy just to update in place. The state of this EC2 was changed from running to stopped. You just go ahead and change it.

**[00:46:26.460] - Yevgeny Pats**

But sometimes you get only part of the information and then you might need to actually do another call to the cloud providers, "Okay, give me all the information now I want to update." A lot of room for questions and how to do it in scalable ways and so on.

**[00:46:44.670] - Alex DeBrie**

Yes, sure thing. Cool. All right, last one. Maybe it already supports the setting poke around the hub too much. But what about non-engineering resources? Maybe if I want to go to my marketing email provider and look up pulling email subscribers or maybe rippling and pulling employees or gusto or whatever. Do you support things outside of cloud providers or is it pretty focused on engineering stuff?

**[00:47:09.180] - Yevgeny Pats**

Right now we don't have those providers just because we were super, you can do AWS and GCP and working with those use cases. But it was always on our mind to also potentially support others. The use case there is usually different. It's more data team, marketing team and it's a question because actually the underlying tech and architecture is exactly that, it's just that our first use case was found infrastructure.

**[00:47:38.280] - Yevgeny Pats**

Then we'll be looking into companies like Airbyte, for example, or Fivetran that are focused on the long tail of small marketing providers. But maybe something that we'll also look into potentially.

**[00:47:53.590] - Alex DeBrie**

Cool. Well, thanks for playing along with me there on some of those even if those weren't the best feature ideas for you, I appreciate walking through that.

**[00:48:00.330] - Yevgeny Pats**

It's going well, actually. Yeah.

**[00:48:02.470] - Alex DeBrie**

Cool. Yevgeny, thanks for coming on Software Engineering Day has been a great discussion. Where can people go to find out more about CloudQuery about you? Where should they be looking?

**[00:48:11.060] - Yevgeny Pats**

Yeah. Feel free to like to jump on our GitHub issues. We also have a Discord which is fairly active. You can go to callgridil/discord or on our website on the discord link and you can reach out to me there on my Twitter or on my LinkedIn. You can find me anywhere basically, monitoring all questions as I'm support team as well.

**[00:48:34.830] - Alex DeBrie**

Perfect. Sounds great. Well, Yevgeny Pats, creator and founder of CloudQuery. Thanks for coming on Software Engineering Daily.

**[00:48:41.180] - Yevgeny Pats**

Awesome. Thanks, Alex. I really appreciate hosting me and I really enjoyed it.

**[00:48:44.960] - Alex DeBrie**

Sure thing.

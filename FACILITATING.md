## Facilitating Onboarding Week (OSS Track)

### Overview
Onboarding Week exists to provide Cloud Foundry engineers dedicated time to explore the platform in a self-paced learning environment.

It has also been used by Labs engineers on the beach to gain a deeper understanding of the platform they use for client projects.

Either way, the idea is to give people who want more experience with CF the space to work through common problems alongside someone who has the same questions as they do, coming away with knowledge drawn from figuring something out themselves (something that is often lacking when paired with an experienced dev).

Whether you are with CF or with Labs, Onboarding Week is intended to be a **facilitated** experience. We want participants to get "productively lost" together, not wander in aimless hair-tearing frustration. If you want to give it a try, but don't have a facilitator, join the [#cf-onboarding-week](https://pivotal.slack.com/messages/cf-onboarding-week/) Slack channel to ask if someone in another office is available to be a remote guide.

### Preparation Suggestions
* Allocate participants for the week.
* Designate workstation(s).
* Setup remote pairing hardware for remote participants.
* Provide one "floater" laptop per pair to facilitate independent docs reading.
* Create and populate *private* Tracker Project(s). Invite participants.
* [Request GCP IAAS project(s)](https://docs.google.com/forms/d/e/1FAIpQLSeJ31997Zma1WtLcCtswiysCFWOG5MXNmlYCpJsiYgdG9kKnA/viewform). Once they're created, add participants as Owners.
* A month before, send out an email to all participants, reminding them to clear their schedules for the week.  Schedule day-long meetings to help avoid people scheduling more meetings for that time.

Subject: CF Onboarding (<DATE>)
```
Howdy folks!
If you're getting this email, it means that you're scheduled to participate in CF Onboarding soon (the week of __________________).

During the week you will be pairing full time with each other, working on stories exploring actual user journeys and workflows that are central to the CF ecosystem.

There will be more communication closer to the start of CF Onboarding, so don't worry too much about the details of preparation just yet.  You will be pairing full time, however, so your teams should expect you to be away for the majority of the week.  We've put a meeting on your calendar to block out this time.  If you generally have a lot of meetings we recommend cancelling or moving as many of them as possible, but if you need to step out for a meeting here or there, that shouldn't be a problem.

If this week doesn't work for you, let us know!  There will be more CF Onboarding weeks and we can move people.

Thanks!
The CF Onboarding Staff
```
* The week before, send out an email to all participants, reminding them that they are signed up for onboarding week and setting expectations about how the week will go. Here is a template that can be used as a starting point:

```
Howdy folks!

If you're getting this email, it means that you're scheduled to participate in CF Onboarding next week (the week of __________________). I wanted to get in touch with everyone and give you an idea about what to expect next week.

For starters, we (_____________ and _____________) will be your facilitators. We'll continue working our "day jobs," but check in a few times a day to answer questions and make sure everyone is doing well. We will also be here for morning stand ups.

You will be part of a group with different backgrounds and skills. The stories can get a bit technical, but by and large you'll be going through actual user journeys -- deploying CF with BOSH, pushing an app, debugging broken systems. Please help each other out as much as possible, write down what you learn, and think about how difficulties you have with the stories could be feedback for us as facilitators or for teams that build the products you're using.

I also wanted to point out that you will be pairing full time with each other. Your teams should expect you to be away for the majority of the week, but if you need to jump out for a meeting here or there, just let your pair know.

Action Items for you:
1. If you have any issues with the scheduling, you can always sign up for a different week, but please try to find someone else to take your spot.
2. Let your anchor and PM know that you'll be away next week for CF Onboarding. Feel free to loop us in if there are any concerns about that.
3. On Monday morning, please meet downstairs outside of Pegasus and Perseus. We'll have pairing stations ready to go for you. We'll start with an Onboarding orientation, we'll set you up with pre-populated Pivotal Tracker projects with stories for you to work on, and we'll send you on your way.

Otherwise, if you have any questions, feel free to reply to this email.

Thanks!
The CF Onboarding Staff
```

### Communication Suggestions
* **Morning Standup:** to review how the week is going / answer questions.
* **After-lunch pow-wow:** more review and questions (for some groups this is indispensable, for others it's less necessary).
* **Second Afternoon Check-in:** Have a facilitator stop by the participants' workstations again later in the afternoon (around 4pm) to see if they have any additional questions.
* **+/∆s ("plus deltas"):** An end-of-day email to everyone involved in the week (e.g. onboarders, managers, interested third parties) reviewing what went well and what they would have liked to change about the day. This is an opportunity to reflect and give prompt feedback on how the week is going.
* **Group Slack:** Use the [#cf-onboarding-week](https://pivotal.slack.com/messages/C3HN8FALB) Slack channel to share resources and post reminders.
* **Retro:** What Pivotal experience could be complete without a retro? Order goodies and treat it like a celebration.

### Monday Briefing Suggestions
* Make sure to introduce everyone to each other in the room and break the ice.  You are each others' greatest resource.
* Make sure to leave space for questions and discussions at each part of the briefing.
* Explain the structure of the week.  The participants will be pairing, working out of a backlog of prepared stories, at a pace they choose.
* Set expectations around the goals of onboarding week.  The objective is to pursue whichever activities will help expand your understanding of how Cloud Foundry works.  Rabbit-holing is allowed.  Gaining deeper understanding is more important than finishing stories.
* Set expectations that no one will finish the backlog, or even come close, and this is expected and OK.
* Explain how there will be opportunities for exploration later in the week.  There a few "core" tracks (`bbl`, `cf-deployment`, `app-dev`).  After that point, it is recommended that participants look at other possible tracks and mindfully choose which they want to work on, since they won't have time to do everything.  Participants can also go rogue and make up their own track if they are interested in a certain question.
* Explain how there will be talks, and solicit the participants for their interests so you can schedule talks for later in the week.

### Structure Suggestions
* If you have more than one pair doing this at once (which is a better experience for all involved if you can swing it), try having them rotate after lunch instead of in the morning. This helps with context retention, knowledge sharing, and getting the day started promptly (instead of settling in with a new pair).
* Don't load all modules at once at the beginning of the week. Start with stories to get their environment set up, then add others as-needed in response to Onboarder questions and requests.
* If they find the last hour or so of the day hard to get though (this is a really draining week for Onboarders) encourage them to use it in whichever way will be most productive. This could mean separating from their pair to read docs, if that's what they need.
* Thursday and Friday are good days for the more free-form exploratory modules that they might not have been confident enough to tackle at the beginning of the week.

### Daily Theme Suggestions
* Be conscious of your pair's needs. Support one another. Teach one another.
* When you learn something new, tell your pair. Get in the habit of being honest about ignorance and progress.
* Learn to be comfortable with the idea of breaking something. If you're afraid to break your environment, you'll have a harder time learning and exploring.
* This is your week. Pursue the topics that interest you.
* How will you apply what you've learned this week to the future?
* Use Todd Sedano's [pairing cards](http://sedano.org/toddsedano/2017/10/23/considerate-pair-programming.html) to prompt a different pairing skill each day.

# rock

The idea behind rock is to make a time tracking tool with only the functionality I've felt has been useful to me.

## Keep the functionality as simple as possible

To me, a timetracking tool needs to do only the following things:
- Start tracking of a task
- Stop tracking of a task
- Provide an overview of how long I have worked on each task

The format in which a timetracking tool logs its data should be human readable. The data it writes should also be easily modifieable by a human; you should be able to correct the time in which you've clocked in or out in case you forgot or made a mistake.

## Provide multiple interfaces

A time tracking tool should have multiple ways to interact with it, depending on the user's needs.

I have currently identified the following modes of interaction which could be useful:
- CLI
- TUI
- HTTP Server

I'm certain there are others out there which I have not considered.

## Other considerations

While not strictly one of the features I expect of a timetracking tool, it would be particularly useful if the tool could make the effort of logging the data in JIRA or similar tools. Nobody enjoys having to open JIRA at the end of the day.

As it stands, I don't consider this a planned feature for `rock`, but I am considering writing a separate tool for this, which can read the `rock` time tracking format.

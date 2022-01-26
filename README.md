# rackmon

This is a tool to output metrics for various sensors that could be
used to monitor a SOHO rack.  It's intended to be used via Telegraf's
`execd` plugin.

It's basically a catchall for what I want to use to monitor the
environment that is not baked-in to Telegraf.

It may not may not be useful to you.  It is only tested on Raspberry
Pis running Raspbian OS.  It should work anywhere you can hook up the
necessary hardware, however.  I try to avoid Raspberry Pi-specific code
and if I cannot it will be noted.  Each component is discrete, so only
the sensor you are using needs to be present and functional.

This is a work in progress.

## Metrics

Currently this supports:

- d1w: Dallas 1-Wire temperature sensor via Linux kernel support.

Planned devices:

- TrippLite UPS via nut - An SMBUS-based temperature/humidity sensor

## LICENSE

This is licensed under the standard MIT License.  That's not negotiable.
Inclusions of libraries that make that impossible are a non-starter.
The terms of that license are required and are the limits of the
requirements.

## Contributions

Suggestions on approaches to testing without the presence of the device(s)
would be appreciated.  As-is, testing is by hand, and not automated and
that's gross. But my creativity there is lacking.

PRs are welcome.  Issues are not very helpful at this stage.  Feature
requests will mostly be closed.  Bug reports with out a PR will be noted,
but likely not addressed promptly.

Adding support for other devices is appreciated, but I personally will
only be targeting what I use because that's all I can test. Submissions
of other device support will only be checked for style at this time.

If you want *me* to add a device that I don't have, you'll need to send
me hardware and be patient :).

I have a job, family, and other toy projects.  Please do not expect me
to be super-responsive.  If you love this and want to own it, you're
welcome to fork and maintain more formally.  I'll shut this down and
refer to yours.

Please be polite and I will be polite as well.  Confrontational or
aggressive language won't be tolerated.  I disdain that behavior so
frequent in the OSS community.

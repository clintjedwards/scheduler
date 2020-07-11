- telemetry statistics on which employees shift timeframe positions
- positions should display stastics on employees who work it the most and least maybe top 3 of each?
- support retrieving schedules for certain timeframes in order of earliest date
- positions should have competancy numbers associated
- schedules should exist in order of date and be able to be retrieve by time span
- for each employee have a dict(or some way to tell which days in the past they were scheduled) which has every day they were scheduled. (Do we need to move to a RMDB?)
- need to create a weighting system so that certain traits about an employee influence their changes to be picked
- store generated dates ranges so that we can retrieve the schedules for them, in order
- employees should have preferred roles
- team leader is just another position

- What should we do about telemetry? It's possible we can write a background process to populate already existing entities with things we care about (dates worked, times worked certain position)
- needs schedule wide rules like employees shouldn't have back to back shifts and etc, this can be presented as generation options

- Eventually we should also make it so that it's easy to place new employees in shifts via the return schedule
  so commit schedule could actully just take an updated schedule

// We should be able to represent unavailability by using extended cron
extended cron is regular cron plus another field for year
http://www.nncron.ru/help/EN/working/cron-format.htm

The user will enter a cron format for all times in availabiliy, it is possible to enter multiple cron formats.
This gives us a large range of times, dates, etc that can be represented. For resourcing limits we cannot
check if all times are within a certain cron timeframe but we can check if a certain fraction of those times are.
To this end we can limit the checks on if an employee can work by taking into account shift start and end time,
the current date and then plugging each of these into the cron checker by 30 min intervals. This should allow
us the best of both worlds for a minimal cost in checking.

The expression parser would work like most cron expression parsers. Break down a time in to indivdual parts
and return true only if all those parts return within certain cron expressions

The library should take a start time and end time and iterate through it by some given time duration, and check against the expression
for each until end

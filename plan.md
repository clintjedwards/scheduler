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

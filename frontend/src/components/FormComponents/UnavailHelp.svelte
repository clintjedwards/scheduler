<div class="text-body">
  <p>
    To set the time periods in which an employee is not available, we write the
    time period in a special format. This format defines the time in which the
    employee should not work for each individual unit of time. We start from the
    smallest unit of time and continue to the largest. You can view the units of
    time in the example table to the left.
  </p>
  <p>
    You are allowed to use commas(representing a list), dashes(representing a
    range), a star(meaning any value), or just a single number to convey what
    times an employee is not available.
  </p>
  <p>
    As an example, let's go through how you would write an employee is
    unavailable from
    <code class="text-sm">12/25/2020 to 12/30/2020</code>.
  </p>
  <ol>
    <li>
      The first time period we define is minute. Since for this example we don't
      want the employee to be unavailable for entire days, we can put a
      <code class="text-sm">*</code>. This signifies that, for this time period,
      all minutes are unschedulable.
    </li>
    <li>
      The second is the hour time period. It follows that we also don't want to
      specify hours. So we put another
      <code class="text-sm">*</code>.
    </li>
    <li>
      The third is day of month or date. For this we could choose to use a list
      like so:
      <code class="text-sm"> 25,26,27,28,29,30</code>. But thats a lot of
      writing. Instead, we use a range: <code class="text-sm">25-30</code>.
    </li>
    <li>
      The third is Day of month. For this we could use a list like so: <code
        class="text-sm"
      >
        25,26,27,28,29,30
      </code>. But thats a lot of writing. Instead we use a range:
      <code class="text-sm">25-30</code>.
    </li>
    <li>
      The fourth is Days of the week. For this we don't mind which days the
      25-30 fall on. So as usual we leave a
      <code class="text-sm">*</code>.
    </li>
    <li>
      Lastly, we need to define the year or else this employee gets Christmas
      off every year. So we just use a single value: <code class="text-sm"
        >2020</code
      >.
    </li>
  </ol>
  <p class="mb-4">
    Altogether the full expression becomes:
    <code class="text-sm">* * 25-30 12 * 2020</code>
  </p>
  <pre
    class="text-gray-600 text-sm">
More Examples:

Unavailable Time                Input
-----------------               -----
12/25/2020 to 12/30/2020        * * 25-30 12 * 2020
Never available                 * * * * * *
Every year on the first of Jan  * * 1 1 * *
All of Jun, July, Aug for 2020  * * * 6,7,8 * 2020
Never schedule for 9am          * 9 * * * *
    </pre>
</div>

/*
Package avail helps with the determination of whether a given time exists within a certain cron expression.

The package uses a subset of the extended cron standard:
https://en.wikipedia.org/wiki/Cron#CRON_expression

Why

It is sometimes useful to represent certain timeframes using cron expressions. For example, the
unavailability of an employee might be a good case. An employee's unavailability must be able to
convey not only specific dates that the employee might not be able to work (Micheal cannot work
on his birthday this year), but also convey repeated dates (Micheal does not work Christmas every year).

Using cron to achieve this allows the representation of situations like Micheal's to be compact and easy
to parse.

This package allows a user to represent Micheal's unavailablity in a cron expression and then efficently
check whether Micheal can or cannot work via a given time.

How

Avail implements/uses a stripped down version of the cron expression syntax as defined below.

    Field           Allowed values  Allowed special characters

    Minutes         0-59            * , -
    Hours           0-23            * , -
    Day of month    1-31            * , -
    Month           1-12            * , -
    Day of week     0-6             * , -
    Year            1970-2100       * , -

Avail accepts a cron expression in the format above, splits it into separate fields, parses it,
and generates map backed sets for each field in order to allow speedy checking of value existance.

This data structure allows avail to take a user supplied time and check that each of the time's
elements exist in the representation of the cron expression.

Usage

Initiate a new avail struct with cron expression and call "able" with a specified time

    avail, _ := New("* * * * * *")

    now := time.Now()

    fmt.Println(avail.Able(now))
    // Output: true

*/
package avail

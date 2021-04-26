const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Check if the market is open now.
financefeast.getClock().then((clock) => {
    console.log('The market is ' + clock.is_open ? 'open.' : 'closed.')
})

// Check when the market was open on Dec. 1, 2018.
const date = '2018-12-01'
financefeast.getCalendar({
    start: date,
    end: date
}).then((calendars) => {
    console.log(`The market opened at ${calendars[0].open} and closed at ${calendars[0].close} on ${date}.`)
})
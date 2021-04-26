const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Get our position in AAPL.
aaplPosition = financefeast.getPosition('AAPL')

// Get a list of all of our positions.
financefeast.getPositions()
    .then((portfolio) => {
        // Print the quantity of shares for each position.
        portfolio.forEach(function (position) {
            console.log(`${position.qty} shares of ${position.symbol}`)
        })
    })


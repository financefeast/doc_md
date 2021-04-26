const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Get our account information.
financefeast.getAccount()
    .then((account) => {
        // Check if our account is restricted from trading.
        if (account.trading_blocked) {
            console.log('Account is currently restricted from trading.')
        }

        // Check how much money we can use to open new positions.
        console.log(`$${account.buying_power} is available as buying power.`)
    })
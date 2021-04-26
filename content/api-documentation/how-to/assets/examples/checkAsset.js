const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Check if AAPL is tradable on the financefeast platform.
financefeast.getAsset('AAPL')
    .then((aaplAsset) => {
        if (aaplAsset.tradable) {
            console.log('We can trade AAPL.')
        }
    })
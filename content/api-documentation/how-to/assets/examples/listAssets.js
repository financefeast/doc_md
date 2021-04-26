const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Get a list of all active assets.
const activeAssets = financefeast.getAssets({
    status: 'active'
}).then((activeAssets) => {
    // Filter the assets down to just those on NASDAQ.
    const nasdaqAssets = activeAssets.filter(asset => asset.exchange == 'NASDAQ')
    console.log(nasdaqAssets)
})

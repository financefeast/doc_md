const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Get the last 100 of our closed orders
const closedOrders = financefeast.getOrders({
    status: 'closed',
    limit: 100,
    nested: true  // show nested multi-leg orders
}).then((closedOrders) => {
    // Get only the closed orders for a particular stock
    const closedAaplOrders = closedOrders.filter(order => order.symbol == 'AAPL')
    console.log(closedAaplOrders)
})


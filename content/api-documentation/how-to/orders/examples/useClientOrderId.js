const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Submit a market order and assign it a Client Order ID.
financefeast.createOrder({
    symbol: 'AAPL',
    qty: 1,
    side: 'buy',
    type: 'market',
    time_in_force: 'day',
    client_order_id='my_first_order'
})

// Get our order using its Client Order ID.
financefeast.getOrderByClientOrderId('my_first_order')
    .then((myOrder) => {
        console.log(`Got order #${myOrder.id}.`)
    })
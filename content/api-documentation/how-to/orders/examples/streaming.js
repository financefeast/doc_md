const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Prepare the websocket connection and subscribe to trade_updates.
financefeast.websocket.onConnect(function () {
    console.log("Connected")
    client.subscribe(['trade_updates'])
    setTimeout(() => {
        client.disconnect()
    }, 30 * 1000)
})
financefeast.websocket.onDisconnect(() => {
    console.log("Disconnected")
})
financefeast.websocket.onStateChange(newState => {
    console.log(`State changed to ${newState}`)
})

// Handle updates on an order you've given a Client Order ID.
const client_order_id = 'my_client_order_id'
financefeast.websocket.onOrderUpdate(orderData => {
    const event = orderData['event']
    if (orderData['order']['client_order_id'] == client_order_id) {
        console.log(`Update for ${client_order_id}. Event: ${event}.`)
    }
})

// Start listening for updates.
financefeast.websocket.connect()
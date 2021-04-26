const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Prepare the websocket connection and subscribe to account_updates.
financefeast.websocket.onConnect(function () {
    console.log("Connected")
    client.subscribe(['account_updates'])
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

// Handle incoming account updates.
financefeast.websocket.onAccountUpdate(accountDdata => {
    // Track the cash balance in our account.
    const balance = accountDdata['cash']
    console.log(`Update for account. Cash balance: ${balance}`)
})

// Start listening for updates.
financefeast.websocket.connect()
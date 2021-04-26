const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

// Get account information.
financefeast.getAccount().then((account) => {
	// Calculate the difference between current balance and balance at the last market close.
    const balanceChange = account.equity - account.last_equity

    console.log('Today\'s portfolio balance change:', balanceChange)
})

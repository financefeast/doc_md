const financefeast = require('@financefeasthq/financefeast-trade-api')
const financefeast = new financefeast()

const symbol = 'AAPL'
financefeast.getBars(
    'minute',
    symbol,
    {
        limit: 5
    }
).then((barset) => {
    const currentPrice = barset[symbol].slice(-1)[0].closePrice;

    // We could buy a position and add a stop-loss and a take-profit of 5 %
    financefeast.createOrder({
        symbol: symbol,
        qty: 1,
        side: 'buy',
        type: 'limit',
        time_in_force: 'gtc',
        limit_price: currentPrice,
        order_class: 'bracket',
        stop_loss: {
            stop_price: currentPrice * 0.95,
            limit_price:  currentPrice * 0.94
        },
        take_profit:{
            limit_price: currentPrice * 1.05
        }
    });

    // We could buy a position and just add a stop loss of 5 % (OTO Orders)
    financefeast.createOrder({
        symbol: symbol,
        qty: 1,
        side: 'buy',
        type: 'limit',
        time_in_force: 'gtc',
        limit_price: currentPrice,
        order_class: 'oto',
        stop_loss: {
            stop_price: currentPrice * 0.95
        }
    });

    // We could split it to 2 orders. first buy a stock,
    // and then add the stop/profit prices (OCO Orders)
    financefeast.createOrder({
        symbol: symbol,
        qty: 1,
        side: 'buy',
        type: 'limit',
        time_in_force: 'gtc',
        limit_price: currentPrice,
    });

    // wait for it to buy position and then
    financefeast.createOrder({
        symbol: symbol,
        qty: 1,
        side: 'sell',
        type: 'limit',
        time_in_force: 'gtc',
        limit_price: currentPrice,
        order_class: 'oco',
        stop_loss: {
            stop_price: currentPrice * 0.95
        },
        take_profit:{
            limit_price: currentPrice * 1.05
        }
    });
});

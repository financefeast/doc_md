---
title: How-To Code Examples
weight: 20
---

# How-To

Welcome to the how-to section of our documentation! Here, you'll find some examples of how to do specific tasks using
the financefeast API in supported programming languages.

Before running these examples, it is recommended that you set the following environment variables to the corresponding 
values found on your customer api dashboard. You can get your client credentials [here]('https://customer.financefeast.io')
```
FF-CLIENT-ID
FF-CLIENT-SECRET
```
On Linux, this can be done with `export ENV_VAR_NAME=DASHBOARD_VALUE`. (You'll probably want to add them to 
your `~/.bash_profile` so they stay set when the machine restarts.) On Windows, you can instead 
use `setx ENV_VAR_NAME "DASHBOARD_VALUE"`. After running these commands, you'll want to open a new command 
prompt window so the variables are available.
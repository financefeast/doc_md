---
title: Account Configurations
weight: 70
---

The account configuration API provides custom configurations about your
trading account settings. These configurations control various allow you to modify settings to suit your trading needs.

For DTMC protection, see [Day Trade Margin Call Protection]({{< ref path="/trading-on-financefeast/user-protections#day-trade-margin-call-dtmc-protection-at-financefeast" >}})

{{< rest-endpoint resource="account-configurations" method="GET" path="/v2/account/configurations" >}}
{{< rest-endpoint resource="account-configurations" method="PATCH" path="/v2/account/configurations" >}}

## AccountConfigurations Entity

### Example
{{< rest-entity-example name="account-configurations-v2">}}

### Properties
{{< rest-entity-desc name="account-configurations-v2">}}

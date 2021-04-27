---
title: Authorization
weight: 25
---

To authorize to protected endpoints you must supply an 'Authorization' http header that was recently returned from 
an authenticate call.

The http header is in the following format
```
Authorization: Bearer YOUR-AUTHORIZATION-CODE
```


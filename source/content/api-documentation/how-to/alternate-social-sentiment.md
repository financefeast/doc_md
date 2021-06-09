---
title: Social Media Sentiment Examples
weight: 160
---

## Get Social Media Sentiment
By sending a GET request to our [`/alternate/social-sentiment`]({{< relref "/api-documentation/api-v1/alternate/alternate-social-sentiment.md" >}}) endpoint you
will get a list of social media posts from social media platforms for the stock ticker searched for. Sentiment is returned per post. Also returned is an aggregation
of sentiment count per day.

{{< code-example exampleId="alternate-social-sentiment" pathURL="/api-documentation/how-to/alternate-social-sentiment/examples">}}


---
title: OAuth Integration
weight: 10
description: Financefeast supports OAuth for integration to our API
---

## <a name="introduction"></a>Introduction

We welcome developers to build applications and products that are powered by Financefeast while also protecting the privacy and security of our users. To build using Financefeast's APIs, please follow the guide below.

Financefeast uses [`OAuth2`] as the framework to authorize Financefeast customers to our API.

* [Terms of Access and Use]({{< relref "#terms-of-access-and-use" >}})
* [Registering Your Application]({{< relref "#application-creation-and-registration" >}})
* [Integration]({{< relref "#integration" >}})
* [Making API Calls]({{< relref "#api-call" >}})
* [Best Practices]({{< relref "#best-practices" >}})

## Terms of Access and Use

* You must read the terms and create a new client in order to connect and use Financefeast's APIs
* All API clients must authenticate with OAuth 2.0.
* You may not imply that the app was developed by Financefeast.
* OAuth clients need to be approved by Financefeast. Please contact partnership@Financefeast.io


## Application Creation and Registration

* Before integrating with Financefeast, you'll first need to create a new OAuth client under your [OAuth Apps](https://identity.financefeast.io/oauth/create_client) page.

<center><img src="./add_client.png" width="80%"></center>

* Once you add your relevant information and create the client, you will see your new client in your [Integrated Applications](https://identity.financefeast.io/#clients) page.

<center><img src="./new_client.png" width="60%"></center>

## Integration

Once the client is approved we will enable it, and then you can start integration testing. You should use OpenID Connect or OIDC to verify customers identity.

To integrate your application with Financefeast, use [`authorization_code`] flow. More details can be found in the [RFC6749](https://datatracker.ietf.org/doc/html/rfc6749)

<center><img src="./authorization_code_flow.png" width="80%"></center>

If you are developing your application in python we recommend using [Authlib](https://docs.authlib.org/en/latest/client/frameworks.html) to handle
the OAuth flows.

### 1. ***End user requests service from application. Application redirect users to request Financefeast Identity***

```
GET https://identity.financefeast.io/oauth/authorize?response_type=code&client_id=SOME_CLIENT_ID&redirect_uri=YOUR_REDIRECT_URI&scope=openid+profile+plan&state=Pxx5fVv8sMue8eymdg1hhCtkvERe0Z&nonce=0jUUjAJskDY4ivq5CgvY
```

When redirecting a user to Financefeast to authorize access to your application, you’ll need to construct the authorization URL with the correct parameters and scopes. Here’s a list of parameters you should always specify:

| Parameter | Description |
| --------- | ----------- |
| `response_type` | **Required** Must be `code` to request an authorization code. |
| `client_id` | **Required** The Client ID you received when you registered the application. |
| `redirect_uri` | **Required** The URL where the user will be sent after authorization. It must match one of the whitelisted redirect URIs for your application. |
| `state` | **Optional** An unguessable random string, used to protect against request forgery attacks.  |
| `nounce` | **Optional** Binds the token with the client. [More here](https://stackoverflow.com/questions/46844285/difference-between-oauth-2-0-state-and-openid-nonce-parameter-why-state-cou)|
| `scope` | **Optional** A space-delimited list of scopes your application requests access to. Generally you will want openid profile plan. |

**Allowed Scopes**:

| Scope           | Description                                             |
| --------------- | ------------------------------------------------------- |
| `openid` | Indicates your application will use OIDC to verify the customers identity |
| `profile`       | Customers email address                        |
| `plan`          | Customers basic plan data                               |



### 2. ***End user authorizes API access for the applications***

From the user side, they will see the following authorization screen:

<center><img src="./authorise_screen.png" width="40%"></center>

### 3. ***Financefeast redirects end user to application with an authorization code.***

If the customer approves access, Financefeast will redirect them back to your `redirect_uri` with a temporary `code` parameter. If you specified a state parameter in step 1, it will be returned as well. The parameter will always match the value specified in step 1. If the values don’t match, the request should not be trusted.

Example of the redirect:

```
GET https://example.com/oauth/callback?code=67f74f5a-a2cc-4ebd-88b4-22453fe07994&state=8e02c9c6a3484fadaaf841fb1df290e1
```

### 4. ***Application receives the authorization code***

You can use this code to exchange for an access token.

### 5. ***Application exchanges the authorization code with an access token from Financefeast***

After you have received the temporary `code`, you can exchange it for an access token. This can be done by making a POST call:

```
POST https://identity.financefeast.io/oauth/token
```

With following parameters:

| Parameter | Description |
| --------- | ----------- |
| `grant_type` | **Required** Must be set to `authorization_code` for an access token request. |
| `code` | **Required** The authorization `code` received in step 4 |
| `client_id` | **Required** The Client ID you received when you registered the application. |
| `client_secret` | **Required** The Client Secret you received when you registered the application. |
| `redirect_uri` | **Required** The redirect URI you used for the authorization code request. |

*Note: this request should take place behind-the-scenes from your backend server and shouldn't be visible to the end users for security purposes.*

The content type must be `application/x-www-form-urlencoded` as defined in [RFC](https://tools.ietf.org/html/rfc6749#section-3.2).

Example request:

```
curl -X POST https://identity.financefeast.io/oauth/token \
  -d 'grant_type=authorization_code&code=67f74f5a-a2cc-4ebd-88b4-22453fe07994&client_id=fc9c55efa3924f369d6c1148e668bbe8&client_secret=5b8027074d8ab434882c0806833e76508861c366&redirect_uri=https://example.com/oauth/callback'
```

After a successful request, a valid access token will be returned with the response:

```
{
  "access_token": "2hvLQ6TQ8MlFd652r8JT7k78gbVNDsNXi5sRS6vSVi",
  "expires_in": 864000,
  "scope": "openid profile plan",
  "token_type": "Bearer",
  "expires_at": 1621634881
}
```

## API Call

Once you have integrated and have a valid access token you can start make calls to Financefeast  API on behalf of the customer.

Example requests:

```
curl -X 'GET' \
  'https://api.financefeast.io/data/eod?ticker=air.nz' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer 2hvLQ6TQ8MlFd652r8JT7k78gbVNDsNXi5sRS6vSVi'
```

## Best Practices

* Once you receive an access token you can store it but you must check for expiry before use by checking the tokens `expires_at` property. If
it has expired request the customer to re-authorize to receive a new token. Tokens have a 10 day lifetime.
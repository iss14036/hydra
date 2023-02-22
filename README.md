Background

Currently, Algo treats third-party authentication the same as a customer or end-user authentication which is using JWT authentication with the same public and private key. The problems that happened are:

We treat first-party and third-party equal (Not recommended due to different profiles and behavior)

We cannot revoke the Token (Client side authentication)

Scope based on user type role (hardcoded in code when there is a need to change)

Hard to set the expiration time for a specific client

As the problems above, we need to find a system that is able to cover those problems. So that we come up with a solution implementing Oauth2 Authentication that Ory Hydra provide

Pre-Requisite

If you don’t familiar with Oauth, please read the RFC documentation HERE. Or read the illustration implementation HERE.

Goal

Implement Oauth2 Authentication with Ory Hydra (Framework) to provide third-party authentication that is reliable, widely used, and easy to maintain, and implement.

What is Ory Hydra?

Ory Hydra is a server implementation of the OAuth 2.0 authorization framework and the OpenID Connect Core 1.0. Existing OAuth2 implementations usually ship as libraries or SDKs such as node-oauth2-server or Ory Fosite, or as fully featured identity solutions with user management and user interfaces, such as Keycloak.

Ory Hydra is a hardened, OpenID Certified OAuth 2.0 Server and OpenID Connect Provider optimized for low latency, high throughput, and low resource consumption. Ory Hydra is not an identity provider (user sign up, user login, password reset flow), but connects to your existing identity provider through a login and consent app. Implementing the login and consent app in a different language is easy, and exemplary consent apps (Node) and SDKs for common languages are provided.

Implementing and using OAuth2 without understanding the whole specification is challenging and prone to errors, even when SDKs are being used. The primary goal of Ory Hydra is to make OAuth 2.0 and OpenID Connect 1.0 better accessible.

Ory Hydra implements the flows described in OAuth2 and OpenID Connect 1.0 without forcing you to use a "Hydra User Management" or some template engine or a predefined front-end. Instead, it relies on HTTP redirection and cryptographic methods to verify user consent allowing you to use Ory Hydra with any authentication endpoint, be it Ory Kratos, authboss, User Frosting or your proprietary Java authentication.

Security

OAuth2 and OAuth2-related specifications are over 400 written pages. Implementing OAuth2 is easy, getting it right is hard. Ory Hydra is trusted by companies all around the world, has a vibrant community and faces millions of requests in production each day. Of course, we also compiled a security guide with more details on cryptography and security concepts. Read the security guide now.

Besides mitigating various attack vectors, such as a compromised database and OAuth 2.0 weaknesses, Ory Hydra is also able to securely manage JSON Web Keys. Click here to read more about security.

Ecosystem

Several guiding principles when it comes to our architecture design:

Minimal dependencies

Runs everywhere

Scales without effort

Minimize room for human and network errors

Who’s using ory hydra?

The Ory community stands on the shoulders of individuals, companies, and maintainers. this community is 1000+ strong and growing rapidly. The Ory stack protects 16.000.000.000+ API requests every month with over 250.000+ active service nodes.

Why Need Ory Hydra?

enable third-party solutions to access your APIs: This is what an OAuth2 Provider does, Hydra is a perfect fit.

you want to limit what type of information your backend services can read from each other. For example, the comment service should only be allowed to fetch user profile updates but shouldn't be able to read user passwords. OAuth 2.0 might make sense for you.

In this implementation, we only implement client_credentials grant types. Because currently, the goal is to provide third-party authentication only communication service-to-service without any end-user (first-party) intervention.

Why use Client Credentials as Authentication Flow?

The first decision is whether the party that requires access to resources is a machine. In the case of machine-to-machine authorization, the Client is also the Resource Owner, so no end-user authorization is needed. An example is a cron job that uses an API to import information to a database. In this example, the cron job is the Client and the Resource Owner since it holds the Client ID and Client Secret and uses them to get an Access Token from the Authorization Server.



In this case, we use this authentication flow because a service or machine represents as third-party.

How does it work in Algo?

[



](https://mermaid.live/edit#pako:eNqFk8FOwzAMhl_FygnE-gIVmjTGYLsgpCJOvaSJt1prk5K4g2nau5M2LWgVEz2l8lfn_3-7J6GsRpEKjx8tGoWPJHdO1rmB8MiWrWnrAl18X1aEhpP5_G5R7WwKa2IorN2T2cHidROhrtQh66N2MoWNYWd9g4phoRR6D292jyayL5YRKtwy2C0MHyxLVH1HGXHucPgskUt0cF-4OSTgmaoqEEwHhBtjGfCrIYf6diQyZRsE8qFAnlEPjiqGg6xIX3SPtf7-JEiP5t6vYL2_H-phsJ-V1NQhnEsmBpZC1sY2Q1gRwspjiCfK6eX-qeMPYtr-ybqCtB4V9o1XMY__fV4Hp9dcJ38XYxjiM_Ll_IojqB6CLlQT7KByyBNFyXhVKLXOTFaAuATfj7Vr0A9cMlkDTPWQDJowaDETNbpakg6LfeoKuQjLExiRhqOWbp-L3JwD1zZaMq40sXUi3cqQ3Ex0a58djRIpuxZHaPgzBur8DeHoFUA) 

sequenceDiagram
    autonumber
    Client->>+Algo: Hit booking API
    Algo->>+Hydra: Introspect Access Token
    Note left of Hydra: Checking access token whether <br> - still active (not expired) <br> - Scope is existed
    alt valid access token
    Hydra-->>Algo: Valid access token
    Algo-->>Algo: Booking Shipment
    Algo-->>Client: Success booking
    else Invalid Scope
    Hydra-->>Algo: Invalid Scope
    Algo-->>Client: Forbidden
    else Expired access token
    Hydra-->>Algo: Expired access token
    Algo-->>Client: Expired access token
    Client->>+Hydra: Get access token by client id and Secret
    Hydra-->>-Client: return access token with scope and expiration time
    end




List of APIs that we need

Add Client Oauth

Update Client Oauth

List Client

Get Detail Client

Update lifetime access token

Generate Access Token

Introspect Access Token

Delete Access Token

API documentation will be added.

Note: these above list is not final

Implementation Challenge

For now, The problem with this implementation is that some services (account, shipment) use the foreign key (account_id or customer_id) of users to do some action (booking, tracking, etc) meanwhile when we implement the client_id and client_secret authentication we have to ensure the client itself only mapping to one user [third party] so that we can store/map the account_id or customer_id.

Limitation

MySQL <= 5.6 / MariaDB

Ory Hydra has issues with MySQL <= 5.6 (but not MySQL 5.7+) and certain MariaDB versions. Read more about this here. Our recommendation is to use MySQL 5.7+ or PostgreSQL.

OAuth 2.0 client secret length

OAuth 2.0 Client Secrets are hashed using BCrypt. BCrypt has, by design, a maximum password length. The Golang BCrypt library has a maximum password length of 73 bytes. Any password longer will be "truncated":

hydra clients create --id long-secret \
  --secret 525348e77144a9cee9a7471a8b67c50ea85b9e3eb377a3c1a3a23db88f9150eefe76e6a339fdbc62b817595f53d72549d9ebe36438f8c2619846b963e9f43a94 \
  --endpoint http://localhost:4445 \
  --token-endpoint-auth-method client_secret_post \
  --grant-types client_credentials

hydra token client --client-id long-secret \
  --client-secret 525348e77144a9cee9a7471a8b67c50ea85b9e3eb377a3c1a3a23db88f9150eefe76e6a3 \
  --endpoint http://localhost:4444

Resource Owner Password Credentials grant type (ROPC)​

Ory Hydra doesn't and won't implement the Resource Owner Password Credentials Grant Type. Read on for context.

The ROPC grant type is discouraged by developers, professionals, and the IETF itself. It was originally added because big legacy corporations (not dropping any names, but they're part of the IETF consortium) didn't want to migrate their authentication infrastructure to the modern web but instead do what they've been doing all along "but OAuth 2.0" and for systems that want to upgrade from OAuth (1.0) to OAuth 2.0.

There are a ton of good reasons why this is a bad flow, they're summarized in this excellent blog article as well.

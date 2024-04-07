## IGDB Credentials

Each 2 month the IGDB client token must be refreshed manually

Run this request with `client_id` and `client_secret` to get the new `access_token` and rewrite it in the config file.

```
POST: https://id.twitch.tv/oauth2/token?client_id=abcdefg12345&client_secret=hijklmn67890&grant_type=client_credentials
```
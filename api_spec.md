# API Routes
## Character information
| Method | Route                         | Notes
|--------|-------------------------------|---
| GET    | /api/characters/              | List all characters
| POST   | /api/characters/              | Adds an individual character or bulk characters
| GET    | /api/characters/:characterID/ | Gets an individual characters information

## Corporation information
| Method | Route                             | Notes
|--------|-----------------------------------|---
| GET    | /api/corporations/                | List all corporations
| POST   | /api/corporations/                | Adds a corporations
| GET    | /api/corporations/:corporationID/ | Gets an individual corporations information

## Alliance Information
| Method | Route                          | Notes
|--------|--------------------------------|---
| GET    | /api/alliances/                | List all alliances
| POST   | /api/alliances/                | Adds a alliances
| GET    | /api/alliances/:corporationID/ | Gets an individual alliances information

## Pubsub routes
Routes used by Google PubSub

| Method | Route                          | Notes
|--------|--------------------------------|---
| POST   | /api/pubsub/characters/update/ | Updates a given characters information


## Updater routes
Routes used to trigger updating information internal to the service

| Method | Route                          | Notes
|--------|--------------------------------|---
| POST   | /api/updaters/characters/      | Finds N characters that need updating and sends to pubsub
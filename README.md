
# Caching Library for User Information

Create a library for caching user information to increase database throughput and availability.
Every request for user data should return the user information while simultaneously 
taking care of database throughput.
It means the library should take care of unnecessary requests for user data if some are pending.

## Why:

Let’s say we have a bunch of user information (identified by id) inside our main database.
Simultaneously there’s thousands of requests per second for user information.
We can consider that database is a bottleneck right. To avoid that we need a cache mechanism.

## Notes:

You can use any language, library, database and framework - anything you want and think is best for this case.

## Example:

If within 1000 requests there are 100 unique user ids then there should be only a maximum 100 requests into the database 
but all 1000 requests should get a response with a user data.

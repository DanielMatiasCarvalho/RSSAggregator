# RSSAggregator

This project is a simple backend webserver in `Go`, which stores posts from RSS Feeds, that are added to it.
It uses a `PostgreSQL` database, as well as `goose` and `sqlc` tools to interact with the database.

It can store users, and use their APIKey as an authentication, store Feeds along with their posts, as well as a user defining which Feeds they which to follow and see their posts.

This project was made during my learning the `Go` language in this course in youtube [Go Programming – Golang Course with Bonus Projects](https://www.youtube.com/watch?v=un6ZyFkqFKo)

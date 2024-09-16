module mutex/main

go 1.23.1

replace mutex/crawl => ./crawl

replace mutex/fetcher => ./fetcher

replace mutex/cache => ./cache

require (
	mutex/cache v0.0.0-00010101000000-000000000000
	mutex/crawl v0.0.0-00010101000000-000000000000
	mutex/fetcher v0.0.0-00010101000000-000000000000
)

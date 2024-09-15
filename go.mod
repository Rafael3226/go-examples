module mutex/main

go 1.21.5

replace mutex/crawl => ./crawl

replace mutex/fetcher => ./fetcher

require (
	mutex/crawl v0.0.0-00010101000000-000000000000
	mutex/fetcher v0.0.0-00010101000000-000000000000
)

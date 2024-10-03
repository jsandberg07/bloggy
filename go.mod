module bloggy

go 1.23.1

require internal/config v1.0.0

require (
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)

replace internal/config => ./internal/config

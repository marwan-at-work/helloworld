module helloworld

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/twitchtv/twirp/v6 v6.0.0
)

replace github.com/twitchtv/twirp/v6 => github.com/marwan-at-work/twirp/v6 v6.0.2

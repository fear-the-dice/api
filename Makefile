dev:
	FTDAUTHKEY=supersecret DB=ftd-test MONGOLAB_URI=mongodb://localhost REDIS_URL=redis://localhost:6379 PORT=:3000 go run server.go

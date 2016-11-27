Comentarismo Gender API

# A API to determine gender on names

# Options
```
GENDER_DEBUG, if true will debug all log entries for gender detection (optional)

LEARNGENDER, if true will use config_genderwords for each available lang, for testing proposes only (optional)
REDIS_HOST, ip addr of the redis instance to be used (required) -> defaults to g7-host
REDIS_PORT, port number of the redis instance to be used (required) defaults to 6379
REDIS_PASSWORD, password for this instance to be used (optional)
```

Running with expected cfg for dev, start port 3004 in debug mode and learn gender names
```
$ GENDER_DEBUG=true LEARNGENDER=true PORT=3004 godep go run main.go
```

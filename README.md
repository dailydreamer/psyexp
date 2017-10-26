# CSV format

ID, picturePicked, DecisionTime (seconds)

# Cross compile

```sh
cd bin
gox -osarch="linux/amd64 windows/386 windows/amd64" ..
```


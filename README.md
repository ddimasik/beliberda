Program generate logs and send them to stdout, logs are plain text and json, use CLI arg `-delay 5` to set log generation time in seconds (5 seconds in this case), use CLI arg `-size 1` to set log message size in kilobytes (1 kb in this case)


You can pass arguments to your Docker container like this:

```bash
docker run -it ddimasik/beliberda -delay 5 -size 1


kubectl run beliberda --image=ddimasik/beliberda -- -delay 5 -size 1
```

Docker repo: `ddimasik/beliberda`
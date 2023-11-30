you can pass arguments to your Docker container like this:

```bash
docker run -it your-image-name -delay 5 -size 2
```

In this command, `-delay 5 -size 2` are the arguments passed to your Go program.

`kubectl run beliberda --image=ddimasik/beliberda -- -delay 5 -size 1`
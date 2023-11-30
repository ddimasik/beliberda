you can pass arguments to your Docker container like this:

```bash
docker run -it ddimasik/beliberda -delay 5 -size 1
```

In this command, `-delay 5 -size 2` are the arguments passed to your Go program.

```bash
kubectl run beliberda --image=ddimasik/beliberda -- -delay 5 -size 1
```
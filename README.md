# banalplatitude
~long-lived ~helloworld


should be fine to cp the below into a ~/bin/bash with docker:

``` bash
#save some typing.
targetimg=docker.pkg.github.com/coip/banalplatitude/listener:latest

hostPort=8080 #configurable by you the reader in this var
svcPort=8080  #configurable by you the reader here in the src: https://github.com/coip/banalplatitude/blob/master/main.go#L18-L20

# pull listener!
docker pull $targetimg

# listen!
docker run              \
  --name acker          \
  -p $hostPort:$svcPort \
  -d                    \
  $targetimg

# showing latest listener entry, note size:
docker images | grep listener

# ~speak to listener
for i in $(seq 0 3)
do 
  curl -s -H "id: $i" localhost:8080
done

# stop listener (primarily here to flush logs)
docker stop acker

# note the lack of brackets on serverside output
docker logs acker

# remove unnecessary container (run live/non-daemon with --rm for likely better lifecycle)
docker rm acker

# remove image
docker rmi $targetimg
```

----


## _if you need to auth to the github pkg docker registry:_

`echo "$PAT" | docker login docker.pkg.github.com -u $USER --password-stdin`

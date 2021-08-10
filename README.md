# banalplatitude
~long-lived helloworld


should be fine to cp the below into a ~/bin/bash with docker:

``` bash
function holdonasec { sleep 4; clear; }

#save some typing.
targetimg=docker.pkg.github.com/coip/banalplatitude/listener:latest
hostPort=8080 #configurable by you the reader in this var
svcPort=8080  #configurable by you the reader here in the src: https://github.com/coip/banalplatitude/blob/master/main.go#L18-L20
clear

# pull listener!
docker pull $targetimg
holdonasec

# listen!
docker run              \
  --name acker          \
  -p $hostPort:$svcPort \
  -d                    \
  $targetimg
holdonasec

# showing latest listener entry, note size:
docker images | grep listener
holdonasec

# ~speak to listener
for i in $(seq 0 3)
do 
  curl -s -H "id: $i" localhost:${hostPort}
done
holdonasec

# stop listener (primarily here to flush logs)
docker stop acker
holdonasec

# note the lack of brackets on serverside output
docker logs acker
holdonasec

# remove unnecessary container (run live/non-daemon with --rm for likely better lifecycle)
docker rm acker
holdonasec

# remove image
docker rmi $targetimg
holdonasec
unset -f holdonasec
```

----


## _if you need to auth to the github pkg docker registry:_

`echo "$PAT" | docker login docker.pkg.github.com -u $USER --password-stdin`

# go-google-gateway
Simple google cloud gateway (don't use in production)

[![Build Status](https://travis-ci.org/yroffin/go-google-gateway.svg?branch=master)](https://travis-ci.org/yroffin/go-google-gateway)

# Developement

With gin (for live reload)

gin -p 3000 -a 8080 -i -- -http=8080

# Dependencies and setup

## Setup on raspberry pi 2 or zero

        pi@raspberrypi:~ $ sudo userdel -r google
        pi@raspberrypi:~ $ sudo useradd -m google
        pi@raspberrypi:~ $ export GITHUB=https://github.com/yroffin/go-google-gateway/releases/download/1.0e
        pi@raspberrypi:~ $ sudo wget ${GITHUB}/go-google-gateway-0.0.1-SNAPSHOT.armel -O /home/google/go-google-gateway-0.0.1-SNAPSHOT.arm
        or
        pi@raspberrypi:~ $ sudo wget ${GITHUB}/go-google-gateway-0.0.1-SNAPSHOT.armhf -O /home/google/go-google-gateway-0.0.1-SNAPSHOT.arm
        pi@raspberrypi:~ $ sudo chmod 755 /home/google/go-google-gateway-0.0.1-SNAPSHOT.arm
        pi@raspberrypi:~ $ sudo chown google:google /home/google/go-google-gateway-0.0.1-SNAPSHOT.arm
        pi@raspberrypi:~ $ sudo wget ${GITHUB}/go-google-gateway-service -O /etc/init.d/go-google-gateway-service
        pi@raspberrypi:~ $ sudo chmod 755 /etc/init.d/go-google-gateway-service
        pi@raspberrypi:~ $ sudo update-rc.d go-google-gateway-service defaults
        pi@raspberrypi:~ $ sudo service go-google-gateway-service restart
        pi@raspberrypi:~ $ curl http://locahost:8081/api/version

# Roadmap

More security
And publish my dialogflow on google

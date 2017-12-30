
WHERE=$(PWD)
MODULE_ARMEL=${WHERE}/go-google-gateway-0.0.1-SNAPSHOT.armel
MODULE_ARMHF=${WHERE}/go-google-gateway-0.0.1-SNAPSHOT.armhf

all: clean ${MODULE_ARMEL} ${MODULE_ARMHF}

clean:
	rm -rf ${WHERE}/armel
	rm -rf ${WHERE}/armhf
	rm -f ${MODULE_ARMHF}
	rm -f ${MODULE_ARMEL}

${MODULE_ARMEL}:
	# module
	CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 CGO_CFLAGS="-march=armv6j -mfloat-abi=soft" go install -installsuffix armel
	mv -f ${GOPATH}/bin/linux_arm/jarvis-go-ext ${GOPATH}/go-google-gateway-SNAPSHOT.armel

${MODULE_ARMHF}:
	# module
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CGO_CFLAGS="" go install -installsuffix armhf
	mv -f ${GOPATH}/bin/linux_arm/jarvis-go-ext ${GOPATH}/go-google-gateway-SNAPSHOT.armhf

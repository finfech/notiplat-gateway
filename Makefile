.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/email-gw cmd/email-gw/main.go cmd/email-gw/wire_gen.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/sms-gw cmd/sms-gw/main.go cmd/sms-gw/wire_gen.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/smsme-gw cmd/smsme-gw/main.go cmd/smsme-gw/wire_gen.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

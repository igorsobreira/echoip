
.PHONY: clean

all: bin/linux-amd64/echoip \
	bin/darwin-amd64/echoip \
	bin/windows-amd64/echoip \
	bin/freebsd-amd64/echoip

bin/%/echoip:
	@bash -c "source ~/dev/golang-crosscompile/crosscompile.bash; \
	mkdir -p bin/$*; \
	go-$* build -o bin/$*/echoip echoip.go"

clean:
	@rm -rf bin/

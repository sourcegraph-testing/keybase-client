COVERWEB=go tool cover -html=/tmp/count.out

fmt:
	go fmt ./...

vet:
	go vet ./...

# KBFS uses stricter linting than client, so the errors that should be
# filtered only for non-kbfs files must use perl regexes, and start
# with "^((?!kbfs))".
lint:
	( \
		which golint >/dev/null 2>&1 \
		|| ( \
			echo "Missing golint" \
			&& exit 1 \
		) \
	) \
	&& ( \
		golint ./... \
		| grep -v "protocol[\/\\]" \
		| grep -v -P -e '^((?!kbfs)).*comment' \
		| grep -v KeysById \
		| grep -v "error should be the last type" \
		| grep -v -P -e '^((?!kbfs)).*stutters' \
		| grep -v -P -e '^((?!kbfs)).*error strings should not be capitalized' \
		| grep -v prot.go \
		| grep -v -P -e '^((?!kbfs)).*redundant if ...; err != nil check, just return error instead' \
		| grep -v -P -e '^((?!kbfs)).*use ALL_CAPS in Go names; use CamelCase' \
		| grep -v runquiet.go \
		| grep -v mocks_test\.go \
		| grep -v mock_codec\.go \
		| grep -v "use underscores in Go names" \
		&& exit 1 \
		|| echo "Lint-free!" \
	)

# to run splint, first do this:  go get stathat.com/c/splint
splint:
	splint -s 100 `find . -name "*.go"`

cover:
	go test -covermode=count ./...

test:
	test/check_or_start_kbweb.sh
	go test ./...
	test/pause_kbweb.sh

testclean:
	test/nuke_kbweb.sh

coverclient:
	$(COVERTEST) ./client
	$(COVERWEB)

coverdaemon:
	$(COVERTEST) ./daemon
	$(COVERWEB)

coverengine:
	$(COVERTEST) ./engine
	$(COVERWEB)

coverlibcmdline:
	$(COVERTEST) ./libcmdline
	$(COVERWEB)

coverlibkb:
	$(COVERTEST) ./libkb
	$(COVERWEB)

coverclean:
	rm /tmp/count.out

gen-deps:
	./gen_deps.pl

linux-clear-devel-secrets:
	secret-tool clear service-base keybase-devel

linux-clear-test-secrets:
	secret-tool clear service-base keybase-devel-test

GOLANGCI_RUN_OPT = --color auto

golangci-lint-kbfs:
	go list -f '{{.Dir}}' ./kbfs/... | fgrep -v dokan | xargs realpath --relative-to=. | xargs golangci-lint run $(GOLANGCI_RUN_OPT) --deadline 5m0s

golangci-lint-nonkbfs:
	go list -f '{{.Dir}}' ./... | fgrep -v protocol  | fgrep -v kbfs | xargs realpath --relative-to=. | xargs golangci-lint run $(GOLANGCI_RUN_OPT) --deadline 10m0s

.PHONY: test splint fmt lint vet cover testclean coverclient coverdaemon coverengine coverlibcmdline coverlibkb coverclean gen-deps

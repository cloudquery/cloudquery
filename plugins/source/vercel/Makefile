# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/vercel
	go run main.go doc ../../../website/tables/vercel
	cp ../../../website/tables/vercel/README.md ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/vercel/README.md
	mv ../../../website/tables/vercel/README.md ../../../website/pages/docs/plugins/sources/vercel/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/vercel/*.md
	rm -rf ../../../website/tables/vercel/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-docs

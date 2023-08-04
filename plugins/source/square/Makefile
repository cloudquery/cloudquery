test:
	pytest .

fmt:
	black . --exclude=cloudquery

fmt-check:
	black --check . --exclude=cloudquery

gen-docs:
	rm -rf ../../../website/tables/square
	python3 main.py doc --format=markdown ../../../website/tables/square
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/square/README.md
	mv ../../../website/tables/square/README.md ../../../website/pages/docs/plugins/sources/square/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/square/*.md
	rm -rf ../../../website/tables/square/*.bak

gen: gen-docs

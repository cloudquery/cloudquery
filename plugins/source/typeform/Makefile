test:
	pytest .

fmt:
	black .

fmt-check:
	black --check .

gen-docs:
	rm -rf ../../../website/tables/typeform
	python3 main.py doc --format=markdown ../../../website/tables/typeform
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/typeform/README.md
	mv ../../../website/tables/typeform/README.md ../../../website/pages/docs/plugins/sources/typeform/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/typeform/*.md
	rm -rf ../../../website/tables/typeform/*.bak

gen: gen-docs

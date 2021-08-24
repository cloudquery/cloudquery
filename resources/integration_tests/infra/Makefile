.PHONY: init
init:
	@./scripts/init.sh

.PHONY: apply
apply:
	@terraform apply --var test_suffix="integration" --var test_prefix="cq-testing"

.PHONY: destroy
destroy:
	@terraform destroy --var test_suffix="integration" --var test_prefix="cq-testing"

.PHONY: clean
clean:
	@./scripts/clean.sh

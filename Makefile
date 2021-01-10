PACKAGE_NAME          := github.com/troian/golang-cross-example
GOLANG_CROSS_VERSION  ?= v1.15.2

.PHONY: release-dry-run
release-dry-run:
	@docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		troian/golang-cross:${GOLANG_CROSS_VERSION} \
		--rm-dist --skip-validate --skip-publish

.PHONY: release
release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		troian/golang-cross:${GOLANG_CROSS_VERSION} \
		release --rm-dist

.PHONY: docker
docker:
	@docker build -t cloudquery:latest .

.PHONY: docker-test
docker-test:
	@docker build -t cloudquery:test -f Dockerfile.test .

.PHONY: docker-publish
	@docker tag cloudquery-test:latest public.ecr.aws/d1r4a9i9/cloudquery-test:latest \
	&& docker push public.ecr.aws/d1r4a9i9/cloudquery-test:latest

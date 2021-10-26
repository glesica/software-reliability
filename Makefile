# This Makefile knows how to build the various
# pieces of this repo, including converting
# Markdown to HTML.

IMAGE_NAME := glesica/software-reliability:latest

README_INPUTS := $(shell find . -type f -name README.md)
README_OUTPUTS := $(README_INPUTS:README.md=index.html)

OTHER_INPUTS := $(shell find . -type f ! -name README.md ! -name index.md -name '*.md')
OTHER_OUTPUTS := $(OTHER_INPUTS:.md=.html)

INCLUDES := $(shell find includes -type f)

SITE_TITLE := [software reliability]
PANDOC := pandoc -s --template includes/template.html -H includes/header.html

.PHONY: build
build: readmes others

.PHONY: build-container
build-container:
	@docker run --mount src="${PWD}",dst=/code,type=bind ${IMAGE_NAME}

.PHONY: container
container:
	@docker build -t ${IMAGE_NAME} .

.PHONY: readmes
readmes: ${README_OUTPUTS}

.PHONY: others
others: ${OTHER_OUTPUTS}

.PHONY: serve
serve:
	@echo "Serving on http://localhost:8080"
	@echo "Use ctrl-c to terminate the server"
	@python3 -m http.server 8080

index.html: README.md ${INCLUDES}
	@echo "building $<"
	@${PANDOC} -V rooturl="." -M pagetitle:"/ ${SITE_TITLE}" -o "$@" "$<"

%/index.html: %/README.md ${INCLUDES} %/media
	@echo "building $<"
	@${PANDOC} -V rooturl="$$(realpath --relative-to=$(@D) .)" -M pagetitle:"/$(@D) ${SITE_TITLE}" -o "$@" "$<"

%.html: %.md ${INCLUDES}
	@echo "building $<"
	@${PANDOC} -V rooturl="$$(realpath --relative-to=$(@D) .)" -M pagetitle:"/$(@D) ${SITE_TITLE}" -o "$@" "$<"

.PRECIOUS: %/media
%/media: media
	@echo "symlinking media $<"
	@ln -s "$$(realpath --relative-to=$(@D) media)" "$@"


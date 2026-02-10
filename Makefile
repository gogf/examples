SHELL := /bin/bash

# Update all goframe module dependencies to their latest versions.
.PHONY: update
update:
	./.make_update.sh

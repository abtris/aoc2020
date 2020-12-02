CURRENTDAY=$(shell date +'%d')

directory = "day${CURRENTDAY}"

all: | $(directory)
	@touch "${directory}/main.go"
	@touch "${directory}/main_test.go"
	@touch "${directory}/input"

$(directory):
	@echo "Folder $(directory) does not exist"
	mkdir -p $@

.PHONY: all

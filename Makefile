CURRENTDAY=$(shell date +'%e')

directory = "day${CURRENTDAY}"

all: | $(directory)
	@touch "day${CURRENTDAY}/main.go"
	@touch "day${CURRENTDAY}/main_test.go"
	@touch "day${CURRENTDAY}/input"

$(directory):
	@echo "Folder $(directory) does not exist"
	mkdir -p $@

.PHONY: all

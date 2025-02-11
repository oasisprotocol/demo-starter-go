#!/usr/bin/env gmake

include ../../common.mk

.PHONY: MessageBox clean

MessageBox:
	@printf "$(CYAN)*** Compiling $(BLUE)$@$(CYAN)...$(OFF)\n"
	@solc --evm-version paris --bin --abi --optimize --overwrite -o . $@.sol
	@mv $@.bin $@.hex
	@printf "$(CYAN)*** Generating Go bindings for $(BLUE)$@$(CYAN)...$(OFF)\n"
	@abigen --abi=$@.abi --bin=$@.hex --pkg message_box --out=$@.go

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@-rm -f MessageBox.abi MessageBox.hex MessageBox.go


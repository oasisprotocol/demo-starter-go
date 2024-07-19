#!/usr/bin/env gmake

include ../common.mk

.PHONY: all MessageBox clean

all: MessageBox
	@printf "$(CYAN)*** Contracts built successfully!$(OFF)\n"

MessageBox:
	@printf "$(CYAN)*** Building contract: $(BLUE)MessageBox$(CYAN)...$(OFF)\n"
	@$(MAKE) -C message-box

clean:
	@printf "$(CYAN)*** Cleaning up...$(OFF)\n"
	@$(MAKE) -C message-box clean


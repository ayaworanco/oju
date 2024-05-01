TARGET = oju
BUILDDIR = build

run: $(BUILDDIR)/Makefile
	@cd $(BUILDDIR) && ./app/$(TARGET)

test: $(BUILDDIR)/Makefile
	@cd $(BUILDDIR) && ctest --output-on-failure --gtest_catch_exceptions=0

build: $(BUILDDIR)/Makefile
	@$(MAKE) -C $(BUILDDIR)

$(BUILDDIR)/Makefile:
	@mkdir -p $(BUILDDIR)
	@cd $(BUILDDIR) && cmake ..

clean:
	@$(MAKE) -C $(BUILDDIR) clean
	@rm -rf $(BUILDDIR)

.PHONY: build clean

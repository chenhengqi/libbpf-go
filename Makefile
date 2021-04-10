OUTPUT := .output
LIBBPF_SRC := $(abspath ./libbpf/src)
LIBBPF_OBJ := $(abspath $(OUTPUT)/libbpf.a)
LIBBPF_GEN_GO := libbpf-gen-go
EXAMPLE := example

all: $(LIBBPF_OBJ) $(LIBBPF_GEN_GO)

$(OUTPUT) $(OUTPUT)/libbpf:
	mkdir -p $@

$(LIBBPF_OBJ): $(wildcard $(LIBBPF_SRC)/*.[ch]) | $(OUTPUT)/libbpf
	$(MAKE) -C $(LIBBPF_SRC) BUILD_STATIC_ONLY=1        \
		OBJDIR=$(dir $@)/libbpf DESTDIR=$(dir $@)       \
		INCLUDEDIR= LIBDIR= UAPIDIR=                    \
		install

$(LIBBPF_GEN_GO): *.go
	go build -o $(LIBBPF_GEN_GO)

.PHONY: example
example: $(LIBBPF_GEN_GO)
	bpftool gen skeleton $(EXAMPLE)/biolatency.bpf.o > $(EXAMPLE)/biolatency.skel.h

run: $(LIBBPF_GEN_GO)
	./$(LIBBPF_GEN_GO) $(EXAMPLE)/biolatency.bpf.o

clean:
	rm -rf $(OUTPUT)
	rm -f $(LIBBPF_GEN_GO)

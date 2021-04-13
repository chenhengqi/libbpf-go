package parse

import "strings"

var types map[string]struct{}

func init() {
	types = make(map[string]struct{})
	// primitive types
	types["void"] = struct{}{}
	types["bool"] = struct{}{}
	types["char"] = struct{}{}
	types["int"] = struct{}{}
	types["long"] = struct{}{}
	types["__u8"] = struct{}{}
	types["__u32"] = struct{}{}
	types["__u64"] = struct{}{}
	types["__s32"] = struct{}{}
	types["__s64"] = struct{}{}
	types["size_t"] = struct{}{}
	types["pid_t"] = struct{}{}

	types["bpf_attach_type"] = struct{}{}
	types["bpf_create_map_attr"] = struct{}{}
	types["bpf_func_id"] = struct{}{}
	types["bpf_insn"] = struct{}{}
	types["bpf_iter_attach_opts"] = struct{}{}
	types["bpf_line_info"] = struct{}{}
	types["bpf_link_create_opts"] = struct{}{}
	types["bpf_linker_opts"] = struct{}{}
	types["bpf_linker"] = struct{}{}
	types["bpf_link"] = struct{}{}
	types["bpf_link_update_opts"] = struct{}{}
	types["bpf_load_program_attr"] = struct{}{}
	types["bpf_map_batch_opts"] = struct{}{}
	types["bpf_map_clear_priv_t"] = struct{}{}
	types["bpf_map_def"] = struct{}{}
	types["bpf_map"] = struct{}{}
	types["bpf_map_type"] = struct{}{}
	types["bpf_object_clear_priv_t"] = struct{}{}
	types["bpf_object_load_attr"] = struct{}{}
	types["bpf_object_open_attr"] = struct{}{}
	types["bpf_object_open_opts"] = struct{}{}
	types["bpf_object_skeleton"] = struct{}{}
	types["bpf_object"] = struct{}{}
	types["bpf_perf_event_print_t"] = struct{}{}
	types["bpf_perf_event_ret"] = struct{}{}
	types["bpf_prog_attach_opts"] = struct{}{}
	types["bpf_prog_bind_opts"] = struct{}{}
	types["bpf_prog_info_linear"] = struct{}{}
	types["bpf_prog_info"] = struct{}{}
	types["bpf_prog_linfo"] = struct{}{}
	types["bpf_prog_load_attr"] = struct{}{}
	types["bpf_program_clear_priv_t"] = struct{}{}
	types["bpf_program_prep_t"] = struct{}{}
	types["bpf_program"] = struct{}{}
	types["bpf_prog_test_run_attr"] = struct{}{}
	types["bpf_prog_type"] = struct{}{}
	types["bpf_stats_type"] = struct{}{}
	types["bpf_test_run_opts"] = struct{}{}
	types["bpf_xdp_set_link_opts"] = struct{}{}
	types["btf_dedup_opts"] = struct{}{}
	types["btf_dump_emit_type_decl_opts"] = struct{}{}
	types["btf_dump_opts"] = struct{}{}
	types["btf_dump_printf_fn_t"] = struct{}{}
	types["btf_dump"] = struct{}{}
	types["btf_endianness"] = struct{}{}
	types["btf_ext"] = struct{}{}
	types["btf_func_linkage"] = struct{}{}
	types["btf_fwd_kind"] = struct{}{}
	types["btf"] = struct{}{}
	types["btf_type"] = struct{}{}
	types["libbpf_print_fn_t"] = struct{}{}
	types["perf_buffer_opts"] = struct{}{}
	types["perf_buffer_raw_opts"] = struct{}{}
	types["perf_buffer"] = struct{}{}
	types["ring_buffer_opts"] = struct{}{}
	types["ring_buffer_sample_fn"] = struct{}{}
	types["ring_buffer"] = struct{}{}
	types["xdp_link_info"] = struct{}{}
	types["xsk_ring_cons"] = struct{}{}
	types["xsk_ring_prod"] = struct{}{}
	types["xsk_socket_config"] = struct{}{}
	types["xsk_socket"] = struct{}{}
	types["xsk_umem_config"] = struct{}{}
	types["xsk_umem"] = struct{}{}
}

func RegisterTypes(ctypes ...string) {
	for _, ctype := range ctypes {
		types[ctype] = struct{}{}
	}
}

// TypeSpec represents a C type specification
type TypeSpec struct {
	IsVoid       bool
	IsStruct     bool
	IsEnum       bool
	IsConst      bool
	IsUnsigned   bool
	IsPointer    bool
	Indirections int    // `*` ~ 1, `**` ~ 2
	Name         string // function name for return type, argument name for arguments
	TypeName     string
}

// GoName return a capitalized name for Go as an exported identifier
func (t *TypeSpec) GoName() string {
	return strings.Title(t.Name)
}

// AUTOGENERATED BY LIBBPF-GO-GEN
// DO NOT EDIT
package libbpf

// #include <bpf/bpf.h>
import "C"
import "unsafe"

var _ unsafe.Pointer

// Bpf_create_map_xattr is a wrapper of bpf_create_map_xattr
//     int bpf_create_map_xattr(const struct bpf_create_map_attr *create_attr)
func Bpf_create_map_xattr(create_attr *Struct_bpf_create_map_attr) int32 {
	return C.bpf_create_map_xattr()
}

// Bpf_create_map_node is a wrapper of bpf_create_map_node
//     int bpf_create_map_node(enum bpf_map_type map_type, const char *name, int key_size, int value_size, int max_entries, __u32 map_flags, int node)
func Bpf_create_map_node(map_type Enum_bpf_map_type, name string, key_size int32, value_size int32, max_entries int32, map_flags uint32, node int32) int32 {
	return C.bpf_create_map_node()
}

// Bpf_create_map_name is a wrapper of bpf_create_map_name
//     int bpf_create_map_name(enum bpf_map_type map_type, const char *name, int key_size, int value_size, int max_entries, __u32 map_flags)
func Bpf_create_map_name(map_type Enum_bpf_map_type, name string, key_size int32, value_size int32, max_entries int32, map_flags uint32) int32 {
	return C.bpf_create_map_name()
}

// Bpf_create_map is a wrapper of bpf_create_map
//     int bpf_create_map(enum bpf_map_type map_type, int key_size, int value_size, int max_entries, __u32 map_flags)
func Bpf_create_map(map_type Enum_bpf_map_type, key_size int32, value_size int32, max_entries int32, map_flags uint32) int32 {
	return C.bpf_create_map()
}

// Bpf_create_map_in_map_node is a wrapper of bpf_create_map_in_map_node
//     int bpf_create_map_in_map_node(enum bpf_map_type map_type, const char *name, int key_size, int inner_map_fd, int max_entries, __u32 map_flags, int node)
func Bpf_create_map_in_map_node(map_type Enum_bpf_map_type, name string, key_size int32, inner_map_fd int32, max_entries int32, map_flags uint32, node int32) int32 {
	return C.bpf_create_map_in_map_node()
}

// Bpf_create_map_in_map is a wrapper of bpf_create_map_in_map
//     int bpf_create_map_in_map(enum bpf_map_type map_type, const char *name, int key_size, int inner_map_fd, int max_entries, __u32 map_flags)
func Bpf_create_map_in_map(map_type Enum_bpf_map_type, name string, key_size int32, inner_map_fd int32, max_entries int32, map_flags uint32) int32 {
	return C.bpf_create_map_in_map()
}

// Bpf_load_program_xattr is a wrapper of bpf_load_program_xattr
//     int bpf_load_program_xattr(const struct bpf_load_program_attr *load_attr, char *log_buf, size_t log_buf_sz)
func Bpf_load_program_xattr(load_attr *Struct_bpf_load_program_attr, log_buf *C.char, log_buf_sz uint) int32 {
	return C.bpf_load_program_xattr()
}

// Bpf_load_program is a wrapper of bpf_load_program
//     int bpf_load_program(enum bpf_prog_type type, const struct bpf_insn *insns, size_t insns_cnt, const char *license, __u32 kern_version, char *log_buf, size_t log_buf_sz)
func Bpf_load_program(type_ Enum_bpf_prog_type, insns *Struct_bpf_insn, insns_cnt uint, license string, kern_version uint32, log_buf *C.char, log_buf_sz uint) int32 {
	return C.bpf_load_program()
}

// Bpf_verify_program is a wrapper of bpf_verify_program
//     int bpf_verify_program(enum bpf_prog_type type, const struct bpf_insn *insns, size_t insns_cnt, __u32 prog_flags, const char *license, __u32 kern_version, char *log_buf, size_t log_buf_sz, int log_level)
func Bpf_verify_program(type_ Enum_bpf_prog_type, insns *Struct_bpf_insn, insns_cnt uint, prog_flags uint32, license string, kern_version uint32, log_buf *C.char, log_buf_sz uint, log_level int32) int32 {
	return C.bpf_verify_program()
}

// Bpf_map_update_elem is a wrapper of bpf_map_update_elem
//     int bpf_map_update_elem(int fd, const void *key, const void *value, __u64 flags)
func Bpf_map_update_elem(fd int32, key unsafe.Pointer, value unsafe.Pointer, flags uint64) int32 {
	return C.bpf_map_update_elem()
}

// Bpf_map_lookup_elem is a wrapper of bpf_map_lookup_elem
//     int bpf_map_lookup_elem(int fd, const void *key, void *value)
func Bpf_map_lookup_elem(fd int32, key unsafe.Pointer, value unsafe.Pointer) int32 {
	return C.bpf_map_lookup_elem()
}

// Bpf_map_lookup_elem_flags is a wrapper of bpf_map_lookup_elem_flags
//     int bpf_map_lookup_elem_flags(int fd, const void *key, void *value, __u64 flags)
func Bpf_map_lookup_elem_flags(fd int32, key unsafe.Pointer, value unsafe.Pointer, flags uint64) int32 {
	return C.bpf_map_lookup_elem_flags()
}

// Bpf_map_lookup_and_delete_elem is a wrapper of bpf_map_lookup_and_delete_elem
//     int bpf_map_lookup_and_delete_elem(int fd, const void *key, void *value)
func Bpf_map_lookup_and_delete_elem(fd int32, key unsafe.Pointer, value unsafe.Pointer) int32 {
	return C.bpf_map_lookup_and_delete_elem()
}

// Bpf_map_delete_elem is a wrapper of bpf_map_delete_elem
//     int bpf_map_delete_elem(int fd, const void *key)
func Bpf_map_delete_elem(fd int32, key unsafe.Pointer) int32 {
	return C.bpf_map_delete_elem()
}

// Bpf_map_get_next_key is a wrapper of bpf_map_get_next_key
//     int bpf_map_get_next_key(int fd, const void *key, void *next_key)
func Bpf_map_get_next_key(fd int32, key unsafe.Pointer, next_key unsafe.Pointer) int32 {
	return C.bpf_map_get_next_key()
}

// Bpf_map_freeze is a wrapper of bpf_map_freeze
//     int bpf_map_freeze(int fd)
func Bpf_map_freeze(fd int32) int32 {
	return C.bpf_map_freeze()
}

// Bpf_map_delete_batch is a wrapper of bpf_map_delete_batch
//     int bpf_map_delete_batch(int fd, void *keys, __u32 *count, const struct bpf_map_batch_opts *opts)
func Bpf_map_delete_batch(fd int32, keys unsafe.Pointer, count *uint32, opts *Struct_bpf_map_batch_opts) int32 {
	return C.bpf_map_delete_batch()
}

// Bpf_map_lookup_batch is a wrapper of bpf_map_lookup_batch
//     int bpf_map_lookup_batch(int fd, void *in_batch, void *out_batch, void *keys, void *values, __u32 *count, const struct bpf_map_batch_opts *opts)
func Bpf_map_lookup_batch(fd int32, in_batch unsafe.Pointer, out_batch unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer, count *uint32, opts *Struct_bpf_map_batch_opts) int32 {
	return C.bpf_map_lookup_batch()
}

// Bpf_map_lookup_and_delete_batch is a wrapper of bpf_map_lookup_and_delete_batch
//     int bpf_map_lookup_and_delete_batch(int fd, void *in_batch, void *out_batch, void *keys, void *values, __u32 *count, const struct bpf_map_batch_opts *opts)
func Bpf_map_lookup_and_delete_batch(fd int32, in_batch unsafe.Pointer, out_batch unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer, count *uint32, opts *Struct_bpf_map_batch_opts) int32 {
	return C.bpf_map_lookup_and_delete_batch()
}

// Bpf_map_update_batch is a wrapper of bpf_map_update_batch
//     int bpf_map_update_batch(int fd, void *keys, void *values, __u32 *count, const struct bpf_map_batch_opts *opts)
func Bpf_map_update_batch(fd int32, keys unsafe.Pointer, values unsafe.Pointer, count *uint32, opts *Struct_bpf_map_batch_opts) int32 {
	return C.bpf_map_update_batch()
}

// Bpf_obj_pin is a wrapper of bpf_obj_pin
//     int bpf_obj_pin(int fd, const char *pathname)
func Bpf_obj_pin(fd int32, pathname string) int32 {
	return C.bpf_obj_pin()
}

// Bpf_obj_get is a wrapper of bpf_obj_get
//     int bpf_obj_get(const char *pathname)
func Bpf_obj_get(pathname string) int32 {
	return C.bpf_obj_get()
}

// Bpf_prog_attach is a wrapper of bpf_prog_attach
//     int bpf_prog_attach(int prog_fd, int attachable_fd, enum bpf_attach_type type, unsigned int flags)
func Bpf_prog_attach(prog_fd int32, attachable_fd int32, type_ Enum_bpf_attach_type, flags int32) int32 {
	return C.bpf_prog_attach()
}

// Bpf_prog_attach_xattr is a wrapper of bpf_prog_attach_xattr
//     int bpf_prog_attach_xattr(int prog_fd, int attachable_fd, enum bpf_attach_type type, const struct bpf_prog_attach_opts *opts)
func Bpf_prog_attach_xattr(prog_fd int32, attachable_fd int32, type_ Enum_bpf_attach_type, opts *Struct_bpf_prog_attach_opts) int32 {
	return C.bpf_prog_attach_xattr()
}

// Bpf_prog_detach is a wrapper of bpf_prog_detach
//     int bpf_prog_detach(int attachable_fd, enum bpf_attach_type type)
func Bpf_prog_detach(attachable_fd int32, type_ Enum_bpf_attach_type) int32 {
	return C.bpf_prog_detach()
}

// Bpf_prog_detach2 is a wrapper of bpf_prog_detach2
//     int bpf_prog_detach2(int prog_fd, int attachable_fd, enum bpf_attach_type type)
func Bpf_prog_detach2(prog_fd int32, attachable_fd int32, type_ Enum_bpf_attach_type) int32 {
	return C.bpf_prog_detach2()
}

// Bpf_link_create is a wrapper of bpf_link_create
//     int bpf_link_create(int prog_fd, int target_fd, enum bpf_attach_type attach_type, const struct bpf_link_create_opts *opts)
func Bpf_link_create(prog_fd int32, target_fd int32, attach_type Enum_bpf_attach_type, opts *Struct_bpf_link_create_opts) int32 {
	return C.bpf_link_create()
}

// Bpf_link_detach is a wrapper of bpf_link_detach
//     int bpf_link_detach(int link_fd)
func Bpf_link_detach(link_fd int32) int32 {
	return C.bpf_link_detach()
}

// Bpf_link_update is a wrapper of bpf_link_update
//     int bpf_link_update(int link_fd, int new_prog_fd, const struct bpf_link_update_opts *opts)
func Bpf_link_update(link_fd int32, new_prog_fd int32, opts *Struct_bpf_link_update_opts) int32 {
	return C.bpf_link_update()
}

// Bpf_iter_create is a wrapper of bpf_iter_create
//     int bpf_iter_create(int link_fd)
func Bpf_iter_create(link_fd int32) int32 {
	return C.bpf_iter_create()
}

// Bpf_prog_test_run_xattr is a wrapper of bpf_prog_test_run_xattr
//     int bpf_prog_test_run_xattr(struct bpf_prog_test_run_attr *test_attr)
func Bpf_prog_test_run_xattr(test_attr *Struct_bpf_prog_test_run_attr) int32 {
	return C.bpf_prog_test_run_xattr()
}

// Bpf_prog_test_run is a wrapper of bpf_prog_test_run
//     int bpf_prog_test_run(int prog_fd, int repeat, void *data, __u32 size, void *data_out, __u32 *size_out, __u32 *retval, __u32 *duration)
func Bpf_prog_test_run(prog_fd int32, repeat int32, data unsafe.Pointer, size uint32, data_out unsafe.Pointer, size_out *uint32, retval *uint32, duration *uint32) int32 {
	return C.bpf_prog_test_run()
}

// Bpf_prog_get_next_id is a wrapper of bpf_prog_get_next_id
//     int bpf_prog_get_next_id(__u32 start_id, __u32 *next_id)
func Bpf_prog_get_next_id(start_id uint32, next_id *uint32) int32 {
	return C.bpf_prog_get_next_id()
}

// Bpf_map_get_next_id is a wrapper of bpf_map_get_next_id
//     int bpf_map_get_next_id(__u32 start_id, __u32 *next_id)
func Bpf_map_get_next_id(start_id uint32, next_id *uint32) int32 {
	return C.bpf_map_get_next_id()
}

// Bpf_btf_get_next_id is a wrapper of bpf_btf_get_next_id
//     int bpf_btf_get_next_id(__u32 start_id, __u32 *next_id)
func Bpf_btf_get_next_id(start_id uint32, next_id *uint32) int32 {
	return C.bpf_btf_get_next_id()
}

// Bpf_link_get_next_id is a wrapper of bpf_link_get_next_id
//     int bpf_link_get_next_id(__u32 start_id, __u32 *next_id)
func Bpf_link_get_next_id(start_id uint32, next_id *uint32) int32 {
	return C.bpf_link_get_next_id()
}

// Bpf_prog_get_fd_by_id is a wrapper of bpf_prog_get_fd_by_id
//     int bpf_prog_get_fd_by_id(__u32 id)
func Bpf_prog_get_fd_by_id(id uint32) int32 {
	return C.bpf_prog_get_fd_by_id()
}

// Bpf_map_get_fd_by_id is a wrapper of bpf_map_get_fd_by_id
//     int bpf_map_get_fd_by_id(__u32 id)
func Bpf_map_get_fd_by_id(id uint32) int32 {
	return C.bpf_map_get_fd_by_id()
}

// Bpf_btf_get_fd_by_id is a wrapper of bpf_btf_get_fd_by_id
//     int bpf_btf_get_fd_by_id(__u32 id)
func Bpf_btf_get_fd_by_id(id uint32) int32 {
	return C.bpf_btf_get_fd_by_id()
}

// Bpf_link_get_fd_by_id is a wrapper of bpf_link_get_fd_by_id
//     int bpf_link_get_fd_by_id(__u32 id)
func Bpf_link_get_fd_by_id(id uint32) int32 {
	return C.bpf_link_get_fd_by_id()
}

// Bpf_obj_get_info_by_fd is a wrapper of bpf_obj_get_info_by_fd
//     int bpf_obj_get_info_by_fd(int bpf_fd, void *info, __u32 *info_len)
func Bpf_obj_get_info_by_fd(bpf_fd int32, info unsafe.Pointer, info_len *uint32) int32 {
	return C.bpf_obj_get_info_by_fd()
}

// Bpf_prog_query is a wrapper of bpf_prog_query
//     int bpf_prog_query(int target_fd, enum bpf_attach_type type, __u32 query_flags, __u32 *attach_flags, __u32 *prog_ids, __u32 *prog_cnt)
func Bpf_prog_query(target_fd int32, type_ Enum_bpf_attach_type, query_flags uint32, attach_flags *uint32, prog_ids *uint32, prog_cnt *uint32) int32 {
	return C.bpf_prog_query()
}

// Bpf_raw_tracepoint_open is a wrapper of bpf_raw_tracepoint_open
//     int bpf_raw_tracepoint_open(const char *name, int prog_fd)
func Bpf_raw_tracepoint_open(name string, prog_fd int32) int32 {
	return C.bpf_raw_tracepoint_open()
}

// Bpf_load_btf is a wrapper of bpf_load_btf
//     int bpf_load_btf(const void *btf, __u32 btf_size, char *log_buf, __u32 log_buf_size, bool do_log)
func Bpf_load_btf(btf unsafe.Pointer, btf_size uint32, log_buf *C.char, log_buf_size uint32, do_log bool) int32 {
	return C.bpf_load_btf()
}

// Bpf_task_fd_query is a wrapper of bpf_task_fd_query
//     int bpf_task_fd_query(int pid, int fd, __u32 flags, char *buf, __u32 *buf_len, __u32 *prog_id, __u32 *fd_type, __u64 *probe_offset, __u64 *probe_addr)
func Bpf_task_fd_query(pid int32, fd int32, flags uint32, buf *C.char, buf_len *uint32, prog_id *uint32, fd_type *uint32, probe_offset *uint64, probe_addr *uint64) int32 {
	return C.bpf_task_fd_query()
}

// Bpf_enable_stats is a wrapper of bpf_enable_stats
//     int bpf_enable_stats(enum bpf_stats_type type)
func Bpf_enable_stats(type_ Enum_bpf_stats_type) int32 {
	return C.bpf_enable_stats()
}

// Bpf_prog_bind_map is a wrapper of bpf_prog_bind_map
//     int bpf_prog_bind_map(int prog_fd, int map_fd, const struct bpf_prog_bind_opts *opts)
func Bpf_prog_bind_map(prog_fd int32, map_fd int32, opts *Struct_bpf_prog_bind_opts) int32 {
	return C.bpf_prog_bind_map()
}

// Bpf_prog_test_run_opts is a wrapper of bpf_prog_test_run_opts
//     int bpf_prog_test_run_opts(int prog_fd, struct bpf_test_run_opts *opts)
func Bpf_prog_test_run_opts(prog_fd int32, opts *Struct_bpf_test_run_opts) int32 {
	return C.bpf_prog_test_run_opts()
}
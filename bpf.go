package libbpf

// #include <bpf/bpf.h>
import "C"


	LIBBPF_API int
bpf_create_map_xattr(const struct bpf_create_map_attr *create_attr);

	LIBBPF_API int bpf_create_map_node(enum bpf_map_type map_type, const char *name,
				   int key_size, int value_size,
				   int max_entries, __u32 map_flags, int node);

	LIBBPF_API int bpf_create_map_name(enum bpf_map_type map_type, const char *name,
				   int key_size, int value_size,
				   int max_entries, __u32 map_flags);

	LIBBPF_API int bpf_create_map(enum bpf_map_type map_type, int key_size,
			      int value_size, int max_entries, __u32 map_flags);

	LIBBPF_API int bpf_create_map_in_map_node(enum bpf_map_type map_type,
					  const char *name, int key_size,
					  int inner_map_fd, int max_entries,
					  __u32 map_flags, int node);

	LIBBPF_API int bpf_create_map_in_map(enum bpf_map_type map_type,
				     const char *name, int key_size,
				     int inner_map_fd, int max_entries,
				     __u32 map_flags);

	LIBBPF_API int
bpf_load_program_xattr(const struct bpf_load_program_attr *load_attr,
		       char *log_buf, size_t log_buf_sz);

	LIBBPF_API int bpf_load_program(enum bpf_prog_type type,
				const struct bpf_insn *insns, size_t insns_cnt,
				const char *license, __u32 kern_version,
				char *log_buf, size_t log_buf_sz);

	LIBBPF_API int bpf_verify_program(enum bpf_prog_type type,
				  const struct bpf_insn *insns,
				  size_t insns_cnt, __u32 prog_flags,
				  const char *license, __u32 kern_version,
				  char *log_buf, size_t log_buf_sz,
				  int log_level);

	LIBBPF_API int bpf_map_update_elem(int fd, const void *key, const void *value,
				   __u64 flags);

	LIBBPF_API int bpf_map_lookup_elem(int fd, const void *key, void *value);

	LIBBPF_API int bpf_map_lookup_elem_flags(int fd, const void *key, void *value,
					 __u64 flags);

	LIBBPF_API int bpf_map_lookup_and_delete_elem(int fd, const void *key,
					      void *value);

	LIBBPF_API int bpf_map_delete_elem(int fd, const void *key);

	LIBBPF_API int bpf_map_get_next_key(int fd, const void *key, void *next_key);

	LIBBPF_API int bpf_map_freeze(int fd);

	LIBBPF_API int bpf_map_delete_batch(int fd, void *keys,
				    __u32 *count,
				    const struct bpf_map_batch_opts *opts);

	LIBBPF_API int bpf_map_lookup_batch(int fd, void *in_batch, void *out_batch,
				    void *keys, void *values, __u32 *count,
				    const struct bpf_map_batch_opts *opts);

	LIBBPF_API int bpf_map_lookup_and_delete_batch(int fd, void *in_batch,
					void *out_batch, void *keys,
					void *values, __u32 *count,
					const struct bpf_map_batch_opts *opts);

	LIBBPF_API int bpf_map_update_batch(int fd, void *keys, void *values,
				    __u32 *count,
				    const struct bpf_map_batch_opts *opts);

	LIBBPF_API int bpf_obj_pin(int fd, const char *pathname);

	LIBBPF_API int bpf_obj_get(const char *pathname);

	LIBBPF_API int bpf_prog_attach(int prog_fd, int attachable_fd,
			       enum bpf_attach_type type, unsigned int flags);

	LIBBPF_API int bpf_prog_attach_xattr(int prog_fd, int attachable_fd,
				     enum bpf_attach_type type,
				     const struct bpf_prog_attach_opts *opts);

	LIBBPF_API int bpf_prog_detach(int attachable_fd, enum bpf_attach_type type);

	LIBBPF_API int bpf_prog_detach2(int prog_fd, int attachable_fd,
				enum bpf_attach_type type);

	LIBBPF_API int bpf_link_create(int prog_fd, int target_fd,
			       enum bpf_attach_type attach_type,
			       const struct bpf_link_create_opts *opts);

	LIBBPF_API int bpf_link_detach(int link_fd);

	LIBBPF_API int bpf_link_update(int link_fd, int new_prog_fd,
			       const struct bpf_link_update_opts *opts);

	LIBBPF_API int bpf_iter_create(int link_fd);

	LIBBPF_API int bpf_prog_test_run_xattr(struct bpf_prog_test_run_attr *test_attr);

	LIBBPF_API int bpf_prog_test_run(int prog_fd, int repeat, void *data,
				 __u32 size, void *data_out, __u32 *size_out,
				 __u32 *retval, __u32 *duration);

	LIBBPF_API int bpf_prog_get_next_id(__u32 start_id, __u32 *next_id);

	LIBBPF_API int bpf_map_get_next_id(__u32 start_id, __u32 *next_id);

	LIBBPF_API int bpf_btf_get_next_id(__u32 start_id, __u32 *next_id);

	LIBBPF_API int bpf_link_get_next_id(__u32 start_id, __u32 *next_id);

	LIBBPF_API int bpf_prog_get_fd_by_id(__u32 id);

	LIBBPF_API int bpf_map_get_fd_by_id(__u32 id);

	LIBBPF_API int bpf_btf_get_fd_by_id(__u32 id);

	LIBBPF_API int bpf_link_get_fd_by_id(__u32 id);

	LIBBPF_API int bpf_obj_get_info_by_fd(int bpf_fd, void *info, __u32 *info_len);

	LIBBPF_API int bpf_prog_query(int target_fd, enum bpf_attach_type type,
			      __u32 query_flags, __u32 *attach_flags,
			      __u32 *prog_ids, __u32 *prog_cnt);

	LIBBPF_API int bpf_raw_tracepoint_open(const char *name, int prog_fd);

	LIBBPF_API int bpf_load_btf(const void *btf, __u32 btf_size, char *log_buf,
			    __u32 log_buf_size, bool do_log);

	LIBBPF_API int bpf_task_fd_query(int pid, int fd, __u32 flags, char *buf,
				 __u32 *buf_len, __u32 *prog_id, __u32 *fd_type,
				 __u64 *probe_offset, __u64 *probe_addr);

	LIBBPF_API int bpf_enable_stats(enum bpf_stats_type type);

	LIBBPF_API int bpf_prog_bind_map(int prog_fd, int map_fd,
				 const struct bpf_prog_bind_opts *opts);

	LIBBPF_API int bpf_prog_test_run_opts(int prog_fd,
				      struct bpf_test_run_opts *opts);


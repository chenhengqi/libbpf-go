package libbpf

// #include <bpf/libbpf.h>
import "C"


	LIBBPF_API int libbpf_strerror(int err, char *buf, size_t size);

	LIBBPF_API libbpf_print_fn_t libbpf_set_print(libbpf_print_fn_t fn);

	LIBBPF_API struct bpf_object *bpf_object__open(const char *path);

	LIBBPF_API struct bpf_object *
bpf_object__open_file(const char *path, const struct bpf_object_open_opts *opts);

	LIBBPF_API struct bpf_object *
bpf_object__open_mem(const void *obj_buf, size_t obj_buf_sz,
		     const struct bpf_object_open_opts *opts);

	LIBBPF_API struct bpf_object *
bpf_object__open_buffer(const void *obj_buf, size_t obj_buf_sz,
			const char *name);

	LIBBPF_API struct bpf_object *
bpf_object__open_xattr(struct bpf_object_open_attr *attr);

	LIBBPF_API int bpf_object__pin_maps(struct bpf_object *obj, const char *path);

	LIBBPF_API int bpf_object__unpin_maps(struct bpf_object *obj,
				      const char *path);

	LIBBPF_API int bpf_object__pin_programs(struct bpf_object *obj,
					const char *path);

	LIBBPF_API int bpf_object__unpin_programs(struct bpf_object *obj,
					  const char *path);

	LIBBPF_API int bpf_object__pin(struct bpf_object *object, const char *path);

	LIBBPF_API void bpf_object__close(struct bpf_object *object);

	LIBBPF_API int bpf_object__load(struct bpf_object *obj);

	LIBBPF_API int bpf_object__load_xattr(struct bpf_object_load_attr *attr);

	LIBBPF_API int bpf_object__unload(struct bpf_object *obj);

	LIBBPF_API const char *bpf_object__name(const struct bpf_object *obj);

	LIBBPF_API unsigned int bpf_object__kversion(const struct bpf_object *obj);

	LIBBPF_API int bpf_object__set_kversion(struct bpf_object *obj, __u32 kern_version);

	LIBBPF_API struct btf *bpf_object__btf(const struct bpf_object *obj);

	LIBBPF_API int bpf_object__btf_fd(const struct bpf_object *obj);

	LIBBPF_API struct bpf_program *
bpf_object__find_program_by_title(const struct bpf_object *obj,
				  const char *title);

	LIBBPF_API struct bpf_program *
bpf_object__find_program_by_name(const struct bpf_object *obj,
				 const char *name);

	LIBBPF_API struct bpf_object *bpf_object__next(struct bpf_object *prev);

	LIBBPF_API int bpf_object__set_priv(struct bpf_object *obj, void *priv,
				    bpf_object_clear_priv_t clear_priv);

	LIBBPF_API void *bpf_object__priv(const struct bpf_object *prog);

	LIBBPF_API int
libbpf_prog_type_by_name(const char *name, enum bpf_prog_type *prog_type,
			 enum bpf_attach_type *expected_attach_type);

	LIBBPF_API int libbpf_attach_type_by_name(const char *name,
					  enum bpf_attach_type *attach_type);

	LIBBPF_API int libbpf_find_vmlinux_btf_id(const char *name,
					  enum bpf_attach_type attach_type);

	LIBBPF_API struct bpf_program *bpf_program__next(struct bpf_program *prog,
						 const struct bpf_object *obj);

	LIBBPF_API struct bpf_program *bpf_program__prev(struct bpf_program *prog,
						 const struct bpf_object *obj);

	LIBBPF_API int bpf_program__set_priv(struct bpf_program *prog, void *priv,
				     bpf_program_clear_priv_t clear_priv);

	LIBBPF_API void *bpf_program__priv(const struct bpf_program *prog);

	LIBBPF_API void bpf_program__set_ifindex(struct bpf_program *prog,
					 __u32 ifindex);

	LIBBPF_API const char *bpf_program__name(const struct bpf_program *prog);

	LIBBPF_API const char *bpf_program__section_name(const struct bpf_program *prog);

	LIBBPF_API LIBBPF_DEPRECATED("BPF program title is confusing term;

	LIBBPF_API bool bpf_program__autoload(const struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_autoload(struct bpf_program *prog, bool autoload);

	LIBBPF_API size_t bpf_program__size(const struct bpf_program *prog);

	LIBBPF_API int bpf_program__load(struct bpf_program *prog, char *license,
				 __u32 kern_version);

	LIBBPF_API int bpf_program__fd(const struct bpf_program *prog);

	LIBBPF_API int bpf_program__pin_instance(struct bpf_program *prog,
					 const char *path,
					 int instance);

	LIBBPF_API int bpf_program__unpin_instance(struct bpf_program *prog,
					   const char *path,
					   int instance);

	LIBBPF_API int bpf_program__pin(struct bpf_program *prog, const char *path);

	LIBBPF_API int bpf_program__unpin(struct bpf_program *prog, const char *path);

	LIBBPF_API void bpf_program__unload(struct bpf_program *prog);

	LIBBPF_API struct bpf_link *bpf_link__open(const char *path);

	LIBBPF_API int bpf_link__fd(const struct bpf_link *link);

	LIBBPF_API const char *bpf_link__pin_path(const struct bpf_link *link);

	LIBBPF_API int bpf_link__pin(struct bpf_link *link, const char *path);

	LIBBPF_API int bpf_link__unpin(struct bpf_link *link);

	LIBBPF_API int bpf_link__update_program(struct bpf_link *link,
					struct bpf_program *prog);

	LIBBPF_API void bpf_link__disconnect(struct bpf_link *link);

	LIBBPF_API int bpf_link__detach(struct bpf_link *link);

	LIBBPF_API int bpf_link__destroy(struct bpf_link *link);

	LIBBPF_API struct bpf_link *
bpf_program__attach(struct bpf_program *prog);

	LIBBPF_API struct bpf_link *
bpf_program__attach_perf_event(struct bpf_program *prog, int pfd);

	LIBBPF_API struct bpf_link *
bpf_program__attach_kprobe(struct bpf_program *prog, bool retprobe,
			   const char *func_name);

	LIBBPF_API struct bpf_link *
bpf_program__attach_uprobe(struct bpf_program *prog, bool retprobe,
			   pid_t pid, const char *binary_path,
			   size_t func_offset);

	LIBBPF_API struct bpf_link *
bpf_program__attach_tracepoint(struct bpf_program *prog,
			       const char *tp_category,
			       const char *tp_name);

	LIBBPF_API struct bpf_link *
bpf_program__attach_raw_tracepoint(struct bpf_program *prog,
				   const char *tp_name);

	LIBBPF_API struct bpf_link *
bpf_program__attach_trace(struct bpf_program *prog);

	LIBBPF_API struct bpf_link *
bpf_program__attach_lsm(struct bpf_program *prog);

	LIBBPF_API struct bpf_link *
bpf_program__attach_cgroup(struct bpf_program *prog, int cgroup_fd);

	LIBBPF_API struct bpf_link *
bpf_program__attach_netns(struct bpf_program *prog, int netns_fd);

	LIBBPF_API struct bpf_link *
bpf_program__attach_xdp(struct bpf_program *prog, int ifindex);

	LIBBPF_API struct bpf_link *
bpf_program__attach_freplace(struct bpf_program *prog,
			     int target_fd, const char *attach_func_name);

	LIBBPF_API struct bpf_link *bpf_map__attach_struct_ops(struct bpf_map *map);

	LIBBPF_API struct bpf_link *
bpf_program__attach_iter(struct bpf_program *prog,
			 const struct bpf_iter_attach_opts *opts);

	LIBBPF_API int bpf_program__set_prep(struct bpf_program *prog, int nr_instance,
				     bpf_program_prep_t prep);

	LIBBPF_API int bpf_program__nth_fd(const struct bpf_program *prog, int n);

	LIBBPF_API int bpf_program__set_socket_filter(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_tracepoint(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_raw_tracepoint(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_kprobe(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_lsm(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_sched_cls(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_sched_act(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_xdp(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_perf_event(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_tracing(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_struct_ops(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_extension(struct bpf_program *prog);

	LIBBPF_API int bpf_program__set_sk_lookup(struct bpf_program *prog);

	LIBBPF_API enum bpf_prog_type bpf_program__get_type(const struct bpf_program *prog);

	LIBBPF_API void bpf_program__set_type(struct bpf_program *prog,
				      enum bpf_prog_type type);

	LIBBPF_API enum bpf_attach_type
bpf_program__get_expected_attach_type(const struct bpf_program *prog);

	LIBBPF_API void
bpf_program__set_expected_attach_type(struct bpf_program *prog,
				      enum bpf_attach_type type);

	LIBBPF_API int
bpf_program__set_attach_target(struct bpf_program *prog, int attach_prog_fd,
			       const char *attach_func_name);

	LIBBPF_API bool bpf_program__is_socket_filter(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_tracepoint(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_raw_tracepoint(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_kprobe(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_lsm(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_sched_cls(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_sched_act(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_xdp(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_perf_event(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_tracing(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_struct_ops(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_extension(const struct bpf_program *prog);

	LIBBPF_API bool bpf_program__is_sk_lookup(const struct bpf_program *prog);

	LIBBPF_API struct bpf_map *
bpf_object__find_map_by_name(const struct bpf_object *obj, const char *name);

	LIBBPF_API int
bpf_object__find_map_fd_by_name(const struct bpf_object *obj, const char *name);

	LIBBPF_API struct bpf_map *
bpf_object__find_map_by_offset(struct bpf_object *obj, size_t offset);

	LIBBPF_API struct bpf_map *
bpf_map__next(const struct bpf_map *map, const struct bpf_object *obj);

	LIBBPF_API struct bpf_map *
bpf_map__prev(const struct bpf_map *map, const struct bpf_object *obj);

	LIBBPF_API int bpf_map__fd(const struct bpf_map *map);

	LIBBPF_API int bpf_map__reuse_fd(struct bpf_map *map, int fd);

	LIBBPF_API const struct bpf_map_def *bpf_map__def(const struct bpf_map *map);

	LIBBPF_API const char *bpf_map__name(const struct bpf_map *map);

	LIBBPF_API enum bpf_map_type bpf_map__type(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_type(struct bpf_map *map, enum bpf_map_type type);

	LIBBPF_API __u32 bpf_map__max_entries(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_max_entries(struct bpf_map *map, __u32 max_entries);

	LIBBPF_API int bpf_map__resize(struct bpf_map *map, __u32 max_entries);

	LIBBPF_API __u32 bpf_map__map_flags(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_map_flags(struct bpf_map *map, __u32 flags);

	LIBBPF_API __u32 bpf_map__numa_node(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_numa_node(struct bpf_map *map, __u32 numa_node);

	LIBBPF_API __u32 bpf_map__key_size(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_key_size(struct bpf_map *map, __u32 size);

	LIBBPF_API __u32 bpf_map__value_size(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_value_size(struct bpf_map *map, __u32 size);

	LIBBPF_API __u32 bpf_map__btf_key_type_id(const struct bpf_map *map);

	LIBBPF_API __u32 bpf_map__btf_value_type_id(const struct bpf_map *map);

	LIBBPF_API __u32 bpf_map__ifindex(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_ifindex(struct bpf_map *map, __u32 ifindex);

	LIBBPF_API int bpf_map__set_priv(struct bpf_map *map, void *priv,
				 bpf_map_clear_priv_t clear_priv);

	LIBBPF_API void *bpf_map__priv(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_initial_value(struct bpf_map *map,
					  const void *data, size_t size);

	LIBBPF_API bool bpf_map__is_offload_neutral(const struct bpf_map *map);

	LIBBPF_API bool bpf_map__is_internal(const struct bpf_map *map);

	LIBBPF_API int bpf_map__set_pin_path(struct bpf_map *map, const char *path);

	LIBBPF_API const char *bpf_map__get_pin_path(const struct bpf_map *map);

	LIBBPF_API bool bpf_map__is_pinned(const struct bpf_map *map);

	LIBBPF_API int bpf_map__pin(struct bpf_map *map, const char *path);

	LIBBPF_API int bpf_map__unpin(struct bpf_map *map, const char *path);

	LIBBPF_API int bpf_map__set_inner_map_fd(struct bpf_map *map, int fd);

	LIBBPF_API long libbpf_get_error(const void *ptr);

	LIBBPF_API int bpf_prog_load_xattr(const struct bpf_prog_load_attr *attr,
				   struct bpf_object **pobj, int *prog_fd);

	LIBBPF_API int bpf_prog_load(const char *file, enum bpf_prog_type type,
			     struct bpf_object **pobj, int *prog_fd);

	LIBBPF_API int bpf_set_link_xdp_fd(int ifindex, int fd, __u32 flags);

	LIBBPF_API int bpf_set_link_xdp_fd_opts(int ifindex, int fd, __u32 flags,
					const struct bpf_xdp_set_link_opts *opts);

	LIBBPF_API int bpf_get_link_xdp_id(int ifindex, __u32 *prog_id, __u32 flags);

	LIBBPF_API int bpf_get_link_xdp_info(int ifindex, struct xdp_link_info *info,
				     size_t info_size, __u32 flags);

	LIBBPF_API struct ring_buffer *
ring_buffer__new(int map_fd, ring_buffer_sample_fn sample_cb, void *ctx,
		 const struct ring_buffer_opts *opts);

	LIBBPF_API void ring_buffer__free(struct ring_buffer *rb);

	LIBBPF_API int ring_buffer__add(struct ring_buffer *rb, int map_fd,
				ring_buffer_sample_fn sample_cb, void *ctx);

	LIBBPF_API int ring_buffer__poll(struct ring_buffer *rb, int timeout_ms);

	LIBBPF_API int ring_buffer__consume(struct ring_buffer *rb);

	LIBBPF_API int ring_buffer__epoll_fd(const struct ring_buffer *rb);

	LIBBPF_API struct perf_buffer *
perf_buffer__new(int map_fd, size_t page_cnt,
		 const struct perf_buffer_opts *opts);

	LIBBPF_API struct perf_buffer *
perf_buffer__new_raw(int map_fd, size_t page_cnt,
		     const struct perf_buffer_raw_opts *opts);

	LIBBPF_API void perf_buffer__free(struct perf_buffer *pb);

	LIBBPF_API int perf_buffer__epoll_fd(const struct perf_buffer *pb);

	LIBBPF_API int perf_buffer__poll(struct perf_buffer *pb, int timeout_ms);

	LIBBPF_API int perf_buffer__consume(struct perf_buffer *pb);

	LIBBPF_API int perf_buffer__consume_buffer(struct perf_buffer *pb, size_t buf_idx);

	LIBBPF_API size_t perf_buffer__buffer_cnt(const struct perf_buffer *pb);

	LIBBPF_API int perf_buffer__buffer_fd(const struct perf_buffer *pb, size_t buf_idx);

	LIBBPF_API enum bpf_perf_event_ret
bpf_perf_event_read_simple(void *mmap_mem, size_t mmap_size, size_t page_size,
			   void **copy_mem, size_t *copy_size,
			   bpf_perf_event_print_t fn, void *private_data);

	LIBBPF_API void bpf_prog_linfo__free(struct bpf_prog_linfo *prog_linfo);

	LIBBPF_API struct bpf_prog_linfo *
bpf_prog_linfo__new(const struct bpf_prog_info *info);

	LIBBPF_API const struct bpf_line_info *
bpf_prog_linfo__lfind_addr_func(const struct bpf_prog_linfo *prog_linfo,
				__u64 addr, __u32 func_idx, __u32 nr_skip);

	LIBBPF_API const struct bpf_line_info *
bpf_prog_linfo__lfind(const struct bpf_prog_linfo *prog_linfo,
		      __u32 insn_off, __u32 nr_skip);

	LIBBPF_API bool bpf_probe_prog_type(enum bpf_prog_type prog_type,
				    __u32 ifindex);

	LIBBPF_API bool bpf_probe_map_type(enum bpf_map_type map_type, __u32 ifindex);

	LIBBPF_API bool bpf_probe_helper(enum bpf_func_id id,
				 enum bpf_prog_type prog_type, __u32 ifindex);

	LIBBPF_API bool bpf_probe_large_insn_limit(__u32 ifindex);

	LIBBPF_API struct bpf_prog_info_linear *
bpf_program__get_prog_info_linear(int fd, __u64 arrays);

	LIBBPF_API void
bpf_program__bpil_addr_to_offs(struct bpf_prog_info_linear *info_linear);

	LIBBPF_API void
bpf_program__bpil_offs_to_addr(struct bpf_prog_info_linear *info_linear);

	LIBBPF_API int libbpf_num_possible_cpus(void);

	LIBBPF_API int
bpf_object__open_skeleton(struct bpf_object_skeleton *s,
			  const struct bpf_object_open_opts *opts);

	LIBBPF_API int bpf_object__load_skeleton(struct bpf_object_skeleton *s);

	LIBBPF_API int bpf_object__attach_skeleton(struct bpf_object_skeleton *s);

	LIBBPF_API void bpf_object__detach_skeleton(struct bpf_object_skeleton *s);

	LIBBPF_API void bpf_object__destroy_skeleton(struct bpf_object_skeleton *s);

	LIBBPF_API struct bpf_linker *bpf_linker__new(const char *filename, struct bpf_linker_opts *opts);

	LIBBPF_API int bpf_linker__add_file(struct bpf_linker *linker, const char *filename);

	LIBBPF_API int bpf_linker__finalize(struct bpf_linker *linker);

	LIBBPF_API void bpf_linker__free(struct bpf_linker *linker);


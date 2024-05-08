/*
 * Copyright (c) 2024, NVIDIA CORPORATION.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * NOTE: this file is generated by generate_ivf_flat.py
 *
 * Make changes there and run in this directory:
 *
 * > python generate_ivf_flat.py
 *
 */

#include <cuvs/neighbors/ivf_flat.hpp>

#include "ivf_flat_search.cuh"

namespace cuvs::neighbors::ivf_flat {

#define CUVS_INST_IVF_FLAT_SEARCH(T, IdxT)                                           \
  void search(raft::resources const& handle,                                         \
              const cuvs::neighbors::ivf_flat::search_params& params,                \
              cuvs::neighbors::ivf_flat::index<T, IdxT>& index,                      \
              raft::device_matrix_view<const T, IdxT, raft::row_major> queries,      \
              raft::device_matrix_view<IdxT, IdxT, raft::row_major> neighbors,       \
              raft::device_matrix_view<float, IdxT, raft::row_major> distances)      \
  {                                                                                  \
    cuvs::neighbors::ivf_flat::detail::search(                                       \
      handle, params, index, queries, neighbors, distances);                         \
  }                                                                                  \
  void search_with_filtering(                                                        \
    raft::resources const& handle,                                                   \
    const search_params& params,                                                     \
    index<T, IdxT>& idx,                                                             \
    raft::device_matrix_view<const T, IdxT, raft::row_major> queries,                \
    raft::device_matrix_view<IdxT, IdxT, raft::row_major> neighbors,                 \
    raft::device_matrix_view<float, IdxT, raft::row_major> distances,                \
    cuvs::neighbors::filtering::bitset_filter<uint32_t, IdxT> sample_filter)      \ 
     \
  {                                                                                  \
    cuvs::neighbors::ivf_flat::detail::search_with_filtering(                        \
      handle, params, idx, queries, neighbors, distances, sample_filter);            \
  }
CUVS_INST_IVF_FLAT_SEARCH(uint8_t, int64_t);

#undef CUVS_INST_IVF_FLAT_SEARCH

}  // namespace cuvs::neighbors::ivf_flat

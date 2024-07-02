package ivf_pq

import (
	"math/rand"
	"rapidsai/cuvs/cuvs/common"
	"testing"
	"time"
)

func TestIvfFlat(t *testing.T) {

	resource, _ := common.NewResource(nil)

	rand.Seed(time.Now().UnixNano())

	NDataPoints := 16
	NFeatures := 8

	TestDataset := make([][]float32, NDataPoints)
	for i := range TestDataset {
		TestDataset[i] = make([]float32, NFeatures)
		for j := range TestDataset[i] {
			TestDataset[i][j] = rand.Float32()
		}
	}

	dataset, _ := common.NewTensor(true, TestDataset)

	IndexParams, _ := CreateIndexParams(2, "L2Expanded", 2.0, 0, 0.5, true)

	index, _ := CreateIndex(IndexParams, &dataset)
	defer index.Close()
	// use the first 4 points from the dataset as queries : will test that we get them back
	// as their own nearest neighbor

	NQueries := 4
	K := 4
	queries, _ := common.NewTensor(true, TestDataset[:NQueries])
	NeighborsDataset := make([][]int64, NQueries)
	for i := range NeighborsDataset {
		NeighborsDataset[i] = make([]int64, K)
	}
	DistancesDataset := make([][]float32, NQueries)
	for i := range DistancesDataset {
		DistancesDataset[i] = make([]float32, K)
	}
	neighbors, _ := common.NewTensor(true, NeighborsDataset)
	distances, _ := common.NewTensor(true, DistancesDataset)

	_, todeviceerr := neighbors.ToDevice(&resource)
	if todeviceerr != nil {
		println(todeviceerr)
	}
	distances.ToDevice(&resource)
	dataset.ToDevice(&resource)

	err := BuildIndex(resource, IndexParams, &dataset, index)
	if err != nil {
		panic(err)
	}
	resource.Sync()

	queries.ToDevice(&resource)

	SearchParams, _ := CreateSearchParams(10)

	err = SearchIndex(resource, SearchParams, index, &queries, &neighbors, &distances)
	if err != nil {
		panic(err)
	}

	neighbors.ToHost(&resource)
	distances.ToHost(&resource)

	resource.Sync()

	// p := (*int64)(unsafe.Pointer(uintptr(neighbors.c_tensor.dl_tensor.data) + uintptr(K*8*3)))
	arr, _ := neighbors.GetArray()
	for i := range arr {
		println(arr[i][0])
		if arr[i][0] != int64(i) {
			t.Error("wrong neighbor, expected", i, "got", arr[i][0])
		}
	}

	arr_dist, _ := distances.GetArray()
	for i := range arr_dist {
		if arr_dist[i][0] >= float32(0.001) || arr_dist[i][0] <= float32(-0.001) {
			t.Error("wrong distance, expected", float32(i), "got", arr_dist[i][0])
		}
	}

}

func TestCompare(t *testing.T) {
	// rwbuilder 와 직접 생성한 rwSet 결과값이 동일한지 테스트
	rwSetBuilder := NewRWSetBuilder()

	// 팁이라면, 트랜잭션 rwSet 내부의 하나의 rwSet은 하나의 key에 대한 read, write 이다.
	// 따라서, 동일한 키 값으로 Add~ 함수를 여러번 선언할 경우 rwSet은 하나만 생성된다.

	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "mint", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "mint", version.NewHeight(2, 1))
	rwSetBuilder.AddToWriteSet("mdl", "mint", []byte("mdl-mint-500"))
	rwSetBuilder.AddToWriteSet("mdl", "mint", []byte("mdl-mint-600"))
	rwSetBuilder.AddToWriteSet("mdl", "mint", []byte("mdl-mint-500"))
	// rwSetBuilder.AddToWriteSet("mdl", "mint", []byte("mdl-mint-500"))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(2, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance2", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance2", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(1, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(2, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(2, 1))
	// rwSetBuilder.AddToReadSet("mdl", "getBalance", version.NewHeight(2, 1))

	// rwSetBuilder.AddToHashedReadSet("mdl", "coll2", "getBalance", version.NewHeight(1, 1))
	// rwSetBuilder.AddToWriteSet("mdl", "mint", []byte("mdl-mint-500"))
	// rwSetBuilder.AddToPvtAndHashedWriteSet("mdl", "coll1", "mint", []byte("pvt-mdl-coll1-mint-500"))

	// logger.Debugf("rwSet: \n %v", rwSetBuilder)

	actualSimRes, _ := rwSetBuilder.GetTxSimulationResults()

	// logger.Debugf("simulationResult: \n %v", actualSimRes)

	pubSet := rwSetBuilder.GetTxReadWriteSet()

	if pubSet != nil {
		pubDataProto, _ := pubSet.toProtoMsg()
		fmt.Println(pubDataProto)
	}

	// logger.Info(*a)
	// rwSetBuilder.GetTxSimulationResults()
	fmt.Println(actualSimRes.PubSimulationResults)

	// // // public test
	// // // KV-RWSet
	// // mdlKVRWSet := &kvrwset.KVRWSet{
	// // 	Reads: []*kvrwset.KVRead{NewKVRead("getBalance", version.NewHeight(1, 1))},
	// // 	// Reads: []*kvrwset.KVRead{NewKVRead("getBalance", version.NewHeight(1, 1)), NewKVRead("getName", version.NewHeight(1, 1))},
	// // 	Writes: []*kvrwset.KVWrite{newKVWrite("mint", []byte("mdl-mint-500"))},
	// // }

	// // mdlKVRWSet2 := &kvrwset.KVRWSet{
	// // 	Reads: []*kvrwset.KVRead{NewKVRead("getBalance", version.NewHeight(2, 1))},
	// // }

	// // // RWSet
	// // mdlRWSet := &rwset.NsReadWriteSet{
	// // 	Namespace: "mdl",
	// // 	Rwset:     serializeTestProtoMsg(t, mdlKVRWSet),
	// // }
	// // // RWSet
	// // mdlRWSet2 := &rwset.NsReadWriteSet{
	// // 	Namespace: "mdl",
	// // 	Rwset:     serializeTestProtoMsg(t, mdlKVRWSet2),
	// // }
	// // log.Println(mdlRWSet2)

	// // expectedTxRWSet := &rwset.TxReadWriteSet{NsRwset: []*rwset.NsReadWriteSet{mdlRWSet}}

	// // pvt test
	// pvtMdlColl1 := &kvrwset.KVRWSet{
	// 	Writes: []*kvrwset.KVWrite{newKVWrite("mint", []byte("pvt-mdl-coll1-mint-500"))},
	// 	// MetadataWrites: []*kvrwset.KVMetadataWrite{{Key: ""}},
	// }

	// // expectedPvtRWSet은 위에 생성하는 코드와 같음.
	// expectedPvtRWSet := &rwset.TxPvtReadWriteSet{
	// 	DataModel: rwset.TxReadWriteSet_KV,
	// 	NsPvtRwset: []*rwset.NsPvtReadWriteSet{
	// 		{
	// 			Namespace: "mdl",
	// 			CollectionPvtRwset: []*rwset.CollectionPvtReadWriteSet{
	// 				{
	// 					CollectionName: "coll1",
	// 					Rwset:          serializeTestProtoMsg(t, pvtMdlColl1),
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// hashedMdlColl1 := &kvrwset.HashedRWSet{
	// 	// HashedReads: []*kvrwset.KVReadHash{
	// 	// 	constructTestPvtKVReadHash(t, "getBalance", version.NewHeight(1, 1)),
	// 	// },
	// 	HashedWrites: []*kvrwset.KVWriteHash{
	// 		constructTestPvtKVWriteHash(t, "mint", []byte("pvt-mdl-coll1-mint-500")),
	// 	},
	// }

	// combinedMdl := &rwset.NsReadWriteSet{
	// 	Namespace: "mdl",
	// 	Rwset:     []byte(""),
	// 	// Rwset:     serializeTestProtoMsg(t, pubMdl),
	// 	CollectionHashedRwset: []*rwset.CollectionHashedReadWriteSet{
	// 		{
	// 			CollectionName: "coll1",
	// 			HashedRwset:    serializeTestProtoMsg(t, hashedMdlColl1),
	// 			PvtRwsetHash:   util.ComputeHash(serializeTestProtoMsg(t, pvtMdlColl1)),
	// 		},
	// 	},
	// }

	// expectedTxRWSet := &rwset.TxReadWriteSet{NsRwset: []*rwset.NsReadWriteSet{combinedMdl}}
	// fmt.Println(expectedTxRWSet)
	// fmt.Println(expectedPvtRWSet)
	// // require.Equal(t, actualSimRes.PubSimulationResults.NsRwset[0], combinedMdl)
	// // require.Equal(t, actualSimRes.PubSimulationResults, expectedTxRWSet)
	// // require.Equal(t, actualSimRes.PvtSimulationResults, expectedPvtRWSet)
}
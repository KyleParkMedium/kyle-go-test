func (s *GossipStateProviderImpl) deliverPayloadsPreprocessingEx() {
	for {
		select {
		// Wait for notification that next seq has arrived
		case <-s.payloads.Ready():
			s.logger.Debugf("[%s] Ready to transfer payloads (blocks) to the ledger, next block number is = [%d]", s.chainID, s.payloads.Next())
			// Collect all subsequent payloads
			for payload := s.payloads.Pop(); payload != nil; payload = s.payloads.Pop() {
				rawBlock := &common.Block{}
				if err := pb.Unmarshal(payload.Data, rawBlock); err != nil {
					s.logger.Errorf("Error getting block with seqNum = %d due to (%+v)...dropping block", payload.SeqNum, errors.WithStack(err))
					continue
				}

				if rawBlock.Data == nil || rawBlock.Header == nil {
					s.logger.Errorf("Block with claimed sequence %d has no header (%v) or data (%v)",
						payload.SeqNum, rawBlock.Header, rawBlock.Data)
					continue
				}
				s.logger.Debugf("[%s] Transferring block [%d] with %d transaction(s) to the ledger", s.chainID, payload.SeqNum, len(rawBlock.Data.Data))

				// Read all private data into slice
				var p util.PvtDataCollections

				if payload.PrivateData != nil {
					err := p.Unmarshal(payload.PrivateData)
					if err != nil {
						s.logger.Errorf("Wasn't able to unmarshal private data for block seqNum = %d due to (%+v)...dropping block", payload.SeqNum, errors.WithStack(err))
						continue
					}
				}

				blockEx := mdl.NewBlockEx(rawBlock)
				if blockEx == nil || blockEx.UnmarshaledBlock == nil || blockEx.ExtInfo == nil {
					s.logger.Errorf("Failed NewBlockEx. BlockEx : (%v), UnmarshaledBlock : (%v), ExtInfo : (%v)",
						blockEx, blockEx.UnmarshaledBlock, blockEx.ExtInfo)
					continue
				}

				blockAndPvtData, retrievedPvtdata, err := s.ledger.PreprocessingCommitEx(blockEx, p)
				if err != nil {
					s.logger.Errorf("Failed preprocessingCommit. (%+v)", err)
					continue
				}

				preprocessed := &PreprocessedPayload{
					SeqNum:           blockAndPvtData.Block.Header.Number,
					BlockAndPvtData:  blockAndPvtData,
					RetrievedPvtdata: retrievedPvtdata,
				}

				s.preprocessed.Push(preprocessed)
			}
		case <-s.stopCh:
			s.logger.Debug("State provider has been stopped, finishing to push new blocks.")
			return
		}
	}
}

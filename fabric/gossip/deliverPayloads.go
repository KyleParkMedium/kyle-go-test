func (s *GossipStateProviderImpl) deliverPayloads() {
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
				if err := s.commitBlock(rawBlock, p); err != nil {
					if executionErr, isExecutionErr := err.(*vsccErrors.VSCCExecutionFailureError); isExecutionErr {
						s.logger.Errorf("Failed executing VSCC due to %v. Aborting chain processing", executionErr)
						return
					}
					s.logger.Panicf("Cannot commit block to the ledger due to %+v", errors.WithStack(err))
				}
			}
		case <-s.stopCh:
			s.logger.Debug("State provider has been stopped, finishing to push new blocks.")
			return
		}
	}
}
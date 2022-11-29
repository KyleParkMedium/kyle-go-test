func (s *GossipStateProviderImpl) deliverPreprocessed() {
	for {
		select {
		// Wait for notification that next seq has arrived
		case <-s.preprocessed.Ready():
			s.logger.Debugf("[%s] Ready to transfer payloads (blocks) to the ledger, next block number is = [%d]", s.chainID, s.preprocessed.Next())
			// Collect all subsequent payloads
			for preprocessed := s.preprocessed.Pop(); preprocessed != nil; preprocessed = s.preprocessed.Pop() {
				if err := s.commitBlockPreprocessing(preprocessed); err != nil {
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
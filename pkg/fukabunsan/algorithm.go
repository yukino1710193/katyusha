package fukabunsan // 負荷分散 - ふかぶんさん - Load Balancing

func (lb *LoadBalancer) LBAlgorithm(lbRequest *LBRequest) string {
	return lbRequest.Targets[0]
}

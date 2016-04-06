package agent_GWC

//GWC->agent control msg
type agent_GWC_msg struct {
	msg_type	uint32
	msg_len		uint32
	UE_IP		uint32
	GW_TEID		uint32
	GW_dst_IP	uint32	
}


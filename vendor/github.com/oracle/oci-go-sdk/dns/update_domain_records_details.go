// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDomainRecordsDetails The representation of UpdateDomainRecordsDetails
type UpdateDomainRecordsDetails struct {
	Items []RecordDetails `mandatory:"false" json:"items"`
}

func (m UpdateDomainRecordsDetails) String() string {
	return common.PointerString(m)
}

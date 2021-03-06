//  Copyright (c) Cloud J Tech, Inc. All rights reserved.
//  Licensed under the GPLv3 License. See License.txt in the project root for license information.
package model

import (
	hclog "github.com/hashicorp/go-hclog"
	"idcos.io/cloud-act2/log"
)

func getLogger() hclog.Logger {
	return log.L().Named("model")
}

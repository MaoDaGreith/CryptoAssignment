package data

import "MaoDaGreith/CryptoAssignment/model"

var SubscribedAddresses = make(map[string]bool)
var Transactions = make(map[string][]model.Transaction)
var CurrentBlock int = 0

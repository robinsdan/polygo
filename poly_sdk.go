/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */
package poly_go_sdk

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/polynetwork/poly-go-sdk/client"
	"github.com/polynetwork/poly-go-sdk/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PolySdk is the main struct for user
type PolySdk struct {
	ChainId uint64
	client.ClientMgr
}

// NewPolySdk return OntologySdk.
func NewPolySdk() *PolySdk {
	polySdk := &PolySdk{}
	return polySdk
}

func (this *PolySdk) SetChainId(chainId uint64) *PolySdk {
	this.ChainId = chainId
	return this
}

// CreateWallet return a new wallet
func (this *PolySdk) CreateWallet(walletFile string) (*Wallet, error) {
	if utils.IsFileExist(walletFile) {
		return nil, fmt.Errorf("wallet:%s has already exist", walletFile)
	}
	return NewWallet(walletFile), nil
}

// OpenWallet return a wallet instance
func (this *PolySdk) OpenWallet(walletFile string) (*Wallet, error) {
	return OpenWallet(walletFile)
}

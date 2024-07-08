package cmd

import (
	"fmt"

	"github.com/dyaksa/encryption-pii/crypt"
)

type CMD struct {
	MACDerivableKeysetPath  *string `env:"MAC_DERIVABLE_KEYSET_PATH,expand" json:"mac_derivable_keyset_path"`
	AEADDerivableKeysetPath *string `env:"AEAD_DERIVABLE_KEYSET_PATH,expand" json:"aead_derivable_keyset_path"`

	getAEAD func() (*crypt.DerivableKeyset[crypt.PrimitiveAEAD], error)
	getBIDX func() (*crypt.DerivableKeyset[crypt.PrimitiveBIDX], error)
}

func New() (cmd *CMD, err error) {
	cmd = &CMD{}

	if err = cmd.initEnv(); err != nil {
		return nil, fmt.Errorf("fail to init env: %w", err)
	}

	cmd.initAEADDerivableKeySet()
	cmd.initMACDerivableKeySet()

	return cmd, nil
}

func (c *CMD) initEnv() error {
	return EnvLoader(c, OptionsEnv{DotEnv: true, Prefix: "ENCRYPT_"})
}

func (c *CMD) initAEADDerivableKeySet() {
	if c.AEADDerivableKeysetPath == nil {
		c.getAEAD = func() (*crypt.DerivableKeyset[crypt.PrimitiveAEAD], error) { return nil, nil }
	}

	a, err := crypt.NewInsecureCleartextDerivableKeyset(*c.AEADDerivableKeysetPath, crypt.NewPrimitiveAEAD)
	c.getAEAD = func() (*crypt.DerivableKeyset[crypt.PrimitiveAEAD], error) { return a, err }
}

func (c *CMD) initMACDerivableKeySet() {
	if c.MACDerivableKeysetPath == nil {
		c.getBIDX = func() (*crypt.DerivableKeyset[crypt.PrimitiveBIDX], error) { return nil, nil }
	}

	b, err := crypt.NewInsecureCleartextDerivableKeyset(*c.MACDerivableKeysetPath, crypt.NewPrimitiveBIDX)
	c.getBIDX = func() (*crypt.DerivableKeyset[crypt.PrimitiveBIDX], error) { return b, err }
}

func (c CMD) AEADDerivableKeyset() (*crypt.DerivableKeyset[crypt.PrimitiveAEAD], error) {
	return c.getAEAD()
}

func (c CMD) MACDerivableKeyset() (*crypt.DerivableKeyset[crypt.PrimitiveBIDX], error) {
	return c.getBIDX()
}
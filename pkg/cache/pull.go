package cache

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/notaryproject/notation-go/dir"
	"github.com/notaryproject/notation-go/registry"
	"github.com/notaryproject/notation/internal/osutil"
	"github.com/opencontainers/go-digest"
)

// PullSignature pulls the signature if not exists in the cache.
func PullSignature(ctx context.Context, sigRepo registry.SignatureRepository, manifestDigest, sigDigest digest.Digest) error {
	sigPath := dir.Path.CachedSignature(manifestDigest, sigDigest)
	if info, err := os.Stat(sigPath); err == nil {
		if info.IsDir() {
			return errors.New("found directory at the signature file path: " + sigPath)
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	// fetch remote if not cached
	sig, err := sigRepo.GetBlob(ctx, sigDigest)
	if err != nil {
		return fmt.Errorf("get signature failure: %v: %v", sigDigest, err)
	}
	if err := osutil.WriteFile(sigPath, sig); err != nil {
		return fmt.Errorf("fail to write signature: %v: %v", sigDigest, err)
	}
	return nil
}

package pb

import "github.com/moby/buildkit/util/apicaps"

var Caps apicaps.CapList

// Every backwards or forwards non-compatible change needs to add a new capability row.
// By default new capabilities should be experimental. After merge a capability is
// considered immutable. After a capability is marked stable it should not be disabled.

const (
	CapSourceImage            apicaps.CapID = "source.image"
	CapSourceImageResolveMode apicaps.CapID = "source.image.resolvemode"
	CapSourceImageLayerLimit  apicaps.CapID = "source.image.layerlimit"

	CapSourceLocal                apicaps.CapID = "source.local"
	CapSourceLocalUnique          apicaps.CapID = "source.local.unique"
	CapSourceLocalSessionID       apicaps.CapID = "source.local.sessionid"
	CapSourceLocalIncludePatterns apicaps.CapID = "source.local.includepatterns"
	CapSourceLocalFollowPaths     apicaps.CapID = "source.local.followpaths"
	CapSourceLocalExcludePatterns apicaps.CapID = "source.local.excludepatterns"
	CapSourceLocalSharedKeyHint   apicaps.CapID = "source.local.sharedkeyhint"
	CapSourceLocalDiffer          apicaps.CapID = "source.local.differ"
	CapSourceMetadataTransfer     apicaps.CapID = "source.local.metadatatransfer"

	CapSourceGit              apicaps.CapID = "source.git"
	CapSourceGitKeepDir       apicaps.CapID = "source.git.keepgitdir"
	CapSourceGitFullURL       apicaps.CapID = "source.git.fullurl"
	CapSourceGitHTTPAuth      apicaps.CapID = "source.git.httpauth"
	CapSourceGitKnownSSHHosts apicaps.CapID = "source.git.knownsshhosts"
	CapSourceGitMountSSHSock  apicaps.CapID = "source.git.mountsshsock"
	CapSourceGitSubdir        apicaps.CapID = "source.git.subdir"

	CapSourceHTTP         apicaps.CapID = "source.http"
	CapSourceHTTPAuth     apicaps.CapID = "source.http.auth"
	CapSourceHTTPChecksum apicaps.CapID = "source.http.checksum"
	CapSourceHTTPPerm     apicaps.CapID = "source.http.perm"
	// NOTE the historical typo
	CapSourceHTTPUIDGID apicaps.CapID = "soruce.http.uidgid"
	CapSourceHTTPHeader apicaps.CapID = "source.http.header"

	CapSourceOCILayout apicaps.CapID = "source.ocilayout"

	CapBuildOpLLBFileName apicaps.CapID = "source.buildop.llbfilename"

	CapExecMetaBase                      apicaps.CapID = "exec.meta.base"
	CapExecMetaCgroupParent              apicaps.CapID = "exec.meta.cgroup.parent"
	CapExecMetaNetwork                   apicaps.CapID = "exec.meta.network"
	CapExecMetaProxy                     apicaps.CapID = "exec.meta.proxyenv"
	CapExecMetaSecurity                  apicaps.CapID = "exec.meta.security"
	CapExecMetaSecurityDeviceWhitelistV1 apicaps.CapID = "exec.meta.security.devices.v1"
	CapExecMetaSetsDefaultPath           apicaps.CapID = "exec.meta.setsdefaultpath"
	CapExecMetaUlimit                    apicaps.CapID = "exec.meta.ulimit"
	CapExecMetaCDI                       apicaps.CapID = "exec.meta.cdi"
	CapExecMetaRemoveMountStubsRecursive apicaps.CapID = "exec.meta.removemountstubs.recursive"
	CapExecMountBind                     apicaps.CapID = "exec.mount.bind"
	CapExecMountBindReadWriteNoOutput    apicaps.CapID = "exec.mount.bind.readwrite-nooutput"
	CapExecMountCache                    apicaps.CapID = "exec.mount.cache"
	CapExecMountCacheSharing             apicaps.CapID = "exec.mount.cache.sharing"
	CapExecMountSelector                 apicaps.CapID = "exec.mount.selector"
	CapExecMountTmpfs                    apicaps.CapID = "exec.mount.tmpfs"
	CapExecMountTmpfsSize                apicaps.CapID = "exec.mount.tmpfs.size"
	CapExecMountSecret                   apicaps.CapID = "exec.mount.secret"
	CapExecMountSSH                      apicaps.CapID = "exec.mount.ssh"
	CapExecMountContentCache             apicaps.CapID = "exec.mount.cache.content"
	CapExecCgroupsMounted                apicaps.CapID = "exec.cgroup"
	CapExecSecretEnv                     apicaps.CapID = "exec.secretenv"
	CapExecValidExitCode                 apicaps.CapID = "exec.validexitcode"

	CapFileBase                               apicaps.CapID = "file.base"
	CapFileRmWildcard                         apicaps.CapID = "file.rm.wildcard"
	CapFileCopyIncludeExcludePatterns         apicaps.CapID = "file.copy.includeexcludepatterns"
	CapFileRmNoFollowSymlink                  apicaps.CapID = "file.rm.nofollowsymlink"
	CapFileCopyAlwaysReplaceExistingDestPaths apicaps.CapID = "file.copy.alwaysreplaceexistingdestpaths"
	CapFileCopyModeStringFormat               apicaps.CapID = "file.copy.modestring"
	CapFileSymlinkCreate                      apicaps.CapID = "file.symlink.create"

	CapConstraints apicaps.CapID = "constraints"
	CapPlatform    apicaps.CapID = "platform"

	CapMetaIgnoreCache apicaps.CapID = "meta.ignorecache"
	CapMetaDescription apicaps.CapID = "meta.description"
	CapMetaExportCache apicaps.CapID = "meta.exportcache"

	CapRemoteCacheGHA    apicaps.CapID = "cache.gha"
	CapRemoteCacheS3     apicaps.CapID = "cache.s3"
	CapRemoteCacheAzBlob apicaps.CapID = "cache.azblob"

	CapMergeOp apicaps.CapID = "mergeop"
	CapDiffOp  apicaps.CapID = "diffop"

	CapAnnotations  apicaps.CapID = "exporter.image.annotations"
	CapAttestations apicaps.CapID = "exporter.image.attestations"

	// CapSourceDateEpoch is the capability to automatically handle the date epoch
	CapSourceDateEpoch apicaps.CapID = "exporter.sourcedateepoch"

	CapMultipleExporters apicaps.CapID = "exporter.multiple"
	CapSessionExporter   apicaps.CapID = "exporter.session"

	CapSourcePolicy apicaps.CapID = "source.policy"

	// GC/Prune controls allow MinFreeSpace and MaxUsedSpace to be set
	CapGCFreeSpaceFilter apicaps.CapID = "gc.freespacefilter"

	// ListenBuildHistory requests support server-side filters
	CapHistoryFilters apicaps.CapID = "history.filter"
)

func init() {
	Caps.Init(apicaps.Cap{
		ID:      CapSourceImage,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceImageResolveMode,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceImageLayerLimit,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocal,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalUnique,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalSessionID,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalIncludePatterns,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalFollowPaths,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalExcludePatterns,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalSharedKeyHint,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceLocalDiffer,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceMetadataTransfer,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGit,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitKeepDir,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitFullURL,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitHTTPAuth,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitKnownSSHHosts,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitMountSSHSock,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceGitSubdir,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTP,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTPChecksum,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTPPerm,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTPAuth,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTPUIDGID,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceHTTPHeader,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceOCILayout,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapBuildOpLLBFileName,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaBase,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaCgroupParent,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaProxy,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaNetwork,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaSetsDefaultPath,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaSecurity,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaSecurityDeviceWhitelistV1,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaUlimit,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMetaCDI,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountBind,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountBindReadWriteNoOutput,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountCache,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountCacheSharing,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountSelector,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountTmpfs,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountTmpfsSize,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountSecret,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountSSH,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecMountContentCache,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecCgroupsMounted,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecSecretEnv,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapExecValidExitCode,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileBase,
		Enabled: true,
		Status:  apicaps.CapStatusPrerelease,
		SupportedHint: map[string]string{
			"docker":   "Docker v19.03",
			"buildkit": "BuildKit v0.5.0",
		},
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileRmWildcard,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileRmNoFollowSymlink,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileCopyIncludeExcludePatterns,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileCopyAlwaysReplaceExistingDestPaths,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapFileSymlinkCreate,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapConstraints,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapPlatform,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapMetaIgnoreCache,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapMetaDescription,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapMetaExportCache,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapRemoteCacheGHA,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapRemoteCacheS3,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapRemoteCacheAzBlob,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapMergeOp,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapDiffOp,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapAnnotations,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapAttestations,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourceDateEpoch,
		Name:    "source date epoch",
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapMultipleExporters,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSessionExporter,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapSourcePolicy,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapGCFreeSpaceFilter,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})

	Caps.Init(apicaps.Cap{
		ID:      CapHistoryFilters,
		Enabled: true,
		Status:  apicaps.CapStatusExperimental,
	})
}

// Copyright 2018 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package libgit

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/kbfs/libfs"
	"github.com/keybase/kbfs/libkbfs"
	"github.com/keybase/kbfs/tlf"
	"github.com/stretchr/testify/require"
	gogit "gopkg.in/src-d/go-git.v4"
)

func TestAutogitNodeWrappersNoRepos(t *testing.T) {
	ctx, config, cancel, tempdir := initConfigForAutogit(t)
	defer cancel()
	defer os.RemoveAll(tempdir)
	defer libkbfs.CheckConfigAndShutdown(ctx, t, config)

	shutdown := StartAutogit(config)
	defer shutdown()

	h, err := libkbfs.ParseTlfHandle(
		ctx, config.KBPKI(), config.MDOps(), "user1", tlf.Private)
	require.NoError(t, err)
	rootFS, err := libfs.NewFS(
		ctx, config, h, libkbfs.MasterBranch, "", "", keybase1.MDPriorityNormal)
	require.NoError(t, err)

	t.Log("Looking at user1's autogit directory should fail if no git repos")
	_, err = rootFS.ReadDir(AutogitRoot)
	require.Error(t, err)
}

func checkAutogitOneFile(t *testing.T, rootFS *libfs.FS) {
	fis, err := rootFS.ReadDir(".kbfs_autogit/test")
	require.NoError(t, err)
	require.Len(t, fis, 1)
	f, err := rootFS.Open(".kbfs_autogit/test/foo")
	require.NoError(t, err)
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	require.Equal(t, "hello", string(data))
}

func checkAutogitTwoFiles(t *testing.T, rootFS *libfs.FS) {
	fis, err := rootFS.ReadDir(".kbfs_autogit/test")
	require.NoError(t, err)
	require.Len(t, fis, 2) // foo and foo2
	f, err := rootFS.Open(".kbfs_autogit/test/foo")
	require.NoError(t, err)
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	require.Equal(t, "hello", string(data))
	f2, err := rootFS.Open(".kbfs_autogit/test/foo2")
	require.NoError(t, err)
	defer f2.Close()
	data2, err := ioutil.ReadAll(f2)
	require.NoError(t, err)
	require.Equal(t, "hello2", string(data2))
}

func TestAutogitRepoNode(t *testing.T) {
	ctx, config, cancel, tempdir := initConfigForAutogit(t)
	defer cancel()
	defer os.RemoveAll(tempdir)
	defer libkbfs.CheckConfigAndShutdown(ctx, t, config)

	am := NewAutogitManager(config)
	defer am.Shutdown()
	rw := rootWrapper{am}
	config.AddRootNodeWrapper(rw.wrap)

	h, err := libkbfs.ParseTlfHandle(
		ctx, config.KBPKI(), config.MDOps(), "user1", tlf.Private)
	require.NoError(t, err)
	rootFS, err := libfs.NewFS(
		ctx, config, h, libkbfs.MasterBranch, "", "", keybase1.MDPriorityNormal)
	require.NoError(t, err)

	t.Log("Init a new repo directly into KBFS.")
	dotgitFS, _, err := GetOrCreateRepoAndID(ctx, config, h, "test", "")
	require.NoError(t, err)
	err = rootFS.MkdirAll("worktree", 0600)
	require.NoError(t, err)
	worktreeFS, err := rootFS.Chroot("worktree")
	require.NoError(t, err)
	dotgitStorage, err := NewGitConfigWithoutRemotesStorer(dotgitFS)
	require.NoError(t, err)
	repo, err := gogit.Init(dotgitStorage, worktreeFS)
	require.NoError(t, err)
	addFileToWorktreeAndCommit(
		t, ctx, config, h, repo, worktreeFS, "foo", "hello")

	t.Log("Use autogit to clone it using ReadDir")
	checkAutogitOneFile(t, rootFS)

	t.Log("Update the source repo and make sure the autogit repos update too")
	addFileToWorktreeAndCommit(
		t, ctx, config, h, repo, worktreeFS, "foo2", "hello2")

	t.Log("Force the source repo to update for the user")
	srcRootNode, _, err := config.KBFSOps().GetOrCreateRootNode(
		ctx, h, libkbfs.MasterBranch)
	require.NoError(t, err)
	err = config.KBFSOps().SyncFromServer(
		ctx, srcRootNode.GetFolderBranch(), nil)
	require.NoError(t, err)

	t.Log("Update the dest repo")
	dstRootNode, _, err := config.KBFSOps().GetOrCreateRootNode(
		ctx, h, libkbfs.MasterBranch)
	require.NoError(t, err)
	err = config.KBFSOps().SyncFromServer(
		ctx, dstRootNode.GetFolderBranch(), nil)
	require.NoError(t, err)

	checkAutogitTwoFiles(t, rootFS)

	t.Log("Switch to branch, check in more files.")
	wt, err := repo.Worktree()
	require.NoError(t, err)
	err = wt.Checkout(&gogit.CheckoutOptions{
		Branch: "refs/heads/dir/test-branch",
		Create: true,
	})
	require.NoError(t, err)
	addFileToWorktreeAndCommit(
		t, ctx, config, h, repo, worktreeFS, "foo3", "hello3")
	err = wt.Checkout(&gogit.CheckoutOptions{Branch: "refs/heads/master"})
	require.NoError(t, err)
	checkAutogitTwoFiles(t, rootFS)

	t.Logf("Check the third file that's only on the branch")
	f3, err := rootFS.Open(
		".kbfs_autogit/test/.kbfs_autogit_branch_dir/" +
			".kbfs_autogit_branch_test-branch/foo3")
	require.NoError(t, err)
	defer f3.Close()
	data3, err := ioutil.ReadAll(f3)
	require.NoError(t, err)
	require.Equal(t, "hello3", string(data3))

	t.Logf("Use colons instead of slashes in the branch name")
	f4, err := rootFS.Open(
		".kbfs_autogit/test/.kbfs_autogit_branch_dir:test-branch/foo3")
	require.NoError(t, err)
	defer f4.Close()
	data4, err := ioutil.ReadAll(f4)
	require.NoError(t, err)
	require.Equal(t, "hello3", string(data4))

	err = dotgitFS.SyncAll()
	require.NoError(t, err)
}

func TestAutogitRepoNodeReadonly(t *testing.T) {
	ctx, config, cancel, tempdir := initConfigForAutogit(t)
	defer cancel()
	defer os.RemoveAll(tempdir)
	defer libkbfs.CheckConfigAndShutdown(ctx, t, config)

	am := NewAutogitManager(config)
	defer am.Shutdown()
	rw := rootWrapper{am}
	config.AddRootNodeWrapper(rw.wrap)

	h, err := libkbfs.ParseTlfHandle(
		ctx, config.KBPKI(), config.MDOps(), "user1", tlf.Public)
	require.NoError(t, err)
	rootFS, err := libfs.NewFS(
		ctx, config, h, libkbfs.MasterBranch, "", "", keybase1.MDPriorityNormal)
	require.NoError(t, err)

	t.Log("Init a new repo directly into KBFS.")
	dotgitFS, _, err := GetOrCreateRepoAndID(ctx, config, h, "test", "")
	require.NoError(t, err)
	err = rootFS.MkdirAll("worktree", 0600)
	require.NoError(t, err)
	worktreeFS, err := rootFS.Chroot("worktree")
	require.NoError(t, err)
	dotgitStorage, err := NewGitConfigWithoutRemotesStorer(dotgitFS)
	require.NoError(t, err)
	repo, err := gogit.Init(dotgitStorage, worktreeFS)
	require.NoError(t, err)
	addFileToWorktreeAndCommit(
		t, ctx, config, h, repo, worktreeFS, "foo", "hello")

	t.Log("Use autogit to open it as another user.")
	config2 := libkbfs.ConfigAsUser(config, "user2")
	defer libkbfs.CheckConfigAndShutdown(ctx, t, config2)
	am2 := NewAutogitManager(config2)
	defer am2.Shutdown()
	rw2 := rootWrapper{am2}
	config2.AddRootNodeWrapper(rw2.wrap)
	rootFS2, err := libfs.NewFS(
		ctx, config2, h, libkbfs.MasterBranch, "", "",
		keybase1.MDPriorityNormal)
	require.NoError(t, err)
	checkAutogitOneFile(t, rootFS2)

	addFileToWorktree(t, repo, worktreeFS, "foo2", "hello2")
	t.Log("Repacking objects to more closely resemble a real kbfsgit push, " +
		"which only creates packfiles")
	err = repo.RepackObjects(&gogit.RepackConfig{})
	require.NoError(t, err)
	objFS, err := dotgitFS.Chroot("objects")
	require.NoError(t, err)
	fis, err := objFS.ReadDir("/")
	require.NoError(t, err)
	for _, fi := range fis {
		if fi.Name() != "pack" {
			err = recursiveDelete(ctx, objFS.(*libfs.FS), fi)
			require.NoError(t, err)
		}
	}
	t.Log("Repacking done")
	commitWorktree(t, ctx, config, h, worktreeFS)

	t.Log("Force the source repo to update for the second user")
	srcRootNode2, _, err := config2.KBFSOps().GetOrCreateRootNode(
		ctx, h, libkbfs.MasterBranch)
	require.NoError(t, err)
	err = config2.KBFSOps().SyncFromServer(
		ctx, srcRootNode2.GetFolderBranch(), nil)
	require.NoError(t, err)

	t.Log("Update the dest repo")
	dstRootNode2, _, err := config2.KBFSOps().GetOrCreateRootNode(
		ctx, h, libkbfs.MasterBranch)
	require.NoError(t, err)
	err = config2.KBFSOps().SyncFromServer(
		ctx, dstRootNode2.GetFolderBranch(), nil)
	require.NoError(t, err)

	checkAutogitTwoFiles(t, rootFS2)
}

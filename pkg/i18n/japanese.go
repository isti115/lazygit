package i18n

const japaneseIntroPopupMessage = `
Thanks for using lazygit! Seriously you rock. Three things to share with you:

 1) If you want to learn about lazygit's features, watch this vid:
      https://youtu.be/CPLdltN7wgE

 2) Be sure to read the latest release notes at:
      https://github.com/jesseduffield/lazygit/releases

 3) If you're using git, that makes you a programmer! With your help we can make
    lazygit better, so consider becoming a contributor and joining the fun at
      https://github.com/jesseduffield/lazygit
    You can also sponsor me and tell me what to work on by clicking the donate
    button at the bottom right.
    Or even just star the repo to share the love!
`

// exporting this so we can use it in tests
func japaneseTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:          "パネルの描画に十分な空間がありません",
		DiffTitle:               "差分",
		FilesTitle:              "ファイル",
		BranchesTitle:           "ブランチ",
		CommitsTitle:            "コミット",
		StashTitle:              "Stash",
		UnstagedChanges:         `ステージされていない変更`,
		StagedChanges:           `ステージされた変更`,
		MainTitle:               "メイン",
		MergeConfirmTitle:       "マージ",
		StagingTitle:            "メインパネル (Staging)",
		MergingTitle:            "メインパネル (Merging)",
		NormalTitle:             "メインパネル (Normal)",
		LogTitle:                "ログ",
		CommitMessage:           "コミットメッセージ",
		CredentialsUsername:     "ユーザ名",
		CredentialsPassword:     "パスワード",
		CredentialsPassphrase:   "SSH鍵のパスフレーズを入力",
		PassUnameWrong:          "パスワード, パスフレーズまたはユーザ名が間違っています。",
		CommitChanges:           "変更をコミット",
		AmendLastCommit:         "最新のコミットにamend",
		AmendLastCommitTitle:    "最新のコミットにamend",
		SureToAmend:             "最新のコミットに変更をamendします。よろしいですか? コミットメッセージはコミットパネルから変更できます。",
		NoCommitToAmend:         "amend可能なコミットが存在しません。",
		CommitChangesWithEditor: "gitエディタを使用して変更をコミット",
		StatusTitle:             "ステータス",
		LcNavigate:              "移動",
		LcMenu:                  "メニュー",
		LcExecute:               "実行",
		LcToggleStaged:          "ステージ/アンステージ",
		LcToggleStagedAll:       "すべての変更をステージ/アンステージ",
		LcToggleTreeView:        "ファイルツリーの表示を切り替え",
		LcOpenMergeTool:         "git mergetoolを開く",
		LcRefresh:               "リフレッシュ",
		LcPush:                  "push",
		LcPull:                  "pull",
		LcScroll:                "スクロール",
		MergeConflictsTitle:     "マージコンフリクト",
		LcCheckout:              "チェックアウト",
		LcFileFilter:            "ファイルをフィルタ (ステージ/アンステージ)",
		FilterStagedFiles:       "ステージされたファイルのみを表示",
		FilterUnstagedFiles:     "ステージされていないファイルのみを表示",
		ResetCommitFilterState:  "フィルタをリセット",
		// NoChangedFiles:                      "No changed files",
		PullWait:                "Pull中...",
		PushWait:                "Push中...",
		FetchWait:               "Fetch中...",
		LcSoftReset:             "softリセット",
		AlreadyCheckedOutBranch: "ブランチはすでにチェックアウトされています。",
		// SureForceCheckout:                   "Are you sure you want force checkout? You will lose all local changes",
		// ForceCheckoutBranch:                 "Force Checkout Branch",
		BranchName:               "ブランチ名",
		NewBranchNameBranchOff:   "新規ブランチ名 ('{{.branchName}}' に作成)",
		CantDeleteCheckOutBranch: "チェックアウト中のブランチは削除できません!",
		DeleteBranch:             "ブランチを削除",
		DeleteBranchMessage:      "ブランチ '{{.selectedBranchName}}' を削除します。よろしいですか?",
		ForceDeleteBranchMessage: "'{{.selectedBranchName}}' はマージされていません。本当に削除しますか?",
		// LcRebaseBranch:                      "rebase checked-out branch onto this branch",
		CantRebaseOntoSelf:        "ブランチを自分自身にリベースすることはできません。",
		CantMergeBranchIntoItself: "ブランチを自分自身にマージすることはできません。",
		// LcForceCheckout:                     "force checkout",
		// LcCheckoutByName:                    "checkout by name",
		LcNewBranch:             "新しいブランチを作成",
		LcDeleteBranch:          "ブランチを削除",
		NoBranchesThisRepo:      "リポジトリにブランチが存在しません",
		CommitMessageConfirm:    "{{.keyBindClose}}: 閉じる, {{.keyBindNewLine}}: 改行, {{.keyBindConfirm}}: 確定",
		CommitWithoutMessageErr: "コミットメッセージを入力してください",
		CloseConfirm:            "{{.keyBindClose}}: 閉じる/キャンセル, {{.keyBindConfirm}}: 確認",
		LcClose:                 "閉じる",
		LcQuit:                  "終了",
		// LcSquashDown:                        "squash down",
		// LcFixupCommit:                       "fixup commit",
		// NoCommitsThisBranch:                 "No commits for this branch",
		// YouNoCommitsToSquash:                "You have no commits to squash with",
		// Fixup:                               "Fixup",
		// SureFixupThisCommit:                 "Are you sure you want to 'fixup' this commit? It will be merged into the commit below",
		// SureSquashThisCommit:                "Are you sure you want to squash this commit into the commit below?",
		// Squash:                              "Squash",
		// LcPickCommit:                        "pick commit (when mid-rebase)",
		LcRevertCommit:       "コミットをrevert",
		LcRewordCommit:       "コミットメッセージを変更",
		LcDeleteCommit:       "コミットを削除",
		LcMoveDownCommit:     "コミットを1つ下に移動",
		LcMoveUpCommit:       "コミットを1つ上に移動",
		LcEditCommit:         "コミットを編集",
		LcAmendToCommit:      "ステージされた変更でamendコミット",
		LcRenameCommitEditor: "エディタでコミットメッセージを編集",
		Error:                "エラー",
		LcSelectHunk:         "hunkを選択",
		LcNavigateConflicts:  "コンフリクトを移動",
		// LcPickHunk:                          "pick hunk",
		// LcPickAllHunks:                      "pick all hunks",
		LcUndo:              "アンドゥ",
		LcUndoReflog:        "アンドゥ (via reflog) (experimental)",
		LcRedoReflog:        "リドゥ (via reflog) (experimental)",
		LcPop:               "pop",
		LcDrop:              "drop",
		LcApply:             "適用",
		NoStashEntries:      "Stashが存在しません",
		StashDrop:           "Stashを削除",
		SureDropStashEntry:  "Stashを削除します。よろしいですか?",
		StashPop:            "Stashをpop",
		SurePopStashEntry:   "Stashをpopします。よろしいですか?",
		StashApply:          "Stashを適用",
		SureApplyStashEntry: "Stashを適用します。よろしいですか?",
		// NoTrackedStagedFilesStash:           "You have no tracked/staged files to stash",
		StashChanges:      "変更をStash",
		LcRenameStash:     "Stashを変更",
		RenameStashPrompt: "Stash名を変更: {{.stashName}}",
		OpenConfig:        "設定ファイルを開く",
		EditConfig:        "設定ファイルを編集",
		ForcePush:         "Force push",
		ForcePushPrompt:   "ブランチがリモートブランチから分岐しています。'esc'でキャンセル, または'enter'でforce pushします。",
		ForcePushDisabled: "ブランチがリモートブランチから分岐しています。force pushは無効化されています。",
		// UpdatesRejectedAndForcePushDisabled: "Updates were rejected and you have disabled force pushing",
		LcCheckForUpdate:                 "更新を確認",
		CheckingForUpdates:               "更新を確認中...",
		UpdateAvailableTitle:             "最新リリース!",
		UpdateAvailable:                  "バージョン {{.newVersion}} をインストールしますか?",
		UpdateInProgressWaitingStatus:    "更新中",
		UpdateCompletedTitle:             "更新完了!",
		UpdateCompleted:                  "更新のインストールに成功しました。lazygitを再起動してください。",
		FailedToRetrieveLatestVersionErr: "バージョン情報の取得に失敗しました",
		OnLatestVersionErr:               "使用中のバージョンは最新です",
		MajorVersionErr:                  "新バージョン ({{.newVersion}}) は現在のバージョン ({{.currentVersion}}) と後方互換性がありません。",
		CouldNotFindBinaryErr:            "{{.url}} にバイナリが存在しませんでした。",
		UpdateFailedErr:                  "更新失敗: {{.errMessage}}",
		ConfirmQuitDuringUpdateTitle:     "現在更新中",
		ConfirmQuitDuringUpdate:          "現在更新を実行中です。終了しますか?",
		MergeToolTitle:                   "マージツール",
		MergeToolPrompt:                  "`git mergetool`を開きます。よろしいですか?",
		IntroPopupMessage:                japaneseIntroPopupMessage,
		// GitconfigParseErr:                   `Gogit failed to parse your gitconfig file due to the presence of unquoted '\' characters. Removing these should fix the issue.`,
		LcEditFile:               `ファイルを編集`,
		LcOpenFile:               `ファイルを開く`,
		LcIgnoreFile:             `.gitignoreに追加`,
		LcRefreshFiles:           `ファイルをリフレッシュ`,
		LcMergeIntoCurrentBranch: `現在のブランチにマージ`,
		ConfirmQuit:              `終了します。よろしいですか?`,
		SwitchRepo:               `最近使用したリポジトリに切り替え`,
		LcAllBranchesLogGraph:    `すべてのブランチログを表示`,
		UnsupportedGitService:    `サポートされていないGitサービスです。`,
		LcCreatePullRequest:      `Pull Requestを作成`,
		LcCopyPullRequestURL:     `Pull RequestのURLをクリップボードにコピー`,
		NoBranchOnRemote:         `ブランチがリモートに存在しません。リモートにpushしてください。`,
		LcFetch:                  `fetch`,
		// NoAutomaticGitFetchTitle:            `No automatic git fetch`,
		// NoAutomaticGitFetchBody:             `Lazygit can't use "git fetch" in a private repo; use 'f' in the files panel to run "git fetch" manually`,
		// FileEnter:                           `stage individual hunks/lines for file, or collapse/expand for directory`,
		// FileStagingRequirements:             `Can only stage individual lines for tracked files`,
		StageSelection:          `選択行をステージ/アンステージ`,
		ResetSelection:          `変更を削除 (git reset)`,
		ToggleDragSelect:        `範囲選択を切り替え`,
		ToggleSelectHunk:        `hunk選択を切り替え`,
		ToggleSelectionForPatch: `行をパッチに追加/削除`,
		ToggleStagingPanel:      `パネルを切り替え`,
		ReturnToFilesPanel:      `ファイル一覧に戻る`,
		// FastForward:                         `fast-forward this branch from its upstream`,
		// Fetching:                            "fetching and fast-forwarding {{.from}} -> {{.to}} ...",
		// FoundConflicts:                      "Conflicts! To abort press 'esc', otherwise press 'enter'",
		// FoundConflictsTitle:                 "Auto-merge failed",
		// PickHunk:                            "pick hunk",
		// PickAllHunks:                        "pick all hunks",
		// ViewMergeRebaseOptions:              "view merge/rebase options",
		// NotMergingOrRebasing:                "You are currently neither rebasing nor merging",
		RecentRepos: "最近使用したリポジトリ",
		// MergeOptionsTitle:                   "Merge Options",
		// RebaseOptionsTitle:                  "Rebase Options",
		CommitMessageTitle:  "コミットメッセージ",
		LocalBranchesTitle:  "ブランチ",
		SearchTitle:         "検索",
		TagsTitle:           "タグ",
		MenuTitle:           "メニュー",
		RemotesTitle:        "リモート",
		RemoteBranchesTitle: "リモートブランチ",
		PatchBuildingTitle:  "メインパネル (Patch Building)",
		InformationTitle:    "Information",
		SecondaryTitle:      "Secondary",
		ReflogCommitsTitle:  "参照ログ",
		GlobalTitle:         "グローバルキーバインド",
		// ConflictsResolved:                   "all merge conflicts resolved. Continue?",
		// RebasingTitle:                       "Rebasing",
		// ConfirmRebase:                       "Are you sure you want to rebase '{{.checkedOutBranch}}' onto '{{.selectedBranch}}'?",
		// ConfirmMerge:                        "Are you sure you want to merge '{{.selectedBranch}}' into '{{.checkedOutBranch}}'?",
		// FwdNoUpstream:                       "Cannot fast-forward a branch with no upstream",
		// FwdNoLocalUpstream:                  "Cannot fast-forward a branch whose remote is not registered locally",
		// FwdCommitsToPush:                    "Cannot fast-forward a branch with commits to push",
		ErrorOccurred: "エラーが発生しました! issueを作成してください: ",
		// NoRoom:                              "Not enough room",
		YouAreHere: "現在位置",
		// LcRewordNotSupported:                "rewording commits while interactively rebasing is not currently supported",
		LcCherryPickCopy:      "コミットをコピー (cherry-pick)",
		LcCherryPickCopyRange: "コミットを範囲コピー (cherry-pick)",
		LcPasteCommits:        "コミットを貼り付け (cherry-pick)",
		// SureCherryPick:                      "Are you sure you want to cherry-pick the copied commits onto this branch?",
		CherryPick: "Cherry-Pick",
		// CannotRebaseOntoFirstCommit:         "You cannot interactive rebase onto the first commit",
		// CannotSquashOntoSecondCommit:        "You cannot squash/fixup onto the second commit",
		Donate:                "支援",
		AskQuestion:           "質問",
		PrevLine:              "前の行を選択",
		NextLine:              "次の行を選択",
		PrevHunk:              "前のhunkを選択",
		NextHunk:              "次のhunkを選択",
		PrevConflict:          "前のコンフリクトを選択",
		NextConflict:          "次のコンフリクトを選択",
		SelectPrevHunk:        "前のhunkを選択",
		SelectNextHunk:        "次のhunkを選択",
		ScrollDown:            "下にスクロール",
		ScrollUp:              "上にスクロール",
		LcScrollUpMainPanel:   "メインパネルを上にスクロール",
		LcScrollDownMainPanel: "メインパネルを下にスクロール",
		AmendCommitTitle:      "amendコミット",
		AmendCommitPrompt:     "ステージされたファイルで現在のコミットをamendします。よろしいですか?",
		DeleteCommitTitle:     "コミットを削除",
		DeleteCommitPrompt:    "選択されたコミットを削除します。よろしいですか?",
		// SquashingStatus:                     "squashing",
		// FixingStatus:                        "fixing up",
		// DeletingStatus:                      "deleting",
		// MovingStatus:                        "moving",
		// RebasingStatus:                      "rebasing",
		// AmendingStatus:                      "amending",
		// CherryPickingStatus:                 "cherry-picking",
		// UndoingStatus:                       "undoing",
		// RedoingStatus:                       "redoing",
		// CheckingOutStatus:                   "checking out",
		// CommittingStatus:                    "committing",
		CommitFiles:                "Commit files",
		SubCommitsDynamicTitle:     "コミット (%s)",
		CommitFilesDynamicTitle:    "Diff files (%s)",
		RemoteBranchesDynamicTitle: "リモートブランチ (%s)",
		// LcViewItemFiles:                     "view selected item's files",
		CommitFilesTitle: "コミットファイル",
		// LcCheckoutCommitFile:                "checkout file",
		// LcDiscardOldFileChange:              "discard this commit's changes to this file",
		DiscardFileChangesTitle: "ファイルの変更を破棄",
		// DiscardFileChangesPrompt:            "Are you sure you want to discard this commit's changes to this file? If this file was created in this commit, it will be deleted",
		// DisabledForGPG:                      "Feature not available for users using GPG",
		CreateRepo: "Gitリポジトリではありません。リポジトリを作成しますか? (y/n): ",
		// AutoStashTitle:                      "Autostash?",
		// AutoStashPrompt:                     "You must stash and pop your changes to bring them across. Do this automatically? (enter/esc)",
		// StashPrefix:                         "Auto-stashing changes for ",
		// LcViewDiscardOptions:                "view 'discard changes' options",
		LcCancel:            "キャンセル",
		LcDiscardAllChanges: "すべての変更を破棄",
		// LcDiscardUnstagedChanges:            "discard unstaged changes",
		// LcDiscardAllChangesToAllFiles:       "nuke working tree",
		// LcDiscardAnyUnstagedChanges:         "discard unstaged changes",
		// LcDiscardUntrackedFiles:             "discard untracked files",
		LcHardReset: "hardリセット",
		// LcViewResetOptions:                  `view reset options`,
		LcCreateFixupCommit: `このコミットに対するfixupコミットを作成`,
		// LcSquashAboveCommits:                `squash all 'fixup!' commits above selected commit (autosquash)`,
		// SquashAboveCommits:                  `Squash all 'fixup!' commits above selected commit (autosquash)`,
		SureSquashAboveCommits:     `{{.commit}}に対するすべての fixup! コミットをsquashします。よろしいですか?`,
		CreateFixupCommit:          `fixupコミットを作成`,
		SureCreateFixupCommit:      `{{.commit}} に対する fixup! コミットを作成します。よろしいですか?`,
		LcExecuteCustomCommand:     "カスタムコマンドを実行",
		CustomCommand:              "カスタムコマンド:",
		LcCommitChangesWithoutHook: "pre-commitフックを実行せずに変更をコミット",
		// SkipHookPrefixNotConfigured:         "You have not configured a commit message prefix for skipping hooks. Set `git.skipHookPrefix = 'WIP'` in your config",
		// LcResetTo:                           `reset to`,
		PressEnterToReturn: "Enterを入力してください",
		// LcViewStashOptions:                  "view stash options",
		LcStashAllChanges: "変更をstash",
		// LcStashStagedChanges:                "stash staged changes",
		// LcStashOptions:                      "Stash options",
		// NotARepository:                      "Error: must be run inside a git repository",
		LcJump:            "パネルに移動",
		LcScrollLeftRight: "左右にスクロール",
		LcScrollLeft:      "左スクロール",
		LcScrollRight:     "右スクロール",
		DiscardPatch:      "パッチを破棄",
		// DiscardPatchConfirm:                 "You can only build a patch from one commit/stash-entry at a time. Discard current patch?",
		// CantPatchWhileRebasingError:         "You cannot build a patch or run patch commands while in a merging or rebasing state",
		// LcToggleAddToPatch:                  "toggle file included in patch",
		// LcToggleAllInPatch:                  "toggle all files included in patch",
		// LcUpdatingPatch:                     "updating patch",
		// ViewPatchOptions:                    "view custom patch options",
		// PatchOptionsTitle:                   "Patch Options",
		// NoPatchError:                        "No patch created yet. To start building a patch, use 'space' on a commit file or enter to add specific lines",
		// LcEnterFile:                         "enter file to add selected lines to the patch (or toggle directory collapsed)",
		// ExitCustomPatchBuilder:    ``,
		EnterUpstream:             `'<remote> <branchname>' の形式でupstreamを入力`,
		InvalidUpstream:           "upstreamの形式が正しくありません。'<remote> <branchname>' の形式で入力してください。",
		ReturnToRemotesList:       `リモート一覧に戻る`,
		LcAddNewRemote:            `リモートを新規追加`,
		LcNewRemoteName:           `新規リモート名:`,
		LcNewRemoteUrl:            `新規リモートURL:`,
		LcEditRemoteName:          `{{.remoteName}} の新しいリモート名を入力:`,
		LcEditRemoteUrl:           `{{.remoteName}} の新しいリモートURLを入力:`,
		LcRemoveRemote:            `リモートを削除`,
		LcRemoveRemotePrompt:      "リモートを削除します。よろしいですか?",
		DeleteRemoteBranch:        "リモートブランチを削除",
		DeleteRemoteBranchMessage: "リモートブランチを削除します。よろしいですか",
		// LcSetUpstream:                       "set as upstream of checked-out branch",
		// SetUpstreamTitle:                    "Set upstream branch",
		// SetUpstreamMessage:                  "Are you sure you want to set the upstream branch of '{{.checkedOut}}' to '{{.selected}}'",
		LcEditRemote:           "リモートを編集",
		LcTagCommit:            "タグを作成",
		TagMenuTitle:           "タグを作成",
		TagNameTitle:           "タグ名:",
		TagMessageTitle:        "タグメッセージ: ",
		LcAnnotatedTag:         "注釈付きタグ",
		LcLightweightTag:       "軽量タグ",
		LcDeleteTag:            "タグを削除",
		DeleteTagTitle:         "タグを削除",
		DeleteTagPrompt:        "タグ '{{.tagName}}' を削除します。よろしいですか?",
		PushTagTitle:           "リモートにタグ '{{.tagName}}' をpush",
		LcPushTag:              "タグをpush",
		LcCreateTag:            "タグを作成",
		CreateTagTitle:         "タグ名:",
		LcFetchRemote:          "リモートをfetch",
		FetchingRemoteStatus:   "リモートをfetch",
		LcCheckoutCommit:       "コミットをチェックアウト",
		SureCheckoutThisCommit: "選択されたコミットをチェックアウトします。よろしいですか?",
		// LcGitFlowOptions:                    "show git-flow options",
		// NotAGitFlowBranch:                   "This does not seem to be a git flow branch",
		// NewGitFlowBranchPrompt:              "new {{.branchType}} name:",
		// IgnoreTracked:                       "Ignore tracked file",
		// IgnoreTrackedPrompt:                 "Are you sure you want to ignore a tracked file?",
		// LcViewResetToUpstreamOptions:        "view upstream reset options",
		LcNextScreenMode:    "次のスクリーンモード (normal/half/fullscreen)",
		LcPrevScreenMode:    "前のスクリーンモード",
		LcStartSearch:       "検索を開始",
		Panel:               "パネル",
		Keybindings:         "キーバインド",
		LcRenameBranch:      "ブランチ名を変更",
		NewBranchNamePrompt: "新しいブランチ名を入力",
		// RenameBranchWarning:                 "This branch is tracking a remote. This action will only rename the local branch name, not the name of the remote branch. Continue?",
		LcOpenMenu: "メニューを開く",
		// LcResetCherryPick:                   "reset cherry-picked (copied) commits selection",
		LcNextTab:               "次のタブ",
		LcPrevTab:               "前のタブ",
		LcCantUndoWhileRebasing: "リベース中はアンドゥできません。",
		LcCantRedoWhileRebasing: "リベース中はリドゥできません。",
		// MustStashWarning:                    "Pulling a patch out into the index requires stashing and unstashing your changes. If something goes wrong, you'll be able to access your files from the stash. Continue?",
		// MustStashTitle:                      "Must stash",
		ConfirmationTitle: "確認パネル",
		LcPrevPage:        "前のページ",
		LcNextPage:        "次のページ",
		LcGotoTop:         "最上部までスクロール",
		LcGotoBottom:      "最下部までスクロール",
		// LcFilteringBy:                       "filtering by",
		// ResetInParentheses:                  "(reset)",
		// LcOpenFilteringMenu:                 "view filter-by-path options",
		// LcFilterBy:                          "filter by",
		// LcExitFilterMode:                    "stop filtering by path",
		// LcFilterPathOption:                  "enter path to filter by",
		// EnterFileName:                       "Enter path:",
		// FilteringMenuTitle:                  "Filtering",
		// MustExitFilterModeTitle:             "Command not available",
		// MustExitFilterModePrompt:            "Command not available in filtered mode. Exit filtered mode?",
		LcDiff: "差分",
		// LcEnterRefToDiff:                    "enter ref to diff",
		LcEnteRefName:    "参照を入力:",
		LcExitDiffMode:   "差分モードを終了",
		DiffingMenuTitle: "差分",
		// LcSwapDiff:                          "reverse diff direction",
		LcOpenDiffingMenu: "差分メニューを開く",
		// // the actual view is the extras view which I intend to give more tabs in future but for now we'll only mention the command log part
		LcOpenExtrasMenu: "コマンドログメニューを開く",
		// LcShowingGitDiff:                    "showing output for:",
		LcCommitDiff:                     "コミットの差分",
		LcCopyCommitShaToClipboard:       "コミットのSHAをクリップボードにコピー",
		LcCommitSha:                      "コミットのSHA",
		LcCommitURL:                      "コミットのURL",
		LcCopyCommitMessageToClipboard:   "コミットメッセージをクリップボードにコピー",
		LcCommitMessage:                  "コミットメッセージ",
		LcCommitAuthor:                   "コミットの作成者名",
		LcCopyCommitAttributeToClipboard: "コミットの情報をコピー",
		LcCopyBranchNameToClipboard:      "ブランチ名をクリップボードにコピー",
		LcCopyFileNameToClipboard:        "ファイル名をクリップボードにコピー",
		LcCopyCommitFileNameToClipboard:  "コミットされたファイル名をクリップボードにコピー",
		LcCopySelectedTexToClipboard:     "選択されたテキストをクリップボードにコピー",
		// LcCommitPrefixPatternError:          "Error in commitPrefix pattern",
		NoFilesStagedTitle:           "ファイルがステージされていません",
		NoFilesStagedPrompt:          "ファイルがステージされていません。すべての変更をコミットしますか?",
		BranchNotFoundTitle:          "ブランチが見つかりませんでした。",
		BranchNotFoundPrompt:         "ブランチが見つかりませんでした。新しくブランチを作成します ",
		UnstageLinesTitle:            "選択行をアンステージ",
		UnstageLinesPrompt:           "選択された行を削除 (git reset) します。よろしいですか? この操作は取り消せません。\nこの警告を無効化するには設定ファイルの 'gui.skipUnstageLineWarning' を true に設定してください。",
		LcCreateNewBranchFromCommit:  "コミットにブランチを作成",
		LcBuildingPatch:              "パッチを構築",
		LcViewCommits:                "コミットを閲覧",
		MinGitVersionError:           "lazygitの実行にはGit 2.0以降のバージョンが必要です。Gitを更新してください。もしくは、lazygitの後方互換性を改善するために https://github.com/jesseduffield/lazygit/issues にissueを作成してください。",
		LcRunningCustomCommandStatus: "カスタムコマンドを実行",
		// LcSubmoduleStashAndReset:            "stash uncommitted submodule changes and update",
		// LcAndResetSubmodules:                "and reset submodules",
		LcEnterSubmodule:               "サブモジュールを開く",
		LcCopySubmoduleNameToClipboard: "サブモジュール名をクリップボードにコピー",
		RemoveSubmodule:                "サブモジュールを削除",
		LcRemoveSubmodule:              "サブモジュールを削除",
		RemoveSubmodulePrompt:          "サブモジュール '%s' とそのディレクトリを削除します。よろしいですか? この操作は取り消せません。",
		LcResettingSubmoduleStatus:     "サブモジュールをリセット",
		LcNewSubmoduleName:             "新規サブモジュール名:",
		LcNewSubmoduleUrl:              "新規サブモジュールのURL:",
		LcNewSubmodulePath:             "新規サブモジュールのパス:",
		LcAddSubmodule:                 "サブモジュールを新規追加",
		LcAddingSubmoduleStatus:        "サブモジュールを新規追加",
		LcUpdateSubmoduleUrl:           "サブモジュール '%s' のURLを更新",
		LcUpdatingSubmoduleUrlStatus:   "URLを更新",
		LcEditSubmoduleUrl:             "サブモジュールのURLを更新",
		LcInitializingSubmoduleStatus:  "サブモジュールを初期化",
		LcInitSubmodule:                "サブモジュールを初期化",
		LcSubmoduleUpdate:              "サブモジュールを更新",
		LcUpdatingSubmoduleStatus:      "サブモジュールを更新",
		LcBulkInitSubmodules:           "サブモジュールを一括初期化",
		LcBulkUpdateSubmodules:         "サブモジュールを一括更新",
		// LcBulkDeinitSubmodules:              "bulk deinit submodules",
		// LcViewBulkSubmoduleOptions:          "view bulk submodule options",
		// LcBulkSubmoduleOptions:              "bulk submodule options",
		// LcRunningCommand:                    "running command",
		// SubCommitsTitle:                     "Sub-commits",
		SubmodulesTitle: "サブモジュール",
		NavigationTitle: "一覧パネルの操作",
		// SuggestionsCheatsheetTitle:          "Suggestions",
		// SuggestionsTitle:                    "Suggestions (press %s to focus)",
		ExtrasTitle: "コマンドログ",
		// PushingTagStatus:                    "pushing tag",
		PullRequestURLCopiedToClipboard:     "pull requestのURLがクリップボードにコピーされました",
		CommitDiffCopiedToClipboard:         "コミットの差分がクリップボードにコピーされました",
		CommitSHACopiedToClipboard:          "コミットのSHAがクリップボードにコピーされました",
		CommitURLCopiedToClipboard:          "コミットのURLがクリップボードにコピーされました",
		CommitMessageCopiedToClipboard:      "コミットメッセージがクリップボードにコピーされました",
		CommitAuthorCopiedToClipboard:       "コミットの作成者名がクリップボードにコピーされました",
		LcCopiedToClipboard:                 "クリップボードにコピーされました",
		ErrCannotEditDirectory:              "ディレクトリは編集できません。",
		ErrStageDirWithInlineMergeConflicts: "マージコンフリクトの発生したファイルを含むディレクトリはステージ/アンステージできません。マージコンフリクトを解決してください。",
		ErrRepositoryMovedOrDeleted:         "リポジトリが見つかりません。すでに削除されたか、移動された可能性があります ¯\\_(ツ)_/¯",
		CommandLog:                          "コマンドログ",
		ToggleShowCommandLog:                "コマンドログの表示/非表示を切り替え",
		FocusCommandLog:                     "コマンドログにフォーカス",
		CommandLogHeader:                    "コマンドログの表示/非表示は '%s' で切り替えられます。\n",
		RandomTip:                           "ランダムTips",
		// SelectParentCommitForMerge:          "Select parent commit for merge",
		ToggleWhitespaceInDiffView:   "空白文字の差分の表示有無を切り替え",
		IgnoringWhitespaceInDiffView: "空白文字の変更は差分画面に表示されません",
		ShowingWhitespaceInDiffView:  "空白文字の変更は差分画面に表示されます",
		// IncreaseContextInDiffView:           "Increase the size of the context shown around changes in the diff view",
		// DecreaseContextInDiffView:           "Decrease the size of the context shown around changes in the diff view",
		CreatePullRequest: "pull requestを作成",
		// CreatePullRequestOptions:            "Create pull request options",
		// LcCreatePullRequestOptions:          "create pull request options",
		LcDefaultBranch:      "デフォルトブランチ",
		LcSelectBranch:       "ブランチを選択",
		SelectConfigFile:     "設定ファイルを選択",
		NoConfigFileFoundErr: "設定ファイルが見つかりませんでした。",
		// LcLoadingFileSuggestions:            "loading file suggestions",
		// LcLoadingCommits:                    "loading commits",
		// MustSpecifyOriginError:              "Must specify a remote if specifying a branch",
		// GitOutput:                           "Git output:",
		// GitCommandFailed:                    "Git command failed. Check command log for details (open with %s)",
		AbortTitle:    "%sを中止",
		AbortPrompt:   "実施中の%sを中止します。よろしいですか?",
		LcOpenLogMenu: "ログメニューを開く",
		LogMenuTitle:  "コミットログオプション",
		// ToggleShowGitGraphAll:               "toggle show whole git graph (pass the `--all` flag to `git log`)",
		ShowGitGraph: "コミットグラフの表示",
		SortCommits:  "コミットの表示順",
		// CantChangeContextSizeError:          "Cannot change context while in patch building mode because we were too lazy to support it when releasing the feature. If you really want it, please let us know!",
		LcOpenCommitInBrowser: "ブラウザでコミットを開く",
		// LcViewBisectOptions:                 "view bisect options",
		// ConfirmRevertCommit:                 "Are you sure you want to revert {{.selectedCommit}}?",
		RewordInEditorTitle: "コミットメッセージをエディタで編集",
		// RewordInEditorPrompt:                "Are you sure you want to reword this commit in your editor?",
		// HardResetAutostashPrompt:            "Are you sure you want to hard reset to '%s'? An auto-stash will be performed if necessary.",
		// CheckoutPrompt:                      "Are you sure you want to checkout '%s'?",
		// UpstreamGone:                        "(upstream gone)",
		Actions: Actions{
			// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
			CheckoutCommit:      "コミットをチェックアウト",
			CheckoutTag:         "タグをチェックアウト",
			CheckoutBranch:      "ブランチをチェックアウト",
			ForceCheckoutBranch: "ブランチを強制的にチェックアウト",
			DeleteBranch:        "ブランチを削除",
			Merge:               "マージ",
			// RebaseBranch:                      "Rebase branch",
			RenameBranch: "ブランチ名を変更",
			CreateBranch: "ブランチを作成",
			// CherryPick:                        "(Cherry-pick) Paste commits",
			CheckoutFile: "ファイルをチェックアウトs",
			// DiscardOldFileChange:              "Discard old file change",
			// SquashCommitDown:                  "Squash commit down",
			FixupCommit:       "fixupコミット",
			RewordCommit:      "コミットメッセージを変更",
			DropCommit:        "コミットを削除",
			EditCommit:        "コミットを編集",
			AmendCommit:       "amendコミット",
			RevertCommit:      "コミットをrevert",
			CreateFixupCommit: "fixupコミットを作成",
			// SquashAllAboveFixupCommits:        "Squash all above fixup commits",
			CreateLightweightTag:              "軽量タグを作成",
			CreateAnnotatedTag:                "注釈付きタグを作成",
			CopyCommitMessageToClipboard:      "コミットメッセージをクリップボードにコピー",
			CopyCommitDiffToClipboard:         "コミットの差分をクリップボードにコピー",
			CopyCommitSHAToClipboard:          "コミットSHAをクリップボードにコピー",
			CopyCommitURLToClipboard:          "コミットのURLをクリップボードにコピー",
			CopyCommitAuthorToClipboard:       "コミットの作成者名をクリップボードにコピー",
			CopyCommitAttributeToClipboard:    "クリップボードにコピー",
			MoveCommitUp:                      "コミットを上に移動",
			MoveCommitDown:                    "コミットを下に移動",
			CustomCommand:                     "カスタムコマンド",
			DiscardAllChangesInDirectory:      "ディレクトリ内のすべての変更を破棄",
			DiscardUnstagedChangesInDirectory: "ディレクトリ内のすべてのステージされていない変更を破棄",
			DiscardAllChangesInFile:           "ファイル内のすべての変更を破棄",
			DiscardAllUnstagedChangesInFile:   "ファイル内のすべてのステージされていない変更を破棄",
			StageFile:                         "ファイルをステージ",
			StageResolvedFiles:                "マージコンフリクトが解決されたすべてのファイルをステージ",
			UnstageFile:                       "ファイルをアンステージ",
			UnstageAllFiles:                   "すべてのファイルをアンステージ",
			StageAllFiles:                     "すべてのファイルをステージ",
			LcIgnoreExcludeFile:               "ファイルをignore",
			Commit:                            "コミット",
			EditFile:                          "ファイルを編集",
			Push:                              "Push",
			Pull:                              "Pull",
			OpenFile:                          "ファイルを開く",
			StashAllChanges:                   "すべての変更をStash",
			StashStagedChanges:                "ステージされた変更をStash",
			GitFlowFinish:                     "Git flow finish",
			GitFlowStart:                      "Git Flow start",
			CopyToClipboard:                   "クリップボードにコピー",
			CopySelectedTextToClipboard:       "選択されたテキストをクリップボードにコピー",
			RemovePatchFromCommit:             "パッチをコミットから削除",
			MovePatchToSelectedCommit:         "パッチを選択したコミットに移動",
			MovePatchIntoIndex:                "パッチをindexに移動",
			MovePatchIntoNewCommit:            "パッチを次のコミットに移動",
			DeleteRemoteBranch:                "リモートブランチを削除",
			SetBranchUpstream:                 "upstreamブランチを設定",
			AddRemote:                         "リモートを追加",
			RemoveRemote:                      "リモートを削除",
			UpdateRemote:                      "リモートを更新",
			ApplyPatch:                        "パッチを適用",
			Stash:                             "Stash",
			RenameStash:                       "Stash名を変更",
			RemoveSubmodule:                   "サブモジュールを削除",
			ResetSubmodule:                    "サブモジュールをリセット",
			AddSubmodule:                      "サブモジュールを追加",
			UpdateSubmoduleUrl:                "サブモジュールのURLを更新",
			InitialiseSubmodule:               "サブモジュールを初期化",
			BulkInitialiseSubmodules:          "サブモジュールを一括初期化",
			BulkUpdateSubmodules:              "サブモジュールを一括更新",
			// BulkDeinitialiseSubmodules:        "Bulk deinitialise submodules",
			UpdateSubmodule: "サブモジュールを更新",
			DeleteTag:       "タグを削除",
			PushTag:         "タグをpush",
			// NukeWorkingTree:                   "Nuke working tree",
			// DiscardUnstagedFileChanges:        "Discard unstaged file changes",
			// RemoveUntrackedFiles:              "Remove untracked files",
			SoftReset:           "Softリセット",
			MixedReset:          "Mixedリセット",
			HardReset:           "Hardリセット",
			FastForwardBranch:   "ブランチをfast forward",
			Undo:                "アンドゥ",
			Redo:                "リドゥ",
			CopyPullRequestURL:  "pull requestのURLをコピー",
			OpenMergeTool:       "マージツールを開く",
			OpenCommitInBrowser: "コミットをブラウザで開く",
			OpenPullRequest:     "pull requestをブラウザで開く",
			StartBisect:         "bisectを開始",
			ResetBisect:         "bisectをリセット",
			BisectSkip:          "bisectをスキップ",
			BisectMark:          "bisectをマーク",
		},
		Bisect: Bisect{
			// Mark:                        "mark %s as %s",
			// MarkStart:                   "mark %s as %s (start bisect)",
			Skip:            "%s をスキップする",
			ResetTitle:      "'git bisect' をリセット",
			ResetPrompt:     "'git bisect' をリセットします。よろしいですか?",
			ResetOption:     "bisectをリセット",
			BisectMenuTitle: "bisect",
			CompleteTitle:   "Bisect完了",
			// CompletePrompt:              "Bisect complete! The following commit introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
			// CompletePromptIndeterminate: "Bisect complete! Some commits were skipped, so any of the following commits may have introduced the change:\n\n%s\n\nDo you want to reset 'git bisect' now?",
		},
	}
}

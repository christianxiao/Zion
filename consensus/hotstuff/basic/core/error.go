package core

import "errors"

var (
	// errInconsistentVote is returned when received subject is different from
	// current subject.
	errInconsistentVote = errors.New("inconsistent vote")

	errInconsistentPrepareQC = errors.New("inconsistent prepare qc")

	errInvalidDigest = errors.New("invalid digest")

	// errNotFromProposer is returned when received message is supposed to be from proposer.
	errNotFromProposer = errors.New("message does not come from proposer")

	errNotToProposer = errors.New("message does not send to proposer")

	// errFutureMessage is returned when current view is earlier than the
	// view of the received message.
	errFutureMessage = errors.New("future message")

	// errOldMessage is returned when the received message's view is earlier
	// than current view.
	errOldMessage = errors.New("old message")

	errOldVote = errors.New("old vote")

	// errInvalidMessage is returned when the message is malformed.
	errInvalidMessage = errors.New("invalid message")

	// errFailedDecodeNewView is returned when the NEWVIEW message is malformed.
	errFailedDecodeNewView = errors.New("failed to decode NEWVIEW")

	// errFailedDecodePrepare is returned when the PREPARE message is malformed.
	errFailedDecodePrepare     = errors.New("failed to decode PREPARE")
	errFailedDecodePrepareVote = errors.New("failed to decode PREPARE_VOTE")

	// errFailedDecodePreCommit is returned when the PRECOMMIT message is malformed.
	errFailedDecodePreCommit = errors.New("failed to decode PRECOMMIT")

	errFailedDecodePreCommitVote = errors.New("faild to decode PRECOMMIT_VOTE")

	// errFailedDecodeCommit is returned when the COMMIT message is malformed.
	errFailedDecodeCommit = errors.New("failed to decode COMMIT")

	errFailedDecodeDecide = errors.New("failed to decode DECIDE")

	errFailedDecodeCommitVote = errors.New("failed to decode COMMIT_VOTE")

	// errInvalidSigner is returned when the message is signed by a validator different than message sender
	errInvalidSigner = errors.New("message not signed by the sender")

	errState = errors.New("error state")

	errProposal = errors.New("proposal hash is not equal to qc digest")

	errMsgTypeInvalid = errors.New("message type invalid")

	errNoRequest = errors.New("no valid request")

	errInvalidProposal = errors.New("invalid proposal")

	errInvalidQCParticipant = errors.New("invalid qc participant")

	errVerifyUnsealedProposal = errors.New("verify unsealed proposal failed")
	errExtend                 = errors.New("proposal extend relationship error")
	errVerifyQC               = errors.New("verify qc error")
	errSafeNode               = errors.New("safeNode checking failed")
	errAddNewViews            = errors.New("add new view error")
	errAddPrepareVote         = errors.New("add prepare vote error")
	errAddPreCommitVote       = errors.New("add pre commit vote error")
	errAddCommitVote          = errors.New("add commit vote error")
)

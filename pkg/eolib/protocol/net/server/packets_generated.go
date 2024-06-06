package server

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	protocol "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
	net "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net"
	pub "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/pub"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// InitInitServerPacket ::  Reply to connection initialization and requests for unencrypted data. This packet is unencrypted.
type InitInitServerPacket struct {
	ReplyCode     InitReply
	ReplyCodeData InitInitReplyCodeData
}

type InitInitReplyCodeData interface {
	protocol.EoData
}

type InitInitReplyCodeDataOutOfDate struct {
	Version net.Version
}

func (s *InitInitReplyCodeDataOutOfDate) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Version : field : Version
	if err = s.Version.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataOutOfDate) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Version : field : Version
	if err = s.Version.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataOk struct {
	Seq1                     int
	Seq2                     int
	ServerEncryptionMultiple int
	ClientEncryptionMultiple int
	PlayerId                 int
	ChallengeResponse        int
}

func (s *InitInitReplyCodeDataOk) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Seq1 : field : byte
	if err = writer.AddByte(s.Seq1); err != nil {
		return
	}
	// Seq2 : field : byte
	if err = writer.AddByte(s.Seq2); err != nil {
		return
	}
	// ServerEncryptionMultiple : field : byte
	if err = writer.AddByte(s.ServerEncryptionMultiple); err != nil {
		return
	}
	// ClientEncryptionMultiple : field : byte
	if err = writer.AddByte(s.ClientEncryptionMultiple); err != nil {
		return
	}
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// ChallengeResponse : field : three
	if err = writer.AddThree(s.ChallengeResponse); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataOk) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Seq1 : field : byte
	s.Seq1 = int(reader.GetByte())
	// Seq2 : field : byte
	s.Seq2 = int(reader.GetByte())
	// ServerEncryptionMultiple : field : byte
	s.ServerEncryptionMultiple = int(reader.GetByte())
	// ClientEncryptionMultiple : field : byte
	s.ClientEncryptionMultiple = int(reader.GetByte())
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// ChallengeResponse : field : three
	s.ChallengeResponse = reader.GetThree()

	return
}

type InitInitReplyCodeDataBanned struct {
	BanType     InitBanType
	BanTypeData InitInitBanTypeData
}

type InitInitBanTypeData interface {
	protocol.EoData
}

// InitInitBanTypeData0 ::  The official client treats any value below 2 as a temporary ban. The official server sends 1, but some game server implementations. erroneously send 0.
type InitInitBanTypeData0 struct {
	MinutesRemaining int
}

func (s *InitInitBanTypeData0) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MinutesRemaining : field : byte
	if err = writer.AddByte(s.MinutesRemaining); err != nil {
		return
	}
	return
}

func (s *InitInitBanTypeData0) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MinutesRemaining : field : byte
	s.MinutesRemaining = int(reader.GetByte())

	return
}

type InitInitBanTypeDataTemporary struct {
	MinutesRemaining int
}

func (s *InitInitBanTypeDataTemporary) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MinutesRemaining : field : byte
	if err = writer.AddByte(s.MinutesRemaining); err != nil {
		return
	}
	return
}

func (s *InitInitBanTypeDataTemporary) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MinutesRemaining : field : byte
	s.MinutesRemaining = int(reader.GetByte())

	return
}

func (s *InitInitReplyCodeDataBanned) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BanType : field : InitBanType
	if err = writer.AddByte(int(s.BanType)); err != nil {
		return
	}
	switch s.BanType {
	case 0:
		switch s.BanTypeData.(type) {
		case *InitInitBanTypeData0:
			if err = s.BanTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.BanType)
			return
		}
	case InitBan_Temporary:
		switch s.BanTypeData.(type) {
		case *InitInitBanTypeDataTemporary:
			if err = s.BanTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.BanType)
			return
		}
	}
	return
}

func (s *InitInitReplyCodeDataBanned) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// BanType : field : InitBanType
	s.BanType = InitBanType(reader.GetByte())
	switch s.BanType {
	case 0:
		s.BanTypeData = &InitInitBanTypeData0{}
		if err = s.BanTypeData.Deserialize(reader); err != nil {
			return
		}
	case InitBan_Temporary:
		s.BanTypeData = &InitInitBanTypeDataTemporary{}
		if err = s.BanTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

type InitInitReplyCodeDataWarpMap struct {
	MapFile MapFile
}

func (s *InitInitReplyCodeDataWarpMap) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataWarpMap) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataFileEmf struct {
	MapFile MapFile
}

func (s *InitInitReplyCodeDataFileEmf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataFileEmf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataFileEif struct {
	PubFile PubFile
}

func (s *InitInitReplyCodeDataFileEif) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataFileEif) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataFileEnf struct {
	PubFile PubFile
}

func (s *InitInitReplyCodeDataFileEnf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataFileEnf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataFileEsf struct {
	PubFile PubFile
}

func (s *InitInitReplyCodeDataFileEsf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataFileEsf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataFileEcf struct {
	PubFile PubFile
}

func (s *InitInitReplyCodeDataFileEcf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataFileEcf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataMapMutation struct {
	MapFile MapFile
}

func (s *InitInitReplyCodeDataMapMutation) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *InitInitReplyCodeDataMapMutation) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

type InitInitReplyCodeDataPlayersList struct {
	PlayersList PlayersList
}

func (s *InitInitReplyCodeDataPlayersList) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *InitInitReplyCodeDataPlayersList) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

type InitInitReplyCodeDataPlayersListFriends struct {
	PlayersList PlayersListFriends
}

func (s *InitInitReplyCodeDataPlayersListFriends) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *InitInitReplyCodeDataPlayersListFriends) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

func (s InitInitServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Init
}

func (s InitInitServerPacket) Action() net.PacketAction {
	return net.PacketAction_Init
}

func (s *InitInitServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : InitReply
	if err = writer.AddByte(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case InitReply_OutOfDate:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataOutOfDate:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_Ok:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataOk:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_Banned:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataBanned:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_WarpMap:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataWarpMap:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_FileEmf:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataFileEmf:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_FileEif:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataFileEif:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_FileEnf:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataFileEnf:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_FileEsf:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataFileEsf:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_FileEcf:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataFileEcf:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_MapMutation:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataMapMutation:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_PlayersList:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataPlayersList:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case InitReply_PlayersListFriends:
		switch s.ReplyCodeData.(type) {
		case *InitInitReplyCodeDataPlayersListFriends:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *InitInitServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : InitReply
	s.ReplyCode = InitReply(reader.GetByte())
	switch s.ReplyCode {
	case InitReply_OutOfDate:
		s.ReplyCodeData = &InitInitReplyCodeDataOutOfDate{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_Ok:
		s.ReplyCodeData = &InitInitReplyCodeDataOk{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_Banned:
		s.ReplyCodeData = &InitInitReplyCodeDataBanned{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_WarpMap:
		s.ReplyCodeData = &InitInitReplyCodeDataWarpMap{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_FileEmf:
		s.ReplyCodeData = &InitInitReplyCodeDataFileEmf{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_FileEif:
		s.ReplyCodeData = &InitInitReplyCodeDataFileEif{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_FileEnf:
		s.ReplyCodeData = &InitInitReplyCodeDataFileEnf{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_FileEsf:
		s.ReplyCodeData = &InitInitReplyCodeDataFileEsf{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_FileEcf:
		s.ReplyCodeData = &InitInitReplyCodeDataFileEcf{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_MapMutation:
		s.ReplyCodeData = &InitInitReplyCodeDataMapMutation{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_PlayersList:
		s.ReplyCodeData = &InitInitReplyCodeDataPlayersList{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case InitReply_PlayersListFriends:
		s.ReplyCodeData = &InitInitReplyCodeDataPlayersListFriends{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// WarpPlayerServerPacket :: Equivalent to INIT_INIT with InitReply.WarpMap.
type WarpPlayerServerPacket struct {
	MapFile MapFile
}

func (s WarpPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *WarpPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WarpPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// WelcomePingServerPacket :: Equivalent to INIT_INIT with InitReply.FileMap.
type WelcomePingServerPacket struct {
	MapFile MapFile
}

func (s WelcomePingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomePingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

func (s *WelcomePingServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomePingServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// WelcomePongServerPacket :: Equivalent to INIT_INIT with InitReply.FileEif.
type WelcomePongServerPacket struct {
	PubFile PubFile
}

func (s WelcomePongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomePongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

func (s *WelcomePongServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomePongServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// WelcomeNet242ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEnf.
type WelcomeNet242ServerPacket struct {
	PubFile PubFile
}

func (s WelcomeNet242ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet242ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net242
}

func (s *WelcomeNet242ServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomeNet242ServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// WelcomeNet243ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEsf.
type WelcomeNet243ServerPacket struct {
	PubFile PubFile
}

func (s WelcomeNet243ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet243ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net243
}

func (s *WelcomeNet243ServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomeNet243ServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// PlayersListServerPacket :: Equivalent to INIT_INIT with InitReply.PlayersList.
type PlayersListServerPacket struct {
	PlayersList PlayersList
}

func (s PlayersListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *PlayersListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *PlayersListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// WarpCreateServerPacket :: Equivalent to INIT_INIT with InitReply.MapMutation.
type WarpCreateServerPacket struct {
	MapFile MapFile
}

func (s WarpCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

func (s *WarpCreateServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapFile : field : MapFile
	if err = s.MapFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WarpCreateServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// PlayersReplyServerPacket :: Equivalent to INIT_INIT with InitReply.PlayersListFriends.
type PlayersReplyServerPacket struct {
	PlayersList PlayersListFriends
}

func (s PlayersReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *PlayersReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *PlayersReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// WelcomeNet244ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEcf.
type WelcomeNet244ServerPacket struct {
	PubFile PubFile
}

func (s WelcomeNet244ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet244ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net244
}

func (s *WelcomeNet244ServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PubFile : field : PubFile
	if err = s.PubFile.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomeNet244ServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}

	return
}

// ConnectionPlayerServerPacket :: Ping request.
type ConnectionPlayerServerPacket struct {
	Seq1 int
	Seq2 int
}

func (s ConnectionPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Connection
}

func (s ConnectionPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *ConnectionPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Seq1 : field : short
	if err = writer.AddShort(s.Seq1); err != nil {
		return
	}
	// Seq2 : field : char
	if err = writer.AddChar(s.Seq2); err != nil {
		return
	}
	return
}

func (s *ConnectionPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Seq1 : field : short
	s.Seq1 = reader.GetShort()
	// Seq2 : field : char
	s.Seq2 = reader.GetChar()

	return
}

// AccountReplyServerPacket :: Reply to client Account-family packets.
type AccountReplyServerPacket struct {
	ReplyCode     AccountReply //  Sometimes an AccountReply code, sometimes a session ID for account creation.
	ReplyCodeData AccountReplyReplyCodeData
}

type AccountReplyReplyCodeData interface {
	protocol.EoData
}

type AccountReplyReplyCodeDataExists struct {
}

func (s *AccountReplyReplyCodeDataExists) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataExists) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type AccountReplyReplyCodeDataNotApproved struct {
}

func (s *AccountReplyReplyCodeDataNotApproved) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataNotApproved) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type AccountReplyReplyCodeDataCreated struct {
}

func (s *AccountReplyReplyCodeDataCreated) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("GO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataCreated) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type AccountReplyReplyCodeDataChangeFailed struct {
}

func (s *AccountReplyReplyCodeDataChangeFailed) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataChangeFailed) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type AccountReplyReplyCodeDataChanged struct {
}

func (s *AccountReplyReplyCodeDataChanged) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataChanged) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type AccountReplyReplyCodeDataRequestDenied struct {
}

func (s *AccountReplyReplyCodeDataRequestDenied) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataRequestDenied) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AccountReplyReplyCodeDataDefault ::  In this case (reply_code > 9), reply_code is a session ID for account creation.
type AccountReplyReplyCodeDataDefault struct {
	SequenceStart int
}

func (s *AccountReplyReplyCodeDataDefault) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SequenceStart : field : char
	if err = writer.AddChar(s.SequenceStart); err != nil {
		return
	}
	//  : field : string
	if err = writer.AddString("OK"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataDefault) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SequenceStart : field : char
	s.SequenceStart = reader.GetChar()
	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

func (s AccountReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Account
}

func (s AccountReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *AccountReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : AccountReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case AccountReply_Exists:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataExists:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case AccountReply_NotApproved:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataNotApproved:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case AccountReply_Created:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataCreated:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case AccountReply_ChangeFailed:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataChangeFailed:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case AccountReply_Changed:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataChanged:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case AccountReply_RequestDenied:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataRequestDenied:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	default:
		switch s.ReplyCodeData.(type) {
		case *AccountReplyReplyCodeDataDefault:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *AccountReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : AccountReply
	s.ReplyCode = AccountReply(reader.GetShort())
	switch s.ReplyCode {
	case AccountReply_Exists:
		s.ReplyCodeData = &AccountReplyReplyCodeDataExists{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case AccountReply_NotApproved:
		s.ReplyCodeData = &AccountReplyReplyCodeDataNotApproved{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case AccountReply_Created:
		s.ReplyCodeData = &AccountReplyReplyCodeDataCreated{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case AccountReply_ChangeFailed:
		s.ReplyCodeData = &AccountReplyReplyCodeDataChangeFailed{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case AccountReply_Changed:
		s.ReplyCodeData = &AccountReplyReplyCodeDataChanged{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case AccountReply_RequestDenied:
		s.ReplyCodeData = &AccountReplyReplyCodeDataRequestDenied{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	default:
		s.ReplyCodeData = &AccountReplyReplyCodeDataDefault{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// CharacterReplyServerPacket :: Reply to client Character-family packets.
type CharacterReplyServerPacket struct {
	ReplyCode     CharacterReply //  Sometimes a CharacterReply code, sometimes a session ID for character creation.
	ReplyCodeData CharacterReplyReplyCodeData
}

type CharacterReplyReplyCodeData interface {
	protocol.EoData
}

type CharacterReplyReplyCodeDataExists struct {
}

func (s *CharacterReplyReplyCodeDataExists) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataExists) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type CharacterReplyReplyCodeDataFull struct {
}

func (s *CharacterReplyReplyCodeDataFull) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataFull) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type CharacterReplyReplyCodeDataFull3 struct {
}

func (s *CharacterReplyReplyCodeDataFull3) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataFull3) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type CharacterReplyReplyCodeDataNotApproved struct {
}

func (s *CharacterReplyReplyCodeDataNotApproved) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataNotApproved) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type CharacterReplyReplyCodeDataOk struct {
	CharactersCount int

	Characters []CharacterSelectionListEntry
}

func (s *CharacterReplyReplyCodeDataOk) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	return
}

func (s *CharacterReplyReplyCodeDataOk) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	//  : field : char
	reader.GetChar()
	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

type CharacterReplyReplyCodeDataDeleted struct {
	CharactersCount int
	Characters      []CharacterSelectionListEntry
}

func (s *CharacterReplyReplyCodeDataDeleted) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	return
}

func (s *CharacterReplyReplyCodeDataDeleted) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// CharacterReplyReplyCodeDataDefault ::  In this case (reply_code > 9), reply_code is a session ID for character creation.
type CharacterReplyReplyCodeDataDefault struct {
}

func (s *CharacterReplyReplyCodeDataDefault) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("OK"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataDefault) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

func (s CharacterReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *CharacterReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// ReplyCode : field : CharacterReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case CharacterReply_Exists:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataExists:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case CharacterReply_Full:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataFull:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case CharacterReply_Full3:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataFull3:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case CharacterReply_NotApproved:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataNotApproved:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case CharacterReply_Ok:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataOk:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case CharacterReply_Deleted:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataDeleted:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	default:
		switch s.ReplyCodeData.(type) {
		case *CharacterReplyReplyCodeDataDefault:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *CharacterReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// ReplyCode : field : CharacterReply
	s.ReplyCode = CharacterReply(reader.GetShort())
	switch s.ReplyCode {
	case CharacterReply_Exists:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataExists{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case CharacterReply_Full:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataFull{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case CharacterReply_Full3:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataFull3{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case CharacterReply_NotApproved:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataNotApproved{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case CharacterReply_Ok:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataOk{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case CharacterReply_Deleted:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataDeleted{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	default:
		s.ReplyCodeData = &CharacterReplyReplyCodeDataDefault{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetIsChunked(false)

	return
}

// CharacterPlayerServerPacket :: Reply to client request to delete a character from the account (Character_Take).
type CharacterPlayerServerPacket struct {
	SessionId   int
	CharacterId int
}

func (s CharacterPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *CharacterPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}
	return
}

func (s *CharacterPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()

	return
}

// LoginReplyServerPacket :: Login reply.
type LoginReplyServerPacket struct {
	ReplyCode     LoginReply
	ReplyCodeData LoginReplyReplyCodeData
}

type LoginReplyReplyCodeData interface {
	protocol.EoData
}

type LoginReplyReplyCodeDataWrongUser struct {
}

func (s *LoginReplyReplyCodeDataWrongUser) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataWrongUser) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type LoginReplyReplyCodeDataWrongUserPassword struct {
}

func (s *LoginReplyReplyCodeDataWrongUserPassword) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataWrongUserPassword) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type LoginReplyReplyCodeDataOk struct {
	CharactersCount int

	Characters []CharacterSelectionListEntry
}

func (s *LoginReplyReplyCodeDataOk) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	return
}

func (s *LoginReplyReplyCodeDataOk) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	//  : field : char
	reader.GetChar()
	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

type LoginReplyReplyCodeDataBanned struct {
}

func (s *LoginReplyReplyCodeDataBanned) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataBanned) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type LoginReplyReplyCodeDataLoggedIn struct {
}

func (s *LoginReplyReplyCodeDataLoggedIn) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataLoggedIn) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

type LoginReplyReplyCodeDataBusy struct {
}

func (s *LoginReplyReplyCodeDataBusy) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataBusy) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

func (s LoginReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Login
}

func (s LoginReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *LoginReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// ReplyCode : field : LoginReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case LoginReply_WrongUser:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataWrongUser:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case LoginReply_WrongUserPassword:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataWrongUserPassword:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case LoginReply_Ok:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataOk:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case LoginReply_Banned:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataBanned:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case LoginReply_LoggedIn:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataLoggedIn:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case LoginReply_Busy:
		switch s.ReplyCodeData.(type) {
		case *LoginReplyReplyCodeDataBusy:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *LoginReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// ReplyCode : field : LoginReply
	s.ReplyCode = LoginReply(reader.GetShort())
	switch s.ReplyCode {
	case LoginReply_WrongUser:
		s.ReplyCodeData = &LoginReplyReplyCodeDataWrongUser{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case LoginReply_WrongUserPassword:
		s.ReplyCodeData = &LoginReplyReplyCodeDataWrongUserPassword{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case LoginReply_Ok:
		s.ReplyCodeData = &LoginReplyReplyCodeDataOk{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case LoginReply_Banned:
		s.ReplyCodeData = &LoginReplyReplyCodeDataBanned{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case LoginReply_LoggedIn:
		s.ReplyCodeData = &LoginReplyReplyCodeDataLoggedIn{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case LoginReply_Busy:
		s.ReplyCodeData = &LoginReplyReplyCodeDataBusy{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetIsChunked(false)

	return
}

// WelcomeReplyServerPacket :: Reply to selecting a character / entering game.
type WelcomeReplyServerPacket struct {
	WelcomeCode     WelcomeCode
	WelcomeCodeData WelcomeReplyWelcomeCodeData
}

type WelcomeReplyWelcomeCodeData interface {
	protocol.EoData
}

type WelcomeReplyWelcomeCodeDataSelectCharacter struct {
	SessionId        int
	CharacterId      int
	MapId            int
	MapRid           []int
	MapFileSize      int
	EifRid           []int
	EifLength        int
	EnfRid           []int
	EnfLength        int
	EsfRid           []int
	EsfLength        int
	EcfRid           []int
	EcfLength        int
	Name             string
	Title            string
	GuildName        string
	GuildRankName    string
	ClassId          int
	GuildTag         string
	Admin            protocol.AdminLevel
	Level            int
	Experience       int
	Usage            int
	Stats            CharacterStatsWelcome
	Equipment        EquipmentWelcome
	GuildRank        int
	Settings         ServerSettings
	LoginMessageCode LoginMessageCode
}

func (s *WelcomeReplyWelcomeCodeDataSelectCharacter) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}
	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	// MapRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.MapRid[ndx]); err != nil {
			return
		}
	}

	// MapFileSize : field : three
	if err = writer.AddThree(s.MapFileSize); err != nil {
		return
	}
	// EifRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.EifRid[ndx]); err != nil {
			return
		}
	}

	// EifLength : field : short
	if err = writer.AddShort(s.EifLength); err != nil {
		return
	}
	// EnfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.EnfRid[ndx]); err != nil {
			return
		}
	}

	// EnfLength : field : short
	if err = writer.AddShort(s.EnfLength); err != nil {
		return
	}
	// EsfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.EsfRid[ndx]); err != nil {
			return
		}
	}

	// EsfLength : field : short
	if err = writer.AddShort(s.EsfLength); err != nil {
		return
	}
	// EcfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.EcfRid[ndx]); err != nil {
			return
		}
	}

	// EcfLength : field : short
	if err = writer.AddShort(s.EcfLength); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Title : field : string
	if err = writer.AddString(s.Title); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildRankName : field : string
	if err = writer.AddString(s.GuildRankName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	// GuildTag : field : string
	if err = writer.AddFixedString(s.GuildTag, 3); err != nil {
		return
	}
	// Admin : field : AdminLevel
	if err = writer.AddChar(int(s.Admin)); err != nil {
		return
	}
	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// Usage : field : int
	if err = writer.AddInt(s.Usage); err != nil {
		return
	}
	// Stats : field : CharacterStatsWelcome
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	// Equipment : field : EquipmentWelcome
	if err = s.Equipment.Serialize(writer); err != nil {
		return
	}
	// GuildRank : field : char
	if err = writer.AddChar(s.GuildRank); err != nil {
		return
	}
	// Settings : field : ServerSettings
	if err = s.Settings.Serialize(writer); err != nil {
		return
	}
	// LoginMessageCode : field : LoginMessageCode
	if err = writer.AddChar(int(s.LoginMessageCode)); err != nil {
		return
	}
	writer.AddByte(0xFF)
	return
}

func (s *WelcomeReplyWelcomeCodeDataSelectCharacter) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// MapRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.MapRid = append(s.MapRid, 0)
		s.MapRid[ndx] = reader.GetShort()
	}

	// MapFileSize : field : three
	s.MapFileSize = reader.GetThree()
	// EifRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.EifRid = append(s.EifRid, 0)
		s.EifRid[ndx] = reader.GetShort()
	}

	// EifLength : field : short
	s.EifLength = reader.GetShort()
	// EnfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.EnfRid = append(s.EnfRid, 0)
		s.EnfRid[ndx] = reader.GetShort()
	}

	// EnfLength : field : short
	s.EnfLength = reader.GetShort()
	// EsfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.EsfRid = append(s.EsfRid, 0)
		s.EsfRid[ndx] = reader.GetShort()
	}

	// EsfLength : field : short
	s.EsfLength = reader.GetShort()
	// EcfRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.EcfRid = append(s.EcfRid, 0)
		s.EcfRid[ndx] = reader.GetShort()
	}

	// EcfLength : field : short
	s.EcfLength = reader.GetShort()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Title : field : string
	if s.Title, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// GuildName : field : string
	if s.GuildName, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// GuildRankName : field : string
	if s.GuildRankName, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetFixedString(3); err != nil {
		return
	}

	// Admin : field : AdminLevel
	s.Admin = protocol.AdminLevel(reader.GetChar())
	// Level : field : char
	s.Level = reader.GetChar()
	// Experience : field : int
	s.Experience = reader.GetInt()
	// Usage : field : int
	s.Usage = reader.GetInt()
	// Stats : field : CharacterStatsWelcome
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	// Equipment : field : EquipmentWelcome
	if err = s.Equipment.Deserialize(reader); err != nil {
		return
	}
	// GuildRank : field : char
	s.GuildRank = reader.GetChar()
	// Settings : field : ServerSettings
	if err = s.Settings.Deserialize(reader); err != nil {
		return
	}
	// LoginMessageCode : field : LoginMessageCode
	s.LoginMessageCode = LoginMessageCode(reader.GetChar())
	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}

	return
}

type WelcomeReplyWelcomeCodeDataEnterGame struct {
	News   []string
	Weight net.Weight
	Items  []net.Item
	Spells []net.Spell
	Nearby NearbyInfo
}

func (s *WelcomeReplyWelcomeCodeDataEnterGame) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.AddByte(0xFF)
	// News : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.News[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// Items : array : Item
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Spells : array : Spell
	for ndx := 0; ndx < len(s.Spells); ndx++ {
		if err = s.Spells[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomeReplyWelcomeCodeDataEnterGame) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// News : array : string
	for ndx := 0; ndx < 9; ndx++ {
		s.News = append(s.News, "")
		if s.News[ndx], err = reader.GetString(); err != nil {
			return
		}

	}

	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// Items : array : Item
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.Item{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Spells : array : Spell
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Spells = append(s.Spells, net.Spell{})
		if err = s.Spells[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}

	return
}

func (s WelcomeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *WelcomeReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WelcomeCode : field : WelcomeCode
	if err = writer.AddShort(int(s.WelcomeCode)); err != nil {
		return
	}
	writer.SanitizeStrings = true
	switch s.WelcomeCode {
	case WelcomeCode_SelectCharacter:
		switch s.WelcomeCodeData.(type) {
		case *WelcomeReplyWelcomeCodeDataSelectCharacter:
			if err = s.WelcomeCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.WelcomeCode)
			return
		}
	case WelcomeCode_EnterGame:
		switch s.WelcomeCodeData.(type) {
		case *WelcomeReplyWelcomeCodeDataEnterGame:
			if err = s.WelcomeCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.WelcomeCode)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *WelcomeReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// WelcomeCode : field : WelcomeCode
	s.WelcomeCode = WelcomeCode(reader.GetShort())
	reader.SetIsChunked(true)
	switch s.WelcomeCode {
	case WelcomeCode_SelectCharacter:
		s.WelcomeCodeData = &WelcomeReplyWelcomeCodeDataSelectCharacter{}
		if err = s.WelcomeCodeData.Deserialize(reader); err != nil {
			return
		}
	case WelcomeCode_EnterGame:
		s.WelcomeCodeData = &WelcomeReplyWelcomeCodeDataEnterGame{}
		if err = s.WelcomeCodeData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetIsChunked(false)

	return
}

// AdminInteractReplyServerPacket :: Incoming admin message.
type AdminInteractReplyServerPacket struct {
	MessageType     AdminMessageType
	MessageTypeData AdminInteractReplyMessageTypeData
}

type AdminInteractReplyMessageTypeData interface {
	protocol.EoData
}

type AdminInteractReplyMessageTypeDataMessage struct {
	PlayerName string
	Message    string
}

func (s *AdminInteractReplyMessageTypeDataMessage) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.AddByte(0xFF)
	return
}

func (s *AdminInteractReplyMessageTypeDataMessage) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}

	return
}

type AdminInteractReplyMessageTypeDataReport struct {
	PlayerName   string
	Message      string
	ReporteeName string
}

func (s *AdminInteractReplyMessageTypeDataReport) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// ReporteeName : field : string
	if err = writer.AddString(s.ReporteeName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	return
}

func (s *AdminInteractReplyMessageTypeDataReport) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}
	// ReporteeName : field : string
	if s.ReporteeName, err = reader.GetString(); err != nil {
		return
	}

	if breakByte := reader.GetByte(); breakByte != 0xFF {
		return fmt.Errorf("missing expected break byte")
	}

	return
}

func (s AdminInteractReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *AdminInteractReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// MessageType : field : AdminMessageType
	if err = writer.AddChar(int(s.MessageType)); err != nil {
		return
	}
	writer.AddByte(0xFF)
	switch s.MessageType {
	case AdminMessage_Message:
		switch s.MessageTypeData.(type) {
		case *AdminInteractReplyMessageTypeDataMessage:
			if err = s.MessageTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.MessageType)
			return
		}
	case AdminMessage_Report:
		switch s.MessageTypeData.(type) {
		case *AdminInteractReplyMessageTypeDataReport:
			if err = s.MessageTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.MessageType)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *AdminInteractReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// MessageType : field : AdminMessageType
	s.MessageType = AdminMessageType(reader.GetChar())
	if err = reader.NextChunk(); err != nil {
		return
	}
	switch s.MessageType {
	case AdminMessage_Message:
		s.MessageTypeData = &AdminInteractReplyMessageTypeDataMessage{}
		if err = s.MessageTypeData.Deserialize(reader); err != nil {
			return
		}
	case AdminMessage_Report:
		s.MessageTypeData = &AdminInteractReplyMessageTypeDataReport{}
		if err = s.MessageTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetIsChunked(false)

	return
}

// AdminInteractRemoveServerPacket :: Nearby player disappearing (admin hide).
type AdminInteractRemoveServerPacket struct {
	PlayerId int
}

func (s AdminInteractRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *AdminInteractRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *AdminInteractRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// AdminInteractAgreeServerPacket :: Nearby player appearing (admin un-hide).
type AdminInteractAgreeServerPacket struct {
	PlayerId int
}

func (s AdminInteractAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *AdminInteractAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *AdminInteractAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// AdminInteractListServerPacket :: Admin character inventory popup.
type AdminInteractListServerPacket struct {
	Name      string
	Usage     int
	GoldBank  int
	Inventory []net.Item
	Bank      []net.ThreeItem
}

func (s AdminInteractListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *AdminInteractListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Usage : field : int
	if err = writer.AddInt(s.Usage); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GoldBank : field : int
	if err = writer.AddInt(s.GoldBank); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Inventory : array : Item
	for ndx := 0; ndx < len(s.Inventory); ndx++ {
		if err = s.Inventory[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Bank : array : ThreeItem
	for ndx := 0; ndx < len(s.Bank); ndx++ {
		if err = s.Bank[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *AdminInteractListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Usage : field : int
	s.Usage = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// GoldBank : field : int
	s.GoldBank = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Inventory : array : Item
	for ndx := 0; ndx < reader.Remaining()/6; ndx++ {
		s.Inventory = append(s.Inventory, net.Item{})
		if err = s.Inventory[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Bank : array : ThreeItem
	for ndx := 0; ndx < reader.Remaining()/5; ndx++ {
		s.Bank = append(s.Bank, net.ThreeItem{})
		if err = s.Bank[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// AdminInteractTellServerPacket :: Admin character info lookup.
type AdminInteractTellServerPacket struct {
	Name      string
	Usage     int
	Exp       int
	Level     int
	MapId     int
	MapCoords BigCoords
	Stats     CharacterStatsInfoLookup
	Weight    net.Weight
}

func (s AdminInteractTellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractTellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

func (s *AdminInteractTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Usage : field : int
	if err = writer.AddInt(s.Usage); err != nil {
		return
	}
	writer.AddByte(0xFF)
	writer.AddByte(0xFF)
	// Exp : field : int
	if err = writer.AddInt(s.Exp); err != nil {
		return
	}
	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	// MapCoords : field : BigCoords
	if err = s.MapCoords.Serialize(writer); err != nil {
		return
	}
	// Stats : field : CharacterStatsInfoLookup
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *AdminInteractTellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Usage : field : int
	s.Usage = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Exp : field : int
	s.Exp = reader.GetInt()
	// Level : field : char
	s.Level = reader.GetChar()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// MapCoords : field : BigCoords
	if err = s.MapCoords.Deserialize(reader); err != nil {
		return
	}
	// Stats : field : CharacterStatsInfoLookup
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// TalkRequestServerPacket :: Guild chat message.
type TalkRequestServerPacket struct {
	PlayerName string
	Message    string
}

func (s TalkRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *TalkRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *TalkRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// TalkOpenServerPacket :: Party chat message.
type TalkOpenServerPacket struct {
	PlayerId int
	Message  string
}

func (s TalkOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *TalkOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkMsgServerPacket :: Global chat message.
type TalkMsgServerPacket struct {
	PlayerName string
	Message    string
}

func (s TalkMsgServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkMsgServerPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

func (s *TalkMsgServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *TalkMsgServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// TalkTellServerPacket :: Private chat message.
type TalkTellServerPacket struct {
	PlayerName string
	Message    string
}

func (s TalkTellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkTellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

func (s *TalkTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *TalkTellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// TalkPlayerServerPacket :: Public chat message.
type TalkPlayerServerPacket struct {
	PlayerId int
	Message  string
}

func (s TalkPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *TalkPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkReplyServerPacket :: Reply to trying to send a private message.
type TalkReplyServerPacket struct {
	ReplyCode TalkReply
	Name      string
}

func (s TalkReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *TalkReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : TalkReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *TalkReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : TalkReply
	s.ReplyCode = TalkReply(reader.GetShort())
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkAdminServerPacket :: Admin chat message.
type TalkAdminServerPacket struct {
	PlayerName string
	Message    string
}

func (s TalkAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

func (s *TalkAdminServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *TalkAdminServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// TalkAnnounceServerPacket :: Admin announcement.
type TalkAnnounceServerPacket struct {
	PlayerName string
	Message    string
}

func (s TalkAnnounceServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAnnounceServerPacket) Action() net.PacketAction {
	return net.PacketAction_Announce
}

func (s *TalkAnnounceServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *TalkAnnounceServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// TalkServerServerPacket :: Server message.
type TalkServerServerPacket struct {
	Message string
}

func (s TalkServerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkServerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Server
}

func (s *TalkServerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkServerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkListServerPacket ::  Global chat backfill. Sent by the official game server when a player opens the global chat tab.
type TalkListServerPacket struct {
	Messages []GlobalBackfillMessage
}

func (s TalkListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *TalkListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Messages : array : GlobalBackfillMessage
	for ndx := 0; ndx < len(s.Messages); ndx++ {
		if err = s.Messages[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *TalkListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Messages : array : GlobalBackfillMessage
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Messages = append(s.Messages, GlobalBackfillMessage{})
		if err = s.Messages[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// MessageOpenServerPacket :: Status bar message.
type MessageOpenServerPacket struct {
	Message string
}

func (s MessageOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *MessageOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *MessageOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// MessageCloseServerPacket :: Server is rebooting.
type MessageCloseServerPacket struct {
}

func (s MessageCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *MessageCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("r"); err != nil {
		return
	}
	return
}

func (s *MessageCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// MessageAcceptServerPacket :: Large message box.
type MessageAcceptServerPacket struct {
	Messages []string
}

func (s MessageAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *MessageAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Messages : array : string
	for ndx := 0; ndx < 4; ndx++ {
		if err = writer.AddString(s.Messages[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *MessageAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Messages : array : string
	for ndx := 0; ndx < 4; ndx++ {
		s.Messages = append(s.Messages, "")
		if s.Messages[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// TalkSpecServerPacket :: Temporary mute applied.
type TalkSpecServerPacket struct {
	AdminName string
}

func (s TalkSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *TalkSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// AdminName : field : string
	if err = writer.AddString(s.AdminName); err != nil {
		return
	}
	return
}

func (s *TalkSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// AdminName : field : string
	if s.AdminName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AttackPlayerServerPacket :: Nearby player attacking.
type AttackPlayerServerPacket struct {
	PlayerId  int
	Direction protocol.Direction
}

func (s AttackPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Attack
}

func (s AttackPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *AttackPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *AttackPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// AttackErrorServerPacket :: Show flood protection message (vestigial).
type AttackErrorServerPacket struct {
}

func (s AttackErrorServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Attack
}

func (s AttackErrorServerPacket) Action() net.PacketAction {
	return net.PacketAction_Error
}

func (s *AttackErrorServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *AttackErrorServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// AvatarReplyServerPacket :: Nearby player hit by another player.
type AvatarReplyServerPacket struct {
	PlayerId     int
	VictimId     int
	Damage       int
	Direction    protocol.Direction
	HpPercentage int
	Dead         bool
}

func (s AvatarReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *AvatarReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// VictimId : field : short
	if err = writer.AddShort(s.VictimId); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Dead : field : bool
	if s.Dead {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	return
}

func (s *AvatarReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// VictimId : field : short
	s.VictimId = reader.GetShort()
	// Damage : field : three
	s.Damage = reader.GetThree()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Dead : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Dead = true
	} else {
		s.Dead = false
	}

	return
}

// ChairPlayerServerPacket :: Nearby player sitting on a chair.
type ChairPlayerServerPacket struct {
	PlayerId  int
	Coords    protocol.Coords
	Direction protocol.Direction
}

func (s ChairPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *ChairPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *ChairPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// ChairReplyServerPacket :: Your character sitting on a chair.
type ChairReplyServerPacket struct {
	PlayerId  int
	Coords    protocol.Coords
	Direction protocol.Direction
}

func (s ChairReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *ChairReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *ChairReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// ChairCloseServerPacket :: Your character standing up from a chair.
type ChairCloseServerPacket struct {
	PlayerId int
	Coords   protocol.Coords
}

func (s ChairCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *ChairCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChairCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// ChairRemoveServerPacket :: Nearby player standing up from a chair.
type ChairRemoveServerPacket struct {
	PlayerId int
	Coords   protocol.Coords
}

func (s ChairRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *ChairRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChairRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// SitPlayerServerPacket :: Nearby player sitting down.
type SitPlayerServerPacket struct {
	PlayerId  int
	Coords    protocol.Coords
	Direction protocol.Direction
}

func (s SitPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *SitPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *SitPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	//  : field : char
	reader.GetChar()

	return
}

// SitCloseServerPacket :: Your character standing up.
type SitCloseServerPacket struct {
	PlayerId int
	Coords   protocol.Coords
}

func (s SitCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *SitCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *SitCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// SitRemoveServerPacket :: Nearby player standing up.
type SitRemoveServerPacket struct {
	PlayerId int
	Coords   protocol.Coords
}

func (s SitRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *SitRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *SitRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// SitReplyServerPacket :: Your character sitting down.
type SitReplyServerPacket struct {
	PlayerId  int
	Coords    protocol.Coords
	Direction protocol.Direction
}

func (s SitReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *SitReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *SitReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	//  : field : char
	reader.GetChar()

	return
}

// EmotePlayerServerPacket :: Nearby player doing an emote.
type EmotePlayerServerPacket struct {
	PlayerId int
	Emote    protocol.Emote
}

func (s EmotePlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Emote
}

func (s EmotePlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *EmotePlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Emote : field : Emote
	if err = writer.AddChar(int(s.Emote)); err != nil {
		return
	}
	return
}

func (s *EmotePlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Emote : field : Emote
	s.Emote = protocol.Emote(reader.GetChar())

	return
}

// EffectPlayerServerPacket :: Nearby player doing an effect.
type EffectPlayerServerPacket struct {
	PlayerId int
	EffectId int
}

func (s EffectPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *EffectPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// EffectId : field : three
	if err = writer.AddThree(s.EffectId); err != nil {
		return
	}
	return
}

func (s *EffectPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// EffectId : field : three
	s.EffectId = reader.GetThree()

	return
}

// FacePlayerServerPacket :: Nearby player facing a direction.
type FacePlayerServerPacket struct {
	PlayerId  int
	Direction protocol.Direction
}

func (s FacePlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Face
}

func (s FacePlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *FacePlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *FacePlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// AvatarRemoveServerPacket :: Nearby player has disappeared from view.
type AvatarRemoveServerPacket struct {
	PlayerId   int
	WarpEffect *WarpEffect
}

func (s AvatarRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *AvatarRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// WarpEffect : field : WarpEffect
	if s.WarpEffect != nil {
		if err = writer.AddChar(int(*s.WarpEffect)); err != nil {
			return
		}
	}
	return
}

func (s *AvatarRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// WarpEffect : field : WarpEffect
	if reader.Remaining() > 0 {
		s.WarpEffect = new(WarpEffect)
		*s.WarpEffect = WarpEffect(reader.GetChar())
	}

	return
}

// PlayersAgreeServerPacket :: Player has appeared in nearby view.
type PlayersAgreeServerPacket struct {
	Nearby NearbyInfo
}

func (s PlayersAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *PlayersAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *PlayersAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}

	return
}

// PlayersRemoveServerPacket :: Nearby player has logged out.
type PlayersRemoveServerPacket struct {
	PlayerId int
}

func (s PlayersRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *PlayersRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *PlayersRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// RangeReplyServerPacket :: Reply to request for information about nearby players and NPCs.
type RangeReplyServerPacket struct {
	Nearby NearbyInfo
}

func (s RangeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Range
}

func (s RangeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *RangeReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *RangeReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}

	return
}

// NpcAgreeServerPacket :: Reply to request for information about nearby NPCs.
type NpcAgreeServerPacket struct {
	NpcsCount int
	Npcs      []NpcMapInfo
}

func (s NpcAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *NpcAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcsCount : length : short
	if err = writer.AddShort(s.NpcsCount); err != nil {
		return
	}
	// Npcs : array : NpcMapInfo
	for ndx := 0; ndx < s.NpcsCount; ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *NpcAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcsCount : length : short
	s.NpcsCount = reader.GetShort()
	// Npcs : array : NpcMapInfo
	for ndx := 0; ndx < s.NpcsCount; ndx++ {
		s.Npcs = append(s.Npcs, NpcMapInfo{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// WalkPlayerServerPacket :: Nearby player has walked.
type WalkPlayerServerPacket struct {
	PlayerId  int
	Direction protocol.Direction
	Coords    protocol.Coords
}

func (s WalkPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *WalkPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// WalkReplyServerPacket :: Players, NPCs, and Items appearing in nearby view.
type WalkReplyServerPacket struct {
	PlayerIds  []int
	NpcIndexes []int
	Items      []ItemMapInfo
}

func (s WalkReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *WalkReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerIds : array : short
	for ndx := 0; ndx < len(s.PlayerIds); ndx++ {
		if err = writer.AddShort(s.PlayerIds[ndx]); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// NpcIndexes : array : char
	for ndx := 0; ndx < len(s.NpcIndexes); ndx++ {
		if err = writer.AddChar(s.NpcIndexes[ndx]); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Items : array : ItemMapInfo
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *WalkReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerIds : array : short
	for ndx := 0; ndx < reader.Remaining()/2; ndx++ {
		s.PlayerIds = append(s.PlayerIds, 0)
		s.PlayerIds[ndx] = reader.GetShort()
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// NpcIndexes : array : char
	for ndx := 0; ndx < reader.Remaining()/1; ndx++ {
		s.NpcIndexes = append(s.NpcIndexes, 0)
		s.NpcIndexes[ndx] = reader.GetChar()
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Items : array : ItemMapInfo
	for ndx := 0; ndx < reader.Remaining()/9; ndx++ {
		s.Items = append(s.Items, ItemMapInfo{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// WalkCloseServerPacket :: Your character has been frozen.
type WalkCloseServerPacket struct {
}

func (s WalkCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *WalkCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *WalkCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// WalkOpenServerPacket :: Your character has been unfrozen.
type WalkOpenServerPacket struct {
}

func (s WalkOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *WalkOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *WalkOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// BankOpenServerPacket :: Open banker NPC interface.
type BankOpenServerPacket struct {
	GoldBank       int
	SessionId      int
	LockerUpgrades int
}

func (s BankOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *BankOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldBank : field : int
	if err = writer.AddInt(s.GoldBank); err != nil {
		return
	}
	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}
	// LockerUpgrades : field : char
	if err = writer.AddChar(s.LockerUpgrades); err != nil {
		return
	}
	return
}

func (s *BankOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldBank : field : int
	s.GoldBank = reader.GetInt()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	// LockerUpgrades : field : char
	s.LockerUpgrades = reader.GetChar()

	return
}

// BankReplyServerPacket :: Update gold counts after deposit/withdraw.
type BankReplyServerPacket struct {
	GoldInventory int
	GoldBank      int
}

func (s BankReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *BankReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldInventory : field : int
	if err = writer.AddInt(s.GoldInventory); err != nil {
		return
	}
	// GoldBank : field : int
	if err = writer.AddInt(s.GoldBank); err != nil {
		return
	}
	return
}

func (s *BankReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldInventory : field : int
	s.GoldInventory = reader.GetInt()
	// GoldBank : field : int
	s.GoldBank = reader.GetInt()

	return
}

// BarberAgreeServerPacket :: Purchasing a new hair style.
type BarberAgreeServerPacket struct {
	GoldAmount int
	Change     AvatarChange
}

func (s BarberAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *BarberAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	// Change : field : AvatarChange
	if err = s.Change.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *BarberAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}

	return
}

// BarberOpenServerPacket :: Response from talking to a barber NPC.
type BarberOpenServerPacket struct {
	SessionId int
}

func (s BarberOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *BarberOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *BarberOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// LockerReplyServerPacket :: Response to adding an item to a bank locker.
type LockerReplyServerPacket struct {
	DepositedItem net.Item
	Weight        net.Weight
	LockerItems   []net.ThreeItem
}

func (s LockerReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *LockerReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// DepositedItem : field : Item
	if err = s.DepositedItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; ndx < len(s.LockerItems); ndx++ {
		if err = s.LockerItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *LockerReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// DepositedItem : field : Item
	if err = s.DepositedItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.LockerItems = append(s.LockerItems, net.ThreeItem{})
		if err = s.LockerItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// LockerGetServerPacket :: Response to taking an item from a bank locker.
type LockerGetServerPacket struct {
	TakenItem   net.ThreeItem
	Weight      net.Weight
	LockerItems []net.ThreeItem
}

func (s LockerGetServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerGetServerPacket) Action() net.PacketAction {
	return net.PacketAction_Get
}

func (s *LockerGetServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; ndx < len(s.LockerItems); ndx++ {
		if err = s.LockerItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *LockerGetServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.LockerItems = append(s.LockerItems, net.ThreeItem{})
		if err = s.LockerItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// LockerOpenServerPacket :: Opening a bank locker.
type LockerOpenServerPacket struct {
	LockerCoords protocol.Coords
	LockerItems  []net.ThreeItem
}

func (s LockerOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *LockerOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Serialize(writer); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; ndx < len(s.LockerItems); ndx++ {
		if err = s.LockerItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *LockerOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	// LockerItems : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.LockerItems = append(s.LockerItems, net.ThreeItem{})
		if err = s.LockerItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// LockerBuyServerPacket :: Response to buying a locker space upgrade from a banker NPC.
type LockerBuyServerPacket struct {
	GoldAmount     int
	LockerUpgrades int
}

func (s LockerBuyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerBuyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

func (s *LockerBuyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	// LockerUpgrades : field : char
	if err = writer.AddChar(s.LockerUpgrades); err != nil {
		return
	}
	return
}

func (s *LockerBuyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// LockerUpgrades : field : char
	s.LockerUpgrades = reader.GetChar()

	return
}

// LockerSpecServerPacket :: Reply to trying to add an item to a full locker.
type LockerSpecServerPacket struct {
	LockerMaxItems int
}

func (s LockerSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *LockerSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerMaxItems : field : char
	if err = writer.AddChar(s.LockerMaxItems); err != nil {
		return
	}
	return
}

func (s *LockerSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// LockerMaxItems : field : char
	s.LockerMaxItems = reader.GetChar()

	return
}

// CitizenReplyServerPacket :: Response to subscribing to a town.
type CitizenReplyServerPacket struct {
	QuestionsWrong int
}

func (s CitizenReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *CitizenReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuestionsWrong : field : char
	if err = writer.AddChar(s.QuestionsWrong); err != nil {
		return
	}
	return
}

func (s *CitizenReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// QuestionsWrong : field : char
	s.QuestionsWrong = reader.GetChar()

	return
}

// CitizenRemoveServerPacket :: Response to giving up citizenship of a town.
type CitizenRemoveServerPacket struct {
	ReplyCode InnUnsubscribeReply
}

func (s CitizenRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *CitizenRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : InnUnsubscribeReply
	if err = writer.AddChar(int(s.ReplyCode)); err != nil {
		return
	}
	return
}

func (s *CitizenRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : InnUnsubscribeReply
	s.ReplyCode = InnUnsubscribeReply(reader.GetChar())

	return
}

// CitizenOpenServerPacket :: Response from talking to a citizenship NPC.
type CitizenOpenServerPacket struct {
	BehaviorId    int
	CurrentHomeId int
	SessionId     int
	Questions     []string
}

func (s CitizenOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *CitizenOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// BehaviorId : field : three
	if err = writer.AddThree(s.BehaviorId); err != nil {
		return
	}
	// CurrentHomeId : field : char
	if err = writer.AddChar(s.CurrentHomeId); err != nil {
		return
	}
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Questions : array : string
	for ndx := 0; ndx < 3; ndx++ {
		if ndx > 0 {
			writer.AddByte(0xFF)
		}

		if err = writer.AddString(s.Questions[ndx]); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *CitizenOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// BehaviorId : field : three
	s.BehaviorId = reader.GetThree()
	// CurrentHomeId : field : char
	s.CurrentHomeId = reader.GetChar()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Questions : array : string
	for ndx := 0; ndx < 3; ndx++ {
		s.Questions = append(s.Questions, "")
		if s.Questions[ndx], err = reader.GetString(); err != nil {
			return
		}

		if ndx+1 < 3 {
			if err = reader.NextChunk(); err != nil {
				return
			}
		}
	}

	reader.SetIsChunked(false)

	return
}

// CitizenRequestServerPacket :: Reply to requesting sleeping at an inn.
type CitizenRequestServerPacket struct {
	Cost int
}

func (s CitizenRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *CitizenRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Cost : field : int
	if err = writer.AddInt(s.Cost); err != nil {
		return
	}
	return
}

func (s *CitizenRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Cost : field : int
	s.Cost = reader.GetInt()

	return
}

// CitizenAcceptServerPacket :: Sleeping at an inn.
type CitizenAcceptServerPacket struct {
	GoldAmount int
}

func (s CitizenAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *CitizenAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *CitizenAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// ShopCreateServerPacket :: Response to crafting an item from a shop.
type ShopCreateServerPacket struct {
	CraftItemId int
	Weight      net.Weight
	Ingredients []net.Item
}

func (s ShopCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

func (s *ShopCreateServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CraftItemId : field : short
	if err = writer.AddShort(s.CraftItemId); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// Ingredients : array : Item
	for ndx := 0; ndx < 4; ndx++ {
		if err = s.Ingredients[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ShopCreateServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CraftItemId : field : short
	s.CraftItemId = reader.GetShort()
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// Ingredients : array : Item
	for ndx := 0; ndx < 4; ndx++ {
		s.Ingredients = append(s.Ingredients, net.Item{})
		if err = s.Ingredients[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ShopBuyServerPacket :: Response to purchasing an item from a shop.
type ShopBuyServerPacket struct {
	GoldAmount int
	BoughtItem net.Item
	Weight     net.Weight
}

func (s ShopBuyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopBuyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

func (s *ShopBuyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	// BoughtItem : field : Item
	if err = s.BoughtItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ShopBuyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// BoughtItem : field : Item
	if err = s.BoughtItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}

	return
}

// ShopSellServerPacket :: Response to selling an item to a shop.
type ShopSellServerPacket struct {
	SoldItem   ShopSoldItem
	GoldAmount int
	Weight     net.Weight
}

func (s ShopSellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopSellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Sell
}

func (s *ShopSellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SoldItem : field : ShopSoldItem
	if err = s.SoldItem.Serialize(writer); err != nil {
		return
	}
	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ShopSellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SoldItem : field : ShopSoldItem
	if err = s.SoldItem.Deserialize(reader); err != nil {
		return
	}
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}

	return
}

// ShopOpenServerPacket :: Response from talking to a shop NPC.
type ShopOpenServerPacket struct {
	SessionId  int
	ShopName   string
	TradeItems []ShopTradeItem
	CraftItems []ShopCraftItem
}

func (s ShopOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *ShopOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// ShopName : field : string
	if err = writer.AddString(s.ShopName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// TradeItems : array : ShopTradeItem
	for ndx := 0; ndx < len(s.TradeItems); ndx++ {
		if err = s.TradeItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// CraftItems : array : ShopCraftItem
	for ndx := 0; ndx < len(s.CraftItems); ndx++ {
		if err = s.CraftItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *ShopOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// ShopName : field : string
	if s.ShopName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// TradeItems : array : ShopTradeItem
	for ndx := 0; ndx < reader.Remaining()/9; ndx++ {
		s.TradeItems = append(s.TradeItems, ShopTradeItem{})
		if err = s.TradeItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// CraftItems : array : ShopCraftItem
	for ndx := 0; ndx < reader.Remaining()/2; ndx++ {
		s.CraftItems = append(s.CraftItems, ShopCraftItem{})
		if err = s.CraftItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// StatSkillOpenServerPacket :: Response from talking to a skill master NPC.
type StatSkillOpenServerPacket struct {
	SessionId int
	ShopName  string
	Skills    []SkillLearn
}

func (s StatSkillOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *StatSkillOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// ShopName : field : string
	if err = writer.AddString(s.ShopName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Skills : array : SkillLearn
	for ndx := 0; ndx < len(s.Skills); ndx++ {
		if err = s.Skills[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *StatSkillOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// ShopName : field : string
	if s.ShopName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Skills : array : SkillLearn
	for ndx := 0; ndx < reader.Remaining()/20; ndx++ {
		s.Skills = append(s.Skills, SkillLearn{})
		if err = s.Skills[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// StatSkillReplyServerPacket :: Response from unsuccessful action at a skill master.
type StatSkillReplyServerPacket struct {
	ReplyCode     SkillMasterReply
	ReplyCodeData StatSkillReplyReplyCodeData
}

type StatSkillReplyReplyCodeData interface {
	protocol.EoData
}

type StatSkillReplyReplyCodeDataWrongClass struct {
	ClassId int
}

func (s *StatSkillReplyReplyCodeDataWrongClass) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	return
}

func (s *StatSkillReplyReplyCodeDataWrongClass) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ClassId : field : char
	s.ClassId = reader.GetChar()

	return
}

func (s StatSkillReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *StatSkillReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : SkillMasterReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case SkillMasterReply_WrongClass:
		switch s.ReplyCodeData.(type) {
		case *StatSkillReplyReplyCodeDataWrongClass:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *StatSkillReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : SkillMasterReply
	s.ReplyCode = SkillMasterReply(reader.GetShort())
	switch s.ReplyCode {
	case SkillMasterReply_WrongClass:
		s.ReplyCodeData = &StatSkillReplyReplyCodeDataWrongClass{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// StatSkillTakeServerPacket :: Response from learning a skill from a skill master.
type StatSkillTakeServerPacket struct {
	SpellId    int
	GoldAmount int
}

func (s StatSkillTakeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillTakeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

func (s *StatSkillTakeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *StatSkillTakeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// StatSkillRemoveServerPacket :: Response to forgetting a skill at a skill master.
type StatSkillRemoveServerPacket struct {
	SpellId int
}

func (s StatSkillRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *StatSkillRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	return
}

func (s *StatSkillRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

// StatSkillPlayerServerPacket :: Response to spending stat points.
type StatSkillPlayerServerPacket struct {
	StatPoints int
	Stats      CharacterStatsUpdate
}

func (s StatSkillPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *StatSkillPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *StatSkillPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

// StatSkillAcceptServerPacket :: Response to spending skill points.
type StatSkillAcceptServerPacket struct {
	SkillPoints int
	Spell       net.Spell
}

func (s StatSkillAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *StatSkillAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// Spell : field : Spell
	if err = s.Spell.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *StatSkillAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// Spell : field : Spell
	if err = s.Spell.Deserialize(reader); err != nil {
		return
	}

	return
}

// StatSkillJunkServerPacket :: Response to resetting stats and skills at a skill master.
type StatSkillJunkServerPacket struct {
	Stats CharacterStatsReset
}

func (s StatSkillJunkServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillJunkServerPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

func (s *StatSkillJunkServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Stats : field : CharacterStatsReset
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *StatSkillJunkServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Stats : field : CharacterStatsReset
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemReplyServerPacket :: Reply to using an item.
type ItemReplyServerPacket struct {
	ItemType     pub.ItemType
	UsedItem     net.Item
	Weight       net.Weight
	ItemTypeData ItemReplyItemTypeData
}

type ItemReplyItemTypeData interface {
	protocol.EoData
}

type ItemReplyItemTypeDataHeal struct {
	HpGain int
	Hp     int
	Tp     int
}

func (s *ItemReplyItemTypeDataHeal) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HpGain : field : int
	if err = writer.AddInt(s.HpGain); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	return
}

func (s *ItemReplyItemTypeDataHeal) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// HpGain : field : int
	s.HpGain = reader.GetInt()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()

	return
}

type ItemReplyItemTypeDataHairDye struct {
	HairColor int
}

func (s *ItemReplyItemTypeDataHairDye) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}
	return
}

func (s *ItemReplyItemTypeDataHairDye) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// HairColor : field : char
	s.HairColor = reader.GetChar()

	return
}

type ItemReplyItemTypeDataEffectPotion struct {
	EffectId int
}

func (s *ItemReplyItemTypeDataEffectPotion) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EffectId : field : short
	if err = writer.AddShort(s.EffectId); err != nil {
		return
	}
	return
}

func (s *ItemReplyItemTypeDataEffectPotion) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// EffectId : field : short
	s.EffectId = reader.GetShort()

	return
}

type ItemReplyItemTypeDataCureCurse struct {
	Stats CharacterStatsEquipmentChange
}

func (s *ItemReplyItemTypeDataCureCurse) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemReplyItemTypeDataCureCurse) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

type ItemReplyItemTypeDataExpReward struct {
	Experience  int
	LevelUp     int //  A value greater than 0 is "new level" and indicates the player leveled up.
	StatPoints  int
	SkillPoints int
	MaxHp       int
	MaxTp       int
	MaxSp       int
}

func (s *ItemReplyItemTypeDataExpReward) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// LevelUp : field : char
	if err = writer.AddChar(s.LevelUp); err != nil {
		return
	}
	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	return
}

func (s *ItemReplyItemTypeDataExpReward) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Experience : field : int
	s.Experience = reader.GetInt()
	// LevelUp : field : char
	s.LevelUp = reader.GetChar()
	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()

	return
}

func (s ItemReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *ItemReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemType : field : ItemType
	if err = writer.AddChar(int(s.ItemType)); err != nil {
		return
	}
	// UsedItem : field : Item
	if err = s.UsedItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	switch s.ItemType {
	case pub.Item_Heal:
		switch s.ItemTypeData.(type) {
		case *ItemReplyItemTypeDataHeal:
			if err = s.ItemTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ItemType)
			return
		}
	case pub.Item_HairDye:
		switch s.ItemTypeData.(type) {
		case *ItemReplyItemTypeDataHairDye:
			if err = s.ItemTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ItemType)
			return
		}
	case pub.Item_EffectPotion:
		switch s.ItemTypeData.(type) {
		case *ItemReplyItemTypeDataEffectPotion:
			if err = s.ItemTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ItemType)
			return
		}
	case pub.Item_CureCurse:
		switch s.ItemTypeData.(type) {
		case *ItemReplyItemTypeDataCureCurse:
			if err = s.ItemTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ItemType)
			return
		}
	case pub.Item_ExpReward:
		switch s.ItemTypeData.(type) {
		case *ItemReplyItemTypeDataExpReward:
			if err = s.ItemTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ItemType)
			return
		}
	}
	return
}

func (s *ItemReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemType : field : ItemType
	s.ItemType = pub.ItemType(reader.GetChar())
	// UsedItem : field : Item
	if err = s.UsedItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	switch s.ItemType {
	case pub.Item_Heal:
		s.ItemTypeData = &ItemReplyItemTypeDataHeal{}
		if err = s.ItemTypeData.Deserialize(reader); err != nil {
			return
		}
	case pub.Item_HairDye:
		s.ItemTypeData = &ItemReplyItemTypeDataHairDye{}
		if err = s.ItemTypeData.Deserialize(reader); err != nil {
			return
		}
	case pub.Item_EffectPotion:
		s.ItemTypeData = &ItemReplyItemTypeDataEffectPotion{}
		if err = s.ItemTypeData.Deserialize(reader); err != nil {
			return
		}
	case pub.Item_CureCurse:
		s.ItemTypeData = &ItemReplyItemTypeDataCureCurse{}
		if err = s.ItemTypeData.Deserialize(reader); err != nil {
			return
		}
	case pub.Item_ExpReward:
		s.ItemTypeData = &ItemReplyItemTypeDataExpReward{}
		if err = s.ItemTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ItemDropServerPacket :: Reply to dropping items on the ground.
type ItemDropServerPacket struct {
	DroppedItem     net.ThreeItem
	RemainingAmount int
	ItemIndex       int
	Coords          protocol.Coords
	Weight          net.Weight
}

func (s ItemDropServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemDropServerPacket) Action() net.PacketAction {
	return net.PacketAction_Drop
}

func (s *ItemDropServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// DroppedItem : field : ThreeItem
	if err = s.DroppedItem.Serialize(writer); err != nil {
		return
	}
	// RemainingAmount : field : int
	if err = writer.AddInt(s.RemainingAmount); err != nil {
		return
	}
	// ItemIndex : field : short
	if err = writer.AddShort(s.ItemIndex); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemDropServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// DroppedItem : field : ThreeItem
	if err = s.DroppedItem.Deserialize(reader); err != nil {
		return
	}
	// RemainingAmount : field : int
	s.RemainingAmount = reader.GetInt()
	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemAddServerPacket :: Item appeared on the ground.
type ItemAddServerPacket struct {
	ItemId     int
	ItemIndex  int
	ItemAmount int
	Coords     protocol.Coords
}

func (s ItemAddServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemAddServerPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

func (s *ItemAddServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// ItemIndex : field : short
	if err = writer.AddShort(s.ItemIndex); err != nil {
		return
	}
	// ItemAmount : field : three
	if err = writer.AddThree(s.ItemAmount); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemAddServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()
	// ItemAmount : field : three
	s.ItemAmount = reader.GetThree()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemRemoveServerPacket :: Item disappeared from the ground.
type ItemRemoveServerPacket struct {
	ItemIndex int
}

func (s ItemRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *ItemRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemIndex : field : short
	if err = writer.AddShort(s.ItemIndex); err != nil {
		return
	}
	return
}

func (s *ItemRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()

	return
}

// ItemJunkServerPacket :: Reply to junking items.
type ItemJunkServerPacket struct {
	JunkedItem      net.ThreeItem
	RemainingAmount int
	Weight          net.Weight
}

func (s ItemJunkServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemJunkServerPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

func (s *ItemJunkServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// JunkedItem : field : ThreeItem
	if err = s.JunkedItem.Serialize(writer); err != nil {
		return
	}
	// RemainingAmount : field : int
	if err = writer.AddInt(s.RemainingAmount); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemJunkServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// JunkedItem : field : ThreeItem
	if err = s.JunkedItem.Deserialize(reader); err != nil {
		return
	}
	// RemainingAmount : field : int
	s.RemainingAmount = reader.GetInt()
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemGetServerPacket :: Reply to taking items from the ground.
type ItemGetServerPacket struct {
	TakenItemIndex int
	TakenItem      net.ThreeItem
	Weight         net.Weight
}

func (s ItemGetServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemGetServerPacket) Action() net.PacketAction {
	return net.PacketAction_Get
}

func (s *ItemGetServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TakenItemIndex : field : short
	if err = writer.AddShort(s.TakenItemIndex); err != nil {
		return
	}
	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemGetServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TakenItemIndex : field : short
	s.TakenItemIndex = reader.GetShort()
	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemObtainServerPacket :: Receive item (from quest).
type ItemObtainServerPacket struct {
	Item          net.ThreeItem
	CurrentWeight int
}

func (s ItemObtainServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemObtainServerPacket) Action() net.PacketAction {
	return net.PacketAction_Obtain
}

func (s *ItemObtainServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Item : field : ThreeItem
	if err = s.Item.Serialize(writer); err != nil {
		return
	}
	// CurrentWeight : field : char
	if err = writer.AddChar(s.CurrentWeight); err != nil {
		return
	}
	return
}

func (s *ItemObtainServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Item : field : ThreeItem
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// CurrentWeight : field : char
	s.CurrentWeight = reader.GetChar()

	return
}

// ItemKickServerPacket :: Lose item (from quest).
type ItemKickServerPacket struct {
	Item          net.Item
	CurrentWeight int
}

func (s ItemKickServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemKickServerPacket) Action() net.PacketAction {
	return net.PacketAction_Kick
}

func (s *ItemKickServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Item : field : Item
	if err = s.Item.Serialize(writer); err != nil {
		return
	}
	// CurrentWeight : field : char
	if err = writer.AddChar(s.CurrentWeight); err != nil {
		return
	}
	return
}

func (s *ItemKickServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Item : field : Item
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// CurrentWeight : field : char
	s.CurrentWeight = reader.GetChar()

	return
}

// ItemAgreeServerPacket :: Reply to using an item that you don't have.
type ItemAgreeServerPacket struct {
	ItemId int
}

func (s ItemAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *ItemAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	return
}

func (s *ItemAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()

	return
}

// ItemSpecServerPacket :: Reply to trying to take a protected item from the ground.
type ItemSpecServerPacket struct {
}

func (s ItemSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *ItemSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}
	return
}

func (s *ItemSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : short
	reader.GetShort()

	return
}

// BoardPlayerServerPacket :: Reply to reading a post on a town board.
type BoardPlayerServerPacket struct {
	PostId   int
	PostBody string
}

func (s BoardPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *BoardPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PostId : field : short
	if err = writer.AddShort(s.PostId); err != nil {
		return
	}
	// PostBody : field : string
	if err = writer.AddString(s.PostBody); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *BoardPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PostId : field : short
	s.PostId = reader.GetShort()
	// PostBody : field : string
	if s.PostBody, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)

	return
}

// BoardOpenServerPacket :: Reply to opening a town board.
type BoardOpenServerPacket struct {
	BoardId    int
	PostsCount int
	Posts      []BoardPostListing
}

func (s BoardOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *BoardOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// BoardId : field : char
	if err = writer.AddChar(s.BoardId); err != nil {
		return
	}
	// PostsCount : length : char
	if err = writer.AddChar(s.PostsCount); err != nil {
		return
	}
	// Posts : array : BoardPostListing
	for ndx := 0; ndx < s.PostsCount; ndx++ {
		if err = s.Posts[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *BoardOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// BoardId : field : char
	s.BoardId = reader.GetChar()
	// PostsCount : length : char
	s.PostsCount = reader.GetChar()
	// Posts : array : BoardPostListing
	for ndx := 0; ndx < s.PostsCount; ndx++ {
		s.Posts = append(s.Posts, BoardPostListing{})
		if err = s.Posts[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// JukeboxAgreeServerPacket :: Reply to successfully requesting a song.
type JukeboxAgreeServerPacket struct {
	GoldAmount int
}

func (s JukeboxAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *JukeboxAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *JukeboxAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// JukeboxReplyServerPacket :: Reply to unsuccessfully requesting a song.
type JukeboxReplyServerPacket struct {
}

func (s JukeboxReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *JukeboxReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : short
	if err = writer.AddShort(1); err != nil {
		return
	}
	return
}

func (s *JukeboxReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : short
	reader.GetShort()

	return
}

// JukeboxOpenServerPacket :: Reply to opening the jukebox listing.
type JukeboxOpenServerPacket struct {
	MapId         int
	JukeboxPlayer string
}

func (s JukeboxOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *JukeboxOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	// JukeboxPlayer : field : string
	if err = writer.AddString(s.JukeboxPlayer); err != nil {
		return
	}
	return
}

func (s *JukeboxOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapId : field : short
	s.MapId = reader.GetShort()
	// JukeboxPlayer : field : string
	if s.JukeboxPlayer, err = reader.GetString(); err != nil {
		return
	}

	return
}

// JukeboxMsgServerPacket :: Someone playing a note with the bard skill nearby.
type JukeboxMsgServerPacket struct {
	PlayerId     int
	Direction    protocol.Direction
	InstrumentId int
	NoteId       int
}

func (s JukeboxMsgServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxMsgServerPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

func (s *JukeboxMsgServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	// InstrumentId : field : char
	if err = writer.AddChar(s.InstrumentId); err != nil {
		return
	}
	// NoteId : field : char
	if err = writer.AddChar(s.NoteId); err != nil {
		return
	}
	return
}

func (s *JukeboxMsgServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// InstrumentId : field : char
	s.InstrumentId = reader.GetChar()
	// NoteId : field : char
	s.NoteId = reader.GetChar()

	return
}

// JukeboxPlayerServerPacket :: Play background music.
type JukeboxPlayerServerPacket struct {
	MfxId int
}

func (s JukeboxPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *JukeboxPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MfxId : field : char
	if err = writer.AddChar(s.MfxId); err != nil {
		return
	}
	return
}

func (s *JukeboxPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MfxId : field : char
	s.MfxId = reader.GetChar()

	return
}

// JukeboxUseServerPacket :: Play jukebox music.
type JukeboxUseServerPacket struct {
	TrackId int
}

func (s JukeboxUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

func (s *JukeboxUseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TrackId : field : short
	if err = writer.AddShort(s.TrackId); err != nil {
		return
	}
	return
}

func (s *JukeboxUseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TrackId : field : short
	s.TrackId = reader.GetShort()

	return
}

// WarpRequestServerPacket :: Warp request from server.
type WarpRequestServerPacket struct {
	WarpType     WarpType
	MapId        int
	WarpTypeData WarpRequestWarpTypeData
	SessionId    int
}

type WarpRequestWarpTypeData interface {
	protocol.EoData
}

type WarpRequestWarpTypeDataMapSwitch struct {
	MapRid      []int
	MapFileSize int
}

func (s *WarpRequestWarpTypeDataMapSwitch) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.MapRid[ndx]); err != nil {
			return
		}
	}

	// MapFileSize : field : three
	if err = writer.AddThree(s.MapFileSize); err != nil {
		return
	}
	return
}

func (s *WarpRequestWarpTypeDataMapSwitch) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.MapRid = append(s.MapRid, 0)
		s.MapRid[ndx] = reader.GetShort()
	}

	// MapFileSize : field : three
	s.MapFileSize = reader.GetThree()

	return
}

func (s WarpRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *WarpRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WarpType : field : WarpType
	if err = writer.AddChar(int(s.WarpType)); err != nil {
		return
	}
	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	switch s.WarpType {
	case Warp_MapSwitch:
		switch s.WarpTypeData.(type) {
		case *WarpRequestWarpTypeDataMapSwitch:
			if err = s.WarpTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.WarpType)
			return
		}
	}
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	return
}

func (s *WarpRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// WarpType : field : WarpType
	s.WarpType = WarpType(reader.GetChar())
	// MapId : field : short
	s.MapId = reader.GetShort()
	switch s.WarpType {
	case Warp_MapSwitch:
		s.WarpTypeData = &WarpRequestWarpTypeDataMapSwitch{}
		if err = s.WarpTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	// SessionId : field : short
	s.SessionId = reader.GetShort()

	return
}

// WarpAgreeServerPacket :: Reply after accepting a warp.
type WarpAgreeServerPacket struct {
	WarpType     WarpType
	WarpTypeData WarpAgreeWarpTypeData
	Nearby       NearbyInfo
}

type WarpAgreeWarpTypeData interface {
	protocol.EoData
}

type WarpAgreeWarpTypeDataMapSwitch struct {
	MapId      int
	WarpEffect WarpEffect
}

func (s *WarpAgreeWarpTypeDataMapSwitch) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	// WarpEffect : field : WarpEffect
	if err = writer.AddChar(int(s.WarpEffect)); err != nil {
		return
	}
	return
}

func (s *WarpAgreeWarpTypeDataMapSwitch) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapId : field : short
	s.MapId = reader.GetShort()
	// WarpEffect : field : WarpEffect
	s.WarpEffect = WarpEffect(reader.GetChar())

	return
}

func (s WarpAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *WarpAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// WarpType : field : WarpType
	if err = writer.AddChar(int(s.WarpType)); err != nil {
		return
	}
	switch s.WarpType {
	case Warp_MapSwitch:
		switch s.WarpTypeData.(type) {
		case *WarpAgreeWarpTypeDataMapSwitch:
			if err = s.WarpTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.WarpType)
			return
		}
	}
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *WarpAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// WarpType : field : WarpType
	s.WarpType = WarpType(reader.GetChar())
	switch s.WarpType {
	case Warp_MapSwitch:
		s.WarpTypeData = &WarpAgreeWarpTypeDataMapSwitch{}
		if err = s.WarpTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// PaperdollReplyServerPacket :: Reply to requesting a paperdoll.
type PaperdollReplyServerPacket struct {
	Details   CharacterDetails
	Equipment EquipmentPaperdoll
	Icon      CharacterIcon
}

func (s PaperdollReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *PaperdollReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Details : field : CharacterDetails
	if err = s.Details.Serialize(writer); err != nil {
		return
	}
	// Equipment : field : EquipmentPaperdoll
	if err = s.Equipment.Serialize(writer); err != nil {
		return
	}
	// Icon : field : CharacterIcon
	if err = writer.AddChar(int(s.Icon)); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *PaperdollReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Details : field : CharacterDetails
	if err = s.Details.Deserialize(reader); err != nil {
		return
	}
	// Equipment : field : EquipmentPaperdoll
	if err = s.Equipment.Deserialize(reader); err != nil {
		return
	}
	// Icon : field : CharacterIcon
	s.Icon = CharacterIcon(reader.GetChar())
	reader.SetIsChunked(false)

	return
}

// PaperdollPingServerPacket :: Failed to equip an item due to being the incorrect class.
type PaperdollPingServerPacket struct {
	ClassId int // The player's current class ID (not the item's required class ID).
}

func (s PaperdollPingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollPingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

func (s *PaperdollPingServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	return
}

func (s *PaperdollPingServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ClassId : field : char
	s.ClassId = reader.GetChar()

	return
}

// PaperdollRemoveServerPacket :: Reply to unequipping an item.
type PaperdollRemoveServerPacket struct {
	Change AvatarChange
	ItemId int
	SubLoc int
	Stats  CharacterStatsEquipmentChange
}

func (s PaperdollRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *PaperdollRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Change : field : AvatarChange
	if err = s.Change.Serialize(writer); err != nil {
		return
	}
	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// SubLoc : field : char
	if err = writer.AddChar(s.SubLoc); err != nil {
		return
	}
	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *PaperdollRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()
	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

// PaperdollAgreeServerPacket :: Reply to equipping an item.
type PaperdollAgreeServerPacket struct {
	Change          AvatarChange
	ItemId          int
	RemainingAmount int
	SubLoc          int
	Stats           CharacterStatsEquipmentChange
}

func (s PaperdollAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *PaperdollAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Change : field : AvatarChange
	if err = s.Change.Serialize(writer); err != nil {
		return
	}
	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// RemainingAmount : field : three
	if err = writer.AddThree(s.RemainingAmount); err != nil {
		return
	}
	// SubLoc : field : char
	if err = writer.AddChar(s.SubLoc); err != nil {
		return
	}
	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *PaperdollAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// RemainingAmount : field : three
	s.RemainingAmount = reader.GetThree()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()
	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

// AvatarAgreeServerPacket :: Nearby player changed appearance.
type AvatarAgreeServerPacket struct {
	Change AvatarChange
}

func (s AvatarAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *AvatarAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Change : field : AvatarChange
	if err = s.Change.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *AvatarAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}

	return
}

// BookReplyServerPacket :: Reply to requesting a book.
type BookReplyServerPacket struct {
	Details    CharacterDetails
	Icon       CharacterIcon
	QuestNames []string
}

func (s BookReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Book
}

func (s BookReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *BookReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Details : field : CharacterDetails
	if err = s.Details.Serialize(writer); err != nil {
		return
	}
	// Icon : field : CharacterIcon
	if err = writer.AddChar(int(s.Icon)); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// QuestNames : array : string
	for ndx := 0; ndx < len(s.QuestNames); ndx++ {
		if err = writer.AddString(s.QuestNames[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *BookReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Details : field : CharacterDetails
	if err = s.Details.Deserialize(reader); err != nil {
		return
	}
	// Icon : field : CharacterIcon
	s.Icon = CharacterIcon(reader.GetChar())
	if err = reader.NextChunk(); err != nil {
		return
	}
	// QuestNames : array : string
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.QuestNames = append(s.QuestNames, "")
		if s.QuestNames[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// MessagePongServerPacket :: #ping command reply.
type MessagePongServerPacket struct {
}

func (s MessagePongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessagePongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

func (s *MessagePongServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}
	return
}

func (s *MessagePongServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : short
	reader.GetShort()

	return
}

// PlayersPingServerPacket :: #find command reply - offline.
type PlayersPingServerPacket struct {
	Name string
}

func (s PlayersPingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersPingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

func (s *PlayersPingServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *PlayersPingServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

// PlayersPongServerPacket :: #find command reply - same map.
type PlayersPongServerPacket struct {
	Name string
}

func (s PlayersPongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersPongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

func (s *PlayersPongServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *PlayersPongServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

// PlayersNet242ServerPacket :: #find command reply - different map.
type PlayersNet242ServerPacket struct {
	Name string
}

func (s PlayersNet242ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersNet242ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net242
}

func (s *PlayersNet242ServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *PlayersNet242ServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

// DoorOpenServerPacket :: Nearby door opening.
type DoorOpenServerPacket struct {
	Coords protocol.Coords
}

func (s DoorOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Door
}

func (s DoorOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *DoorOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *DoorOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	//  : field : char
	reader.GetChar()

	return
}

// DoorCloseServerPacket :: Reply to trying to open a locked door.
type DoorCloseServerPacket struct {
	Key int
}

func (s DoorCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Door
}

func (s DoorCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *DoorCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Key : field : char
	if err = writer.AddChar(s.Key); err != nil {
		return
	}
	return
}

func (s *DoorCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Key : field : char
	s.Key = reader.GetChar()

	return
}

// ChestOpenServerPacket :: Reply to opening a chest.
type ChestOpenServerPacket struct {
	Coords protocol.Coords
	Items  []net.ThreeItem
}

func (s ChestOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *ChestOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ChestOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.ThreeItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ChestReplyServerPacket :: Reply to placing an item in to a chest.
type ChestReplyServerPacket struct {
	AddedItemId     int
	RemainingAmount int
	Weight          net.Weight
	Items           []net.ThreeItem
}

func (s ChestReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *ChestReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// AddedItemId : field : short
	if err = writer.AddShort(s.AddedItemId); err != nil {
		return
	}
	// RemainingAmount : field : int
	if err = writer.AddInt(s.RemainingAmount); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ChestReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// AddedItemId : field : short
	s.AddedItemId = reader.GetShort()
	// RemainingAmount : field : int
	s.RemainingAmount = reader.GetInt()
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.ThreeItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ChestGetServerPacket :: Reply to removing an item from a chest.
type ChestGetServerPacket struct {
	TakenItem net.ThreeItem
	Weight    net.Weight
	Items     []net.ThreeItem
}

func (s ChestGetServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestGetServerPacket) Action() net.PacketAction {
	return net.PacketAction_Get
}

func (s *ChestGetServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Serialize(writer); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Serialize(writer); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ChestGetServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TakenItem : field : ThreeItem
	if err = s.TakenItem.Deserialize(reader); err != nil {
		return
	}
	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// Items : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.ThreeItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ChestAgreeServerPacket :: Chest contents updating.
type ChestAgreeServerPacket struct {
	Items []net.ThreeItem
}

func (s ChestAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *ChestAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Items : array : ThreeItem
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ChestAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Items : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.ThreeItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ChestSpecServerPacket :: Reply to trying to add an item to a full chest.
type ChestSpecServerPacket struct {
}

func (s ChestSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *ChestSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(0); err != nil {
		return
	}
	return
}

func (s *ChestSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// ChestCloseServerPacket ::  Reply to trying to interact with a locked or "broken" chest. The official client assumes a broken chest if the packet is under 2 bytes in length.
type ChestCloseServerPacket struct {
	Key *int // Sent if the player is trying to interact with a locked chest.

}

func (s ChestCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *ChestCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Key : field : short
	if s.Key != nil {
		if err = writer.AddShort(*s.Key); err != nil {
			return
		}
	}
	//  : dummy : string
	if err = writer.AddString("N"); err != nil {
		return
	}
	return
}

func (s *ChestCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Key : field : short
	if reader.Remaining() > 0 {
		s.Key = new(int)
		*s.Key = reader.GetShort()
	}
	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// RefreshReplyServerPacket :: Reply to request for new info about nearby objects.
type RefreshReplyServerPacket struct {
	Nearby NearbyInfo
}

func (s RefreshReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Refresh
}

func (s RefreshReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *RefreshReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *RefreshReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}

	return
}

// PartyRequestServerPacket :: Received party invite / join request.
type PartyRequestServerPacket struct {
	RequestType     net.PartyRequestType
	InviterPlayerId int
	PlayerName      string
}

func (s PartyRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *PartyRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// RequestType : field : PartyRequestType
	if err = writer.AddChar(int(s.RequestType)); err != nil {
		return
	}
	// InviterPlayerId : field : short
	if err = writer.AddShort(s.InviterPlayerId); err != nil {
		return
	}
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	return
}

func (s *PartyRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// PartyReplyServerPacket :: Failed party invite / join request.
type PartyReplyServerPacket struct {
	ReplyCode     PartyReplyCode
	ReplyCodeData PartyReplyReplyCodeData
}

type PartyReplyReplyCodeData interface {
	protocol.EoData
}

type PartyReplyReplyCodeDataAlreadyInAnotherParty struct {
	PlayerName string
}

func (s *PartyReplyReplyCodeDataAlreadyInAnotherParty) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	return
}

func (s *PartyReplyReplyCodeDataAlreadyInAnotherParty) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	return
}

type PartyReplyReplyCodeDataAlreadyInYourParty struct {
	PlayerName string
}

func (s *PartyReplyReplyCodeDataAlreadyInYourParty) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	return
}

func (s *PartyReplyReplyCodeDataAlreadyInYourParty) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	return
}

func (s PartyReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *PartyReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : PartyReplyCode
	if err = writer.AddChar(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case PartyReplyCode_AlreadyInAnotherParty:
		switch s.ReplyCodeData.(type) {
		case *PartyReplyReplyCodeDataAlreadyInAnotherParty:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case PartyReplyCode_AlreadyInYourParty:
		switch s.ReplyCodeData.(type) {
		case *PartyReplyReplyCodeDataAlreadyInYourParty:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *PartyReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : PartyReplyCode
	s.ReplyCode = PartyReplyCode(reader.GetChar())
	switch s.ReplyCode {
	case PartyReplyCode_AlreadyInAnotherParty:
		s.ReplyCodeData = &PartyReplyReplyCodeDataAlreadyInAnotherParty{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case PartyReplyCode_AlreadyInYourParty:
		s.ReplyCodeData = &PartyReplyReplyCodeDataAlreadyInYourParty{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// PartyCreateServerPacket :: Member list received when party is first joined.
type PartyCreateServerPacket struct {
	Members []PartyMember
}

func (s PartyCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

func (s *PartyCreateServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Members : array : PartyMember
	for ndx := 0; ndx < len(s.Members); ndx++ {
		if err = s.Members[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PartyCreateServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Members : array : PartyMember
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Members = append(s.Members, PartyMember{})
		if err = s.Members[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// PartyAddServerPacket :: New player joined the party.
type PartyAddServerPacket struct {
	Member PartyMember
}

func (s PartyAddServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyAddServerPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

func (s *PartyAddServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Member : field : PartyMember
	if err = s.Member.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *PartyAddServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Member : field : PartyMember
	if err = s.Member.Deserialize(reader); err != nil {
		return
	}

	return
}

// PartyRemoveServerPacket :: Player left the party.
type PartyRemoveServerPacket struct {
	PlayerId int
}

func (s PartyRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

func (s *PartyRemoveServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *PartyRemoveServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// PartyCloseServerPacket :: Left / disbanded a party.
type PartyCloseServerPacket struct {
}

func (s PartyCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *PartyCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *PartyCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// PartyListServerPacket :: Party member list update.
type PartyListServerPacket struct {
	Members []PartyMember
}

func (s PartyListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *PartyListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Members : array : PartyMember
	for ndx := 0; ndx < len(s.Members); ndx++ {
		if err = s.Members[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PartyListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Members : array : PartyMember
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Members = append(s.Members, PartyMember{})
		if err = s.Members[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// PartyAgreeServerPacket :: Party member list update.
type PartyAgreeServerPacket struct {
	PlayerId     int
	HpPercentage int
}

func (s PartyAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *PartyAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	return
}

func (s *PartyAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()

	return
}

// PartyTargetGroupServerPacket :: Updated experience and level-ups from party experience.
type PartyTargetGroupServerPacket struct {
	Gains []PartyExpShare
}

func (s PartyTargetGroupServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyTargetGroupServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetGroup
}

func (s *PartyTargetGroupServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Gains : array : PartyExpShare
	for ndx := 0; ndx < len(s.Gains); ndx++ {
		if err = s.Gains[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *PartyTargetGroupServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Gains : array : PartyExpShare
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Gains = append(s.Gains, PartyExpShare{})
		if err = s.Gains[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// GuildReplyServerPacket :: Generic guild reply messages.
type GuildReplyServerPacket struct {
	ReplyCode     GuildReply
	ReplyCodeData GuildReplyReplyCodeData
}

type GuildReplyReplyCodeData interface {
	protocol.EoData
}

type GuildReplyReplyCodeDataCreateAdd struct {
	Name string
}

func (s *GuildReplyReplyCodeDataCreateAdd) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *GuildReplyReplyCodeDataCreateAdd) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

type GuildReplyReplyCodeDataCreateAddConfirm struct {
	Name string
}

func (s *GuildReplyReplyCodeDataCreateAddConfirm) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *GuildReplyReplyCodeDataCreateAddConfirm) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

type GuildReplyReplyCodeDataJoinRequest struct {
	PlayerId int
	Name     string
}

func (s *GuildReplyReplyCodeDataJoinRequest) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *GuildReplyReplyCodeDataJoinRequest) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

func (s GuildReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *GuildReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : GuildReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case GuildReply_CreateAdd:
		switch s.ReplyCodeData.(type) {
		case *GuildReplyReplyCodeDataCreateAdd:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case GuildReply_CreateAddConfirm:
		switch s.ReplyCodeData.(type) {
		case *GuildReplyReplyCodeDataCreateAddConfirm:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	case GuildReply_JoinRequest:
		switch s.ReplyCodeData.(type) {
		case *GuildReplyReplyCodeDataJoinRequest:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *GuildReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : GuildReply
	s.ReplyCode = GuildReply(reader.GetShort())
	switch s.ReplyCode {
	case GuildReply_CreateAdd:
		s.ReplyCodeData = &GuildReplyReplyCodeDataCreateAdd{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case GuildReply_CreateAddConfirm:
		s.ReplyCodeData = &GuildReplyReplyCodeDataCreateAddConfirm{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	case GuildReply_JoinRequest:
		s.ReplyCodeData = &GuildReplyReplyCodeDataJoinRequest{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// GuildRequestServerPacket :: Guild create request.
type GuildRequestServerPacket struct {
	PlayerId      int
	GuildIdentity string
}

func (s GuildRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *GuildRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// GuildIdentity : field : string
	if err = writer.AddString(s.GuildIdentity); err != nil {
		return
	}
	return
}

func (s *GuildRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GuildCreateServerPacket :: Guild created.
type GuildCreateServerPacket struct {
	LeaderPlayerId int
	GuildTag       string
	GuildName      string
	RankName       string
	GoldAmount     int
}

func (s GuildCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

func (s *GuildCreateServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// LeaderPlayerId : field : short
	if err = writer.AddShort(s.LeaderPlayerId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// RankName : field : string
	if err = writer.AddString(s.RankName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *GuildCreateServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// LeaderPlayerId : field : short
	s.LeaderPlayerId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// GuildName : field : string
	if s.GuildName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// RankName : field : string
	if s.RankName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	reader.SetIsChunked(false)

	return
}

// GuildTakeServerPacket :: Get guild description reply.
type GuildTakeServerPacket struct {
	Description string
}

func (s GuildTakeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTakeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

func (s *GuildTakeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	return
}

func (s *GuildTakeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GuildRankServerPacket :: Get guild rank list reply.
type GuildRankServerPacket struct {
	Ranks []string
}

func (s GuildRankServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRankServerPacket) Action() net.PacketAction {
	return net.PacketAction_Rank
}

func (s *GuildRankServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.Ranks[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildRankServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		s.Ranks = append(s.Ranks, "")
		if s.Ranks[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// GuildSellServerPacket :: Get guild bank reply.
type GuildSellServerPacket struct {
	GoldAmount int
}

func (s GuildSellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildSellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Sell
}

func (s *GuildSellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *GuildSellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// GuildBuyServerPacket :: Deposit guild bank list reply.
type GuildBuyServerPacket struct {
	GoldAmount int
}

func (s GuildBuyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildBuyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

func (s *GuildBuyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *GuildBuyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// GuildOpenServerPacket :: Talk to guild master NPC reply.
type GuildOpenServerPacket struct {
	SessionId int
}

func (s GuildOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *GuildOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}
	return
}

func (s *GuildOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : three
	s.SessionId = reader.GetThree()

	return
}

// GuildTellServerPacket :: Get guild member list reply.
type GuildTellServerPacket struct {
	MembersCount int
	Members      []GuildMember
}

func (s GuildTellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

func (s *GuildTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// MembersCount : length : short
	if err = writer.AddShort(s.MembersCount); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Members : array : GuildMember
	for ndx := 0; ndx < s.MembersCount; ndx++ {
		if err = s.Members[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildTellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// MembersCount : length : short
	s.MembersCount = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Members : array : GuildMember
	for ndx := 0; ndx < s.MembersCount; ndx++ {
		s.Members = append(s.Members, GuildMember{})
		if err = s.Members[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// GuildReportServerPacket :: Get guild info reply.
type GuildReportServerPacket struct {
	Name        string
	Tag         string
	CreateDate  string
	Description string
	Wealth      string
	Ranks       []string
	StaffCount  int
	Staff       []GuildStaff
}

func (s GuildReportServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildReportServerPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

func (s *GuildReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Tag : field : string
	if err = writer.AddString(s.Tag); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// CreateDate : field : string
	if err = writer.AddString(s.CreateDate); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Wealth : field : string
	if err = writer.AddString(s.Wealth); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.Ranks[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	// StaffCount : length : short
	if err = writer.AddShort(s.StaffCount); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Staff : array : GuildStaff
	for ndx := 0; ndx < s.StaffCount; ndx++ {
		if err = s.Staff[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Tag : field : string
	if s.Tag, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// CreateDate : field : string
	if s.CreateDate, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Wealth : field : string
	if s.Wealth, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		s.Ranks = append(s.Ranks, "")
		if s.Ranks[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	// StaffCount : length : short
	s.StaffCount = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Staff : array : GuildStaff
	for ndx := 0; ndx < s.StaffCount; ndx++ {
		s.Staff = append(s.Staff, GuildStaff{})
		if err = s.Staff[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// GuildAgreeServerPacket :: Joined guild info.
type GuildAgreeServerPacket struct {
	RecruiterId int
	GuildTag    string
	GuildName   string
	RankName    string
}

func (s GuildAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *GuildAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// RecruiterId : field : short
	if err = writer.AddShort(s.RecruiterId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// RankName : field : string
	if err = writer.AddString(s.RankName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *GuildAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// RecruiterId : field : short
	s.RecruiterId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// GuildName : field : string
	if s.GuildName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// RankName : field : string
	if s.RankName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// GuildAcceptServerPacket :: Update guild rank.
type GuildAcceptServerPacket struct {
	Rank int
}

func (s GuildAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *GuildAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Rank : field : char
	if err = writer.AddChar(s.Rank); err != nil {
		return
	}
	return
}

func (s *GuildAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Rank : field : char
	s.Rank = reader.GetChar()

	return
}

// GuildKickServerPacket :: Left the guild.
type GuildKickServerPacket struct {
}

func (s GuildKickServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildKickServerPacket) Action() net.PacketAction {
	return net.PacketAction_Kick
}

func (s *GuildKickServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *GuildKickServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// SpellRequestServerPacket :: Nearby player chanting a spell.
type SpellRequestServerPacket struct {
	PlayerId int
	SpellId  int
}

func (s SpellRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *SpellRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	return
}

func (s *SpellRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

// SpellTargetSelfServerPacket :: Nearby player self-casted a spell.
type SpellTargetSelfServerPacket struct {
	PlayerId     int
	SpellId      int
	SpellHealHp  int
	HpPercentage int
	Hp           *int // The official client reads this if the packet is larger than 12 bytes.
	Tp           *int // The official client reads this if the packet is larger than 12 bytes.
}

func (s SpellTargetSelfServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetSelfServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetSelf
}

func (s *SpellTargetSelfServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// SpellHealHp : field : int
	if err = writer.AddInt(s.SpellHealHp); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Hp : field : short
	if s.Hp != nil {
		if err = writer.AddShort(*s.Hp); err != nil {
			return
		}
	}
	// Tp : field : short
	if s.Tp != nil {
		if err = writer.AddShort(*s.Tp); err != nil {
			return
		}
	}
	return
}

func (s *SpellTargetSelfServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// SpellHealHp : field : int
	s.SpellHealHp = reader.GetInt()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Hp : field : short
	if reader.Remaining() > 0 {
		s.Hp = new(int)
		*s.Hp = reader.GetShort()
	}
	// Tp : field : short
	if reader.Remaining() > 0 {
		s.Tp = new(int)
		*s.Tp = reader.GetShort()
	}

	return
}

// SpellPlayerServerPacket :: Nearby player raising their arm to cast a spell (vestigial).
type SpellPlayerServerPacket struct {
	PlayerId  int
	Direction protocol.Direction
}

func (s SpellPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *SpellPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *SpellPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// SpellErrorServerPacket :: Show flood protection message (vestigial).
type SpellErrorServerPacket struct {
}

func (s SpellErrorServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellErrorServerPacket) Action() net.PacketAction {
	return net.PacketAction_Error
}

func (s *SpellErrorServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *SpellErrorServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// AvatarAdminServerPacket :: Nearby player hit by a damage spell from a player.
type AvatarAdminServerPacket struct {
	CasterId        int
	VictimId        int
	CasterDirection protocol.Direction
	Damage          int
	HpPercentage    int
	VictimDied      bool
	SpellId         int
}

func (s AvatarAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

func (s *AvatarAdminServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CasterId : field : short
	if err = writer.AddShort(s.CasterId); err != nil {
		return
	}
	// VictimId : field : short
	if err = writer.AddShort(s.VictimId); err != nil {
		return
	}
	// CasterDirection : field : Direction
	if err = writer.AddChar(int(s.CasterDirection)); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// VictimDied : field : bool
	if s.VictimDied {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	return
}

func (s *AvatarAdminServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CasterId : field : short
	s.CasterId = reader.GetShort()
	// VictimId : field : short
	s.VictimId = reader.GetShort()
	// CasterDirection : field : Direction
	s.CasterDirection = protocol.Direction(reader.GetChar())
	// Damage : field : three
	s.Damage = reader.GetThree()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// VictimDied : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.VictimDied = true
	} else {
		s.VictimDied = false
	}
	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

// SpellTargetGroupServerPacket :: Nearby player(s) hit by a group heal spell from a player.
type SpellTargetGroupServerPacket struct {
	SpellId     int
	CasterId    int
	CasterTp    int
	SpellHealHp int
	Players     []GroupHealTargetPlayer
}

func (s SpellTargetGroupServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetGroupServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetGroup
}

func (s *SpellTargetGroupServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// CasterId : field : short
	if err = writer.AddShort(s.CasterId); err != nil {
		return
	}
	// CasterTp : field : short
	if err = writer.AddShort(s.CasterTp); err != nil {
		return
	}
	// SpellHealHp : field : short
	if err = writer.AddShort(s.SpellHealHp); err != nil {
		return
	}
	// Players : array : GroupHealTargetPlayer
	for ndx := 0; ndx < len(s.Players); ndx++ {
		if err = s.Players[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *SpellTargetGroupServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// CasterId : field : short
	s.CasterId = reader.GetShort()
	// CasterTp : field : short
	s.CasterTp = reader.GetShort()
	// SpellHealHp : field : short
	s.SpellHealHp = reader.GetShort()
	// Players : array : GroupHealTargetPlayer
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Players = append(s.Players, GroupHealTargetPlayer{})
		if err = s.Players[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// SpellTargetOtherServerPacket :: Nearby player hit by a heal spell from a player.
type SpellTargetOtherServerPacket struct {
	VictimId        int
	CasterId        int
	CasterDirection protocol.Direction
	SpellId         int
	SpellHealHp     int
	HpPercentage    int
	Hp              *int
}

func (s SpellTargetOtherServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetOtherServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetOther
}

func (s *SpellTargetOtherServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// VictimId : field : short
	if err = writer.AddShort(s.VictimId); err != nil {
		return
	}
	// CasterId : field : short
	if err = writer.AddShort(s.CasterId); err != nil {
		return
	}
	// CasterDirection : field : Direction
	if err = writer.AddChar(int(s.CasterDirection)); err != nil {
		return
	}
	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// SpellHealHp : field : int
	if err = writer.AddInt(s.SpellHealHp); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Hp : field : short
	if s.Hp != nil {
		if err = writer.AddShort(*s.Hp); err != nil {
			return
		}
	}
	return
}

func (s *SpellTargetOtherServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// VictimId : field : short
	s.VictimId = reader.GetShort()
	// CasterId : field : short
	s.CasterId = reader.GetShort()
	// CasterDirection : field : Direction
	s.CasterDirection = protocol.Direction(reader.GetChar())
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// SpellHealHp : field : int
	s.SpellHealHp = reader.GetInt()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Hp : field : short
	if reader.Remaining() > 0 {
		s.Hp = new(int)
		*s.Hp = reader.GetShort()
	}

	return
}

// TradeRequestServerPacket :: Trade request from another player.
type TradeRequestServerPacket struct {
	PartnerPlayerId   int
	PartnerPlayerName string
}

func (s TradeRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *TradeRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : char
	if err = writer.AddChar(138); err != nil {
		return
	}
	// PartnerPlayerId : field : short
	if err = writer.AddShort(s.PartnerPlayerId); err != nil {
		return
	}
	// PartnerPlayerName : field : string
	if err = writer.AddString(s.PartnerPlayerName); err != nil {
		return
	}
	return
}

func (s *TradeRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : char
	reader.GetChar()
	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	// PartnerPlayerName : field : string
	if s.PartnerPlayerName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TradeOpenServerPacket :: Trade window opens.
type TradeOpenServerPacket struct {
	PartnerPlayerId   int
	PartnerPlayerName string
	YourPlayerId      int
	YourPlayerName    string
}

func (s TradeOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *TradeOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PartnerPlayerId : field : short
	if err = writer.AddShort(s.PartnerPlayerId); err != nil {
		return
	}
	// PartnerPlayerName : field : string
	if err = writer.AddString(s.PartnerPlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// YourPlayerId : field : short
	if err = writer.AddShort(s.YourPlayerId); err != nil {
		return
	}
	// YourPlayerName : field : string
	if err = writer.AddString(s.YourPlayerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *TradeOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	// PartnerPlayerName : field : string
	if s.PartnerPlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// YourPlayerId : field : short
	s.YourPlayerId = reader.GetShort()
	// YourPlayerName : field : string
	if s.YourPlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// TradeReplyServerPacket :: Trade updated (items changed).
type TradeReplyServerPacket struct {
	TradeData TradeItemData
}

func (s TradeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *TradeReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *TradeReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}

	return
}

// TradeAdminServerPacket :: Trade updated (items changed while trade was accepted).
type TradeAdminServerPacket struct {
	TradeData TradeItemData
}

func (s TradeAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

func (s *TradeAdminServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *TradeAdminServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}

	return
}

// TradeUseServerPacket :: Trade completed.
type TradeUseServerPacket struct {
	TradeData TradeItemData
}

func (s TradeUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

func (s *TradeUseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *TradeUseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}

	return
}

// TradeSpecServerPacket :: Own agree state updated.
type TradeSpecServerPacket struct {
	Agree bool
}

func (s TradeSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *TradeSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Agree : field : bool
	if s.Agree {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	return
}

func (s *TradeSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Agree : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Agree = true
	} else {
		s.Agree = false
	}

	return
}

// TradeAgreeServerPacket :: Partner agree state updated.
type TradeAgreeServerPacket struct {
	PartnerPlayerId int
	Agree           bool
}

func (s TradeAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *TradeAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PartnerPlayerId : field : short
	if err = writer.AddShort(s.PartnerPlayerId); err != nil {
		return
	}
	// Agree : field : bool
	if s.Agree {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	return
}

func (s *TradeAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	// Agree : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Agree = true
	} else {
		s.Agree = false
	}

	return
}

// TradeCloseServerPacket :: Partner closed trade window.
type TradeCloseServerPacket struct {
	PartnerPlayerId int
}

func (s TradeCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

func (s *TradeCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PartnerPlayerId : field : short
	if err = writer.AddShort(s.PartnerPlayerId); err != nil {
		return
	}
	return
}

func (s *TradeCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()

	return
}

// NpcReplyServerPacket :: Nearby NPC hit by a player.
type NpcReplyServerPacket struct {
	PlayerId            int
	PlayerDirection     protocol.Direction
	NpcIndex            int
	Damage              int
	HpPercentage        int
	KillStealProtection *NpcKillStealProtectionState
}

func (s NpcReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *NpcReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// PlayerDirection : field : Direction
	if err = writer.AddChar(int(s.PlayerDirection)); err != nil {
		return
	}
	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// HpPercentage : field : short
	if err = writer.AddShort(s.HpPercentage); err != nil {
		return
	}
	// KillStealProtection : field : NpcKillStealProtectionState
	if s.KillStealProtection != nil {
		if err = writer.AddChar(int(*s.KillStealProtection)); err != nil {
			return
		}
	}
	return
}

func (s *NpcReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// PlayerDirection : field : Direction
	s.PlayerDirection = protocol.Direction(reader.GetChar())
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// Damage : field : three
	s.Damage = reader.GetThree()
	// HpPercentage : field : short
	s.HpPercentage = reader.GetShort()
	// KillStealProtection : field : NpcKillStealProtectionState
	if reader.Remaining() > 0 {
		s.KillStealProtection = new(NpcKillStealProtectionState)
		*s.KillStealProtection = NpcKillStealProtectionState(reader.GetChar())
	}

	return
}

// CastReplyServerPacket :: Nearby NPC hit by a spell from a player.
type CastReplyServerPacket struct {
	SpellId             int
	CasterId            int
	CasterDirection     protocol.Direction
	NpcIndex            int
	Damage              int
	HpPercentage        int
	CasterTp            int
	KillStealProtection *NpcKillStealProtectionState
}

func (s CastReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Cast
}

func (s CastReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *CastReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// CasterId : field : short
	if err = writer.AddShort(s.CasterId); err != nil {
		return
	}
	// CasterDirection : field : Direction
	if err = writer.AddChar(int(s.CasterDirection)); err != nil {
		return
	}
	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// HpPercentage : field : short
	if err = writer.AddShort(s.HpPercentage); err != nil {
		return
	}
	// CasterTp : field : short
	if err = writer.AddShort(s.CasterTp); err != nil {
		return
	}
	// KillStealProtection : field : NpcKillStealProtectionState
	if s.KillStealProtection != nil {
		if err = writer.AddChar(int(*s.KillStealProtection)); err != nil {
			return
		}
	}
	return
}

func (s *CastReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// CasterId : field : short
	s.CasterId = reader.GetShort()
	// CasterDirection : field : Direction
	s.CasterDirection = protocol.Direction(reader.GetChar())
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// Damage : field : three
	s.Damage = reader.GetThree()
	// HpPercentage : field : short
	s.HpPercentage = reader.GetShort()
	// CasterTp : field : short
	s.CasterTp = reader.GetShort()
	// KillStealProtection : field : NpcKillStealProtectionState
	if reader.Remaining() > 0 {
		s.KillStealProtection = new(NpcKillStealProtectionState)
		*s.KillStealProtection = NpcKillStealProtectionState(reader.GetChar())
	}

	return
}

// NpcSpecServerPacket :: Nearby NPC killed by player.
type NpcSpecServerPacket struct {
	NpcKilledData NpcKilledData
	Experience    *int
}

func (s NpcSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *NpcSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Serialize(writer); err != nil {
		return
	}
	// Experience : field : int
	if s.Experience != nil {
		if err = writer.AddInt(*s.Experience); err != nil {
			return
		}
	}
	return
}

func (s *NpcSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Deserialize(reader); err != nil {
		return
	}
	// Experience : field : int
	if reader.Remaining() > 0 {
		s.Experience = new(int)
		*s.Experience = reader.GetInt()
	}

	return
}

// NpcAcceptServerPacket :: Nearby NPC killed by player and you leveled up.
type NpcAcceptServerPacket struct {
	NpcKilledData NpcKilledData
	Experience    int
	LevelUp       LevelUpStats
}

func (s NpcAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *NpcAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Serialize(writer); err != nil {
		return
	}
	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// LevelUp : field : LevelUpStats
	if err = s.LevelUp.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *NpcAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Deserialize(reader); err != nil {
		return
	}
	// Experience : field : int
	s.Experience = reader.GetInt()
	// LevelUp : field : LevelUpStats
	if err = s.LevelUp.Deserialize(reader); err != nil {
		return
	}

	return
}

// CastSpecServerPacket :: Nearby NPC killed by player spell.
type CastSpecServerPacket struct {
	SpellId       int
	NpcKilledData NpcKilledData
	CasterTp      int
	Experience    *int
}

func (s CastSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Cast
}

func (s CastSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *CastSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Serialize(writer); err != nil {
		return
	}
	// CasterTp : field : short
	if err = writer.AddShort(s.CasterTp); err != nil {
		return
	}
	// Experience : field : int
	if s.Experience != nil {
		if err = writer.AddInt(*s.Experience); err != nil {
			return
		}
	}
	return
}

func (s *CastSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Deserialize(reader); err != nil {
		return
	}
	// CasterTp : field : short
	s.CasterTp = reader.GetShort()
	// Experience : field : int
	if reader.Remaining() > 0 {
		s.Experience = new(int)
		*s.Experience = reader.GetInt()
	}

	return
}

// CastAcceptServerPacket :: Nearby NPC killed by player spell and you leveled up.
type CastAcceptServerPacket struct {
	SpellId       int
	NpcKilledData NpcKilledData
	CasterTp      int
	Experience    int
	LevelUp       LevelUpStats
}

func (s CastAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Cast
}

func (s CastAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *CastAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Serialize(writer); err != nil {
		return
	}
	// CasterTp : field : short
	if err = writer.AddShort(s.CasterTp); err != nil {
		return
	}
	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// LevelUp : field : LevelUpStats
	if err = s.LevelUp.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CastAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Deserialize(reader); err != nil {
		return
	}
	// CasterTp : field : short
	s.CasterTp = reader.GetShort()
	// Experience : field : int
	s.Experience = reader.GetInt()
	// LevelUp : field : LevelUpStats
	if err = s.LevelUp.Deserialize(reader); err != nil {
		return
	}

	return
}

// NpcJunkServerPacket :: Clearing all boss children.
type NpcJunkServerPacket struct {
	NpcId int
}

func (s NpcJunkServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcJunkServerPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

func (s *NpcJunkServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcId : field : short
	if err = writer.AddShort(s.NpcId); err != nil {
		return
	}
	return
}

func (s *NpcJunkServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcId : field : short
	s.NpcId = reader.GetShort()

	return
}

// NpcPlayerServerPacket :: Main NPC update message.
type NpcPlayerServerPacket struct {
	Positions []NpcUpdatePosition
	Attacks   []NpcUpdateAttack
	Chats     []NpcUpdateChat
	Hp        *int
	Tp        *int
}

func (s NpcPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *NpcPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Positions : array : NpcUpdatePosition
	for ndx := 0; ndx < len(s.Positions); ndx++ {
		if err = s.Positions[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Attacks : array : NpcUpdateAttack
	for ndx := 0; ndx < len(s.Attacks); ndx++ {
		if err = s.Attacks[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Chats : array : NpcUpdateChat
	for ndx := 0; ndx < len(s.Chats); ndx++ {
		if err = s.Chats[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(0xFF)
	// Hp : field : short
	if s.Hp != nil {
		if err = writer.AddShort(*s.Hp); err != nil {
			return
		}
	}
	// Tp : field : short
	if s.Tp != nil {
		if err = writer.AddShort(*s.Tp); err != nil {
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *NpcPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Positions : array : NpcUpdatePosition
	for ndx := 0; ndx < reader.Remaining()/4; ndx++ {
		s.Positions = append(s.Positions, NpcUpdatePosition{})
		if err = s.Positions[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Attacks : array : NpcUpdateAttack
	for ndx := 0; ndx < reader.Remaining()/9; ndx++ {
		s.Attacks = append(s.Attacks, NpcUpdateAttack{})
		if err = s.Attacks[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Chats : array : NpcUpdateChat
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Chats = append(s.Chats, NpcUpdateChat{})
		if err = s.Chats[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Hp : field : short
	if reader.Remaining() > 0 {
		s.Hp = new(int)
		*s.Hp = reader.GetShort()
	}
	// Tp : field : short
	if reader.Remaining() > 0 {
		s.Tp = new(int)
		*s.Tp = reader.GetShort()
	}
	reader.SetIsChunked(false)

	return
}

// NpcDialogServerPacket :: NPC chat message.
type NpcDialogServerPacket struct {
	NpcIndex int
	Message  string
}

func (s NpcDialogServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcDialogServerPacket) Action() net.PacketAction {
	return net.PacketAction_Dialog
}

func (s *NpcDialogServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *NpcDialogServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// QuestReportServerPacket :: NPC chat messages.
type QuestReportServerPacket struct {
	NpcId    int
	Messages []string
}

func (s QuestReportServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestReportServerPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

func (s *QuestReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// NpcId : field : short
	if err = writer.AddShort(s.NpcId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Messages : array : string
	for ndx := 0; ndx < len(s.Messages); ndx++ {
		if err = writer.AddString(s.Messages[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *QuestReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// NpcId : field : short
	s.NpcId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Messages : array : string
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Messages = append(s.Messages, "")
		if s.Messages[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// QuestDialogServerPacket :: Quest selection dialog.
type QuestDialogServerPacket struct {
	QuestCount    int
	BehaviorId    int
	QuestId       int
	SessionId     int
	DialogId      int
	QuestEntries  []DialogQuestEntry
	DialogEntries []DialogEntry
}

func (s QuestDialogServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestDialogServerPacket) Action() net.PacketAction {
	return net.PacketAction_Dialog
}

func (s *QuestDialogServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// QuestCount : length : char
	if err = writer.AddChar(s.QuestCount); err != nil {
		return
	}
	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	// QuestId : field : short
	if err = writer.AddShort(s.QuestId); err != nil {
		return
	}
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// DialogId : field : short
	if err = writer.AddShort(s.DialogId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// QuestEntries : array : DialogQuestEntry
	for ndx := 0; ndx < s.QuestCount; ndx++ {
		if err = s.QuestEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	// DialogEntries : array : DialogEntry
	for ndx := 0; ndx < len(s.DialogEntries); ndx++ {
		if err = s.DialogEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	writer.SanitizeStrings = false
	return
}

func (s *QuestDialogServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// QuestCount : length : char
	s.QuestCount = reader.GetChar()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	// QuestId : field : short
	s.QuestId = reader.GetShort()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// DialogId : field : short
	s.DialogId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// QuestEntries : array : DialogQuestEntry
	for ndx := 0; ndx < s.QuestCount; ndx++ {
		s.QuestEntries = append(s.QuestEntries, DialogQuestEntry{})
		if err = s.QuestEntries[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	// DialogEntries : array : DialogEntry
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.DialogEntries = append(s.DialogEntries, DialogEntry{})
		if err = s.DialogEntries[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)

	return
}

// QuestListServerPacket :: Quest history / progress reply.
type QuestListServerPacket struct {
	Page        net.QuestPage
	QuestsCount int
	PageData    QuestListPageData
}

type QuestListPageData interface {
	protocol.EoData
}

type QuestListPageDataProgress struct {
	QuestProgressEntries []QuestProgressEntry
}

func (s *QuestListPageDataProgress) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuestProgressEntries : array : QuestProgressEntry
	for ndx := 0; ndx < len(s.QuestProgressEntries); ndx++ {
		if err = s.QuestProgressEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	return
}

func (s *QuestListPageDataProgress) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// QuestProgressEntries : array : QuestProgressEntry
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.QuestProgressEntries = append(s.QuestProgressEntries, QuestProgressEntry{})
		if err = s.QuestProgressEntries[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

type QuestListPageDataHistory struct {
	CompletedQuests []string
}

func (s *QuestListPageDataHistory) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CompletedQuests : array : string
	for ndx := 0; ndx < len(s.CompletedQuests); ndx++ {
		if err = writer.AddString(s.CompletedQuests[ndx]); err != nil {
			return
		}
		writer.AddByte(0xFF)
	}

	return
}

func (s *QuestListPageDataHistory) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// CompletedQuests : array : string
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.CompletedQuests = append(s.CompletedQuests, "")
		if s.CompletedQuests[ndx], err = reader.GetString(); err != nil {
			return
		}

	}

	return
}

func (s QuestListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *QuestListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Page : field : QuestPage
	if err = writer.AddChar(int(s.Page)); err != nil {
		return
	}
	// QuestsCount : field : short
	if err = writer.AddShort(s.QuestsCount); err != nil {
		return
	}
	switch s.Page {
	case net.QuestPage_Progress:
		switch s.PageData.(type) {
		case *QuestListPageDataProgress:
			if err = s.PageData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.Page)
			return
		}
	case net.QuestPage_History:
		switch s.PageData.(type) {
		case *QuestListPageDataHistory:
			if err = s.PageData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.Page)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *QuestListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// Page : field : QuestPage
	s.Page = net.QuestPage(reader.GetChar())
	// QuestsCount : field : short
	s.QuestsCount = reader.GetShort()
	switch s.Page {
	case net.QuestPage_Progress:
		s.PageData = &QuestListPageDataProgress{}
		if err = s.PageData.Deserialize(reader); err != nil {
			return
		}
	case net.QuestPage_History:
		s.PageData = &QuestListPageDataHistory{}
		if err = s.PageData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetIsChunked(false)

	return
}

// ItemAcceptServerPacket :: Nearby player leveled up from quest.
type ItemAcceptServerPacket struct {
	PlayerId int
}

func (s ItemAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *ItemAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *ItemAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// ArenaDropServerPacket :: "Arena is blocked" message.
type ArenaDropServerPacket struct {
}

func (s ArenaDropServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaDropServerPacket) Action() net.PacketAction {
	return net.PacketAction_Drop
}

func (s *ArenaDropServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("N"); err != nil {
		return
	}
	return
}

func (s *ArenaDropServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// ArenaUseServerPacket :: Arena start message.
type ArenaUseServerPacket struct {
	PlayersCount int
}

func (s ArenaUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

func (s *ArenaUseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayersCount : field : char
	if err = writer.AddChar(s.PlayersCount); err != nil {
		return
	}
	return
}

func (s *ArenaUseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayersCount : field : char
	s.PlayersCount = reader.GetChar()

	return
}

// ArenaSpecServerPacket :: Arena kill message.
type ArenaSpecServerPacket struct {
	PlayerId   int
	Direction  protocol.Direction
	KillsCount int
	KillerName string
	VictimName string
}

func (s ArenaSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *ArenaSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// KillsCount : field : int
	if err = writer.AddInt(s.KillsCount); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// KillerName : field : string
	if err = writer.AddString(s.KillerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// VictimName : field : string
	if err = writer.AddString(s.VictimName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *ArenaSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	if err = reader.NextChunk(); err != nil {
		return
	}
	// KillsCount : field : int
	s.KillsCount = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// KillerName : field : string
	if s.KillerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// VictimName : field : string
	if s.VictimName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// ArenaAcceptServerPacket :: Arena win message.
type ArenaAcceptServerPacket struct {
	WinnerName string
	KillsCount int
	KillerName string
	VictimName string
}

func (s ArenaAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

func (s *ArenaAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// WinnerName : field : string
	if err = writer.AddString(s.WinnerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// KillsCount : field : int
	if err = writer.AddInt(s.KillsCount); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// KillerName : field : string
	if err = writer.AddString(s.KillerName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	// VictimName : field : string
	if err = writer.AddString(s.VictimName); err != nil {
		return
	}
	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *ArenaAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	reader.SetIsChunked(true)
	// WinnerName : field : string
	if s.WinnerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// KillsCount : field : int
	s.KillsCount = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// KillerName : field : string
	if s.KillerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// VictimName : field : string
	if s.VictimName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)

	return
}

// MarriageOpenServerPacket :: Response from talking to a law NPC.
type MarriageOpenServerPacket struct {
	SessionId int
}

func (s MarriageOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *MarriageOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}
	return
}

func (s *MarriageOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : three
	s.SessionId = reader.GetThree()

	return
}

// MarriageReplyServerPacket :: Reply to client Marriage-family packets.
type MarriageReplyServerPacket struct {
	ReplyCode     MarriageReply
	ReplyCodeData MarriageReplyReplyCodeData
}

type MarriageReplyReplyCodeData interface {
	protocol.EoData
}

type MarriageReplyReplyCodeDataSuccess struct {
	GoldAmount int
}

func (s *MarriageReplyReplyCodeDataSuccess) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}
	return
}

func (s *MarriageReplyReplyCodeDataSuccess) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

func (s MarriageReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *MarriageReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : MarriageReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	switch s.ReplyCode {
	case MarriageReply_Success:
		switch s.ReplyCodeData.(type) {
		case *MarriageReplyReplyCodeDataSuccess:
			if err = s.ReplyCodeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyCode)
			return
		}
	}
	return
}

func (s *MarriageReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : MarriageReply
	s.ReplyCode = MarriageReply(reader.GetShort())
	switch s.ReplyCode {
	case MarriageReply_Success:
		s.ReplyCodeData = &MarriageReplyReplyCodeDataSuccess{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// PriestOpenServerPacket :: Response from talking to a priest NPC.
type PriestOpenServerPacket struct {
	SessionId int
}

func (s PriestOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

func (s *PriestOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *PriestOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// PriestReplyServerPacket :: Reply to client Priest-family packets.
type PriestReplyServerPacket struct {
	ReplyCode PriestReply
}

func (s PriestReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *PriestReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ReplyCode : field : PriestReply
	if err = writer.AddShort(int(s.ReplyCode)); err != nil {
		return
	}
	return
}

func (s *PriestReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ReplyCode : field : PriestReply
	s.ReplyCode = PriestReply(reader.GetShort())

	return
}

// PriestRequestServerPacket :: Wedding request.
type PriestRequestServerPacket struct {
	SessionId   int
	PartnerName string
}

func (s PriestRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

func (s *PriestRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	// PartnerName : field : string
	if err = writer.AddString(s.PartnerName); err != nil {
		return
	}
	return
}

func (s *PriestRequestServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// PartnerName : field : string
	if s.PartnerName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// RecoverPlayerServerPacket :: HP/TP update.
type RecoverPlayerServerPacket struct {
	Hp int
	Tp int
}

func (s RecoverPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *RecoverPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	return
}

func (s *RecoverPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Hp : field : short
	s.Hp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()

	return
}

// RecoverAgreeServerPacket :: Nearby player gained HP.
type RecoverAgreeServerPacket struct {
	PlayerId     int
	HealHp       int
	HpPercentage int
}

func (s RecoverAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *RecoverAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// HealHp : field : int
	if err = writer.AddInt(s.HealHp); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	return
}

func (s *RecoverAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HealHp : field : int
	s.HealHp = reader.GetInt()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()

	return
}

// RecoverListServerPacket :: Stats update.
type RecoverListServerPacket struct {
	ClassId int
	Stats   CharacterStatsUpdate
}

func (s RecoverListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

func (s *RecoverListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ClassId : field : short
	if err = writer.AddShort(s.ClassId); err != nil {
		return
	}
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *RecoverListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ClassId : field : short
	s.ClassId = reader.GetShort()
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}

	return
}

// RecoverReplyServerPacket :: Karma/experience update.
type RecoverReplyServerPacket struct {
	Experience  int
	Karma       int
	LevelUp     *int //  A value greater than 0 is "new level" and indicates the player leveled up. The official client reads this if the packet is larger than 6 bytes.
	StatPoints  *int // The official client reads this if the player leveled up.
	SkillPoints *int // The official client reads this if the player leveled up.
}

func (s RecoverReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

func (s *RecoverReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// Karma : field : short
	if err = writer.AddShort(s.Karma); err != nil {
		return
	}
	// LevelUp : field : char
	if s.LevelUp != nil {
		if err = writer.AddChar(*s.LevelUp); err != nil {
			return
		}
	}
	// StatPoints : field : short
	if s.StatPoints != nil {
		if err = writer.AddShort(*s.StatPoints); err != nil {
			return
		}
	}
	// SkillPoints : field : short
	if s.SkillPoints != nil {
		if err = writer.AddShort(*s.SkillPoints); err != nil {
			return
		}
	}
	return
}

func (s *RecoverReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Experience : field : int
	s.Experience = reader.GetInt()
	// Karma : field : short
	s.Karma = reader.GetShort()
	// LevelUp : field : char
	if reader.Remaining() > 0 {
		s.LevelUp = new(int)
		*s.LevelUp = reader.GetChar()
	}
	// StatPoints : field : short
	if reader.Remaining() > 0 {
		s.StatPoints = new(int)
		*s.StatPoints = reader.GetShort()
	}
	// SkillPoints : field : short
	if reader.Remaining() > 0 {
		s.SkillPoints = new(int)
		*s.SkillPoints = reader.GetShort()
	}

	return
}

// RecoverTargetGroupServerPacket :: Updated stats when levelling up from party experience.
type RecoverTargetGroupServerPacket struct {
	StatPoints  int
	SkillPoints int
	MaxHp       int
	MaxTp       int
	MaxSp       int
}

func (s RecoverTargetGroupServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverTargetGroupServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetGroup
}

func (s *RecoverTargetGroupServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	return
}

func (s *RecoverTargetGroupServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()

	return
}

// EffectUseServerPacket :: Map effect.
type EffectUseServerPacket struct {
	Effect     MapEffect
	EffectData EffectUseEffectData
}

type EffectUseEffectData interface {
	protocol.EoData
}

type EffectUseEffectDataQuake struct {
	QuakeStrength int
}

func (s *EffectUseEffectDataQuake) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuakeStrength : field : char
	if err = writer.AddChar(s.QuakeStrength); err != nil {
		return
	}
	return
}

func (s *EffectUseEffectDataQuake) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// QuakeStrength : field : char
	s.QuakeStrength = reader.GetChar()

	return
}

func (s EffectUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

func (s *EffectUseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Effect : field : MapEffect
	if err = writer.AddChar(int(s.Effect)); err != nil {
		return
	}
	switch s.Effect {
	case MapEffect_Quake:
		switch s.EffectData.(type) {
		case *EffectUseEffectDataQuake:
			if err = s.EffectData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.Effect)
			return
		}
	}
	return
}

func (s *EffectUseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Effect : field : MapEffect
	s.Effect = MapEffect(reader.GetChar())
	switch s.Effect {
	case MapEffect_Quake:
		s.EffectData = &EffectUseEffectDataQuake{}
		if err = s.EffectData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EffectAgreeServerPacket :: Map tile effect.
type EffectAgreeServerPacket struct {
	Coords   protocol.Coords
	EffectId int
}

func (s EffectAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

func (s *EffectAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// EffectId : field : short
	if err = writer.AddShort(s.EffectId); err != nil {
		return
	}
	return
}

func (s *EffectAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// EffectId : field : short
	s.EffectId = reader.GetShort()

	return
}

// EffectTargetOtherServerPacket :: Map drain damage.
type EffectTargetOtherServerPacket struct {
	Damage int
	Hp     int
	MaxHp  int
	Others []MapDrainDamageOther
}

func (s EffectTargetOtherServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectTargetOtherServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetOther
}

func (s *EffectTargetOtherServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Damage : field : short
	if err = writer.AddShort(s.Damage); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// Others : array : MapDrainDamageOther
	for ndx := 0; ndx < len(s.Others); ndx++ {
		if err = s.Others[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *EffectTargetOtherServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Damage : field : short
	s.Damage = reader.GetShort()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// Others : array : MapDrainDamageOther
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Others = append(s.Others, MapDrainDamageOther{})
		if err = s.Others[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EffectReportServerPacket :: Map spike timer.
type EffectReportServerPacket struct {
}

func (s EffectReportServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectReportServerPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

func (s *EffectReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *EffectReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// EffectSpecServerPacket :: Taking spike or tp drain damage.
type EffectSpecServerPacket struct {
	MapDamageType     MapDamageType
	MapDamageTypeData EffectSpecMapDamageTypeData
}

type EffectSpecMapDamageTypeData interface {
	protocol.EoData
}

type EffectSpecMapDamageTypeDataTpDrain struct {
	TpDamage int
	Tp       int
	MaxTp    int
}

func (s *EffectSpecMapDamageTypeDataTpDrain) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TpDamage : field : short
	if err = writer.AddShort(s.TpDamage); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	return
}

func (s *EffectSpecMapDamageTypeDataTpDrain) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// TpDamage : field : short
	s.TpDamage = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()

	return
}

type EffectSpecMapDamageTypeDataSpikes struct {
	HpDamage int
	Hp       int
	MaxHp    int
}

func (s *EffectSpecMapDamageTypeDataSpikes) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HpDamage : field : short
	if err = writer.AddShort(s.HpDamage); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	return
}

func (s *EffectSpecMapDamageTypeDataSpikes) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// HpDamage : field : short
	s.HpDamage = reader.GetShort()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()

	return
}

func (s EffectSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

func (s *EffectSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapDamageType : field : MapDamageType
	if err = writer.AddChar(int(s.MapDamageType)); err != nil {
		return
	}
	switch s.MapDamageType {
	case MapDamage_TpDrain:
		switch s.MapDamageTypeData.(type) {
		case *EffectSpecMapDamageTypeDataTpDrain:
			if err = s.MapDamageTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.MapDamageType)
			return
		}
	case MapDamage_Spikes:
		switch s.MapDamageTypeData.(type) {
		case *EffectSpecMapDamageTypeDataSpikes:
			if err = s.MapDamageTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.MapDamageType)
			return
		}
	}
	return
}

func (s *EffectSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MapDamageType : field : MapDamageType
	s.MapDamageType = MapDamageType(reader.GetChar())
	switch s.MapDamageType {
	case MapDamage_TpDrain:
		s.MapDamageTypeData = &EffectSpecMapDamageTypeDataTpDrain{}
		if err = s.MapDamageTypeData.Deserialize(reader); err != nil {
			return
		}
	case MapDamage_Spikes:
		s.MapDamageTypeData = &EffectSpecMapDamageTypeDataSpikes{}
		if err = s.MapDamageTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EffectAdminServerPacket :: Nearby character taking spike damage.
type EffectAdminServerPacket struct {
	PlayerId     int
	HpPercentage int
	Died         bool
	Damage       int
}

func (s EffectAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

func (s *EffectAdminServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Died : field : bool
	if s.Died {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	return
}

func (s *EffectAdminServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Died : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Died = true
	} else {
		s.Died = false
	}
	// Damage : field : three
	s.Damage = reader.GetThree()

	return
}

// MusicPlayerServerPacket :: Sound effect.
type MusicPlayerServerPacket struct {
	SoundId int
}

func (s MusicPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Music
}

func (s MusicPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

func (s *MusicPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SoundId : field : char
	if err = writer.AddChar(s.SoundId); err != nil {
		return
	}
	return
}

func (s *MusicPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SoundId : field : char
	s.SoundId = reader.GetChar()

	return
}

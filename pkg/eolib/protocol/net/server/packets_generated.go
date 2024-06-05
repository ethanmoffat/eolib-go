package server

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/pub"
)

// InitInitServerPacket ::  Reply to connection initialization and requests for unencrypted data. This packet is unencrypted.
type InitInitServerPacket struct {
	byteSize int

	ReplyCode     InitReply
	ReplyCodeData InitInitReplyCodeData
}

type InitInitReplyCodeData interface {
	protocol.EoData
}

type InitInitReplyCodeDataOutOfDate struct {
	byteSize int

	Version net.Version
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataOutOfDate) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Version : field : Version
	if err = s.Version.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataOk struct {
	byteSize int

	Seq1                     int
	Seq2                     int
	ServerEncryptionMultiple int
	ClientEncryptionMultiple int
	PlayerId                 int
	ChallengeResponse        int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataOk) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataBanned struct {
	byteSize int

	BanType     InitBanType
	BanTypeData InitInitBanTypeData
}

type InitInitBanTypeData interface {
	protocol.EoData
}

// InitInitBanTypeData0 ::  The official client treats any value below 2 as a temporary ban. The official server sends 1, but some game server implementations. erroneously send 0.
type InitInitBanTypeData0 struct {
	byteSize int

	MinutesRemaining int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitBanTypeData0) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MinutesRemaining : field : byte
	s.MinutesRemaining = int(reader.GetByte())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitBanTypeDataTemporary struct {
	byteSize int

	MinutesRemaining int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitBanTypeDataTemporary) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MinutesRemaining : field : byte
	s.MinutesRemaining = int(reader.GetByte())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataBanned) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataWarpMap struct {
	byteSize int

	MapFile MapFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataWarpMap) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataFileEmf struct {
	byteSize int

	MapFile MapFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataFileEmf) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataFileEif struct {
	byteSize int

	PubFile PubFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataFileEif) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataFileEnf struct {
	byteSize int

	PubFile PubFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataFileEnf) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataFileEsf struct {
	byteSize int

	PubFile PubFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataFileEsf) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataFileEcf struct {
	byteSize int

	PubFile PubFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataFileEcf) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataMapMutation struct {
	byteSize int

	MapFile MapFile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataMapMutation) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataPlayersList struct {
	byteSize int

	PlayersList PlayersList
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataPlayersList) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type InitInitReplyCodeDataPlayersListFriends struct {
	byteSize int

	PlayersList PlayersListFriends
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitReplyCodeDataPlayersListFriends) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s InitInitServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Init
}

func (s InitInitServerPacket) Action() net.PacketAction {
	return net.PacketAction_Init
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpPlayerServerPacket :: Equivalent to INIT_INIT with InitReply.WarpMap.
type WarpPlayerServerPacket struct {
	byteSize int

	MapFile MapFile
}

func (s WarpPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomePingServerPacket :: Equivalent to INIT_INIT with InitReply.FileMap.
type WelcomePingServerPacket struct {
	byteSize int

	MapFile MapFile
}

func (s WelcomePingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomePingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomePingServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomePongServerPacket :: Equivalent to INIT_INIT with InitReply.FileEif.
type WelcomePongServerPacket struct {
	byteSize int

	PubFile PubFile
}

func (s WelcomePongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomePongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomePongServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeNet242ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEnf.
type WelcomeNet242ServerPacket struct {
	byteSize int

	PubFile PubFile
}

func (s WelcomeNet242ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet242ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net242
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeNet242ServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeNet243ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEsf.
type WelcomeNet243ServerPacket struct {
	byteSize int

	PubFile PubFile
}

func (s WelcomeNet243ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet243ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net243
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeNet243ServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersListServerPacket :: Equivalent to INIT_INIT with InitReply.PlayersList.
type PlayersListServerPacket struct {
	byteSize int

	PlayersList PlayersList
}

func (s PlayersListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersListServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersList : field : PlayersList
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpCreateServerPacket :: Equivalent to INIT_INIT with InitReply.MapMutation.
type WarpCreateServerPacket struct {
	byteSize int

	MapFile MapFile
}

func (s WarpCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpCreateServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapFile : field : MapFile
	if err = s.MapFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersReplyServerPacket :: Equivalent to INIT_INIT with InitReply.PlayersListFriends.
type PlayersReplyServerPacket struct {
	byteSize int

	PlayersList PlayersListFriends
}

func (s PlayersReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersList : field : PlayersListFriends
	if err = s.PlayersList.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeNet244ServerPacket :: Equivalent to INIT_INIT with InitReply.FileEcf.
type WelcomeNet244ServerPacket struct {
	byteSize int

	PubFile PubFile
}

func (s WelcomeNet244ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeNet244ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net244
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeNet244ServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PubFile : field : PubFile
	if err = s.PubFile.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ConnectionPlayerServerPacket :: Ping request.
type ConnectionPlayerServerPacket struct {
	byteSize int

	Seq1 int
	Seq2 int
}

func (s ConnectionPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Connection
}

func (s ConnectionPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ConnectionPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Seq1 : field : short
	s.Seq1 = reader.GetShort()
	// Seq2 : field : char
	s.Seq2 = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AccountReplyServerPacket :: Reply to client Account-family packets.
type AccountReplyServerPacket struct {
	byteSize int

	ReplyCode     AccountReply //  Sometimes an AccountReply code, sometimes a session ID for account creation.
	ReplyCodeData AccountReplyReplyCodeData
}

type AccountReplyReplyCodeData interface {
	protocol.EoData
}

type AccountReplyReplyCodeDataExists struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataExists) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataExists) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataExists) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AccountReplyReplyCodeDataNotApproved struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataNotApproved) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataNotApproved) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataNotApproved) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AccountReplyReplyCodeDataCreated struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataCreated) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataCreated) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GO : field : string
	if err = writer.AddString("GO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataCreated) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// GO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AccountReplyReplyCodeDataChangeFailed struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataChangeFailed) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataChangeFailed) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataChangeFailed) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AccountReplyReplyCodeDataChanged struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataChanged) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataChanged) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataChanged) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AccountReplyReplyCodeDataRequestDenied struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataRequestDenied) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataRequestDenied) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataRequestDenied) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AccountReplyReplyCodeDataDefault ::  In this case (reply_code > 9), reply_code is a session ID for account creation.
type AccountReplyReplyCodeDataDefault struct {
	byteSize int

	SequenceStart int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyReplyCodeDataDefault) ByteSize() int {
	return s.byteSize
}

func (s *AccountReplyReplyCodeDataDefault) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SequenceStart : field : char
	if err = writer.AddChar(s.SequenceStart); err != nil {
		return
	}
	// OK : field : string
	if err = writer.AddString("OK"); err != nil {
		return
	}
	return
}

func (s *AccountReplyReplyCodeDataDefault) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SequenceStart : field : char
	s.SequenceStart = reader.GetChar()
	// OK : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s AccountReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Account
}

func (s AccountReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterReplyServerPacket :: Reply to client Character-family packets.
type CharacterReplyServerPacket struct {
	byteSize int

	ReplyCode     CharacterReply //  Sometimes a CharacterReply code, sometimes a session ID for character creation.
	ReplyCodeData CharacterReplyReplyCodeData
}

type CharacterReplyReplyCodeData interface {
	protocol.EoData
}

type CharacterReplyReplyCodeDataExists struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataExists) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataExists) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataExists) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type CharacterReplyReplyCodeDataFull struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataFull) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataFull) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataFull) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type CharacterReplyReplyCodeDataFull3 struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataFull3) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataFull3) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataFull3) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type CharacterReplyReplyCodeDataNotApproved struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataNotApproved) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataNotApproved) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataNotApproved) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type CharacterReplyReplyCodeDataOk struct {
	byteSize int

	CharactersCount int

	Characters []CharacterSelectionListEntry
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataOk) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataOk) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	writer.AddByte(255)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *CharacterReplyReplyCodeDataOk) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	// 0 : field : char
	reader.GetChar()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type CharacterReplyReplyCodeDataDeleted struct {
	byteSize int

	CharactersCount int
	Characters      []CharacterSelectionListEntry
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataDeleted) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataDeleted) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	writer.AddByte(255)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *CharacterReplyReplyCodeDataDeleted) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterReplyReplyCodeDataDefault ::  In this case (reply_code > 9), reply_code is a session ID for character creation.
type CharacterReplyReplyCodeDataDefault struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyReplyCodeDataDefault) ByteSize() int {
	return s.byteSize
}

func (s *CharacterReplyReplyCodeDataDefault) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// OK : field : string
	if err = writer.AddString("OK"); err != nil {
		return
	}
	return
}

func (s *CharacterReplyReplyCodeDataDefault) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// OK : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s CharacterReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterPlayerServerPacket :: Reply to client request to delete a character from the account (Character_Take).
type CharacterPlayerServerPacket struct {
	byteSize int

	SessionId   int
	CharacterId int
}

func (s CharacterPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LoginReplyServerPacket :: Login reply.
type LoginReplyServerPacket struct {
	byteSize int

	ReplyCode     LoginReply
	ReplyCodeData LoginReplyReplyCodeData
}

type LoginReplyReplyCodeData interface {
	protocol.EoData
}

type LoginReplyReplyCodeDataWrongUser struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataWrongUser) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataWrongUser) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataWrongUser) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type LoginReplyReplyCodeDataWrongUserPassword struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataWrongUserPassword) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataWrongUserPassword) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataWrongUserPassword) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type LoginReplyReplyCodeDataOk struct {
	byteSize int

	CharactersCount int

	Characters []CharacterSelectionListEntry
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataOk) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataOk) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(s.CharactersCount); err != nil {
		return
	}
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	writer.AddByte(255)
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *LoginReplyReplyCodeDataOk) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharactersCount : length : char
	s.CharactersCount = reader.GetChar()
	// 0 : field : char
	reader.GetChar()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Characters : array : CharacterSelectionListEntry
	for ndx := 0; ndx < s.CharactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterSelectionListEntry{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type LoginReplyReplyCodeDataBanned struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataBanned) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataBanned) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataBanned) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type LoginReplyReplyCodeDataLoggedIn struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataLoggedIn) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataLoggedIn) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataLoggedIn) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type LoginReplyReplyCodeDataBusy struct {
	byteSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyReplyCodeDataBusy) ByteSize() int {
	return s.byteSize
}

func (s *LoginReplyReplyCodeDataBusy) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NO : field : string
	if err = writer.AddString("NO"); err != nil {
		return
	}
	return
}

func (s *LoginReplyReplyCodeDataBusy) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NO : field : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s LoginReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Login
}

func (s LoginReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeReplyServerPacket :: Reply to selecting a character / entering game.
type WelcomeReplyServerPacket struct {
	byteSize int

	WelcomeCode     WelcomeCode
	WelcomeCodeData WelcomeReplyWelcomeCodeData
}

type WelcomeReplyWelcomeCodeData interface {
	protocol.EoData
}

type WelcomeReplyWelcomeCodeDataSelectCharacter struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeReplyWelcomeCodeDataSelectCharacter) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// Title : field : string
	if err = writer.AddString(s.Title); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildRankName : field : string
	if err = writer.AddString(s.GuildRankName); err != nil {
		return
	}
	writer.AddByte(255)
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
	writer.AddByte(255)
	return
}

func (s *WelcomeReplyWelcomeCodeDataSelectCharacter) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Title : field : string
	if s.Title, err = reader.GetString(); err != nil {
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
	// GuildRankName : field : string
	if s.GuildRankName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
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
	if err = reader.NextChunk(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type WelcomeReplyWelcomeCodeDataEnterGame struct {
	byteSize int

	News   []string
	Weight net.Weight
	Items  []net.Item
	Spells []net.Spell
	Nearby NearbyInfo
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeReplyWelcomeCodeDataEnterGame) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeReplyWelcomeCodeDataEnterGame) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.AddByte(255)
	// News : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.News[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
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

	writer.AddByte(255)
	// Spells : array : Spell
	for ndx := 0; ndx < len(s.Spells); ndx++ {
		if err = s.Spells[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WelcomeReplyWelcomeCodeDataEnterGame) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// News : array : string
	for ndx := 0; ndx < 9; ndx++ {
		s.News = append(s.News, "")
		if s.News[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	// Weight : field : Weight
	if err = s.Weight.Deserialize(reader); err != nil {
		return
	}
	// Items : array : Item
	ItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < ItemsRemaining/6; ndx++ {
		s.Items = append(s.Items, net.Item{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Spells : array : Spell
	SpellsRemaining := reader.Remaining()
	for ndx := 0; ndx < SpellsRemaining/4; ndx++ {
		s.Spells = append(s.Spells, net.Spell{})
		if err = s.Spells[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s WelcomeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractReplyServerPacket :: Incoming admin message.
type AdminInteractReplyServerPacket struct {
	byteSize int

	MessageType     AdminMessageType
	MessageTypeData AdminInteractReplyMessageTypeData
}

type AdminInteractReplyMessageTypeData interface {
	protocol.EoData
}

type AdminInteractReplyMessageTypeDataMessage struct {
	byteSize int

	PlayerName string
	Message    string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractReplyMessageTypeDataMessage) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractReplyMessageTypeDataMessage) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.AddByte(255)
	return
}

func (s *AdminInteractReplyMessageTypeDataMessage) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	if err = reader.NextChunk(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type AdminInteractReplyMessageTypeDataReport struct {
	byteSize int

	PlayerName   string
	Message      string
	ReporteeName string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractReplyMessageTypeDataReport) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractReplyMessageTypeDataReport) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.AddByte(255)
	// ReporteeName : field : string
	if err = writer.AddString(s.ReporteeName); err != nil {
		return
	}
	writer.AddByte(255)
	return
}

func (s *AdminInteractReplyMessageTypeDataReport) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	if err = reader.NextChunk(); err != nil {
		return
	}
	// ReporteeName : field : string
	if s.ReporteeName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s AdminInteractReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractReplyServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// MessageType : field : AdminMessageType
	if err = writer.AddChar(int(s.MessageType)); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractRemoveServerPacket :: Nearby player disappearing (admin hide).
type AdminInteractRemoveServerPacket struct {
	byteSize int

	PlayerId int
}

func (s AdminInteractRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractAgreeServerPacket :: Nearby player appearing (admin un-hide).
type AdminInteractAgreeServerPacket struct {
	byteSize int

	PlayerId int
}

func (s AdminInteractAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractListServerPacket :: Admin character inventory popup.
type AdminInteractListServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractListServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractListServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Usage : field : int
	if err = writer.AddInt(s.Usage); err != nil {
		return
	}
	writer.AddByte(255)
	// GoldBank : field : int
	if err = writer.AddInt(s.GoldBank); err != nil {
		return
	}
	writer.AddByte(255)
	// Inventory : array : Item
	for ndx := 0; ndx < len(s.Inventory); ndx++ {
		if err = s.Inventory[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	InventoryRemaining := reader.Remaining()
	for ndx := 0; ndx < InventoryRemaining/6; ndx++ {
		s.Inventory = append(s.Inventory, net.Item{})
		if err = s.Inventory[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Bank : array : ThreeItem
	BankRemaining := reader.Remaining()
	for ndx := 0; ndx < BankRemaining/5; ndx++ {
		s.Bank = append(s.Bank, net.ThreeItem{})
		if err = s.Bank[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractTellServerPacket :: Admin character info lookup.
type AdminInteractTellServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractTellServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Usage : field : int
	if err = writer.AddInt(s.Usage); err != nil {
		return
	}
	writer.AddByte(255)
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkRequestServerPacket :: Guild chat message.
type TalkRequestServerPacket struct {
	byteSize int

	PlayerName string
	Message    string
}

func (s TalkRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkRequestServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkOpenServerPacket :: Party chat message.
type TalkOpenServerPacket struct {
	byteSize int

	PlayerId int
	Message  string
}

func (s TalkOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkMsgServerPacket :: Global chat message.
type TalkMsgServerPacket struct {
	byteSize int

	PlayerName string
	Message    string
}

func (s TalkMsgServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkMsgServerPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkMsgServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkMsgServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkTellServerPacket :: Private chat message.
type TalkTellServerPacket struct {
	byteSize int

	PlayerName string
	Message    string
}

func (s TalkTellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkTellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkTellServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkPlayerServerPacket :: Public chat message.
type TalkPlayerServerPacket struct {
	byteSize int

	PlayerId int
	Message  string
}

func (s TalkPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkReplyServerPacket :: Reply to trying to send a private message.
type TalkReplyServerPacket struct {
	byteSize int

	ReplyCode TalkReply
	Name      string
}

func (s TalkReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ReplyCode : field : TalkReply
	s.ReplyCode = TalkReply(reader.GetShort())
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkAdminServerPacket :: Admin chat message.
type TalkAdminServerPacket struct {
	byteSize int

	PlayerName string
	Message    string
}

func (s TalkAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkAdminServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkAdminServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkAnnounceServerPacket :: Admin announcement.
type TalkAnnounceServerPacket struct {
	byteSize int

	PlayerName string
	Message    string
}

func (s TalkAnnounceServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAnnounceServerPacket) Action() net.PacketAction {
	return net.PacketAction_Announce
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkAnnounceServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkAnnounceServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkServerServerPacket :: Server message.
type TalkServerServerPacket struct {
	byteSize int

	Message string
}

func (s TalkServerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkServerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Server
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkServerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkListServerPacket ::  Global chat backfill. Sent by the official game server when a player opens the global chat tab.
type TalkListServerPacket struct {
	byteSize int

	Messages []GlobalBackfillMessage
}

func (s TalkListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkListServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *TalkListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MessageOpenServerPacket :: Status bar message.
type MessageOpenServerPacket struct {
	byteSize int

	Message string
}

func (s MessageOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MessageOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MessageCloseServerPacket :: Server is rebooting.
type MessageCloseServerPacket struct {
	byteSize int
}

func (s MessageCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MessageCloseServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *MessageCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// r : dummy : string
	if err = writer.AddString("r"); err != nil {
		return
	}
	return
}

func (s *MessageCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// r : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MessageAcceptServerPacket :: Large message box.
type MessageAcceptServerPacket struct {
	byteSize int

	Messages []string
}

func (s MessageAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessageAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MessageAcceptServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *MessageAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkSpecServerPacket :: Temporary mute applied.
type TalkSpecServerPacket struct {
	byteSize int

	AdminName string
}

func (s TalkSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// AdminName : field : string
	if s.AdminName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AttackPlayerServerPacket :: Nearby player attacking.
type AttackPlayerServerPacket struct {
	byteSize int

	PlayerId  int
	Direction protocol.Direction
}

func (s AttackPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Attack
}

func (s AttackPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AttackPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AttackErrorServerPacket :: Show flood protection message (vestigial).
type AttackErrorServerPacket struct {
	byteSize int
}

func (s AttackErrorServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Attack
}

func (s AttackErrorServerPacket) Action() net.PacketAction {
	return net.PacketAction_Error
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AttackErrorServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *AttackErrorServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *AttackErrorServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AvatarReplyServerPacket :: Nearby player hit by another player.
type AvatarReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AvatarReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChairPlayerServerPacket :: Nearby player sitting on a chair.
type ChairPlayerServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChairReplyServerPacket :: Your character sitting on a chair.
type ChairReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChairCloseServerPacket :: Your character standing up from a chair.
type ChairCloseServerPacket struct {
	byteSize int

	PlayerId int
	Coords   protocol.Coords
}

func (s ChairCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairCloseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChairRemoveServerPacket :: Nearby player standing up from a chair.
type ChairRemoveServerPacket struct {
	byteSize int

	PlayerId int
	Coords   protocol.Coords
}

func (s ChairRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SitPlayerServerPacket :: Nearby player sitting down.
type SitPlayerServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitPlayerServerPacket) ByteSize() int {
	return s.byteSize
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
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *SitPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// 0 : field : char
	reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SitCloseServerPacket :: Your character standing up.
type SitCloseServerPacket struct {
	byteSize int

	PlayerId int
	Coords   protocol.Coords
}

func (s SitCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitCloseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SitRemoveServerPacket :: Nearby player standing up.
type SitRemoveServerPacket struct {
	byteSize int

	PlayerId int
	Coords   protocol.Coords
}

func (s SitRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SitReplyServerPacket :: Your character sitting down.
type SitReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitReplyServerPacket) ByteSize() int {
	return s.byteSize
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
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *SitReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// 0 : field : char
	reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EmotePlayerServerPacket :: Nearby player doing an emote.
type EmotePlayerServerPacket struct {
	byteSize int

	PlayerId int
	Emote    protocol.Emote
}

func (s EmotePlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Emote
}

func (s EmotePlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EmotePlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Emote : field : Emote
	s.Emote = protocol.Emote(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectPlayerServerPacket :: Effects playing on nearby players.
type EffectPlayerServerPacket struct {
	byteSize int

	Effects []PlayerEffect
}

func (s EffectPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectPlayerServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *EffectPlayerServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Effects : array : PlayerEffect
	for ndx := 0; ndx < len(s.Effects); ndx++ {
		if err = s.Effects[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *EffectPlayerServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Effects : array : PlayerEffect
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Effects = append(s.Effects, PlayerEffect{})
		if err = s.Effects[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// FacePlayerServerPacket :: Nearby player facing a direction.
type FacePlayerServerPacket struct {
	byteSize int

	PlayerId  int
	Direction protocol.Direction
}

func (s FacePlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Face
}

func (s FacePlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *FacePlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AvatarRemoveServerPacket :: Nearby player has disappeared from view.
type AvatarRemoveServerPacket struct {
	byteSize int

	PlayerId   int
	WarpEffect *WarpEffect
}

func (s AvatarRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AvatarRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// WarpEffect : field : WarpEffect
	if reader.Remaining() > 0 {
		s.WarpEffect = new(WarpEffect)
		*s.WarpEffect = WarpEffect(reader.GetChar())
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersAgreeServerPacket :: Player has appeared in nearby view.
type PlayersAgreeServerPacket struct {
	byteSize int

	Nearby NearbyInfo
}

func (s PlayersAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersRemoveServerPacket :: Nearby player has logged out.
type PlayersRemoveServerPacket struct {
	byteSize int

	PlayerId int
}

func (s PlayersRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RangeReplyServerPacket :: Reply to request for information about nearby players and NPCs.
type RangeReplyServerPacket struct {
	byteSize int

	Nearby NearbyInfo
}

func (s RangeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Range
}

func (s RangeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RangeReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcAgreeServerPacket :: Reply to request for information about nearby NPCs.
type NpcAgreeServerPacket struct {
	byteSize int

	NpcsCount int
	Npcs      []NpcMapInfo
}

func (s NpcAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcAgreeServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *NpcAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcsCount : length : char
	if err = writer.AddChar(s.NpcsCount); err != nil {
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

	readerStartPosition := reader.Position()
	// NpcsCount : length : char
	s.NpcsCount = reader.GetChar()
	// Npcs : array : NpcMapInfo
	for ndx := 0; ndx < s.NpcsCount; ndx++ {
		s.Npcs = append(s.Npcs, NpcMapInfo{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkPlayerServerPacket :: Nearby player has walked.
type WalkPlayerServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkReplyServerPacket :: Players, NPCs, and Items appearing in nearby view.
type WalkReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	writer.AddByte(255)
	// NpcIndexes : array : char
	for ndx := 0; ndx < len(s.NpcIndexes); ndx++ {
		if err = writer.AddChar(s.NpcIndexes[ndx]); err != nil {
			return
		}
	}

	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayerIds : array : short
	PlayerIdsRemaining := reader.Remaining()
	for ndx := 0; ndx < PlayerIdsRemaining/2; ndx++ {
		s.PlayerIds = append(s.PlayerIds, 0)
		s.PlayerIds[ndx] = reader.GetShort()
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// NpcIndexes : array : char
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.NpcIndexes = append(s.NpcIndexes, 0)
		s.NpcIndexes[ndx] = reader.GetChar()
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Items : array : ItemMapInfo
	ItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < ItemsRemaining/9; ndx++ {
		s.Items = append(s.Items, ItemMapInfo{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkCloseServerPacket :: Your character has been frozen.
type WalkCloseServerPacket struct {
	byteSize int
}

func (s WalkCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkCloseServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *WalkCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// S : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *WalkCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// S : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkOpenServerPacket :: Your character has been unfrozen.
type WalkOpenServerPacket struct {
	byteSize int
}

func (s WalkOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkOpenServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *WalkOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// S : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *WalkOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// S : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BankOpenServerPacket :: Open banker NPC interface.
type BankOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BankOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldBank : field : int
	s.GoldBank = reader.GetInt()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	// LockerUpgrades : field : char
	s.LockerUpgrades = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BankReplyServerPacket :: Update gold counts after deposit/withdraw.
type BankReplyServerPacket struct {
	byteSize int

	GoldInventory int
	GoldBank      int
}

func (s BankReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BankReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldInventory : field : int
	s.GoldInventory = reader.GetInt()
	// GoldBank : field : int
	s.GoldBank = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BarberAgreeServerPacket :: Purchasing a new hair style.
type BarberAgreeServerPacket struct {
	byteSize int

	GoldAmount int
	Change     AvatarChange
}

func (s BarberAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BarberAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BarberOpenServerPacket :: Response from talking to a barber NPC.
type BarberOpenServerPacket struct {
	byteSize int

	SessionId int
}

func (s BarberOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BarberOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerReplyServerPacket :: Response to adding an item to a bank locker.
type LockerReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerGetServerPacket :: Response to taking an item from a bank locker.
type LockerGetServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerGetServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerOpenServerPacket :: Opening a bank locker.
type LockerOpenServerPacket struct {
	byteSize int

	LockerCoords protocol.Coords
	LockerItems  []net.ThreeItem
}

func (s LockerOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerBuyServerPacket :: Response to buying a locker space upgrade from a banker NPC.
type LockerBuyServerPacket struct {
	byteSize int

	GoldAmount     int
	LockerUpgrades int
}

func (s LockerBuyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerBuyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerBuyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	// LockerUpgrades : field : char
	s.LockerUpgrades = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerSpecServerPacket :: Reply to trying to add an item to a full locker.
type LockerSpecServerPacket struct {
	byteSize int

	LockerMaxItems int
}

func (s LockerSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// LockerMaxItems : field : char
	s.LockerMaxItems = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenReplyServerPacket :: Response to subscribing to a town.
type CitizenReplyServerPacket struct {
	byteSize int

	QuestionsWrong int
}

func (s CitizenReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// QuestionsWrong : field : char
	s.QuestionsWrong = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenRemoveServerPacket :: Response to giving up citizenship of a town.
type CitizenRemoveServerPacket struct {
	byteSize int

	ReplyCode InnUnsubscribeReply
}

func (s CitizenRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ReplyCode : field : InnUnsubscribeReply
	s.ReplyCode = InnUnsubscribeReply(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenOpenServerPacket :: Response from talking to a citizenship NPC.
type CitizenOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenOpenServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// Questions : array : string
	for ndx := 0; ndx < 3; ndx++ {
		if ndx > 0 {
			writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenRequestServerPacket :: Reply to requesting sleeping at an inn.
type CitizenRequestServerPacket struct {
	byteSize int

	Cost int
}

func (s CitizenRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Cost : field : int
	s.Cost = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenAcceptServerPacket :: Sleeping at an inn.
type CitizenAcceptServerPacket struct {
	byteSize int

	GoldAmount int
}

func (s CitizenAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopCreateServerPacket :: Response to crafting an item from a shop.
type ShopCreateServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopCreateServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopBuyServerPacket :: Response to purchasing an item from a shop.
type ShopBuyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopBuyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopSellServerPacket :: Response to selling an item to a shop.
type ShopSellServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopSellServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopOpenServerPacket :: Response from talking to a shop NPC.
type ShopOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopOpenServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// TradeItems : array : ShopTradeItem
	for ndx := 0; ndx < len(s.TradeItems); ndx++ {
		if err = s.TradeItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
	// CraftItems : array : ShopCraftItem
	for ndx := 0; ndx < len(s.CraftItems); ndx++ {
		if err = s.CraftItems[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *ShopOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	TradeItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < TradeItemsRemaining/9; ndx++ {
		s.TradeItems = append(s.TradeItems, ShopTradeItem{})
		if err = s.TradeItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// CraftItems : array : ShopCraftItem
	CraftItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < CraftItemsRemaining/14; ndx++ {
		s.CraftItems = append(s.CraftItems, ShopCraftItem{})
		if err = s.CraftItems[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillOpenServerPacket :: Response from talking to a skill master NPC.
type StatSkillOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillOpenServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	SkillsRemaining := reader.Remaining()
	for ndx := 0; ndx < SkillsRemaining/28; ndx++ {
		s.Skills = append(s.Skills, SkillLearn{})
		if err = s.Skills[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillReplyServerPacket :: Response from unsuccessful action at a skill master.
type StatSkillReplyServerPacket struct {
	byteSize int

	ReplyCode     SkillMasterReply
	ReplyCodeData StatSkillReplyReplyCodeData
}

type StatSkillReplyReplyCodeData interface {
	protocol.EoData
}

type StatSkillReplyReplyCodeDataWrongClass struct {
	byteSize int

	ClassId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillReplyReplyCodeDataWrongClass) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s StatSkillReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ReplyCode : field : SkillMasterReply
	s.ReplyCode = SkillMasterReply(reader.GetShort())
	switch s.ReplyCode {
	case SkillMasterReply_WrongClass:
		s.ReplyCodeData = &StatSkillReplyReplyCodeDataWrongClass{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillTakeServerPacket :: Response from learning a skill from a skill master.
type StatSkillTakeServerPacket struct {
	byteSize int

	SpellId    int
	GoldAmount int
}

func (s StatSkillTakeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillTakeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillTakeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillRemoveServerPacket :: Response to forgetting a skill at a skill master.
type StatSkillRemoveServerPacket struct {
	byteSize int

	SpellId int
}

func (s StatSkillRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillPlayerServerPacket :: Response to spending stat points.
type StatSkillPlayerServerPacket struct {
	byteSize int

	StatPoints int
	Stats      CharacterStatsUpdate
}

func (s StatSkillPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillAcceptServerPacket :: Response to spending skill points.
type StatSkillAcceptServerPacket struct {
	byteSize int

	SkillPoints int
	Spell       net.Spell
}

func (s StatSkillAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// Spell : field : Spell
	if err = s.Spell.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillJunkServerPacket :: Response to resetting stats and skills at a skill master.
type StatSkillJunkServerPacket struct {
	byteSize int

	Stats CharacterStatsReset
}

func (s StatSkillJunkServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillJunkServerPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillJunkServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Stats : field : CharacterStatsReset
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemReplyServerPacket :: Reply to using an item.
type ItemReplyServerPacket struct {
	byteSize int

	ItemType     pub.ItemType
	UsedItem     net.Item
	Weight       net.Weight
	ItemTypeData ItemReplyItemTypeData
}

type ItemReplyItemTypeData interface {
	protocol.EoData
}

type ItemReplyItemTypeDataHeal struct {
	byteSize int

	HpGain int
	Hp     int
	Tp     int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyItemTypeDataHeal) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// HpGain : field : int
	s.HpGain = reader.GetInt()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ItemReplyItemTypeDataHairDye struct {
	byteSize int

	HairColor int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyItemTypeDataHairDye) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ItemReplyItemTypeDataEffectPotion struct {
	byteSize int

	EffectId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyItemTypeDataEffectPotion) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// EffectId : field : short
	s.EffectId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ItemReplyItemTypeDataCureCurse struct {
	byteSize int

	Stats CharacterStatsEquipmentChange
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyItemTypeDataCureCurse) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Stats : field : CharacterStatsEquipmentChange
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ItemReplyItemTypeDataExpReward struct {
	byteSize int

	Experience  int
	LevelUp     int //  A value greater than 0 is "new level" and indicates the player leveled up.
	StatPoints  int
	SkillPoints int
	MaxHp       int
	MaxTp       int
	MaxSp       int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyItemTypeDataExpReward) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s ItemReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemDropServerPacket :: Reply to dropping items on the ground.
type ItemDropServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemDropServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemAddServerPacket :: Item appeared on the ground.
type ItemAddServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemAddServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemRemoveServerPacket :: Item disappeared from the ground.
type ItemRemoveServerPacket struct {
	byteSize int

	ItemIndex int
}

func (s ItemRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemJunkServerPacket :: Reply to junking items.
type ItemJunkServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemJunkServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemGetServerPacket :: Reply to taking items from the ground.
type ItemGetServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemGetServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemObtainServerPacket :: Receive item (from quest).
type ItemObtainServerPacket struct {
	byteSize int

	Item          net.ThreeItem
	CurrentWeight int
}

func (s ItemObtainServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemObtainServerPacket) Action() net.PacketAction {
	return net.PacketAction_Obtain
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemObtainServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Item : field : ThreeItem
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// CurrentWeight : field : char
	s.CurrentWeight = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemKickServerPacket :: Lose item (from quest).
type ItemKickServerPacket struct {
	byteSize int

	Item          net.Item
	CurrentWeight int
}

func (s ItemKickServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemKickServerPacket) Action() net.PacketAction {
	return net.PacketAction_Kick
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemKickServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Item : field : Item
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// CurrentWeight : field : char
	s.CurrentWeight = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemAgreeServerPacket :: Reply to using an item that you don't have.
type ItemAgreeServerPacket struct {
	byteSize int

	ItemId int
}

func (s ItemAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemSpecServerPacket :: Reply to trying to take a protected item from the ground.
type ItemSpecServerPacket struct {
	byteSize int
}

func (s ItemSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemSpecServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *ItemSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 2 : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}
	return
}

func (s *ItemSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 2 : dummy : short
	reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardPlayerServerPacket :: Reply to reading a post on a town board.
type BoardPlayerServerPacket struct {
	byteSize int

	PostId   int
	PostBody string
}

func (s BoardPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PostId : field : short
	s.PostId = reader.GetShort()
	// PostBody : field : string
	if s.PostBody, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardOpenServerPacket :: Reply to opening a town board.
type BoardOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardOpenServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *BoardOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxAgreeServerPacket :: Reply to successfully requesting a song.
type JukeboxAgreeServerPacket struct {
	byteSize int

	GoldAmount int
}

func (s JukeboxAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxReplyServerPacket :: Reply to unsuccessfully requesting a song.
type JukeboxReplyServerPacket struct {
	byteSize int
}

func (s JukeboxReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxReplyServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *JukeboxReplyServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 1 : dummy : short
	if err = writer.AddShort(1); err != nil {
		return
	}
	return
}

func (s *JukeboxReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 1 : dummy : short
	reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxOpenServerPacket :: Reply to opening the jukebox listing.
type JukeboxOpenServerPacket struct {
	byteSize int

	MapId         int
	JukeboxPlayer string
}

func (s JukeboxOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// JukeboxPlayer : field : string
	if s.JukeboxPlayer, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxMsgServerPacket :: Someone playing a note with the bard skill nearby.
type JukeboxMsgServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxMsgServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// InstrumentId : field : char
	s.InstrumentId = reader.GetChar()
	// NoteId : field : char
	s.NoteId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxPlayerServerPacket :: Play background music.
type JukeboxPlayerServerPacket struct {
	byteSize int

	MfxId int
}

func (s JukeboxPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MfxId : field : char
	s.MfxId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxUseServerPacket :: Play jukebox music.
type JukeboxUseServerPacket struct {
	byteSize int

	TrackId int // This value is 1-indexed.
}

func (s JukeboxUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxUseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// TrackId : field : short
	s.TrackId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpRequestServerPacket :: Warp request from server.
type WarpRequestServerPacket struct {
	byteSize int

	WarpType     WarpType
	MapId        int
	WarpTypeData WarpRequestWarpTypeData
	SessionId    int
}

type WarpRequestWarpTypeData interface {
	protocol.EoData
}

type WarpRequestWarpTypeDataMapSwitch struct {
	byteSize int

	MapRid      []int
	MapFileSize int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpRequestWarpTypeDataMapSwitch) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapRid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.MapRid = append(s.MapRid, 0)
		s.MapRid[ndx] = reader.GetShort()
	}

	// MapFileSize : field : three
	s.MapFileSize = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s WarpRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpAgreeServerPacket :: Reply after accepting a warp.
type WarpAgreeServerPacket struct {
	byteSize int

	WarpType     WarpType
	WarpTypeData WarpAgreeWarpTypeData
	Nearby       NearbyInfo
}

type WarpAgreeWarpTypeData interface {
	protocol.EoData
}

type WarpAgreeWarpTypeDataMapSwitch struct {
	byteSize int

	MapId      int
	WarpEffect WarpEffect
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpAgreeWarpTypeDataMapSwitch) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// WarpEffect : field : WarpEffect
	s.WarpEffect = WarpEffect(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s WarpAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollReplyServerPacket :: Reply to requesting a paperdoll.
type PaperdollReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollPingServerPacket :: Failed to equip an item due to being the incorrect class.
type PaperdollPingServerPacket struct {
	byteSize int

	ClassId int // The player's current class ID (not the item's required class ID).
}

func (s PaperdollPingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollPingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollPingServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollRemoveServerPacket :: Reply to unequipping an item.
type PaperdollRemoveServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollAgreeServerPacket :: Reply to equipping an item.
type PaperdollAgreeServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AvatarAgreeServerPacket :: Nearby player changed appearance.
type AvatarAgreeServerPacket struct {
	byteSize int

	Change AvatarChange
}

func (s AvatarAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Avatar
}

func (s AvatarAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AvatarAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Change : field : AvatarChange
	if err = s.Change.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BookReplyServerPacket :: Reply to requesting a book.
type BookReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BookReplyServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// QuestNames : array : string
	for ndx := 0; ndx < len(s.QuestNames); ndx++ {
		if err = writer.AddString(s.QuestNames[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *BookReplyServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MessagePongServerPacket :: #ping command reply.
type MessagePongServerPacket struct {
	byteSize int
}

func (s MessagePongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessagePongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MessagePongServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *MessagePongServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 2 : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}
	return
}

func (s *MessagePongServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 2 : dummy : short
	reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersPingServerPacket :: #find command reply - offline.
type PlayersPingServerPacket struct {
	byteSize int

	Name string
}

func (s PlayersPingServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersPingServerPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersPingServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersPongServerPacket :: #find command reply - same map.
type PlayersPongServerPacket struct {
	byteSize int

	Name string
}

func (s PlayersPongServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersPongServerPacket) Action() net.PacketAction {
	return net.PacketAction_Pong
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersPongServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersNet242ServerPacket :: #find command reply - different map.
type PlayersNet242ServerPacket struct {
	byteSize int

	Name string
}

func (s PlayersNet242ServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersNet242ServerPacket) Action() net.PacketAction {
	return net.PacketAction_Net242
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersNet242ServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// DoorOpenServerPacket :: Nearby door opening.
type DoorOpenServerPacket struct {
	byteSize int

	Coords protocol.Coords
}

func (s DoorOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Door
}

func (s DoorOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *DoorOpenServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *DoorOpenServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *DoorOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// 0 : field : char
	reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// DoorCloseServerPacket :: Reply to trying to open a locked door.
type DoorCloseServerPacket struct {
	byteSize int

	Key int
}

func (s DoorCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Door
}

func (s DoorCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *DoorCloseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Key : field : char
	s.Key = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestOpenServerPacket :: Reply to opening a chest.
type ChestOpenServerPacket struct {
	byteSize int

	Coords protocol.Coords
	Items  []net.ThreeItem
}

func (s ChestOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestReplyServerPacket :: Reply to placing an item in to a chest.
type ChestReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestGetServerPacket :: Reply to removing an item from a chest.
type ChestGetServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestGetServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestAgreeServerPacket :: Chest contents updating.
type ChestAgreeServerPacket struct {
	byteSize int

	Items []net.ThreeItem
}

func (s ChestAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Items : array : ThreeItem
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, net.ThreeItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestSpecServerPacket :: Reply to trying to add an item to a full chest.
type ChestSpecServerPacket struct {
	byteSize int
}

func (s ChestSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestSpecServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *ChestSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 0 : dummy : byte
	if err = writer.AddByte(0); err != nil {
		return
	}
	return
}

func (s *ChestSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 0 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestCloseServerPacket ::  Reply to trying to interact with a locked or "broken" chest. The official client assumes a broken chest if the packet is under 2 bytes in length.
type ChestCloseServerPacket struct {
	byteSize int

	Key *int // Sent if the player is trying to interact with a locked chest.
}

func (s ChestCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestCloseServerPacket) ByteSize() int {
	return s.byteSize
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
	// N : dummy : string
	if err = writer.AddString("N"); err != nil {
		return
	}
	return
}

func (s *ChestCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Key : field : short
	if reader.Remaining() > 0 {
		s.Key = new(int)
		*s.Key = reader.GetShort()
	}
	// N : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RefreshReplyServerPacket :: Reply to request for new info about nearby objects.
type RefreshReplyServerPacket struct {
	byteSize int

	Nearby NearbyInfo
}

func (s RefreshReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Refresh
}

func (s RefreshReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RefreshReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Nearby : field : NearbyInfo
	if err = s.Nearby.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyRequestServerPacket :: Received party invite / join request.
type PartyRequestServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyReplyServerPacket :: Failed party invite / join request.
type PartyReplyServerPacket struct {
	byteSize int

	ReplyCode     PartyReplyCode
	ReplyCodeData PartyReplyReplyCodeData
}

type PartyReplyReplyCodeData interface {
	protocol.EoData
}

type PartyReplyReplyCodeDataAlreadyInAnotherParty struct {
	byteSize int

	PlayerName string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyReplyReplyCodeDataAlreadyInAnotherParty) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type PartyReplyReplyCodeDataAlreadyInYourParty struct {
	byteSize int

	PlayerName string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyReplyReplyCodeDataAlreadyInYourParty) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s PartyReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyCreateServerPacket :: Member list received when party is first joined.
type PartyCreateServerPacket struct {
	byteSize int

	Members []PartyMember
}

func (s PartyCreateServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyCreateServerPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyCreateServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PartyCreateServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyAddServerPacket :: New player joined the party.
type PartyAddServerPacket struct {
	byteSize int

	Member PartyMember
}

func (s PartyAddServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyAddServerPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyAddServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Member : field : PartyMember
	if err = s.Member.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyRemoveServerPacket :: Player left the party.
type PartyRemoveServerPacket struct {
	byteSize int

	PlayerId int
}

func (s PartyRemoveServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyRemoveServerPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyRemoveServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyCloseServerPacket :: Left / disbanded a party.
type PartyCloseServerPacket struct {
	byteSize int
}

func (s PartyCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyCloseServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *PartyCloseServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *PartyCloseServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyListServerPacket :: Party member list update.
type PartyListServerPacket struct {
	byteSize int

	Members []PartyMember
}

func (s PartyListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyListServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PartyListServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyAgreeServerPacket :: Party member list update.
type PartyAgreeServerPacket struct {
	byteSize int

	PlayerId     int
	HpPercentage int
}

func (s PartyAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyTargetGroupServerPacket :: Updated experience and level-ups from party experience.
type PartyTargetGroupServerPacket struct {
	byteSize int

	Gains []PartyExpShare
}

func (s PartyTargetGroupServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyTargetGroupServerPacket) Action() net.PacketAction {
	return net.PacketAction_TargetGroup
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyTargetGroupServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Gains : array : PartyExpShare
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Gains = append(s.Gains, PartyExpShare{})
		if err = s.Gains[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildReplyServerPacket :: Generic guild reply messages.
type GuildReplyServerPacket struct {
	byteSize int

	ReplyCode     GuildReply
	ReplyCodeData GuildReplyReplyCodeData
}

type GuildReplyReplyCodeData interface {
	protocol.EoData
}

type GuildReplyReplyCodeDataCreateAdd struct {
	byteSize int

	Name string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReplyReplyCodeDataCreateAdd) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type GuildReplyReplyCodeDataCreateAddConfirm struct {
	byteSize int

	Name string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReplyReplyCodeDataCreateAddConfirm) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type GuildReplyReplyCodeDataJoinRequest struct {
	byteSize int

	PlayerId int
	Name     string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReplyReplyCodeDataJoinRequest) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s GuildReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildRequestServerPacket :: Guild create request.
type GuildRequestServerPacket struct {
	byteSize int

	PlayerId      int
	GuildIdentity string
}

func (s GuildRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildCreateServerPacket :: Guild created.
type GuildCreateServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildCreateServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildCreateServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// LeaderPlayerId : field : short
	if err = writer.AddShort(s.LeaderPlayerId); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(255)
	// RankName : field : string
	if err = writer.AddString(s.RankName); err != nil {
		return
	}
	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildTakeServerPacket :: Get guild description reply.
type GuildTakeServerPacket struct {
	byteSize int

	Description string
}

func (s GuildTakeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTakeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildTakeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildRankServerPacket :: Get guild rank list reply.
type GuildRankServerPacket struct {
	byteSize int

	Ranks []string
}

func (s GuildRankServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRankServerPacket) Action() net.PacketAction {
	return net.PacketAction_Rank
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildRankServerPacket) ByteSize() int {
	return s.byteSize
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
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildRankServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildSellServerPacket :: Get guild bank reply.
type GuildSellServerPacket struct {
	byteSize int

	GoldAmount int
}

func (s GuildSellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildSellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Sell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildSellServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildBuyServerPacket :: Deposit guild bank reply.
type GuildBuyServerPacket struct {
	byteSize int

	GoldAmount int
}

func (s GuildBuyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildBuyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildBuyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildOpenServerPacket :: Talk to guild master NPC reply.
type GuildOpenServerPacket struct {
	byteSize int

	SessionId int
}

func (s GuildOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildTellServerPacket :: Get guild member list reply.
type GuildTellServerPacket struct {
	byteSize int

	MembersCount int
	Members      []GuildMember
}

func (s GuildTellServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTellServerPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildTellServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildTellServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// MembersCount : length : short
	if err = writer.AddShort(s.MembersCount); err != nil {
		return
	}
	writer.AddByte(255)
	// Members : array : GuildMember
	for ndx := 0; ndx < s.MembersCount; ndx++ {
		if err = s.Members[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildTellServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildReportServerPacket :: Get guild info reply.
type GuildReportServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReportServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Tag : field : string
	if err = writer.AddString(s.Tag); err != nil {
		return
	}
	writer.AddByte(255)
	// CreateDate : field : string
	if err = writer.AddString(s.CreateDate); err != nil {
		return
	}
	writer.AddByte(255)
	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	writer.AddByte(255)
	// Wealth : field : string
	if err = writer.AddString(s.Wealth); err != nil {
		return
	}
	writer.AddByte(255)
	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.Ranks[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	// StaffCount : length : short
	if err = writer.AddShort(s.StaffCount); err != nil {
		return
	}
	writer.AddByte(255)
	// Staff : array : GuildStaff
	for ndx := 0; ndx < s.StaffCount; ndx++ {
		if err = s.Staff[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *GuildReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildAgreeServerPacket :: Joined guild info.
type GuildAgreeServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAgreeServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// RecruiterId : field : short
	if err = writer.AddShort(s.RecruiterId); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildName : field : string
	if err = writer.AddString(s.GuildName); err != nil {
		return
	}
	writer.AddByte(255)
	// RankName : field : string
	if err = writer.AddString(s.RankName); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *GuildAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildAcceptServerPacket :: Update guild rank.
type GuildAcceptServerPacket struct {
	byteSize int

	Rank int
}

func (s GuildAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Rank : field : char
	s.Rank = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildKickServerPacket :: Left the guild.
type GuildKickServerPacket struct {
	byteSize int
}

func (s GuildKickServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildKickServerPacket) Action() net.PacketAction {
	return net.PacketAction_Kick
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildKickServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildKickServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *GuildKickServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellRequestServerPacket :: Nearby player chanting a spell.
type SpellRequestServerPacket struct {
	byteSize int

	PlayerId int
	SpellId  int
}

func (s SpellRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetSelfServerPacket :: Nearby player self-casted a spell.
type SpellTargetSelfServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetSelfServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellPlayerServerPacket :: Nearby player raising their arm to cast a spell (vestigial).
type SpellPlayerServerPacket struct {
	byteSize int

	PlayerId  int
	Direction protocol.Direction
}

func (s SpellPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellErrorServerPacket :: Show flood protection message (vestigial).
type SpellErrorServerPacket struct {
	byteSize int
}

func (s SpellErrorServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellErrorServerPacket) Action() net.PacketAction {
	return net.PacketAction_Error
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellErrorServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellErrorServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *SpellErrorServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AvatarAdminServerPacket :: Nearby player hit by a damage spell from a player.
type AvatarAdminServerPacket struct {
	byteSize int

	CasterId        int
	VictimId        int
	Damage          int
	CasterDirection protocol.Direction
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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AvatarAdminServerPacket) ByteSize() int {
	return s.byteSize
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
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// CasterDirection : field : Direction
	if err = writer.AddChar(int(s.CasterDirection)); err != nil {
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

	readerStartPosition := reader.Position()
	// CasterId : field : short
	s.CasterId = reader.GetShort()
	// VictimId : field : short
	s.VictimId = reader.GetShort()
	// Damage : field : three
	s.Damage = reader.GetThree()
	// CasterDirection : field : Direction
	s.CasterDirection = protocol.Direction(reader.GetChar())
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetGroupServerPacket :: Nearby player(s) hit by a group heal spell from a player.
type SpellTargetGroupServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetGroupServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetOtherServerPacket :: Nearby player hit by a heal spell from a player.
type SpellTargetOtherServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetOtherServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeRequestServerPacket :: Trade request from another player.
type TradeRequestServerPacket struct {
	byteSize int

	PartnerPlayerId   int
	PartnerPlayerName string
}

func (s TradeRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeRequestServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeRequestServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 138 : field : char
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

	readerStartPosition := reader.Position()
	// 138 : field : char
	reader.GetChar()
	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	// PartnerPlayerName : field : string
	if s.PartnerPlayerName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeOpenServerPacket :: Trade window opens.
type TradeOpenServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeOpenServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// YourPlayerId : field : short
	if err = writer.AddShort(s.YourPlayerId); err != nil {
		return
	}
	// YourPlayerName : field : string
	if err = writer.AddString(s.YourPlayerName); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *TradeOpenServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeReplyServerPacket :: Trade updated (items changed).
type TradeReplyServerPacket struct {
	byteSize int

	TradeData TradeItemData
}

func (s TradeReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeAdminServerPacket :: Trade updated (items changed while trade was accepted).
type TradeAdminServerPacket struct {
	byteSize int

	TradeData TradeItemData
}

func (s TradeAdminServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAdminServerPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeAdminServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeUseServerPacket :: Trade completed.
type TradeUseServerPacket struct {
	byteSize int

	TradeData TradeItemData
}

func (s TradeUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeUseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// TradeData : field : TradeItemData
	if err = s.TradeData.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeSpecServerPacket :: Own agree state updated.
type TradeSpecServerPacket struct {
	byteSize int

	Agree bool
}

func (s TradeSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Agree : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Agree = true
	} else {
		s.Agree = false
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeAgreeServerPacket :: Partner agree state updated.
type TradeAgreeServerPacket struct {
	byteSize int

	PartnerPlayerId int
	Agree           bool
}

func (s TradeAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	// Agree : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Agree = true
	} else {
		s.Agree = false
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeCloseServerPacket :: Partner closed trade window.
type TradeCloseServerPacket struct {
	byteSize int

	PartnerPlayerId int
}

func (s TradeCloseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeCloseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeCloseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PartnerPlayerId : field : short
	s.PartnerPlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcReplyServerPacket :: Nearby NPC hit by a player.
type NpcReplyServerPacket struct {
	byteSize int

	PlayerId            int
	PlayerDirection     protocol.Direction
	NpcIndex            int
	Damage              int
	HpPercentage        int
	KillStealProtection *NpcKillStealProtectionState // This field should be sent to the attacker, but not nearby players.
}

func (s NpcReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CastReplyServerPacket :: Nearby NPC hit by a spell from a player.
type CastReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CastReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcSpecServerPacket :: Nearby NPC killed by player.
type NpcSpecServerPacket struct {
	byteSize int

	NpcKilledData NpcKilledData
	Experience    *int
}

func (s NpcSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// NpcKilledData : field : NpcKilledData
	if err = s.NpcKilledData.Deserialize(reader); err != nil {
		return
	}
	// Experience : field : int
	if reader.Remaining() > 0 {
		s.Experience = new(int)
		*s.Experience = reader.GetInt()
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcAcceptServerPacket :: Nearby NPC killed by player and you leveled up.
type NpcAcceptServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CastSpecServerPacket :: Nearby NPC killed by player spell.
type CastSpecServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CastSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CastAcceptServerPacket :: Nearby NPC killed by player spell and you leveled up.
type CastAcceptServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CastAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcJunkServerPacket :: Clearing all boss children.
type NpcJunkServerPacket struct {
	byteSize int

	NpcId int
}

func (s NpcJunkServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcJunkServerPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcJunkServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// NpcId : field : short
	s.NpcId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcPlayerServerPacket :: Main NPC update message.
type NpcPlayerServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	writer.AddByte(255)
	// Attacks : array : NpcUpdateAttack
	for ndx := 0; ndx < len(s.Attacks); ndx++ {
		if err = s.Attacks[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
	// Chats : array : NpcUpdateChat
	for ndx := 0; ndx < len(s.Chats); ndx++ {
		if err = s.Chats[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
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

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Positions : array : NpcUpdatePosition
	PositionsRemaining := reader.Remaining()
	for ndx := 0; ndx < PositionsRemaining/4; ndx++ {
		s.Positions = append(s.Positions, NpcUpdatePosition{})
		if err = s.Positions[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Attacks : array : NpcUpdateAttack
	AttacksRemaining := reader.Remaining()
	for ndx := 0; ndx < AttacksRemaining/9; ndx++ {
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcDialogServerPacket :: NPC chat message.
type NpcDialogServerPacket struct {
	byteSize int

	NpcIndex int
	Message  string
}

func (s NpcDialogServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Npc
}

func (s NpcDialogServerPacket) Action() net.PacketAction {
	return net.PacketAction_Dialog
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcDialogServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestReportServerPacket :: NPC chat messages.
type QuestReportServerPacket struct {
	byteSize int

	NpcId    int
	Messages []string
}

func (s QuestReportServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestReportServerPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestReportServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *QuestReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// NpcId : field : short
	if err = writer.AddShort(s.NpcId); err != nil {
		return
	}
	writer.AddByte(255)
	// Messages : array : string
	for ndx := 0; ndx < len(s.Messages); ndx++ {
		if err = writer.AddString(s.Messages[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *QuestReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestDialogServerPacket :: Quest selection dialog.
type QuestDialogServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestDialogServerPacket) ByteSize() int {
	return s.byteSize
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
	writer.AddByte(255)
	// QuestEntries : array : DialogQuestEntry
	for ndx := 0; ndx < s.QuestCount; ndx++ {
		if err = s.QuestEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	// DialogEntries : array : DialogEntry
	for ndx := 0; ndx < len(s.DialogEntries); ndx++ {
		if err = s.DialogEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *QuestDialogServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestListServerPacket :: Quest history / progress reply.
type QuestListServerPacket struct {
	byteSize int

	Page        net.QuestPage
	QuestsCount int
	PageData    QuestListPageData
}

type QuestListPageData interface {
	protocol.EoData
}

type QuestListPageDataProgress struct {
	byteSize int

	QuestProgressEntries []QuestProgressEntry
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestListPageDataProgress) ByteSize() int {
	return s.byteSize
}

func (s *QuestListPageDataProgress) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuestProgressEntries : array : QuestProgressEntry
	for ndx := 0; ndx < len(s.QuestProgressEntries); ndx++ {
		if err = s.QuestProgressEntries[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *QuestListPageDataProgress) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// QuestProgressEntries : array : QuestProgressEntry
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.QuestProgressEntries = append(s.QuestProgressEntries, QuestProgressEntry{})
		if err = s.QuestProgressEntries[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

type QuestListPageDataHistory struct {
	byteSize int

	CompletedQuests []string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestListPageDataHistory) ByteSize() int {
	return s.byteSize
}

func (s *QuestListPageDataHistory) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CompletedQuests : array : string
	for ndx := 0; ndx < len(s.CompletedQuests); ndx++ {
		if err = writer.AddString(s.CompletedQuests[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *QuestListPageDataHistory) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CompletedQuests : array : string
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.CompletedQuests = append(s.CompletedQuests, "")
		if s.CompletedQuests[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s QuestListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestListServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemAcceptServerPacket :: Nearby player leveled up from quest.
type ItemAcceptServerPacket struct {
	byteSize int

	PlayerId int
}

func (s ItemAcceptServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemAcceptServerPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemAcceptServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ArenaDropServerPacket :: "Arena is blocked" message.
type ArenaDropServerPacket struct {
	byteSize int
}

func (s ArenaDropServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaDropServerPacket) Action() net.PacketAction {
	return net.PacketAction_Drop
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ArenaDropServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *ArenaDropServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// N : dummy : string
	if err = writer.AddString("N"); err != nil {
		return
	}
	return
}

func (s *ArenaDropServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// N : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ArenaUseServerPacket :: Arena start message.
type ArenaUseServerPacket struct {
	byteSize int

	PlayersCount int
}

func (s ArenaUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Arena
}

func (s ArenaUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ArenaUseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayersCount : field : char
	s.PlayersCount = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ArenaSpecServerPacket :: Arena kill message.
type ArenaSpecServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ArenaSpecServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *ArenaSpecServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	writer.AddByte(255)
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	writer.AddByte(255)
	// KillsCount : field : int
	if err = writer.AddInt(s.KillsCount); err != nil {
		return
	}
	writer.AddByte(255)
	// KillerName : field : string
	if err = writer.AddString(s.KillerName); err != nil {
		return
	}
	writer.AddByte(255)
	// VictimName : field : string
	if err = writer.AddString(s.VictimName); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *ArenaSpecServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ArenaAcceptServerPacket :: Arena win message.
type ArenaAcceptServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ArenaAcceptServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *ArenaAcceptServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// WinnerName : field : string
	if err = writer.AddString(s.WinnerName); err != nil {
		return
	}
	writer.AddByte(255)
	// KillsCount : field : int
	if err = writer.AddInt(s.KillsCount); err != nil {
		return
	}
	writer.AddByte(255)
	// KillerName : field : string
	if err = writer.AddString(s.KillerName); err != nil {
		return
	}
	writer.AddByte(255)
	// VictimName : field : string
	if err = writer.AddString(s.VictimName); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *ArenaAcceptServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MarriageOpenServerPacket :: Response from talking to a law NPC.
type MarriageOpenServerPacket struct {
	byteSize int

	SessionId int
}

func (s MarriageOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MarriageOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MarriageReplyServerPacket :: Reply to client Marriage-family packets.
type MarriageReplyServerPacket struct {
	byteSize int

	ReplyCode     MarriageReply
	ReplyCodeData MarriageReplyReplyCodeData
}

type MarriageReplyReplyCodeData interface {
	protocol.EoData
}

type MarriageReplyReplyCodeDataSuccess struct {
	byteSize int

	GoldAmount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MarriageReplyReplyCodeDataSuccess) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s MarriageReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MarriageReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ReplyCode : field : MarriageReply
	s.ReplyCode = MarriageReply(reader.GetShort())
	switch s.ReplyCode {
	case MarriageReply_Success:
		s.ReplyCodeData = &MarriageReplyReplyCodeDataSuccess{}
		if err = s.ReplyCodeData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestOpenServerPacket :: Response from talking to a priest NPC.
type PriestOpenServerPacket struct {
	byteSize int

	SessionId int
}

func (s PriestOpenServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestOpenServerPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestOpenServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestReplyServerPacket :: Reply to client Priest-family packets.
type PriestReplyServerPacket struct {
	byteSize int

	ReplyCode PriestReply
}

func (s PriestReplyServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestReplyServerPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ReplyCode : field : PriestReply
	s.ReplyCode = PriestReply(reader.GetShort())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestRequestServerPacket :: Wedding request.
type PriestRequestServerPacket struct {
	byteSize int

	SessionId   int
	PartnerName string
}

func (s PriestRequestServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestRequestServerPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestRequestServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// PartnerName : field : string
	if s.PartnerName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RecoverPlayerServerPacket :: HP/TP update.
type RecoverPlayerServerPacket struct {
	byteSize int

	Hp int
	Tp int
}

func (s RecoverPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RecoverPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RecoverAgreeServerPacket :: Nearby player gained HP.
type RecoverAgreeServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RecoverAgreeServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HealHp : field : int
	s.HealHp = reader.GetInt()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RecoverListServerPacket :: Stats update.
type RecoverListServerPacket struct {
	byteSize int

	ClassId int
	Stats   CharacterStatsUpdate
}

func (s RecoverListServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Recover
}

func (s RecoverListServerPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RecoverListServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// ClassId : field : short
	s.ClassId = reader.GetShort()
	// Stats : field : CharacterStatsUpdate
	if err = s.Stats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RecoverReplyServerPacket :: Karma/experience update.
type RecoverReplyServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RecoverReplyServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RecoverTargetGroupServerPacket :: Updated stats when levelling up from party experience.
type RecoverTargetGroupServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RecoverTargetGroupServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectUseServerPacket :: Map effect.
type EffectUseServerPacket struct {
	byteSize int

	Effect     MapEffect
	EffectData EffectUseEffectData
}

type EffectUseEffectData interface {
	protocol.EoData
}

type EffectUseEffectDataQuake struct {
	byteSize int

	QuakeStrength int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectUseEffectDataQuake) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// QuakeStrength : field : char
	s.QuakeStrength = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s EffectUseServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectUseServerPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectUseServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// Effect : field : MapEffect
	s.Effect = MapEffect(reader.GetChar())
	switch s.Effect {
	case MapEffect_Quake:
		s.EffectData = &EffectUseEffectDataQuake{}
		if err = s.EffectData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectAgreeServerPacket :: Effects playing on nearby tiles.
type EffectAgreeServerPacket struct {
	byteSize int

	Effects []TileEffect
}

func (s EffectAgreeServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectAgreeServerPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectAgreeServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *EffectAgreeServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Effects : array : TileEffect
	for ndx := 0; ndx < len(s.Effects); ndx++ {
		if err = s.Effects[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *EffectAgreeServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Effects : array : TileEffect
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Effects = append(s.Effects, TileEffect{})
		if err = s.Effects[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectTargetOtherServerPacket :: Map drain damage.
type EffectTargetOtherServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectTargetOtherServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectReportServerPacket :: Map spike timer.
type EffectReportServerPacket struct {
	byteSize int
}

func (s EffectReportServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectReportServerPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectReportServerPacket) ByteSize() int {
	return s.byteSize
}

func (s *EffectReportServerPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// S : dummy : string
	if err = writer.AddString("S"); err != nil {
		return
	}
	return
}

func (s *EffectReportServerPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// S : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectSpecServerPacket :: Taking spike or tp drain damage.
type EffectSpecServerPacket struct {
	byteSize int

	MapDamageType     MapDamageType
	MapDamageTypeData EffectSpecMapDamageTypeData
}

type EffectSpecMapDamageTypeData interface {
	protocol.EoData
}

type EffectSpecMapDamageTypeDataTpDrain struct {
	byteSize int

	TpDamage int
	Tp       int
	MaxTp    int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectSpecMapDamageTypeDataTpDrain) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// TpDamage : field : short
	s.TpDamage = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type EffectSpecMapDamageTypeDataSpikes struct {
	byteSize int

	HpDamage int
	Hp       int
	MaxHp    int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectSpecMapDamageTypeDataSpikes) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// HpDamage : field : short
	s.HpDamage = reader.GetShort()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s EffectSpecServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Effect
}

func (s EffectSpecServerPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectSpecServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EffectAdminServerPacket :: Nearby character taking spike damage.
type EffectAdminServerPacket struct {
	byteSize int

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

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EffectAdminServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MusicPlayerServerPacket :: Sound effect.
type MusicPlayerServerPacket struct {
	byteSize int

	SoundId int
}

func (s MusicPlayerServerPacket) Family() net.PacketFamily {
	return net.PacketFamily_Music
}

func (s MusicPlayerServerPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MusicPlayerServerPacket) ByteSize() int {
	return s.byteSize
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

	readerStartPosition := reader.Position()
	// SoundId : field : char
	s.SoundId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

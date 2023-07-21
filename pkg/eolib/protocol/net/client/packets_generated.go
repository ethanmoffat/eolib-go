package client

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	protocol "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
	net "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// InitInitClientPacket ::  Connection initialization request. This packet is unencrypted.
type InitInitClientPacket struct {
	Challenge int
	Version   net.Version

	HdidLength int
	Hdid       string
}

func (s *InitInitClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Challenge : field : three
	if err = writer.AddThree(s.Challenge); err != nil {
		return
	}

	// Version : field : Version
	if err = s.Version.Serialize(writer); err != nil {
		return
	}
	//  : field : char
	if err = writer.AddChar(112); err != nil {
		return
	}

	// HdidLength : length : char
	if err = writer.AddChar(s.HdidLength); err != nil {
		return
	}

	// Hdid : field : string
	if err = writer.AddFixedString(s.Hdid, s.HdidLength); err != nil {
		return
	}

	return
}

func (s *InitInitClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Challenge : field : three
	s.Challenge = reader.GetThree()
	// Version : field : Version
	if err = s.Version.Deserialize(reader); err != nil {
		return
	}
	//  : field : char
	reader.GetChar()
	// HdidLength : length : char
	s.HdidLength = reader.GetChar()
	// Hdid : field : string
	if s.Hdid, err = reader.GetFixedString(s.HdidLength); err != nil {
		return
	}

	return
}

// ConnectionAcceptClientPacket :: Confirm initialization data.
type ConnectionAcceptClientPacket struct {
	ClientEncryptionMultiple int
	ServerEncryptionMultiple int
	PlayerId                 int
}

func (s *ConnectionAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ClientEncryptionMultiple : field : short
	if err = writer.AddShort(s.ClientEncryptionMultiple); err != nil {
		return
	}

	// ServerEncryptionMultiple : field : short
	if err = writer.AddShort(s.ServerEncryptionMultiple); err != nil {
		return
	}

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *ConnectionAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ClientEncryptionMultiple : field : short
	s.ClientEncryptionMultiple = reader.GetShort()
	// ServerEncryptionMultiple : field : short
	s.ServerEncryptionMultiple = reader.GetShort()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// ConnectionPingClientPacket :: Ping reply.
type ConnectionPingClientPacket struct {
}

func (s *ConnectionPingClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("k"); err != nil {
		return
	}

	return
}

func (s *ConnectionPingClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AccountRequestClientPacket :: Request creating an account.
type AccountRequestClientPacket struct {
	Username string
}

func (s *AccountRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}

	return
}

func (s *AccountRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Username : field : string
	if s.Username, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AccountCreateClientPacket :: Confirm creating an account.
type AccountCreateClientPacket struct {
	SessionId int
	Username  string
	Password  string
	FullName  string
	Location  string
	Email     string
	Computer  string
	Hdid      string
}

func (s *AccountCreateClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Password : field : string
	if err = writer.AddString(s.Password); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// FullName : field : string
	if err = writer.AddString(s.FullName); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Location : field : string
	if err = writer.AddString(s.Location); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Email : field : string
	if err = writer.AddString(s.Email); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Computer : field : string
	if err = writer.AddString(s.Computer); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Hdid : field : string
	if err = writer.AddString(s.Hdid); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *AccountCreateClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Username : field : string
	if s.Username, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Password : field : string
	if s.Password, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// FullName : field : string
	if s.FullName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Location : field : string
	if s.Location, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Email : field : string
	if s.Email, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Computer : field : string
	if s.Computer, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Hdid : field : string
	if s.Hdid, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// AccountAgreeClientPacket :: Change password.
type AccountAgreeClientPacket struct {
	Username    string
	OldPassword string
	NewPassword string
}

func (s *AccountAgreeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// OldPassword : field : string
	if err = writer.AddString(s.OldPassword); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// NewPassword : field : string
	if err = writer.AddString(s.NewPassword); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *AccountAgreeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// Username : field : string
	if s.Username, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// OldPassword : field : string
	if s.OldPassword, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// NewPassword : field : string
	if s.NewPassword, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// CharacterRequestClientPacket :: Request to create a character.
type CharacterRequestClientPacket struct {
	RequestString string
}

func (s *CharacterRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// RequestString : field : string
	if err = writer.AddString(s.RequestString); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *CharacterRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// RequestString : field : string
	if s.RequestString, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// CharacterCreateClientPacket :: Confirm creating a character.
type CharacterCreateClientPacket struct {
	SessionId int
	Gender    protocol.Gender
	HairStyle int
	HairColor int
	Skin      int
	Name      string
}

func (s *CharacterCreateClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	// Gender : field : Gender:short
	if err = writer.AddChar(int(s.Gender)); err != nil {
		return
	}

	// HairStyle : field : short
	if err = writer.AddShort(s.HairStyle); err != nil {
		return
	}

	// HairColor : field : short
	if err = writer.AddShort(s.HairColor); err != nil {
		return
	}

	// Skin : field : short
	if err = writer.AddShort(s.Skin); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *CharacterCreateClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// Gender : field : Gender:short
	s.Gender = protocol.Gender(reader.GetChar())
	// HairStyle : field : short
	s.HairStyle = reader.GetShort()
	// HairColor : field : short
	s.HairColor = reader.GetShort()
	// Skin : field : short
	s.Skin = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// CharacterTakeClientPacket :: Request to delete a character from an account.
type CharacterTakeClientPacket struct {
	CharacterId int
}

func (s *CharacterTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}

	return
}

func (s *CharacterTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// CharacterId : field : int
	s.CharacterId = reader.GetInt()

	return
}

// CharacterRemoveClientPacket :: Confirm deleting character from an account.
type CharacterRemoveClientPacket struct {
	SessionId   int
	CharacterId int
}

func (s *CharacterRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
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

func (s *CharacterRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()

	return
}

// LoginRequestClientPacket :: Login request.
type LoginRequestClientPacket struct {
	Username string
	Password string
}

func (s *LoginRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Password : field : string
	if err = writer.AddString(s.Password); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *LoginRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// Username : field : string
	if s.Username, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Password : field : string
	if s.Password, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// WelcomeRequestClientPacket :: Selected a character.
type WelcomeRequestClientPacket struct {
	CharacterId int
}

func (s *WelcomeRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}

	return
}

func (s *WelcomeRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// CharacterId : field : int
	s.CharacterId = reader.GetInt()

	return
}

// WelcomeMsgClientPacket :: Entering game.
type WelcomeMsgClientPacket struct {
	SessionId   int
	CharacterId int
}

func (s *WelcomeMsgClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}

	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}

	return
}

func (s *WelcomeMsgClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : three
	s.SessionId = reader.GetThree()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()

	return
}

// WelcomeAgreeClientPacket :: Requesting a file.
type WelcomeAgreeClientPacket struct {
	FileType     FileType
	SessionId    int
	FileTypeData WelcomeAgreeFileTypeData
}

type WelcomeAgreeFileTypeData interface {
	protocol.EoData
}

type WelcomeAgreeFileTypeDataEmf struct {
	FileId int
}

func (s *WelcomeAgreeFileTypeDataEmf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : short
	if err = writer.AddShort(s.FileId); err != nil {
		return
	}

	return
}

func (s *WelcomeAgreeFileTypeDataEmf) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileId : field : short
	s.FileId = reader.GetShort()

	return
}

type WelcomeAgreeFileTypeDataEif struct {
	FileId int
}

func (s *WelcomeAgreeFileTypeDataEif) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}

	return
}

func (s *WelcomeAgreeFileTypeDataEif) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileId : field : char
	s.FileId = reader.GetChar()

	return
}

type WelcomeAgreeFileTypeDataEnf struct {
	FileId int
}

func (s *WelcomeAgreeFileTypeDataEnf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}

	return
}

func (s *WelcomeAgreeFileTypeDataEnf) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileId : field : char
	s.FileId = reader.GetChar()

	return
}

type WelcomeAgreeFileTypeDataEsf struct {
	FileId int
}

func (s *WelcomeAgreeFileTypeDataEsf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}

	return
}

func (s *WelcomeAgreeFileTypeDataEsf) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileId : field : char
	s.FileId = reader.GetChar()

	return
}

type WelcomeAgreeFileTypeDataEcf struct {
	FileId int
}

func (s *WelcomeAgreeFileTypeDataEcf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}

	return
}

func (s *WelcomeAgreeFileTypeDataEcf) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileId : field : char
	s.FileId = reader.GetChar()

	return
}

func (s *WelcomeAgreeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileType : field : FileType
	if err = writer.AddChar(int(s.FileType)); err != nil {
		return
	}

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	switch s.FileType {
	case File_Emf:
		switch s.FileTypeData.(type) {
		case *WelcomeAgreeFileTypeDataEmf:
			if err = s.FileTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.FileType)
			return
		}
	case File_Eif:
		switch s.FileTypeData.(type) {
		case *WelcomeAgreeFileTypeDataEif:
			if err = s.FileTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.FileType)
			return
		}
	case File_Enf:
		switch s.FileTypeData.(type) {
		case *WelcomeAgreeFileTypeDataEnf:
			if err = s.FileTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.FileType)
			return
		}
	case File_Esf:
		switch s.FileTypeData.(type) {
		case *WelcomeAgreeFileTypeDataEsf:
			if err = s.FileTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.FileType)
			return
		}
	case File_Ecf:
		switch s.FileTypeData.(type) {
		case *WelcomeAgreeFileTypeDataEcf:
			if err = s.FileTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.FileType)
			return
		}
	}
	return
}

func (s *WelcomeAgreeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// FileType : field : FileType
	s.FileType = FileType(reader.GetChar())
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	switch s.FileType {
	case File_Emf:
		s.FileTypeData = &WelcomeAgreeFileTypeDataEmf{}
		if err = s.FileTypeData.Deserialize(reader); err != nil {
			return
		}
	case File_Eif:
		s.FileTypeData = &WelcomeAgreeFileTypeDataEif{}
		if err = s.FileTypeData.Deserialize(reader); err != nil {
			return
		}
	case File_Enf:
		s.FileTypeData = &WelcomeAgreeFileTypeDataEnf{}
		if err = s.FileTypeData.Deserialize(reader); err != nil {
			return
		}
	case File_Esf:
		s.FileTypeData = &WelcomeAgreeFileTypeDataEsf{}
		if err = s.FileTypeData.Deserialize(reader); err != nil {
			return
		}
	case File_Ecf:
		s.FileTypeData = &WelcomeAgreeFileTypeDataEcf{}
		if err = s.FileTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// AdminInteractTellClientPacket :: Talk to admin.
type AdminInteractTellClientPacket struct {
	Message string
}

func (s *AdminInteractTellClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *AdminInteractTellClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AdminInteractReportClientPacket :: Report character.
type AdminInteractReportClientPacket struct {
	Reportee string
	Message  string
}

func (s *AdminInteractReportClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Reportee : field : string
	if err = writer.AddString(s.Reportee); err != nil {
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

func (s *AdminInteractReportClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// Reportee : field : string
	if s.Reportee, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetChunkedReadingMode(false)

	return
}

// GlobalRemoveClientPacket :: Enable whispers.
type GlobalRemoveClientPacket struct {
}

func (s *GlobalRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("n"); err != nil {
		return
	}

	return
}

func (s *GlobalRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GlobalPlayerClientPacket :: Disable whispers.
type GlobalPlayerClientPacket struct {
}

func (s *GlobalPlayerClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("y"); err != nil {
		return
	}

	return
}

func (s *GlobalPlayerClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GlobalOpenClientPacket :: Opened global tab.
type GlobalOpenClientPacket struct {
}

func (s *GlobalOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("y"); err != nil {
		return
	}

	return
}

func (s *GlobalOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GlobalCloseClientPacket :: Closed global tab.
type GlobalCloseClientPacket struct {
}

func (s *GlobalCloseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : string
	if err = writer.AddString("n"); err != nil {
		return
	}

	return
}

func (s *GlobalCloseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkRequestClientPacket :: Guild chat message.
type TalkRequestClientPacket struct {
	Message string
}

func (s *TalkRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkOpenClientPacket :: Party chat message.
type TalkOpenClientPacket struct {
	Message string
}

func (s *TalkOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkMsgClientPacket :: Global chat message.
type TalkMsgClientPacket struct {
	Message string
}

func (s *TalkMsgClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkMsgClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkTellClientPacket :: Private chat message.
type TalkTellClientPacket struct {
	Name    string
	Message string
}

func (s *TalkTellClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
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

func (s *TalkTellClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetChunkedReadingMode(false)

	return
}

// TalkReportClientPacket :: Public chat message.
type TalkReportClientPacket struct {
	Message string
}

func (s *TalkReportClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkReportClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkPlayerClientPacket :: Public chat message - alias of TALK_REPORT (vestigial).
type TalkPlayerClientPacket struct {
	Message string
}

func (s *TalkPlayerClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkPlayerClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkUseClientPacket :: Public chat message - alias of TALK_REPORT (vestigial).
type TalkUseClientPacket struct {
	Message string
}

func (s *TalkUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkAdminClientPacket :: Admin chat message.
type TalkAdminClientPacket struct {
	Message string
}

func (s *TalkAdminClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkAdminClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// TalkAnnounceClientPacket :: Admin announcement.
type TalkAnnounceClientPacket struct {
	Message string
}

func (s *TalkAnnounceClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}

	return
}

func (s *TalkAnnounceClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	return
}

// AttackUseClientPacket :: Attacking.
type AttackUseClientPacket struct {
	Direction protocol.Direction
	Timestamp int
}

func (s *AttackUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	return
}

func (s *AttackUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()

	return
}

// ChairRequestClientPacket :: Sitting on a chair.
type ChairRequestClientPacket struct {
	SitAction     SitAction
	SitActionData ChairRequestSitActionData
}

type ChairRequestSitActionData interface {
	protocol.EoData
}

type ChairRequestSitActionDataSit struct {
	Coords protocol.Coords
}

func (s *ChairRequestSitActionDataSit) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChairRequestSitActionDataSit) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

func (s *ChairRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SitAction : field : SitAction
	if err = writer.AddChar(int(s.SitAction)); err != nil {
		return
	}

	switch s.SitAction {
	case SitAction_Sit:
		switch s.SitActionData.(type) {
		case *ChairRequestSitActionDataSit:
			if err = s.SitActionData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.SitAction)
			return
		}
	}
	return
}

func (s *ChairRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SitAction : field : SitAction
	s.SitAction = SitAction(reader.GetChar())
	switch s.SitAction {
	case SitAction_Sit:
		s.SitActionData = &ChairRequestSitActionDataSit{}
		if err = s.SitActionData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// SitRequestClientPacket :: Sit/stand request.
type SitRequestClientPacket struct {
	SitAction     SitAction
	SitActionData SitRequestSitActionData
}

type SitRequestSitActionData interface {
	protocol.EoData
}

type SitRequestSitActionDataSit struct {
	CursorCoords protocol.Coords // The coordinates of the map cursor.
}

func (s *SitRequestSitActionDataSit) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CursorCoords : field : Coords
	if err = s.CursorCoords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *SitRequestSitActionDataSit) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// CursorCoords : field : Coords
	if err = s.CursorCoords.Deserialize(reader); err != nil {
		return
	}

	return
}

func (s *SitRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SitAction : field : SitAction
	if err = writer.AddChar(int(s.SitAction)); err != nil {
		return
	}

	switch s.SitAction {
	case SitAction_Sit:
		switch s.SitActionData.(type) {
		case *SitRequestSitActionDataSit:
			if err = s.SitActionData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.SitAction)
			return
		}
	}
	return
}

func (s *SitRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SitAction : field : SitAction
	s.SitAction = SitAction(reader.GetChar())
	switch s.SitAction {
	case SitAction_Sit:
		s.SitActionData = &SitRequestSitActionDataSit{}
		if err = s.SitActionData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EmoteReportClientPacket :: Doing an emote.
type EmoteReportClientPacket struct {
	Emote protocol.Emote
}

func (s *EmoteReportClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Emote : field : Emote
	if err = writer.AddChar(int(s.Emote)); err != nil {
		return
	}

	return
}

func (s *EmoteReportClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Emote : field : Emote
	s.Emote = protocol.Emote(reader.GetChar())

	return
}

// FacePlayerClientPacket :: Facing a direction.
type FacePlayerClientPacket struct {
	Direction protocol.Direction
}

func (s *FacePlayerClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}

	return
}

func (s *FacePlayerClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// WalkAdminClientPacket :: Walking with #nowall.
type WalkAdminClientPacket struct {
	WalkAction WalkAction
}

func (s *WalkAdminClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkAdminClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}

	return
}

// WalkSpecClientPacket :: Walking through a player.
type WalkSpecClientPacket struct {
	WalkAction WalkAction
}

func (s *WalkSpecClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkSpecClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}

	return
}

// WalkPlayerClientPacket :: Walking.
type WalkPlayerClientPacket struct {
	WalkAction WalkAction
}

func (s *WalkPlayerClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkPlayerClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}

	return
}

// BankOpenClientPacket :: Talked to a banker NPC.
type BankOpenClientPacket struct {
	NpcIndex int
}

func (s *BankOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *BankOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// BankAddClientPacket :: Depositing gold.
type BankAddClientPacket struct {
	Amount int
}

func (s *BankAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}

	return
}

func (s *BankAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Amount : field : int
	s.Amount = reader.GetInt()

	return
}

// BankTakeClientPacket :: Withdrawing gold.
type BankTakeClientPacket struct {
	Amount int
}

func (s *BankTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}

	return
}

func (s *BankTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Amount : field : int
	s.Amount = reader.GetInt()

	return
}

// BarberBuyClientPacket :: Purchasing a hair-style.
type BarberBuyClientPacket struct {
	HairStyle int
	HairColor int
	SessionId int
}

func (s *BarberBuyClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HairStyle : field : char
	if err = writer.AddChar(s.HairStyle); err != nil {
		return
	}

	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *BarberBuyClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// HairStyle : field : char
	s.HairStyle = reader.GetChar()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// BarberOpenClientPacket :: Talking to a barber NPC.
type BarberOpenClientPacket struct {
	NpcIndex int
}

func (s *BarberOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *BarberOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// LockerAddClientPacket :: Adding an item to a bank locker.
type LockerAddClientPacket struct {
	LockerCoords protocol.Coords
	DepositItem  net.ThreeItem
}

func (s *LockerAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Serialize(writer); err != nil {
		return
	}
	// DepositItem : field : ThreeItem
	if err = s.DepositItem.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *LockerAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	// DepositItem : field : ThreeItem
	if err = s.DepositItem.Deserialize(reader); err != nil {
		return
	}

	return
}

// LockerTakeClientPacket :: Taking an item from a bank locker.
type LockerTakeClientPacket struct {
	LockerCoords protocol.Coords
	TakeItemId   int
}

func (s *LockerTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Serialize(writer); err != nil {
		return
	}
	// TakeItemId : field : short
	if err = writer.AddShort(s.TakeItemId); err != nil {
		return
	}

	return
}

func (s *LockerTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	// TakeItemId : field : short
	s.TakeItemId = reader.GetShort()

	return
}

// LockerOpenClientPacket :: Opening a bank locker.
type LockerOpenClientPacket struct {
	LockerCoords protocol.Coords
}

func (s *LockerOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *LockerOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}

	return
}

// LockerBuyClientPacket :: Buying a locker space upgrade from a banker NPC.
type LockerBuyClientPacket struct {
}

func (s *LockerBuyClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : char
	if err = writer.AddChar(1); err != nil {
		return
	}

	return
}

func (s *LockerBuyClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : char
	reader.GetChar()

	return
}

// CitizenRequestClientPacket :: Request sleeping at an inn.
type CitizenRequestClientPacket struct {
	SessionId  int
	BehaviorId int
}

func (s *CitizenRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}

	return
}

func (s *CitizenRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()

	return
}

// CitizenAcceptClientPacket :: Confirm sleeping at an inn.
type CitizenAcceptClientPacket struct {
	SessionId  int
	BehaviorId int
}

func (s *CitizenAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}

	return
}

func (s *CitizenAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()

	return
}

// CitizenReplyClientPacket :: Subscribing to a town.
type CitizenReplyClientPacket struct {
	SessionId  int
	BehaviorId int
	Answers    []string
}

func (s *CitizenReplyClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Answers : array : string
	for ndx := 0; ndx < 3; ndx++ {
		if ndx > 0 {
			writer.AddByte(0xFF)
		}

		if err = writer.AddString(s.Answers[ndx]); err != nil {
			return
		}

	}

	writer.SanitizeStrings = false
	return
}

func (s *CitizenReplyClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Answers : array : string
	for ndx := 0; ndx < 3; ndx++ {
		if s.Answers[ndx], err = reader.GetString(); err != nil {
			return
		}

		if ndx+1 < 3 {
			if err = reader.NextChunk(); err != nil {
				return
			}
		}
	}

	reader.SetChunkedReadingMode(false)

	return
}

// CitizenRemoveClientPacket :: Giving up citizenship of a town.
type CitizenRemoveClientPacket struct {
	BehaviorId int
}

func (s *CitizenRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}

	return
}

func (s *CitizenRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()

	return
}

// CitizenOpenClientPacket :: Talking to a citizenship NPC.
type CitizenOpenClientPacket struct {
	NpcIndex int
}

func (s *CitizenOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *CitizenOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// ShopCreateClientPacket :: Crafting an item from a shop.
type ShopCreateClientPacket struct {
	CraftItemId int
	SessionId   int
}

func (s *ShopCreateClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CraftItemId : field : short
	if err = writer.AddShort(s.CraftItemId); err != nil {
		return
	}

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *ShopCreateClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// CraftItemId : field : short
	s.CraftItemId = reader.GetShort()
	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// ShopBuyClientPacket :: Purchasing an item from a shop.
type ShopBuyClientPacket struct {
	BuyItem   net.Item
	SessionId int
}

func (s *ShopBuyClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BuyItem : field : Item
	if err = s.BuyItem.Serialize(writer); err != nil {
		return
	}
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *ShopBuyClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// BuyItem : field : Item
	if err = s.BuyItem.Deserialize(reader); err != nil {
		return
	}
	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// ShopSellClientPacket :: Selling an item to a shop.
type ShopSellClientPacket struct {
	SellItem  net.Item
	SessionId int
}

func (s *ShopSellClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SellItem : field : Item
	if err = s.SellItem.Serialize(writer); err != nil {
		return
	}
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *ShopSellClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SellItem : field : Item
	if err = s.SellItem.Deserialize(reader); err != nil {
		return
	}
	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// ShopOpenClientPacket :: Talking to a shop NPC.
type ShopOpenClientPacket struct {
	NpcIndex int
}

func (s *ShopOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *ShopOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// StatSkillOpenClientPacket :: Talking to a skill master NPC.
type StatSkillOpenClientPacket struct {
	NpcIndex int
}

func (s *StatSkillOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *StatSkillOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// StatSkillTakeClientPacket :: Learning a skill from a skill master NPC.
type StatSkillTakeClientPacket struct {
	SessionId int
	SpellId   int
}

func (s *StatSkillTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	return
}

func (s *StatSkillTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

// StatSkillRemoveClientPacket :: Forgetting a skill at a skill master NPC.
type StatSkillRemoveClientPacket struct {
	SessionId int
	SpellId   int
}

func (s *StatSkillRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	return
}

func (s *StatSkillRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

// StatSkillAddClientPacket :: Spending a stat point on a stat or skill.
type StatSkillAddClientPacket struct {
	ActionType     TrainType
	ActionTypeData StatSkillAddActionTypeData
}

type StatSkillAddActionTypeData interface {
	protocol.EoData
}

type StatSkillAddActionTypeDataStat struct {
	StatId StatId
}

func (s *StatSkillAddActionTypeDataStat) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// StatId : field : StatId
	if err = writer.AddShort(int(s.StatId)); err != nil {
		return
	}

	return
}

func (s *StatSkillAddActionTypeDataStat) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// StatId : field : StatId
	s.StatId = StatId(reader.GetShort())

	return
}

type StatSkillAddActionTypeDataSkill struct {
	SpellId int
}

func (s *StatSkillAddActionTypeDataSkill) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	return
}

func (s *StatSkillAddActionTypeDataSkill) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()

	return
}

func (s *StatSkillAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ActionType : field : TrainType
	if err = writer.AddChar(int(s.ActionType)); err != nil {
		return
	}

	switch s.ActionType {
	case Train_Stat:
		switch s.ActionTypeData.(type) {
		case *StatSkillAddActionTypeDataStat:
			if err = s.ActionTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ActionType)
			return
		}
	case Train_Skill:
		switch s.ActionTypeData.(type) {
		case *StatSkillAddActionTypeDataSkill:
			if err = s.ActionTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ActionType)
			return
		}
	}
	return
}

func (s *StatSkillAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ActionType : field : TrainType
	s.ActionType = TrainType(reader.GetChar())
	switch s.ActionType {
	case Train_Stat:
		s.ActionTypeData = &StatSkillAddActionTypeDataStat{}
		if err = s.ActionTypeData.Deserialize(reader); err != nil {
			return
		}
	case Train_Skill:
		s.ActionTypeData = &StatSkillAddActionTypeDataSkill{}
		if err = s.ActionTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// StatSkillJunkClientPacket :: Resetting stats at a skill master.
type StatSkillJunkClientPacket struct {
	SessionId int
}

func (s *StatSkillJunkClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *StatSkillJunkClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// ItemUseClientPacket :: Using an item.
type ItemUseClientPacket struct {
	ItemId int
}

func (s *ItemUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}

	return
}

func (s *ItemUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()

	return
}

// ItemDropClientPacket :: Dropping items on the ground.
type ItemDropClientPacket struct {
	Item   net.ThreeItem
	Coords ByteCoords //  The official client sends 255 byte values for the coords if an item is dropped via. the GUI button. 255 values here should be interpreted to mean "drop at current coords". Otherwise, the x and y fields contain encoded numbers that must be explicitly. decoded to get the actual x and y values.
}

func (s *ItemDropClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Item : field : ThreeItem
	if err = s.Item.Serialize(writer); err != nil {
		return
	}
	// Coords : field : ByteCoords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemDropClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Item : field : ThreeItem
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// Coords : field : ByteCoords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemJunkClientPacket :: Junking items.
type ItemJunkClientPacket struct {
	Item net.Item
}

func (s *ItemJunkClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Item : field : Item
	if err = s.Item.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemJunkClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Item : field : Item
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}

	return
}

// ItemGetClientPacket :: Taking items from the ground.
type ItemGetClientPacket struct {
	ItemIndex int
}

func (s *ItemGetClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemIndex : field : short
	if err = writer.AddShort(s.ItemIndex); err != nil {
		return
	}

	return
}

func (s *ItemGetClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()

	return
}

// BoardRemoveClientPacket :: Removing a post from a town board.
type BoardRemoveClientPacket struct {
	BoardId int
	PostId  int
}

func (s *BoardRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}

	// PostId : field : short
	if err = writer.AddShort(s.PostId); err != nil {
		return
	}

	return
}

func (s *BoardRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// BoardId : field : short
	s.BoardId = reader.GetShort()
	// PostId : field : short
	s.PostId = reader.GetShort()

	return
}

// BoardCreateClientPacket :: Posting a new message to a town board.
type BoardCreateClientPacket struct {
	BoardId     int
	PostSubject string
	PostBody    string
}

func (s *BoardCreateClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// PostSubject : field : string
	if err = writer.AddString(s.PostSubject); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// PostBody : field : string
	if err = writer.AddString(s.PostBody); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *BoardCreateClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// BoardId : field : short
	s.BoardId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// PostSubject : field : string
	if s.PostSubject, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// PostBody : field : string
	if s.PostBody, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// BoardTakeClientPacket :: Reading a post on a town board.
type BoardTakeClientPacket struct {
	BoardId int
	PostId  int
}

func (s *BoardTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}

	// PostId : field : short
	if err = writer.AddShort(s.PostId); err != nil {
		return
	}

	return
}

func (s *BoardTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// BoardId : field : short
	s.BoardId = reader.GetShort()
	// PostId : field : short
	s.PostId = reader.GetShort()

	return
}

// BoardOpenClientPacket :: Opening a town board.
type BoardOpenClientPacket struct {
	BoardId int
}

func (s *BoardOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}

	return
}

func (s *BoardOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// BoardId : field : short
	s.BoardId = reader.GetShort()

	return
}

// JukeboxOpenClientPacket :: Opening the jukebox listing.
type JukeboxOpenClientPacket struct {
	Coords protocol.Coords
}

func (s *JukeboxOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *JukeboxOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// JukeboxMsgClientPacket :: Requesting a song on a jukebox.
type JukeboxMsgClientPacket struct {
	TrackId int
}

func (s *JukeboxMsgClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	// TrackId : field : short
	if err = writer.AddShort(s.TrackId); err != nil {
		return
	}

	writer.SanitizeStrings = false
	return
}

func (s *JukeboxMsgClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	//  : field : char
	reader.GetChar()
	//  : field : char
	reader.GetChar()
	// TrackId : field : short
	s.TrackId = reader.GetShort()
	reader.SetChunkedReadingMode(false)

	return
}

// JukeboxUseClientPacket :: Playing a note with the bard skill.
type JukeboxUseClientPacket struct {
	InstrumentId int
	NoteId       int
}

func (s *JukeboxUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

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

func (s *JukeboxUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// InstrumentId : field : char
	s.InstrumentId = reader.GetChar()
	// NoteId : field : char
	s.NoteId = reader.GetChar()

	return
}

// WarpAcceptClientPacket :: Accept a warp request from the server.
type WarpAcceptClientPacket struct {
	MapId     int
	SessionId int
}

func (s *WarpAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	return
}

func (s *WarpAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// MapId : field : short
	s.MapId = reader.GetShort()
	// SessionId : field : short
	s.SessionId = reader.GetShort()

	return
}

// WarpTakeClientPacket :: Request to download a copy of the map.
type WarpTakeClientPacket struct {
	MapId     int
	SessionId int
}

func (s *WarpTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	return
}

func (s *WarpTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// MapId : field : short
	s.MapId = reader.GetShort()
	// SessionId : field : short
	s.SessionId = reader.GetShort()

	return
}

// PaperdollRequestClientPacket :: Request for a player's paperdoll.
type PaperdollRequestClientPacket struct {
	PlayerId int
}

func (s *PaperdollRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *PaperdollRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// PaperdollRemoveClientPacket :: Unequipping an item.
type PaperdollRemoveClientPacket struct {
	ItemId int
	SubLoc int
}

func (s *PaperdollRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}

	// SubLoc : field : char
	if err = writer.AddChar(s.SubLoc); err != nil {
		return
	}

	return
}

func (s *PaperdollRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()

	return
}

// PaperdollAddClientPacket :: Equipping an item.
type PaperdollAddClientPacket struct {
	ItemId int
	SubLoc int
}

func (s *PaperdollAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}

	// SubLoc : field : char
	if err = writer.AddChar(s.SubLoc); err != nil {
		return
	}

	return
}

func (s *PaperdollAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()

	return
}

// BookRequestClientPacket :: Request for a player's book.
type BookRequestClientPacket struct {
	PlayerId int
}

func (s *BookRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *BookRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// MessagePingClientPacket :: #ping command request.
type MessagePingClientPacket struct {
}

func (s *MessagePingClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}

	return
}

func (s *MessagePingClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : short
	reader.GetShort()

	return
}

// PlayersAcceptClientPacket :: #find command request.
type PlayersAcceptClientPacket struct {
	Name string
}

func (s *PlayersAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}

	return
}

func (s *PlayersAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	return
}

// PlayersRequestClientPacket :: Requesting a list of online players.
type PlayersRequestClientPacket struct {
}

func (s *PlayersRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}

	return
}

func (s *PlayersRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// PlayersListClientPacket :: Requesting a list of online friends.
type PlayersListClientPacket struct {
}

func (s *PlayersListClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}

	return
}

func (s *PlayersListClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// DoorOpenClientPacket :: Opening a door.
type DoorOpenClientPacket struct {
	Coords protocol.Coords
}

func (s *DoorOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *DoorOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// ChestOpenClientPacket :: Opening a chest.
type ChestOpenClientPacket struct {
	Coords protocol.Coords
}

func (s *ChestOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChestOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}

// ChestAddClientPacket :: Placing an item in to a chest.
type ChestAddClientPacket struct {
	Coords  protocol.Coords
	AddItem net.ThreeItem
}

func (s *ChestAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// AddItem : field : ThreeItem
	if err = s.AddItem.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChestAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// AddItem : field : ThreeItem
	if err = s.AddItem.Deserialize(reader); err != nil {
		return
	}

	return
}

// ChestTakeClientPacket :: Taking an item from a chest.
type ChestTakeClientPacket struct {
	Coords     protocol.Coords
	TakeItemId int
}

func (s *ChestTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// TakeItemId : field : short
	if err = writer.AddShort(s.TakeItemId); err != nil {
		return
	}

	return
}

func (s *ChestTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// TakeItemId : field : short
	s.TakeItemId = reader.GetShort()

	return
}

// RefreshRequestClientPacket :: Requesting new info about nearby objects.
type RefreshRequestClientPacket struct {
}

func (s *RefreshRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}

	return
}

func (s *RefreshRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : byte
	reader.GetByte()

	return
}

// RangeRequestClientPacket :: Requesting info about nearby players and NPCs.
type RangeRequestClientPacket struct {
	PlayerIds  []int
	NpcIndexes []int
}

func (s *RangeRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
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

	writer.SanitizeStrings = false
	return
}

func (s *RangeRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// PlayerIds : array : short
	for ndx := 0; ndx < reader.Remaining()/2; ndx++ {
		s.PlayerIds[ndx] = reader.GetShort()
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// NpcIndexes : array : char
	for ndx := 0; ndx < reader.Remaining()/1; ndx++ {
		s.NpcIndexes[ndx] = reader.GetChar()
	}

	reader.SetChunkedReadingMode(false)

	return
}

// PlayerRangeRequestClientPacket :: Requesting info about nearby players.
type PlayerRangeRequestClientPacket struct {
	PlayerIds []int
}

func (s *PlayerRangeRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerIds : array : short
	for ndx := 0; ndx < len(s.PlayerIds); ndx++ {
		if err = writer.AddShort(s.PlayerIds[ndx]); err != nil {
			return
		}

	}

	return
}

func (s *PlayerRangeRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// PlayerIds : array : short
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.PlayerIds[ndx] = reader.GetShort()
	}

	return
}

// NpcRangeRequestClientPacket :: Requesting info about nearby NPCs.
type NpcRangeRequestClientPacket struct {
	NpcIndexesLength int

	NpcIndexes []int
}

func (s *NpcRangeRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndexesLength : length : char
	if err = writer.AddChar(s.NpcIndexesLength); err != nil {
		return
	}

	//  : field : byte
	if err = writer.AddByte(255); err != nil {
		return
	}

	// NpcIndexes : array : char
	for ndx := 0; ndx < s.NpcIndexesLength; ndx++ {
		if err = writer.AddChar(s.NpcIndexes[ndx]); err != nil {
			return
		}

	}

	return
}

func (s *NpcRangeRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndexesLength : length : char
	s.NpcIndexesLength = reader.GetChar()
	//  : field : byte
	reader.GetByte()
	// NpcIndexes : array : char
	for ndx := 0; ndx < s.NpcIndexesLength; ndx++ {
		s.NpcIndexes[ndx] = reader.GetChar()
	}

	return
}

// PartyRequestClientPacket :: Send party invite / join request.
type PartyRequestClientPacket struct {
	RequestType net.PartyRequestType
	PlayerId    int
}

func (s *PartyRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// RequestType : field : PartyRequestType
	if err = writer.AddChar(int(s.RequestType)); err != nil {
		return
	}

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *PartyRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// PartyAcceptClientPacket :: Accept party invite / join request.
type PartyAcceptClientPacket struct {
	RequestType     net.PartyRequestType
	InviterPlayerId int
}

func (s *PartyAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
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

	return
}

func (s *PartyAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()

	return
}

// PartyRemoveClientPacket :: Remove player from a party.
type PartyRemoveClientPacket struct {
	PlayerId int
}

func (s *PartyRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *PartyRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// PartyTakeClientPacket :: Request updated party info.
type PartyTakeClientPacket struct {
	MembersCount int
}

func (s *PartyTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MembersCount : field : char
	if err = writer.AddChar(s.MembersCount); err != nil {
		return
	}

	return
}

func (s *PartyTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// MembersCount : field : char
	s.MembersCount = reader.GetChar()

	return
}

// GuildRequestClientPacket :: Requested to create a guild.
type GuildRequestClientPacket struct {
	SessionId int
	GuildTag  string
	GuildName string
}

func (s *GuildRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
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
	writer.SanitizeStrings = false
	return
}

func (s *GuildRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
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
	reader.SetChunkedReadingMode(false)

	return
}

// GuildAcceptClientPacket :: Accept pending guild creation invite.
type GuildAcceptClientPacket struct {
	InviterPlayerId int
}

func (s *GuildAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : int
	if err = writer.AddInt(20202); err != nil {
		return
	}

	// InviterPlayerId : field : short
	if err = writer.AddShort(s.InviterPlayerId); err != nil {
		return
	}

	return
}

func (s *GuildAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : field : int
	reader.GetInt()
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()

	return
}

// GuildRemoveClientPacket :: Leave guild.
type GuildRemoveClientPacket struct {
	SessionId int
}

func (s *GuildRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *GuildRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// GuildAgreeClientPacket :: Update the guild description or rank list.
type GuildAgreeClientPacket struct {
	SessionId    int
	InfoType     GuildInfoType
	InfoTypeData GuildAgreeInfoTypeData
}

type GuildAgreeInfoTypeData interface {
	protocol.EoData
}

type GuildAgreeInfoTypeDataDescription struct {
	Description string
}

func (s *GuildAgreeInfoTypeDataDescription) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}

	return
}

func (s *GuildAgreeInfoTypeDataDescription) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	return
}

type GuildAgreeInfoTypeDataRanks struct {
	Ranks []string
}

func (s *GuildAgreeInfoTypeDataRanks) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if err = writer.AddString(s.Ranks[ndx]); err != nil {
			return
		}

		writer.AddByte(0xFF)
	}

	return
}

func (s *GuildAgreeInfoTypeDataRanks) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if s.Ranks[ndx], err = reader.GetString(); err != nil {
			return
		}

	}

	return
}

func (s *GuildAgreeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// InfoType : field : GuildInfoType
	if err = writer.AddShort(int(s.InfoType)); err != nil {
		return
	}

	switch s.InfoType {
	case GuildInfo_Description:
		switch s.InfoTypeData.(type) {
		case *GuildAgreeInfoTypeDataDescription:
			if err = s.InfoTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.InfoType)
			return
		}
	case GuildInfo_Ranks:
		switch s.InfoTypeData.(type) {
		case *GuildAgreeInfoTypeDataRanks:
			if err = s.InfoTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.InfoType)
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *GuildAgreeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// InfoType : field : GuildInfoType
	s.InfoType = GuildInfoType(reader.GetShort())
	switch s.InfoType {
	case GuildInfo_Description:
		s.InfoTypeData = &GuildAgreeInfoTypeDataDescription{}
		if err = s.InfoTypeData.Deserialize(reader); err != nil {
			return
		}
	case GuildInfo_Ranks:
		s.InfoTypeData = &GuildAgreeInfoTypeDataRanks{}
		if err = s.InfoTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	reader.SetChunkedReadingMode(false)

	return
}

// GuildCreateClientPacket :: Final confirm creating a guild.
type GuildCreateClientPacket struct {
	SessionId   int
	GuildTag    string
	GuildName   string
	Description string
}

func (s *GuildCreateClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
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
	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *GuildCreateClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
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
	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// GuildPlayerClientPacket :: Request to join a guild.
type GuildPlayerClientPacket struct {
	SessionId     int
	GuildTag      string
	RecruiterName string
}

func (s *GuildPlayerClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// RecruiterName : field : string
	if err = writer.AddString(s.RecruiterName); err != nil {
		return
	}

	writer.AddByte(0xFF)
	writer.SanitizeStrings = false
	return
}

func (s *GuildPlayerClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
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
	// RecruiterName : field : string
	if s.RecruiterName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetChunkedReadingMode(false)

	return
}

// GuildTakeClientPacket :: Request guild description, rank list, or bank balance.
type GuildTakeClientPacket struct {
	SessionId int
	InfoType  GuildInfoType
}

func (s *GuildTakeClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// InfoType : field : GuildInfoType
	if err = writer.AddShort(int(s.InfoType)); err != nil {
		return
	}

	return
}

func (s *GuildTakeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// InfoType : field : GuildInfoType
	s.InfoType = GuildInfoType(reader.GetShort())

	return
}

// GuildUseClientPacket :: Accepted a join request.
type GuildUseClientPacket struct {
	PlayerId int
}

func (s *GuildUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *GuildUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// GuildBuyClientPacket :: Deposit gold in to the guild bank.
type GuildBuyClientPacket struct {
	SessionId  int
	GoldAmount int
}

func (s *GuildBuyClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// GoldAmount : field : int
	if err = writer.AddInt(s.GoldAmount); err != nil {
		return
	}

	return
}

func (s *GuildBuyClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()

	return
}

// GuildOpenClientPacket :: Talking to a guild master NPC.
type GuildOpenClientPacket struct {
	NpcIndex int
}

func (s *GuildOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *GuildOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// GuildTellClientPacket :: Requested member list of a guild.
type GuildTellClientPacket struct {
	SessionId     int
	GuildIdentity string
}

func (s *GuildTellClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// GuildIdentity : field : string
	if err = writer.AddString(s.GuildIdentity); err != nil {
		return
	}

	return
}

func (s *GuildTellClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GuildReportClientPacket :: Requested general information of a guild.
type GuildReportClientPacket struct {
	SessionId     int
	GuildIdentity string
}

func (s *GuildReportClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// GuildIdentity : field : string
	if err = writer.AddString(s.GuildIdentity); err != nil {
		return
	}

	return
}

func (s *GuildReportClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GuildJunkClientPacket :: Disband guild.
type GuildJunkClientPacket struct {
	SessionId int
}

func (s *GuildJunkClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *GuildJunkClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

// GuildKickClientPacket :: Kick member from guild.
type GuildKickClientPacket struct {
	SessionId  int
	MemberName string
}

func (s *GuildKickClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// MemberName : field : string
	if err = writer.AddString(s.MemberName); err != nil {
		return
	}

	return
}

func (s *GuildKickClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// MemberName : field : string
	if s.MemberName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// GuildRankClientPacket :: Update a member's rank.
type GuildRankClientPacket struct {
	SessionId  int
	Rank       int
	MemberName string
}

func (s *GuildRankClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	// Rank : field : char
	if err = writer.AddChar(s.Rank); err != nil {
		return
	}

	// MemberName : field : string
	if err = writer.AddString(s.MemberName); err != nil {
		return
	}

	return
}

func (s *GuildRankClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// Rank : field : char
	s.Rank = reader.GetChar()
	// MemberName : field : string
	if s.MemberName, err = reader.GetString(); err != nil {
		return
	}

	return
}

// SpellRequestClientPacket :: Begin spell chanting.
type SpellRequestClientPacket struct {
	SpellId   int
	Timestamp int
}

func (s *SpellRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	return
}

func (s *SpellRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()

	return
}

// SpellTargetSelfClientPacket :: Self-targeted spell cast.
type SpellTargetSelfClientPacket struct {
	Direction protocol.Direction
	SpellId   int
	Timestamp int
}

func (s *SpellTargetSelfClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	return
}

func (s *SpellTargetSelfClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()

	return
}

// SpellTargetOtherClientPacket :: Targeted spell cast.
type SpellTargetOtherClientPacket struct {
	TargetType        SpellTargetType
	PreviousTimestamp int
	SpellId           int
	VictimId          int
	Timestamp         int
}

func (s *SpellTargetOtherClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// TargetType : field : SpellTargetType
	if err = writer.AddChar(int(s.TargetType)); err != nil {
		return
	}

	// PreviousTimestamp : field : three
	if err = writer.AddThree(s.PreviousTimestamp); err != nil {
		return
	}

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	// VictimId : field : short
	if err = writer.AddShort(s.VictimId); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	return
}

func (s *SpellTargetOtherClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// TargetType : field : SpellTargetType
	s.TargetType = SpellTargetType(reader.GetChar())
	// PreviousTimestamp : field : three
	s.PreviousTimestamp = reader.GetThree()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// VictimId : field : short
	s.VictimId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()

	return
}

// SpellTargetGroupClientPacket :: Group spell cast.
type SpellTargetGroupClientPacket struct {
	SpellId   int
	Timestamp int
}

func (s *SpellTargetGroupClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	return
}

func (s *SpellTargetGroupClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()

	return
}

// SpellUseClientPacket :: Raise arm to cast a spell (vestigial).
type SpellUseClientPacket struct {
	Direction protocol.Direction
}

func (s *SpellUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}

	return
}

func (s *SpellUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())

	return
}

// TradeRequestClientPacket :: Requesting a trade with another player.
type TradeRequestClientPacket struct {
	PlayerId int
}

func (s *TradeRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : char
	if err = writer.AddChar(138); err != nil {
		return
	}

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *TradeRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : field : char
	reader.GetChar()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// TradeAcceptClientPacket :: Accepting a trade request.
type TradeAcceptClientPacket struct {
	PlayerId int
}

func (s *TradeAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}

	return
}

func (s *TradeAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : field : char
	reader.GetChar()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()

	return
}

// TradeRemoveClientPacket :: Remove an item from the trade screen.
type TradeRemoveClientPacket struct {
	ItemId int
}

func (s *TradeRemoveClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}

	return
}

func (s *TradeRemoveClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()

	return
}

// TradeAgreeClientPacket :: Mark trade as agreed.
type TradeAgreeClientPacket struct {
	Agree bool
}

func (s *TradeAgreeClientPacket) Serialize(writer data.EoWriter) (err error) {
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

func (s *TradeAgreeClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Agree : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Agree = true
	} else {
		s.Agree = false
	}

	return
}

// TradeAddClientPacket :: Add an item to the trade screen.
type TradeAddClientPacket struct {
	AddItem net.Item
}

func (s *TradeAddClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// AddItem : field : Item
	if err = s.AddItem.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *TradeAddClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// AddItem : field : Item
	if err = s.AddItem.Deserialize(reader); err != nil {
		return
	}

	return
}

// TradeCloseClientPacket :: Cancel the trade.
type TradeCloseClientPacket struct {
}

func (s *TradeCloseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : dummy : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	return
}

func (s *TradeCloseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	//  : dummy : char
	reader.GetChar()

	return
}

// QuestUseClientPacket :: Talking to a quest NPC.
type QuestUseClientPacket struct {
	NpcIndex int
	QuestId  int //  Quest ID is 0 unless the player explicitly selects a quest from the quest switcher.
}

func (s *QuestUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	// QuestId : field : short
	if err = writer.AddShort(s.QuestId); err != nil {
		return
	}

	return
}

func (s *QuestUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// QuestId : field : short
	s.QuestId = reader.GetShort()

	return
}

// QuestAcceptClientPacket :: Response to a quest NPC dialog.
type QuestAcceptClientPacket struct {
	SessionId     int
	DialogId      int
	QuestId       int
	NpcIndex      int
	ReplyType     DialogReply
	ReplyTypeData QuestAcceptReplyTypeData
}

type QuestAcceptReplyTypeData interface {
	protocol.EoData
}

type QuestAcceptReplyTypeDataLink struct {
	Action int
}

func (s *QuestAcceptReplyTypeDataLink) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Action : field : char
	if err = writer.AddChar(s.Action); err != nil {
		return
	}

	return
}

func (s *QuestAcceptReplyTypeDataLink) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Action : field : char
	s.Action = reader.GetChar()

	return
}

func (s *QuestAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	// DialogId : field : short
	if err = writer.AddShort(s.DialogId); err != nil {
		return
	}

	// QuestId : field : short
	if err = writer.AddShort(s.QuestId); err != nil {
		return
	}

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	// ReplyType : field : DialogReply
	if err = writer.AddChar(int(s.ReplyType)); err != nil {
		return
	}

	switch s.ReplyType {
	case DialogReply_Link:
		switch s.ReplyTypeData.(type) {
		case *QuestAcceptReplyTypeDataLink:
			if err = s.ReplyTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ReplyType)
			return
		}
	}
	return
}

func (s *QuestAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// DialogId : field : short
	s.DialogId = reader.GetShort()
	// QuestId : field : short
	s.QuestId = reader.GetShort()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// ReplyType : field : DialogReply
	s.ReplyType = DialogReply(reader.GetChar())
	switch s.ReplyType {
	case DialogReply_Link:
		s.ReplyTypeData = &QuestAcceptReplyTypeDataLink{}
		if err = s.ReplyTypeData.Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// QuestListClientPacket :: Quest history / progress request.
type QuestListClientPacket struct {
	Page net.QuestPage
}

func (s *QuestListClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Page : field : QuestPage
	if err = writer.AddChar(int(s.Page)); err != nil {
		return
	}

	return
}

func (s *QuestListClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// Page : field : QuestPage
	s.Page = net.QuestPage(reader.GetChar())

	return
}

// MarriageOpenClientPacket :: Talking to a law NPC.
type MarriageOpenClientPacket struct {
	NpcIndex int
}

func (s *MarriageOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *MarriageOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()

	return
}

// MarriageRequestClientPacket :: Requesting marriage approval.
type MarriageRequestClientPacket struct {
	RequestType MarriageRequestType
	SessionId   int
	Name        string
}

func (s *MarriageRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// RequestType : field : MarriageRequestType
	if err = writer.AddChar(int(s.RequestType)); err != nil {
		return
	}

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}

	writer.SanitizeStrings = false
	return
}

func (s *MarriageRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// RequestType : field : MarriageRequestType
	s.RequestType = MarriageRequestType(reader.GetChar())
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	reader.SetChunkedReadingMode(false)

	return
}

// PriestAcceptClientPacket :: Accepting a marriage request.
type PriestAcceptClientPacket struct {
	SessionId int
}

func (s *PriestAcceptClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}

	return
}

func (s *PriestAcceptClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : short
	s.SessionId = reader.GetShort()

	return
}

// PriestOpenClientPacket :: Talking to a priest NPC.
type PriestOpenClientPacket struct {
	NpcIndex int
}

func (s *PriestOpenClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : int
	if err = writer.AddInt(s.NpcIndex); err != nil {
		return
	}

	return
}

func (s *PriestOpenClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// NpcIndex : field : int
	s.NpcIndex = reader.GetInt()

	return
}

// PriestRequestClientPacket :: Requesting marriage at a priest.
type PriestRequestClientPacket struct {
	SessionId int
	Name      string
}

func (s *PriestRequestClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	writer.AddByte(0xFF)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}

	writer.SanitizeStrings = false
	return
}

func (s *PriestRequestClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	reader.SetChunkedReadingMode(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	reader.SetChunkedReadingMode(false)

	return
}

// PriestUseClientPacket :: Saying "I do" at a wedding.
type PriestUseClientPacket struct {
	SessionId int
}

func (s *PriestUseClientPacket) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}

	return
}

func (s *PriestUseClientPacket) Deserialize(reader data.EoReader) (err error) {
	oldChunkedReadingMode := reader.GetChunkedReadingMode()
	defer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()

	// SessionId : field : int
	s.SessionId = reader.GetInt()

	return
}

package client

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net"
)

// InitInitClientPacket ::  Connection initialization request. This packet is unencrypted.
type InitInitClientPacket struct {
	byteSize int

	Challenge int
	Version   net.Version

	Hdid string
}

func (s InitInitClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Init
}

func (s InitInitClientPacket) Action() net.PacketAction {
	return net.PacketAction_Init
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *InitInitClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *InitInitClientPacket) Serialize(writer *data.EoWriter) (err error) {
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
	// 112 : field : char
	if err = writer.AddChar(112); err != nil {
		return
	}
	// HdidLength : length : char
	if err = writer.AddChar(len(s.Hdid)); err != nil {
		return
	}
	// Hdid : field : string
	if err = writer.AddFixedString(s.Hdid, len(s.Hdid)); err != nil {
		return
	}
	return
}

func (s *InitInitClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Challenge : field : three
	s.Challenge = reader.GetThree()
	// Version : field : Version
	if err = s.Version.Deserialize(reader); err != nil {
		return
	}
	// 112 : field : char
	reader.GetChar()
	// HdidLength : length : char
	hdidLength := reader.GetChar()
	// Hdid : field : string
	if s.Hdid, err = reader.GetFixedString(hdidLength); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ConnectionAcceptClientPacket :: Confirm initialization data.
type ConnectionAcceptClientPacket struct {
	byteSize int

	ClientEncryptionMultiple int
	ServerEncryptionMultiple int
	PlayerId                 int
}

func (s ConnectionAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Connection
}

func (s ConnectionAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ConnectionAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ConnectionAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ConnectionAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ClientEncryptionMultiple : field : short
	s.ClientEncryptionMultiple = reader.GetShort()
	// ServerEncryptionMultiple : field : short
	s.ServerEncryptionMultiple = reader.GetShort()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ConnectionPingClientPacket :: Ping reply.
type ConnectionPingClientPacket struct {
	byteSize int
}

func (s ConnectionPingClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Connection
}

func (s ConnectionPingClientPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ConnectionPingClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ConnectionPingClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// k : dummy : string
	if err = writer.AddString("k"); err != nil {
		return
	}
	return
}

func (s *ConnectionPingClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// k : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AccountRequestClientPacket :: Request creating an account.
type AccountRequestClientPacket struct {
	byteSize int

	Username string
}

func (s AccountRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Account
}

func (s AccountRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AccountRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}
	return
}

func (s *AccountRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Username : field : string
	if s.Username, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AccountCreateClientPacket :: Confirm creating an account.
type AccountCreateClientPacket struct {
	byteSize int

	SessionId int
	Username  string
	Password  string
	FullName  string
	Location  string
	Email     string
	Computer  string
	Hdid      string
}

func (s AccountCreateClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Account
}

func (s AccountCreateClientPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountCreateClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AccountCreateClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	writer.AddByte(255)
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}
	writer.AddByte(255)
	// Password : field : string
	if err = writer.AddString(s.Password); err != nil {
		return
	}
	writer.AddByte(255)
	// FullName : field : string
	if err = writer.AddString(s.FullName); err != nil {
		return
	}
	writer.AddByte(255)
	// Location : field : string
	if err = writer.AddString(s.Location); err != nil {
		return
	}
	writer.AddByte(255)
	// Email : field : string
	if err = writer.AddString(s.Email); err != nil {
		return
	}
	writer.AddByte(255)
	// Computer : field : string
	if err = writer.AddString(s.Computer); err != nil {
		return
	}
	writer.AddByte(255)
	// Hdid : field : string
	if err = writer.AddString(s.Hdid); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *AccountCreateClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AccountAgreeClientPacket :: Change password.
type AccountAgreeClientPacket struct {
	byteSize int

	Username    string
	OldPassword string
	NewPassword string
}

func (s AccountAgreeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Account
}

func (s AccountAgreeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AccountAgreeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AccountAgreeClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}
	writer.AddByte(255)
	// OldPassword : field : string
	if err = writer.AddString(s.OldPassword); err != nil {
		return
	}
	writer.AddByte(255)
	// NewPassword : field : string
	if err = writer.AddString(s.NewPassword); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *AccountAgreeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterRequestClientPacket :: Request to create a character.
type CharacterRequestClientPacket struct {
	byteSize int

	RequestString string
}

func (s CharacterRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CharacterRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// RequestString : field : string
	if err = writer.AddString(s.RequestString); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *CharacterRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// RequestString : field : string
	if s.RequestString, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterCreateClientPacket :: Confirm creating a character.
type CharacterCreateClientPacket struct {
	byteSize int

	SessionId int
	Gender    protocol.Gender
	HairStyle int
	HairColor int
	Skin      int
	Name      string
}

func (s CharacterCreateClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterCreateClientPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterCreateClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CharacterCreateClientPacket) Serialize(writer *data.EoWriter) (err error) {
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
	writer.AddByte(255)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *CharacterCreateClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterTakeClientPacket :: Request to delete a character from an account.
type CharacterTakeClientPacket struct {
	byteSize int

	CharacterId int
}

func (s CharacterTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CharacterTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}
	return
}

func (s *CharacterTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterRemoveClientPacket :: Confirm deleting character from an account.
type CharacterRemoveClientPacket struct {
	byteSize int

	SessionId   int
	CharacterId int
}

func (s CharacterRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Character
}

func (s CharacterRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CharacterRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *CharacterRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// LoginRequestClientPacket :: Login request.
type LoginRequestClientPacket struct {
	byteSize int

	Username string
	Password string
}

func (s LoginRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Login
}

func (s LoginRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LoginRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *LoginRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Username : field : string
	if err = writer.AddString(s.Username); err != nil {
		return
	}
	writer.AddByte(255)
	// Password : field : string
	if err = writer.AddString(s.Password); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *LoginRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeRequestClientPacket :: Selected a character.
type WelcomeRequestClientPacket struct {
	byteSize int

	CharacterId int
}

func (s WelcomeRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharacterId : field : int
	if err = writer.AddInt(s.CharacterId); err != nil {
		return
	}
	return
}

func (s *WelcomeRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeMsgClientPacket :: Entering game.
type WelcomeMsgClientPacket struct {
	byteSize int

	SessionId   int
	CharacterId int
}

func (s WelcomeMsgClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeMsgClientPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeMsgClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeMsgClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *WelcomeMsgClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	// CharacterId : field : int
	s.CharacterId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WelcomeAgreeClientPacket :: Requesting a file.
type WelcomeAgreeClientPacket struct {
	byteSize int

	FileType     FileType
	SessionId    int
	FileTypeData WelcomeAgreeFileTypeData
}

type WelcomeAgreeFileTypeData interface {
	protocol.EoData
}

type WelcomeAgreeFileTypeDataEmf struct {
	byteSize int

	FileId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeFileTypeDataEmf) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeFileTypeDataEmf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : short
	if err = writer.AddShort(s.FileId); err != nil {
		return
	}
	return
}

func (s *WelcomeAgreeFileTypeDataEmf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : short
	s.FileId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type WelcomeAgreeFileTypeDataEif struct {
	byteSize int

	FileId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeFileTypeDataEif) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeFileTypeDataEif) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}
	return
}

func (s *WelcomeAgreeFileTypeDataEif) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : char
	s.FileId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type WelcomeAgreeFileTypeDataEnf struct {
	byteSize int

	FileId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeFileTypeDataEnf) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeFileTypeDataEnf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}
	return
}

func (s *WelcomeAgreeFileTypeDataEnf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : char
	s.FileId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type WelcomeAgreeFileTypeDataEsf struct {
	byteSize int

	FileId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeFileTypeDataEsf) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeFileTypeDataEsf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}
	return
}

func (s *WelcomeAgreeFileTypeDataEsf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : char
	s.FileId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type WelcomeAgreeFileTypeDataEcf struct {
	byteSize int

	FileId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeFileTypeDataEcf) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeFileTypeDataEcf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}
	return
}

func (s *WelcomeAgreeFileTypeDataEcf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : char
	s.FileId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s WelcomeAgreeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Welcome
}

func (s WelcomeAgreeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WelcomeAgreeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WelcomeAgreeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *WelcomeAgreeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AdminInteractTellClientPacket :: Talk to admin.
type AdminInteractTellClientPacket struct {
	byteSize int

	Message string
}

func (s AdminInteractTellClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractTellClientPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractTellClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractTellClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *AdminInteractTellClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// AdminInteractReportClientPacket :: Report character.
type AdminInteractReportClientPacket struct {
	byteSize int

	Reportee string
	Message  string
}

func (s AdminInteractReportClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_AdminInteract
}

func (s AdminInteractReportClientPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AdminInteractReportClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AdminInteractReportClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Reportee : field : string
	if err = writer.AddString(s.Reportee); err != nil {
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

func (s *AdminInteractReportClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GlobalRemoveClientPacket :: Enable whispers.
type GlobalRemoveClientPacket struct {
	byteSize int
}

func (s GlobalRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Global
}

func (s GlobalRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GlobalRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GlobalRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// n : dummy : string
	if err = writer.AddString("n"); err != nil {
		return
	}
	return
}

func (s *GlobalRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// n : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GlobalPlayerClientPacket :: Disable whispers.
type GlobalPlayerClientPacket struct {
	byteSize int
}

func (s GlobalPlayerClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Global
}

func (s GlobalPlayerClientPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GlobalPlayerClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GlobalPlayerClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// y : dummy : string
	if err = writer.AddString("y"); err != nil {
		return
	}
	return
}

func (s *GlobalPlayerClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// y : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GlobalOpenClientPacket :: Opened global tab.
type GlobalOpenClientPacket struct {
	byteSize int
}

func (s GlobalOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Global
}

func (s GlobalOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GlobalOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GlobalOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// y : dummy : string
	if err = writer.AddString("y"); err != nil {
		return
	}
	return
}

func (s *GlobalOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// y : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GlobalCloseClientPacket :: Closed global tab.
type GlobalCloseClientPacket struct {
	byteSize int
}

func (s GlobalCloseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Global
}

func (s GlobalCloseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GlobalCloseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GlobalCloseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// n : dummy : string
	if err = writer.AddString("n"); err != nil {
		return
	}
	return
}

func (s *GlobalCloseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// n : dummy : string
	if _, err = reader.GetString(); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkRequestClientPacket :: Guild chat message.
type TalkRequestClientPacket struct {
	byteSize int

	Message string
}

func (s TalkRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkOpenClientPacket :: Party chat message.
type TalkOpenClientPacket struct {
	byteSize int

	Message string
}

func (s TalkOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkMsgClientPacket :: Global chat message.
type TalkMsgClientPacket struct {
	byteSize int

	Message string
}

func (s TalkMsgClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkMsgClientPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkMsgClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkMsgClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkMsgClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkTellClientPacket :: Private chat message.
type TalkTellClientPacket struct {
	byteSize int

	Name    string
	Message string
}

func (s TalkTellClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkTellClientPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkTellClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkTellClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
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

func (s *TalkTellClientPacket) Deserialize(reader *data.EoReader) (err error) {
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
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TalkReportClientPacket :: Public chat message.
type TalkReportClientPacket struct {
	byteSize int

	Message string
}

func (s TalkReportClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkReportClientPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkReportClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkReportClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkReportClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkPlayerClientPacket :: Public chat message - alias of TALK_REPORT (vestigial).
type TalkPlayerClientPacket struct {
	byteSize int

	Message string
}

func (s TalkPlayerClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkPlayerClientPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkPlayerClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkPlayerClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkPlayerClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkUseClientPacket :: Public chat message - alias of TALK_REPORT (vestigial).
type TalkUseClientPacket struct {
	byteSize int

	Message string
}

func (s TalkUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkAdminClientPacket :: Admin chat message.
type TalkAdminClientPacket struct {
	byteSize int

	Message string
}

func (s TalkAdminClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAdminClientPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkAdminClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkAdminClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkAdminClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TalkAnnounceClientPacket :: Admin announcement.
type TalkAnnounceClientPacket struct {
	byteSize int

	Message string
}

func (s TalkAnnounceClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Talk
}

func (s TalkAnnounceClientPacket) Action() net.PacketAction {
	return net.PacketAction_Announce
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TalkAnnounceClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TalkAnnounceClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	return
}

func (s *TalkAnnounceClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// AttackUseClientPacket :: Attacking.
type AttackUseClientPacket struct {
	byteSize int

	Direction protocol.Direction
	Timestamp int
}

func (s AttackUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Attack
}

func (s AttackUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AttackUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *AttackUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *AttackUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChairRequestClientPacket :: Sitting on a chair.
type ChairRequestClientPacket struct {
	byteSize int

	SitAction     SitAction
	SitActionData ChairRequestSitActionData
}

type ChairRequestSitActionData interface {
	protocol.EoData
}

type ChairRequestSitActionDataSit struct {
	byteSize int

	Coords protocol.Coords
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairRequestSitActionDataSit) ByteSize() int {
	return s.byteSize
}

func (s *ChairRequestSitActionDataSit) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChairRequestSitActionDataSit) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s ChairRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chair
}

func (s ChairRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChairRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ChairRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ChairRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SitAction : field : SitAction
	s.SitAction = SitAction(reader.GetChar())
	switch s.SitAction {
	case SitAction_Sit:
		s.SitActionData = &ChairRequestSitActionDataSit{}
		if err = s.SitActionData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SitRequestClientPacket :: Sit/stand request.
type SitRequestClientPacket struct {
	byteSize int

	SitAction     SitAction
	SitActionData SitRequestSitActionData
}

type SitRequestSitActionData interface {
	protocol.EoData
}

type SitRequestSitActionDataSit struct {
	byteSize int

	CursorCoords protocol.Coords // The coordinates of the map cursor.
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitRequestSitActionDataSit) ByteSize() int {
	return s.byteSize
}

func (s *SitRequestSitActionDataSit) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CursorCoords : field : Coords
	if err = s.CursorCoords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *SitRequestSitActionDataSit) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CursorCoords : field : Coords
	if err = s.CursorCoords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s SitRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Sit
}

func (s SitRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SitRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SitRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *SitRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SitAction : field : SitAction
	s.SitAction = SitAction(reader.GetChar())
	switch s.SitAction {
	case SitAction_Sit:
		s.SitActionData = &SitRequestSitActionDataSit{}
		if err = s.SitActionData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EmoteReportClientPacket :: Doing an emote.
type EmoteReportClientPacket struct {
	byteSize int

	Emote protocol.Emote
}

func (s EmoteReportClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Emote
}

func (s EmoteReportClientPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EmoteReportClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *EmoteReportClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Emote : field : Emote
	if err = writer.AddChar(int(s.Emote)); err != nil {
		return
	}
	return
}

func (s *EmoteReportClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Emote : field : Emote
	s.Emote = protocol.Emote(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// FacePlayerClientPacket :: Facing a direction.
type FacePlayerClientPacket struct {
	byteSize int

	Direction protocol.Direction
}

func (s FacePlayerClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Face
}

func (s FacePlayerClientPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *FacePlayerClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *FacePlayerClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *FacePlayerClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkAdminClientPacket :: Walking with #nowall.
type WalkAdminClientPacket struct {
	byteSize int

	WalkAction WalkAction
}

func (s WalkAdminClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkAdminClientPacket) Action() net.PacketAction {
	return net.PacketAction_Admin
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkAdminClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WalkAdminClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkAdminClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkSpecClientPacket :: Walking through a player.
type WalkSpecClientPacket struct {
	byteSize int

	WalkAction WalkAction
}

func (s WalkSpecClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkSpecClientPacket) Action() net.PacketAction {
	return net.PacketAction_Spec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkSpecClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WalkSpecClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkSpecClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkPlayerClientPacket :: Walking.
type WalkPlayerClientPacket struct {
	byteSize int

	WalkAction WalkAction
}

func (s WalkPlayerClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Walk
}

func (s WalkPlayerClientPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkPlayerClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WalkPlayerClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// WalkAction : field : WalkAction
	if err = s.WalkAction.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkPlayerClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// WalkAction : field : WalkAction
	if err = s.WalkAction.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BankOpenClientPacket :: Talked to a banker NPC.
type BankOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s BankOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BankOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BankOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *BankOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BankAddClientPacket :: Depositing gold.
type BankAddClientPacket struct {
	byteSize int

	Amount    int
	SessionId int
}

func (s BankAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BankAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BankAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}
	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}
	return
}

func (s *BankAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Amount : field : int
	s.Amount = reader.GetInt()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BankTakeClientPacket :: Withdrawing gold.
type BankTakeClientPacket struct {
	byteSize int

	Amount    int
	SessionId int
}

func (s BankTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Bank
}

func (s BankTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BankTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BankTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}
	// SessionId : field : three
	if err = writer.AddThree(s.SessionId); err != nil {
		return
	}
	return
}

func (s *BankTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Amount : field : int
	s.Amount = reader.GetInt()
	// SessionId : field : three
	s.SessionId = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BarberBuyClientPacket :: Purchasing a hair-style.
type BarberBuyClientPacket struct {
	byteSize int

	HairStyle int
	HairColor int
	SessionId int
}

func (s BarberBuyClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberBuyClientPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BarberBuyClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BarberBuyClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *BarberBuyClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// HairStyle : field : char
	s.HairStyle = reader.GetChar()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BarberOpenClientPacket :: Talking to a barber NPC.
type BarberOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s BarberOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Barber
}

func (s BarberOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BarberOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BarberOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *BarberOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerAddClientPacket :: Adding an item to a bank locker.
type LockerAddClientPacket struct {
	byteSize int

	LockerCoords protocol.Coords
	DepositItem  net.ThreeItem
}

func (s LockerAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *LockerAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *LockerAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	// DepositItem : field : ThreeItem
	if err = s.DepositItem.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerTakeClientPacket :: Taking an item from a bank locker.
type LockerTakeClientPacket struct {
	byteSize int

	LockerCoords protocol.Coords
	TakeItemId   int
}

func (s LockerTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *LockerTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *LockerTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	// TakeItemId : field : short
	s.TakeItemId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerOpenClientPacket :: Opening a bank locker.
type LockerOpenClientPacket struct {
	byteSize int

	LockerCoords protocol.Coords
}

func (s LockerOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *LockerOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LockerCoords : field : Coords
	if err = s.LockerCoords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *LockerOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// LockerCoords : field : Coords
	if err = s.LockerCoords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LockerBuyClientPacket :: Buying a locker space upgrade from a banker NPC.
type LockerBuyClientPacket struct {
	byteSize int
}

func (s LockerBuyClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Locker
}

func (s LockerBuyClientPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LockerBuyClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *LockerBuyClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 1 : dummy : char
	if err = writer.AddChar(1); err != nil {
		return
	}
	return
}

func (s *LockerBuyClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 1 : dummy : char
	reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenRequestClientPacket :: Request sleeping at an inn.
type CitizenRequestClientPacket struct {
	byteSize int

	SessionId  int
	BehaviorId int
}

func (s CitizenRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CitizenRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *CitizenRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenAcceptClientPacket :: Confirm sleeping at an inn.
type CitizenAcceptClientPacket struct {
	byteSize int

	SessionId  int
	BehaviorId int
}

func (s CitizenAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CitizenAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *CitizenAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenReplyClientPacket :: Subscribing to a town.
type CitizenReplyClientPacket struct {
	byteSize int

	SessionId  int
	BehaviorId int
	Answers    []string
}

func (s CitizenReplyClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenReplyClientPacket) Action() net.PacketAction {
	return net.PacketAction_Reply
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenReplyClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CitizenReplyClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	writer.AddByte(255)
	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	writer.AddByte(255)
	// Answers : array : string
	for ndx := 0; ndx < 3; ndx++ {
		if len(s.Answers) != 3 {
			err = fmt.Errorf("expected Answers with length 3, got %d", len(s.Answers))
			return
		}

		if ndx > 0 {
			writer.AddByte(255)
		}

		if err = writer.AddString(s.Answers[ndx]); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *CitizenReplyClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
		s.Answers = append(s.Answers, "")
		if s.Answers[ndx], err = reader.GetString(); err != nil {
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

// CitizenRemoveClientPacket :: Giving up citizenship of a town.
type CitizenRemoveClientPacket struct {
	byteSize int

	BehaviorId int
}

func (s CitizenRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CitizenRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	return
}

func (s *CitizenRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CitizenOpenClientPacket :: Talking to a citizenship NPC.
type CitizenOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s CitizenOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Citizen
}

func (s CitizenOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CitizenOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *CitizenOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *CitizenOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopCreateClientPacket :: Crafting an item from a shop.
type ShopCreateClientPacket struct {
	byteSize int

	CraftItemId int
	SessionId   int
}

func (s ShopCreateClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopCreateClientPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopCreateClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ShopCreateClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ShopCreateClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CraftItemId : field : short
	s.CraftItemId = reader.GetShort()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopBuyClientPacket :: Purchasing an item from a shop.
type ShopBuyClientPacket struct {
	byteSize int

	BuyItem   net.Item
	SessionId int
}

func (s ShopBuyClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopBuyClientPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopBuyClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ShopBuyClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ShopBuyClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BuyItem : field : Item
	if err = s.BuyItem.Deserialize(reader); err != nil {
		return
	}
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopSellClientPacket :: Selling an item to a shop.
type ShopSellClientPacket struct {
	byteSize int

	SellItem  net.Item
	SessionId int
}

func (s ShopSellClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopSellClientPacket) Action() net.PacketAction {
	return net.PacketAction_Sell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopSellClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ShopSellClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ShopSellClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SellItem : field : Item
	if err = s.SellItem.Deserialize(reader); err != nil {
		return
	}
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopOpenClientPacket :: Talking to a shop NPC.
type ShopOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s ShopOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Shop
}

func (s ShopOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ShopOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *ShopOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillOpenClientPacket :: Talking to a skill master NPC.
type StatSkillOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s StatSkillOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *StatSkillOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillTakeClientPacket :: Learning a skill from a skill master NPC.
type StatSkillTakeClientPacket struct {
	byteSize int

	SessionId int
	SpellId   int
}

func (s StatSkillTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *StatSkillTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillRemoveClientPacket :: Forgetting a skill at a skill master NPC.
type StatSkillRemoveClientPacket struct {
	byteSize int

	SessionId int
	SpellId   int
}

func (s StatSkillRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *StatSkillRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillAddClientPacket :: Spending a stat point on a stat or skill.
type StatSkillAddClientPacket struct {
	byteSize int

	ActionType     TrainType
	ActionTypeData StatSkillAddActionTypeData
}

type StatSkillAddActionTypeData interface {
	protocol.EoData
}

type StatSkillAddActionTypeDataStat struct {
	byteSize int

	StatId StatId
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillAddActionTypeDataStat) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillAddActionTypeDataStat) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// StatId : field : StatId
	if err = writer.AddShort(int(s.StatId)); err != nil {
		return
	}
	return
}

func (s *StatSkillAddActionTypeDataStat) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// StatId : field : StatId
	s.StatId = StatId(reader.GetShort())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type StatSkillAddActionTypeDataSkill struct {
	byteSize int

	SpellId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillAddActionTypeDataSkill) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillAddActionTypeDataSkill) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SpellId : field : short
	if err = writer.AddShort(s.SpellId); err != nil {
		return
	}
	return
}

func (s *StatSkillAddActionTypeDataSkill) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s StatSkillAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *StatSkillAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// StatSkillJunkClientPacket :: Resetting stats at a skill master.
type StatSkillJunkClientPacket struct {
	byteSize int

	SessionId int
}

func (s StatSkillJunkClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_StatSkill
}

func (s StatSkillJunkClientPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *StatSkillJunkClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *StatSkillJunkClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *StatSkillJunkClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemUseClientPacket :: Using an item.
type ItemUseClientPacket struct {
	byteSize int

	ItemId int
}

func (s ItemUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ItemUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	return
}

func (s *ItemUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemDropClientPacket :: Dropping items on the ground.
type ItemDropClientPacket struct {
	byteSize int

	Item   net.ThreeItem
	Coords ByteCoords //  The official client sends 255 byte values for the coords if an item is dropped via. the GUI button. 255 values here should be interpreted to mean "drop at current coords". Otherwise, the x and y fields contain encoded numbers that must be explicitly. decoded to get the actual x and y values.
}

func (s ItemDropClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemDropClientPacket) Action() net.PacketAction {
	return net.PacketAction_Drop
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemDropClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ItemDropClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ItemDropClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Item : field : ThreeItem
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	// Coords : field : ByteCoords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemJunkClientPacket :: Junking items.
type ItemJunkClientPacket struct {
	byteSize int

	Item net.Item
}

func (s ItemJunkClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemJunkClientPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemJunkClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ItemJunkClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Item : field : Item
	if err = s.Item.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ItemJunkClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Item : field : Item
	if err = s.Item.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemGetClientPacket :: Taking items from the ground.
type ItemGetClientPacket struct {
	byteSize int

	ItemIndex int
}

func (s ItemGetClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Item
}

func (s ItemGetClientPacket) Action() net.PacketAction {
	return net.PacketAction_Get
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemGetClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ItemGetClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemIndex : field : short
	if err = writer.AddShort(s.ItemIndex); err != nil {
		return
	}
	return
}

func (s *ItemGetClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemIndex : field : short
	s.ItemIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardRemoveClientPacket :: Removing a post from a town board.
type BoardRemoveClientPacket struct {
	byteSize int

	BoardId int
	PostId  int
}

func (s BoardRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BoardRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *BoardRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BoardId : field : short
	s.BoardId = reader.GetShort()
	// PostId : field : short
	s.PostId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardCreateClientPacket :: Posting a new message to a town board.
type BoardCreateClientPacket struct {
	byteSize int

	BoardId     int
	PostSubject string
	PostBody    string
}

func (s BoardCreateClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardCreateClientPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardCreateClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BoardCreateClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}
	writer.AddByte(255)
	// PostSubject : field : string
	if err = writer.AddString(s.PostSubject); err != nil {
		return
	}
	writer.AddByte(255)
	// PostBody : field : string
	if err = writer.AddString(s.PostBody); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *BoardCreateClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardTakeClientPacket :: Reading a post on a town board.
type BoardTakeClientPacket struct {
	byteSize int

	BoardId int
	PostId  int
}

func (s BoardTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BoardTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *BoardTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BoardId : field : short
	s.BoardId = reader.GetShort()
	// PostId : field : short
	s.PostId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardOpenClientPacket :: Opening a town board.
type BoardOpenClientPacket struct {
	byteSize int

	BoardId int
}

func (s BoardOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Board
}

func (s BoardOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BoardOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BoardId : field : short
	if err = writer.AddShort(s.BoardId); err != nil {
		return
	}
	return
}

func (s *BoardOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BoardId : field : short
	s.BoardId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxOpenClientPacket :: Opening the jukebox listing.
type JukeboxOpenClientPacket struct {
	byteSize int

	Coords protocol.Coords
}

func (s JukeboxOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *JukeboxOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *JukeboxOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxMsgClientPacket :: Requesting a song on a jukebox.
type JukeboxMsgClientPacket struct {
	byteSize int

	TrackId int // This value is 0-indexed.
}

func (s JukeboxMsgClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxMsgClientPacket) Action() net.PacketAction {
	return net.PacketAction_Msg
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxMsgClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *JukeboxMsgClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	// TrackId : field : short
	if err = writer.AddShort(s.TrackId); err != nil {
		return
	}
	return
}

func (s *JukeboxMsgClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 0 : field : char
	reader.GetChar()
	// 0 : field : char
	reader.GetChar()
	// TrackId : field : short
	s.TrackId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// JukeboxUseClientPacket :: Playing a note with the bard skill.
type JukeboxUseClientPacket struct {
	byteSize int

	InstrumentId int
	NoteId       int
}

func (s JukeboxUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Jukebox
}

func (s JukeboxUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *JukeboxUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *JukeboxUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *JukeboxUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// InstrumentId : field : char
	s.InstrumentId = reader.GetChar()
	// NoteId : field : char
	s.NoteId = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpAcceptClientPacket :: Accept a warp request from the server.
type WarpAcceptClientPacket struct {
	byteSize int

	MapId     int
	SessionId int
}

func (s WarpAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WarpAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *WarpAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WarpTakeClientPacket :: Request to download a copy of the map.
type WarpTakeClientPacket struct {
	byteSize int

	MapId     int
	SessionId int
}

func (s WarpTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Warp
}

func (s WarpTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WarpTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *WarpTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *WarpTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollRequestClientPacket :: Request for a player's paperdoll.
type PaperdollRequestClientPacket struct {
	byteSize int

	PlayerId int
}

func (s PaperdollRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PaperdollRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *PaperdollRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollRemoveClientPacket :: Unequipping an item.
type PaperdollRemoveClientPacket struct {
	byteSize int

	ItemId int
	SubLoc int
}

func (s PaperdollRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PaperdollRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *PaperdollRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PaperdollAddClientPacket :: Equipping an item.
type PaperdollAddClientPacket struct {
	byteSize int

	ItemId int
	SubLoc int
}

func (s PaperdollAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Paperdoll
}

func (s PaperdollAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PaperdollAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PaperdollAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *PaperdollAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SubLoc : field : char
	s.SubLoc = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BookRequestClientPacket :: Request for a player's book.
type BookRequestClientPacket struct {
	byteSize int

	PlayerId int
}

func (s BookRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Book
}

func (s BookRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BookRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *BookRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *BookRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MessagePingClientPacket :: #ping command request.
type MessagePingClientPacket struct {
	byteSize int
}

func (s MessagePingClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Message
}

func (s MessagePingClientPacket) Action() net.PacketAction {
	return net.PacketAction_Ping
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MessagePingClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *MessagePingClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 2 : dummy : short
	if err = writer.AddShort(2); err != nil {
		return
	}
	return
}

func (s *MessagePingClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 2 : dummy : short
	reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersAcceptClientPacket :: #find command request.
type PlayersAcceptClientPacket struct {
	byteSize int

	Name string
}

func (s PlayersAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PlayersAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *PlayersAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// PlayersRequestClientPacket :: Requesting a list of online players.
type PlayersRequestClientPacket struct {
	byteSize int
}

func (s PlayersRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PlayersRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *PlayersRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersListClientPacket :: Requesting a list of online friends.
type PlayersListClientPacket struct {
	byteSize int
}

func (s PlayersListClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Players
}

func (s PlayersListClientPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersListClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PlayersListClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *PlayersListClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// DoorOpenClientPacket :: Opening a door.
type DoorOpenClientPacket struct {
	byteSize int

	Coords protocol.Coords
}

func (s DoorOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Door
}

func (s DoorOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *DoorOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *DoorOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *DoorOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestOpenClientPacket :: Opening a chest.
type ChestOpenClientPacket struct {
	byteSize int

	Coords protocol.Coords
}

func (s ChestOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ChestOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChestOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestAddClientPacket :: Placing an item in to a chest.
type ChestAddClientPacket struct {
	byteSize int

	Coords  protocol.Coords
	AddItem net.ThreeItem
}

func (s ChestAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ChestAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ChestAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// AddItem : field : ThreeItem
	if err = s.AddItem.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ChestTakeClientPacket :: Taking an item from a chest.
type ChestTakeClientPacket struct {
	byteSize int

	Coords     protocol.Coords
	TakeItemId int
}

func (s ChestTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Chest
}

func (s ChestTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChestTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *ChestTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ChestTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// TakeItemId : field : short
	s.TakeItemId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RefreshRequestClientPacket :: Requesting new info about nearby objects.
type RefreshRequestClientPacket struct {
	byteSize int
}

func (s RefreshRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Refresh
}

func (s RefreshRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RefreshRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *RefreshRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 255 : dummy : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	return
}

func (s *RefreshRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 255 : dummy : byte
	reader.GetByte()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// RangeRequestClientPacket :: Requesting info about nearby players and NPCs.
type RangeRequestClientPacket struct {
	byteSize int

	PlayerIds  []int
	NpcIndexes []int
}

func (s RangeRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Range
}

func (s RangeRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *RangeRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *RangeRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

	writer.SanitizeStrings = false
	return
}

func (s *RangeRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayerRangeRequestClientPacket :: Requesting info about nearby players.
type PlayerRangeRequestClientPacket struct {
	byteSize int

	PlayerIds []int
}

func (s PlayerRangeRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_PlayerRange
}

func (s PlayerRangeRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayerRangeRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PlayerRangeRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *PlayerRangeRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerIds : array : short
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.PlayerIds = append(s.PlayerIds, 0)
		s.PlayerIds[ndx] = reader.GetShort()
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcRangeRequestClientPacket :: Requesting info about nearby NPCs.
type NpcRangeRequestClientPacket struct {
	byteSize int

	NpcIndexes []int
}

func (s NpcRangeRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_NpcRange
}

func (s NpcRangeRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcRangeRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *NpcRangeRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndexesLength : length : char
	if err = writer.AddChar(len(s.NpcIndexes)); err != nil {
		return
	}
	// 255 : field : byte
	if err = writer.AddByte(255); err != nil {
		return
	}
	// NpcIndexes : array : char
	for ndx := 0; ndx < len(s.NpcIndexes); ndx++ {
		if err = writer.AddChar(s.NpcIndexes[ndx]); err != nil {
			return
		}
	}

	return
}

func (s *NpcRangeRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndexesLength : length : char
	npcIndexesLength := reader.GetChar()
	// 255 : field : byte
	reader.GetByte()
	// NpcIndexes : array : char
	for ndx := 0; ndx < npcIndexesLength; ndx++ {
		s.NpcIndexes = append(s.NpcIndexes, 0)
		s.NpcIndexes[ndx] = reader.GetChar()
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyRequestClientPacket :: Send party invite / join request.
type PartyRequestClientPacket struct {
	byteSize int

	RequestType net.PartyRequestType
	PlayerId    int
}

func (s PartyRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PartyRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *PartyRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyAcceptClientPacket :: Accept party invite / join request.
type PartyAcceptClientPacket struct {
	byteSize int

	RequestType     net.PartyRequestType
	InviterPlayerId int
}

func (s PartyAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PartyAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *PartyAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// RequestType : field : PartyRequestType
	s.RequestType = net.PartyRequestType(reader.GetChar())
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyRemoveClientPacket :: Remove player from a party.
type PartyRemoveClientPacket struct {
	byteSize int

	PlayerId int
}

func (s PartyRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PartyRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *PartyRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyTakeClientPacket :: Request updated party info.
type PartyTakeClientPacket struct {
	byteSize int

	MembersCount int
}

func (s PartyTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Party
}

func (s PartyTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PartyTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MembersCount : field : char
	if err = writer.AddChar(s.MembersCount); err != nil {
		return
	}
	return
}

func (s *PartyTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// MembersCount : field : char
	s.MembersCount = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildRequestClientPacket :: Requested to create a guild.
type GuildRequestClientPacket struct {
	byteSize int

	SessionId int
	GuildTag  string
	GuildName string
}

func (s GuildRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
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
	writer.SanitizeStrings = false
	return
}

func (s *GuildRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildAcceptClientPacket :: Accept pending guild creation invite.
type GuildAcceptClientPacket struct {
	byteSize int

	InviterPlayerId int
}

func (s GuildAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 20202 : field : int
	if err = writer.AddInt(20202); err != nil {
		return
	}
	// InviterPlayerId : field : short
	if err = writer.AddShort(s.InviterPlayerId); err != nil {
		return
	}
	return
}

func (s *GuildAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 20202 : field : int
	reader.GetInt()
	// InviterPlayerId : field : short
	s.InviterPlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildRemoveClientPacket :: Leave guild.
type GuildRemoveClientPacket struct {
	byteSize int

	SessionId int
}

func (s GuildRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *GuildRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildAgreeClientPacket :: Update the guild description or rank list.
type GuildAgreeClientPacket struct {
	byteSize int

	SessionId    int
	InfoType     GuildInfoType
	InfoTypeData GuildAgreeInfoTypeData
}

type GuildAgreeInfoTypeData interface {
	protocol.EoData
}

type GuildAgreeInfoTypeDataDescription struct {
	byteSize int

	Description string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAgreeInfoTypeDataDescription) ByteSize() int {
	return s.byteSize
}

func (s *GuildAgreeInfoTypeDataDescription) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	return
}

func (s *GuildAgreeInfoTypeDataDescription) Deserialize(reader *data.EoReader) (err error) {
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

type GuildAgreeInfoTypeDataRanks struct {
	byteSize int

	Ranks []string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAgreeInfoTypeDataRanks) ByteSize() int {
	return s.byteSize
}

func (s *GuildAgreeInfoTypeDataRanks) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Ranks : array : string
	for ndx := 0; ndx < 9; ndx++ {
		if len(s.Ranks) != 9 {
			err = fmt.Errorf("expected Ranks with length 9, got %d", len(s.Ranks))
			return
		}

		if err = writer.AddString(s.Ranks[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	return
}

func (s *GuildAgreeInfoTypeDataRanks) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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

	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s GuildAgreeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildAgreeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildAgreeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildAgreeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildAgreeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildCreateClientPacket :: Final confirm creating a guild.
type GuildCreateClientPacket struct {
	byteSize int

	SessionId   int
	GuildTag    string
	GuildName   string
	Description string
}

func (s GuildCreateClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildCreateClientPacket) Action() net.PacketAction {
	return net.PacketAction_Create
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildCreateClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildCreateClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
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
	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *GuildCreateClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildPlayerClientPacket :: Request to join a guild.
type GuildPlayerClientPacket struct {
	byteSize int

	SessionId     int
	GuildTag      string
	RecruiterName string
}

func (s GuildPlayerClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildPlayerClientPacket) Action() net.PacketAction {
	return net.PacketAction_Player
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildPlayerClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildPlayerClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.AddByte(255)
	// RecruiterName : field : string
	if err = writer.AddString(s.RecruiterName); err != nil {
		return
	}
	writer.AddByte(255)
	writer.SanitizeStrings = false
	return
}

func (s *GuildPlayerClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildTakeClientPacket :: Request guild description, rank list, or bank balance.
type GuildTakeClientPacket struct {
	byteSize int

	SessionId int
	InfoType  GuildInfoType
	GuildTag  string
}

func (s GuildTakeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTakeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Take
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildTakeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildTakeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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
	// GuildTag : field : string
	if len(s.GuildTag) != 3 {
		err = fmt.Errorf("expected GuildTag with length 3, got %d", len(s.GuildTag))
		return
	}
	if err = writer.AddFixedString(s.GuildTag, 3); err != nil {
		return
	}
	return
}

func (s *GuildTakeClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// InfoType : field : GuildInfoType
	s.InfoType = GuildInfoType(reader.GetShort())
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetFixedString(3); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildUseClientPacket :: Accepted a join request.
type GuildUseClientPacket struct {
	byteSize int

	PlayerId int
}

func (s GuildUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *GuildUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildBuyClientPacket :: Deposit gold in to the guild bank.
type GuildBuyClientPacket struct {
	byteSize int

	SessionId  int
	GoldAmount int
}

func (s GuildBuyClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildBuyClientPacket) Action() net.PacketAction {
	return net.PacketAction_Buy
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildBuyClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildBuyClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildBuyClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GoldAmount : field : int
	s.GoldAmount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildOpenClientPacket :: Talking to a guild master NPC.
type GuildOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s GuildOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *GuildOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildTellClientPacket :: Requested member list of a guild.
type GuildTellClientPacket struct {
	byteSize int

	SessionId     int
	GuildIdentity string
}

func (s GuildTellClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildTellClientPacket) Action() net.PacketAction {
	return net.PacketAction_Tell
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildTellClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildTellClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildTellClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildReportClientPacket :: Requested general information of a guild.
type GuildReportClientPacket struct {
	byteSize int

	SessionId     int
	GuildIdentity string
}

func (s GuildReportClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildReportClientPacket) Action() net.PacketAction {
	return net.PacketAction_Report
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildReportClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildReportClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildReportClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// GuildIdentity : field : string
	if s.GuildIdentity, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildJunkClientPacket :: Disband guild.
type GuildJunkClientPacket struct {
	byteSize int

	SessionId int
}

func (s GuildJunkClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildJunkClientPacket) Action() net.PacketAction {
	return net.PacketAction_Junk
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildJunkClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildJunkClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *GuildJunkClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildKickClientPacket :: Kick member from guild.
type GuildKickClientPacket struct {
	byteSize int

	SessionId  int
	MemberName string
}

func (s GuildKickClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildKickClientPacket) Action() net.PacketAction {
	return net.PacketAction_Kick
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildKickClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildKickClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildKickClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// MemberName : field : string
	if s.MemberName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildRankClientPacket :: Update a member's rank.
type GuildRankClientPacket struct {
	byteSize int

	SessionId  int
	Rank       int
	MemberName string
}

func (s GuildRankClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Guild
}

func (s GuildRankClientPacket) Action() net.PacketAction {
	return net.PacketAction_Rank
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildRankClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *GuildRankClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *GuildRankClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	// Rank : field : char
	s.Rank = reader.GetChar()
	// MemberName : field : string
	if s.MemberName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellRequestClientPacket :: Begin spell chanting.
type SpellRequestClientPacket struct {
	byteSize int

	SpellId   int
	Timestamp int
}

func (s SpellRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *SpellRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetSelfClientPacket :: Self-targeted spell cast.
type SpellTargetSelfClientPacket struct {
	byteSize int

	Direction protocol.Direction
	SpellId   int
	Timestamp int
}

func (s SpellTargetSelfClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetSelfClientPacket) Action() net.PacketAction {
	return net.PacketAction_TargetSelf
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetSelfClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellTargetSelfClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *SpellTargetSelfClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetOtherClientPacket :: Targeted spell cast.
type SpellTargetOtherClientPacket struct {
	byteSize int

	TargetType        SpellTargetType
	PreviousTimestamp int
	SpellId           int
	VictimId          int
	Timestamp         int
}

func (s SpellTargetOtherClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetOtherClientPacket) Action() net.PacketAction {
	return net.PacketAction_TargetOther
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetOtherClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellTargetOtherClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *SpellTargetOtherClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellTargetGroupClientPacket :: Group spell cast.
type SpellTargetGroupClientPacket struct {
	byteSize int

	SpellId   int
	Timestamp int
}

func (s SpellTargetGroupClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellTargetGroupClientPacket) Action() net.PacketAction {
	return net.PacketAction_TargetGroup
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellTargetGroupClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellTargetGroupClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *SpellTargetGroupClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SpellId : field : short
	s.SpellId = reader.GetShort()
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SpellUseClientPacket :: Raise arm to cast a spell (vestigial).
type SpellUseClientPacket struct {
	byteSize int

	Direction protocol.Direction
}

func (s SpellUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Spell
}

func (s SpellUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SpellUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *SpellUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *SpellUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeRequestClientPacket :: Requesting a trade with another player.
type TradeRequestClientPacket struct {
	byteSize int

	PlayerId int
}

func (s TradeRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 138 : field : char
	if err = writer.AddChar(138); err != nil {
		return
	}
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *TradeRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 138 : field : char
	reader.GetChar()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeAcceptClientPacket :: Accepting a trade request.
type TradeAcceptClientPacket struct {
	byteSize int

	PlayerId int
}

func (s TradeAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	return
}

func (s *TradeAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 0 : field : char
	reader.GetChar()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeRemoveClientPacket :: Remove an item from the trade screen.
type TradeRemoveClientPacket struct {
	byteSize int

	ItemId int
}

func (s TradeRemoveClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeRemoveClientPacket) Action() net.PacketAction {
	return net.PacketAction_Remove
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeRemoveClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeRemoveClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	return
}

func (s *TradeRemoveClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeAgreeClientPacket :: Mark trade as agreed.
type TradeAgreeClientPacket struct {
	byteSize int

	Agree bool
}

func (s TradeAgreeClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAgreeClientPacket) Action() net.PacketAction {
	return net.PacketAction_Agree
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeAgreeClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeAgreeClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *TradeAgreeClientPacket) Deserialize(reader *data.EoReader) (err error) {
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

// TradeAddClientPacket :: Add an item to the trade screen.
type TradeAddClientPacket struct {
	byteSize int

	AddItem net.Item
}

func (s TradeAddClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeAddClientPacket) Action() net.PacketAction {
	return net.PacketAction_Add
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeAddClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeAddClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// AddItem : field : Item
	if err = s.AddItem.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *TradeAddClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// AddItem : field : Item
	if err = s.AddItem.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeCloseClientPacket :: Cancel the trade.
type TradeCloseClientPacket struct {
	byteSize int
}

func (s TradeCloseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Trade
}

func (s TradeCloseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Close
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeCloseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *TradeCloseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// 0 : dummy : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	return
}

func (s *TradeCloseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// 0 : dummy : char
	reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestUseClientPacket :: Talking to a quest NPC.
type QuestUseClientPacket struct {
	byteSize int

	NpcIndex int
	QuestId  int //  Quest ID is 0 unless the player explicitly selects a quest from the quest switcher.
}

func (s QuestUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *QuestUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *QuestUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// QuestId : field : short
	s.QuestId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestAcceptClientPacket :: Response to a quest NPC dialog.
type QuestAcceptClientPacket struct {
	byteSize int

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
	byteSize int

	Action int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestAcceptReplyTypeDataLink) ByteSize() int {
	return s.byteSize
}

func (s *QuestAcceptReplyTypeDataLink) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Action : field : char
	if err = writer.AddChar(s.Action); err != nil {
		return
	}
	return
}

func (s *QuestAcceptReplyTypeDataLink) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Action : field : char
	s.Action = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

func (s QuestAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *QuestAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
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

func (s *QuestAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestListClientPacket :: Quest history / progress request.
type QuestListClientPacket struct {
	byteSize int

	Page net.QuestPage
}

func (s QuestListClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Quest
}

func (s QuestListClientPacket) Action() net.PacketAction {
	return net.PacketAction_List
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestListClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *QuestListClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Page : field : QuestPage
	if err = writer.AddChar(int(s.Page)); err != nil {
		return
	}
	return
}

func (s *QuestListClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Page : field : QuestPage
	s.Page = net.QuestPage(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MarriageOpenClientPacket :: Talking to a law NPC.
type MarriageOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s MarriageOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MarriageOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *MarriageOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *MarriageOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MarriageRequestClientPacket :: Requesting marriage approval.
type MarriageRequestClientPacket struct {
	byteSize int

	RequestType MarriageRequestType
	SessionId   int
	Name        string
}

func (s MarriageRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Marriage
}

func (s MarriageRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MarriageRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *MarriageRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
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
	writer.AddByte(255)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *MarriageRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
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

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestAcceptClientPacket :: Accepting a marriage request.
type PriestAcceptClientPacket struct {
	byteSize int

	SessionId int
}

func (s PriestAcceptClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestAcceptClientPacket) Action() net.PacketAction {
	return net.PacketAction_Accept
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestAcceptClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PriestAcceptClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : short
	if err = writer.AddShort(s.SessionId); err != nil {
		return
	}
	return
}

func (s *PriestAcceptClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : short
	s.SessionId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestOpenClientPacket :: Talking to a priest NPC.
type PriestOpenClientPacket struct {
	byteSize int

	NpcIndex int
}

func (s PriestOpenClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestOpenClientPacket) Action() net.PacketAction {
	return net.PacketAction_Open
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestOpenClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PriestOpenClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : int
	if err = writer.AddInt(s.NpcIndex); err != nil {
		return
	}
	return
}

func (s *PriestOpenClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : int
	s.NpcIndex = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestRequestClientPacket :: Requesting marriage at a priest.
type PriestRequestClientPacket struct {
	byteSize int

	SessionId int
	Name      string
}

func (s PriestRequestClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestRequestClientPacket) Action() net.PacketAction {
	return net.PacketAction_Request
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestRequestClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PriestRequestClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	writer.AddByte(255)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *PriestRequestClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PriestUseClientPacket :: Saying "I do" at a wedding.
type PriestUseClientPacket struct {
	byteSize int

	SessionId int
}

func (s PriestUseClientPacket) Family() net.PacketFamily {
	return net.PacketFamily_Priest
}

func (s PriestUseClientPacket) Action() net.PacketAction {
	return net.PacketAction_Use
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PriestUseClientPacket) ByteSize() int {
	return s.byteSize
}

func (s *PriestUseClientPacket) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SessionId : field : int
	if err = writer.AddInt(s.SessionId); err != nil {
		return
	}
	return
}

func (s *PriestUseClientPacket) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// SessionId : field : int
	s.SessionId = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

package mtproto

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func (m *MTProto) ContactsGetContacts(hash string) (*TL, error) {
	return m.InvokeSync(TL_contacts_getContacts{
		Hash: hash,
	})
}

func (m *MTProto) ContactsGetTopPeers(correspondents, botsPM, botsInline, groups, channels bool, offset, limit, hash int32) (*TL, error) {
	tl, err := m.InvokeSync(TL_contacts_getTopPeers{
		Correspondents: correspondents,
		Bots_pm:        botsPM,
		Bots_inline:    botsInline,
		Groups:         groups,
		Channels:       channels,
		Offset:         offset,
		Limit:          limit,
		Hash:           hash,
	})

	if err != nil {
		return nil, err
	}

	switch (*tl).(type) {
	case TL_contacts_topPeersNotModified:
	case TL_contacts_topPeers:
	default:
		return nil, errors.New("MTProto::ContactsGetTopPeers error: Unknown type")
	}

	return tl, nil
}

func (m *MTProto) ContactsSearch(q string, limit int) (*TL_contacts_found, error) {
	var res TL_contacts_found
	tl, err := m.InvokeSync(TL_contacts_search{
		Q:     q,
		Limit: int32(limit),
	})

	if err != nil {
		return nil, err
	}

	switch (*tl).(type) {
	case TL_contacts_found:
		res = (*tl).(TL_contacts_found)
		return &res, nil
	default:
		return nil, errors.New("MTProto::ContactsGetTopPeers error: Unknown type")
	}
	return nil, errors.New("notmytype")
}

func (m *MTProto) ImportContacts(larens []*TL_inputPhoneContact) {
	_contacts := []TL{}
	for idx := range larens {
		_contacts = append(_contacts, *(larens[idx]))
	}
	m.InvokeSync(TL_contacts_importContacts{
		Contacts: _contacts,
		Replace:  TL_boolTrue{},
	})
}

func (m *MTProto) DeleteContact(inputUser *TL_inputUser) {
	time.Sleep(time.Second * 1)
	m.InvokeSync(TL_contacts_deleteContact{
		Id: *inputUser,
	})
}
func (m *MTProto) DeleteContactList(inputUser []*TL_inputUser) {
	tl := []TL{}
	for idx := range inputUser {
		tl = append(tl, *(inputUser[idx]))
	}
	m.InvokeSync(TL_contacts_deleteContacts{
		Id: tl,
	})
}

func (m *MTProto) InviteToChannel(inputUsers []TL, channel TL_inputChannel) {
	time.Sleep(time.Second * 1) // make sure

	// invite to channel
	tlmain := TL_channels_inviteToChannel{
		Channel: channel,
		Users:   inputUsers,
	}
	_, err := m.InvokeSync(tlmain)
	if err != nil {
		fmt.Printf("err->%v\n", err)
		if strings.Contains(err.Error(), "CHAT_WRITE_FORBIDDEN") || strings.Contains(err.Error(), "PEER_FLOOD") || strings.Contains(err.Error(), "USER_BANNED_IN_CHANNEL") {
			fmt.Printf("willexit->%v\n", err)
			os.Exit(1)
		}
	}
}

func (m *MTProto) JoinChannel(channel TL_inputChannel) {
	time.Sleep(time.Second * 1) // make sure

	// invite to channel
	tlmain := TL_channels_joinChannel{
		Channel: channel,
	}
	_, err := m.InvokeSync(tlmain)
	if err != nil {
		fmt.Printf("err->%v\n", err)
		if strings.Contains(err.Error(), "CHAT_WRITE_FORBIDDEN") || strings.Contains(err.Error(), "PEER_FLOOD") {
			fmt.Printf("willexit->%v\n", err)
			os.Exit(1)
		}
	}
}
func (m *MTProto) ResolveName(run TL_contacts_resolveUsername) int64 {
	time.Sleep(time.Second * 10) // make sure

	// invite to channel
	tlmain := run
	tl, err := m.InvokeSync(tlmain)
	if err != nil {
		fmt.Printf("err->%v\n", err)
		if strings.Contains(err.Error(), "CHAT_WRITE_FORBIDDEN") || strings.Contains(err.Error(), "PEER_FLOOD") {
			fmt.Printf("willexit->%v\n", err)
			os.Exit(1)
		}
	}
	peer := (*tl).(TL_contacts_resolvedPeer)
	for idx := range peer.Chats {
		if tl_channel, ok := peer.Chats[idx].(TL_channel); ok {
			return tl_channel.Access_hash
		}
	}
	return 88
}

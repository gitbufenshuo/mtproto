package mtproto

import "errors"

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
	})
}

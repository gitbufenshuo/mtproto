package mtproto

import (
	"errors"
	"fmt"
)

func (m *MTProto) AuthSendCode(phonenumber string) (*TL_auth_sentCode, error) {
	var authSentCode TL_auth_sentCode
	tl, err := m.InvokeSync(TL_auth_sendCode{
		Allow_flashcall: false,
		Phone_number:    phonenumber,
		Current_number:  TL_boolTrue{},
		Api_id:          m.id,
		Api_hash:        m.hash,
	})
	fmt.Println("m.InvokeSync_over", err)
	if err != nil {
		return nil, err
	}

	switch (*tl).(type) {
	case TL_auth_sentCode:
		authSentCode = (*tl).(TL_auth_sentCode)
	default:
		return nil, fmt.Errorf("Got: %T", *tl)
	}

	return &authSentCode, nil
}

func (m *MTProto) AuthSignIn(phoneNumber, phoneCode, phoneCodeHash string) (*TL_auth_authorization, error) {
	if phoneNumber == "" || phoneCode == "" || phoneCodeHash == "" {
		return nil, errors.New("MRProto::AuthSignIn one of function parameters is empty")
	}

	tl, err := m.InvokeSync(TL_auth_signIn{
		Phone_number:    phoneNumber,
		Phone_code_hash: phoneCodeHash,
		Phone_code:      phoneCode,
	})

	if err != nil {
		return nil, err
	}

	auth, ok := (*tl).(TL_auth_authorization)

	if !ok {
		return nil, fmt.Errorf("RPC: %#v", *tl)
	}

	return &auth, nil
}

// Phone_number:    m.String(),
// Phone_code_hash: m.String(),
// Phone_code:      m.String(),
// First_name:      m.String(),
// Last_name:       m.String(),
func (m *MTProto) AuthSignUp(phoneNumber, phoneCode, phoneCodeHash, fname, lname string) (*TL_auth_authorization, error) {
	if phoneNumber == "" || phoneCode == "" || phoneCodeHash == "" {
		return nil, errors.New("MRProto::AuthSignUp one of function parameters is empty")
	}

	tl, err := m.InvokeSync(TL_auth_signUp{
		Phone_number:    phoneNumber,
		Phone_code_hash: phoneCodeHash,
		Phone_code:      phoneCode,
		First_name:      fname,
		Last_name:       lname,
	})

	if err != nil {
		return nil, err
	}

	auth, ok := (*tl).(TL_auth_authorization)

	if !ok {
		return nil, fmt.Errorf("RPC: %#v", *tl)
	}

	return &auth, nil
}

func (m *MTProto) AuthLogOut() (bool, error) {
	var result bool

	tl, err := m.InvokeSync(TL_auth_logOut{})
	if err != nil {
		return result, err
	}

	result, err = ToBool(*tl)
	if err != nil {
		return result, err
	}

	return result, err
}

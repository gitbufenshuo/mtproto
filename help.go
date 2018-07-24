package mtproto

import "fmt"

func (m *MTProto) HelpGetConfig() (*TL_config, error) {
	var config TL_config
	tl, err := m.InvokeSync(TL_auth_sendCode{})
	fmt.Println("m.InvokeSync_over", err)
	if err != nil {
		return nil, err
	}

	switch (*tl).(type) {
	case TL_config:
		config = (*tl).(TL_config)
	default:
		return nil, fmt.Errorf("Got: %T", *tl)
	}

	return &config, nil
}

package main

type OutlineKey struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Port      int64  `json:"port"`
	Method    string `json:"method"`
	AccessURL string `json:"accessUrl"`
}

type BytesTransferred struct {
	BytesTransferredByUserId map[string]int64 `json:"bytesTransferredByUserId"`
}

func NewKey() *OutlineKey {
	return &OutlineKey{}
}

func (k *OutlineKey) SetID(id string) *OutlineKey {
	k.ID = id
	return k
}

func (k *OutlineKey) SetName(name string) *OutlineKey {
	k.Name = name
	return k
}

func (k *OutlineKey) SetPassword(password string) *OutlineKey {
	k.Password = password
	return k
}

func (k *OutlineKey) SetPort(port int64) *OutlineKey {
	k.Port = port
	return k
}

func (k *OutlineKey) SetMethod(method string) *OutlineKey {
	k.Method = method
	return k
}

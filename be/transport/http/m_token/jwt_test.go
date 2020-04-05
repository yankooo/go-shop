package m_token

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(0, "")
	fmt.Printf("token: %+v, err %+v", token, err)

	fmt.Printf("res: %+v", ParseToken(token))
}

package kvalid

import "github.com/google/uuid"

// IsValidUUID 验证是否为有效的 UUID
func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

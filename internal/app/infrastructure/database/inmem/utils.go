package inmem

import (
	"time"

	"github.com/xabsvoid/eis/internal/app/domain/model/valueobject"
)

func refresh(md valueobject.Metadata) valueobject.Metadata {
	md.UpdatedAt = time.Now().UTC()
	if md.CreatedAt.IsZero() {
		md.CreatedAt = md.UpdatedAt
	}
	return md
}

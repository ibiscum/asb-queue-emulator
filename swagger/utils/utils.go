package utils

import (
	"asb-queue-emulator/pkg/broker/abstract"
	"time"

	"github.com/google/uuid"
)

type HandlerContext struct {
	MQBroker abstract.MQBroker
}

type BrokerProperties struct {
	ContentType             string
	CorrelationId           string
	DeadLetterSource        string
	DeliveryCount           int
	EnqueuedSequenceNumber  int64
	EnqueuedTimeUtc         time.Time
	ExpiresAtUtc            time.Time
	IsBodyConsumed          bool
	LockedUntilUtc          time.Time
	ForcePersistence        bool
	Label                   string
	MessageId               string
	LockToken               uuid.UUID
	PartitionKey            string
	ReplyTo                 string
	ReplyToSessionId        string
	ScheduledEnqueueTimeUtc time.Time
	SequenceNumber          int64
	SessionId               string
	Size                    int64
	// enum: [Active, Deferred, Scheduled]
	State           string
	TimeToLive      string
	To              string
	ViaPartitionKey string
}

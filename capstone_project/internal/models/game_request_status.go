package models

type GameRequestStatus string

const (
    Pending  GameRequestStatus = "pending"
    Accepted GameRequestStatus = "accepted"
    Rejected GameRequestStatus = "rejected"
)

package services

import (
	"NetGuardServer/config"
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
	"github.com/google/uuid"
)

// NotificationService defines the interface for notification business logic
type NotificationService interface {
	SendServerDownNotification(serverID uuid.UUID, serverName, serverURL string, reportedBy uuid.UUID) error
}

// notificationService implements NotificationService
type notificationService struct {
	fcmClient *fcm.Client
}

// NewNotificationService creates a new notification service instance
func NewNotificationService() NotificationService {
	ctx := context.Background()
	client, err := fcm.NewClient(
		ctx,
		fcm.WithCredentialsFile(config.AppConfig.FirebaseServiceAccountPath),
	)
	if err != nil {
		log.Printf("Failed to create FCM client: %v", err)
		return &notificationService{fcmClient: nil}
	}

	return &notificationService{
		fcmClient: client,
	}
}

// SendServerDownNotification sends FCM notification when server is down
func (s *notificationService) SendServerDownNotification(serverID uuid.UUID, serverName, serverURL string, reportedBy uuid.UUID) error {
	if s.fcmClient == nil {
		return fmt.Errorf("FCM client not initialized")
	}

	ctx := context.Background()

	// Send to topic "serverdown" - all users subscribed to this topic will receive the notification
	message := &messaging.Message{
		Data: map[string]string{
			"title":       fmt.Sprintf("Server DOWN: %s", serverName),
			"body":        serverURL,
			"server_id":   serverID.String(),
			"server_name": serverName,
			"server_url":  serverURL,
			"status":      "DOWN",
			"reported_by": reportedBy.String(),
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:     fmt.Sprintf("Server DOWN: %s", serverName),
				Body:      serverURL,
				ChannelID: "server_status",
				Priority:  messaging.AndroidNotificationPriority(messaging.PriorityHigh),
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10", // silent notification
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					ContentAvailable: true, // required for silent push
				},
			},
		},
		Topic: "serverdown",
	}

	// Send the message
	resp, err := s.fcmClient.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send FCM message: %w", err)
	}

	log.Printf("FCM notification sent successfully. Success: %d, Failure: %d",
		resp.SuccessCount, resp.FailureCount)

	if resp.FailureCount > 0 {
		return fmt.Errorf("FCM notification partially failed: %d success, %d failure", resp.SuccessCount, resp.FailureCount)
	}

	return nil
}

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"slentry/internal/http/middlewares"
	"slentry/pkg/env"
	"slentry/pkg/sentry"
	"slentry/pkg/slack"
)

var Server *fiber.App

func Start() {
	webhook := env.SlackWebhook.MustGet()
	endpointSecret := env.EndpointSecret.MustGet()

	Server = fiber.New(fiber.Config{
		AppName:                 "Slentry",
		Immutable:               true,
		EnableTrustedProxyCheck: true,
		ServerHeader:            "Slentry",
	})

	Server.Use(
		middlewares.Log,
	)

	Server.Post("/webhooks/sentry/"+endpointSecret, func(ctx *fiber.Ctx) error {
		var payload sentry.Payload
		if err := ctx.BodyParser(&payload); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(Error{Message: "Failed to parse payload."})
		}

		culprit := payload.Culprit
		if culprit == "" {
			culprit = "Unknown"
		}

		message := payload.Message
		if message == "" {
			message = "No message provided by Sentry."
		}

		webhookPayload := slack.Payload{
			Blocks: []slack.Block{
				{
					Type: "section",
					Text: &slack.Text{
						Type: "mrkdwn",
						Text: "*<" + payload.URL + "|" + payload.Event.Title + ">*",
					},
				},
				{
					Type: "context",
					Elements: []slack.Element{
						{
							Type: "plain_text",
							Text: "Culprit: " + culprit,
						},
					},
				},
				{
					Type: "section",
					Text: &slack.Text{
						Type: "mrkdwn",
						Text: "```\n" + message + "\n```",
					},
				},
				{
					Type: "context",
					Elements: []slack.Element{
						{
							Type: "plain_text",
							Text: "Project: " + payload.ProjectName,
						},
						{
							Type: "plain_text",
							Text: "Level: " + payload.Level,
						},
					},
				},
			},
		}
		err := slack.SendWebhook(webhook, &webhookPayload)

		if err != nil {
			log.Err(err).Msg("Failed to send Slack entry.")
			return err
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	})

	_ = Server.Listen(":9950") // k:float
}

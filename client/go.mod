module github.com/supperdoggy/diploma_university_statistics_tool/client

go 1.20

replace github.com/supperdoggy/diploma_university_statistics_tool/models => ../models

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/supperdoggy/diploma_university_statistics_tool/models v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.24.0
	gopkg.in/tucnak/telebot.v2 v2.5.0
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

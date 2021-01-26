package notificationtemplate

type Helper func(from string) (to string)

type NamedHelper struct {
	Name   string
	Helper Helper
}
type Formatter func(format string, raw string) (formatted string)

type NamedFormatter struct {
	Name      string
	Formatter Formatter
}

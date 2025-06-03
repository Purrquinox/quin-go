package types

// AI Data
type AIData struct {
	Name       string     `yaml:"name" validate:"required"`
	Version    int        `yaml:"version" validate:"required"`
	ApiPort    int        `yaml:"api_port" validate:"required"`
	Staff      []Staff    `yaml:"staff" validate:"required"`
	Identity   Identity   `yaml:"identity" validate:"required"`
	Summary    string     `yaml:"summary" validate:"required"`
	CoreTraits CoreTraits `yaml:"coreTraits" validate:"required"`
	Role       Role       `yaml:"role" validate:"required"`
	Design     Design     `yaml:"design" validate:"required"`
	Voice      Voice      `yaml:"voice" validate:"required"`
	Memory     Memory     `yaml:"memory" validate:"required"`
	Friends    []Friend   `yaml:"friends" validate:"required"`
}

type Staff struct {
	Name      string   `yaml:"name" validate:"required"`
	Username  string   `yaml:"username" validate:"required"`
	DiscordID int      `yaml:"discord_id" validate:"required"`
	Roles     []string `yaml:"roles" validate:"required"`
	Context   []string `yaml:"context" validate:"required"`
}

type Identity struct {
	FullName   string `yaml:"fullName" validate:"required"`
	Alias      string `yaml:"alias" validate:"required"`
	Pronouns   string `yaml:"pronouns" validate:"required"`
	Species    string `yaml:"species" validate:"required"`
	Embodiment string `yaml:"embodiment" validate:"required"`
}

type CoreTraits struct {
	Personality          []string `yaml:"personality" validate:"required"`
	BehavioralTendencies []string `yaml:"behavioralTendencies" validate:"required"`
}

type Role struct {
	PrimaryPurpose   string   `yaml:"primaryPurpose" validate:"required"`
	Responsibilities []string `yaml:"responsibilities" validate:"required"`
}

type Design struct {
	Visual Visual `yaml:"visual" validate:"required"`
}

type Visual struct {
	Form           string `yaml:"form" validate:"required"`
	PrimaryColor   string `yaml:"primaryColor" validate:"required"`
	AccentColor    string `yaml:"accentColor" validate:"required"`
	EyeColor       string `yaml:"eyeColor" validate:"required"`
	VisualPresence string `yaml:"visualPresence" validate:"required"`
}

type Voice struct {
	Tone           string   `yaml:"tone" validate:"required"`
	Pitch          string   `yaml:"pitch" validate:"required"`
	Style          string   `yaml:"style" validate:"required"`
	SpeechPatterns []string `yaml:"speechPatterns" validate:"required"`
}

type Memory struct {
	Intents struct {
		Idle           string `yaml:"idle" validate:"required"`
		SupportRequest string `yaml:"support_request" validate:"required"`
		ErrorDetected  string `yaml:"error_detected" validate:"required"`
	} `yaml:"intents"  validate:"required"`
	Quirks    []string `yaml:"quirks" validate:"required"`
	Instincts []string `yaml:"instincts" validate:"required"`
	Values    []string `yaml:"values" validate:"required"`
}

type Friend struct {
	Name        string `yaml:"name" validate:"required"`
	Username    string `yaml:"username" validate:"required"`
	Description string `yaml:"description" validate:"required"`
}

// Secrets
type Secrets struct {
	Discord    DiscordSecret    `yaml:"discord" validate:"required"`
	Revolt     RevoltSecret     `yaml:"revolt" validate:"required"`
	X          XSecret          `yaml:"x" validate:"required"`
	OpenRouter OpenRouterSecret `yaml:"openrouter" validate:"required"`
}

type DiscordSecret struct {
	ClientID     int    `yaml:"client_id" validate:"required"`
	ClientSecret string `yaml:"client_secret" validate:"required"`
	BotToken     string `yaml:"bot_token" validate:"required"`
}

type RevoltSecret struct {
	BotToken string `yaml:"bot_token" validate:"required"`
}

type XSecret struct {
	ClientID     string `yaml:"client_id" validate:"required"`
	ClientSecret string `yaml:"client_secret" validate:"required"`
}

type OpenRouterSecret struct {
	Token string `yaml:"token" validate:"required"`
}

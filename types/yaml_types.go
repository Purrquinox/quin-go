package types

// AI Data
type AIData struct {
	Name        string    `json:"name"`
	Version     int       `json:"version"`
	ApiPort     int       `json:"api_port"`
	Staff       []Staff   `json:"staff"`
	Identity    Identity  `json:"identity"`
	Summary     string    `json:"summary"`
	CoreTraits  CoreTraits `json:"coreTraits"`
	Role        Role      `json:"role"`
	Design      Design    `json:"design"`
	Voice       Voice     `json:"voice"`
	Memory      Memory    `json:"memory"`
	Friends     []Friend  `json:"friends"`
}

type Staff struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	DiscordID int     `json:"discord_id"`
	Roles    []string `json:"roles"`
	Context  []string `json:"context"`
}

type Identity struct {
	FullName    string `json:"fullName"`
	Alias       string `json:"alias"`
	Pronouns    string `json:"pronouns"`
	Species     string `json:"species"`
	Embodiment  string `json:"embodiment"`
}

type CoreTraits struct {
	Personality         []string `json:"personality"`
	BehavioralTendencies []string `json:"behavioralTendencies"`
}

type Role struct {
	PrimaryPurpose  string   `json:"primaryPurpose"`
	Responsibilities []string `json:"responsibilities"`
}

type Design struct {
	Visual Visual `json:"visual"`
}

type Visual struct {
	Form            string `json:"form"`
	PrimaryColor    string `json:"primaryColor"`
	AccentColor     string `json:"accentColor"`
	EyeColor        string `json:"eyeColor"`
	VisualPresence   string `json:"visualPresence"`
}

type Voice struct {
	Tone          string   `json:"tone"`
	Pitch         string   `json:"pitch"`
	Style         string   `json:"style"`
	SpeechPatterns []string `json:"speechPatterns"`
}

type Memory struct {
	Intents struct {
		Idle            string `json:"idle"`
		SupportRequest  string `json:"support_request"`
		ErrorDetected   string `json:"error_detected"`
	} `json:"intents"`
	Quirks   []string `json:"quirks"`
	Instincts []string `json:"instincts"`
	Values    []string `json:"values"`
}

type Friend struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

// Secrets 
type Secrets struct {
	Discord      DiscordSecret   `json:"discord"`
	Revolt       RevoltSecret    `json:"revolt"`
	X            XSecret         `json:"x"`
	OpenRouter   OpenRouterSecret `json:"openrouter"`
}

type DiscordSecret struct {
	ClientID     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	BotToken     string `json:"bot_token"`
}

type RevoltSecret struct {
	BotToken string `json:"bot_token"`
}

type XSecret struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OpenRouterSecret struct {
	Token string `json:"token"`
}

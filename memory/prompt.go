package memory

import (
	"fmt"
	"quin/state"
	"strings"
)

// Personality returns a formatted string describing the current personality state.
func Personality() string {
	return fmt.Sprintf(`You are %s, also known as "%s".

Species: %s  
Pronouns: %s  
Embodiment: %s

Summary: %s

🎭 Personality Traits:
%s

📈 Behavioral Tendencies:
%s

🧠 Instincts:
%s

🎨 Visual Design:
- Form: %s
- Primary Color: %s
- Accent Color: %s
- Eye Color: %s
- Presence: %s

🗣️ Voice:
- Tone: %s
- Pitch: %s
- Style: %s
- Patterns: %s

🎯 Purpose:
- %s
- Responsibilities: %s

⚙️ Quirks: %s
🎵 Idle Behavior: %s
🧩 Support Behavior: %s
🚨 Error Response: %s

💖 Values: %s

👥 Staff:
%s

🧑‍🤝‍🧑 Friends:
%s
`,
		state.Data.Identity.FullName,
		state.Data.Identity.Alias,
		state.Data.Identity.Species,
		state.Data.Identity.Pronouns,
		state.Data.Identity.Embodiment,
		state.Data.Summary,
		strings.Join(state.Data.CoreTraits.Personality, ", "),
		strings.Join(state.Data.CoreTraits.BehavioralTendencies, ", "),
		strings.Join(state.Data.Memory.Instincts, ", "),
		state.Data.Design.Visual.Form,
		state.Data.Design.Visual.PrimaryColor,
		state.Data.Design.Visual.AccentColor,
		state.Data.Design.Visual.EyeColor,
		state.Data.Design.Visual.VisualPresence,
		state.Data.Voice.Tone,
		state.Data.Voice.Pitch,
		state.Data.Voice.Style,
		strings.Join(state.Data.Voice.SpeechPatterns, "; "),
		state.Data.Role.PrimaryPurpose,
		strings.Join(state.Data.Role.Responsibilities, "; "),
		strings.Join(state.Data.Memory.Quirks, ", "),
		state.Data.Memory.Intents.Idle,
		state.Data.Memory.Intents.SupportRequest,
		state.Data.Memory.Intents.ErrorDetected,
		strings.Join(state.Data.Memory.Values, ", "),
		func() string {
			var out []string
			for _, s := range state.Data.Staff {
				out = append(out, fmt.Sprintf("- %s (%s): Roles: %s | Context: %s", s.Name, s.Username, strings.Join(s.Roles, ", "), strings.Join(s.Context, ", ")))
			}
			return strings.Join(out, "\n")
		}(),
		func() string {
			var out []string
			for _, f := range state.Data.Friends {
				out = append(out, fmt.Sprintf("- %s (%s): %s", f.Name, f.Username, f.Description))
			}
			return strings.Join(out, "\n")
		}(),
	)
}

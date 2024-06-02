package info

import "testing"

func stringPtr(s string) *string {
	return &s
}

func TestValidateLinks(t *testing.T) {
	t.Run("No validation for less than 2 links", func(t *testing.T) {
		links := []Link{
			{Name: stringPtr("twitter"), URL: stringPtr("https://twitter.com/example")},
		}
		err := ValidateLinks(links)
		if err != nil {
			t.Errorf("Expected no validation (and thus no error) for fewer than 2 links, got %v", err)
		}
	})

	t.Run("Invalid prefix for URL", func(t *testing.T) {
		links := []Link{
			{Name: stringPtr("discord"), URL: stringPtr("http://twitter.com/example")},
			{Name: stringPtr("discord"), URL: stringPtr("http//discord.io/example")},
		}
		err := ValidateLinks(links)
		println(err.Error())
		if err == nil {
			t.Errorf("Expected error for invalid URL prefix, got nil")
		}
	})

	t.Run("Invalid link name", func(t *testing.T) {
		links := []Link{
			{Name: stringPtr("unknownLink"), URL: stringPtr("https://unknown.com")},
			{Name: stringPtr("twitter"), URL: stringPtr("https://twitter.com/example")},
		}
		err := ValidateLinks(links)
		if err == nil {
			t.Errorf("Expected error for invalid link name, got nil")
		}
	})

	t.Run("URL without https:// prefix", func(t *testing.T) {
		links := []Link{
			{Name: stringPtr("twitter"), URL: stringPtr("twitter.com/example")},
			{Name: stringPtr("twitter"), URL: stringPtr("https://twitter.com/example")},
		}
		err := ValidateLinks(links)
		if err == nil {
			t.Errorf("Expected error for missing https:// prefix, got nil")
		}
	})

	var allowedLinkKeys = map[string][]string{
		"twitter":       {"https://twitter.com/example"},
		"medium":        {"https://medium.com/@example"},
		"telegram":      {"https://t.me/example"},
		"github":        {"https://github.com/example"},
		"whitepaper":    {"https://somewebsite.com/whitepaper"},
		"telegram_news": {"https://t.me/example"},
		"discord":       {"https://discord.com/example", "https://discord.gg/example"},
		"reddit":        {"https://reddit.com/example"},
		"facebook":      {"https://facebook.com/example"},
		"youtube":       {"https://youtube.com/example"},
		"coinmarketcap": {"https://coinmarketcap.com/example"},
		"coingecko":     {"https://coingecko.com/example"},
		"blog":          {"https://example.com"},
		"forum":         {"https://example.com"},
		"docs":          {"https://example.com"},
		"source_code":   {"https://example.com"},
	}

	for key, domains := range allowedLinkKeys {
		for _, domain := range domains {
			t.Run("Valid link for "+key+" with domain "+domain, func(t *testing.T) {
				links := []Link{
					{Name: stringPtr(key), URL: stringPtr(domain)},
					{Name: stringPtr("twitter"), URL: stringPtr("https://twitter.com/example")}, // adding a second link to meet the 2 links requirement
				}
				err := ValidateLinks(links)
				if err != nil {
					t.Errorf("Expected no error for valid link key %s with domain %s, got %v", key, domain, err)
				}
			})
		}
	}
}

func main() {
	t := &testing.T{}
	TestValidateLinks(t)
}

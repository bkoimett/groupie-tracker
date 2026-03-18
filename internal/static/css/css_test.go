package css

import (
	"os"
	"strings"
	"testing"
)

func TestCSSMediaQueries(t *testing.T) {
	// Read CSS file
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for responsive design
	expectedQueries := []string{
		"@media",
		"max-width: 768px",
		"grid-template-columns",
	}

	for _, query := range expectedQueries {
		if !strings.Contains(cssContent, query) {
			t.Errorf("CSS missing expected media query or property: %q", query)
		}
	}
}

func TestCSSClasses(t *testing.T) {
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for required CSS classes from templates - UPDATED to match new design
	expectedClasses := []string{
		".artist-grid",
		".artist-card",
		".artist-image",
		".artist-info",
		".details-link",
		".error-container", // Changed from .error-page to .error-container
		".home-button",
		".theme-toggle", // Added new class
		".theme-btn",    // Added new class
	}

	for _, class := range expectedClasses {
		if !strings.Contains(cssContent, class) {
			t.Errorf("CSS missing expected class: %q", class)
		}
	}
}

// Add test for theme toggle
func TestThemeClasses(t *testing.T) {
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for theme-related classes
	themeClasses := []string{
		"[data-theme=\"dark\"]",
		"--accent-primary",
		"--bg-primary",
		"--text-primary",
	}

	for _, class := range themeClasses {
		if !strings.Contains(cssContent, class) {
			t.Errorf("CSS missing theme class/variable: %q", class)
		}
	}
}

// Test responsive breakpoints
func TestResponsiveBreakpoints(t *testing.T) {
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for mobile breakpoint
	if !strings.Contains(cssContent, "@media (max-width: 768px)") {
		t.Error("CSS missing mobile breakpoint @media (max-width: 768px)")
	}

	// Check for small mobile breakpoint
	if !strings.Contains(cssContent, "@media (max-width: 480px)") {
		t.Error("CSS missing small mobile breakpoint @media (max-width: 480px)")
	}
}

// Test CSS variables
func TestCSSVariables(t *testing.T) {
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for CSS custom properties
	requiredVars := []string{
		"--bg-primary",
		"--text-primary",
		"--accent-primary",
		"--border-color",
	}

	for _, v := range requiredVars {
		if !strings.Contains(cssContent, v) {
			t.Errorf("CSS missing required variable: %s", v)
		}
	}
}

// Test animations
func TestAnimations(t *testing.T) {
	css, err := os.ReadFile("style.css")
	if err != nil {
		t.Skip("CSS file not found, skipping test")
	}

	cssContent := string(css)

	// Check for animations
	if !strings.Contains(cssContent, "@keyframes fadeIn") {
		t.Error("CSS missing fadeIn animation")
	}

	if !strings.Contains(cssContent, "transition:") {
		t.Error("CSS missing transitions")
	}
}

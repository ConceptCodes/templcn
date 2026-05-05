package ui

import (
	"bytes"
	"context"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func renderToString(c templ.Component) (string, error) {
	var buf bytes.Buffer
	err := c.Render(context.Background(), &buf)
	return buf.String(), err
}

func mustRender(t *testing.T, c templ.Component) string {
	t.Helper()
	s, err := renderToString(c)
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	return s
}

func attrValue(html, attr string) string {
	pattern := regexp.MustCompile(attr + `="([^"]*)"`)
	matches := pattern.FindStringSubmatch(html)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}

func hasAttr(html, attr string) bool {
	pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(attr) + `(?:=|\s|\/|>)`)
	return pattern.MatchString(html)
}

func hasClass(html, class string) bool {
	cls := attrValue(html, "class")
	for _, c := range strings.Split(cls, " ") {
		if c == class {
			return true
		}
	}
	return false
}

func hasAllClasses(html string, classes ...string) bool {
	for _, c := range classes {
		if !hasClass(html, c) {
			return false
		}
	}
	return true
}

func classContains(html, substr string) bool {
	cls := attrValue(html, "class")
	return strings.Contains(cls, substr)
}

func htmlEscapeClass(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}

func extractTag(html string) string {
	re := regexp.MustCompile(`<(\w+)`)
	matches := re.FindStringSubmatch(html)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}

func TestCN(t *testing.T) {
	tests := []struct {
		inputs []string
		want   string
	}{
		{[]string{"a", "b", "c"}, "a b c"},
		{[]string{"a", "", "b"}, "a b"},
		{[]string{"", ""}, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "  b  "}, "a b"},
	}
	for _, tt := range tests {
		got := cn(tt.inputs...)
		if got != tt.want {
			t.Errorf("cn(%v) = %q, want %q", tt.inputs, got, tt.want)
		}
	}
}

func TestCloneAttributes(t *testing.T) {
	original := templ.Attributes{"class": "foo", "id": "bar"}
	cloned := cloneAttributes(original)
	if cloned["class"] != original["class"] {
		t.Error("cloneAttributes did not copy class")
	}
	cloned["class"] = "baz"
	if original["class"] == "baz" {
		t.Error("cloneAttributes did not deep-copy, original was mutated")
	}
}

func TestAttrsFromDOMProps_SetsDataSlot(t *testing.T) {
	attrs := attrsFromDOMProps(DOMProps{}, "test-slot", "base-class")
	if attrs["data-slot"] != "test-slot" {
		t.Errorf("expected data-slot=test-slot, got %v", attrs["data-slot"])
	}
}

func TestAttrsFromDOMProps_SetsID(t *testing.T) {
	attrs := attrsFromDOMProps(DOMProps{ID: "my-id"}, "", "")
	if attrs["id"] != "my-id" {
		t.Errorf("expected id=my-id, got %v", attrs["id"])
	}
}

func TestAttrsFromDOMProps_MergesClass(t *testing.T) {
	attrs := attrsFromDOMProps(DOMProps{Class: "extra"}, "slot", "base")
	cls := attrs["class"].(string)
	if !strings.Contains(cls, "base") || !strings.Contains(cls, "extra") {
		t.Errorf("expected merged classes, got %q", cls)
	}
}

func TestRenderElement(t *testing.T) {
	var buf bytes.Buffer
	attrs := templ.Attributes{"class": "test", "data-slot": "el"}
	children := templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "hello")
		return err
	})
	err := renderElement(context.Background(), &buf, "div", attrs, children)
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if !strings.Contains(s, "<div") {
		t.Errorf("expected <div, got %q", s)
	}
	if !strings.Contains(s, "</div>") {
		t.Errorf("expected </div>, got %q", s)
	}
	if !strings.Contains(s, "hello") {
		t.Errorf("expected hello, got %q", s)
	}
}

func TestRenderVoidElement(t *testing.T) {
	var buf bytes.Buffer
	attrs := templ.Attributes{"class": "test"}
	err := renderVoidElement(context.Background(), &buf, "hr", attrs)
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if !strings.Contains(s, "<hr") {
		t.Errorf("expected <hr, got %q", s)
	}
	if !strings.Contains(s, " />") {
		t.Errorf("expected void element closing, got %q", s)
	}
	if strings.Contains(s, "</hr>") {
		t.Errorf("void element should not have closing tag, got %q", s)
	}
}

func TestRenderTextElement(t *testing.T) {
	var buf bytes.Buffer
	attrs := templ.Attributes{"class": "test"}
	err := renderTextElement(context.Background(), &buf, "span", attrs, "hello & world")
	if err != nil {
		t.Fatal(err)
	}
	s := buf.String()
	if !strings.Contains(s, "hello &amp; world") {
		t.Errorf("expected escaped text, got %q", s)
	}
}

func TestButton_DefaultVariant(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{}))
	tag := extractTag(html)
	if tag != "button" {
		t.Errorf("expected <button>, got <%s>", tag)
	}
	if !hasAttr(html, "data-slot") {
		t.Error("Button missing data-slot attribute")
	}
	if attrValue(html, "data-slot") != "button" {
		t.Errorf("expected data-slot=button, got %q", attrValue(html, "data-slot"))
	}
}

func TestButton_AllVariants(t *testing.T) {
	variants := []ButtonVariant{
		ButtonVariantDefault,
		ButtonVariantDestructive,
		ButtonVariantOutline,
		ButtonVariantSecondary,
		ButtonVariantGhost,
		ButtonVariantLink,
	}
	for _, v := range variants {
		html := mustRender(t, Button(ButtonProps{Variant: v}))
		cls := attrValue(html, "class")
		expectedClasses := buttonVariantClasses[v]
		if !strings.Contains(cls, expectedClasses) {
			t.Errorf("variant %s: expected class %q in %q", v, expectedClasses, cls)
		}
	}
}

func TestButton_AllSizes(t *testing.T) {
	sizes := []ButtonSize{
		ButtonSizeDefault,
		ButtonSizeXS,
		ButtonSizeSM,
		ButtonSizeLG,
		ButtonSizeIcon,
		ButtonSizeIconXS,
		ButtonSizeIconSM,
		ButtonSizeIconLG,
	}
	for _, s := range sizes {
		html := mustRender(t, Button(ButtonProps{Size: s}))
		cls := attrValue(html, "class")
		for _, fragment := range strings.Split(buttonSizeClasses[s], " ") {
			if !strings.Contains(cls, fragment) && !strings.Contains(cls, htmlEscapeClass(fragment)) {
				t.Errorf("size %s: expected class fragment %q in %q", s, fragment, cls)
			}
		}
	}
}

func TestButton_Disabled(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{Disabled: true}))
	if !hasAttr(html, "disabled") {
		t.Error("disabled button missing disabled attribute")
	}
}

func TestButton_Type(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{}))
	if attrValue(html, "type") != "button" {
		t.Error("default button type should be 'button'")
	}
	html = mustRender(t, Button(ButtonProps{Type: "submit"}))
	if attrValue(html, "type") != "submit" {
		t.Error("submit button type should be 'submit'")
	}
}

func TestButton_Href(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{Href: "/path"}))
	tag := extractTag(html)
	if tag != "a" {
		t.Errorf("button with href should render <a>, got <%s>", tag)
	}
	if attrValue(html, "href") != "/path" {
		t.Error("anchor button missing href")
	}
	if hasAttr(html, "disabled") {
		t.Error("anchor button should not have disabled attribute")
	}
}

func TestButton_Label(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{Label: "Click me"}))
	if !strings.Contains(html, "Click me") {
		t.Error("button should contain label text")
	}
}

func TestButton_CustomElement(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{DOMProps: DOMProps{Element: "span"}}))
	tag := extractTag(html)
	if tag != "span" {
		t.Errorf("expected <span>, got <%s>", tag)
	}
}

func TestButton_ClassOverride(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{DOMProps: DOMProps{Class: "my-custom"}}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "my-custom") {
		t.Errorf("custom class should be appended, got %q", cls)
	}
	if !strings.Contains(cls, "bg-primary") {
		t.Errorf("base variant classes should remain, got %q", cls)
	}
}

func TestButton_DataAttributes(t *testing.T) {
	html := mustRender(t, Button(ButtonProps{Variant: ButtonVariantDestructive, Size: ButtonSizeLG}))
	if attrValue(html, "data-variant") != "destructive" {
		t.Errorf("expected data-variant=destructive, got %q", attrValue(html, "data-variant"))
	}
	if attrValue(html, "data-size") != "lg" {
		t.Errorf("expected data-size=lg, got %q", attrValue(html, "data-size"))
	}
}

func TestBadge_DefaultVariant(t *testing.T) {
	html := mustRender(t, Badge(BadgeProps{}))
	tag := extractTag(html)
	if tag != "span" {
		t.Errorf("expected <span>, got <%s>", tag)
	}
}

func TestBadge_AllVariants(t *testing.T) {
	variants := []BadgeVariant{
		BadgeVariantDefault,
		BadgeVariantSecondary,
		BadgeVariantDestructive,
		BadgeVariantOutline,
		BadgeVariantGhost,
		BadgeVariantLink,
	}
	for _, v := range variants {
		html := mustRender(t, Badge(BadgeProps{Variant: v}))
		cls := attrValue(html, "class")
		for _, fragment := range strings.Split(badgeVariantClasses[v], " ") {
			if !strings.Contains(cls, fragment) && !strings.Contains(cls, htmlEscapeClass(fragment)) {
				t.Errorf("badge variant %s: expected class fragment %q in %q", v, fragment, cls)
			}
		}
	}
}

func TestBadge_Href(t *testing.T) {
	html := mustRender(t, Badge(BadgeProps{Href: "/link"}))
	tag := extractTag(html)
	if tag != "a" {
		t.Errorf("badge with href should render <a>, got <%s>", tag)
	}
}

func TestBadge_DataSlot(t *testing.T) {
	html := mustRender(t, Badge(BadgeProps{}))
	if attrValue(html, "data-slot") != "badge" {
		t.Errorf("expected data-slot=badge, got %q", attrValue(html, "data-slot"))
	}
}

func TestBadge_DataVariant(t *testing.T) {
	html := mustRender(t, Badge(BadgeProps{Variant: BadgeVariantOutline}))
	if attrValue(html, "data-variant") != "outline" {
		t.Errorf("expected data-variant=outline, got %q", attrValue(html, "data-variant"))
	}
}

func TestCard_Root(t *testing.T) {
	html := mustRender(t, Card(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("expected <div>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "card" {
		t.Error("Card missing data-slot=card")
	}
}

func TestCard_RootClasses(t *testing.T) {
	html := mustRender(t, Card(DOMProps{}))
	required := []string{"flex", "flex-col", "gap-6", "rounded-xl", "border", "bg-card", "py-6", "text-card-foreground", "shadow-sm"}
	for _, c := range required {
		if !hasClass(html, c) {
			t.Errorf("Card missing required class %q", c)
		}
	}
}

func TestCard_HeaderClasses(t *testing.T) {
	html := mustRender(t, CardHeader(DOMProps{}))
	if attrValue(html, "data-slot") != "card-header" {
		t.Error("CardHeader missing data-slot")
	}
}

func TestCard_Title(t *testing.T) {
	html := mustRender(t, CardTitle(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("CardTitle should render <div>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "card-title" {
		t.Error("CardTitle missing data-slot")
	}
}

func TestCard_Description(t *testing.T) {
	html := mustRender(t, CardDescription(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("CardDescription should render <div>, got <%s>", tag)
	}
}

func TestCard_Action(t *testing.T) {
	html := mustRender(t, CardAction(DOMProps{}))
	if attrValue(html, "data-slot") != "card-action" {
		t.Error("CardAction missing data-slot")
	}
}

func TestCard_Content(t *testing.T) {
	html := mustRender(t, CardContent(DOMProps{}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "px-6") {
		t.Errorf("CardContent should have px-6, got %q", cls)
	}
}

func TestCard_Footer(t *testing.T) {
	html := mustRender(t, CardFooter(DOMProps{}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "flex") || !strings.Contains(cls, "px-6") {
		t.Errorf("CardFooter missing expected classes, got %q", cls)
	}
}

func TestAlert_DefaultVariant(t *testing.T) {
	html := mustRender(t, Alert(AlertProps{}))
	if !hasAttr(html, "role") || attrValue(html, "role") != "alert" {
		t.Error("Alert should have role=alert")
	}
	if attrValue(html, "data-slot") != "alert" {
		t.Error("Alert missing data-slot=alert")
	}
}

func TestAlert_DestructiveVariant(t *testing.T) {
	html := mustRender(t, Alert(AlertProps{Variant: AlertVariantDestructive}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "text-destructive") {
		t.Errorf("destructive alert should contain text-destructive, got %q", cls)
	}
}

func TestAlert_Title(t *testing.T) {
	html := mustRender(t, AlertTitle(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("AlertTitle should render <div>, got <%s>", tag)
	}
}

func TestAlert_Description(t *testing.T) {
	html := mustRender(t, AlertDescription(DOMProps{}))
	if attrValue(html, "data-slot") != "alert-description" {
		t.Error("AlertDescription missing data-slot")
	}
}

func TestSeparator_Default(t *testing.T) {
	html := mustRender(t, Separator(SeparatorProps{}))
	tag := extractTag(html)
	if tag != "hr" {
		t.Errorf("expected <hr>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "separator" {
		t.Error("Separator missing data-slot")
	}
}

func TestSeparator_Horizontal(t *testing.T) {
	html := mustRender(t, Separator(SeparatorProps{Orientation: SeparatorOrientationHorizontal}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "h-px") || !strings.Contains(cls, "w-full") {
		t.Errorf("horizontal separator should have h-px w-full, got %q", cls)
	}
}

func TestSeparator_Vertical(t *testing.T) {
	html := mustRender(t, Separator(SeparatorProps{Orientation: SeparatorOrientationVertical}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "h-full") || !strings.Contains(cls, "w-px") {
		t.Errorf("vertical separator should have h-full w-px, got %q", cls)
	}
}

func TestSeparator_Decorative(t *testing.T) {
	html := mustRender(t, Separator(SeparatorProps{Decorative: true}))
	if !hasAttr(html, "aria-hidden") {
		t.Error("decorative separator should have aria-hidden")
	}
}

func TestSeparator_NonDecorative(t *testing.T) {
	html := mustRender(t, Separator(SeparatorProps{Decorative: false}))
	if attrValue(html, "role") != "separator" {
		t.Error("non-decorative separator should have role=separator")
	}
}

func TestSkeleton(t *testing.T) {
	html := mustRender(t, Skeleton(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("expected <div>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "skeleton" {
		t.Error("Skeleton missing data-slot")
	}
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "animate-pulse") {
		t.Errorf("Skeleton should have animate-pulse, got %q", cls)
	}
	if !strings.Contains(cls, "bg-accent") {
		t.Errorf("Skeleton should have bg-accent (not bg-muted), got %q", cls)
	}
}

func TestSpinner_Default(t *testing.T) {
	html := mustRender(t, Spinner(SpinnerProps{}))
	if attrValue(html, "data-slot") != "spinner" {
		t.Error("Spinner missing data-slot")
	}
	if !hasAttr(html, "aria-hidden") {
		t.Error("Spinner without label should have aria-hidden")
	}
}

func TestSpinner_WithLabel(t *testing.T) {
	html := mustRender(t, Spinner(SpinnerProps{Label: "Loading"}))
	if attrValue(html, "role") != "status" {
		t.Error("Spinner with label should have role=status")
	}
	if attrValue(html, "aria-label") != "Loading" {
		t.Error("Spinner should have aria-label matching Label prop")
	}
}

func TestSpinner_Sizes(t *testing.T) {
	sizes := map[SpinnerSize]string{
		SpinnerSizeSM: "size-4",
		SpinnerSizeMD: "size-5",
		SpinnerSizeLG: "size-6",
	}
	for size, expected := range sizes {
		html := mustRender(t, Spinner(SpinnerProps{Size: size}))
		cls := attrValue(html, "class")
		if !strings.Contains(cls, expected) {
			t.Errorf("Spinner size %s: expected %q in classes, got %q", size, expected, cls)
		}
	}
}

func TestKbd(t *testing.T) {
	html := mustRender(t, Kbd(KbdProps{Text: "Ctrl"}))
	tag := extractTag(html)
	if tag != "kbd" {
		t.Errorf("expected <kbd>, got <%s>", tag)
	}
	if !strings.Contains(html, "Ctrl") {
		t.Error("Kbd should contain text")
	}
	if attrValue(html, "data-slot") != "kbd" {
		t.Error("Kbd missing data-slot")
	}
}

func TestKbd_Sizes(t *testing.T) {
	tests := []struct {
		size     string
		expected string
	}{
		{"sm", "text-[0.65rem]"},
		{"", "text-[0.7rem]"},
		{"lg", "text-[0.75rem]"},
	}
	for _, tt := range tests {
		html := mustRender(t, Kbd(KbdProps{Text: "K", Size: tt.size}))
		cls := attrValue(html, "class")
		if !strings.Contains(cls, tt.expected) {
			t.Errorf("Kbd size %q: expected %q in %q", tt.size, tt.expected, cls)
		}
	}
}

func TestLabel(t *testing.T) {
	html := mustRender(t, Label(LabelProps{}))
	tag := extractTag(html)
	if tag != "label" {
		t.Errorf("expected <label>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "label" {
		t.Error("Label missing data-slot")
	}
}

func TestLabel_For(t *testing.T) {
	html := mustRender(t, Label(LabelProps{For: "input-id"}))
	if attrValue(html, "for") != "input-id" {
		t.Error("Label should set for attribute")
	}
}

func TestLabel_Classes(t *testing.T) {
	html := mustRender(t, Label(LabelProps{}))
	cls := attrValue(html, "class")
	required := []string{"flex", "items-center", "gap-2", "text-sm", "leading-none", "font-medium", "select-none", "peer-disabled:opacity-50"}
	for _, c := range required {
		if !strings.Contains(cls, c) {
			t.Errorf("Label missing required class %q in %q", c, cls)
		}
	}
}

func TestProgress(t *testing.T) {
	html := mustRender(t, Progress(ProgressProps{Value: 50}))
	if attrValue(html, "data-slot") != "progress" {
		t.Error("Progress missing data-slot")
	}
	if attrValue(html, "role") != "progressbar" {
		t.Error("Progress should have role=progressbar")
	}
	if attrValue(html, "aria-valuemin") != "0" {
		t.Error("Progress should have aria-valuemin=0")
	}
	if attrValue(html, "aria-valuemax") != "100" {
		t.Error("Progress should have aria-valuemax=100")
	}
	if attrValue(html, "aria-valuenow") != "50" {
		t.Error("Progress should have aria-valuenow=50")
	}
}

func TestProgress_TrackClass(t *testing.T) {
	html := mustRender(t, Progress(ProgressProps{Value: 0}))
	cls := attrValue(html, "class")
	if !strings.Contains(cls, "bg-primary/20") {
		t.Errorf("Progress track should use bg-primary/20, got %q", cls)
	}
}

func TestProgress_Max(t *testing.T) {
	html := mustRender(t, Progress(ProgressProps{Value: 3, Max: 10}))
	if attrValue(html, "aria-valuemax") != "10" {
		t.Errorf("expected aria-valuemax=10, got %q", attrValue(html, "aria-valuemax"))
	}
}

func TestProgress_Clamp(t *testing.T) {
	html := mustRender(t, Progress(ProgressProps{Value: 150, Max: 100}))
	if attrValue(html, "aria-valuenow") != "100" {
		t.Error("Progress value should be clamped to max")
	}
	html = mustRender(t, Progress(ProgressProps{Value: -10, Max: 100}))
	if attrValue(html, "aria-valuenow") != "0" {
		t.Error("Progress value should be clamped to 0")
	}
}

func TestProgress_IndicatorDataSlot(t *testing.T) {
	html := mustRender(t, Progress(ProgressProps{Value: 50}))
	if !strings.Contains(html, `data-slot="progress-indicator"`) {
		t.Error("Progress indicator missing data-slot=progress-indicator")
	}
}

func TestTypography_H1(t *testing.T) {
	html := mustRender(t, H1(DOMProps{}))
	tag := extractTag(html)
	if tag != "h1" {
		t.Errorf("expected <h1>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "h1" {
		t.Error("H1 missing data-slot")
	}
}

func TestTypography_H2(t *testing.T) {
	html := mustRender(t, H2(DOMProps{}))
	tag := extractTag(html)
	if tag != "h2" {
		t.Errorf("expected <h2>, got <%s>", tag)
	}
}

func TestTypography_H3(t *testing.T) {
	html := mustRender(t, H3(DOMProps{}))
	tag := extractTag(html)
	if tag != "h3" {
		t.Errorf("expected <h3>, got <%s>", tag)
	}
}

func TestTypography_H4(t *testing.T) {
	html := mustRender(t, H4(DOMProps{}))
	tag := extractTag(html)
	if tag != "h4" {
		t.Errorf("expected <h4>, got <%s>", tag)
	}
}

func TestTypography_P(t *testing.T) {
	html := mustRender(t, P(DOMProps{}))
	tag := extractTag(html)
	if tag != "p" {
		t.Errorf("expected <p>, got <%s>", tag)
	}
}

func TestTypography_Blockquote(t *testing.T) {
	html := mustRender(t, Blockquote(DOMProps{}))
	tag := extractTag(html)
	if tag != "blockquote" {
		t.Errorf("expected <blockquote>, got <%s>", tag)
	}
}

func TestTypography_InlineCode(t *testing.T) {
	html := mustRender(t, InlineCode(DOMProps{}))
	tag := extractTag(html)
	if tag != "code" {
		t.Errorf("expected <code>, got <%s>", tag)
	}
}

func TestTypography_Lead(t *testing.T) {
	html := mustRender(t, Lead(DOMProps{}))
	tag := extractTag(html)
	if tag != "p" {
		t.Errorf("expected <p>, got <%s>", tag)
	}
	if !strings.Contains(attrValue(html, "class"), "text-muted-foreground") {
		t.Error("Lead should have text-muted-foreground")
	}
}

func TestTypography_Small(t *testing.T) {
	html := mustRender(t, Small(DOMProps{}))
	tag := extractTag(html)
	if tag != "small" {
		t.Errorf("expected <small>, got <%s>", tag)
	}
}

func TestTypography_List(t *testing.T) {
	html := mustRender(t, List(DOMProps{}))
	tag := extractTag(html)
	if tag != "ul" {
		t.Errorf("expected <ul>, got <%s>", tag)
	}
}

func TestTypography_TableProse(t *testing.T) {
	html := mustRender(t, TableProse(DOMProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("expected <div>, got <%s>", tag)
	}
}

func TestTypography_AllHaveDataSlots(t *testing.T) {
	components := map[string]templ.Component{
		"h1":             H1(DOMProps{}),
		"h2":             H2(DOMProps{}),
		"h3":             H3(DOMProps{}),
		"h4":             H4(DOMProps{}),
		"p":              P(DOMProps{}),
		"blockquote":     Blockquote(DOMProps{}),
		"inline-code":    InlineCode(DOMProps{}),
		"lead":           Lead(DOMProps{}),
		"large":          Large(DOMProps{}),
		"small":          Small(DOMProps{}),
		"muted":          Muted(DOMProps{}),
		"list":           List(DOMProps{}),
		"table-prose":    TableProse(DOMProps{}),
	}
	for slot, comp := range components {
		html := mustRender(t, comp)
		if attrValue(html, "data-slot") != slot {
			t.Errorf("expected data-slot=%q, got %q", slot, attrValue(html, "data-slot"))
		}
	}
}

func TestAspectRatio_Default(t *testing.T) {
	html := mustRender(t, AspectRatio(AspectRatioProps{}))
	tag := extractTag(html)
	if tag != "div" {
		t.Errorf("expected <div>, got <%s>", tag)
	}
	if attrValue(html, "data-slot") != "aspect-ratio" {
		t.Error("AspectRatio missing data-slot")
	}
}

func TestAspectRatio_Ratio(t *testing.T) {
	html := mustRender(t, AspectRatio(AspectRatioProps{Ratio: 16.0 / 9.0}))
	style := attrValue(html, "style")
	if !strings.Contains(style, "aspect-ratio:") {
		t.Errorf("expected aspect-ratio in style, got %q", style)
	}
}

func TestAspectRatio_DefaultRatio(t *testing.T) {
	html := mustRender(t, AspectRatio(AspectRatioProps{Ratio: 0}))
	style := attrValue(html, "style")
	if !strings.Contains(style, "aspect-ratio: 1") {
		t.Errorf("default ratio should be 1, got %q", style)
	}
}
